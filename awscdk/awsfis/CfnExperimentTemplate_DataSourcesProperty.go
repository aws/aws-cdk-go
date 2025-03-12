package awsfis


// Describes the data sources for the experiment report.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourcesProperty := &DataSourcesProperty{
//   	CloudWatchDashboards: []interface{}{
//   		&CloudWatchDashboardProperty{
//   			DashboardIdentifier: jsii.String("dashboardIdentifier"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-datasources.html
//
type CfnExperimentTemplate_DataSourcesProperty struct {
	// The CloudWatch dashboards to include as data sources in the experiment report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-datasources.html#cfn-fis-experimenttemplate-datasources-cloudwatchdashboards
	//
	CloudWatchDashboards interface{} `field:"optional" json:"cloudWatchDashboards" yaml:"cloudWatchDashboards"`
}

