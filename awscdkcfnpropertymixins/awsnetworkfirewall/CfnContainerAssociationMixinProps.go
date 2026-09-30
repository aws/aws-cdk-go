package awsnetworkfirewall

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnContainerAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnContainerAssociationMixinProps := &CfnContainerAssociationMixinProps{
//   	ContainerAssociationName: jsii.String("containerAssociationName"),
//   	ContainerMonitoringConfigurations: []interface{}{
//   		&ContainerMonitoringConfigurationProperty{
//   			AttributeFilters: []interface{}{
//   				&ContainerAttributeProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			ClusterArn: jsii.String("clusterArn"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-containerassociation.html
//
type CfnContainerAssociationMixinProps struct {
	// The descriptive name of the container association.
	//
	// You can't change the name of a container association after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-containerassociation.html#cfn-networkfirewall-containerassociation-containerassociationname
	//
	ContainerAssociationName *string `field:"optional" json:"containerAssociationName" yaml:"containerAssociationName"`
	// The monitoring configurations for the container association.
	//
	// Each configuration specifies an Amazon ECS or Amazon EKS cluster to monitor and optional attribute filters to narrow which containers are tracked.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-containerassociation.html#cfn-networkfirewall-containerassociation-containermonitoringconfigurations
	//
	ContainerMonitoringConfigurations interface{} `field:"optional" json:"containerMonitoringConfigurations" yaml:"containerMonitoringConfigurations"`
	// A description of the container association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-containerassociation.html#cfn-networkfirewall-containerassociation-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-containerassociation.html#cfn-networkfirewall-containerassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of containers to monitor.
	//
	// You can't change the container type after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-containerassociation.html#cfn-networkfirewall-containerassociation-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

