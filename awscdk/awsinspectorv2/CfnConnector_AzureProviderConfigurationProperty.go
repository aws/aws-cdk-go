package awsinspectorv2


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
//   	ScopeConfiguration: &AzureScopeConfigurationMapProperty{
//   		ContainerImageScanning: &ScopeConfigurationProperty{
//   			ScopeType: jsii.String("scopeType"),
//
//   			// the properties below are optional
//   			ScopeValues: []*string{
//   				jsii.String("scopeValues"),
//   			},
//   			State: jsii.String("state"),
//   			StateReason: jsii.String("stateReason"),
//   		},
//   		ServerlessScanning: &ScopeConfigurationProperty{
//   			ScopeType: jsii.String("scopeType"),
//
//   			// the properties below are optional
//   			ScopeValues: []*string{
//   				jsii.String("scopeValues"),
//   			},
//   			State: jsii.String("state"),
//   			StateReason: jsii.String("stateReason"),
//   		},
//   		VmScanning: &ScopeConfigurationProperty{
//   			ScopeType: jsii.String("scopeType"),
//
//   			// the properties below are optional
//   			ScopeValues: []*string{
//   				jsii.String("scopeValues"),
//   			},
//   			State: jsii.String("state"),
//   			StateReason: jsii.String("stateReason"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AutoInstallVmScanner: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-azureproviderconfiguration.html
//
type CfnConnector_AzureProviderConfigurationProperty struct {
	// The ARN of the AWS Config connector used for Azure resource discovery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-azureproviderconfiguration.html#cfn-inspectorv2-connector-azureproviderconfiguration-awsconfigconnectorarn
	//
	AwsConfigConnectorArn *string `field:"required" json:"awsConfigConnectorArn" yaml:"awsConfigConnectorArn"`
	// List of Azure regions to scan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-azureproviderconfiguration.html#cfn-inspectorv2-connector-azureproviderconfiguration-azureregions
	//
	AzureRegions *[]*string `field:"required" json:"azureRegions" yaml:"azureRegions"`
	// Defines which resource types to scan and at what scope level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-azureproviderconfiguration.html#cfn-inspectorv2-connector-azureproviderconfiguration-scopeconfiguration
	//
	ScopeConfiguration interface{} `field:"required" json:"scopeConfiguration" yaml:"scopeConfiguration"`
	// Whether to automatically install the VM scanner.
	//
	// Defaults to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-azureproviderconfiguration.html#cfn-inspectorv2-connector-azureproviderconfiguration-autoinstallvmscanner
	//
	// Default: - true.
	//
	AutoInstallVmScanner interface{} `field:"optional" json:"autoInstallVmScanner" yaml:"autoInstallVmScanner"`
}

