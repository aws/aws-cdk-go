package awslocation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   androidAppProperty := &AndroidAppProperty{
//   	CertificateFingerprint: jsii.String("certificateFingerprint"),
//   	Package: jsii.String("package"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-apikey-androidapp.html
//
type CfnAPIKey_AndroidAppProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-apikey-androidapp.html#cfn-location-apikey-androidapp-certificatefingerprint
	//
	CertificateFingerprint *string `field:"required" json:"certificateFingerprint" yaml:"certificateFingerprint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-apikey-androidapp.html#cfn-location-apikey-androidapp-package
	//
	Package *string `field:"required" json:"package" yaml:"package"`
}

