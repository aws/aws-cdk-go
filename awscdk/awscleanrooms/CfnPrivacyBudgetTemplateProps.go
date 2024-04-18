package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPrivacyBudgetTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPrivacyBudgetTemplateProps := &CfnPrivacyBudgetTemplateProps{
//   	AutoRefresh: jsii.String("autoRefresh"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Parameters: &ParametersProperty{
//   		Epsilon: jsii.Number(123),
//   		UsersNoisePerQuery: jsii.Number(123),
//   	},
//   	PrivacyBudgetType: jsii.String("privacyBudgetType"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-privacybudgettemplate.html
//
type CfnPrivacyBudgetTemplateProps struct {
	// How often the privacy budget refreshes.
	//
	// > If you plan to regularly bring new data into the collaboration, use `CALENDAR_MONTH` to automatically get a new privacy budget for the collaboration every calendar month. Choosing this option allows arbitrary amounts of information to be revealed about rows of the data when repeatedly queried across refreshes. Avoid choosing this if the same rows will be repeatedly queried between privacy budget refreshes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-privacybudgettemplate.html#cfn-cleanrooms-privacybudgettemplate-autorefresh
	//
	AutoRefresh *string `field:"required" json:"autoRefresh" yaml:"autoRefresh"`
	// The identifier for a membership resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-privacybudgettemplate.html#cfn-cleanrooms-privacybudgettemplate-membershipidentifier
	//
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
	// Specifies the epislon and noise parameters for the privacy budget template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-privacybudgettemplate.html#cfn-cleanrooms-privacybudgettemplate-parameters
	//
	Parameters interface{} `field:"required" json:"parameters" yaml:"parameters"`
	// Specifies the type of the privacy budget template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-privacybudgettemplate.html#cfn-cleanrooms-privacybudgettemplate-privacybudgettype
	//
	PrivacyBudgetType *string `field:"required" json:"privacyBudgetType" yaml:"privacyBudgetType"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms privacy budget template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-privacybudgettemplate.html#cfn-cleanrooms-privacybudgettemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

