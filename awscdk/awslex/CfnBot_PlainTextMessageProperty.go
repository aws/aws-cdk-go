package awslex


// Defines an ASCII text message to send to the user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   plainTextMessageProperty := &PlainTextMessageProperty{
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-plaintextmessage.html
//
type CfnBot_PlainTextMessageProperty struct {
	// The message to send to the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-plaintextmessage.html#cfn-lex-bot-plaintextmessage-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

