package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDashboard`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDashboardProps := &cfnDashboardProps{
//   	dashboardDefinition: jsii.String("dashboardDefinition"),
//   	dashboardDescription: jsii.String("dashboardDescription"),
//   	dashboardName: jsii.String("dashboardName"),
//
//   	// the properties below are optional
//   	projectId: jsii.String("projectId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDashboardProps struct {
	// The dashboard definition specified in a JSON literal.
	//
	// For detailed information, see [Creating dashboards (CLI)](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/create-dashboards-using-aws-cli.html) in the *AWS IoT SiteWise User Guide* .
	DashboardDefinition *string `field:"required" json:"dashboardDefinition" yaml:"dashboardDefinition"`
	// A description for the dashboard.
	DashboardDescription *string `field:"required" json:"dashboardDescription" yaml:"dashboardDescription"`
	// A friendly name for the dashboard.
	DashboardName *string `field:"required" json:"dashboardName" yaml:"dashboardName"`
	// The ID of the project in which to create the dashboard.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// A list of key-value pairs that contain metadata for the dashboard.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

