package awswellarchitected


// An update to a profile question.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   profileQuestionUpdateProperty := &ProfileQuestionUpdateProperty{
//   	QuestionId: jsii.String("questionId"),
//   	SelectedChoiceIds: []*string{
//   		jsii.String("selectedChoiceIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wellarchitected-profile-profilequestionupdate.html
//
type CfnProfile_ProfileQuestionUpdateProperty struct {
	// The ID of the question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wellarchitected-profile-profilequestionupdate.html#cfn-wellarchitected-profile-profilequestionupdate-questionid
	//
	QuestionId *string `field:"optional" json:"questionId" yaml:"questionId"`
	// The selected choices.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wellarchitected-profile-profilequestionupdate.html#cfn-wellarchitected-profile-profilequestionupdate-selectedchoiceids
	//
	SelectedChoiceIds *[]*string `field:"optional" json:"selectedChoiceIds" yaml:"selectedChoiceIds"`
}

