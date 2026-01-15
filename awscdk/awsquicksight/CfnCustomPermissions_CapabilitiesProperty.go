package awsquicksight


// A set of actions that correspond to Amazon Quick Sight permissions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capabilitiesProperty := &CapabilitiesProperty{
//   	Action: jsii.String("action"),
//   	AddOrRunAnomalyDetectionForAnalyses: jsii.String("addOrRunAnomalyDetectionForAnalyses"),
//   	Analysis: jsii.String("analysis"),
//   	Automate: jsii.String("automate"),
//   	ChatAgent: jsii.String("chatAgent"),
//   	CreateAndUpdateDashboardEmailReports: jsii.String("createAndUpdateDashboardEmailReports"),
//   	CreateAndUpdateDatasets: jsii.String("createAndUpdateDatasets"),
//   	CreateAndUpdateDataSources: jsii.String("createAndUpdateDataSources"),
//   	CreateAndUpdateThemes: jsii.String("createAndUpdateThemes"),
//   	CreateAndUpdateThresholdAlerts: jsii.String("createAndUpdateThresholdAlerts"),
//   	CreateChatAgents: jsii.String("createChatAgents"),
//   	CreateSharedFolders: jsii.String("createSharedFolders"),
//   	CreateSpiceDataset: jsii.String("createSpiceDataset"),
//   	Dashboard: jsii.String("dashboard"),
//   	ExportToCsv: jsii.String("exportToCsv"),
//   	ExportToCsvInScheduledReports: jsii.String("exportToCsvInScheduledReports"),
//   	ExportToExcel: jsii.String("exportToExcel"),
//   	ExportToExcelInScheduledReports: jsii.String("exportToExcelInScheduledReports"),
//   	ExportToPdf: jsii.String("exportToPdf"),
//   	ExportToPdfInScheduledReports: jsii.String("exportToPdfInScheduledReports"),
//   	Flow: jsii.String("flow"),
//   	IncludeContentInScheduledReportsEmail: jsii.String("includeContentInScheduledReportsEmail"),
//   	KnowledgeBase: jsii.String("knowledgeBase"),
//   	PerformFlowUiTask: jsii.String("performFlowUiTask"),
//   	PrintReports: jsii.String("printReports"),
//   	PublishWithoutApproval: jsii.String("publishWithoutApproval"),
//   	RenameSharedFolders: jsii.String("renameSharedFolders"),
//   	Research: jsii.String("research"),
//   	ShareAnalyses: jsii.String("shareAnalyses"),
//   	ShareDashboards: jsii.String("shareDashboards"),
//   	ShareDatasets: jsii.String("shareDatasets"),
//   	ShareDataSources: jsii.String("shareDataSources"),
//   	Space: jsii.String("space"),
//   	SubscribeDashboardEmailReports: jsii.String("subscribeDashboardEmailReports"),
//   	UseAgentWebSearch: jsii.String("useAgentWebSearch"),
//   	UseBedrockModels: jsii.String("useBedrockModels"),
//   	ViewAccountSpiceCapacity: jsii.String("viewAccountSpiceCapacity"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html
//
type CfnCustomPermissions_CapabilitiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The ability to add or run anomaly detection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-addorrunanomalydetectionforanalyses
	//
	AddOrRunAnomalyDetectionForAnalyses *string `field:"optional" json:"addOrRunAnomalyDetectionForAnalyses" yaml:"addOrRunAnomalyDetectionForAnalyses"`
	// The ability to perform analysis-related actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-analysis
	//
	Analysis *string `field:"optional" json:"analysis" yaml:"analysis"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-automate
	//
	Automate *string `field:"optional" json:"automate" yaml:"automate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-chatagent
	//
	ChatAgent *string `field:"optional" json:"chatAgent" yaml:"chatAgent"`
	// The ability to create and update email reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatedashboardemailreports
	//
	CreateAndUpdateDashboardEmailReports *string `field:"optional" json:"createAndUpdateDashboardEmailReports" yaml:"createAndUpdateDashboardEmailReports"`
	// The ability to create and update datasets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatedatasets
	//
	CreateAndUpdateDatasets *string `field:"optional" json:"createAndUpdateDatasets" yaml:"createAndUpdateDatasets"`
	// The ability to create and update data sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatedatasources
	//
	CreateAndUpdateDataSources *string `field:"optional" json:"createAndUpdateDataSources" yaml:"createAndUpdateDataSources"`
	// The ability to export to Create and Update themes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatethemes
	//
	CreateAndUpdateThemes *string `field:"optional" json:"createAndUpdateThemes" yaml:"createAndUpdateThemes"`
	// The ability to create and update threshold alerts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatethresholdalerts
	//
	CreateAndUpdateThresholdAlerts *string `field:"optional" json:"createAndUpdateThresholdAlerts" yaml:"createAndUpdateThresholdAlerts"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createchatagents
	//
	CreateChatAgents *string `field:"optional" json:"createChatAgents" yaml:"createChatAgents"`
	// The ability to create shared folders.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createsharedfolders
	//
	CreateSharedFolders *string `field:"optional" json:"createSharedFolders" yaml:"createSharedFolders"`
	// The ability to create a SPICE dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createspicedataset
	//
	CreateSpiceDataset *string `field:"optional" json:"createSpiceDataset" yaml:"createSpiceDataset"`
	// The ability to perform dashboard-related actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-dashboard
	//
	Dashboard *string `field:"optional" json:"dashboard" yaml:"dashboard"`
	// The ability to export to CSV files from the UI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-exporttocsv
	//
	ExportToCsv *string `field:"optional" json:"exportToCsv" yaml:"exportToCsv"`
	// The ability to export to CSV files in scheduled email reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-exporttocsvinscheduledreports
	//
	ExportToCsvInScheduledReports *string `field:"optional" json:"exportToCsvInScheduledReports" yaml:"exportToCsvInScheduledReports"`
	// The ability to export to Excel files from the UI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-exporttoexcel
	//
	ExportToExcel *string `field:"optional" json:"exportToExcel" yaml:"exportToExcel"`
	// The ability to export to Excel files in scheduled email reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-exporttoexcelinscheduledreports
	//
	ExportToExcelInScheduledReports *string `field:"optional" json:"exportToExcelInScheduledReports" yaml:"exportToExcelInScheduledReports"`
	// The ability to export to PDF files from the UI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-exporttopdf
	//
	ExportToPdf *string `field:"optional" json:"exportToPdf" yaml:"exportToPdf"`
	// The ability to export to PDF files in scheduled email reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-exporttopdfinscheduledreports
	//
	ExportToPdfInScheduledReports *string `field:"optional" json:"exportToPdfInScheduledReports" yaml:"exportToPdfInScheduledReports"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-flow
	//
	Flow *string `field:"optional" json:"flow" yaml:"flow"`
	// The ability to include content in scheduled email reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-includecontentinscheduledreportsemail
	//
	IncludeContentInScheduledReportsEmail *string `field:"optional" json:"includeContentInScheduledReportsEmail" yaml:"includeContentInScheduledReportsEmail"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-knowledgebase
	//
	KnowledgeBase *string `field:"optional" json:"knowledgeBase" yaml:"knowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-performflowuitask
	//
	PerformFlowUiTask *string `field:"optional" json:"performFlowUiTask" yaml:"performFlowUiTask"`
	// The ability to print reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-printreports
	//
	PrintReports *string `field:"optional" json:"printReports" yaml:"printReports"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-publishwithoutapproval
	//
	PublishWithoutApproval *string `field:"optional" json:"publishWithoutApproval" yaml:"publishWithoutApproval"`
	// The ability to rename shared folders.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-renamesharedfolders
	//
	RenameSharedFolders *string `field:"optional" json:"renameSharedFolders" yaml:"renameSharedFolders"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-research
	//
	Research *string `field:"optional" json:"research" yaml:"research"`
	// The ability to share analyses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareanalyses
	//
	ShareAnalyses *string `field:"optional" json:"shareAnalyses" yaml:"shareAnalyses"`
	// The ability to share dashboards.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharedashboards
	//
	ShareDashboards *string `field:"optional" json:"shareDashboards" yaml:"shareDashboards"`
	// The ability to share datasets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharedatasets
	//
	ShareDatasets *string `field:"optional" json:"shareDatasets" yaml:"shareDatasets"`
	// The ability to share data sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharedatasources
	//
	ShareDataSources *string `field:"optional" json:"shareDataSources" yaml:"shareDataSources"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-space
	//
	Space *string `field:"optional" json:"space" yaml:"space"`
	// The ability to subscribe to email reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-subscribedashboardemailreports
	//
	SubscribeDashboardEmailReports *string `field:"optional" json:"subscribeDashboardEmailReports" yaml:"subscribeDashboardEmailReports"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useagentwebsearch
	//
	UseAgentWebSearch *string `field:"optional" json:"useAgentWebSearch" yaml:"useAgentWebSearch"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usebedrockmodels
	//
	UseBedrockModels *string `field:"optional" json:"useBedrockModels" yaml:"useBedrockModels"`
	// The ability to view account SPICE capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-viewaccountspicecapacity
	//
	ViewAccountSpiceCapacity *string `field:"optional" json:"viewAccountSpiceCapacity" yaml:"viewAccountSpiceCapacity"`
}

