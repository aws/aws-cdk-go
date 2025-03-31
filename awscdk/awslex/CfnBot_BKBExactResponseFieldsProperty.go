package awslex


// Contains the names of the fields used for an exact response to the user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bKBExactResponseFieldsProperty := &BKBExactResponseFieldsProperty{
//   	AnswerField: jsii.String("answerField"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bkbexactresponsefields.html
//
type CfnBot_BKBExactResponseFieldsProperty struct {
	// The answer field used for an exact response from Bedrock Knowledge Store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bkbexactresponsefields.html#cfn-lex-bot-bkbexactresponsefields-answerfield
	//
	AnswerField *string `field:"optional" json:"answerField" yaml:"answerField"`
}

