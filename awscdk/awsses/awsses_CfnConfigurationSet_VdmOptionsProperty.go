package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vdmOptionsProperty := &vdmOptionsProperty{
//   	dashboardOptions: &dashboardOptionsProperty{
//   		engagementMetrics: jsii.String("engagementMetrics"),
//   	},
//   	guardianOptions: &guardianOptionsProperty{
//   		optimizedSharedDelivery: jsii.String("optimizedSharedDelivery"),
//   	},
//   }
//
type CfnConfigurationSet_VdmOptionsProperty struct {
	// `CfnConfigurationSet.VdmOptionsProperty.DashboardOptions`.
	DashboardOptions interface{} `field:"optional" json:"dashboardOptions" yaml:"dashboardOptions"`
	// `CfnConfigurationSet.VdmOptionsProperty.GuardianOptions`.
	GuardianOptions interface{} `field:"optional" json:"guardianOptions" yaml:"guardianOptions"`
}

