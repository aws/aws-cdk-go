package awsinspectorv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   providerConfigurationProperty := &ProviderConfigurationProperty{
//   	Azure: &AzureProviderConfigurationProperty{
//   		AutoInstallVmScanner: jsii.Boolean(false),
//   		AwsConfigConnectorArn: jsii.String("awsConfigConnectorArn"),
//   		AzureRegions: []*string{
//   			jsii.String("azureRegions"),
//   		},
//   		ScopeConfiguration: &AzureScopeConfigurationMapProperty{
//   			ContainerImageScanning: &ScopeConfigurationProperty{
//   				ScopeType: jsii.String("scopeType"),
//   				ScopeValues: []*string{
//   					jsii.String("scopeValues"),
//   				},
//   				State: jsii.String("state"),
//   				StateReason: jsii.String("stateReason"),
//   			},
//   			ServerlessScanning: &ScopeConfigurationProperty{
//   				ScopeType: jsii.String("scopeType"),
//   				ScopeValues: []*string{
//   					jsii.String("scopeValues"),
//   				},
//   				State: jsii.String("state"),
//   				StateReason: jsii.String("stateReason"),
//   			},
//   			VmScanning: &ScopeConfigurationProperty{
//   				ScopeType: jsii.String("scopeType"),
//   				ScopeValues: []*string{
//   					jsii.String("scopeValues"),
//   				},
//   				State: jsii.String("state"),
//   				StateReason: jsii.String("stateReason"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-providerconfiguration.html
//
type CfnConnectorPropsMixin_ProviderConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-providerconfiguration.html#cfn-inspectorv2-connector-providerconfiguration-azure
	//
	Azure interface{} `field:"optional" json:"azure" yaml:"azure"`
}

