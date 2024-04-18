package awssecuritylake


// The AWS identity used to access your data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriberIdentityProperty := &SubscriberIdentityProperty{
//   	ExternalId: jsii.String("externalId"),
//   	Principal: jsii.String("principal"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-subscriberidentity.html
//
type CfnSubscriber_SubscriberIdentityProperty struct {
	// The external ID used to establish trust relationship with the AWS identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-subscriberidentity.html#cfn-securitylake-subscriber-subscriberidentity-externalid
	//
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// The AWS identity principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-subscriberidentity.html#cfn-securitylake-subscriber-subscriberidentity-principal
	//
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

