package awstransfer


// A reference to a Agreement resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agreementReference := &AgreementReference{
//   	AgreementArn: jsii.String("agreementArn"),
//   	AgreementId: jsii.String("agreementId"),
//   	ServerId: jsii.String("serverId"),
//   }
//
type AgreementReference struct {
	// The ARN of the Agreement resource.
	AgreementArn *string `field:"required" json:"agreementArn" yaml:"agreementArn"`
	// The AgreementId of the Agreement resource.
	AgreementId *string `field:"required" json:"agreementId" yaml:"agreementId"`
	// The ServerId of the Agreement resource.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
}

