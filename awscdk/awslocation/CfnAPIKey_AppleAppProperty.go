package awslocation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appleAppProperty := &AppleAppProperty{
//   	BundleId: jsii.String("bundleId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-apikey-appleapp.html
//
type CfnAPIKey_AppleAppProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-apikey-appleapp.html#cfn-location-apikey-appleapp-bundleid
	//
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
}

