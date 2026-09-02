package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wafConfigurationProperty := &WafConfigurationProperty{
//   	FailureMode: jsii.String("failureMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-wafconfiguration.html
//
type CfnGateway_WafConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-wafconfiguration.html#cfn-bedrockagentcore-gateway-wafconfiguration-failuremode
	//
	FailureMode *string `field:"optional" json:"failureMode" yaml:"failureMode"`
}

