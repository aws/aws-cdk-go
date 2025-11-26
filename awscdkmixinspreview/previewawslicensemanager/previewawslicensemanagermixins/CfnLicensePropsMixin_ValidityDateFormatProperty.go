package previewawslicensemanagermixins


// Date and time range during which the license is valid, in ISO8601-UTC format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   validityDateFormatProperty := &ValidityDateFormatProperty{
//   	Begin: jsii.String("begin"),
//   	End: jsii.String("end"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-validitydateformat.html
//
type CfnLicensePropsMixin_ValidityDateFormatProperty struct {
	// Start of the time range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-validitydateformat.html#cfn-licensemanager-license-validitydateformat-begin
	//
	Begin *string `field:"optional" json:"begin" yaml:"begin"`
	// End of the time range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-validitydateformat.html#cfn-licensemanager-license-validitydateformat-end
	//
	End *string `field:"optional" json:"end" yaml:"end"`
}

