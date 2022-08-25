package awslicensemanager


// Details about a provisional configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionalConfigurationProperty := &provisionalConfigurationProperty{
//   	maxTimeToLiveInMinutes: jsii.Number(123),
//   }
//
type CfnLicense_ProvisionalConfigurationProperty struct {
	// Maximum time for the provisional configuration, in minutes.
	MaxTimeToLiveInMinutes *float64 `field:"required" json:"maxTimeToLiveInMinutes" yaml:"maxTimeToLiveInMinutes"`
}

