package awsacmpca


// Defines the X.509 `CertificatePolicies` extension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyInformationProperty := &PolicyInformationProperty{
//   	CertPolicyId: jsii.String("certPolicyId"),
//
//   	// the properties below are optional
//   	PolicyQualifiers: []interface{}{
//   		&PolicyQualifierInfoProperty{
//   			PolicyQualifierId: jsii.String("policyQualifierId"),
//   			Qualifier: &QualifierProperty{
//   				CpsUri: jsii.String("cpsUri"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-policyinformation.html
//
type CfnCertificate_PolicyInformationProperty struct {
	// Specifies the object identifier (OID) of the certificate policy under which the certificate was issued.
	//
	// For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-policyinformation.html#cfn-acmpca-certificate-policyinformation-certpolicyid
	//
	CertPolicyId *string `field:"required" json:"certPolicyId" yaml:"certPolicyId"`
	// Modifies the given `CertPolicyId` with a qualifier.
	//
	// AWS Private CA supports the certification practice statement (CPS) qualifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-policyinformation.html#cfn-acmpca-certificate-policyinformation-policyqualifiers
	//
	PolicyQualifiers interface{} `field:"optional" json:"policyQualifiers" yaml:"policyQualifiers"`
}

