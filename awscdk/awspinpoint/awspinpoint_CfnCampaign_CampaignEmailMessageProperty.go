package awspinpoint


// Specifies the content and "From" address for an email message that's sent to recipients of a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignEmailMessageProperty := &campaignEmailMessageProperty{
//   	body: jsii.String("body"),
//   	fromAddress: jsii.String("fromAddress"),
//   	htmlBody: jsii.String("htmlBody"),
//   	title: jsii.String("title"),
//   }
//
type CfnCampaign_CampaignEmailMessageProperty struct {
	// The body of the email for recipients whose email clients don't render HTML content.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The verified email address to send the email from.
	//
	// The default address is the `FromAddress` specified for the email channel for the application.
	FromAddress *string `field:"optional" json:"fromAddress" yaml:"fromAddress"`
	// The body of the email, in HTML format, for recipients whose email clients render HTML content.
	HtmlBody *string `field:"optional" json:"htmlBody" yaml:"htmlBody"`
	// The subject line, or title, of the email.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

