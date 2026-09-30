package awsnetworkfirewall


// Contains the monitoring configuration for a single cluster in a container association.
//
// Specifies the cluster ARN and optional attribute filters to narrow which containers are tracked.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerMonitoringConfigurationProperty := &ContainerMonitoringConfigurationProperty{
//   	ClusterArn: jsii.String("clusterArn"),
//
//   	// the properties below are optional
//   	AttributeFilters: []interface{}{
//   		&ContainerAttributeProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-containerassociation-containermonitoringconfiguration.html
//
type CfnContainerAssociation_ContainerMonitoringConfigurationProperty struct {
	// The ARN of the Amazon ECS or Amazon EKS cluster to monitor.
	//
	// The cluster must be in the same Region and account as the container association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-containerassociation-containermonitoringconfiguration.html#cfn-networkfirewall-containerassociation-containermonitoringconfiguration-clusterarn
	//
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// Key-value pairs that filter which containers are tracked.
	//
	// For Amazon EKS, you can filter by namespace and Kubernetes labels. For Amazon ECS, you can filter by container instance attributes (EC2 launch type only).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-containerassociation-containermonitoringconfiguration.html#cfn-networkfirewall-containerassociation-containermonitoringconfiguration-attributefilters
	//
	AttributeFilters interface{} `field:"optional" json:"attributeFilters" yaml:"attributeFilters"`
}

