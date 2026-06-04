package awsquicksight


// A set of actions that correspond to Amazon Quick Sight permissions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capabilitiesProperty := &CapabilitiesProperty{
//   	AccessAppsNativeDataStore: jsii.String("accessAppsNativeDataStore"),
//   	Action: jsii.String("action"),
//   	AddOrRunAnomalyDetectionForAnalyses: jsii.String("addOrRunAnomalyDetectionForAnalyses"),
//   	AmazonBedrockArsAction: jsii.String("amazonBedrockArsAction"),
//   	AmazonBedrockFsAction: jsii.String("amazonBedrockFsAction"),
//   	AmazonBedrockKrsAction: jsii.String("amazonBedrockKrsAction"),
//   	AmazonSThreeAction: jsii.String("amazonSThreeAction"),
//   	Analysis: jsii.String("analysis"),
//   	ApproveFlowShareRequests: jsii.String("approveFlowShareRequests"),
//   	Apps: jsii.String("apps"),
//   	AsanaAction: jsii.String("asanaAction"),
//   	Automate: jsii.String("automate"),
//   	BambooHrAction: jsii.String("bambooHrAction"),
//   	BoxAgentAction: jsii.String("boxAgentAction"),
//   	BuildCalculatedFieldWithQ: jsii.String("buildCalculatedFieldWithQ"),
//   	CanvaAgentAction: jsii.String("canvaAgentAction"),
//   	ChatAgent: jsii.String("chatAgent"),
//   	ComprehendAction: jsii.String("comprehendAction"),
//   	ComprehendMedicalAction: jsii.String("comprehendMedicalAction"),
//   	ConfluenceAction: jsii.String("confluenceAction"),
//   	CreateAndUpdateAmazonBedrockArsAction: jsii.String("createAndUpdateAmazonBedrockArsAction"),
//   	CreateAndUpdateAmazonBedrockFsAction: jsii.String("createAndUpdateAmazonBedrockFsAction"),
//   	CreateAndUpdateAmazonBedrockKrsAction: jsii.String("createAndUpdateAmazonBedrockKrsAction"),
//   	CreateAndUpdateAmazonSThreeAction: jsii.String("createAndUpdateAmazonSThreeAction"),
//   	CreateAndUpdateApps: jsii.String("createAndUpdateApps"),
//   	CreateAndUpdateAsanaAction: jsii.String("createAndUpdateAsanaAction"),
//   	CreateAndUpdateBambooHrAction: jsii.String("createAndUpdateBambooHrAction"),
//   	CreateAndUpdateBoxAgentAction: jsii.String("createAndUpdateBoxAgentAction"),
//   	CreateAndUpdateCanvaAgentAction: jsii.String("createAndUpdateCanvaAgentAction"),
//   	CreateAndUpdateComprehendAction: jsii.String("createAndUpdateComprehendAction"),
//   	CreateAndUpdateComprehendMedicalAction: jsii.String("createAndUpdateComprehendMedicalAction"),
//   	CreateAndUpdateConfluenceAction: jsii.String("createAndUpdateConfluenceAction"),
//   	CreateAndUpdateDashboardEmailReports: jsii.String("createAndUpdateDashboardEmailReports"),
//   	CreateAndUpdateDatasets: jsii.String("createAndUpdateDatasets"),
//   	CreateAndUpdateDataSources: jsii.String("createAndUpdateDataSources"),
//   	CreateAndUpdateFactSetAction: jsii.String("createAndUpdateFactSetAction"),
//   	CreateAndUpdateGenericHttpAction: jsii.String("createAndUpdateGenericHttpAction"),
//   	CreateAndUpdateGithubAction: jsii.String("createAndUpdateGithubAction"),
//   	CreateAndUpdateGoogleCalendarAction: jsii.String("createAndUpdateGoogleCalendarAction"),
//   	CreateAndUpdateHubspotAction: jsii.String("createAndUpdateHubspotAction"),
//   	CreateAndUpdateHuggingFaceAction: jsii.String("createAndUpdateHuggingFaceAction"),
//   	CreateAndUpdateIntercomAction: jsii.String("createAndUpdateIntercomAction"),
//   	CreateAndUpdateJiraAction: jsii.String("createAndUpdateJiraAction"),
//   	CreateAndUpdateKnowledgeBases: jsii.String("createAndUpdateKnowledgeBases"),
//   	CreateAndUpdateLinearAction: jsii.String("createAndUpdateLinearAction"),
//   	CreateAndUpdateMcpAction: jsii.String("createAndUpdateMcpAction"),
//   	CreateAndUpdateMondayAction: jsii.String("createAndUpdateMondayAction"),
//   	CreateAndUpdateMsExchangeAction: jsii.String("createAndUpdateMsExchangeAction"),
//   	CreateAndUpdateMsTeamsAction: jsii.String("createAndUpdateMsTeamsAction"),
//   	CreateAndUpdateNewRelicAction: jsii.String("createAndUpdateNewRelicAction"),
//   	CreateAndUpdateNotionAction: jsii.String("createAndUpdateNotionAction"),
//   	CreateAndUpdateOneDriveAction: jsii.String("createAndUpdateOneDriveAction"),
//   	CreateAndUpdateOpenApiAction: jsii.String("createAndUpdateOpenApiAction"),
//   	CreateAndUpdatePagerDutyAction: jsii.String("createAndUpdatePagerDutyAction"),
//   	CreateAndUpdateSalesforceAction: jsii.String("createAndUpdateSalesforceAction"),
//   	CreateAndUpdateSandPGlobalEnergyAction: jsii.String("createAndUpdateSandPGlobalEnergyAction"),
//   	CreateAndUpdateSandPgmiAction: jsii.String("createAndUpdateSandPgmiAction"),
//   	CreateAndUpdateSapBillOfMaterialAction: jsii.String("createAndUpdateSapBillOfMaterialAction"),
//   	CreateAndUpdateSapBusinessPartnerAction: jsii.String("createAndUpdateSapBusinessPartnerAction"),
//   	CreateAndUpdateSapMaterialStockAction: jsii.String("createAndUpdateSapMaterialStockAction"),
//   	CreateAndUpdateSapPhysicalInventoryAction: jsii.String("createAndUpdateSapPhysicalInventoryAction"),
//   	CreateAndUpdateSapProductMasterDataAction: jsii.String("createAndUpdateSapProductMasterDataAction"),
//   	CreateAndUpdateServiceNowAction: jsii.String("createAndUpdateServiceNowAction"),
//   	CreateAndUpdateSharePointAction: jsii.String("createAndUpdateSharePointAction"),
//   	CreateAndUpdateSlackAction: jsii.String("createAndUpdateSlackAction"),
//   	CreateAndUpdateSmartsheetAction: jsii.String("createAndUpdateSmartsheetAction"),
//   	CreateAndUpdateTextractAction: jsii.String("createAndUpdateTextractAction"),
//   	CreateAndUpdateThemes: jsii.String("createAndUpdateThemes"),
//   	CreateAndUpdateThresholdAlerts: jsii.String("createAndUpdateThresholdAlerts"),
//   	CreateAndUpdateZendeskAction: jsii.String("createAndUpdateZendeskAction"),
//   	CreateChatAgents: jsii.String("createChatAgents"),
//   	CreateDashboardExecutiveSummaryWithQ: jsii.String("createDashboardExecutiveSummaryWithQ"),
//   	CreateSharedFolders: jsii.String("createSharedFolders"),
//   	CreateSpaces: jsii.String("createSpaces"),
//   	CreateSpiceDataset: jsii.String("createSpiceDataset"),
//   	Dashboard: jsii.String("dashboard"),
//   	EditVisualWithQ: jsii.String("editVisualWithQ"),
//   	ExportToCsv: jsii.String("exportToCsv"),
//   	ExportToCsvInScheduledReports: jsii.String("exportToCsvInScheduledReports"),
//   	ExportToExcel: jsii.String("exportToExcel"),
//   	ExportToExcelInScheduledReports: jsii.String("exportToExcelInScheduledReports"),
//   	ExportToPdf: jsii.String("exportToPdf"),
//   	ExportToPdfInScheduledReports: jsii.String("exportToPdfInScheduledReports"),
//   	Extension: jsii.String("extension"),
//   	FactSetAction: jsii.String("factSetAction"),
//   	Flow: jsii.String("flow"),
//   	GenericHttpAction: jsii.String("genericHttpAction"),
//   	GithubAction: jsii.String("githubAction"),
//   	GoogleCalendarAction: jsii.String("googleCalendarAction"),
//   	HubspotAction: jsii.String("hubspotAction"),
//   	HuggingFaceAction: jsii.String("huggingFaceAction"),
//   	IncludeContentInScheduledReportsEmail: jsii.String("includeContentInScheduledReportsEmail"),
//   	IntercomAction: jsii.String("intercomAction"),
//   	InvokeAppsAiInference: jsii.String("invokeAppsAiInference"),
//   	JiraAction: jsii.String("jiraAction"),
//   	KnowledgeBase: jsii.String("knowledgeBase"),
//   	LinearAction: jsii.String("linearAction"),
//   	ManageSharedFolders: jsii.String("manageSharedFolders"),
//   	McpAction: jsii.String("mcpAction"),
//   	MondayAction: jsii.String("mondayAction"),
//   	MsExchangeAction: jsii.String("msExchangeAction"),
//   	MsTeamsAction: jsii.String("msTeamsAction"),
//   	NewRelicAction: jsii.String("newRelicAction"),
//   	NotionAction: jsii.String("notionAction"),
//   	OneDriveAction: jsii.String("oneDriveAction"),
//   	OpenApiAction: jsii.String("openApiAction"),
//   	PagerDutyAction: jsii.String("pagerDutyAction"),
//   	PerformFlowUiTask: jsii.String("performFlowUiTask"),
//   	PrintReports: jsii.String("printReports"),
//   	PublishWithoutApproval: jsii.String("publishWithoutApproval"),
//   	RenameSharedFolders: jsii.String("renameSharedFolders"),
//   	Research: jsii.String("research"),
//   	SalesforceAction: jsii.String("salesforceAction"),
//   	SandPGlobalEnergyAction: jsii.String("sandPGlobalEnergyAction"),
//   	SandPgmiAction: jsii.String("sandPgmiAction"),
//   	SapBillOfMaterialAction: jsii.String("sapBillOfMaterialAction"),
//   	SapBusinessPartnerAction: jsii.String("sapBusinessPartnerAction"),
//   	SapMaterialStockAction: jsii.String("sapMaterialStockAction"),
//   	SapPhysicalInventoryAction: jsii.String("sapPhysicalInventoryAction"),
//   	SapProductMasterDataAction: jsii.String("sapProductMasterDataAction"),
//   	ServiceNowAction: jsii.String("serviceNowAction"),
//   	ShareAmazonBedrockArsAction: jsii.String("shareAmazonBedrockArsAction"),
//   	ShareAmazonBedrockFsAction: jsii.String("shareAmazonBedrockFsAction"),
//   	ShareAmazonBedrockKrsAction: jsii.String("shareAmazonBedrockKrsAction"),
//   	ShareAmazonSThreeAction: jsii.String("shareAmazonSThreeAction"),
//   	ShareAnalyses: jsii.String("shareAnalyses"),
//   	ShareApps: jsii.String("shareApps"),
//   	ShareAsanaAction: jsii.String("shareAsanaAction"),
//   	ShareBambooHrAction: jsii.String("shareBambooHrAction"),
//   	ShareBoxAgentAction: jsii.String("shareBoxAgentAction"),
//   	ShareCanvaAgentAction: jsii.String("shareCanvaAgentAction"),
//   	ShareChatAgents: jsii.String("shareChatAgents"),
//   	ShareComprehendAction: jsii.String("shareComprehendAction"),
//   	ShareComprehendMedicalAction: jsii.String("shareComprehendMedicalAction"),
//   	ShareConfluenceAction: jsii.String("shareConfluenceAction"),
//   	ShareDashboards: jsii.String("shareDashboards"),
//   	ShareDatasets: jsii.String("shareDatasets"),
//   	ShareDataSources: jsii.String("shareDataSources"),
//   	ShareFactSetAction: jsii.String("shareFactSetAction"),
//   	ShareGenericHttpAction: jsii.String("shareGenericHttpAction"),
//   	ShareGithubAction: jsii.String("shareGithubAction"),
//   	ShareGoogleCalendarAction: jsii.String("shareGoogleCalendarAction"),
//   	ShareHubspotAction: jsii.String("shareHubspotAction"),
//   	ShareHuggingFaceAction: jsii.String("shareHuggingFaceAction"),
//   	ShareIntercomAction: jsii.String("shareIntercomAction"),
//   	ShareJiraAction: jsii.String("shareJiraAction"),
//   	ShareKnowledgeBases: jsii.String("shareKnowledgeBases"),
//   	ShareLinearAction: jsii.String("shareLinearAction"),
//   	ShareMcpAction: jsii.String("shareMcpAction"),
//   	ShareMondayAction: jsii.String("shareMondayAction"),
//   	ShareMsExchangeAction: jsii.String("shareMsExchangeAction"),
//   	ShareMsTeamsAction: jsii.String("shareMsTeamsAction"),
//   	ShareNewRelicAction: jsii.String("shareNewRelicAction"),
//   	ShareNotionAction: jsii.String("shareNotionAction"),
//   	ShareOneDriveAction: jsii.String("shareOneDriveAction"),
//   	ShareOpenApiAction: jsii.String("shareOpenApiAction"),
//   	SharePagerDutyAction: jsii.String("sharePagerDutyAction"),
//   	SharePointAction: jsii.String("sharePointAction"),
//   	ShareSalesforceAction: jsii.String("shareSalesforceAction"),
//   	ShareSandPGlobalEnergyAction: jsii.String("shareSandPGlobalEnergyAction"),
//   	ShareSandPgmiAction: jsii.String("shareSandPgmiAction"),
//   	ShareSapBillOfMaterialAction: jsii.String("shareSapBillOfMaterialAction"),
//   	ShareSapBusinessPartnerAction: jsii.String("shareSapBusinessPartnerAction"),
//   	ShareSapMaterialStockAction: jsii.String("shareSapMaterialStockAction"),
//   	ShareSapPhysicalInventoryAction: jsii.String("shareSapPhysicalInventoryAction"),
//   	ShareSapProductMasterDataAction: jsii.String("shareSapProductMasterDataAction"),
//   	ShareServiceNowAction: jsii.String("shareServiceNowAction"),
//   	ShareSharePointAction: jsii.String("shareSharePointAction"),
//   	ShareSlackAction: jsii.String("shareSlackAction"),
//   	ShareSmartsheetAction: jsii.String("shareSmartsheetAction"),
//   	ShareSpaces: jsii.String("shareSpaces"),
//   	ShareTextractAction: jsii.String("shareTextractAction"),
//   	ShareZendeskAction: jsii.String("shareZendeskAction"),
//   	SlackAction: jsii.String("slackAction"),
//   	SmartsheetAction: jsii.String("smartsheetAction"),
//   	Space: jsii.String("space"),
//   	SubscribeDashboardEmailReports: jsii.String("subscribeDashboardEmailReports"),
//   	TextractAction: jsii.String("textractAction"),
//   	Topic: jsii.String("topic"),
//   	UseAgentWebSearch: jsii.String("useAgentWebSearch"),
//   	UseAmazonBedrockArsAction: jsii.String("useAmazonBedrockArsAction"),
//   	UseAmazonBedrockFsAction: jsii.String("useAmazonBedrockFsAction"),
//   	UseAmazonBedrockKrsAction: jsii.String("useAmazonBedrockKrsAction"),
//   	UseAmazonSThreeAction: jsii.String("useAmazonSThreeAction"),
//   	UseAsanaAction: jsii.String("useAsanaAction"),
//   	UseBambooHrAction: jsii.String("useBambooHrAction"),
//   	UseBedrockModels: jsii.String("useBedrockModels"),
//   	UseBoxAgentAction: jsii.String("useBoxAgentAction"),
//   	UseCanvaAgentAction: jsii.String("useCanvaAgentAction"),
//   	UseComprehendAction: jsii.String("useComprehendAction"),
//   	UseComprehendMedicalAction: jsii.String("useComprehendMedicalAction"),
//   	UseConfluenceAction: jsii.String("useConfluenceAction"),
//   	UseFactSetAction: jsii.String("useFactSetAction"),
//   	UseGenericHttpAction: jsii.String("useGenericHttpAction"),
//   	UseGithubAction: jsii.String("useGithubAction"),
//   	UseGoogleCalendarAction: jsii.String("useGoogleCalendarAction"),
//   	UseHubspotAction: jsii.String("useHubspotAction"),
//   	UseHuggingFaceAction: jsii.String("useHuggingFaceAction"),
//   	UseIntercomAction: jsii.String("useIntercomAction"),
//   	UseJiraAction: jsii.String("useJiraAction"),
//   	UseLinearAction: jsii.String("useLinearAction"),
//   	UseMcpAction: jsii.String("useMcpAction"),
//   	UseMondayAction: jsii.String("useMondayAction"),
//   	UseMsExchangeAction: jsii.String("useMsExchangeAction"),
//   	UseMsTeamsAction: jsii.String("useMsTeamsAction"),
//   	UseNewRelicAction: jsii.String("useNewRelicAction"),
//   	UseNotionAction: jsii.String("useNotionAction"),
//   	UseOneDriveAction: jsii.String("useOneDriveAction"),
//   	UseOpenApiAction: jsii.String("useOpenApiAction"),
//   	UsePagerDutyAction: jsii.String("usePagerDutyAction"),
//   	UseSalesforceAction: jsii.String("useSalesforceAction"),
//   	UseSandPGlobalEnergyAction: jsii.String("useSandPGlobalEnergyAction"),
//   	UseSandPgmiAction: jsii.String("useSandPgmiAction"),
//   	UseSapBillOfMaterialAction: jsii.String("useSapBillOfMaterialAction"),
//   	UseSapBusinessPartnerAction: jsii.String("useSapBusinessPartnerAction"),
//   	UseSapMaterialStockAction: jsii.String("useSapMaterialStockAction"),
//   	UseSapPhysicalInventoryAction: jsii.String("useSapPhysicalInventoryAction"),
//   	UseSapProductMasterDataAction: jsii.String("useSapProductMasterDataAction"),
//   	UseServiceNowAction: jsii.String("useServiceNowAction"),
//   	UseSharePointAction: jsii.String("useSharePointAction"),
//   	UseSlackAction: jsii.String("useSlackAction"),
//   	UseSmartsheetAction: jsii.String("useSmartsheetAction"),
//   	UseTextractAction: jsii.String("useTextractAction"),
//   	UseZendeskAction: jsii.String("useZendeskAction"),
//   	ViewAccountSpiceCapacity: jsii.String("viewAccountSpiceCapacity"),
//   	ZendeskAction: jsii.String("zendeskAction"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html
//
type CfnCustomPermissions_CapabilitiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-accessappsnativedatastore
	//
	AccessAppsNativeDataStore *string `field:"optional" json:"accessAppsNativeDataStore" yaml:"accessAppsNativeDataStore"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The ability to add or run anomaly detection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-addorrunanomalydetectionforanalyses
	//
	AddOrRunAnomalyDetectionForAnalyses *string `field:"optional" json:"addOrRunAnomalyDetectionForAnalyses" yaml:"addOrRunAnomalyDetectionForAnalyses"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-amazonbedrockarsaction
	//
	AmazonBedrockArsAction *string `field:"optional" json:"amazonBedrockArsAction" yaml:"amazonBedrockArsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-amazonbedrockfsaction
	//
	AmazonBedrockFsAction *string `field:"optional" json:"amazonBedrockFsAction" yaml:"amazonBedrockFsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-amazonbedrockkrsaction
	//
	AmazonBedrockKrsAction *string `field:"optional" json:"amazonBedrockKrsAction" yaml:"amazonBedrockKrsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-amazonsthreeaction
	//
	AmazonSThreeAction *string `field:"optional" json:"amazonSThreeAction" yaml:"amazonSThreeAction"`
	// The ability to perform analysis-related actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-analysis
	//
	Analysis *string `field:"optional" json:"analysis" yaml:"analysis"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-approveflowsharerequests
	//
	ApproveFlowShareRequests *string `field:"optional" json:"approveFlowShareRequests" yaml:"approveFlowShareRequests"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-apps
	//
	Apps *string `field:"optional" json:"apps" yaml:"apps"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-asanaaction
	//
	AsanaAction *string `field:"optional" json:"asanaAction" yaml:"asanaAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-automate
	//
	Automate *string `field:"optional" json:"automate" yaml:"automate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-bamboohraction
	//
	BambooHrAction *string `field:"optional" json:"bambooHrAction" yaml:"bambooHrAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-boxagentaction
	//
	BoxAgentAction *string `field:"optional" json:"boxAgentAction" yaml:"boxAgentAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-buildcalculatedfieldwithq
	//
	BuildCalculatedFieldWithQ *string `field:"optional" json:"buildCalculatedFieldWithQ" yaml:"buildCalculatedFieldWithQ"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-canvaagentaction
	//
	CanvaAgentAction *string `field:"optional" json:"canvaAgentAction" yaml:"canvaAgentAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-chatagent
	//
	ChatAgent *string `field:"optional" json:"chatAgent" yaml:"chatAgent"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-comprehendaction
	//
	ComprehendAction *string `field:"optional" json:"comprehendAction" yaml:"comprehendAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-comprehendmedicalaction
	//
	ComprehendMedicalAction *string `field:"optional" json:"comprehendMedicalAction" yaml:"comprehendMedicalAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-confluenceaction
	//
	ConfluenceAction *string `field:"optional" json:"confluenceAction" yaml:"confluenceAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateamazonbedrockarsaction
	//
	CreateAndUpdateAmazonBedrockArsAction *string `field:"optional" json:"createAndUpdateAmazonBedrockArsAction" yaml:"createAndUpdateAmazonBedrockArsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateamazonbedrockfsaction
	//
	CreateAndUpdateAmazonBedrockFsAction *string `field:"optional" json:"createAndUpdateAmazonBedrockFsAction" yaml:"createAndUpdateAmazonBedrockFsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateamazonbedrockkrsaction
	//
	CreateAndUpdateAmazonBedrockKrsAction *string `field:"optional" json:"createAndUpdateAmazonBedrockKrsAction" yaml:"createAndUpdateAmazonBedrockKrsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateamazonsthreeaction
	//
	CreateAndUpdateAmazonSThreeAction *string `field:"optional" json:"createAndUpdateAmazonSThreeAction" yaml:"createAndUpdateAmazonSThreeAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateapps
	//
	CreateAndUpdateApps *string `field:"optional" json:"createAndUpdateApps" yaml:"createAndUpdateApps"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateasanaaction
	//
	CreateAndUpdateAsanaAction *string `field:"optional" json:"createAndUpdateAsanaAction" yaml:"createAndUpdateAsanaAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatebamboohraction
	//
	CreateAndUpdateBambooHrAction *string `field:"optional" json:"createAndUpdateBambooHrAction" yaml:"createAndUpdateBambooHrAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateboxagentaction
	//
	CreateAndUpdateBoxAgentAction *string `field:"optional" json:"createAndUpdateBoxAgentAction" yaml:"createAndUpdateBoxAgentAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatecanvaagentaction
	//
	CreateAndUpdateCanvaAgentAction *string `field:"optional" json:"createAndUpdateCanvaAgentAction" yaml:"createAndUpdateCanvaAgentAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatecomprehendaction
	//
	CreateAndUpdateComprehendAction *string `field:"optional" json:"createAndUpdateComprehendAction" yaml:"createAndUpdateComprehendAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatecomprehendmedicalaction
	//
	CreateAndUpdateComprehendMedicalAction *string `field:"optional" json:"createAndUpdateComprehendMedicalAction" yaml:"createAndUpdateComprehendMedicalAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateconfluenceaction
	//
	CreateAndUpdateConfluenceAction *string `field:"optional" json:"createAndUpdateConfluenceAction" yaml:"createAndUpdateConfluenceAction"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatefactsetaction
	//
	CreateAndUpdateFactSetAction *string `field:"optional" json:"createAndUpdateFactSetAction" yaml:"createAndUpdateFactSetAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdategenerichttpaction
	//
	CreateAndUpdateGenericHttpAction *string `field:"optional" json:"createAndUpdateGenericHttpAction" yaml:"createAndUpdateGenericHttpAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdategithubaction
	//
	CreateAndUpdateGithubAction *string `field:"optional" json:"createAndUpdateGithubAction" yaml:"createAndUpdateGithubAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdategooglecalendaraction
	//
	CreateAndUpdateGoogleCalendarAction *string `field:"optional" json:"createAndUpdateGoogleCalendarAction" yaml:"createAndUpdateGoogleCalendarAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatehubspotaction
	//
	CreateAndUpdateHubspotAction *string `field:"optional" json:"createAndUpdateHubspotAction" yaml:"createAndUpdateHubspotAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatehuggingfaceaction
	//
	CreateAndUpdateHuggingFaceAction *string `field:"optional" json:"createAndUpdateHuggingFaceAction" yaml:"createAndUpdateHuggingFaceAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateintercomaction
	//
	CreateAndUpdateIntercomAction *string `field:"optional" json:"createAndUpdateIntercomAction" yaml:"createAndUpdateIntercomAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatejiraaction
	//
	CreateAndUpdateJiraAction *string `field:"optional" json:"createAndUpdateJiraAction" yaml:"createAndUpdateJiraAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateknowledgebases
	//
	CreateAndUpdateKnowledgeBases *string `field:"optional" json:"createAndUpdateKnowledgeBases" yaml:"createAndUpdateKnowledgeBases"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatelinearaction
	//
	CreateAndUpdateLinearAction *string `field:"optional" json:"createAndUpdateLinearAction" yaml:"createAndUpdateLinearAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatemcpaction
	//
	CreateAndUpdateMcpAction *string `field:"optional" json:"createAndUpdateMcpAction" yaml:"createAndUpdateMcpAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatemondayaction
	//
	CreateAndUpdateMondayAction *string `field:"optional" json:"createAndUpdateMondayAction" yaml:"createAndUpdateMondayAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatemsexchangeaction
	//
	CreateAndUpdateMsExchangeAction *string `field:"optional" json:"createAndUpdateMsExchangeAction" yaml:"createAndUpdateMsExchangeAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatemsteamsaction
	//
	CreateAndUpdateMsTeamsAction *string `field:"optional" json:"createAndUpdateMsTeamsAction" yaml:"createAndUpdateMsTeamsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatenewrelicaction
	//
	CreateAndUpdateNewRelicAction *string `field:"optional" json:"createAndUpdateNewRelicAction" yaml:"createAndUpdateNewRelicAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatenotionaction
	//
	CreateAndUpdateNotionAction *string `field:"optional" json:"createAndUpdateNotionAction" yaml:"createAndUpdateNotionAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateonedriveaction
	//
	CreateAndUpdateOneDriveAction *string `field:"optional" json:"createAndUpdateOneDriveAction" yaml:"createAndUpdateOneDriveAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateopenapiaction
	//
	CreateAndUpdateOpenApiAction *string `field:"optional" json:"createAndUpdateOpenApiAction" yaml:"createAndUpdateOpenApiAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatepagerdutyaction
	//
	CreateAndUpdatePagerDutyAction *string `field:"optional" json:"createAndUpdatePagerDutyAction" yaml:"createAndUpdatePagerDutyAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatesalesforceaction
	//
	CreateAndUpdateSalesforceAction *string `field:"optional" json:"createAndUpdateSalesforceAction" yaml:"createAndUpdateSalesforceAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatesandpglobalenergyaction
	//
	CreateAndUpdateSandPGlobalEnergyAction *string `field:"optional" json:"createAndUpdateSandPGlobalEnergyAction" yaml:"createAndUpdateSandPGlobalEnergyAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatesandpgmiaction
	//
	CreateAndUpdateSandPgmiAction *string `field:"optional" json:"createAndUpdateSandPgmiAction" yaml:"createAndUpdateSandPgmiAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatesapbillofmaterialaction
	//
	CreateAndUpdateSapBillOfMaterialAction *string `field:"optional" json:"createAndUpdateSapBillOfMaterialAction" yaml:"createAndUpdateSapBillOfMaterialAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatesapbusinesspartneraction
	//
	CreateAndUpdateSapBusinessPartnerAction *string `field:"optional" json:"createAndUpdateSapBusinessPartnerAction" yaml:"createAndUpdateSapBusinessPartnerAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatesapmaterialstockaction
	//
	CreateAndUpdateSapMaterialStockAction *string `field:"optional" json:"createAndUpdateSapMaterialStockAction" yaml:"createAndUpdateSapMaterialStockAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatesapphysicalinventoryaction
	//
	CreateAndUpdateSapPhysicalInventoryAction *string `field:"optional" json:"createAndUpdateSapPhysicalInventoryAction" yaml:"createAndUpdateSapPhysicalInventoryAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatesapproductmasterdataaction
	//
	CreateAndUpdateSapProductMasterDataAction *string `field:"optional" json:"createAndUpdateSapProductMasterDataAction" yaml:"createAndUpdateSapProductMasterDataAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateservicenowaction
	//
	CreateAndUpdateServiceNowAction *string `field:"optional" json:"createAndUpdateServiceNowAction" yaml:"createAndUpdateServiceNowAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatesharepointaction
	//
	CreateAndUpdateSharePointAction *string `field:"optional" json:"createAndUpdateSharePointAction" yaml:"createAndUpdateSharePointAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateslackaction
	//
	CreateAndUpdateSlackAction *string `field:"optional" json:"createAndUpdateSlackAction" yaml:"createAndUpdateSlackAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatesmartsheetaction
	//
	CreateAndUpdateSmartsheetAction *string `field:"optional" json:"createAndUpdateSmartsheetAction" yaml:"createAndUpdateSmartsheetAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatetextractaction
	//
	CreateAndUpdateTextractAction *string `field:"optional" json:"createAndUpdateTextractAction" yaml:"createAndUpdateTextractAction"`
	// The ability to export to Create and Update themes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatethemes
	//
	CreateAndUpdateThemes *string `field:"optional" json:"createAndUpdateThemes" yaml:"createAndUpdateThemes"`
	// The ability to create and update threshold alerts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatethresholdalerts
	//
	CreateAndUpdateThresholdAlerts *string `field:"optional" json:"createAndUpdateThresholdAlerts" yaml:"createAndUpdateThresholdAlerts"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatezendeskaction
	//
	CreateAndUpdateZendeskAction *string `field:"optional" json:"createAndUpdateZendeskAction" yaml:"createAndUpdateZendeskAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createchatagents
	//
	CreateChatAgents *string `field:"optional" json:"createChatAgents" yaml:"createChatAgents"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createdashboardexecutivesummarywithq
	//
	CreateDashboardExecutiveSummaryWithQ *string `field:"optional" json:"createDashboardExecutiveSummaryWithQ" yaml:"createDashboardExecutiveSummaryWithQ"`
	// The ability to create shared folders.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createsharedfolders
	//
	CreateSharedFolders *string `field:"optional" json:"createSharedFolders" yaml:"createSharedFolders"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createspaces
	//
	CreateSpaces *string `field:"optional" json:"createSpaces" yaml:"createSpaces"`
	// The ability to create a SPICE dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createspicedataset
	//
	CreateSpiceDataset *string `field:"optional" json:"createSpiceDataset" yaml:"createSpiceDataset"`
	// The ability to perform dashboard-related actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-dashboard
	//
	Dashboard *string `field:"optional" json:"dashboard" yaml:"dashboard"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-editvisualwithq
	//
	EditVisualWithQ *string `field:"optional" json:"editVisualWithQ" yaml:"editVisualWithQ"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-extension
	//
	Extension *string `field:"optional" json:"extension" yaml:"extension"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-factsetaction
	//
	FactSetAction *string `field:"optional" json:"factSetAction" yaml:"factSetAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-flow
	//
	Flow *string `field:"optional" json:"flow" yaml:"flow"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-generichttpaction
	//
	GenericHttpAction *string `field:"optional" json:"genericHttpAction" yaml:"genericHttpAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-githubaction
	//
	GithubAction *string `field:"optional" json:"githubAction" yaml:"githubAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-googlecalendaraction
	//
	GoogleCalendarAction *string `field:"optional" json:"googleCalendarAction" yaml:"googleCalendarAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-hubspotaction
	//
	HubspotAction *string `field:"optional" json:"hubspotAction" yaml:"hubspotAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-huggingfaceaction
	//
	HuggingFaceAction *string `field:"optional" json:"huggingFaceAction" yaml:"huggingFaceAction"`
	// The ability to include content in scheduled email reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-includecontentinscheduledreportsemail
	//
	IncludeContentInScheduledReportsEmail *string `field:"optional" json:"includeContentInScheduledReportsEmail" yaml:"includeContentInScheduledReportsEmail"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-intercomaction
	//
	IntercomAction *string `field:"optional" json:"intercomAction" yaml:"intercomAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-invokeappsaiinference
	//
	InvokeAppsAiInference *string `field:"optional" json:"invokeAppsAiInference" yaml:"invokeAppsAiInference"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-jiraaction
	//
	JiraAction *string `field:"optional" json:"jiraAction" yaml:"jiraAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-knowledgebase
	//
	KnowledgeBase *string `field:"optional" json:"knowledgeBase" yaml:"knowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-linearaction
	//
	LinearAction *string `field:"optional" json:"linearAction" yaml:"linearAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-managesharedfolders
	//
	ManageSharedFolders *string `field:"optional" json:"manageSharedFolders" yaml:"manageSharedFolders"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-mcpaction
	//
	McpAction *string `field:"optional" json:"mcpAction" yaml:"mcpAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-mondayaction
	//
	MondayAction *string `field:"optional" json:"mondayAction" yaml:"mondayAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-msexchangeaction
	//
	MsExchangeAction *string `field:"optional" json:"msExchangeAction" yaml:"msExchangeAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-msteamsaction
	//
	MsTeamsAction *string `field:"optional" json:"msTeamsAction" yaml:"msTeamsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-newrelicaction
	//
	NewRelicAction *string `field:"optional" json:"newRelicAction" yaml:"newRelicAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-notionaction
	//
	NotionAction *string `field:"optional" json:"notionAction" yaml:"notionAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-onedriveaction
	//
	OneDriveAction *string `field:"optional" json:"oneDriveAction" yaml:"oneDriveAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-openapiaction
	//
	OpenApiAction *string `field:"optional" json:"openApiAction" yaml:"openApiAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-pagerdutyaction
	//
	PagerDutyAction *string `field:"optional" json:"pagerDutyAction" yaml:"pagerDutyAction"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-salesforceaction
	//
	SalesforceAction *string `field:"optional" json:"salesforceAction" yaml:"salesforceAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sandpglobalenergyaction
	//
	SandPGlobalEnergyAction *string `field:"optional" json:"sandPGlobalEnergyAction" yaml:"sandPGlobalEnergyAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sandpgmiaction
	//
	SandPgmiAction *string `field:"optional" json:"sandPgmiAction" yaml:"sandPgmiAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sapbillofmaterialaction
	//
	SapBillOfMaterialAction *string `field:"optional" json:"sapBillOfMaterialAction" yaml:"sapBillOfMaterialAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sapbusinesspartneraction
	//
	SapBusinessPartnerAction *string `field:"optional" json:"sapBusinessPartnerAction" yaml:"sapBusinessPartnerAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sapmaterialstockaction
	//
	SapMaterialStockAction *string `field:"optional" json:"sapMaterialStockAction" yaml:"sapMaterialStockAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sapphysicalinventoryaction
	//
	SapPhysicalInventoryAction *string `field:"optional" json:"sapPhysicalInventoryAction" yaml:"sapPhysicalInventoryAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sapproductmasterdataaction
	//
	SapProductMasterDataAction *string `field:"optional" json:"sapProductMasterDataAction" yaml:"sapProductMasterDataAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-servicenowaction
	//
	ServiceNowAction *string `field:"optional" json:"serviceNowAction" yaml:"serviceNowAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareamazonbedrockarsaction
	//
	ShareAmazonBedrockArsAction *string `field:"optional" json:"shareAmazonBedrockArsAction" yaml:"shareAmazonBedrockArsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareamazonbedrockfsaction
	//
	ShareAmazonBedrockFsAction *string `field:"optional" json:"shareAmazonBedrockFsAction" yaml:"shareAmazonBedrockFsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareamazonbedrockkrsaction
	//
	ShareAmazonBedrockKrsAction *string `field:"optional" json:"shareAmazonBedrockKrsAction" yaml:"shareAmazonBedrockKrsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareamazonsthreeaction
	//
	ShareAmazonSThreeAction *string `field:"optional" json:"shareAmazonSThreeAction" yaml:"shareAmazonSThreeAction"`
	// The ability to share analyses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareanalyses
	//
	ShareAnalyses *string `field:"optional" json:"shareAnalyses" yaml:"shareAnalyses"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareapps
	//
	ShareApps *string `field:"optional" json:"shareApps" yaml:"shareApps"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareasanaaction
	//
	ShareAsanaAction *string `field:"optional" json:"shareAsanaAction" yaml:"shareAsanaAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharebamboohraction
	//
	ShareBambooHrAction *string `field:"optional" json:"shareBambooHrAction" yaml:"shareBambooHrAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareboxagentaction
	//
	ShareBoxAgentAction *string `field:"optional" json:"shareBoxAgentAction" yaml:"shareBoxAgentAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharecanvaagentaction
	//
	ShareCanvaAgentAction *string `field:"optional" json:"shareCanvaAgentAction" yaml:"shareCanvaAgentAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharechatagents
	//
	ShareChatAgents *string `field:"optional" json:"shareChatAgents" yaml:"shareChatAgents"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharecomprehendaction
	//
	ShareComprehendAction *string `field:"optional" json:"shareComprehendAction" yaml:"shareComprehendAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharecomprehendmedicalaction
	//
	ShareComprehendMedicalAction *string `field:"optional" json:"shareComprehendMedicalAction" yaml:"shareComprehendMedicalAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareconfluenceaction
	//
	ShareConfluenceAction *string `field:"optional" json:"shareConfluenceAction" yaml:"shareConfluenceAction"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharefactsetaction
	//
	ShareFactSetAction *string `field:"optional" json:"shareFactSetAction" yaml:"shareFactSetAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharegenerichttpaction
	//
	ShareGenericHttpAction *string `field:"optional" json:"shareGenericHttpAction" yaml:"shareGenericHttpAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharegithubaction
	//
	ShareGithubAction *string `field:"optional" json:"shareGithubAction" yaml:"shareGithubAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharegooglecalendaraction
	//
	ShareGoogleCalendarAction *string `field:"optional" json:"shareGoogleCalendarAction" yaml:"shareGoogleCalendarAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharehubspotaction
	//
	ShareHubspotAction *string `field:"optional" json:"shareHubspotAction" yaml:"shareHubspotAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharehuggingfaceaction
	//
	ShareHuggingFaceAction *string `field:"optional" json:"shareHuggingFaceAction" yaml:"shareHuggingFaceAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareintercomaction
	//
	ShareIntercomAction *string `field:"optional" json:"shareIntercomAction" yaml:"shareIntercomAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharejiraaction
	//
	ShareJiraAction *string `field:"optional" json:"shareJiraAction" yaml:"shareJiraAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareknowledgebases
	//
	ShareKnowledgeBases *string `field:"optional" json:"shareKnowledgeBases" yaml:"shareKnowledgeBases"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharelinearaction
	//
	ShareLinearAction *string `field:"optional" json:"shareLinearAction" yaml:"shareLinearAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharemcpaction
	//
	ShareMcpAction *string `field:"optional" json:"shareMcpAction" yaml:"shareMcpAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharemondayaction
	//
	ShareMondayAction *string `field:"optional" json:"shareMondayAction" yaml:"shareMondayAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharemsexchangeaction
	//
	ShareMsExchangeAction *string `field:"optional" json:"shareMsExchangeAction" yaml:"shareMsExchangeAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharemsteamsaction
	//
	ShareMsTeamsAction *string `field:"optional" json:"shareMsTeamsAction" yaml:"shareMsTeamsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharenewrelicaction
	//
	ShareNewRelicAction *string `field:"optional" json:"shareNewRelicAction" yaml:"shareNewRelicAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharenotionaction
	//
	ShareNotionAction *string `field:"optional" json:"shareNotionAction" yaml:"shareNotionAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareonedriveaction
	//
	ShareOneDriveAction *string `field:"optional" json:"shareOneDriveAction" yaml:"shareOneDriveAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareopenapiaction
	//
	ShareOpenApiAction *string `field:"optional" json:"shareOpenApiAction" yaml:"shareOpenApiAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharepagerdutyaction
	//
	SharePagerDutyAction *string `field:"optional" json:"sharePagerDutyAction" yaml:"sharePagerDutyAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharepointaction
	//
	SharePointAction *string `field:"optional" json:"sharePointAction" yaml:"sharePointAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesalesforceaction
	//
	ShareSalesforceAction *string `field:"optional" json:"shareSalesforceAction" yaml:"shareSalesforceAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesandpglobalenergyaction
	//
	ShareSandPGlobalEnergyAction *string `field:"optional" json:"shareSandPGlobalEnergyAction" yaml:"shareSandPGlobalEnergyAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesandpgmiaction
	//
	ShareSandPgmiAction *string `field:"optional" json:"shareSandPgmiAction" yaml:"shareSandPgmiAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesapbillofmaterialaction
	//
	ShareSapBillOfMaterialAction *string `field:"optional" json:"shareSapBillOfMaterialAction" yaml:"shareSapBillOfMaterialAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesapbusinesspartneraction
	//
	ShareSapBusinessPartnerAction *string `field:"optional" json:"shareSapBusinessPartnerAction" yaml:"shareSapBusinessPartnerAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesapmaterialstockaction
	//
	ShareSapMaterialStockAction *string `field:"optional" json:"shareSapMaterialStockAction" yaml:"shareSapMaterialStockAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesapphysicalinventoryaction
	//
	ShareSapPhysicalInventoryAction *string `field:"optional" json:"shareSapPhysicalInventoryAction" yaml:"shareSapPhysicalInventoryAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesapproductmasterdataaction
	//
	ShareSapProductMasterDataAction *string `field:"optional" json:"shareSapProductMasterDataAction" yaml:"shareSapProductMasterDataAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareservicenowaction
	//
	ShareServiceNowAction *string `field:"optional" json:"shareServiceNowAction" yaml:"shareServiceNowAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesharepointaction
	//
	ShareSharePointAction *string `field:"optional" json:"shareSharePointAction" yaml:"shareSharePointAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareslackaction
	//
	ShareSlackAction *string `field:"optional" json:"shareSlackAction" yaml:"shareSlackAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesmartsheetaction
	//
	ShareSmartsheetAction *string `field:"optional" json:"shareSmartsheetAction" yaml:"shareSmartsheetAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharespaces
	//
	ShareSpaces *string `field:"optional" json:"shareSpaces" yaml:"shareSpaces"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharetextractaction
	//
	ShareTextractAction *string `field:"optional" json:"shareTextractAction" yaml:"shareTextractAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharezendeskaction
	//
	ShareZendeskAction *string `field:"optional" json:"shareZendeskAction" yaml:"shareZendeskAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-slackaction
	//
	SlackAction *string `field:"optional" json:"slackAction" yaml:"slackAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-smartsheetaction
	//
	SmartsheetAction *string `field:"optional" json:"smartsheetAction" yaml:"smartsheetAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-space
	//
	Space *string `field:"optional" json:"space" yaml:"space"`
	// The ability to subscribe to email reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-subscribedashboardemailreports
	//
	SubscribeDashboardEmailReports *string `field:"optional" json:"subscribeDashboardEmailReports" yaml:"subscribeDashboardEmailReports"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-textractaction
	//
	TextractAction *string `field:"optional" json:"textractAction" yaml:"textractAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-topic
	//
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useagentwebsearch
	//
	UseAgentWebSearch *string `field:"optional" json:"useAgentWebSearch" yaml:"useAgentWebSearch"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useamazonbedrockarsaction
	//
	UseAmazonBedrockArsAction *string `field:"optional" json:"useAmazonBedrockArsAction" yaml:"useAmazonBedrockArsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useamazonbedrockfsaction
	//
	UseAmazonBedrockFsAction *string `field:"optional" json:"useAmazonBedrockFsAction" yaml:"useAmazonBedrockFsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useamazonbedrockkrsaction
	//
	UseAmazonBedrockKrsAction *string `field:"optional" json:"useAmazonBedrockKrsAction" yaml:"useAmazonBedrockKrsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useamazonsthreeaction
	//
	UseAmazonSThreeAction *string `field:"optional" json:"useAmazonSThreeAction" yaml:"useAmazonSThreeAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useasanaaction
	//
	UseAsanaAction *string `field:"optional" json:"useAsanaAction" yaml:"useAsanaAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usebamboohraction
	//
	UseBambooHrAction *string `field:"optional" json:"useBambooHrAction" yaml:"useBambooHrAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usebedrockmodels
	//
	UseBedrockModels *string `field:"optional" json:"useBedrockModels" yaml:"useBedrockModels"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useboxagentaction
	//
	UseBoxAgentAction *string `field:"optional" json:"useBoxAgentAction" yaml:"useBoxAgentAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usecanvaagentaction
	//
	UseCanvaAgentAction *string `field:"optional" json:"useCanvaAgentAction" yaml:"useCanvaAgentAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usecomprehendaction
	//
	UseComprehendAction *string `field:"optional" json:"useComprehendAction" yaml:"useComprehendAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usecomprehendmedicalaction
	//
	UseComprehendMedicalAction *string `field:"optional" json:"useComprehendMedicalAction" yaml:"useComprehendMedicalAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useconfluenceaction
	//
	UseConfluenceAction *string `field:"optional" json:"useConfluenceAction" yaml:"useConfluenceAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usefactsetaction
	//
	UseFactSetAction *string `field:"optional" json:"useFactSetAction" yaml:"useFactSetAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usegenerichttpaction
	//
	UseGenericHttpAction *string `field:"optional" json:"useGenericHttpAction" yaml:"useGenericHttpAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usegithubaction
	//
	UseGithubAction *string `field:"optional" json:"useGithubAction" yaml:"useGithubAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usegooglecalendaraction
	//
	UseGoogleCalendarAction *string `field:"optional" json:"useGoogleCalendarAction" yaml:"useGoogleCalendarAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usehubspotaction
	//
	UseHubspotAction *string `field:"optional" json:"useHubspotAction" yaml:"useHubspotAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usehuggingfaceaction
	//
	UseHuggingFaceAction *string `field:"optional" json:"useHuggingFaceAction" yaml:"useHuggingFaceAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useintercomaction
	//
	UseIntercomAction *string `field:"optional" json:"useIntercomAction" yaml:"useIntercomAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usejiraaction
	//
	UseJiraAction *string `field:"optional" json:"useJiraAction" yaml:"useJiraAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-uselinearaction
	//
	UseLinearAction *string `field:"optional" json:"useLinearAction" yaml:"useLinearAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usemcpaction
	//
	UseMcpAction *string `field:"optional" json:"useMcpAction" yaml:"useMcpAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usemondayaction
	//
	UseMondayAction *string `field:"optional" json:"useMondayAction" yaml:"useMondayAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usemsexchangeaction
	//
	UseMsExchangeAction *string `field:"optional" json:"useMsExchangeAction" yaml:"useMsExchangeAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usemsteamsaction
	//
	UseMsTeamsAction *string `field:"optional" json:"useMsTeamsAction" yaml:"useMsTeamsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usenewrelicaction
	//
	UseNewRelicAction *string `field:"optional" json:"useNewRelicAction" yaml:"useNewRelicAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usenotionaction
	//
	UseNotionAction *string `field:"optional" json:"useNotionAction" yaml:"useNotionAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useonedriveaction
	//
	UseOneDriveAction *string `field:"optional" json:"useOneDriveAction" yaml:"useOneDriveAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useopenapiaction
	//
	UseOpenApiAction *string `field:"optional" json:"useOpenApiAction" yaml:"useOpenApiAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usepagerdutyaction
	//
	UsePagerDutyAction *string `field:"optional" json:"usePagerDutyAction" yaml:"usePagerDutyAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usesalesforceaction
	//
	UseSalesforceAction *string `field:"optional" json:"useSalesforceAction" yaml:"useSalesforceAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usesandpglobalenergyaction
	//
	UseSandPGlobalEnergyAction *string `field:"optional" json:"useSandPGlobalEnergyAction" yaml:"useSandPGlobalEnergyAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usesandpgmiaction
	//
	UseSandPgmiAction *string `field:"optional" json:"useSandPgmiAction" yaml:"useSandPgmiAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usesapbillofmaterialaction
	//
	UseSapBillOfMaterialAction *string `field:"optional" json:"useSapBillOfMaterialAction" yaml:"useSapBillOfMaterialAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usesapbusinesspartneraction
	//
	UseSapBusinessPartnerAction *string `field:"optional" json:"useSapBusinessPartnerAction" yaml:"useSapBusinessPartnerAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usesapmaterialstockaction
	//
	UseSapMaterialStockAction *string `field:"optional" json:"useSapMaterialStockAction" yaml:"useSapMaterialStockAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usesapphysicalinventoryaction
	//
	UseSapPhysicalInventoryAction *string `field:"optional" json:"useSapPhysicalInventoryAction" yaml:"useSapPhysicalInventoryAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usesapproductmasterdataaction
	//
	UseSapProductMasterDataAction *string `field:"optional" json:"useSapProductMasterDataAction" yaml:"useSapProductMasterDataAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useservicenowaction
	//
	UseServiceNowAction *string `field:"optional" json:"useServiceNowAction" yaml:"useServiceNowAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usesharepointaction
	//
	UseSharePointAction *string `field:"optional" json:"useSharePointAction" yaml:"useSharePointAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useslackaction
	//
	UseSlackAction *string `field:"optional" json:"useSlackAction" yaml:"useSlackAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usesmartsheetaction
	//
	UseSmartsheetAction *string `field:"optional" json:"useSmartsheetAction" yaml:"useSmartsheetAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usetextractaction
	//
	UseTextractAction *string `field:"optional" json:"useTextractAction" yaml:"useTextractAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usezendeskaction
	//
	UseZendeskAction *string `field:"optional" json:"useZendeskAction" yaml:"useZendeskAction"`
	// The ability to view account SPICE capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-viewaccountspicecapacity
	//
	ViewAccountSpiceCapacity *string `field:"optional" json:"viewAccountSpiceCapacity" yaml:"viewAccountSpiceCapacity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-zendeskaction
	//
	ZendeskAction *string `field:"optional" json:"zendeskAction" yaml:"zendeskAction"`
}

