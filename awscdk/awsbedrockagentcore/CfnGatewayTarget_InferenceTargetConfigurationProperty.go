package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceTargetConfigurationProperty := &InferenceTargetConfigurationProperty{
//   	Connector: &InferenceConnectorTargetConfigurationProperty{
//   		Source: &InferenceConnectorSourceProperty{
//   			ConnectorId: jsii.String("connectorId"),
//   		},
//   	},
//   	Provider: &InferenceProviderTargetConfigurationProperty{
//   		Endpoint: jsii.String("endpoint"),
//
//   		// the properties below are optional
//   		ModelMapping: &ModelMappingProperty{
//   			ProviderPrefix: &ProviderPrefixProperty{
//   				Separator: jsii.String("separator"),
//   				Strip: jsii.Boolean(false),
//   			},
//   		},
//   		Operations: []interface{}{
//   			&InferenceOperationConfigurationProperty{
//   				Path: jsii.String("path"),
//
//   				// the properties below are optional
//   				Models: []interface{}{
//   					&ModelEntryProperty{
//   						Model: jsii.String("model"),
//   					},
//   				},
//   				ProviderPath: jsii.String("providerPath"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferencetargetconfiguration.html
//
type CfnGatewayTarget_InferenceTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferencetargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferencetargetconfiguration-connector
	//
	Connector interface{} `field:"optional" json:"connector" yaml:"connector"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferencetargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferencetargetconfiguration-provider
	//
	Provider interface{} `field:"optional" json:"provider" yaml:"provider"`
}

