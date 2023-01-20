package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardVersionProperty := &dashboardVersionProperty{
//   	arn: jsii.String("arn"),
//   	createdTime: jsii.String("createdTime"),
//   	dataSetArns: []*string{
//   		jsii.String("dataSetArns"),
//   	},
//   	description: jsii.String("description"),
//   	errors: []interface{}{
//   		&dashboardErrorProperty{
//   			message: jsii.String("message"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	sheets: []interface{}{
//   		&sheetProperty{
//   			name: jsii.String("name"),
//   			sheetId: jsii.String("sheetId"),
//   		},
//   	},
//   	sourceEntityArn: jsii.String("sourceEntityArn"),
//   	status: jsii.String("status"),
//   	themeArn: jsii.String("themeArn"),
//   	versionNumber: jsii.Number(123),
//   }
//
type CfnDashboard_DashboardVersionProperty struct {
	// `CfnDashboard.DashboardVersionProperty.Arn`.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// `CfnDashboard.DashboardVersionProperty.CreatedTime`.
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// `CfnDashboard.DashboardVersionProperty.DataSetArns`.
	DataSetArns *[]*string `field:"optional" json:"dataSetArns" yaml:"dataSetArns"`
	// `CfnDashboard.DashboardVersionProperty.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `CfnDashboard.DashboardVersionProperty.Errors`.
	Errors interface{} `field:"optional" json:"errors" yaml:"errors"`
	// `CfnDashboard.DashboardVersionProperty.Sheets`.
	Sheets interface{} `field:"optional" json:"sheets" yaml:"sheets"`
	// `CfnDashboard.DashboardVersionProperty.SourceEntityArn`.
	SourceEntityArn *string `field:"optional" json:"sourceEntityArn" yaml:"sourceEntityArn"`
	// `CfnDashboard.DashboardVersionProperty.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// `CfnDashboard.DashboardVersionProperty.ThemeArn`.
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
	// `CfnDashboard.DashboardVersionProperty.VersionNumber`.
	VersionNumber *float64 `field:"optional" json:"versionNumber" yaml:"versionNumber"`
}

