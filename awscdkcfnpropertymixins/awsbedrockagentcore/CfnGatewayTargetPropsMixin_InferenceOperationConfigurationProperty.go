package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inferenceOperationConfigurationProperty := &InferenceOperationConfigurationProperty{
//   	Models: []interface{}{
//   		&ModelEntryProperty{
//   			Model: jsii.String("model"),
//   		},
//   	},
//   	Path: jsii.String("path"),
//   	ProviderPath: jsii.String("providerPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceoperationconfiguration.html
//
type CfnGatewayTargetPropsMixin_InferenceOperationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceoperationconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferenceoperationconfiguration-models
	//
	Models interface{} `field:"optional" json:"models" yaml:"models"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceoperationconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferenceoperationconfiguration-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceoperationconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferenceoperationconfiguration-providerpath
	//
	ProviderPath *string `field:"optional" json:"providerPath" yaml:"providerPath"`
}

