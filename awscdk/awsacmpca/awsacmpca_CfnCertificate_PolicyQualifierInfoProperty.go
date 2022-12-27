package awsacmpca


// Modifies the `CertPolicyId` of a `PolicyInformation` object with a qualifier.
//
// ACM Private CA supports the certification practice statement (CPS) qualifier.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyQualifierInfoProperty := &policyQualifierInfoProperty{
//   	policyQualifierId: jsii.String("policyQualifierId"),
//   	qualifier: &qualifierProperty{
//   		cpsUri: jsii.String("cpsUri"),
//   	},
//   }
//
type CfnCertificate_PolicyQualifierInfoProperty struct {
	// Identifies the qualifier modifying a `CertPolicyId` .
	PolicyQualifierId *string `field:"required" json:"policyQualifierId" yaml:"policyQualifierId"`
	// Defines the qualifier type.
	//
	// ACM Private CA supports the use of a URI for a CPS qualifier in this field.
	Qualifier interface{} `field:"required" json:"qualifier" yaml:"qualifier"`
}

