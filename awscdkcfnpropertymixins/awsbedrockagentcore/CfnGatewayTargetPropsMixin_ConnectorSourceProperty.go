package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   connectorSourceProperty := &ConnectorSourceProperty{
//   	ConnectorId: jsii.String("connectorId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectorsource.html
//
type CfnGatewayTargetPropsMixin_ConnectorSourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectorsource.html#cfn-bedrockagentcore-gatewaytarget-connectorsource-connectorid
	//
	ConnectorId *string `field:"optional" json:"connectorId" yaml:"connectorId"`
}

