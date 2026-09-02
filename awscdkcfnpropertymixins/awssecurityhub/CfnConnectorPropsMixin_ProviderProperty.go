package awssecurityhub


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   providerProperty := &ProviderProperty{
//   	Azure: &AzureProviderConfigurationProperty{
//   		AwsConfigConnectorArn: jsii.String("awsConfigConnectorArn"),
//   		AzureRegions: []*string{
//   			jsii.String("azureRegions"),
//   		},
//   		ScopeConfiguration: &AzureScopeConfigurationProperty{
//   			ScopeType: jsii.String("scopeType"),
//   			ScopeValues: []*string{
//   				jsii.String("scopeValues"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connector-provider.html
//
type CfnConnectorPropsMixin_ProviderProperty struct {
	// The configuration for connecting to an Azure environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connector-provider.html#cfn-securityhub-connector-provider-azure
	//
	Azure interface{} `field:"optional" json:"azure" yaml:"azure"`
}

