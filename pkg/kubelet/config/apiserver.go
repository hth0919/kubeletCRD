/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	corev1 "k8s.io/api/core/v1"
	ketiv1 "k8s.io/api/keti/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
	kubetypes "k8s.io/kubernetes/pkg/kubelet/types"
)

// NewSourceApiserver creates a config source that watches and pulls from the apiserver.
func NewSourceApiserver(c clientset.Interface, nodeName types.NodeName, updates chan<- interface{}) {
	lw := cache.NewListWatchFromClient(c.KetiV1().RESTClient(), "pods", metav1.NamespaceAll, fields.Everything())
	klog.Infoln(lw.List(metav1.ListOptions{}))
	newSourceApiserverFromLW(c, lw, nodeName, updates)
}

// newSourceApiserverFromLW holds creates a config source that watches and pulls from the apiserver.
func newSourceApiserverFromLW(c clientset.Interface,lw cache.ListerWatcher, nodeName types.NodeName, updates chan<- interface{}) {
	send := func(objs []interface{}) {
		var pods []*corev1.Pod
		var pod *corev1.Pod
		for _, o := range objs {
			pod = (*corev1.Pod)(o.(*ketiv1.Pod))
			pod.APIVersion = "v1"
			klog.Infoln("BEFORE : ", pod.Spec)
			pod.Spec = podConstructor(c, o.(*ketiv1.Pod))
			klog.Infoln("AFTER : ", pod.Spec)
			if pod.Spec.NodeName == string(nodeName) {
				pods = append(pods, pod)
			}
		}
		updates <- kubetypes.PodUpdate{Pods: pods, Op: kubetypes.SET, Source: kubetypes.ApiserverSource}
	}
	r := cache.NewReflector(lw, &ketiv1.Pod{}, cache.NewUndeltaStore(send, cache.MetaNamespaceKeyFunc), 0)
	go r.Run(wait.NeverStop)
}

func podConstructor(c clientset.Interface, pod *ketiv1.Pod) corev1.PodSpec{
	corepod := &corev1.Pod{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                       pod.Name,
			Namespace:                  pod.Namespace,
		},
		Spec:       corev1.PodSpec{},
		Status:     corev1.PodStatus{},
	}
	corepodSpec := pod.Spec
	corepod.Spec = corepodSpec
	klog.Infoln(corepod)
	corepod, err := c.CoreV1().Pods(pod.Namespace).Create(corepod)
	if err != nil {
		klog.Infoln(err)
	}
	result := corepod.Spec
	err = c.CoreV1().Pods(pod.Namespace).Delete(pod.Name, &metav1.DeleteOptions{})
	if err != nil {
		klog.Infoln(err)
	}
	return result
}
