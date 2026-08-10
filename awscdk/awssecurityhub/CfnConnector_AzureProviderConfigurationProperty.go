package awssecurityhub


// The configuration for connecting to an Azure environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   azureProviderConfigurationProperty := &AzureProviderConfigurationProperty{
//   	AwsConfigConnectorArn: jsii.String("awsConfigConnectorArn"),
//   	AzureRegions: []*string{
//   		jsii.String("azureRegions"),
//   	},
//   	ScopeConfiguration: &AzureScopeConfigurationProperty{
//   		ScopeType: jsii.String("scopeType"),
//
//   		// the properties below are optional
//   		ScopeValues: []*string{
//   			jsii.String("scopeValues"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connector-azureproviderconfiguration.html
//
type CfnConnector_AzureProviderConfigurationProperty struct {
	// The ARN of the multi-cloud configuration connector used to establish the connection to Azure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connector-azureproviderconfiguration.html#cfn-securityhub-connector-azureproviderconfiguration-awsconfigconnectorarn
	//
	AwsConfigConnectorArn *string `field:"required" json:"awsConfigConnectorArn" yaml:"awsConfigConnectorArn"`
	// The list of Azure regions to monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connector-azureproviderconfiguration.html#cfn-securityhub-connector-azureproviderconfiguration-azureregions
	//
	AzureRegions *[]*string `field:"required" json:"azureRegions" yaml:"azureRegions"`
	// The scope configuration for an Azure connector, defining the tenant or subscription scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connector-azureproviderconfiguration.html#cfn-securityhub-connector-azureproviderconfiguration-scopeconfiguration
	//
	ScopeConfiguration interface{} `field:"required" json:"scopeConfiguration" yaml:"scopeConfiguration"`
}

