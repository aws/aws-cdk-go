package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDashboard`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDashboardProps := &CfnDashboardProps{
//   	DashboardDefinition: jsii.String("dashboardDefinition"),
//   	DashboardDescription: jsii.String("dashboardDescription"),
//   	DashboardName: jsii.String("dashboardName"),
//
//   	// the properties below are optional
//   	ProjectId: jsii.String("projectId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dashboard.html
//
type CfnDashboardProps struct {
	// The dashboard definition specified in a JSON literal.
	//
	// - AWS IoT SiteWise Monitor (Classic) see [Create dashboards ( AWS CLI )](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/create-dashboards-using-aws-cli.html)
	// - AWS IoT SiteWise Monitor (AI-aware) see [Create dashboards ( AWS CLI )](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/create-dashboards-ai-dashboard-cli.html)
	//
	// in the *AWS IoT SiteWise User Guide*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dashboard.html#cfn-iotsitewise-dashboard-dashboarddefinition
	//
	DashboardDefinition *string `field:"required" json:"dashboardDefinition" yaml:"dashboardDefinition"`
	// A description for the dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dashboard.html#cfn-iotsitewise-dashboard-dashboarddescription
	//
	DashboardDescription *string `field:"required" json:"dashboardDescription" yaml:"dashboardDescription"`
	// A friendly name for the dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dashboard.html#cfn-iotsitewise-dashboard-dashboardname
	//
	DashboardName *string `field:"required" json:"dashboardName" yaml:"dashboardName"`
	// The ID of the project in which to create the dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dashboard.html#cfn-iotsitewise-dashboard-projectid
	//
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// A list of key-value pairs that contain metadata for the dashboard.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dashboard.html#cfn-iotsitewise-dashboard-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

