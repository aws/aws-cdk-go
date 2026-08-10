package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inferenceTargetConfigurationProperty := &InferenceTargetConfigurationProperty{
//   	Connector: &InferenceConnectorTargetConfigurationProperty{
//   		Source: &InferenceConnectorSourceProperty{
//   			ConnectorId: jsii.String("connectorId"),
//   		},
//   	},
//   	Provider: &InferenceProviderTargetConfigurationProperty{
//   		Endpoint: jsii.String("endpoint"),
//   		ModelMapping: &ModelMappingProperty{
//   			ProviderPrefix: &ProviderPrefixProperty{
//   				Separator: jsii.String("separator"),
//   				Strip: jsii.Boolean(false),
//   			},
//   		},
//   		Operations: []interface{}{
//   			&InferenceOperationConfigurationProperty{
//   				Models: []interface{}{
//   					&ModelEntryProperty{
//   						Model: jsii.String("model"),
//   					},
//   				},
//   				Path: jsii.String("path"),
//   				ProviderPath: jsii.String("providerPath"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferencetargetconfiguration.html
//
type CfnGatewayTargetPropsMixin_InferenceTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferencetargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferencetargetconfiguration-connector
	//
	Connector interface{} `field:"optional" json:"connector" yaml:"connector"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferencetargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferencetargetconfiguration-provider
	//
	Provider interface{} `field:"optional" json:"provider" yaml:"provider"`
}

