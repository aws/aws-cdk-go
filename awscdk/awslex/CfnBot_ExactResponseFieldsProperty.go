package awslex


// Contains the names of the fields used for an exact response to the user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   exactResponseFieldsProperty := &ExactResponseFieldsProperty{
//   	AnswerField: jsii.String("answerField"),
//   	QuestionField: jsii.String("questionField"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-exactresponsefields.html
//
type CfnBot_ExactResponseFieldsProperty struct {
	// The name of the field that contains the answer to the query made to the OpenSearch Service database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-exactresponsefields.html#cfn-lex-bot-exactresponsefields-answerfield
	//
	AnswerField *string `field:"optional" json:"answerField" yaml:"answerField"`
	// The name of the field that contains the query made to the OpenSearch Service database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-exactresponsefields.html#cfn-lex-bot-exactresponsefields-questionfield
	//
	QuestionField *string `field:"optional" json:"questionField" yaml:"questionField"`
}

