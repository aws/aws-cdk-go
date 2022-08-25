package awslicensemanager


// Details about a borrow configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   borrowConfigurationProperty := &borrowConfigurationProperty{
//   	allowEarlyCheckIn: jsii.Boolean(false),
//   	maxTimeToLiveInMinutes: jsii.Number(123),
//   }
//
type CfnLicense_BorrowConfigurationProperty struct {
	// Indicates whether early check-ins are allowed.
	AllowEarlyCheckIn interface{} `field:"required" json:"allowEarlyCheckIn" yaml:"allowEarlyCheckIn"`
	// Maximum time for the borrow configuration, in minutes.
	MaxTimeToLiveInMinutes *float64 `field:"required" json:"maxTimeToLiveInMinutes" yaml:"maxTimeToLiveInMinutes"`
}

