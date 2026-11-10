package awsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsquicksight"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a custom permissions profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomPermissions := awscdk.Aws_quicksight.NewCfnCustomPermissions(this, jsii.String("MyCfnCustomPermissions"), &CfnCustomPermissionsProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	CustomPermissionsName: jsii.String("customPermissionsName"),
//
//   	// the properties below are optional
//   	Capabilities: &CapabilitiesProperty{
//   		AccessAppsNativeDataStore: jsii.String("accessAppsNativeDataStore"),
//   		Action: jsii.String("action"),
//   		AddOrRunAnomalyDetectionForAnalyses: jsii.String("addOrRunAnomalyDetectionForAnalyses"),
//   		AdobeAnalyticsDataSource: jsii.String("adobeAnalyticsDataSource"),
//   		AmazonBedrockArsAction: jsii.String("amazonBedrockArsAction"),
//   		AmazonBedrockFsAction: jsii.String("amazonBedrockFsAction"),
//   		AmazonBedrockKrsAction: jsii.String("amazonBedrockKrsAction"),
//   		AmazonSThreeAction: jsii.String("amazonSThreeAction"),
//   		Analysis: jsii.String("analysis"),
//   		ApproveFlowShareRequests: jsii.String("approveFlowShareRequests"),
//   		Apps: jsii.String("apps"),
//   		AsanaAction: jsii.String("asanaAction"),
//   		AthenaDataSource: jsii.String("athenaDataSource"),
//   		AuroraDataSource: jsii.String("auroraDataSource"),
//   		Automate: jsii.String("automate"),
//   		BambooHrAction: jsii.String("bambooHrAction"),
//   		BedrockManagedKnowledgeBase: jsii.String("bedrockManagedKnowledgeBase"),
//   		BoxAgentAction: jsii.String("boxAgentAction"),
//   		BoxKnowledgeBase: jsii.String("boxKnowledgeBase"),
//   		BuildCalculatedFieldWithQ: jsii.String("buildCalculatedFieldWithQ"),
//   		CanvaAgentAction: jsii.String("canvaAgentAction"),
//   		ChatAgent: jsii.String("chatAgent"),
//   		ComprehendAction: jsii.String("comprehendAction"),
//   		ComprehendMedicalAction: jsii.String("comprehendMedicalAction"),
//   		ConfluenceAction: jsii.String("confluenceAction"),
//   		ConfluenceKnowledgeBase: jsii.String("confluenceKnowledgeBase"),
//   		CreateAdobeAnalyticsDataSource: jsii.String("createAdobeAnalyticsDataSource"),
//   		CreateAndUpdateAmazonBedrockArsAction: jsii.String("createAndUpdateAmazonBedrockArsAction"),
//   		CreateAndUpdateAmazonBedrockFsAction: jsii.String("createAndUpdateAmazonBedrockFsAction"),
//   		CreateAndUpdateAmazonBedrockKrsAction: jsii.String("createAndUpdateAmazonBedrockKrsAction"),
//   		CreateAndUpdateAmazonSThreeAction: jsii.String("createAndUpdateAmazonSThreeAction"),
//   		CreateAndUpdateApps: jsii.String("createAndUpdateApps"),
//   		CreateAndUpdateAsanaAction: jsii.String("createAndUpdateAsanaAction"),
//   		CreateAndUpdateBambooHrAction: jsii.String("createAndUpdateBambooHrAction"),
//   		CreateAndUpdateBedrockManagedKnowledgeBase: jsii.String("createAndUpdateBedrockManagedKnowledgeBase"),
//   		CreateAndUpdateBoxAgentAction: jsii.String("createAndUpdateBoxAgentAction"),
//   		CreateAndUpdateBoxKnowledgeBase: jsii.String("createAndUpdateBoxKnowledgeBase"),
//   		CreateAndUpdateCanvaAgentAction: jsii.String("createAndUpdateCanvaAgentAction"),
//   		CreateAndUpdateComprehendAction: jsii.String("createAndUpdateComprehendAction"),
//   		CreateAndUpdateComprehendMedicalAction: jsii.String("createAndUpdateComprehendMedicalAction"),
//   		CreateAndUpdateConfluenceAction: jsii.String("createAndUpdateConfluenceAction"),
//   		CreateAndUpdateConfluenceKnowledgeBase: jsii.String("createAndUpdateConfluenceKnowledgeBase"),
//   		CreateAndUpdateDashboardEmailReports: jsii.String("createAndUpdateDashboardEmailReports"),
//   		CreateAndUpdateDatasets: jsii.String("createAndUpdateDatasets"),
//   		CreateAndUpdateDataSources: jsii.String("createAndUpdateDataSources"),
//   		CreateAndUpdateFactSetAction: jsii.String("createAndUpdateFactSetAction"),
//   		CreateAndUpdateGenericHttpAction: jsii.String("createAndUpdateGenericHttpAction"),
//   		CreateAndUpdateGithubAction: jsii.String("createAndUpdateGithubAction"),
//   		CreateAndUpdateGoogleCalendarAction: jsii.String("createAndUpdateGoogleCalendarAction"),
//   		CreateAndUpdateGoogleDriveKnowledgeBase: jsii.String("createAndUpdateGoogleDriveKnowledgeBase"),
//   		CreateAndUpdateHubspotAction: jsii.String("createAndUpdateHubspotAction"),
//   		CreateAndUpdateHuggingFaceAction: jsii.String("createAndUpdateHuggingFaceAction"),
//   		CreateAndUpdateIdcKnowledgeBase: jsii.String("createAndUpdateIdcKnowledgeBase"),
//   		CreateAndUpdateIntercomAction: jsii.String("createAndUpdateIntercomAction"),
//   		CreateAndUpdateJiraAction: jsii.String("createAndUpdateJiraAction"),
//   		CreateAndUpdateKnowledgeBases: jsii.String("createAndUpdateKnowledgeBases"),
//   		CreateAndUpdateLinearAction: jsii.String("createAndUpdateLinearAction"),
//   		CreateAndUpdateMcpAction: jsii.String("createAndUpdateMcpAction"),
//   		CreateAndUpdateMondayAction: jsii.String("createAndUpdateMondayAction"),
//   		CreateAndUpdateMsExchangeAction: jsii.String("createAndUpdateMsExchangeAction"),
//   		CreateAndUpdateMsTeamsAction: jsii.String("createAndUpdateMsTeamsAction"),
//   		CreateAndUpdateNewRelicAction: jsii.String("createAndUpdateNewRelicAction"),
//   		CreateAndUpdateNotionAction: jsii.String("createAndUpdateNotionAction"),
//   		CreateAndUpdateOneDriveAction: jsii.String("createAndUpdateOneDriveAction"),
//   		CreateAndUpdateOneDriveKnowledgeBase: jsii.String("createAndUpdateOneDriveKnowledgeBase"),
//   		CreateAndUpdateOpenApiAction: jsii.String("createAndUpdateOpenApiAction"),
//   		CreateAndUpdatePagerDutyAction: jsii.String("createAndUpdatePagerDutyAction"),
//   		CreateAndUpdateQBusinessKnowledgeBase: jsii.String("createAndUpdateQBusinessKnowledgeBase"),
//   		CreateAndUpdateS3KnowledgeBase: jsii.String("createAndUpdateS3KnowledgeBase"),
//   		CreateAndUpdateSalesforceAction: jsii.String("createAndUpdateSalesforceAction"),
//   		CreateAndUpdateSandPGlobalEnergyAction: jsii.String("createAndUpdateSandPGlobalEnergyAction"),
//   		CreateAndUpdateSandPgmiAction: jsii.String("createAndUpdateSandPgmiAction"),
//   		CreateAndUpdateSapBillOfMaterialAction: jsii.String("createAndUpdateSapBillOfMaterialAction"),
//   		CreateAndUpdateSapBusinessPartnerAction: jsii.String("createAndUpdateSapBusinessPartnerAction"),
//   		CreateAndUpdateSapMaterialStockAction: jsii.String("createAndUpdateSapMaterialStockAction"),
//   		CreateAndUpdateSapPhysicalInventoryAction: jsii.String("createAndUpdateSapPhysicalInventoryAction"),
//   		CreateAndUpdateSapProductMasterDataAction: jsii.String("createAndUpdateSapProductMasterDataAction"),
//   		CreateAndUpdateServiceNowAction: jsii.String("createAndUpdateServiceNowAction"),
//   		CreateAndUpdateSharePointAction: jsii.String("createAndUpdateSharePointAction"),
//   		CreateAndUpdateSharePointKnowledgeBase: jsii.String("createAndUpdateSharePointKnowledgeBase"),
//   		CreateAndUpdateSlackAction: jsii.String("createAndUpdateSlackAction"),
//   		CreateAndUpdateSmartsheetAction: jsii.String("createAndUpdateSmartsheetAction"),
//   		CreateAndUpdateTextractAction: jsii.String("createAndUpdateTextractAction"),
//   		CreateAndUpdateThemes: jsii.String("createAndUpdateThemes"),
//   		CreateAndUpdateThresholdAlerts: jsii.String("createAndUpdateThresholdAlerts"),
//   		CreateAndUpdateWebCrawlerKnowledgeBase: jsii.String("createAndUpdateWebCrawlerKnowledgeBase"),
//   		CreateAndUpdateZendeskAction: jsii.String("createAndUpdateZendeskAction"),
//   		CreateAthenaDataSource: jsii.String("createAthenaDataSource"),
//   		CreateAuroraDataSource: jsii.String("createAuroraDataSource"),
//   		CreateChatAgents: jsii.String("createChatAgents"),
//   		CreateDashboardExecutiveSummaryWithQ: jsii.String("createDashboardExecutiveSummaryWithQ"),
//   		CreateDatabricksDataSource: jsii.String("createDatabricksDataSource"),
//   		CreateDb2DataSource: jsii.String("createDb2DataSource"),
//   		CreateDenodoDataSource: jsii.String("createDenodoDataSource"),
//   		CreateDocumentDbDataSource: jsii.String("createDocumentDbDataSource"),
//   		CreateDremioDataSource: jsii.String("createDremioDataSource"),
//   		CreateDynamoDbDataSource: jsii.String("createDynamoDbDataSource"),
//   		CreateExasolDataSource: jsii.String("createExasolDataSource"),
//   		CreateFileDataSource: jsii.String("createFileDataSource"),
//   		CreateGitHubDataSource: jsii.String("createGitHubDataSource"),
//   		CreateGoogleAnalyticsDataSource: jsii.String("createGoogleAnalyticsDataSource"),
//   		CreateGoogleBigQueryDataSource: jsii.String("createGoogleBigQueryDataSource"),
//   		CreateGoogleSheetsDataSource: jsii.String("createGoogleSheetsDataSource"),
//   		CreateImpalaDataSource: jsii.String("createImpalaDataSource"),
//   		CreateJiraDataSource: jsii.String("createJiraDataSource"),
//   		CreateMariaDbDataSource: jsii.String("createMariaDbDataSource"),
//   		CreateMongoAtlasDataSource: jsii.String("createMongoAtlasDataSource"),
//   		CreateMongoDbDataSource: jsii.String("createMongoDbDataSource"),
//   		CreateMySqlDataSource: jsii.String("createMySqlDataSource"),
//   		CreateOpenSearchDataSource: jsii.String("createOpenSearchDataSource"),
//   		CreateOracleDataSource: jsii.String("createOracleDataSource"),
//   		CreatePayPalDataSource: jsii.String("createPayPalDataSource"),
//   		CreatePostgreSqlDataSource: jsii.String("createPostgreSqlDataSource"),
//   		CreatePrestoDataSource: jsii.String("createPrestoDataSource"),
//   		CreateRadiantDataSource: jsii.String("createRadiantDataSource"),
//   		CreateRdsDataSource: jsii.String("createRdsDataSource"),
//   		CreateRedshiftAutoDiscoveredDataSource: jsii.String("createRedshiftAutoDiscoveredDataSource"),
//   		CreateRedshiftManualDataSource: jsii.String("createRedshiftManualDataSource"),
//   		CreateS3AnalyticsDataSource: jsii.String("createS3AnalyticsDataSource"),
//   		CreateS3DataSource: jsii.String("createS3DataSource"),
//   		CreateS3TablesDataSource: jsii.String("createS3TablesDataSource"),
//   		CreateSalesforceDataSource: jsii.String("createSalesforceDataSource"),
//   		CreateSapHanaDataSource: jsii.String("createSapHanaDataSource"),
//   		CreateServiceNowDataSource: jsii.String("createServiceNowDataSource"),
//   		CreateSharedFolders: jsii.String("createSharedFolders"),
//   		CreateSnowflakeDataSource: jsii.String("createSnowflakeDataSource"),
//   		CreateSpaces: jsii.String("createSpaces"),
//   		CreateSparkDataSource: jsii.String("createSparkDataSource"),
//   		CreateSpiceDataset: jsii.String("createSpiceDataset"),
//   		CreateSqlServerDataSource: jsii.String("createSqlServerDataSource"),
//   		CreateSquareDataSource: jsii.String("createSquareDataSource"),
//   		CreateStarburstDataSource: jsii.String("createStarburstDataSource"),
//   		CreateTeradataDataSource: jsii.String("createTeradataDataSource"),
//   		CreateTimestreamDataSource: jsii.String("createTimestreamDataSource"),
//   		CreateTrinoDataSource: jsii.String("createTrinoDataSource"),
//   		CreateTwitterDataSource: jsii.String("createTwitterDataSource"),
//   		Dashboard: jsii.String("dashboard"),
//   		DatabricksDataSource: jsii.String("databricksDataSource"),
//   		Db2DataSource: jsii.String("db2DataSource"),
//   		DenodoDataSource: jsii.String("denodoDataSource"),
//   		DocumentDbDataSource: jsii.String("documentDbDataSource"),
//   		DremioDataSource: jsii.String("dremioDataSource"),
//   		DynamoDbDataSource: jsii.String("dynamoDbDataSource"),
//   		EditVisualWithQ: jsii.String("editVisualWithQ"),
//   		ExasolDataSource: jsii.String("exasolDataSource"),
//   		ExportToCsv: jsii.String("exportToCsv"),
//   		ExportToCsvInScheduledReports: jsii.String("exportToCsvInScheduledReports"),
//   		ExportToExcel: jsii.String("exportToExcel"),
//   		ExportToExcelInScheduledReports: jsii.String("exportToExcelInScheduledReports"),
//   		ExportToPdf: jsii.String("exportToPdf"),
//   		ExportToPdfInScheduledReports: jsii.String("exportToPdfInScheduledReports"),
//   		Extension: jsii.String("extension"),
//   		FactSetAction: jsii.String("factSetAction"),
//   		FileDataSource: jsii.String("fileDataSource"),
//   		Flow: jsii.String("flow"),
//   		GenericHttpAction: jsii.String("genericHttpAction"),
//   		GithubAction: jsii.String("githubAction"),
//   		GitHubDataSource: jsii.String("gitHubDataSource"),
//   		GoogleAnalyticsDataSource: jsii.String("googleAnalyticsDataSource"),
//   		GoogleBigQueryDataSource: jsii.String("googleBigQueryDataSource"),
//   		GoogleCalendarAction: jsii.String("googleCalendarAction"),
//   		GoogleDriveKnowledgeBase: jsii.String("googleDriveKnowledgeBase"),
//   		GoogleSheetsDataSource: jsii.String("googleSheetsDataSource"),
//   		HubspotAction: jsii.String("hubspotAction"),
//   		HuggingFaceAction: jsii.String("huggingFaceAction"),
//   		IdcKnowledgeBase: jsii.String("idcKnowledgeBase"),
//   		ImpalaDataSource: jsii.String("impalaDataSource"),
//   		InboundEmailTrigger: jsii.String("inboundEmailTrigger"),
//   		IncludeContentInScheduledReportsEmail: jsii.String("includeContentInScheduledReportsEmail"),
//   		IntercomAction: jsii.String("intercomAction"),
//   		InvokeAppsAiInference: jsii.String("invokeAppsAiInference"),
//   		JiraAction: jsii.String("jiraAction"),
//   		JiraDataSource: jsii.String("jiraDataSource"),
//   		KnowledgeBase: jsii.String("knowledgeBase"),
//   		LinearAction: jsii.String("linearAction"),
//   		ManageSharedFolders: jsii.String("manageSharedFolders"),
//   		MariaDbDataSource: jsii.String("mariaDbDataSource"),
//   		McpAction: jsii.String("mcpAction"),
//   		MondayAction: jsii.String("mondayAction"),
//   		MongoAtlasDataSource: jsii.String("mongoAtlasDataSource"),
//   		MongoDbDataSource: jsii.String("mongoDbDataSource"),
//   		MsExchangeAction: jsii.String("msExchangeAction"),
//   		MsTeamsAction: jsii.String("msTeamsAction"),
//   		MySqlDataSource: jsii.String("mySqlDataSource"),
//   		NewRelicAction: jsii.String("newRelicAction"),
//   		NotionAction: jsii.String("notionAction"),
//   		OneDriveAction: jsii.String("oneDriveAction"),
//   		OneDriveKnowledgeBase: jsii.String("oneDriveKnowledgeBase"),
//   		OpenApiAction: jsii.String("openApiAction"),
//   		OpenSearchDataSource: jsii.String("openSearchDataSource"),
//   		OracleDataSource: jsii.String("oracleDataSource"),
//   		PagerDutyAction: jsii.String("pagerDutyAction"),
//   		PayPalDataSource: jsii.String("payPalDataSource"),
//   		PerformFlowUiTask: jsii.String("performFlowUiTask"),
//   		PostgreSqlDataSource: jsii.String("postgreSqlDataSource"),
//   		PrestoDataSource: jsii.String("prestoDataSource"),
//   		PrintReports: jsii.String("printReports"),
//   		PublishWithoutApproval: jsii.String("publishWithoutApproval"),
//   		QBusinessKnowledgeBase: jsii.String("qBusinessKnowledgeBase"),
//   		QuickEventTrigger: jsii.String("quickEventTrigger"),
//   		RadiantDataSource: jsii.String("radiantDataSource"),
//   		RdsDataSource: jsii.String("rdsDataSource"),
//   		RedshiftAutoDiscoveredDataSource: jsii.String("redshiftAutoDiscoveredDataSource"),
//   		RedshiftManualDataSource: jsii.String("redshiftManualDataSource"),
//   		RenameSharedFolders: jsii.String("renameSharedFolders"),
//   		Research: jsii.String("research"),
//   		S3AnalyticsDataSource: jsii.String("s3AnalyticsDataSource"),
//   		S3DataSource: jsii.String("s3DataSource"),
//   		S3KnowledgeBase: jsii.String("s3KnowledgeBase"),
//   		S3TablesDataSource: jsii.String("s3TablesDataSource"),
//   		SalesforceAction: jsii.String("salesforceAction"),
//   		SalesforceDataSource: jsii.String("salesforceDataSource"),
//   		SandPGlobalEnergyAction: jsii.String("sandPGlobalEnergyAction"),
//   		SandPgmiAction: jsii.String("sandPgmiAction"),
//   		SapBillOfMaterialAction: jsii.String("sapBillOfMaterialAction"),
//   		SapBusinessPartnerAction: jsii.String("sapBusinessPartnerAction"),
//   		SapHanaDataSource: jsii.String("sapHanaDataSource"),
//   		SapMaterialStockAction: jsii.String("sapMaterialStockAction"),
//   		SapPhysicalInventoryAction: jsii.String("sapPhysicalInventoryAction"),
//   		SapProductMasterDataAction: jsii.String("sapProductMasterDataAction"),
//   		ScheduleTrigger: jsii.String("scheduleTrigger"),
//   		ServiceNowAction: jsii.String("serviceNowAction"),
//   		ServiceNowDataSource: jsii.String("serviceNowDataSource"),
//   		ShareAdobeAnalyticsDataSource: jsii.String("shareAdobeAnalyticsDataSource"),
//   		ShareAmazonBedrockArsAction: jsii.String("shareAmazonBedrockArsAction"),
//   		ShareAmazonBedrockFsAction: jsii.String("shareAmazonBedrockFsAction"),
//   		ShareAmazonBedrockKrsAction: jsii.String("shareAmazonBedrockKrsAction"),
//   		ShareAmazonSThreeAction: jsii.String("shareAmazonSThreeAction"),
//   		ShareAnalyses: jsii.String("shareAnalyses"),
//   		ShareApps: jsii.String("shareApps"),
//   		ShareAsanaAction: jsii.String("shareAsanaAction"),
//   		ShareAthenaDataSource: jsii.String("shareAthenaDataSource"),
//   		ShareAuroraDataSource: jsii.String("shareAuroraDataSource"),
//   		ShareBambooHrAction: jsii.String("shareBambooHrAction"),
//   		ShareBedrockManagedKnowledgeBase: jsii.String("shareBedrockManagedKnowledgeBase"),
//   		ShareBoxAgentAction: jsii.String("shareBoxAgentAction"),
//   		ShareBoxKnowledgeBase: jsii.String("shareBoxKnowledgeBase"),
//   		ShareCanvaAgentAction: jsii.String("shareCanvaAgentAction"),
//   		ShareChatAgents: jsii.String("shareChatAgents"),
//   		ShareComprehendAction: jsii.String("shareComprehendAction"),
//   		ShareComprehendMedicalAction: jsii.String("shareComprehendMedicalAction"),
//   		ShareConfluenceAction: jsii.String("shareConfluenceAction"),
//   		ShareConfluenceKnowledgeBase: jsii.String("shareConfluenceKnowledgeBase"),
//   		ShareDashboards: jsii.String("shareDashboards"),
//   		ShareDatabricksDataSource: jsii.String("shareDatabricksDataSource"),
//   		ShareDatasets: jsii.String("shareDatasets"),
//   		ShareDataSources: jsii.String("shareDataSources"),
//   		ShareDb2DataSource: jsii.String("shareDb2DataSource"),
//   		ShareDenodoDataSource: jsii.String("shareDenodoDataSource"),
//   		ShareDocumentDbDataSource: jsii.String("shareDocumentDbDataSource"),
//   		ShareDremioDataSource: jsii.String("shareDremioDataSource"),
//   		ShareDynamoDbDataSource: jsii.String("shareDynamoDbDataSource"),
//   		ShareExasolDataSource: jsii.String("shareExasolDataSource"),
//   		ShareFactSetAction: jsii.String("shareFactSetAction"),
//   		ShareFileDataSource: jsii.String("shareFileDataSource"),
//   		ShareGenericHttpAction: jsii.String("shareGenericHttpAction"),
//   		ShareGithubAction: jsii.String("shareGithubAction"),
//   		ShareGitHubDataSource: jsii.String("shareGitHubDataSource"),
//   		ShareGoogleAnalyticsDataSource: jsii.String("shareGoogleAnalyticsDataSource"),
//   		ShareGoogleBigQueryDataSource: jsii.String("shareGoogleBigQueryDataSource"),
//   		ShareGoogleCalendarAction: jsii.String("shareGoogleCalendarAction"),
//   		ShareGoogleDriveKnowledgeBase: jsii.String("shareGoogleDriveKnowledgeBase"),
//   		ShareGoogleSheetsDataSource: jsii.String("shareGoogleSheetsDataSource"),
//   		ShareHubspotAction: jsii.String("shareHubspotAction"),
//   		ShareHuggingFaceAction: jsii.String("shareHuggingFaceAction"),
//   		ShareIdcKnowledgeBase: jsii.String("shareIdcKnowledgeBase"),
//   		ShareImpalaDataSource: jsii.String("shareImpalaDataSource"),
//   		ShareIntercomAction: jsii.String("shareIntercomAction"),
//   		ShareJiraAction: jsii.String("shareJiraAction"),
//   		ShareJiraDataSource: jsii.String("shareJiraDataSource"),
//   		ShareKnowledgeBases: jsii.String("shareKnowledgeBases"),
//   		ShareLinearAction: jsii.String("shareLinearAction"),
//   		ShareMariaDbDataSource: jsii.String("shareMariaDbDataSource"),
//   		ShareMcpAction: jsii.String("shareMcpAction"),
//   		ShareMondayAction: jsii.String("shareMondayAction"),
//   		ShareMongoAtlasDataSource: jsii.String("shareMongoAtlasDataSource"),
//   		ShareMongoDbDataSource: jsii.String("shareMongoDbDataSource"),
//   		ShareMsExchangeAction: jsii.String("shareMsExchangeAction"),
//   		ShareMsTeamsAction: jsii.String("shareMsTeamsAction"),
//   		ShareMySqlDataSource: jsii.String("shareMySqlDataSource"),
//   		ShareNewRelicAction: jsii.String("shareNewRelicAction"),
//   		ShareNotionAction: jsii.String("shareNotionAction"),
//   		ShareOneDriveAction: jsii.String("shareOneDriveAction"),
//   		ShareOneDriveKnowledgeBase: jsii.String("shareOneDriveKnowledgeBase"),
//   		ShareOpenApiAction: jsii.String("shareOpenApiAction"),
//   		ShareOpenSearchDataSource: jsii.String("shareOpenSearchDataSource"),
//   		ShareOracleDataSource: jsii.String("shareOracleDataSource"),
//   		SharePagerDutyAction: jsii.String("sharePagerDutyAction"),
//   		SharePayPalDataSource: jsii.String("sharePayPalDataSource"),
//   		SharePointAction: jsii.String("sharePointAction"),
//   		SharePointKnowledgeBase: jsii.String("sharePointKnowledgeBase"),
//   		SharePostgreSqlDataSource: jsii.String("sharePostgreSqlDataSource"),
//   		SharePrestoDataSource: jsii.String("sharePrestoDataSource"),
//   		ShareQBusinessKnowledgeBase: jsii.String("shareQBusinessKnowledgeBase"),
//   		ShareRadiantDataSource: jsii.String("shareRadiantDataSource"),
//   		ShareRdsDataSource: jsii.String("shareRdsDataSource"),
//   		ShareRedshiftAutoDiscoveredDataSource: jsii.String("shareRedshiftAutoDiscoveredDataSource"),
//   		ShareRedshiftManualDataSource: jsii.String("shareRedshiftManualDataSource"),
//   		ShareS3AnalyticsDataSource: jsii.String("shareS3AnalyticsDataSource"),
//   		ShareS3DataSource: jsii.String("shareS3DataSource"),
//   		ShareS3KnowledgeBase: jsii.String("shareS3KnowledgeBase"),
//   		ShareS3TablesDataSource: jsii.String("shareS3TablesDataSource"),
//   		ShareSalesforceAction: jsii.String("shareSalesforceAction"),
//   		ShareSalesforceDataSource: jsii.String("shareSalesforceDataSource"),
//   		ShareSandPGlobalEnergyAction: jsii.String("shareSandPGlobalEnergyAction"),
//   		ShareSandPgmiAction: jsii.String("shareSandPgmiAction"),
//   		ShareSapBillOfMaterialAction: jsii.String("shareSapBillOfMaterialAction"),
//   		ShareSapBusinessPartnerAction: jsii.String("shareSapBusinessPartnerAction"),
//   		ShareSapHanaDataSource: jsii.String("shareSapHanaDataSource"),
//   		ShareSapMaterialStockAction: jsii.String("shareSapMaterialStockAction"),
//   		ShareSapPhysicalInventoryAction: jsii.String("shareSapPhysicalInventoryAction"),
//   		ShareSapProductMasterDataAction: jsii.String("shareSapProductMasterDataAction"),
//   		ShareServiceNowAction: jsii.String("shareServiceNowAction"),
//   		ShareServiceNowDataSource: jsii.String("shareServiceNowDataSource"),
//   		ShareSharePointAction: jsii.String("shareSharePointAction"),
//   		ShareSharePointKnowledgeBase: jsii.String("shareSharePointKnowledgeBase"),
//   		ShareSlackAction: jsii.String("shareSlackAction"),
//   		ShareSmartsheetAction: jsii.String("shareSmartsheetAction"),
//   		ShareSnowflakeDataSource: jsii.String("shareSnowflakeDataSource"),
//   		ShareSpaces: jsii.String("shareSpaces"),
//   		ShareSparkDataSource: jsii.String("shareSparkDataSource"),
//   		ShareSqlServerDataSource: jsii.String("shareSqlServerDataSource"),
//   		ShareSquareDataSource: jsii.String("shareSquareDataSource"),
//   		ShareStarburstDataSource: jsii.String("shareStarburstDataSource"),
//   		ShareTeradataDataSource: jsii.String("shareTeradataDataSource"),
//   		ShareTextractAction: jsii.String("shareTextractAction"),
//   		ShareTimestreamDataSource: jsii.String("shareTimestreamDataSource"),
//   		ShareTrinoDataSource: jsii.String("shareTrinoDataSource"),
//   		ShareTwitterDataSource: jsii.String("shareTwitterDataSource"),
//   		ShareWebCrawlerKnowledgeBase: jsii.String("shareWebCrawlerKnowledgeBase"),
//   		ShareZendeskAction: jsii.String("shareZendeskAction"),
//   		SlackAction: jsii.String("slackAction"),
//   		SmartsheetAction: jsii.String("smartsheetAction"),
//   		SnowflakeDataSource: jsii.String("snowflakeDataSource"),
//   		Space: jsii.String("space"),
//   		SparkDataSource: jsii.String("sparkDataSource"),
//   		SqlServerDataSource: jsii.String("sqlServerDataSource"),
//   		SquareDataSource: jsii.String("squareDataSource"),
//   		StarburstDataSource: jsii.String("starburstDataSource"),
//   		SubscribeDashboardEmailReports: jsii.String("subscribeDashboardEmailReports"),
//   		TeradataDataSource: jsii.String("teradataDataSource"),
//   		TextractAction: jsii.String("textractAction"),
//   		TimestreamDataSource: jsii.String("timestreamDataSource"),
//   		Topic: jsii.String("topic"),
//   		Trigger: jsii.String("trigger"),
//   		TrinoDataSource: jsii.String("trinoDataSource"),
//   		TwitterDataSource: jsii.String("twitterDataSource"),
//   		UpdateAdobeAnalyticsDataSource: jsii.String("updateAdobeAnalyticsDataSource"),
//   		UpdateAthenaDataSource: jsii.String("updateAthenaDataSource"),
//   		UpdateAuroraDataSource: jsii.String("updateAuroraDataSource"),
//   		UpdateDatabricksDataSource: jsii.String("updateDatabricksDataSource"),
//   		UpdateDb2DataSource: jsii.String("updateDb2DataSource"),
//   		UpdateDenodoDataSource: jsii.String("updateDenodoDataSource"),
//   		UpdateDocumentDbDataSource: jsii.String("updateDocumentDbDataSource"),
//   		UpdateDremioDataSource: jsii.String("updateDremioDataSource"),
//   		UpdateDynamoDbDataSource: jsii.String("updateDynamoDbDataSource"),
//   		UpdateExasolDataSource: jsii.String("updateExasolDataSource"),
//   		UpdateFileDataSource: jsii.String("updateFileDataSource"),
//   		UpdateGitHubDataSource: jsii.String("updateGitHubDataSource"),
//   		UpdateGoogleAnalyticsDataSource: jsii.String("updateGoogleAnalyticsDataSource"),
//   		UpdateGoogleBigQueryDataSource: jsii.String("updateGoogleBigQueryDataSource"),
//   		UpdateGoogleSheetsDataSource: jsii.String("updateGoogleSheetsDataSource"),
//   		UpdateImpalaDataSource: jsii.String("updateImpalaDataSource"),
//   		UpdateJiraDataSource: jsii.String("updateJiraDataSource"),
//   		UpdateMariaDbDataSource: jsii.String("updateMariaDbDataSource"),
//   		UpdateMongoAtlasDataSource: jsii.String("updateMongoAtlasDataSource"),
//   		UpdateMongoDbDataSource: jsii.String("updateMongoDbDataSource"),
//   		UpdateMySqlDataSource: jsii.String("updateMySqlDataSource"),
//   		UpdateOpenSearchDataSource: jsii.String("updateOpenSearchDataSource"),
//   		UpdateOracleDataSource: jsii.String("updateOracleDataSource"),
//   		UpdatePayPalDataSource: jsii.String("updatePayPalDataSource"),
//   		UpdatePostgreSqlDataSource: jsii.String("updatePostgreSqlDataSource"),
//   		UpdatePrestoDataSource: jsii.String("updatePrestoDataSource"),
//   		UpdateRadiantDataSource: jsii.String("updateRadiantDataSource"),
//   		UpdateRdsDataSource: jsii.String("updateRdsDataSource"),
//   		UpdateRedshiftAutoDiscoveredDataSource: jsii.String("updateRedshiftAutoDiscoveredDataSource"),
//   		UpdateRedshiftManualDataSource: jsii.String("updateRedshiftManualDataSource"),
//   		UpdateS3AnalyticsDataSource: jsii.String("updateS3AnalyticsDataSource"),
//   		UpdateS3DataSource: jsii.String("updateS3DataSource"),
//   		UpdateS3TablesDataSource: jsii.String("updateS3TablesDataSource"),
//   		UpdateSalesforceDataSource: jsii.String("updateSalesforceDataSource"),
//   		UpdateSapHanaDataSource: jsii.String("updateSapHanaDataSource"),
//   		UpdateServiceNowDataSource: jsii.String("updateServiceNowDataSource"),
//   		UpdateSnowflakeDataSource: jsii.String("updateSnowflakeDataSource"),
//   		UpdateSparkDataSource: jsii.String("updateSparkDataSource"),
//   		UpdateSqlServerDataSource: jsii.String("updateSqlServerDataSource"),
//   		UpdateSquareDataSource: jsii.String("updateSquareDataSource"),
//   		UpdateStarburstDataSource: jsii.String("updateStarburstDataSource"),
//   		UpdateTeradataDataSource: jsii.String("updateTeradataDataSource"),
//   		UpdateTimestreamDataSource: jsii.String("updateTimestreamDataSource"),
//   		UpdateTrinoDataSource: jsii.String("updateTrinoDataSource"),
//   		UpdateTwitterDataSource: jsii.String("updateTwitterDataSource"),
//   		UseAgentWebSearch: jsii.String("useAgentWebSearch"),
//   		UseAmazonBedrockArsAction: jsii.String("useAmazonBedrockArsAction"),
//   		UseAmazonBedrockFsAction: jsii.String("useAmazonBedrockFsAction"),
//   		UseAmazonBedrockKrsAction: jsii.String("useAmazonBedrockKrsAction"),
//   		UseAmazonSThreeAction: jsii.String("useAmazonSThreeAction"),
//   		UseAsanaAction: jsii.String("useAsanaAction"),
//   		UseBambooHrAction: jsii.String("useBambooHrAction"),
//   		UseBedrockManagedKnowledgeBase: jsii.String("useBedrockManagedKnowledgeBase"),
//   		UseBedrockModels: jsii.String("useBedrockModels"),
//   		UseBoxAgentAction: jsii.String("useBoxAgentAction"),
//   		UseBoxKnowledgeBase: jsii.String("useBoxKnowledgeBase"),
//   		UseBrowserExtension: jsii.String("useBrowserExtension"),
//   		UseCanvaAgentAction: jsii.String("useCanvaAgentAction"),
//   		UseComprehendAction: jsii.String("useComprehendAction"),
//   		UseComprehendMedicalAction: jsii.String("useComprehendMedicalAction"),
//   		UseConfluenceAction: jsii.String("useConfluenceAction"),
//   		UseConfluenceKnowledgeBase: jsii.String("useConfluenceKnowledgeBase"),
//   		UseExcelAddInExtension: jsii.String("useExcelAddInExtension"),
//   		UseFactSetAction: jsii.String("useFactSetAction"),
//   		UseGenericHttpAction: jsii.String("useGenericHttpAction"),
//   		UseGithubAction: jsii.String("useGithubAction"),
//   		UseGoogleCalendarAction: jsii.String("useGoogleCalendarAction"),
//   		UseGoogleDriveKnowledgeBase: jsii.String("useGoogleDriveKnowledgeBase"),
//   		UseHubspotAction: jsii.String("useHubspotAction"),
//   		UseHuggingFaceAction: jsii.String("useHuggingFaceAction"),
//   		UseIdcKnowledgeBase: jsii.String("useIdcKnowledgeBase"),
//   		UseIntercomAction: jsii.String("useIntercomAction"),
//   		UseJiraAction: jsii.String("useJiraAction"),
//   		UseLinearAction: jsii.String("useLinearAction"),
//   		UseMcpAction: jsii.String("useMcpAction"),
//   		UseMondayAction: jsii.String("useMondayAction"),
//   		UseMsExchangeAction: jsii.String("useMsExchangeAction"),
//   		UseMsTeamsAction: jsii.String("useMsTeamsAction"),
//   		UseNewRelicAction: jsii.String("useNewRelicAction"),
//   		UseNotionAction: jsii.String("useNotionAction"),
//   		UseOneDriveAction: jsii.String("useOneDriveAction"),
//   		UseOneDriveKnowledgeBase: jsii.String("useOneDriveKnowledgeBase"),
//   		UseOpenApiAction: jsii.String("useOpenApiAction"),
//   		UseOutlookAddInExtension: jsii.String("useOutlookAddInExtension"),
//   		UsePagerDutyAction: jsii.String("usePagerDutyAction"),
//   		UsePowerpointAddInExtension: jsii.String("usePowerpointAddInExtension"),
//   		UseQBusinessKnowledgeBase: jsii.String("useQBusinessKnowledgeBase"),
//   		UseS3KnowledgeBase: jsii.String("useS3KnowledgeBase"),
//   		UseSalesforceAction: jsii.String("useSalesforceAction"),
//   		UseSandPGlobalEnergyAction: jsii.String("useSandPGlobalEnergyAction"),
//   		UseSandPgmiAction: jsii.String("useSandPgmiAction"),
//   		UseSapBillOfMaterialAction: jsii.String("useSapBillOfMaterialAction"),
//   		UseSapBusinessPartnerAction: jsii.String("useSapBusinessPartnerAction"),
//   		UseSapMaterialStockAction: jsii.String("useSapMaterialStockAction"),
//   		UseSapPhysicalInventoryAction: jsii.String("useSapPhysicalInventoryAction"),
//   		UseSapProductMasterDataAction: jsii.String("useSapProductMasterDataAction"),
//   		UseServiceNowAction: jsii.String("useServiceNowAction"),
//   		UseSharePointAction: jsii.String("useSharePointAction"),
//   		UseSharePointKnowledgeBase: jsii.String("useSharePointKnowledgeBase"),
//   		UseSlackAction: jsii.String("useSlackAction"),
//   		UseSmartsheetAction: jsii.String("useSmartsheetAction"),
//   		UseTextractAction: jsii.String("useTextractAction"),
//   		UseWebCrawlerKnowledgeBase: jsii.String("useWebCrawlerKnowledgeBase"),
//   		UseWordAddInExtension: jsii.String("useWordAddInExtension"),
//   		UseZendeskAction: jsii.String("useZendeskAction"),
//   		ViewAccountSpiceCapacity: jsii.String("viewAccountSpiceCapacity"),
//   		WebCrawlerKnowledgeBase: jsii.String("webCrawlerKnowledgeBase"),
//   		ZendeskAction: jsii.String("zendeskAction"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-custompermissions.html
//
type CfnCustomPermissions interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsquicksight.ICustomPermissionsRef
	awscdk.ITaggableV2
	// The Amazon Resource Name (ARN) of the custom permissions profile.
	AttrArn() *string
	// The ID of the AWS account that contains the custom permission configuration that you want to update.
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	// A set of actions in the custom permissions profile.
	Capabilities() interface{}
	SetCapabilities(val interface{})
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnPropertyNames() *map[string]*string
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The name of the custom permissions profile.
	CustomPermissionsName() *string
	SetCustomPermissionsName(val *string)
	// A reference to a CustomPermissions resource.
	CustomPermissionsRef() *interfacesawsquicksight.CustomPermissionsReference
	Env() *interfaces.ResourceEnvironment
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags to associate with the custom permissions profile.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This method has been renamed to `addResourceDependency` to more clearly
	// set it apart from `construct.node.addDependency`. See the documentation
	// of that function for more details.
	// Deprecated: Use `addResourceDependency` instead.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method has been renamed to `addResourceDependency`, which makes it
	// more clear that this method operates at a different level from the
	// construct-level `construct.node.addDependency()` mechanism.
	// Deprecated: Use `addResourceDependency` instead.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method only adds dependencies between L1 resources. If you are
	// looking for a generic construct-to-construct dependency mechanism that works
	// for all constructs including L2s, use `construct.node.addDependency` instead.
	AddResourceDependency(target awscdk.CfnResource, reason *string)
	// Sets the cross-stack reference strength for this resource.
	//
	// When set, any cross-stack reference to this resource will use the specified
	// strength instead of the global default from the consuming stack's context.
	ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength)
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	CfnPropertyName(cdkPropertyName *string) *string
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources and stacks this resource depends on.
	//
	// For resources depended on directly, returns the `CfnResource` object. For
	// dependencies on other stacks, returns the `Stack` object. The order of the
	// array is not guaranteed.
	ObtainDependencies() *[]interface{}
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	// Deprecated: Use `removeResourceDependency` instead.
	RemoveDependency(target awscdk.CfnResource)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveResourceDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnCustomPermissions
type jsiiProxy_CfnCustomPermissions struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsquicksightICustomPermissionsRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnCustomPermissions) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) Capabilities() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) CfnPropertyNames() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"cfnPropertyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) CustomPermissionsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPermissionsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) CustomPermissionsRef() *interfacesawsquicksight.CustomPermissionsReference {
	var returns *interfacesawsquicksight.CustomPermissionsReference
	_jsii_.Get(
		j,
		"customPermissionsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissions) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::QuickSight::CustomPermissions`.
func NewCfnCustomPermissions(scope constructs.Construct, id *string, props *CfnCustomPermissionsProps) CfnCustomPermissions {
	_init_.Initialize()

	if err := validateNewCfnCustomPermissionsParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCustomPermissions{}

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnCustomPermissions",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::QuickSight::CustomPermissions`.
func NewCfnCustomPermissions_Override(c CfnCustomPermissions, scope constructs.Construct, id *string, props *CfnCustomPermissionsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnCustomPermissions",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCustomPermissions)SetAwsAccountId(val *string) {
	if err := j.validateSetAwsAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnCustomPermissions)SetCapabilities(val interface{}) {
	if err := j.validateSetCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capabilities",
		val,
	)
}

func (j *jsiiProxy_CfnCustomPermissions)SetCustomPermissionsName(val *string) {
	if err := j.validateSetCustomPermissionsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customPermissionsName",
		val,
	)
}

func (j *jsiiProxy_CfnCustomPermissions)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func CfnCustomPermissions_ArnForCustomPermissions(resource interfacesawsquicksight.ICustomPermissionsRef) *string {
	_init_.Initialize()

	if err := validateCfnCustomPermissions_ArnForCustomPermissionsParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnCustomPermissions",
		"arnForCustomPermissions",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnCustomPermissions.
func CfnCustomPermissions_IsCfnCustomPermissions(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomPermissions_IsCfnCustomPermissionsParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnCustomPermissions",
		"isCfnCustomPermissions",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCustomPermissions_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomPermissions_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnCustomPermissions",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnCustomPermissions_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomPermissions_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnCustomPermissions",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnCustomPermissions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomPermissions_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnCustomPermissions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCustomPermissions_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_quicksight.CfnCustomPermissions",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCustomPermissions) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) AddResourceDependency(target awscdk.CfnResource, reason *string) {
	if err := c.validateAddResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addResourceDependency",
		[]interface{}{target, reason},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength) {
	if err := c.validateApplyCrossStackReferenceStrengthParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyCrossStackReferenceStrength",
		[]interface{}{strength},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) CfnPropertyName(cdkPropertyName *string) *string {
	if err := c.validateCfnPropertyNameParameters(cdkPropertyName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"cfnPropertyName",
		[]interface{}{cdkPropertyName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomPermissions) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomPermissions) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomPermissions) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomPermissions) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) RemoveResourceDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeResourceDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomPermissions) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomPermissions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomPermissions) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnCustomPermissions) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

