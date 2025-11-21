package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var default_ interface{}
//
//   maintenanceConfigurationProperty := &MaintenanceConfigurationProperty{
//   	Default: default_,
//   	PreferredDayTime: &PreferredDayTimeMaintenanceConfigurationProperty{
//   		Day: jsii.String("day"),
//   		Time: jsii.String("time"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-maintenanceconfiguration.html
//
type CfnRouterInput_MaintenanceConfigurationProperty struct {
	// Configuration settings for default maintenance scheduling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-maintenanceconfiguration.html#cfn-mediaconnect-routerinput-maintenanceconfiguration-default
	//
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// Configuration for preferred day and time maintenance settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-maintenanceconfiguration.html#cfn-mediaconnect-routerinput-maintenanceconfiguration-preferreddaytime
	//
	PreferredDayTime interface{} `field:"optional" json:"preferredDayTime" yaml:"preferredDayTime"`
}

