package mixinsawsses


// Properties for CfnVdmAttributesPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVdmAttributesMixinProps := &CfnVdmAttributesMixinProps{
//   	DashboardAttributes: &DashboardAttributesProperty{
//   		EngagementMetrics: jsii.String("engagementMetrics"),
//   	},
//   	GuardianAttributes: &GuardianAttributesProperty{
//   		OptimizedSharedDelivery: jsii.String("optimizedSharedDelivery"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-vdmattributes.html
//
type CfnVdmAttributesMixinProps struct {
	// Specifies additional settings for your VDM configuration as applicable to the Dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-vdmattributes.html#cfn-ses-vdmattributes-dashboardattributes
	//
	DashboardAttributes interface{} `field:"optional" json:"dashboardAttributes" yaml:"dashboardAttributes"`
	// Specifies additional settings for your VDM configuration as applicable to the Guardian.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-vdmattributes.html#cfn-ses-vdmattributes-guardianattributes
	//
	GuardianAttributes interface{} `field:"optional" json:"guardianAttributes" yaml:"guardianAttributes"`
}

