package awsacmpca


// Defines the X.509 `CertificatePolicies` extension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyInformationProperty := &policyInformationProperty{
//   	certPolicyId: jsii.String("certPolicyId"),
//
//   	// the properties below are optional
//   	policyQualifiers: []interface{}{
//   		&policyQualifierInfoProperty{
//   			policyQualifierId: jsii.String("policyQualifierId"),
//   			qualifier: &qualifierProperty{
//   				cpsUri: jsii.String("cpsUri"),
//   			},
//   		},
//   	},
//   }
//
type CfnCertificate_PolicyInformationProperty struct {
	// Specifies the object identifier (OID) of the certificate policy under which the certificate was issued.
	//
	// For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
	CertPolicyId *string `field:"required" json:"certPolicyId" yaml:"certPolicyId"`
	// Modifies the given `CertPolicyId` with a qualifier.
	//
	// ACM Private CA supports the certification practice statement (CPS) qualifier.
	PolicyQualifiers interface{} `field:"optional" json:"policyQualifiers" yaml:"policyQualifiers"`
}

