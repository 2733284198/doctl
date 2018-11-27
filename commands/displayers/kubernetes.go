package displayers

import (
	"fmt"
	"io"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/digitalocean/doctl/do"
)

type KubernetesClusters struct {
	KubernetesClusters do.KubernetesClusters
}

var _ Displayable = &KubernetesClusters{}

func (clusters *KubernetesClusters) JSON(out io.Writer) error {
	return writeJSON(clusters.KubernetesClusters, out)
}

func (clusters *KubernetesClusters) Cols() []string {
	return []string{
		"ID",
		"Name",
		"Region",
		"Version",
		"Status",
		"Endpoint",
		"IPv4",
		"ClusterSubnet",
		"ServiceSubnet",
		"Tags",
		"Created",
		"Updated",
		"NodePools",
	}
}

func (clusters *KubernetesClusters) ColMap() map[string]string {
	return map[string]string{
		"ID":            "ID",
		"Name":          "Name",
		"Region":        "Region",
		"Version":       "Version",
		"ClusterSubnet": "Cluster Subnet",
		"ServiceSubnet": "Service Subnet",
		"IPv4":          "IPv4",
		"Endpoint":      "Endpoint",
		"Tags":          "Tags",
		"Status":        "Status",
		"Created":       "Created At",
		"Updated":       "Updated At",
		"NodePools":     "Node Pools",
	}
}

func (clusters *KubernetesClusters) KV() []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(clusters.KubernetesClusters))

	for _, cluster := range clusters.KubernetesClusters {
		tags := strings.Join(cluster.Tags, ",")
		nodePools := make([]string, 0, len(cluster.NodePools))
		for _, pool := range cluster.NodePools {
			nodePools = append(nodePools, pool.Name)
		}
		if cluster.Status == nil {
			cluster.Status = new(godo.KubernetesClusterStatus)
		}

		o := map[string]interface{}{
			"ID":            cluster.ID,
			"Name":          cluster.Name,
			"Region":        cluster.RegionSlug,
			"Version":       cluster.VersionSlug,
			"ClusterSubnet": cluster.ClusterSubnet,
			"ServiceSubnet": cluster.ServiceSubnet,
			"IPv4":          cluster.IPv4,
			"Endpoint":      cluster.Endpoint,
			"Tags":          tags,
			"Status":        cluster.Status.State,
			"Created":       cluster.CreatedAt,
			"Updated":       cluster.UpdatedAt,
			"NodePools":     fmt.Sprintf(strings.Join(nodePools, " ")),
		}
		out = append(out, o)
	}

	return out
}

type KubernetesNodePools struct {
	KubernetesNodePools do.KubernetesNodePools
}

var _ Displayable = &KubernetesNodePools{}

func (nodePools *KubernetesNodePools) JSON(out io.Writer) error {
	return writeJSON(nodePools.KubernetesNodePools, out)
}

func (nodePools *KubernetesNodePools) Cols() []string {
	return []string{
		"ID",
		"Name",
		"Size",
		"Count",
		"Tags",
		"Nodes",
	}
}

func (nodePools *KubernetesNodePools) ColMap() map[string]string {
	return map[string]string{
		"ID":    "ID",
		"Name":  "Name",
		"Size":  "Size",
		"Count": "Count",
		"Tags":  "Tags",
		"Nodes": "Nodes",
	}
}

func (nodePools *KubernetesNodePools) KV() []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(nodePools.KubernetesNodePools))

	for _, nodePools := range nodePools.KubernetesNodePools {
		tags := strings.Join(nodePools.Tags, ",")
		nodes := make([]string, 0, len(nodePools.Nodes))
		for _, node := range nodePools.Nodes {
			nodes = append(nodes, node.Name)
		}

		o := map[string]interface{}{
			"ID":    nodePools.ID,
			"Name":  nodePools.Name,
			"Size":  nodePools.Size,
			"Count": nodePools.Count,
			"Tags":  tags,
			"Nodes": nodes,
		}
		out = append(out, o)
	}

	return out
}

type KubernetesVersions struct {
	KubernetesVersions do.KubernetesVersions
}

var _ Displayable = &KubernetesVersions{}

func (versions *KubernetesVersions) JSON(out io.Writer) error {
	return writeJSON(versions.KubernetesVersions, out)
}

func (versions *KubernetesVersions) Cols() []string {
	return []string{
		"Slug",
		"KubernetesVersion",
	}
}

func (versions *KubernetesVersions) ColMap() map[string]string {
	return map[string]string{
		"Slug":              "Slug",
		"KubernetesVersion": "Kubernetes Version",
	}
}

func (versions *KubernetesVersions) KV() []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(versions.KubernetesVersions))

	for _, version := range versions.KubernetesVersions {

		o := map[string]interface{}{
			"Slug":              version.KubernetesVersion.Slug,
			"KubernetesVersion": version.KubernetesVersion.KubernetesVersion,
		}
		out = append(out, o)
	}

	return out
}
