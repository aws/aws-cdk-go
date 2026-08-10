package awswellarchitected

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProfileProps := &CfnProfileProps{
//   	ProfileDescription: jsii.String("profileDescription"),
//   	ProfileName: jsii.String("profileName"),
//   	ProfileQuestions: []interface{}{
//   		&ProfileQuestionUpdateProperty{
//   			QuestionId: jsii.String("questionId"),
//   			SelectedChoiceIds: []*string{
//   				jsii.String("selectedChoiceIds"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-profile.html
//
type CfnProfileProps struct {
	// The profile description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-profile.html#cfn-wellarchitected-profile-profiledescription
	//
	ProfileDescription *string `field:"required" json:"profileDescription" yaml:"profileDescription"`
	// The name of the profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-profile.html#cfn-wellarchitected-profile-profilename
	//
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	// The profile questions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-profile.html#cfn-wellarchitected-profile-profilequestions
	//
	ProfileQuestions interface{} `field:"required" json:"profileQuestions" yaml:"profileQuestions"`
	// The tags assigned to the profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-profile.html#cfn-wellarchitected-profile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

