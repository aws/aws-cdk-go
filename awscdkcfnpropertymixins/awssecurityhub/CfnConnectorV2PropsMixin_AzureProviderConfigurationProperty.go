package awssecurityhub


// The configuration settings required to establish an integration between AWS Security Hub and Azure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   azureProviderConfigurationProperty := &AzureProviderConfigurationProperty{
//   	AwsConfigConnectorArn: jsii.String("awsConfigConnectorArn"),
//   	AzureRegions: []*string{
//   		jsii.String("azureRegions"),
//   	},
//   	ScopeConfiguration: &AzureScopeConfigurationProperty{
//   		ScopeType: jsii.String("scopeType"),
//   		ScopeValues: []*string{
//   			jsii.String("scopeValues"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-azureproviderconfiguration.html
//
type CfnConnectorV2PropsMixin_AzureProviderConfigurationProperty struct {
	// The ARN of the AWS Config connector used for the Azure integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-azureproviderconfiguration.html#cfn-securityhub-connectorv2-azureproviderconfiguration-awsconfigconnectorarn
	//
	AwsConfigConnectorArn *string `field:"optional" json:"awsConfigConnectorArn" yaml:"awsConfigConnectorArn"`
	// The list of Azure regions to include in the connector scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-azureproviderconfiguration.html#cfn-securityhub-connectorv2-azureproviderconfiguration-azureregions
	//
	AzureRegions *[]*string `field:"optional" json:"azureRegions" yaml:"azureRegions"`
	// The scope configuration for an Azure connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-azureproviderconfiguration.html#cfn-securityhub-connectorv2-azureproviderconfiguration-scopeconfiguration
	//
	ScopeConfiguration interface{} `field:"optional" json:"scopeConfiguration" yaml:"scopeConfiguration"`
}

