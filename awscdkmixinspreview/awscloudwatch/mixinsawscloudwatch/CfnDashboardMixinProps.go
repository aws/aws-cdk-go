package mixinsawscloudwatch


// Properties for CfnDashboardPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDashboardMixinProps := &CfnDashboardMixinProps{
//   	DashboardBody: jsii.String("dashboardBody"),
//   	DashboardName: jsii.String("dashboardName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-dashboard.html
//
type CfnDashboardMixinProps struct {
	// The detailed information about the dashboard in JSON format, including the widgets to include and their location on the dashboard.
	//
	// This parameter is required.
	//
	// For more information about the syntax, see [Dashboard Body Structure and Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-dashboard.html#cfn-cloudwatch-dashboard-dashboardbody
	//
	DashboardBody *string `field:"optional" json:"dashboardBody" yaml:"dashboardBody"`
	// The name of the dashboard.
	//
	// The name must be between 1 and 255 characters. If you do not specify a name, one will be generated automatically.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-dashboard.html#cfn-cloudwatch-dashboard-dashboardname
	//
	DashboardName *string `field:"optional" json:"dashboardName" yaml:"dashboardName"`
}

