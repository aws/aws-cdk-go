package interfacesawsvpclattice


// A reference to a DomainVerification resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainVerificationReference := &DomainVerificationReference{
//   	DomainVerificationArn: jsii.String("domainVerificationArn"),
//   }
//
type DomainVerificationReference struct {
	// The Arn of the DomainVerification resource.
	DomainVerificationArn *string `field:"required" json:"domainVerificationArn" yaml:"domainVerificationArn"`
}

