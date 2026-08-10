package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceOperationConfigurationProperty := &InferenceOperationConfigurationProperty{
//   	Path: jsii.String("path"),
//
//   	// the properties below are optional
//   	Models: []interface{}{
//   		&ModelEntryProperty{
//   			Model: jsii.String("model"),
//   		},
//   	},
//   	ProviderPath: jsii.String("providerPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceoperationconfiguration.html
//
type CfnGatewayTarget_InferenceOperationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceoperationconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferenceoperationconfiguration-path
	//
	Path *string `field:"required" json:"path" yaml:"path"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceoperationconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferenceoperationconfiguration-models
	//
	Models interface{} `field:"optional" json:"models" yaml:"models"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceoperationconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferenceoperationconfiguration-providerpath
	//
	ProviderPath *string `field:"optional" json:"providerPath" yaml:"providerPath"`
}

