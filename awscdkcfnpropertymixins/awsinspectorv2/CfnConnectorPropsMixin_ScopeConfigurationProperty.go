package awsinspectorv2


// Defines the scope of Azure resources to monitor for a specific resource type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   scopeConfigurationProperty := &ScopeConfigurationProperty{
//   	ScopeType: jsii.String("scopeType"),
//   	ScopeValues: []*string{
//   		jsii.String("scopeValues"),
//   	},
//   	State: jsii.String("state"),
//   	StateReason: jsii.String("stateReason"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-scopeconfiguration.html
//
type CfnConnectorPropsMixin_ScopeConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-scopeconfiguration.html#cfn-inspectorv2-connector-scopeconfiguration-scopetype
	//
	ScopeType *string `field:"optional" json:"scopeType" yaml:"scopeType"`
	// List of subscription IDs.
	//
	// Empty for TENANT scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-scopeconfiguration.html#cfn-inspectorv2-connector-scopeconfiguration-scopevalues
	//
	ScopeValues *[]*string `field:"optional" json:"scopeValues" yaml:"scopeValues"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-scopeconfiguration.html#cfn-inspectorv2-connector-scopeconfiguration-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// Reason for the current scope state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-connector-scopeconfiguration.html#cfn-inspectorv2-connector-scopeconfiguration-statereason
	//
	StateReason *string `field:"optional" json:"stateReason" yaml:"stateReason"`
}

