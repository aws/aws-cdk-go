package mixinsawslex


// A key/value pair representing session-specific context information.
//
// It contains application information passed between Amazon Lex V2 and a client application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sessionAttributeProperty := &SessionAttributeProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-sessionattribute.html
//
type CfnBotPropsMixin_SessionAttributeProperty struct {
	// The name of the session attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-sessionattribute.html#cfn-lex-bot-sessionattribute-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The session-specific context information for the session attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-sessionattribute.html#cfn-lex-bot-sessionattribute-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

