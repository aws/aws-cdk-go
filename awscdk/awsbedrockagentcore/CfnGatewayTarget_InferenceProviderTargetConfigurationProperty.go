package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceProviderTargetConfigurationProperty := &InferenceProviderTargetConfigurationProperty{
//   	Endpoint: jsii.String("endpoint"),
//
//   	// the properties below are optional
//   	ModelMapping: &ModelMappingProperty{
//   		ProviderPrefix: &ProviderPrefixProperty{
//   			Separator: jsii.String("separator"),
//   			Strip: jsii.Boolean(false),
//   		},
//   	},
//   	Operations: []interface{}{
//   		&InferenceOperationConfigurationProperty{
//   			Path: jsii.String("path"),
//
//   			// the properties below are optional
//   			Models: []interface{}{
//   				&ModelEntryProperty{
//   					Model: jsii.String("model"),
//   				},
//   			},
//   			ProviderPath: jsii.String("providerPath"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceprovidertargetconfiguration.html
//
type CfnGatewayTarget_InferenceProviderTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceprovidertargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferenceprovidertargetconfiguration-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceprovidertargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferenceprovidertargetconfiguration-modelmapping
	//
	ModelMapping interface{} `field:"optional" json:"modelMapping" yaml:"modelMapping"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceprovidertargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferenceprovidertargetconfiguration-operations
	//
	Operations interface{} `field:"optional" json:"operations" yaml:"operations"`
}

