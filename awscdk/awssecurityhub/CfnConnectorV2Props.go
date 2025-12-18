package awssecurityhub


// Properties for defining a `CfnConnectorV2`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectorV2Props := &CfnConnectorV2Props{
//   	Name: jsii.String("name"),
//   	Provider: &ProviderProperty{
//   		JiraCloud: &JiraCloudProperty{
//   			ProjectKey: jsii.String("projectKey"),
//
//   			// the properties below are optional
//   			AuthStatus: jsii.String("authStatus"),
//   			AuthUrl: jsii.String("authUrl"),
//   			CloudId: jsii.String("cloudId"),
//   			Domain: jsii.String("domain"),
//   		},
//   		ServiceNow: &ServiceNowProperty{
//   			InstanceName: jsii.String("instanceName"),
//   			SecretArn: jsii.String("secretArn"),
//
//   			// the properties below are optional
//   			AuthStatus: jsii.String("authStatus"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connectorv2.html
//
type CfnConnectorV2Props struct {
	// The unique name of the connectorV2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connectorv2.html#cfn-securityhub-connectorv2-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The third-party provider detail for a service configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connectorv2.html#cfn-securityhub-connectorv2-provider
	//
	Provider interface{} `field:"required" json:"provider" yaml:"provider"`
	// The description of the connectorV2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connectorv2.html#cfn-securityhub-connectorv2-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of KMS key used to encrypt secrets for the connectorV2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connectorv2.html#cfn-securityhub-connectorv2-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The tags to add to the connectorV2 when you create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connectorv2.html#cfn-securityhub-connectorv2-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

