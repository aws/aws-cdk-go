package awsopensearchserverless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCollectionGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCollectionGroupProps := &CfnCollectionGroupProps{
//   	Name: jsii.String("name"),
//   	StandbyReplicas: jsii.String("standbyReplicas"),
//
//   	// the properties below are optional
//   	CapacityLimits: &CapacityLimitsProperty{
//   		MaxIndexingCapacityInOcu: jsii.Number(123),
//   		MaxSearchCapacityInOcu: jsii.Number(123),
//   		MinIndexingCapacityInOcu: jsii.Number(123),
//   		MinSearchCapacityInOcu: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collectiongroup.html
//
type CfnCollectionGroupProps struct {
	// The name of the collection group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collectiongroup.html#cfn-opensearchserverless-collectiongroup-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicates whether standby replicas are used for the collection group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collectiongroup.html#cfn-opensearchserverless-collectiongroup-standbyreplicas
	//
	StandbyReplicas *string `field:"required" json:"standbyReplicas" yaml:"standbyReplicas"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collectiongroup.html#cfn-opensearchserverless-collectiongroup-capacitylimits
	//
	CapacityLimits interface{} `field:"optional" json:"capacityLimits" yaml:"capacityLimits"`
	// The description of the collection group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collectiongroup.html#cfn-opensearchserverless-collectiongroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collectiongroup.html#cfn-opensearchserverless-collectiongroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

