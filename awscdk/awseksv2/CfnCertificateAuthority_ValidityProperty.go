package awseksv2


// The validity period of the certificate authority.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   validityProperty := &ValidityProperty{
//   	NotAfter: jsii.String("notAfter"),
//   	NotBefore: jsii.String("notBefore"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-certificateauthority-validity.html
//
type CfnCertificateAuthority_ValidityProperty struct {
	// The end of the validity period for the certificate authority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-certificateauthority-validity.html#cfn-eks-certificateauthority-validity-notafter
	//
	NotAfter *string `field:"optional" json:"notAfter" yaml:"notAfter"`
	// The start of the validity period for the certificate authority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-certificateauthority-validity.html#cfn-eks-certificateauthority-validity-notbefore
	//
	NotBefore *string `field:"optional" json:"notBefore" yaml:"notBefore"`
}

