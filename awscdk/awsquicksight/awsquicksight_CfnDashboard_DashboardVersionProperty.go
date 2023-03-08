package awsquicksight


// Dashboard version.
//
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
	// The Amazon Resource Name (ARN) of the resource.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The time that this dashboard version was created.
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// The Amazon Resource Numbers (ARNs) for the datasets that are associated with this version of the dashboard.
	DataSetArns *[]*string `field:"optional" json:"dataSetArns" yaml:"dataSetArns"`
	// Description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Errors associated with this dashboard version.
	Errors interface{} `field:"optional" json:"errors" yaml:"errors"`
	// A list of the associated sheets with the unique identifier and name of each sheet.
	Sheets interface{} `field:"optional" json:"sheets" yaml:"sheets"`
	// Source entity ARN.
	SourceEntityArn *string `field:"optional" json:"sourceEntityArn" yaml:"sourceEntityArn"`
	// The HTTP status of the request.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The ARN of the theme associated with a version of the dashboard.
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
	// Version number for this version of the dashboard.
	VersionNumber *float64 `field:"optional" json:"versionNumber" yaml:"versionNumber"`
}

