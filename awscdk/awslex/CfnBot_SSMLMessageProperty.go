package awslex


// Defines a Speech Synthesis Markup Language (SSML) prompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sSMLMessageProperty := &SSMLMessageProperty{
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-ssmlmessage.html
//
type CfnBot_SSMLMessageProperty struct {
	// The SSML text that defines the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-ssmlmessage.html#cfn-lex-bot-ssmlmessage-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

