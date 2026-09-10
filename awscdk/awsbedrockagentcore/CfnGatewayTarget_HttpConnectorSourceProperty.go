package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpConnectorSourceProperty := &HttpConnectorSourceProperty{
//   	ConnectorId: jsii.String("connectorId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-httpconnectorsource.html
//
type CfnGatewayTarget_HttpConnectorSourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-httpconnectorsource.html#cfn-bedrockagentcore-gatewaytarget-httpconnectorsource-connectorid
	//
	ConnectorId *string `field:"required" json:"connectorId" yaml:"connectorId"`
}

