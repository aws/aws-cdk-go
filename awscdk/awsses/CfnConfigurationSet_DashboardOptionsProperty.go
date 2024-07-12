package awsses


// An object containing additional settings for your VDM configuration as applicable to the Dashboard.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardOptionsProperty := &DashboardOptionsProperty{
//   	EngagementMetrics: jsii.String("engagementMetrics"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-dashboardoptions.html
//
type CfnConfigurationSet_DashboardOptionsProperty struct {
	// Specifies the status of your VDM engagement metrics collection. Can be one of the following:.
	//
	// - `ENABLED` – Amazon SES enables engagement metrics for the configuration set.
	// - `DISABLED` – Amazon SES disables engagement metrics for the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-dashboardoptions.html#cfn-ses-configurationset-dashboardoptions-engagementmetrics
	//
	EngagementMetrics *string `field:"required" json:"engagementMetrics" yaml:"engagementMetrics"`
}

