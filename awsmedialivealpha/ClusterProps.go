package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for creating a MediaLive Cluster.
//
// Example:
//   var stack Stack
//   var instanceRole IRole
//
//
//   cluster := medialive.NewCluster(stack, jsii.String("Cluster"), &ClusterProps{
//   	ClusterName: jsii.String("on-prem-cluster"),
//   	ClusterType: medialive.ClusterType_ON_PREMISES(),
//   	InstanceRole: InstanceRole,
//   })
//
// Experimental.
type ClusterProps struct {
	// The IAM role for nodes in the cluster.
	//
	// [disable-awslint:prefer-ref-interface].
	// Experimental.
	InstanceRole awsiam.IRole `field:"required" json:"instanceRole" yaml:"instanceRole"`
	// The name of the cluster.
	// Default: - auto-generated name.
	//
	// Experimental.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The hardware type for the cluster.
	// Default: ClusterType.ON_PREMISES
	//
	// Experimental.
	ClusterType ClusterType `field:"optional" json:"clusterType" yaml:"clusterType"`
	// Network settings for the cluster - only required if your networking setup requires it.
	// See: https://docs.aws.amazon.com/medialive/latest/ug/emla-deploy-identify-network-requirements.html
	//
	// Default: - no network settings.
	//
	// Experimental.
	NetworkSettings *ClusterNetworkSettings `field:"optional" json:"networkSettings" yaml:"networkSettings"`
	// Tags to add to the cluster.
	// Default: - no tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

