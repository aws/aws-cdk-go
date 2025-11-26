package previewawscleanroomsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPrivacyBudgetTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPrivacyBudgetTemplateMixinProps := &CfnPrivacyBudgetTemplateMixinProps{
//   	AutoRefresh: jsii.String("autoRefresh"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Parameters: &ParametersProperty{
//   		BudgetParameters: []interface{}{
//   			&BudgetParameterProperty{
//   				AutoRefresh: jsii.String("autoRefresh"),
//   				Budget: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Epsilon: jsii.Number(123),
//   		ResourceArn: jsii.String("resourceArn"),
//   		UsersNoisePerQuery: jsii.Number(123),
//   	},
//   	PrivacyBudgetType: jsii.String("privacyBudgetType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-privacybudgettemplate.html
//
type CfnPrivacyBudgetTemplateMixinProps struct {
	// How often the privacy budget refreshes.
	//
	// > If you plan to regularly bring new data into the collaboration, use `CALENDAR_MONTH` to automatically get a new privacy budget for the collaboration every calendar month. Choosing this option allows arbitrary amounts of information to be revealed about rows of the data when repeatedly queried across refreshes. Avoid choosing this if the same rows will be repeatedly queried between privacy budget refreshes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-privacybudgettemplate.html#cfn-cleanrooms-privacybudgettemplate-autorefresh
	//
	AutoRefresh *string `field:"optional" json:"autoRefresh" yaml:"autoRefresh"`
	// The identifier for a membership resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-privacybudgettemplate.html#cfn-cleanrooms-privacybudgettemplate-membershipidentifier
	//
	MembershipIdentifier *string `field:"optional" json:"membershipIdentifier" yaml:"membershipIdentifier"`
	// Specifies the epsilon and noise parameters for the privacy budget template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-privacybudgettemplate.html#cfn-cleanrooms-privacybudgettemplate-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Specifies the type of the privacy budget template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-privacybudgettemplate.html#cfn-cleanrooms-privacybudgettemplate-privacybudgettype
	//
	PrivacyBudgetType *string `field:"optional" json:"privacyBudgetType" yaml:"privacyBudgetType"`
	// An optional label that you can assign to a resource when you create it.
	//
	// Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-privacybudgettemplate.html#cfn-cleanrooms-privacybudgettemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

