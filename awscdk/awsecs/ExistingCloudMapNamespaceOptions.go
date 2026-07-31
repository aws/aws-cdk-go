package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery"
)

// The options for using an existing AWS Cloud Map namespace as the default namespace of a cluster.
//
// Example:
//   var vpc Vpc
//
//
//   // Create or reference an existing namespace
//   existingNamespace := cloudmap.NewPrivateDnsNamespace(this, jsii.String("Namespace"), &PrivateDnsNamespaceProps{
//   	Name: jsii.String("example.local"),
//   	Vpc: Vpc,
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   // Use the existing namespace as the default
//   cluster.AddExistingDefaultCloudMapNamespace(&ExistingCloudMapNamespaceOptions{
//   	Namespace: existingNamespace,
//   	UseForServiceConnect: jsii.Boolean(true),
//   })
//
type ExistingCloudMapNamespaceOptions struct {
	// This property specifies whether to set the provided namespace as the service connect default in the cluster properties.
	// Default: false.
	//
	UseForServiceConnect *bool `field:"optional" json:"useForServiceConnect" yaml:"useForServiceConnect"`
	// The existing Cloud Map namespace to use as the cluster's default namespace.
	//
	// The full `INamespace` is required (rather than a ref) because the cluster needs the
	// namespace name, ARN and type to wire up Service Discovery and Service Connect.
	//
	// [disable-awslint:prefer-ref-interface].
	Namespace awsservicediscovery.INamespace `field:"required" json:"namespace" yaml:"namespace"`
}

