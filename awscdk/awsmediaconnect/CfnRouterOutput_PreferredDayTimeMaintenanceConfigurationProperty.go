package awsmediaconnect


// Configuration for preferred day and time maintenance settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   preferredDayTimeMaintenanceConfigurationProperty := &PreferredDayTimeMaintenanceConfigurationProperty{
//   	Day: jsii.String("day"),
//   	Time: jsii.String("time"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-preferreddaytimemaintenanceconfiguration.html
//
type CfnRouterOutput_PreferredDayTimeMaintenanceConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-preferreddaytimemaintenanceconfiguration.html#cfn-mediaconnect-routeroutput-preferreddaytimemaintenanceconfiguration-day
	//
	Day *string `field:"required" json:"day" yaml:"day"`
	// The preferred time for maintenance operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-preferreddaytimemaintenanceconfiguration.html#cfn-mediaconnect-routeroutput-preferreddaytimemaintenanceconfiguration-time
	//
	Time *string `field:"required" json:"time" yaml:"time"`
}

