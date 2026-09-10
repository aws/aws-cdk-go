package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpConnectorTargetConfigurationProperty := &HttpConnectorTargetConfigurationProperty{
//   	Source: &HttpConnectorSourceProperty{
//   		ConnectorId: jsii.String("connectorId"),
//   	},
//
//   	// the properties below are optional
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-httpconnectortargetconfiguration.html
//
type CfnGatewayTarget_HttpConnectorTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-httpconnectortargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-httpconnectortargetconfiguration-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-httpconnectortargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-httpconnectortargetconfiguration-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

