package awslicensemanager


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   licenseAssetGroupPropertyProperty := &LicenseAssetGroupPropertyProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetgroup-licenseassetgroupproperty.html
//
type CfnLicenseAssetGroup_LicenseAssetGroupPropertyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetgroup-licenseassetgroupproperty.html#cfn-licensemanager-licenseassetgroup-licenseassetgroupproperty-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetgroup-licenseassetgroupproperty.html#cfn-licensemanager-licenseassetgroup-licenseassetgroupproperty-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

