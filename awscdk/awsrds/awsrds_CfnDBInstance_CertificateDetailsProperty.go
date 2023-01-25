package awsrds


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateDetailsProperty := &certificateDetailsProperty{
//   	caIdentifier: jsii.String("caIdentifier"),
//   	validTill: jsii.String("validTill"),
//   }
//
type CfnDBInstance_CertificateDetailsProperty struct {
	// `CfnDBInstance.CertificateDetailsProperty.CAIdentifier`.
	CaIdentifier *string `field:"optional" json:"caIdentifier" yaml:"caIdentifier"`
	// `CfnDBInstance.CertificateDetailsProperty.ValidTill`.
	ValidTill *string `field:"optional" json:"validTill" yaml:"validTill"`
}

