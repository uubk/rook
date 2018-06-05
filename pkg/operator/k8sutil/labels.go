/*
Copyright 2018 The Rook Authors. All rights reserved.

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

// Package k8sutil for Kubernetes helpers.
package k8sutil

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
)

// Try to determine current node and list all node labels for it
func GetCurrentNodeLabels(clientset kubernetes.Interface) map[string]string {
	hostname := os.Getenv("ROOK_NODE_NAME")
	logger.Infof("Node name for node labels: %s", hostname)
	node, err := clientset.CoreV1().Nodes().Get(hostname, metav1.GetOptions{})
	if err != nil {
		logger.Warningf("Couldn't get labels of current node! (%s)", err)
		return make(map[string]string)
	}
	return node.Labels
}
