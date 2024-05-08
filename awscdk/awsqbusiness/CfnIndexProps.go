package awsqbusiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIndex`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIndexProps := &CfnIndexProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	DisplayName: jsii.String("displayName"),
//
//   	// the properties below are optional
//   	CapacityConfiguration: &IndexCapacityConfigurationProperty{
//   		Units: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	DocumentAttributeConfigurations: []interface{}{
//   		&DocumentAttributeConfigurationProperty{
//   			Name: jsii.String("name"),
//   			Search: jsii.String("search"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-index.html
//
type CfnIndexProps struct {
	// The identifier of the Amazon Q Business application using the index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-index.html#cfn-qbusiness-index-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The name of the index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-index.html#cfn-qbusiness-index-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The capacity units you want to provision for your index.
	//
	// You can add and remove capacity to fit your usage needs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-index.html#cfn-qbusiness-index-capacityconfiguration
	//
	CapacityConfiguration interface{} `field:"optional" json:"capacityConfiguration" yaml:"capacityConfiguration"`
	// A description for the Amazon Q Business index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-index.html#cfn-qbusiness-index-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration information for document attributes.
	//
	// Document attributes are metadata or fields associated with your documents. For example, the company department name associated with each document.
	//
	// For more information, see [Understanding document attributes](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/doc-attributes.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-index.html#cfn-qbusiness-index-documentattributeconfigurations
	//
	DocumentAttributeConfigurations interface{} `field:"optional" json:"documentAttributeConfigurations" yaml:"documentAttributeConfigurations"`
	// A list of key-value pairs that identify or categorize the index.
	//
	// You can also use tags to help control access to the index. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + -
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-index.html#cfn-qbusiness-index-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The index type that's suitable for your needs.
	//
	// For more information on what's included in each type of index, see [Amazon Q Business tiers](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/tiers.html#index-tiers) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-index.html#cfn-qbusiness-index-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

