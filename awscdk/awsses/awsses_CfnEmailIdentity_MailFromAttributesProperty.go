package awsses


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
type CfnEmailIdentity_MailFromAttributesProperty struct {
	// `CfnEmailIdentity.MailFromAttributesProperty.BehaviorOnMxFailure`.
	BehaviorOnMxFailure *string `field:"optional" json:"behaviorOnMxFailure" yaml:"behaviorOnMxFailure"`
	// `CfnEmailIdentity.MailFromAttributesProperty.MailFromDomain`.
	MailFromDomain *string `field:"optional" json:"mailFromDomain" yaml:"mailFromDomain"`
}

