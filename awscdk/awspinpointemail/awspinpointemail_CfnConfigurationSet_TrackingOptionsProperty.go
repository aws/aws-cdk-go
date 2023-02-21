package awspinpointemail


// An object that defines the tracking options for a configuration set.
//
// When you use Amazon Pinpoint to send an email, it contains an invisible image that's used to track when recipients open your email. If your email contains links, those links are changed slightly in order to track when recipients click them.
//
// These images and links include references to a domain operated by AWS . You can optionally configure Amazon Pinpoint to use a domain that you operate for these images and links.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trackingOptionsProperty := &trackingOptionsProperty{
//   	customRedirectDomain: jsii.String("customRedirectDomain"),
//   }
//
type CfnConfigurationSet_TrackingOptionsProperty struct {
	// The domain that you want to use for tracking open and click events.
	CustomRedirectDomain *string `field:"optional" json:"customRedirectDomain" yaml:"customRedirectDomain"`
}

