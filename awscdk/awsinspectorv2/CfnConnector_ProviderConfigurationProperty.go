package awsinspectorv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   providerConfigurationProperty := &ProviderConfigurationProperty{
//   	Azure: &AzureProviderConfigurationProperty{
//   		AwsConfigConnectorArn: jsii.String("awsConfigConnectorArn"),
//   		AzureRegions: []*string{
//   			jsii.String("azureRegions"),
//   		},
//   		ScopeConfiguration: &AzureScopeConfigurationMapProperty{
//   			ContainerImageScanning: &ScopeConfigurationProperty{
//   				ScopeType: jsii.String("scopeType"),
//
//   				// the properties below are optional
//   				ScopeValues: []*string{
//   					jsii.String("scopeValues"),
//   				},
//   				State: jsii.String("state"),
//   				StateReason: jsii.String("stateReason"),
//   			},
//   			ServerlessScanning: &ScopeConfigurationProperty{
//   				ScopeType: jsii.String("scopeType"),
//
//   				// the properties below are optional
//   				ScopeValues: []*string{
//   					jsii.String("scopeValues"),
//   				},
//   				State: jsii.String("state"),
//   				StateReason: jsii.String("stateReason"),
//   			},
//   			VmScanning: &ScopeConfigurationProperty{
//   				ScopeType: jsii.String("scopeType"),
//
//   				// the properties below are optional
//   				ScopeValues: []*string{
//   					jsii.String("scopeValues"),
//   				},
//   				State: jsii.String("state"),
//   				StateReason: jsii.String("stateReason"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		AutoInstallVmScanner: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-providerconfiguration.html
//
type CfnConnector_ProviderConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-providerconfiguration.html#cfn-inspectorv2-connector-providerconfiguration-azure
	//
	Azure interface{} `field:"required" json:"azure" yaml:"azure"`
}

