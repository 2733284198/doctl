/*
Copyright 2018 The Doctl Authors All rights reserved.
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

package do

import (
	"context"

	"github.com/digitalocean/godo"
)

// KubernetesCluster wraps a godo KubernetesCluster.
type KubernetesCluster struct {
	*godo.KubernetesCluster
}

// KubernetesClusters is a slice of KubernetesCluster.
type KubernetesClusters []KubernetesCluster

// KubernetesNodePool wraps a godo KubernetesNodePool.
type KubernetesNodePool struct {
	*godo.KubernetesNodePool
}

// KubernetesNodePools is a slice of KubernetesNodePool.
type KubernetesNodePools []KubernetesNodePool

// KubernetesService is the godo KubernetesService interface.
type KubernetesService interface {
	Get(clusterID string) (*KubernetesCluster, error)
	GetKubeConfig(clusterID string) ([]byte, error)
	List() (KubernetesClusters, error)
	Create(create *godo.KubernetesClusterCreateRequest) (*KubernetesCluster, error)
	Update(clusterID string, update *godo.KubernetesClusterUpdateRequest) (*KubernetesCluster, error)
	Delete(clusterID string) error

	CreateNodePool(clusterID string, req *godo.KubernetesNodePoolCreateRequest) (*KubernetesNodePool, error)
	GetNodePool(clusterID, poolID string) (*KubernetesNodePool, error)
	ListNodePools(clusterID string, opts *godo.ListOptions) (KubernetesNodePools, error)
	UpdateNodePool(clusterID, poolID string, req *godo.KubernetesNodePoolUpdateRequest) (*KubernetesNodePool, error)
	RecycleNodePoolNodes(clusterID, poolID string, req *godo.KubernetesNodePoolRecycleNodesRequest) error
	DeleteNodePool(clusterID, poolID string) error

	GetOptions() (*godo.KubernetesOptions, error)
}

var _ KubernetesService = &kubernetesClusterService{}

type kubernetesClusterService struct {
	client godo.KubernetesService
}

// NewKubernetesService builds an instance of KubernetesService.
func NewKubernetesService(client *godo.Client) KubernetesService {
	return &kubernetesClusterService{
		client: client.Kubernetes,
	}
}

func (k8s *kubernetesClusterService) Get(clusterID string) (*KubernetesCluster, error) {
	cluster, _, err := k8s.client.Get(context.TODO(), clusterID)
	if err != nil {
		return nil, err
	}

	return &KubernetesCluster{KubernetesCluster: cluster}, nil
}

func (k8s *kubernetesClusterService) GetKubeConfig(clusterID string) ([]byte, error) {
	config, _, err := k8s.client.GetKubeConfig(context.TODO(), clusterID)
	if err != nil {
		return nil, err
	}

	return config.KubeconfigYAML, nil
}

func (k8s *kubernetesClusterService) List() (KubernetesClusters, error) {
	f := func(opt *godo.ListOptions) ([]interface{}, *godo.Response, error) {
		list, resp, err := k8s.client.List(context.TODO(), opt)
		if err != nil {
			return nil, nil, err
		}

		si := make([]interface{}, 0, len(list))
		for _, item := range list {
			si = append(si, item)
		}

		return si, resp, err
	}

	si, err := PaginateResp(f)
	if err != nil {
		return nil, err
	}

	list := make([]KubernetesCluster, 0, len(si))
	for _, item := range si {
		a := item.(godo.KubernetesCluster)
		list = append(list, KubernetesCluster{KubernetesCluster: &a})
	}

	return list, nil
}

func (k8s *kubernetesClusterService) Create(create *godo.KubernetesClusterCreateRequest) (*KubernetesCluster, error) {
	cluster, _, err := k8s.client.Create(context.TODO(), create)
	if err != nil {
		return nil, err
	}
	return &KubernetesCluster{KubernetesCluster: cluster}, nil
}

func (k8s *kubernetesClusterService) Update(clusterID string, update *godo.KubernetesClusterUpdateRequest) (*KubernetesCluster, error) {
	cluster, _, err := k8s.client.Update(context.TODO(), clusterID, update)
	if err != nil {
		return nil, err
	}
	return &KubernetesCluster{KubernetesCluster: cluster}, nil
}

func (k8s *kubernetesClusterService) Delete(clusterID string) error {
	_, err := k8s.client.Delete(context.TODO(), clusterID)
	return err
}

func (k8s *kubernetesClusterService) CreateNodePool(clusterID string, req *godo.KubernetesNodePoolCreateRequest) (*KubernetesNodePool, error) {
	pool, _, err := k8s.client.CreateNodePool(context.TODO(), clusterID, req)
	if err != nil {
		return nil, err
	}
	return &KubernetesNodePool{KubernetesNodePool: pool}, nil
}

func (k8s *kubernetesClusterService) GetNodePool(clusterID, poolID string) (*KubernetesNodePool, error) {
	pool, _, err := k8s.client.GetNodePool(context.TODO(), clusterID, poolID)
	if err != nil {
		return nil, err
	}
	return &KubernetesNodePool{KubernetesNodePool: pool}, nil
}

func (k8s *kubernetesClusterService) ListNodePools(clusterID string, opts *godo.ListOptions) (KubernetesNodePools, error) {
	f := func(opt *godo.ListOptions) ([]interface{}, *godo.Response, error) {
		list, resp, err := k8s.client.ListNodePools(context.TODO(), clusterID, opt)
		if err != nil {
			return nil, nil, err
		}

		si := make([]interface{}, 0, len(list))
		for _, item := range list {
			si = append(si, item)
		}

		return si, resp, err
	}

	si, err := PaginateResp(f)
	if err != nil {
		return nil, err
	}

	list := make([]KubernetesNodePool, len(si))
	for _, item := range si {
		a := item.(godo.KubernetesNodePool)
		list = append(list, KubernetesNodePool{KubernetesNodePool: &a})
	}

	return list, nil

}

func (k8s *kubernetesClusterService) UpdateNodePool(clusterID, poolID string, req *godo.KubernetesNodePoolUpdateRequest) (*KubernetesNodePool, error) {
	pool, _, err := k8s.client.UpdateNodePool(context.TODO(), clusterID, poolID, req)
	if err != nil {
		return nil, err
	}
	return &KubernetesNodePool{KubernetesNodePool: pool}, nil
}

func (k8s *kubernetesClusterService) RecycleNodePoolNodes(clusterID, poolID string, req *godo.KubernetesNodePoolRecycleNodesRequest) error {
	_, err := k8s.client.RecycleNodePoolNodes(context.TODO(), clusterID, poolID, req)
	return err
}

func (k8s *kubernetesClusterService) DeleteNodePool(clusterID, poolID string) error {
	_, err := k8s.client.DeleteNodePool(context.TODO(), clusterID, poolID)
	return err
}

func (k8s *kubernetesClusterService) GetOptions() (*godo.KubernetesOptions, error) {
	opts, _, err := k8s.client.GetOptions(context.TODO())
	return opts, err
}
