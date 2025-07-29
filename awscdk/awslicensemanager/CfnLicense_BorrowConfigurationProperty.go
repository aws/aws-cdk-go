package awslicensemanager


// Details about a borrow configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   borrowConfigurationProperty := &BorrowConfigurationProperty{
//   	AllowEarlyCheckIn: jsii.Boolean(false),
//   	MaxTimeToLiveInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-borrowconfiguration.html
//
type CfnLicense_BorrowConfigurationProperty struct {
	// Indicates whether early check-ins are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-borrowconfiguration.html#cfn-licensemanager-license-borrowconfiguration-allowearlycheckin
	//
	AllowEarlyCheckIn interface{} `field:"required" json:"allowEarlyCheckIn" yaml:"allowEarlyCheckIn"`
	// Maximum time for the borrow configuration, in minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-borrowconfiguration.html#cfn-licensemanager-license-borrowconfiguration-maxtimetoliveinminutes
	//
	MaxTimeToLiveInMinutes *float64 `field:"required" json:"maxTimeToLiveInMinutes" yaml:"maxTimeToLiveInMinutes"`
}

