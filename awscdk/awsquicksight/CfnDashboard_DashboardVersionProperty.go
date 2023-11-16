package awsquicksight


// Dashboard version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardVersionProperty := &DashboardVersionProperty{
//   	Arn: jsii.String("arn"),
//   	CreatedTime: jsii.String("createdTime"),
//   	DataSetArns: []*string{
//   		jsii.String("dataSetArns"),
//   	},
//   	Description: jsii.String("description"),
//   	Errors: []interface{}{
//   		&DashboardErrorProperty{
//   			Message: jsii.String("message"),
//   			Type: jsii.String("type"),
//   			ViolatedEntities: []interface{}{
//   				&EntityProperty{
//   					Path: jsii.String("path"),
//   				},
//   			},
//   		},
//   	},
//   	Options: &AssetOptionsProperty{
//   		Timezone: jsii.String("timezone"),
//   		WeekStart: jsii.String("weekStart"),
//   	},
//   	Sheets: []interface{}{
//   		&SheetProperty{
//   			Name: jsii.String("name"),
//   			SheetId: jsii.String("sheetId"),
//   		},
//   	},
//   	SourceEntityArn: jsii.String("sourceEntityArn"),
//   	Status: jsii.String("status"),
//   	ThemeArn: jsii.String("themeArn"),
//   	VersionNumber: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardversion.html
//
type CfnDashboard_DashboardVersionProperty struct {
	// The Amazon Resource Name (ARN) of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardversion.html#cfn-quicksight-dashboard-dashboardversion-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The time that this dashboard version was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardversion.html#cfn-quicksight-dashboard-dashboardversion-createdtime
	//
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// The Amazon Resource Numbers (ARNs) for the datasets that are associated with this version of the dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardversion.html#cfn-quicksight-dashboard-dashboardversion-datasetarns
	//
	DataSetArns *[]*string `field:"optional" json:"dataSetArns" yaml:"dataSetArns"`
	// Description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardversion.html#cfn-quicksight-dashboard-dashboardversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Errors associated with this dashboard version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardversion.html#cfn-quicksight-dashboard-dashboardversion-errors
	//
	Errors interface{} `field:"optional" json:"errors" yaml:"errors"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardversion.html#cfn-quicksight-dashboard-dashboardversion-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// A list of the associated sheets with the unique identifier and name of each sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardversion.html#cfn-quicksight-dashboard-dashboardversion-sheets
	//
	Sheets interface{} `field:"optional" json:"sheets" yaml:"sheets"`
	// Source entity ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardversion.html#cfn-quicksight-dashboard-dashboardversion-sourceentityarn
	//
	SourceEntityArn *string `field:"optional" json:"sourceEntityArn" yaml:"sourceEntityArn"`
	// The HTTP status of the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardversion.html#cfn-quicksight-dashboard-dashboardversion-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The ARN of the theme associated with a version of the dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardversion.html#cfn-quicksight-dashboard-dashboardversion-themearn
	//
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
	// Version number for this version of the dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardversion.html#cfn-quicksight-dashboard-dashboardversion-versionnumber
	//
	VersionNumber *float64 `field:"optional" json:"versionNumber" yaml:"versionNumber"`
}

