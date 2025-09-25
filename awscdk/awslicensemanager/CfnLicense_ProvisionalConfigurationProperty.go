package awslicensemanager


// Details about a provisional configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionalConfigurationProperty := &ProvisionalConfigurationProperty{
//   	MaxTimeToLiveInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-provisionalconfiguration.html
//
type CfnLicense_ProvisionalConfigurationProperty struct {
	// Maximum time for the provisional configuration, in minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-provisionalconfiguration.html#cfn-licensemanager-license-provisionalconfiguration-maxtimetoliveinminutes
	//
	MaxTimeToLiveInMinutes *float64 `field:"required" json:"maxTimeToLiveInMinutes" yaml:"maxTimeToLiveInMinutes"`
}

