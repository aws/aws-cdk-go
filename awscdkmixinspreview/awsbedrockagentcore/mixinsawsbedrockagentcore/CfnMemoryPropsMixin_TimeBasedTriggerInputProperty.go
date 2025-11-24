package mixinsawsbedrockagentcore


// The memory trigger condition input for the time based trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timeBasedTriggerInputProperty := &TimeBasedTriggerInputProperty{
//   	IdleSessionTimeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-timebasedtriggerinput.html
//
type CfnMemoryPropsMixin_TimeBasedTriggerInputProperty struct {
	// The memory trigger condition input for the session timeout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-timebasedtriggerinput.html#cfn-bedrockagentcore-memory-timebasedtriggerinput-idlesessiontimeout
	//
	IdleSessionTimeout *float64 `field:"optional" json:"idleSessionTimeout" yaml:"idleSessionTimeout"`
}

