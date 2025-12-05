package previewawssecurityhubmixins


// Properties for CfnConnectorV2PropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConnectorV2MixinProps := &CfnConnectorV2MixinProps{
//   	Description: jsii.String("description"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Name: jsii.String("name"),
//   	Provider: &ProviderProperty{
//   		JiraCloud: &JiraCloudProperty{
//   			AuthStatus: jsii.String("authStatus"),
//   			AuthUrl: jsii.String("authUrl"),
//   			CloudId: jsii.String("cloudId"),
//   			Domain: jsii.String("domain"),
//   			ProjectKey: jsii.String("projectKey"),
//   		},
//   		ServiceNow: &ServiceNowProperty{
//   			AuthStatus: jsii.String("authStatus"),
//   			InstanceName: jsii.String("instanceName"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connectorv2.html
//
type CfnConnectorV2MixinProps struct {
	// A description of the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connectorv2.html#cfn-securityhub-connectorv2-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of KMS key used for the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connectorv2.html#cfn-securityhub-connectorv2-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The name of the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connectorv2.html#cfn-securityhub-connectorv2-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The provider configuration of the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connectorv2.html#cfn-securityhub-connectorv2-provider
	//
	Provider interface{} `field:"optional" json:"provider" yaml:"provider"`
	// A key-value pair to associate with a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connectorv2.html#cfn-securityhub-connectorv2-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

