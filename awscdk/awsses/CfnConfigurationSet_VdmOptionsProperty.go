package awsses


// The Virtual Deliverability Manager (VDM) options that apply to a configuration set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vdmOptionsProperty := &VdmOptionsProperty{
//   	DashboardOptions: &DashboardOptionsProperty{
//   		EngagementMetrics: jsii.String("engagementMetrics"),
//   	},
//   	GuardianOptions: &GuardianOptionsProperty{
//   		OptimizedSharedDelivery: jsii.String("optimizedSharedDelivery"),
//   	},
//   }
//
type CfnConfigurationSet_VdmOptionsProperty struct {
	// Settings for your VDM configuration as applicable to the Dashboard.
	DashboardOptions interface{} `field:"optional" json:"dashboardOptions" yaml:"dashboardOptions"`
	// Settings for your VDM configuration as applicable to the Guardian.
	GuardianOptions interface{} `field:"optional" json:"guardianOptions" yaml:"guardianOptions"`
}

