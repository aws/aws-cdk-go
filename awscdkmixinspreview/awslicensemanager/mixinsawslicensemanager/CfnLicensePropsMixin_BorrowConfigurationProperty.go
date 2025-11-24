package mixinsawslicensemanager


// Details about a borrow configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   borrowConfigurationProperty := &BorrowConfigurationProperty{
//   	AllowEarlyCheckIn: jsii.Boolean(false),
//   	MaxTimeToLiveInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-borrowconfiguration.html
//
type CfnLicensePropsMixin_BorrowConfigurationProperty struct {
	// Indicates whether early check-ins are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-borrowconfiguration.html#cfn-licensemanager-license-borrowconfiguration-allowearlycheckin
	//
	AllowEarlyCheckIn interface{} `field:"optional" json:"allowEarlyCheckIn" yaml:"allowEarlyCheckIn"`
	// Maximum time for the borrow configuration, in minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-borrowconfiguration.html#cfn-licensemanager-license-borrowconfiguration-maxtimetoliveinminutes
	//
	MaxTimeToLiveInMinutes *float64 `field:"optional" json:"maxTimeToLiveInMinutes" yaml:"maxTimeToLiveInMinutes"`
}

