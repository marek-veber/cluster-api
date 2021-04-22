/*
Copyright 2021 The Kubernetes Authors.

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

package framework

import (
	"context"

	. "github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/cluster-api/controllers/noderefutil"
)

// WaitForNodesReadyInput is the input for WaitForNodesReady.
type WaitForNodesReadyInput struct {
	Lister            Lister
	KubernetesVersion string
	Count             int
	WaitForNodesReady []interface{}
}

// WaitForNodesReady waits until all nodes match with the Kubernetes version and are ready.
func WaitForNodesReady(ctx context.Context, input WaitForNodesReadyInput) {
	Eventually(func() (bool, error) {
		nodeList := &corev1.NodeList{}
		if err := input.Lister.List(ctx, nodeList); err != nil {
			return false, err
		}
		nodeReadyCount := 0
		for _, node := range nodeList.Items {
			n := node
			if node.Status.NodeInfo.KubeletVersion == input.KubernetesVersion && noderefutil.IsNodeReady(&n) {
				nodeReadyCount++
			}
		}
		return input.Count == nodeReadyCount, nil
	}, input.WaitForNodesReady...).Should(BeTrue())
}
