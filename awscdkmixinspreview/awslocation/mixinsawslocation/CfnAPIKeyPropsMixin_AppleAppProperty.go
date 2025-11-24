package mixinsawslocation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   appleAppProperty := &AppleAppProperty{
//   	BundleId: jsii.String("bundleId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-apikey-appleapp.html
//
type CfnAPIKeyPropsMixin_AppleAppProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-apikey-appleapp.html#cfn-location-apikey-appleapp-bundleid
	//
	BundleId *string `field:"optional" json:"bundleId" yaml:"bundleId"`
}

