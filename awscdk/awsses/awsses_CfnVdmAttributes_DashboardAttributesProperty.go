package awsses


// Settings for your VDM configuration as applicable to the Dashboard.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardAttributesProperty := &DashboardAttributesProperty{
//   	EngagementMetrics: jsii.String("engagementMetrics"),
//   }
//
type CfnVdmAttributes_DashboardAttributesProperty struct {
	// Specifies the status of your VDM engagement metrics collection. Can be one of the following:.
	//
	// - `ENABLED` – Amazon SES enables engagement metrics for your account.
	// - `DISABLED` – Amazon SES disables engagement metrics for your account.
	EngagementMetrics *string `field:"optional" json:"engagementMetrics" yaml:"engagementMetrics"`
}

