package previewawssecuritylakemixins


// Specify the AWS account ID and external ID that the subscriber will use to access source data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   subscriberIdentityProperty := &SubscriberIdentityProperty{
//   	ExternalId: jsii.String("externalId"),
//   	Principal: jsii.String("principal"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-subscriberidentity.html
//
type CfnSubscriberPropsMixin_SubscriberIdentityProperty struct {
	// The external ID is a unique identifier that the subscriber provides to you.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-subscriberidentity.html#cfn-securitylake-subscriber-subscriberidentity-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Principals can include accounts, users, roles, federated users, or AWS services.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-subscriberidentity.html#cfn-securitylake-subscriber-subscriberidentity-principal
	//
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
}

