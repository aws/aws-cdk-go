package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamingConfigurationProperty := &StreamingConfigurationProperty{
//   	EnableResponseStreaming: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-streamingconfiguration.html
//
type CfnGateway_StreamingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-streamingconfiguration.html#cfn-bedrockagentcore-gateway-streamingconfiguration-enableresponsestreaming
	//
	EnableResponseStreaming interface{} `field:"optional" json:"enableResponseStreaming" yaml:"enableResponseStreaming"`
}

