package awssecurityhub


// Properties for defining a `CfnConnector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectorProps := &CfnConnectorProps{
//   	Name: jsii.String("name"),
//   	Provider: &ProviderProperty{
//   		Azure: &AzureProviderConfigurationProperty{
//   			AwsConfigConnectorArn: jsii.String("awsConfigConnectorArn"),
//   			AzureRegions: []*string{
//   				jsii.String("azureRegions"),
//   			},
//   			ScopeConfiguration: &AzureScopeConfigurationProperty{
//   				ScopeType: jsii.String("scopeType"),
//
//   				// the properties below are optional
//   				ScopeValues: []*string{
//   					jsii.String("scopeValues"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connector.html
//
type CfnConnectorProps struct {
	// The name of the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connector.html#cfn-securityhub-connector-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connector.html#cfn-securityhub-connector-provider
	//
	Provider interface{} `field:"required" json:"provider" yaml:"provider"`
	// The description of the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connector.html#cfn-securityhub-connector-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A key-value pair to associate with a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connector.html#cfn-securityhub-connector-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

