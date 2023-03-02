package awsquicksight

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
//   	awsAccountId: jsii.String("awsAccountId"),
//   	dashboardId: jsii.String("dashboardId"),
//   	sourceEntity: &dashboardSourceEntityProperty{
//   		sourceTemplate: &dashboardSourceTemplateProperty{
//   			arn: jsii.String("arn"),
//   			dataSetReferences: []interface{}{
//   				&dataSetReferenceProperty{
//   					dataSetArn: jsii.String("dataSetArn"),
//   					dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	dashboardPublishOptions: &dashboardPublishOptionsProperty{
//   		adHocFilteringOption: &adHocFilteringOptionProperty{
//   			availabilityStatus: jsii.String("availabilityStatus"),
//   		},
//   		exportToCsvOption: &exportToCSVOptionProperty{
//   			availabilityStatus: jsii.String("availabilityStatus"),
//   		},
//   		sheetControlsOption: &sheetControlsOptionProperty{
//   			visibilityState: jsii.String("visibilityState"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	parameters: &parametersProperty{
//   		dateTimeParameters: []interface{}{
//   			&dateTimeParameterProperty{
//   				name: jsii.String("name"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		decimalParameters: []interface{}{
//   			&decimalParameterProperty{
//   				name: jsii.String("name"),
//   				values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		integerParameters: []interface{}{
//   			&integerParameterProperty{
//   				name: jsii.String("name"),
//   				values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		stringParameters: []interface{}{
//   			&stringParameterProperty{
//   				name: jsii.String("name"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	themeArn: jsii.String("themeArn"),
//   	versionDescription: jsii.String("versionDescription"),
//   }
//
type CfnDashboardProps struct {
	// The ID of the AWS account where you want to create the dashboard.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The ID for the dashboard, also added to the IAM policy.
	DashboardId *string `field:"required" json:"dashboardId" yaml:"dashboardId"`
	// The entity that you are using as a source when you create the dashboard.
	//
	// In `SourceEntity` , you specify the type of object that you want to use. You can only create a dashboard from a template, so you use a `SourceTemplate` entity. If you need to create a dashboard from an analysis, first convert the analysis to a template by using the `CreateTemplate` API operation. For `SourceTemplate` , specify the Amazon Resource Name (ARN) of the source template. The `SourceTemplate` ARN can contain any AWS account; and any QuickSight-supported AWS Region .
	//
	// Use the `DataSetReferences` entity within `SourceTemplate` to list the replacement datasets for the placeholders listed in the original. The schema in each dataset must match its placeholder.
	SourceEntity interface{} `field:"required" json:"sourceEntity" yaml:"sourceEntity"`
	// Options for publishing the dashboard when you create it:.
	//
	// - `AvailabilityStatus` for `AdHocFilteringOption` - This status can be either `ENABLED` or `DISABLED` . When this is set to `DISABLED` , Amazon QuickSight disables the left filter pane on the published dashboard, which can be used for ad hoc (one-time) filtering. This option is `ENABLED` by default.
	// - `AvailabilityStatus` for `ExportToCSVOption` - This status can be either `ENABLED` or `DISABLED` . The visual option to export data to .CSV format isn't enabled when this is set to `DISABLED` . This option is `ENABLED` by default.
	// - `VisibilityState` for `SheetControlsOption` - This visibility state can be either `COLLAPSED` or `EXPANDED` . This option is `COLLAPSED` by default.
	DashboardPublishOptions interface{} `field:"optional" json:"dashboardPublishOptions" yaml:"dashboardPublishOptions"`
	// The display name of the dashboard.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parameters for the creation of the dashboard, which you want to use to override the default settings.
	//
	// A dashboard can have any type of parameters, and some parameters might accept multiple values.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A structure that contains the permissions of the dashboard.
	//
	// You can use this structure for granting permissions by providing a list of IAM action information for each principal ARN.
	//
	// To specify no permissions, omit the permissions list.
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the dashboard.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon Resource Name (ARN) of the theme that is being used for this dashboard.
	//
	// If you add a value for this field, it overrides the value that is used in the source entity. The theme ARN must exist in the same AWS account where you create the dashboard.
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
	// A description for the first version of the dashboard being created.
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

