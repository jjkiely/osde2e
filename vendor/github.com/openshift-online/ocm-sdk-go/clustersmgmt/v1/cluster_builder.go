/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	time "time"
)

// ClusterBuilder contains the data and logic needed to build 'cluster' objects.
//
// Definition of an _OpenShift_ cluster.
//
// The `cloud_provider` attribute is a reference to the cloud provider. When a
// cluster is retrieved it will be a link to the cloud provider, containing only
// the kind, id and href attributes:
//
// [source,json]
// ----
// {
//   "cloud_provider": {
//     "kind": "CloudProviderLink",
//     "id": "123",
//     "href": "/api/clusters_mgmt/v1/cloud_providers/123"
//   }
// }
// ----
//
// When a cluster is created this is optional, and if used it should contain the
// identifier of the cloud provider to use:
//
// [source,json]
// ----
// {
//   "cloud_provider": {
//     "id": "123",
//   }
// }
// ----
//
// If not included, then the cluster will be created using the default cloud
// provider, which is currently Amazon Web Services.
//
// The region attribute is mandatory when a cluster is created.
//
// The `aws.access_key_id`, `aws.secret_access_key` and `dns.base_domain`
// attributes are mandatory when creation a cluster with your own Amazon Web
// Services account.
type ClusterBuilder struct {
	id                                *string
	href                              *string
	link                              bool
	api                               *ClusterAPIBuilder
	aws                               *AWSBuilder
	awsInfrastructureAccessRoleGrants *AWSInfrastructureAccessRoleGrantListBuilder
	ccs                               *CCSBuilder
	dns                               *DNSBuilder
	gcp                               *GCPBuilder
	addons                            *AddOnInstallationListBuilder
	cloudProvider                     *CloudProviderBuilder
	clusterAdminEnabled               *bool
	console                           *ClusterConsoleBuilder
	creationTimestamp                 *time.Time
	displayName                       *string
	etcdEncryption                    *bool
	expirationTimestamp               *time.Time
	externalID                        *string
	externalConfiguration             *ExternalConfigurationBuilder
	flavour                           *FlavourBuilder
	groups                            *GroupListBuilder
	healthState                       *ClusterHealthState
	identityProviders                 *IdentityProviderListBuilder
	ingresses                         *IngressListBuilder
	loadBalancerQuota                 *int
	machinePools                      *MachinePoolListBuilder
	managed                           *bool
	metrics                           *ClusterMetricsBuilder
	multiAZ                           *bool
	name                              *string
	network                           *NetworkBuilder
	nodeDrainGracePeriod              *ValueBuilder
	nodes                             *ClusterNodesBuilder
	openshiftVersion                  *string
	product                           *ProductBuilder
	properties                        map[string]string
	provisionShard                    *ProvisionShardBuilder
	region                            *CloudRegionBuilder
	state                             *ClusterState
	status                            *ClusterStatusBuilder
	storageQuota                      *ValueBuilder
	subscription                      *SubscriptionBuilder
	upgradeChannelGroup               *string
	version                           *VersionBuilder
}

// NewCluster creates a new builder of 'cluster' objects.
func NewCluster() *ClusterBuilder {
	return new(ClusterBuilder)
}

// ID sets the identifier of the object.
func (b *ClusterBuilder) ID(value string) *ClusterBuilder {
	b.id = &value
	return b
}

// HREF sets the link to the object.
func (b *ClusterBuilder) HREF(value string) *ClusterBuilder {
	b.href = &value
	return b
}

// Link sets the flag that indicates if this is a link.
func (b *ClusterBuilder) Link(value bool) *ClusterBuilder {
	b.link = value
	return b
}

// API sets the value of the 'API' attribute to the given value.
//
// Information about the API of a cluster.
func (b *ClusterBuilder) API(value *ClusterAPIBuilder) *ClusterBuilder {
	b.api = value
	return b
}

// AWS sets the value of the 'AWS' attribute to the given value.
//
// _Amazon Web Services_ specific settings of a cluster.
func (b *ClusterBuilder) AWS(value *AWSBuilder) *ClusterBuilder {
	b.aws = value
	return b
}

// AWSInfrastructureAccessRoleGrants sets the value of the 'AWS_infrastructure_access_role_grants' attribute to the given values.
//
//
func (b *ClusterBuilder) AWSInfrastructureAccessRoleGrants(value *AWSInfrastructureAccessRoleGrantListBuilder) *ClusterBuilder {
	b.awsInfrastructureAccessRoleGrants = value
	return b
}

// CCS sets the value of the 'CCS' attribute to the given value.
//
//
func (b *ClusterBuilder) CCS(value *CCSBuilder) *ClusterBuilder {
	b.ccs = value
	return b
}

// DNS sets the value of the 'DNS' attribute to the given value.
//
// DNS settings of the cluster.
func (b *ClusterBuilder) DNS(value *DNSBuilder) *ClusterBuilder {
	b.dns = value
	return b
}

// GCP sets the value of the 'GCP' attribute to the given value.
//
// Google cloud platform settings of a cluster.
func (b *ClusterBuilder) GCP(value *GCPBuilder) *ClusterBuilder {
	b.gcp = value
	return b
}

// Addons sets the value of the 'addons' attribute to the given values.
//
//
func (b *ClusterBuilder) Addons(value *AddOnInstallationListBuilder) *ClusterBuilder {
	b.addons = value
	return b
}

// CloudProvider sets the value of the 'cloud_provider' attribute to the given value.
//
// Cloud provider.
func (b *ClusterBuilder) CloudProvider(value *CloudProviderBuilder) *ClusterBuilder {
	b.cloudProvider = value
	return b
}

// ClusterAdminEnabled sets the value of the 'cluster_admin_enabled' attribute to the given value.
//
//
func (b *ClusterBuilder) ClusterAdminEnabled(value bool) *ClusterBuilder {
	b.clusterAdminEnabled = &value
	return b
}

// Console sets the value of the 'console' attribute to the given value.
//
// Information about the console of a cluster.
func (b *ClusterBuilder) Console(value *ClusterConsoleBuilder) *ClusterBuilder {
	b.console = value
	return b
}

// CreationTimestamp sets the value of the 'creation_timestamp' attribute to the given value.
//
//
func (b *ClusterBuilder) CreationTimestamp(value time.Time) *ClusterBuilder {
	b.creationTimestamp = &value
	return b
}

// DisplayName sets the value of the 'display_name' attribute to the given value.
//
//
func (b *ClusterBuilder) DisplayName(value string) *ClusterBuilder {
	b.displayName = &value
	return b
}

// EtcdEncryption sets the value of the 'etcd_encryption' attribute to the given value.
//
//
func (b *ClusterBuilder) EtcdEncryption(value bool) *ClusterBuilder {
	b.etcdEncryption = &value
	return b
}

// ExpirationTimestamp sets the value of the 'expiration_timestamp' attribute to the given value.
//
//
func (b *ClusterBuilder) ExpirationTimestamp(value time.Time) *ClusterBuilder {
	b.expirationTimestamp = &value
	return b
}

// ExternalID sets the value of the 'external_ID' attribute to the given value.
//
//
func (b *ClusterBuilder) ExternalID(value string) *ClusterBuilder {
	b.externalID = &value
	return b
}

// ExternalConfiguration sets the value of the 'external_configuration' attribute to the given value.
//
// Representation of cluster external configuration.
func (b *ClusterBuilder) ExternalConfiguration(value *ExternalConfigurationBuilder) *ClusterBuilder {
	b.externalConfiguration = value
	return b
}

// Flavour sets the value of the 'flavour' attribute to the given value.
//
// Set of predefined properties of a cluster. For example, a _huge_ flavour can be a cluster
// with 10 infra nodes and 1000 compute nodes.
func (b *ClusterBuilder) Flavour(value *FlavourBuilder) *ClusterBuilder {
	b.flavour = value
	return b
}

// Groups sets the value of the 'groups' attribute to the given values.
//
//
func (b *ClusterBuilder) Groups(value *GroupListBuilder) *ClusterBuilder {
	b.groups = value
	return b
}

// HealthState sets the value of the 'health_state' attribute to the given value.
//
// ClusterHealthState indicates the health of a cluster.
func (b *ClusterBuilder) HealthState(value ClusterHealthState) *ClusterBuilder {
	b.healthState = &value
	return b
}

// IdentityProviders sets the value of the 'identity_providers' attribute to the given values.
//
//
func (b *ClusterBuilder) IdentityProviders(value *IdentityProviderListBuilder) *ClusterBuilder {
	b.identityProviders = value
	return b
}

// Ingresses sets the value of the 'ingresses' attribute to the given values.
//
//
func (b *ClusterBuilder) Ingresses(value *IngressListBuilder) *ClusterBuilder {
	b.ingresses = value
	return b
}

// LoadBalancerQuota sets the value of the 'load_balancer_quota' attribute to the given value.
//
//
func (b *ClusterBuilder) LoadBalancerQuota(value int) *ClusterBuilder {
	b.loadBalancerQuota = &value
	return b
}

// MachinePools sets the value of the 'machine_pools' attribute to the given values.
//
//
func (b *ClusterBuilder) MachinePools(value *MachinePoolListBuilder) *ClusterBuilder {
	b.machinePools = value
	return b
}

// Managed sets the value of the 'managed' attribute to the given value.
//
//
func (b *ClusterBuilder) Managed(value bool) *ClusterBuilder {
	b.managed = &value
	return b
}

// Metrics sets the value of the 'metrics' attribute to the given value.
//
// Cluster metrics received via telemetry.
func (b *ClusterBuilder) Metrics(value *ClusterMetricsBuilder) *ClusterBuilder {
	b.metrics = value
	return b
}

// MultiAZ sets the value of the 'multi_AZ' attribute to the given value.
//
//
func (b *ClusterBuilder) MultiAZ(value bool) *ClusterBuilder {
	b.multiAZ = &value
	return b
}

// Name sets the value of the 'name' attribute to the given value.
//
//
func (b *ClusterBuilder) Name(value string) *ClusterBuilder {
	b.name = &value
	return b
}

// Network sets the value of the 'network' attribute to the given value.
//
// Network configuration of a cluster.
func (b *ClusterBuilder) Network(value *NetworkBuilder) *ClusterBuilder {
	b.network = value
	return b
}

// NodeDrainGracePeriod sets the value of the 'node_drain_grace_period' attribute to the given value.
//
// Numeric value and the unit used to measure it.
//
// Units are not mandatory, and they're not specified for some resources. For
// resources that use bytes, the accepted units are:
//
// - 1 B = 1 byte
// - 1 KB = 10^3 bytes
// - 1 MB = 10^6 bytes
// - 1 GB = 10^9 bytes
// - 1 TB = 10^12 bytes
// - 1 PB = 10^15 bytes
//
// - 1 B = 1 byte
// - 1 KiB = 2^10 bytes
// - 1 MiB = 2^20 bytes
// - 1 GiB = 2^30 bytes
// - 1 TiB = 2^40 bytes
// - 1 PiB = 2^50 bytes
func (b *ClusterBuilder) NodeDrainGracePeriod(value *ValueBuilder) *ClusterBuilder {
	b.nodeDrainGracePeriod = value
	return b
}

// Nodes sets the value of the 'nodes' attribute to the given value.
//
// Counts of different classes of nodes inside a cluster.
func (b *ClusterBuilder) Nodes(value *ClusterNodesBuilder) *ClusterBuilder {
	b.nodes = value
	return b
}

// OpenshiftVersion sets the value of the 'openshift_version' attribute to the given value.
//
//
func (b *ClusterBuilder) OpenshiftVersion(value string) *ClusterBuilder {
	b.openshiftVersion = &value
	return b
}

// Product sets the value of the 'product' attribute to the given value.
//
// Representation of an product that can be selected as a cluster type.
func (b *ClusterBuilder) Product(value *ProductBuilder) *ClusterBuilder {
	b.product = value
	return b
}

// Properties sets the value of the 'properties' attribute to the given value.
//
//
func (b *ClusterBuilder) Properties(value map[string]string) *ClusterBuilder {
	b.properties = value
	return b
}

// ProvisionShard sets the value of the 'provision_shard' attribute to the given value.
//
// Contains the properties of the provision shard, including AWS and GCP related configurations
func (b *ClusterBuilder) ProvisionShard(value *ProvisionShardBuilder) *ClusterBuilder {
	b.provisionShard = value
	return b
}

// Region sets the value of the 'region' attribute to the given value.
//
// Description of a region of a cloud provider.
func (b *ClusterBuilder) Region(value *CloudRegionBuilder) *ClusterBuilder {
	b.region = value
	return b
}

// State sets the value of the 'state' attribute to the given value.
//
// Overall state of a cluster.
func (b *ClusterBuilder) State(value ClusterState) *ClusterBuilder {
	b.state = &value
	return b
}

// Status sets the value of the 'status' attribute to the given value.
//
// Detailed status of a cluster.
func (b *ClusterBuilder) Status(value *ClusterStatusBuilder) *ClusterBuilder {
	b.status = value
	return b
}

// StorageQuota sets the value of the 'storage_quota' attribute to the given value.
//
// Numeric value and the unit used to measure it.
//
// Units are not mandatory, and they're not specified for some resources. For
// resources that use bytes, the accepted units are:
//
// - 1 B = 1 byte
// - 1 KB = 10^3 bytes
// - 1 MB = 10^6 bytes
// - 1 GB = 10^9 bytes
// - 1 TB = 10^12 bytes
// - 1 PB = 10^15 bytes
//
// - 1 B = 1 byte
// - 1 KiB = 2^10 bytes
// - 1 MiB = 2^20 bytes
// - 1 GiB = 2^30 bytes
// - 1 TiB = 2^40 bytes
// - 1 PiB = 2^50 bytes
func (b *ClusterBuilder) StorageQuota(value *ValueBuilder) *ClusterBuilder {
	b.storageQuota = value
	return b
}

// Subscription sets the value of the 'subscription' attribute to the given value.
//
// Definition of a subscription.
func (b *ClusterBuilder) Subscription(value *SubscriptionBuilder) *ClusterBuilder {
	b.subscription = value
	return b
}

// UpgradeChannelGroup sets the value of the 'upgrade_channel_group' attribute to the given value.
//
//
func (b *ClusterBuilder) UpgradeChannelGroup(value string) *ClusterBuilder {
	b.upgradeChannelGroup = &value
	return b
}

// Version sets the value of the 'version' attribute to the given value.
//
// Representation of an _OpenShift_ version.
func (b *ClusterBuilder) Version(value *VersionBuilder) *ClusterBuilder {
	b.version = value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ClusterBuilder) Copy(object *Cluster) *ClusterBuilder {
	if object == nil {
		return b
	}
	b.id = object.id
	b.href = object.href
	b.link = object.link
	if object.api != nil {
		b.api = NewClusterAPI().Copy(object.api)
	} else {
		b.api = nil
	}
	if object.aws != nil {
		b.aws = NewAWS().Copy(object.aws)
	} else {
		b.aws = nil
	}
	if object.awsInfrastructureAccessRoleGrants != nil {
		b.awsInfrastructureAccessRoleGrants = NewAWSInfrastructureAccessRoleGrantList().Copy(object.awsInfrastructureAccessRoleGrants)
	} else {
		b.awsInfrastructureAccessRoleGrants = nil
	}
	if object.ccs != nil {
		b.ccs = NewCCS().Copy(object.ccs)
	} else {
		b.ccs = nil
	}
	if object.dns != nil {
		b.dns = NewDNS().Copy(object.dns)
	} else {
		b.dns = nil
	}
	if object.gcp != nil {
		b.gcp = NewGCP().Copy(object.gcp)
	} else {
		b.gcp = nil
	}
	if object.addons != nil {
		b.addons = NewAddOnInstallationList().Copy(object.addons)
	} else {
		b.addons = nil
	}
	if object.cloudProvider != nil {
		b.cloudProvider = NewCloudProvider().Copy(object.cloudProvider)
	} else {
		b.cloudProvider = nil
	}
	b.clusterAdminEnabled = object.clusterAdminEnabled
	if object.console != nil {
		b.console = NewClusterConsole().Copy(object.console)
	} else {
		b.console = nil
	}
	b.creationTimestamp = object.creationTimestamp
	b.displayName = object.displayName
	b.etcdEncryption = object.etcdEncryption
	b.expirationTimestamp = object.expirationTimestamp
	b.externalID = object.externalID
	if object.externalConfiguration != nil {
		b.externalConfiguration = NewExternalConfiguration().Copy(object.externalConfiguration)
	} else {
		b.externalConfiguration = nil
	}
	if object.flavour != nil {
		b.flavour = NewFlavour().Copy(object.flavour)
	} else {
		b.flavour = nil
	}
	if object.groups != nil {
		b.groups = NewGroupList().Copy(object.groups)
	} else {
		b.groups = nil
	}
	b.healthState = object.healthState
	if object.identityProviders != nil {
		b.identityProviders = NewIdentityProviderList().Copy(object.identityProviders)
	} else {
		b.identityProviders = nil
	}
	if object.ingresses != nil {
		b.ingresses = NewIngressList().Copy(object.ingresses)
	} else {
		b.ingresses = nil
	}
	b.loadBalancerQuota = object.loadBalancerQuota
	if object.machinePools != nil {
		b.machinePools = NewMachinePoolList().Copy(object.machinePools)
	} else {
		b.machinePools = nil
	}
	b.managed = object.managed
	if object.metrics != nil {
		b.metrics = NewClusterMetrics().Copy(object.metrics)
	} else {
		b.metrics = nil
	}
	b.multiAZ = object.multiAZ
	b.name = object.name
	if object.network != nil {
		b.network = NewNetwork().Copy(object.network)
	} else {
		b.network = nil
	}
	if object.nodeDrainGracePeriod != nil {
		b.nodeDrainGracePeriod = NewValue().Copy(object.nodeDrainGracePeriod)
	} else {
		b.nodeDrainGracePeriod = nil
	}
	if object.nodes != nil {
		b.nodes = NewClusterNodes().Copy(object.nodes)
	} else {
		b.nodes = nil
	}
	b.openshiftVersion = object.openshiftVersion
	if object.product != nil {
		b.product = NewProduct().Copy(object.product)
	} else {
		b.product = nil
	}
	if len(object.properties) > 0 {
		b.properties = make(map[string]string)
		for k, v := range object.properties {
			b.properties[k] = v
		}
	} else {
		b.properties = nil
	}
	if object.provisionShard != nil {
		b.provisionShard = NewProvisionShard().Copy(object.provisionShard)
	} else {
		b.provisionShard = nil
	}
	if object.region != nil {
		b.region = NewCloudRegion().Copy(object.region)
	} else {
		b.region = nil
	}
	b.state = object.state
	if object.status != nil {
		b.status = NewClusterStatus().Copy(object.status)
	} else {
		b.status = nil
	}
	if object.storageQuota != nil {
		b.storageQuota = NewValue().Copy(object.storageQuota)
	} else {
		b.storageQuota = nil
	}
	if object.subscription != nil {
		b.subscription = NewSubscription().Copy(object.subscription)
	} else {
		b.subscription = nil
	}
	b.upgradeChannelGroup = object.upgradeChannelGroup
	if object.version != nil {
		b.version = NewVersion().Copy(object.version)
	} else {
		b.version = nil
	}
	return b
}

// Build creates a 'cluster' object using the configuration stored in the builder.
func (b *ClusterBuilder) Build() (object *Cluster, err error) {
	object = new(Cluster)
	object.id = b.id
	object.href = b.href
	object.link = b.link
	if b.api != nil {
		object.api, err = b.api.Build()
		if err != nil {
			return
		}
	}
	if b.aws != nil {
		object.aws, err = b.aws.Build()
		if err != nil {
			return
		}
	}
	if b.awsInfrastructureAccessRoleGrants != nil {
		object.awsInfrastructureAccessRoleGrants, err = b.awsInfrastructureAccessRoleGrants.Build()
		if err != nil {
			return
		}
	}
	if b.ccs != nil {
		object.ccs, err = b.ccs.Build()
		if err != nil {
			return
		}
	}
	if b.dns != nil {
		object.dns, err = b.dns.Build()
		if err != nil {
			return
		}
	}
	if b.gcp != nil {
		object.gcp, err = b.gcp.Build()
		if err != nil {
			return
		}
	}
	if b.addons != nil {
		object.addons, err = b.addons.Build()
		if err != nil {
			return
		}
	}
	if b.cloudProvider != nil {
		object.cloudProvider, err = b.cloudProvider.Build()
		if err != nil {
			return
		}
	}
	object.clusterAdminEnabled = b.clusterAdminEnabled
	if b.console != nil {
		object.console, err = b.console.Build()
		if err != nil {
			return
		}
	}
	object.creationTimestamp = b.creationTimestamp
	object.displayName = b.displayName
	object.etcdEncryption = b.etcdEncryption
	object.expirationTimestamp = b.expirationTimestamp
	object.externalID = b.externalID
	if b.externalConfiguration != nil {
		object.externalConfiguration, err = b.externalConfiguration.Build()
		if err != nil {
			return
		}
	}
	if b.flavour != nil {
		object.flavour, err = b.flavour.Build()
		if err != nil {
			return
		}
	}
	if b.groups != nil {
		object.groups, err = b.groups.Build()
		if err != nil {
			return
		}
	}
	object.healthState = b.healthState
	if b.identityProviders != nil {
		object.identityProviders, err = b.identityProviders.Build()
		if err != nil {
			return
		}
	}
	if b.ingresses != nil {
		object.ingresses, err = b.ingresses.Build()
		if err != nil {
			return
		}
	}
	object.loadBalancerQuota = b.loadBalancerQuota
	if b.machinePools != nil {
		object.machinePools, err = b.machinePools.Build()
		if err != nil {
			return
		}
	}
	object.managed = b.managed
	if b.metrics != nil {
		object.metrics, err = b.metrics.Build()
		if err != nil {
			return
		}
	}
	object.multiAZ = b.multiAZ
	object.name = b.name
	if b.network != nil {
		object.network, err = b.network.Build()
		if err != nil {
			return
		}
	}
	if b.nodeDrainGracePeriod != nil {
		object.nodeDrainGracePeriod, err = b.nodeDrainGracePeriod.Build()
		if err != nil {
			return
		}
	}
	if b.nodes != nil {
		object.nodes, err = b.nodes.Build()
		if err != nil {
			return
		}
	}
	object.openshiftVersion = b.openshiftVersion
	if b.product != nil {
		object.product, err = b.product.Build()
		if err != nil {
			return
		}
	}
	if b.properties != nil {
		object.properties = make(map[string]string)
		for k, v := range b.properties {
			object.properties[k] = v
		}
	}
	if b.provisionShard != nil {
		object.provisionShard, err = b.provisionShard.Build()
		if err != nil {
			return
		}
	}
	if b.region != nil {
		object.region, err = b.region.Build()
		if err != nil {
			return
		}
	}
	object.state = b.state
	if b.status != nil {
		object.status, err = b.status.Build()
		if err != nil {
			return
		}
	}
	if b.storageQuota != nil {
		object.storageQuota, err = b.storageQuota.Build()
		if err != nil {
			return
		}
	}
	if b.subscription != nil {
		object.subscription, err = b.subscription.Build()
		if err != nil {
			return
		}
	}
	object.upgradeChannelGroup = b.upgradeChannelGroup
	if b.version != nil {
		object.version, err = b.version.Build()
		if err != nil {
			return
		}
	}
	return
}
