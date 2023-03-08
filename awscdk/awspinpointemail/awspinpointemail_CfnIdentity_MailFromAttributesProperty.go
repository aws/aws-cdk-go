package awspinpointemail


// A list of attributes that are associated with a MAIL FROM domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mailFromAttributesProperty := &mailFromAttributesProperty{
//   	behaviorOnMxFailure: jsii.String("behaviorOnMxFailure"),
//   	mailFromDomain: jsii.String("mailFromDomain"),
//   }
//
type CfnIdentity_MailFromAttributesProperty struct {
	// The action that Amazon Pinpoint to takes if it can't read the required MX record for a custom MAIL FROM domain.
	//
	// When you set this value to `UseDefaultValue` , Amazon Pinpoint uses *amazonses.com* as the MAIL FROM domain. When you set this value to `RejectMessage` , Amazon Pinpoint returns a `MailFromDomainNotVerified` error, and doesn't attempt to deliver the email.
	//
	// These behaviors are taken when the custom MAIL FROM domain configuration is in the `Pending` , `Failed` , and `TemporaryFailure` states.
	BehaviorOnMxFailure *string `field:"optional" json:"behaviorOnMxFailure" yaml:"behaviorOnMxFailure"`
	// The name of a domain that an email identity uses as a custom MAIL FROM domain.
	MailFromDomain *string `field:"optional" json:"mailFromDomain" yaml:"mailFromDomain"`
}

