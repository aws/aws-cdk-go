package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCustomPermissions`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomPermissionsProps := &CfnCustomPermissionsProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	CustomPermissionsName: jsii.String("customPermissionsName"),
//
//   	// the properties below are optional
//   	Capabilities: &CapabilitiesProperty{
//   		Action: jsii.String("action"),
//   		AddOrRunAnomalyDetectionForAnalyses: jsii.String("addOrRunAnomalyDetectionForAnalyses"),
//   		Analysis: jsii.String("analysis"),
//   		Automate: jsii.String("automate"),
//   		ChatAgent: jsii.String("chatAgent"),
//   		CreateAndUpdateDashboardEmailReports: jsii.String("createAndUpdateDashboardEmailReports"),
//   		CreateAndUpdateDatasets: jsii.String("createAndUpdateDatasets"),
//   		CreateAndUpdateDataSources: jsii.String("createAndUpdateDataSources"),
//   		CreateAndUpdateThemes: jsii.String("createAndUpdateThemes"),
//   		CreateAndUpdateThresholdAlerts: jsii.String("createAndUpdateThresholdAlerts"),
//   		CreateChatAgents: jsii.String("createChatAgents"),
//   		CreateSharedFolders: jsii.String("createSharedFolders"),
//   		CreateSpiceDataset: jsii.String("createSpiceDataset"),
//   		Dashboard: jsii.String("dashboard"),
//   		ExportToCsv: jsii.String("exportToCsv"),
//   		ExportToCsvInScheduledReports: jsii.String("exportToCsvInScheduledReports"),
//   		ExportToExcel: jsii.String("exportToExcel"),
//   		ExportToExcelInScheduledReports: jsii.String("exportToExcelInScheduledReports"),
//   		ExportToPdf: jsii.String("exportToPdf"),
//   		ExportToPdfInScheduledReports: jsii.String("exportToPdfInScheduledReports"),
//   		Flow: jsii.String("flow"),
//   		IncludeContentInScheduledReportsEmail: jsii.String("includeContentInScheduledReportsEmail"),
//   		KnowledgeBase: jsii.String("knowledgeBase"),
//   		PerformFlowUiTask: jsii.String("performFlowUiTask"),
//   		PrintReports: jsii.String("printReports"),
//   		PublishWithoutApproval: jsii.String("publishWithoutApproval"),
//   		RenameSharedFolders: jsii.String("renameSharedFolders"),
//   		Research: jsii.String("research"),
//   		ShareAnalyses: jsii.String("shareAnalyses"),
//   		ShareDashboards: jsii.String("shareDashboards"),
//   		ShareDatasets: jsii.String("shareDatasets"),
//   		ShareDataSources: jsii.String("shareDataSources"),
//   		Space: jsii.String("space"),
//   		SubscribeDashboardEmailReports: jsii.String("subscribeDashboardEmailReports"),
//   		UseAgentWebSearch: jsii.String("useAgentWebSearch"),
//   		UseBedrockModels: jsii.String("useBedrockModels"),
//   		ViewAccountSpiceCapacity: jsii.String("viewAccountSpiceCapacity"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-custompermissions.html
//
type CfnCustomPermissionsProps struct {
	// The ID of the AWS account that contains the custom permission configuration that you want to update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-custompermissions.html#cfn-quicksight-custompermissions-awsaccountid
	//
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The name of the custom permissions profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-custompermissions.html#cfn-quicksight-custompermissions-custompermissionsname
	//
	CustomPermissionsName *string `field:"required" json:"customPermissionsName" yaml:"customPermissionsName"`
	// A set of actions in the custom permissions profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-custompermissions.html#cfn-quicksight-custompermissions-capabilities
	//
	Capabilities interface{} `field:"optional" json:"capabilities" yaml:"capabilities"`
	// The tags to associate with the custom permissions profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-custompermissions.html#cfn-quicksight-custompermissions-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

