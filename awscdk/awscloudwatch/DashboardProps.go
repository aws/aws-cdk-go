package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a CloudWatch Dashboard.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   dashboard := cw.NewDashboard(this, jsii.String("Dash"), &DashboardProps{
//   	DefaultInterval: awscdk.Duration_Days(jsii.Number(7)),
//   	Variables: []iVariable{
//   		cw.NewDashboardVariable(&DashboardVariableOptions{
//   			Id: jsii.String("region2"),
//   			Type: cw.VariableType_PATTERN,
//   			Label: jsii.String("RegionPattern"),
//   			InputType: cw.VariableInputType_INPUT,
//   			Value: jsii.String("us-east-1"),
//   			DefaultValue: cw.DefaultValue_Value(jsii.String("us-east-1")),
//   			Visible: jsii.Boolean(true),
//   		}),
//   	},
//   })
//
type DashboardProps struct {
	// Name of the dashboard.
	//
	// If set, must only contain alphanumerics, dash (-) and underscore (_).
	// Default: - automatically generated name.
	//
	DashboardName *string `field:"optional" json:"dashboardName" yaml:"dashboardName"`
	// Interval duration for metrics. You can specify defaultInterval with the relative time(eg. cdk.Duration.days(7)).
	//
	// Both properties `defaultInterval` and `start` cannot be set at once.
	// Default: When the dashboard loads, the defaultInterval time will be the default time range.
	//
	DefaultInterval awscdk.Duration `field:"optional" json:"defaultInterval" yaml:"defaultInterval"`
	// The end of the time range to use for each widget on the dashboard when the dashboard loads.
	//
	// If you specify a value for end, you must also specify a value for start.
	// Specify an absolute time in the ISO 8601 format. For example, 2018-12-17T06:00:00.000Z.
	// Default: When the dashboard loads, the end date will be the current time.
	//
	End *string `field:"optional" json:"end" yaml:"end"`
	// Use this field to specify the period for the graphs when the dashboard loads.
	//
	// Specifying `Auto` causes the period of all graphs on the dashboard to automatically adapt to the time range of the dashboard.
	// Specifying `Inherit` ensures that the period set for each graph is always obeyed.
	// Default: Auto.
	//
	PeriodOverride PeriodOverride `field:"optional" json:"periodOverride" yaml:"periodOverride"`
	// The start of the time range to use for each widget on the dashboard.
	//
	// You can specify start without specifying end to specify a relative time range that ends with the current time.
	// In this case, the value of start must begin with -P, and you can use M, H, D, W and M as abbreviations for
	// minutes, hours, days, weeks and months. For example, -PT8H shows the last 8 hours and -P3M shows the last three months.
	// You can also use start along with an end field, to specify an absolute time range.
	// When specifying an absolute time range, use the ISO 8601 format. For example, 2018-12-17T06:00:00.000Z.
	//
	// Both properties `defaultInterval` and `start` cannot be set at once.
	// Default: When the dashboard loads, the start time will be the default time range.
	//
	Start *string `field:"optional" json:"start" yaml:"start"`
	// A list of dashboard variables.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_dashboard_variables.html#cloudwatch_dashboard_variables_types
	//
	// Default: - No variables.
	//
	Variables *[]IVariable `field:"optional" json:"variables" yaml:"variables"`
	// Initial set of widgets on the dashboard.
	//
	// One array represents a row of widgets.
	// Default: - No widgets.
	//
	Widgets *[]*[]IWidget `field:"optional" json:"widgets" yaml:"widgets"`
}

