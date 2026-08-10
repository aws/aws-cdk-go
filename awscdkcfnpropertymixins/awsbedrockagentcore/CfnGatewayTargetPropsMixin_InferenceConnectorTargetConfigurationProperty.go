package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inferenceConnectorTargetConfigurationProperty := &InferenceConnectorTargetConfigurationProperty{
//   	Source: &InferenceConnectorSourceProperty{
//   		ConnectorId: jsii.String("connectorId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceconnectortargetconfiguration.html
//
type CfnGatewayTargetPropsMixin_InferenceConnectorTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceconnectortargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-inferenceconnectortargetconfiguration-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

