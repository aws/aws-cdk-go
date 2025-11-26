package previewawsquicksightmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDashboardPropsMixin.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html
//
type CfnDashboardMixinProps struct {
	// The ID of the AWS account where you want to create the dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The ID for the dashboard, also added to the IAM policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-dashboardid
	//
	DashboardId *string `field:"optional" json:"dashboardId" yaml:"dashboardId"`
	// Options for publishing the dashboard when you create it:.
	//
	// - `AvailabilityStatus` for `AdHocFilteringOption` - This status can be either `ENABLED` or `DISABLED` . When this is set to `DISABLED` , Amazon Quick Sight disables the left filter pane on the published dashboard, which can be used for ad hoc (one-time) filtering. This option is `ENABLED` by default.
	// - `AvailabilityStatus` for `ExportToCSVOption` - This status can be either `ENABLED` or `DISABLED` . The visual option to export data to .CSV format isn't enabled when this is set to `DISABLED` . This option is `ENABLED` by default.
	// - `VisibilityState` for `SheetControlsOption` - This visibility state can be either `COLLAPSED` or `EXPANDED` . This option is `COLLAPSED` by default.
	// - `AvailabilityStatus` for `QuickSuiteActionsOption` - This status can be either `ENABLED` or `DISABLED` . Features related to Actions in Amazon Quick Suite on dashboards are disabled when this is set to `DISABLED` . This option is `DISABLED` by default.
	// - `AvailabilityStatus` for `ExecutiveSummaryOption` - This status can be either `ENABLED` or `DISABLED` . The option to build an executive summary is disabled when this is set to `DISABLED` . This option is `ENABLED` by default.
	// - `AvailabilityStatus` for `DataStoriesSharingOption` - This status can be either `ENABLED` or `DISABLED` . The option to share a data story is disabled when this is set to `DISABLED` . This option is `ENABLED` by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-dashboardpublishoptions
	//
	DashboardPublishOptions interface{} `field:"optional" json:"dashboardPublishOptions" yaml:"dashboardPublishOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-definition
	//
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-folderarns
	//
	FolderArns *[]*string `field:"optional" json:"folderArns" yaml:"folderArns"`
	// A list of analysis Amazon Resource Names (ARNs) to be linked to the dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-linkentities
	//
	LinkEntities *[]*string `field:"optional" json:"linkEntities" yaml:"linkEntities"`
	// A structure that contains the link sharing configurations that you want to apply overrides to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-linksharingconfiguration
	//
	LinkSharingConfiguration interface{} `field:"optional" json:"linkSharingConfiguration" yaml:"linkSharingConfiguration"`
	// The display name of the dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parameters for the creation of the dashboard, which you want to use to override the default settings.
	//
	// A dashboard can have any type of parameters, and some parameters might accept multiple values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A structure that contains the permissions of the dashboard.
	//
	// You can use this structure for granting permissions by providing a list of IAM action information for each principal ARN.
	//
	// To specify no permissions, omit the permissions list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// The entity that you are using as a source when you create the dashboard.
	//
	// In `SourceEntity` , you specify the type of object that you want to use. You can only create a dashboard from a template, so you use a `SourceTemplate` entity. If you need to create a dashboard from an analysis, first convert the analysis to a template by using the `CreateTemplate` API operation. For `SourceTemplate` , specify the Amazon Resource Name (ARN) of the source template. The `SourceTemplate` ARN can contain any AWS account; and any QuickSight-supported AWS Region .
	//
	// Use the `DataSetReferences` entity within `SourceTemplate` to list the replacement datasets for the placeholders listed in the original. The schema in each dataset must match its placeholder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-sourceentity
	//
	SourceEntity interface{} `field:"optional" json:"sourceEntity" yaml:"sourceEntity"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon Resource Name (ARN) of the theme that is being used for this dashboard.
	//
	// If you add a value for this field, it overrides the value that is used in the source entity. The theme ARN must exist in the same AWS account where you create the dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-themearn
	//
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
	// The option to relax the validation that is required to create and update analyses, dashboards, and templates with definition objects.
	//
	// When you set this value to `LENIENT` , validation is skipped for specific errors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-validationstrategy
	//
	ValidationStrategy interface{} `field:"optional" json:"validationStrategy" yaml:"validationStrategy"`
	// A description for the first version of the dashboard being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-versiondescription
	//
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

