package previewawsfismixins


// The CloudWatch dashboards to include as data sources in the experiment report.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudWatchDashboardProperty := &CloudWatchDashboardProperty{
//   	DashboardIdentifier: jsii.String("dashboardIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-cloudwatchdashboard.html
//
type CfnExperimentTemplatePropsMixin_CloudWatchDashboardProperty struct {
	// The Amazon Resource Name (ARN) of the CloudWatch dashboard to include in the experiment report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-cloudwatchdashboard.html#cfn-fis-experimenttemplate-cloudwatchdashboard-dashboardidentifier
	//
	DashboardIdentifier *string `field:"optional" json:"dashboardIdentifier" yaml:"dashboardIdentifier"`
}

