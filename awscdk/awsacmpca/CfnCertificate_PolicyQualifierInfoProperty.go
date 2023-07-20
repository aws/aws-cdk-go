package awsacmpca


// Modifies the `CertPolicyId` of a `PolicyInformation` object with a qualifier.
//
// AWS Private CA supports the certification practice statement (CPS) qualifier.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyQualifierInfoProperty := &PolicyQualifierInfoProperty{
//   	PolicyQualifierId: jsii.String("policyQualifierId"),
//   	Qualifier: &QualifierProperty{
//   		CpsUri: jsii.String("cpsUri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-policyqualifierinfo.html
//
type CfnCertificate_PolicyQualifierInfoProperty struct {
	// Identifies the qualifier modifying a `CertPolicyId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-policyqualifierinfo.html#cfn-acmpca-certificate-policyqualifierinfo-policyqualifierid
	//
	PolicyQualifierId *string `field:"required" json:"policyQualifierId" yaml:"policyQualifierId"`
	// Defines the qualifier type.
	//
	// AWS Private CA supports the use of a URI for a CPS qualifier in this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-policyqualifierinfo.html#cfn-acmpca-certificate-policyqualifierinfo-qualifier
	//
	Qualifier interface{} `field:"required" json:"qualifier" yaml:"qualifier"`
}

