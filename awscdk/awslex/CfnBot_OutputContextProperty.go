package awslex


// Describes a session context that is activated when an intent is fulfilled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputContextProperty := &OutputContextProperty{
//   	Name: jsii.String("name"),
//   	TimeToLiveInSeconds: jsii.Number(123),
//   	TurnsToLive: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-outputcontext.html
//
type CfnBot_OutputContextProperty struct {
	// The name of the output context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-outputcontext.html#cfn-lex-bot-outputcontext-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The amount of time, in seconds, that the output context should remain active.
	//
	// The time is figured from the first time the context is sent to the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-outputcontext.html#cfn-lex-bot-outputcontext-timetoliveinseconds
	//
	TimeToLiveInSeconds *float64 `field:"required" json:"timeToLiveInSeconds" yaml:"timeToLiveInSeconds"`
	// The number of conversation turns that the output context should remain active.
	//
	// The number of turns is counted from the first time that the context is sent to the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-outputcontext.html#cfn-lex-bot-outputcontext-turnstolive
	//
	TurnsToLive *float64 `field:"required" json:"turnsToLive" yaml:"turnsToLive"`
}

