package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCustomVerificationEmailTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomVerificationEmailTemplateProps := &CfnCustomVerificationEmailTemplateProps{
//   	FailureRedirectionUrl: jsii.String("failureRedirectionUrl"),
//   	FromEmailAddress: jsii.String("fromEmailAddress"),
//   	SuccessRedirectionUrl: jsii.String("successRedirectionUrl"),
//   	TemplateContent: jsii.String("templateContent"),
//   	TemplateName: jsii.String("templateName"),
//   	TemplateSubject: jsii.String("templateSubject"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html
//
type CfnCustomVerificationEmailTemplateProps struct {
	// The URL that the recipient of the verification email is sent to if his or her address is not successfully verified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html#cfn-ses-customverificationemailtemplate-failureredirectionurl
	//
	FailureRedirectionUrl *string `field:"required" json:"failureRedirectionUrl" yaml:"failureRedirectionUrl"`
	// The email address that the custom verification email is sent from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html#cfn-ses-customverificationemailtemplate-fromemailaddress
	//
	FromEmailAddress *string `field:"required" json:"fromEmailAddress" yaml:"fromEmailAddress"`
	// The URL that the recipient of the verification email is sent to if his or her address is successfully verified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html#cfn-ses-customverificationemailtemplate-successredirectionurl
	//
	SuccessRedirectionUrl *string `field:"required" json:"successRedirectionUrl" yaml:"successRedirectionUrl"`
	// The content of the custom verification email.
	//
	// The total size of the email must be less than 10 MB. The message body may contain HTML, with some limitations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html#cfn-ses-customverificationemailtemplate-templatecontent
	//
	TemplateContent *string `field:"required" json:"templateContent" yaml:"templateContent"`
	// The name of the custom verification email template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html#cfn-ses-customverificationemailtemplate-templatename
	//
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// The subject line of the custom verification email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html#cfn-ses-customverificationemailtemplate-templatesubject
	//
	TemplateSubject *string `field:"required" json:"templateSubject" yaml:"templateSubject"`
	// The tags (keys and values) associated with the tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-customverificationemailtemplate.html#cfn-ses-customverificationemailtemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

