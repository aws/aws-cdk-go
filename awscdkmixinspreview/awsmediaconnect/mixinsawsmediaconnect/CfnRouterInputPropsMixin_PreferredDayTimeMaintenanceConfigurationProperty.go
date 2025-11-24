package mixinsawsmediaconnect


// Configuration for preferred day and time maintenance settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   preferredDayTimeMaintenanceConfigurationProperty := &PreferredDayTimeMaintenanceConfigurationProperty{
//   	Day: jsii.String("day"),
//   	Time: jsii.String("time"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-preferreddaytimemaintenanceconfiguration.html
//
type CfnRouterInputPropsMixin_PreferredDayTimeMaintenanceConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-preferreddaytimemaintenanceconfiguration.html#cfn-mediaconnect-routerinput-preferreddaytimemaintenanceconfiguration-day
	//
	Day *string `field:"optional" json:"day" yaml:"day"`
	// The preferred time for maintenance operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-preferreddaytimemaintenanceconfiguration.html#cfn-mediaconnect-routerinput-preferreddaytimemaintenanceconfiguration-time
	//
	Time *string `field:"optional" json:"time" yaml:"time"`
}

