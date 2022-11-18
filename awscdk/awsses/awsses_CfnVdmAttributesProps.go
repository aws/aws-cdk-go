package awsses


// Properties for defining a `CfnVdmAttributes`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVdmAttributesProps := &cfnVdmAttributesProps{
//   	dashboardAttributes: &dashboardAttributesProperty{
//   		engagementMetrics: jsii.String("engagementMetrics"),
//   	},
//   	guardianAttributes: &guardianAttributesProperty{
//   		optimizedSharedDelivery: jsii.String("optimizedSharedDelivery"),
//   	},
//   }
//
type CfnVdmAttributesProps struct {
	// `AWS::SES::VdmAttributes.DashboardAttributes`.
	DashboardAttributes interface{} `field:"optional" json:"dashboardAttributes" yaml:"dashboardAttributes"`
	// `AWS::SES::VdmAttributes.GuardianAttributes`.
	GuardianAttributes interface{} `field:"optional" json:"guardianAttributes" yaml:"guardianAttributes"`
}

