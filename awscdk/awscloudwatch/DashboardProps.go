package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a CloudWatch Dashboard.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import cw "github.com/aws/aws-cdk-go/awscdk"
//
//
//   dashboard := cw.NewDashboard(stack, jsii.String("Dash"), &DashboardProps{
//   	DefaultInterval: cdk.duration_Days(jsii.Number(7)),
//   })
//
type DashboardProps struct {
	// Name of the dashboard.
	//
	// If set, must only contain alphanumerics, dash (-) and underscore (_).
	DashboardName *string `field:"optional" json:"dashboardName" yaml:"dashboardName"`
	// Interval duration for metrics.
	//
	// You can specify defaultInterval with the relative time(eg. cdk.Duration.days(7)).
	DefaultInterval awscdk.Duration `field:"optional" json:"defaultInterval" yaml:"defaultInterval"`
	// The end of the time range to use for each widget on the dashboard when the dashboard loads.
	//
	// If you specify a value for end, you must also specify a value for start.
	// Specify an absolute time in the ISO 8601 format. For example, 2018-12-17T06:00:00.000Z.
	End *string `field:"optional" json:"end" yaml:"end"`
	// Use this field to specify the period for the graphs when the dashboard loads.
	//
	// Specifying `Auto` causes the period of all graphs on the dashboard to automatically adapt to the time range of the dashboard.
	// Specifying `Inherit` ensures that the period set for each graph is always obeyed.
	PeriodOverride PeriodOverride `field:"optional" json:"periodOverride" yaml:"periodOverride"`
	// The start of the time range to use for each widget on the dashboard.
	//
	// You can specify start without specifying end to specify a relative time range that ends with the current time.
	// In this case, the value of start must begin with -P, and you can use M, H, D, W and M as abbreviations for
	// minutes, hours, days, weeks and months. For example, -PT8H shows the last 8 hours and -P3M shows the last three months.
	// You can also use start along with an end field, to specify an absolute time range.
	// When specifying an absolute time range, use the ISO 8601 format. For example, 2018-12-17T06:00:00.000Z.
	Start *string `field:"optional" json:"start" yaml:"start"`
	// Initial set of widgets on the dashboard.
	//
	// One array represents a row of widgets.
	Widgets *[]*[]IWidget `field:"optional" json:"widgets" yaml:"widgets"`
}

