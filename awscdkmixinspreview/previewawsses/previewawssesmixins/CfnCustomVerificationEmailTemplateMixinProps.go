package previewawssesmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCustomVerificationEmailTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCustomVerificationEmailTemplateMixinProps := &CfnCustomVerificationEmailTemplateMixinProps{
//   	FailureRedirectionUrl: jsii.String("failureRedirectionUrl"),
//   	FromEmailAddress: jsii.String("fromEmailAddress"),
//   	SuccessRedirectionUrl: jsii.String("successRedirectionUrl"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TemplateContent: jsii.String("templateContent"),
//   	TemplateName: jsii.String("templateName"),
//   	TemplateSubject: jsii.String("templateSubject"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html
//
type CfnCustomVerificationEmailTemplateMixinProps struct {
	// The URL that the recipient of the verification email is sent to if his or her address is not successfully verified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html#cfn-ses-customverificationemailtemplate-failureredirectionurl
	//
	FailureRedirectionUrl *string `field:"optional" json:"failureRedirectionUrl" yaml:"failureRedirectionUrl"`
	// The email address that the custom verification email is sent from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html#cfn-ses-customverificationemailtemplate-fromemailaddress
	//
	FromEmailAddress *string `field:"optional" json:"fromEmailAddress" yaml:"fromEmailAddress"`
	// The URL that the recipient of the verification email is sent to if his or her address is successfully verified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html#cfn-ses-customverificationemailtemplate-successredirectionurl
	//
	SuccessRedirectionUrl *string `field:"optional" json:"successRedirectionUrl" yaml:"successRedirectionUrl"`
	// The tags (keys and values) associated with the tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html#cfn-ses-customverificationemailtemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The content of the custom verification email.
	//
	// The total size of the email must be less than 10 MB. The message body may contain HTML, with some limitations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html#cfn-ses-customverificationemailtemplate-templatecontent
	//
	TemplateContent *string `field:"optional" json:"templateContent" yaml:"templateContent"`
	// The name of the custom verification email template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html#cfn-ses-customverificationemailtemplate-templatename
	//
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
	// The subject line of the custom verification email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html#cfn-ses-customverificationemailtemplate-templatesubject
	//
	TemplateSubject *string `field:"optional" json:"templateSubject" yaml:"templateSubject"`
}

