package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApiKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiKeyProps := &CfnApiKeyProps{
//   	CustomerId: jsii.String("customerId"),
//   	Description: jsii.String("description"),
//   	Enabled: jsii.Boolean(false),
//   	GenerateDistinctId: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	StageKeys: []interface{}{
//   		&StageKeyProperty{
//   			RestApiId: jsii.String("restApiId"),
//   			StageName: jsii.String("stageName"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html
//
type CfnApiKeyProps struct {
	// An AWS Marketplace customer identifier, when integrating with the AWS SaaS Marketplace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apikey-customerid
	//
	CustomerId *string `field:"optional" json:"customerId" yaml:"customerId"`
	// The description of the ApiKey.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apikey-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the ApiKey can be used by callers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apikey-enabled
	//
	// Default: - false.
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies whether ( `true` ) or not ( `false` ) the key identifier is distinct from the created API key value.
	//
	// This parameter is deprecated and should not be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apikey-generatedistinctid
	//
	GenerateDistinctId interface{} `field:"optional" json:"generateDistinctId" yaml:"generateDistinctId"`
	// A name for the API key.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apikey-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// DEPRECATED FOR USAGE PLANS - Specifies stages associated with the API key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apikey-stagekeys
	//
	StageKeys interface{} `field:"optional" json:"stageKeys" yaml:"stageKeys"`
	// The key-value map of strings.
	//
	// The valid character set is [a-zA-Z+-=._:/]. The tag key can be up to 128 characters and must not start with `aws:` . The tag value can be up to 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apikey-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies a value of the API key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apikey-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

