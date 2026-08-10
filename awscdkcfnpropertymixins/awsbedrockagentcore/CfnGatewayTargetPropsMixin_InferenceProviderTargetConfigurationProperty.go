package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inferenceProviderTargetConfigurationProperty := &InferenceProviderTargetConfigurationProperty{
//   	Endpoint: jsii.String("endpoint"),
//   	ModelMapping: &ModelMappingProperty{
//   		ProviderPrefix: &ProviderPrefixProperty{
//   			Separator: jsii.String("separator"),
//   			Strip: jsii.Boolean(false),
//   		},
//   	},
//   	Operations: []interface{}{
//   		&InferenceOperationConfigurationProperty{
//   			Models: []interface{}{
//   				&ModelEntryProperty{
//   					Model: jsii.String("model"),
//   				},
//   			},
//   			Path: jsii.String("path"),
//   			ProviderPath: jsii.String("providerPath"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceprovidertargetconfiguration.html
//
type CfnGatewayTargetPropsMixin_InferenceProviderTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceprovidertargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferenceprovidertargetconfiguration-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceprovidertargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferenceprovidertargetconfiguration-modelmapping
	//
	ModelMapping interface{} `field:"optional" json:"modelMapping" yaml:"modelMapping"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceprovidertargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferenceprovidertargetconfiguration-operations
	//
	Operations interface{} `field:"optional" json:"operations" yaml:"operations"`
}

