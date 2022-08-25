package awsacmpca


// Defines a `PolicyInformation` qualifier.
//
// ACM Private CA supports the [certification practice statement (CPS) qualifier](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.4) defined in RFC 5280.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   qualifierProperty := &qualifierProperty{
//   	cpsUri: jsii.String("cpsUri"),
//   }
//
type CfnCertificate_QualifierProperty struct {
	// Contains a pointer to a certification practice statement (CPS) published by the CA.
	CpsUri *string `field:"required" json:"cpsUri" yaml:"cpsUri"`
}

