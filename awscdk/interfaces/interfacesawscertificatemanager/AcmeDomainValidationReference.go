package interfacesawscertificatemanager


// A reference to a AcmeDomainValidation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acmeDomainValidationReference := &AcmeDomainValidationReference{
//   	AcmeDomainValidationArn: jsii.String("acmeDomainValidationArn"),
//   }
//
type AcmeDomainValidationReference struct {
	// The Arn of the AcmeDomainValidation resource.
	AcmeDomainValidationArn *string `field:"required" json:"acmeDomainValidationArn" yaml:"acmeDomainValidationArn"`
}

