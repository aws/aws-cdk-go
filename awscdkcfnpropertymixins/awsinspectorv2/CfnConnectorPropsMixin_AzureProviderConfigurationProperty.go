package awsinspectorv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   azureProviderConfigurationProperty := &AzureProviderConfigurationProperty{
//   	AutoInstallVmScanner: jsii.Boolean(false),
//   	AwsConfigConnectorArn: jsii.String("awsConfigConnectorArn"),
//   	AzureRegions: []*string{
//   		jsii.String("azureRegions"),
//   	},
//   	ScopeConfiguration: &AzureScopeConfigurationMapProperty{
//   		ContainerImageScanning: &ScopeConfigurationProperty{
//   			ScopeType: jsii.String("scopeType"),
//   			ScopeValues: []*string{
//   				jsii.String("scopeValues"),
//   			},
//   			State: jsii.String("state"),
//   			StateReason: jsii.String("stateReason"),
//   		},
//   		ServerlessScanning: &ScopeConfigurationProperty{
//   			ScopeType: jsii.String("scopeType"),
//   			ScopeValues: []*string{
//   				jsii.String("scopeValues"),
//   			},
//   			State: jsii.String("state"),
//   			StateReason: jsii.String("stateReason"),
//   		},
//   		VmScanning: &ScopeConfigurationProperty{
//   			ScopeType: jsii.String("scopeType"),
//   			ScopeValues: []*string{
//   				jsii.String("scopeValues"),
//   			},
//   			State: jsii.String("state"),
//   			StateReason: jsii.String("stateReason"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-azureproviderconfiguration.html
//
type CfnConnectorPropsMixin_AzureProviderConfigurationProperty struct {
	// Whether to automatically install the VM scanner.
	//
	// Defaults to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-azureproviderconfiguration.html#cfn-inspectorv2-connector-azureproviderconfiguration-autoinstallvmscanner
	//
	// Default: - true.
	//
	AutoInstallVmScanner interface{} `field:"optional" json:"autoInstallVmScanner" yaml:"autoInstallVmScanner"`
	// The ARN of the AWS Config connector used for Azure resource discovery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-azureproviderconfiguration.html#cfn-inspectorv2-connector-azureproviderconfiguration-awsconfigconnectorarn
	//
	AwsConfigConnectorArn *string `field:"optional" json:"awsConfigConnectorArn" yaml:"awsConfigConnectorArn"`
	// List of Azure regions to scan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-azureproviderconfiguration.html#cfn-inspectorv2-connector-azureproviderconfiguration-azureregions
	//
	AzureRegions *[]*string `field:"optional" json:"azureRegions" yaml:"azureRegions"`
	// Defines which resource types to scan and at what scope level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-azureproviderconfiguration.html#cfn-inspectorv2-connector-azureproviderconfiguration-scopeconfiguration
	//
	ScopeConfiguration interface{} `field:"optional" json:"scopeConfiguration" yaml:"scopeConfiguration"`
}

