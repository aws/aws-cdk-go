package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardOptionsProperty := &dashboardOptionsProperty{
//   	engagementMetrics: jsii.String("engagementMetrics"),
//   }
//
type CfnConfigurationSet_DashboardOptionsProperty struct {
	// `CfnConfigurationSet.DashboardOptionsProperty.EngagementMetrics`.
	EngagementMetrics *string `field:"required" json:"engagementMetrics" yaml:"engagementMetrics"`
}

