package mixinsawslocation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   androidAppProperty := &AndroidAppProperty{
//   	CertificateFingerprint: jsii.String("certificateFingerprint"),
//   	Package: jsii.String("package"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-apikey-androidapp.html
//
type CfnAPIKeyPropsMixin_AndroidAppProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-apikey-androidapp.html#cfn-location-apikey-androidapp-certificatefingerprint
	//
	CertificateFingerprint *string `field:"optional" json:"certificateFingerprint" yaml:"certificateFingerprint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-apikey-androidapp.html#cfn-location-apikey-androidapp-package
	//
	Package *string `field:"optional" json:"package" yaml:"package"`
}

