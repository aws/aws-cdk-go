package awscloudwatch


// Properties for defining a `CfnIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIntegrationProps := &CfnIntegrationProps{
//   	IntegrationType: jsii.String("integrationType"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Credential: &IntegrationCredentialProperty{
//   		ApiKeyCredential: &ApiKeyCredentialProperty{
//   			ApiKeyValue: jsii.String("apiKeyValue"),
//   		},
//   	},
//   	IntegrationAttributes: map[string]*string{
//   		"integrationAttributesKey": jsii.String("integrationAttributes"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-integration.html
//
type CfnIntegrationProps struct {
	// The type of integration.
	//
	// Immutable after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-integration.html#cfn-cloudwatch-integration-integrationtype
	//
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
	// The customer-supplied name of the integration.
	//
	// Unique per account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-integration.html#cfn-cloudwatch-integration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Write-only API-key credential material for the integration.
	//
	// Stored encrypted and never returned on read.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-integration.html#cfn-cloudwatch-integration-credential
	//
	Credential interface{} `field:"optional" json:"credential" yaml:"credential"`
	// Type-specific configuration attributes (max 50 entries).
	//
	// EXTERNAL_AGENT requires 'catalogId'. On update, supplied attributes merge into the existing map; omitting a key does not remove it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-integration.html#cfn-cloudwatch-integration-integrationattributes
	//
	IntegrationAttributes interface{} `field:"optional" json:"integrationAttributes" yaml:"integrationAttributes"`
	// IAM role ARN the integration assumes.
	//
	// Required for AWS_INTEGRATION (a valid, service-assumable IAM role ARN). Not used by AWS_CONFIG_SLREC or EXTERNAL_AGENT.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-integration.html#cfn-cloudwatch-integration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Tags associated with the integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-integration.html#cfn-cloudwatch-integration-tags
	//
	Tags *[]*CfnIntegration_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

