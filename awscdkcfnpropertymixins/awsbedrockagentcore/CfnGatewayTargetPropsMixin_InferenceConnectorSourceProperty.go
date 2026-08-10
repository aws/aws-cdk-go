package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inferenceConnectorSourceProperty := &InferenceConnectorSourceProperty{
//   	ConnectorId: jsii.String("connectorId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceconnectorsource.html
//
type CfnGatewayTargetPropsMixin_InferenceConnectorSourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-inferenceconnectorsource.html#cfn-bedrockagentcore-gatewaytarget-inferenceconnectorsource-connectorid
	//
	ConnectorId *string `field:"optional" json:"connectorId" yaml:"connectorId"`
}

