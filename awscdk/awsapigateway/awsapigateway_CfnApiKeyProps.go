package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnApiKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiKeyProps := &cfnApiKeyProps{
//   	customerId: jsii.String("customerId"),
//   	description: jsii.String("description"),
//   	enabled: jsii.Boolean(false),
//   	generateDistinctId: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	stageKeys: []interface{}{
//   		&stageKeyProperty{
//   			restApiId: jsii.String("restApiId"),
//   			stageName: jsii.String("stageName"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	value: jsii.String("value"),
//   }
//
type CfnApiKeyProps struct {
	// An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.
	CustomerId *string `field:"optional" json:"customerId" yaml:"customerId"`
	// A description of the purpose of the API key.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether the API key can be used by clients.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies whether the key identifier is distinct from the created API key value.
	//
	// This parameter is deprecated and should not be used.
	GenerateDistinctId interface{} `field:"optional" json:"generateDistinctId" yaml:"generateDistinctId"`
	// A name for the API key.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of stages to associate with this API key.
	StageKeys interface{} `field:"optional" json:"stageKeys" yaml:"stageKeys"`
	// An array of arbitrary tags (key-value pairs) to associate with the API key.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The value of the API key.
	//
	// Must be at least 20 characters long.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

