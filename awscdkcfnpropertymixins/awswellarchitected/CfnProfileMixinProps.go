package awswellarchitected

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnProfilePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnProfileMixinProps := &CfnProfileMixinProps{
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
type CfnProfileMixinProps struct {
	// The profile description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-profile.html#cfn-wellarchitected-profile-profiledescription
	//
	ProfileDescription *string `field:"optional" json:"profileDescription" yaml:"profileDescription"`
	// The name of the profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-profile.html#cfn-wellarchitected-profile-profilename
	//
	ProfileName *string `field:"optional" json:"profileName" yaml:"profileName"`
	// The profile questions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-profile.html#cfn-wellarchitected-profile-profilequestions
	//
	ProfileQuestions interface{} `field:"optional" json:"profileQuestions" yaml:"profileQuestions"`
	// The tags assigned to the profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-profile.html#cfn-wellarchitected-profile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

