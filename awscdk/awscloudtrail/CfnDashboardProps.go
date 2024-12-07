package awscloudtrail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDashboard`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDashboardProps := &CfnDashboardProps{
//   	Name: jsii.String("name"),
//   	RefreshSchedule: &RefreshScheduleProperty{
//   		Frequency: &FrequencyProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   		Status: jsii.String("status"),
//   		TimeOfDay: jsii.String("timeOfDay"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TerminationProtectionEnabled: jsii.Boolean(false),
//   	Widgets: []interface{}{
//   		&WidgetProperty{
//   			QueryStatement: jsii.String("queryStatement"),
//
//   			// the properties below are optional
//   			QueryParameters: []*string{
//   				jsii.String("queryParameters"),
//   			},
//   			ViewProperties: map[string]*string{
//   				"viewPropertiesKey": jsii.String("viewProperties"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-dashboard.html
//
type CfnDashboardProps struct {
	// The name of the dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-dashboard.html#cfn-cloudtrail-dashboard-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Configures the automatic refresh schedule for the dashboard.
	//
	// Includes the frequency unit (DAYS or HOURS) and value, as well as the status (ENABLED or DISABLED) of the refresh schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-dashboard.html#cfn-cloudtrail-dashboard-refreshschedule
	//
	RefreshSchedule interface{} `field:"optional" json:"refreshSchedule" yaml:"refreshSchedule"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-dashboard.html#cfn-cloudtrail-dashboard-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Indicates whether the dashboard is protected from termination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-dashboard.html#cfn-cloudtrail-dashboard-terminationprotectionenabled
	//
	TerminationProtectionEnabled interface{} `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
	// List of widgets on the dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-dashboard.html#cfn-cloudtrail-dashboard-widgets
	//
	Widgets interface{} `field:"optional" json:"widgets" yaml:"widgets"`
}

