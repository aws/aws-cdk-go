package mixinsawsses


// An object containing additional settings for your VDM configuration as applicable to the Dashboard.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dashboardAttributesProperty := &DashboardAttributesProperty{
//   	EngagementMetrics: jsii.String("engagementMetrics"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-vdmattributes-dashboardattributes.html
//
type CfnVdmAttributesPropsMixin_DashboardAttributesProperty struct {
	// Specifies the status of your VDM engagement metrics collection. Can be one of the following:.
	//
	// - `ENABLED` – Amazon SES enables engagement metrics for your account.
	// - `DISABLED` – Amazon SES disables engagement metrics for your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-vdmattributes-dashboardattributes.html#cfn-ses-vdmattributes-dashboardattributes-engagementmetrics
	//
	EngagementMetrics *string `field:"optional" json:"engagementMetrics" yaml:"engagementMetrics"`
}

