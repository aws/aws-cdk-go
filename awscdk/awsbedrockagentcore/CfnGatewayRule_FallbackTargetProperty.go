package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fallbackTargetProperty := &FallbackTargetProperty{
//   	TargetName: jsii.String("targetName"),
//
//   	// the properties below are optional
//   	FirstByteTimeoutMilliseconds: jsii.Number(123),
//   	Model: jsii.String("model"),
//   	TriggerOn: &FallbackTriggersProperty{
//   		Errors: []*string{
//   			jsii.String("errors"),
//   		},
//   		StatusCodes: []*string{
//   			jsii.String("statusCodes"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-fallbacktarget.html
//
type CfnGatewayRule_FallbackTargetProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-fallbacktarget.html#cfn-bedrockagentcore-gatewayrule-fallbacktarget-targetname
	//
	TargetName *string `field:"required" json:"targetName" yaml:"targetName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-fallbacktarget.html#cfn-bedrockagentcore-gatewayrule-fallbacktarget-firstbytetimeoutmilliseconds
	//
	FirstByteTimeoutMilliseconds *float64 `field:"optional" json:"firstByteTimeoutMilliseconds" yaml:"firstByteTimeoutMilliseconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-fallbacktarget.html#cfn-bedrockagentcore-gatewayrule-fallbacktarget-model
	//
	Model *string `field:"optional" json:"model" yaml:"model"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-fallbacktarget.html#cfn-bedrockagentcore-gatewayrule-fallbacktarget-triggeron
	//
	TriggerOn interface{} `field:"optional" json:"triggerOn" yaml:"triggerOn"`
}

