package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardAttributesProperty := &dashboardAttributesProperty{
//   	engagementMetrics: jsii.String("engagementMetrics"),
//   }
//
type CfnVdmAttributes_DashboardAttributesProperty struct {
	// `CfnVdmAttributes.DashboardAttributesProperty.EngagementMetrics`.
	EngagementMetrics *string `field:"optional" json:"engagementMetrics" yaml:"engagementMetrics"`
}

