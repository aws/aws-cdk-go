package awsinspectorv2


// Defines which resource types to scan and at what scope level.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   azureScopeConfigurationMapProperty := &AzureScopeConfigurationMapProperty{
//   	ContainerImageScanning: &ScopeConfigurationProperty{
//   		ScopeType: jsii.String("scopeType"),
//
//   		// the properties below are optional
//   		ScopeValues: []*string{
//   			jsii.String("scopeValues"),
//   		},
//   		State: jsii.String("state"),
//   		StateReason: jsii.String("stateReason"),
//   	},
//   	ServerlessScanning: &ScopeConfigurationProperty{
//   		ScopeType: jsii.String("scopeType"),
//
//   		// the properties below are optional
//   		ScopeValues: []*string{
//   			jsii.String("scopeValues"),
//   		},
//   		State: jsii.String("state"),
//   		StateReason: jsii.String("stateReason"),
//   	},
//   	VmScanning: &ScopeConfigurationProperty{
//   		ScopeType: jsii.String("scopeType"),
//
//   		// the properties below are optional
//   		ScopeValues: []*string{
//   			jsii.String("scopeValues"),
//   		},
//   		State: jsii.String("state"),
//   		StateReason: jsii.String("stateReason"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-azurescopeconfigurationmap.html
//
type CfnConnector_AzureScopeConfigurationMapProperty struct {
	// Defines the scope of Azure resources to monitor for a specific resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-azurescopeconfigurationmap.html#cfn-inspectorv2-connector-azurescopeconfigurationmap-containerimagescanning
	//
	ContainerImageScanning interface{} `field:"optional" json:"containerImageScanning" yaml:"containerImageScanning"`
	// Defines the scope of Azure resources to monitor for a specific resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-azurescopeconfigurationmap.html#cfn-inspectorv2-connector-azurescopeconfigurationmap-serverlessscanning
	//
	ServerlessScanning interface{} `field:"optional" json:"serverlessScanning" yaml:"serverlessScanning"`
	// Defines the scope of Azure resources to monitor for a specific resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-azurescopeconfigurationmap.html#cfn-inspectorv2-connector-azurescopeconfigurationmap-vmscanning
	//
	VmScanning interface{} `field:"optional" json:"vmScanning" yaml:"vmScanning"`
}

