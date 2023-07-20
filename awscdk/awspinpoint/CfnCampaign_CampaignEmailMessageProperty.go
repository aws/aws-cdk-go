package awspinpoint


// Specifies the content and "From" address for an email message that's sent to recipients of a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignEmailMessageProperty := &CampaignEmailMessageProperty{
//   	Body: jsii.String("body"),
//   	FromAddress: jsii.String("fromAddress"),
//   	HtmlBody: jsii.String("htmlBody"),
//   	Title: jsii.String("title"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignemailmessage.html
//
type CfnCampaign_CampaignEmailMessageProperty struct {
	// The body of the email for recipients whose email clients don't render HTML content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignemailmessage.html#cfn-pinpoint-campaign-campaignemailmessage-body
	//
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The verified email address to send the email from.
	//
	// The default address is the `FromAddress` specified for the email channel for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignemailmessage.html#cfn-pinpoint-campaign-campaignemailmessage-fromaddress
	//
	FromAddress *string `field:"optional" json:"fromAddress" yaml:"fromAddress"`
	// The body of the email, in HTML format, for recipients whose email clients render HTML content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignemailmessage.html#cfn-pinpoint-campaign-campaignemailmessage-htmlbody
	//
	HtmlBody *string `field:"optional" json:"htmlBody" yaml:"htmlBody"`
	// The subject line, or title, of the email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignemailmessage.html#cfn-pinpoint-campaign-campaignemailmessage-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
}

