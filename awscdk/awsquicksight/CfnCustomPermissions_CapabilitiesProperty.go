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
//   	AdobeAnalyticsDataSource: jsii.String("adobeAnalyticsDataSource"),
//   	AmazonBedrockArsAction: jsii.String("amazonBedrockArsAction"),
//   	AmazonBedrockFsAction: jsii.String("amazonBedrockFsAction"),
//   	AmazonBedrockKrsAction: jsii.String("amazonBedrockKrsAction"),
//   	AmazonSThreeAction: jsii.String("amazonSThreeAction"),
//   	Analysis: jsii.String("analysis"),
//   	ApproveFlowShareRequests: jsii.String("approveFlowShareRequests"),
//   	Apps: jsii.String("apps"),
//   	AsanaAction: jsii.String("asanaAction"),
//   	AthenaDataSource: jsii.String("athenaDataSource"),
//   	AuroraDataSource: jsii.String("auroraDataSource"),
//   	Automate: jsii.String("automate"),
//   	BambooHrAction: jsii.String("bambooHrAction"),
//   	BedrockManagedKnowledgeBase: jsii.String("bedrockManagedKnowledgeBase"),
//   	BoxAgentAction: jsii.String("boxAgentAction"),
//   	BoxKnowledgeBase: jsii.String("boxKnowledgeBase"),
//   	BuildCalculatedFieldWithQ: jsii.String("buildCalculatedFieldWithQ"),
//   	CanvaAgentAction: jsii.String("canvaAgentAction"),
//   	ChatAgent: jsii.String("chatAgent"),
//   	ComprehendAction: jsii.String("comprehendAction"),
//   	ComprehendMedicalAction: jsii.String("comprehendMedicalAction"),
//   	ConfluenceAction: jsii.String("confluenceAction"),
//   	ConfluenceKnowledgeBase: jsii.String("confluenceKnowledgeBase"),
//   	CreateAdobeAnalyticsDataSource: jsii.String("createAdobeAnalyticsDataSource"),
//   	CreateAndUpdateAmazonBedrockArsAction: jsii.String("createAndUpdateAmazonBedrockArsAction"),
//   	CreateAndUpdateAmazonBedrockFsAction: jsii.String("createAndUpdateAmazonBedrockFsAction"),
//   	CreateAndUpdateAmazonBedrockKrsAction: jsii.String("createAndUpdateAmazonBedrockKrsAction"),
//   	CreateAndUpdateAmazonSThreeAction: jsii.String("createAndUpdateAmazonSThreeAction"),
//   	CreateAndUpdateApps: jsii.String("createAndUpdateApps"),
//   	CreateAndUpdateAsanaAction: jsii.String("createAndUpdateAsanaAction"),
//   	CreateAndUpdateBambooHrAction: jsii.String("createAndUpdateBambooHrAction"),
//   	CreateAndUpdateBedrockManagedKnowledgeBase: jsii.String("createAndUpdateBedrockManagedKnowledgeBase"),
//   	CreateAndUpdateBoxAgentAction: jsii.String("createAndUpdateBoxAgentAction"),
//   	CreateAndUpdateBoxKnowledgeBase: jsii.String("createAndUpdateBoxKnowledgeBase"),
//   	CreateAndUpdateCanvaAgentAction: jsii.String("createAndUpdateCanvaAgentAction"),
//   	CreateAndUpdateComprehendAction: jsii.String("createAndUpdateComprehendAction"),
//   	CreateAndUpdateComprehendMedicalAction: jsii.String("createAndUpdateComprehendMedicalAction"),
//   	CreateAndUpdateConfluenceAction: jsii.String("createAndUpdateConfluenceAction"),
//   	CreateAndUpdateConfluenceKnowledgeBase: jsii.String("createAndUpdateConfluenceKnowledgeBase"),
//   	CreateAndUpdateDashboardEmailReports: jsii.String("createAndUpdateDashboardEmailReports"),
//   	CreateAndUpdateDatasets: jsii.String("createAndUpdateDatasets"),
//   	CreateAndUpdateDataSources: jsii.String("createAndUpdateDataSources"),
//   	CreateAndUpdateFactSetAction: jsii.String("createAndUpdateFactSetAction"),
//   	CreateAndUpdateGenericHttpAction: jsii.String("createAndUpdateGenericHttpAction"),
//   	CreateAndUpdateGithubAction: jsii.String("createAndUpdateGithubAction"),
//   	CreateAndUpdateGoogleCalendarAction: jsii.String("createAndUpdateGoogleCalendarAction"),
//   	CreateAndUpdateGoogleDriveKnowledgeBase: jsii.String("createAndUpdateGoogleDriveKnowledgeBase"),
//   	CreateAndUpdateHubspotAction: jsii.String("createAndUpdateHubspotAction"),
//   	CreateAndUpdateHuggingFaceAction: jsii.String("createAndUpdateHuggingFaceAction"),
//   	CreateAndUpdateIdcKnowledgeBase: jsii.String("createAndUpdateIdcKnowledgeBase"),
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
//   	CreateAndUpdateOneDriveKnowledgeBase: jsii.String("createAndUpdateOneDriveKnowledgeBase"),
//   	CreateAndUpdateOpenApiAction: jsii.String("createAndUpdateOpenApiAction"),
//   	CreateAndUpdatePagerDutyAction: jsii.String("createAndUpdatePagerDutyAction"),
//   	CreateAndUpdateQBusinessKnowledgeBase: jsii.String("createAndUpdateQBusinessKnowledgeBase"),
//   	CreateAndUpdateS3KnowledgeBase: jsii.String("createAndUpdateS3KnowledgeBase"),
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
//   	CreateAndUpdateSharePointKnowledgeBase: jsii.String("createAndUpdateSharePointKnowledgeBase"),
//   	CreateAndUpdateSlackAction: jsii.String("createAndUpdateSlackAction"),
//   	CreateAndUpdateSmartsheetAction: jsii.String("createAndUpdateSmartsheetAction"),
//   	CreateAndUpdateTextractAction: jsii.String("createAndUpdateTextractAction"),
//   	CreateAndUpdateThemes: jsii.String("createAndUpdateThemes"),
//   	CreateAndUpdateThresholdAlerts: jsii.String("createAndUpdateThresholdAlerts"),
//   	CreateAndUpdateWebCrawlerKnowledgeBase: jsii.String("createAndUpdateWebCrawlerKnowledgeBase"),
//   	CreateAndUpdateZendeskAction: jsii.String("createAndUpdateZendeskAction"),
//   	CreateAthenaDataSource: jsii.String("createAthenaDataSource"),
//   	CreateAuroraDataSource: jsii.String("createAuroraDataSource"),
//   	CreateChatAgents: jsii.String("createChatAgents"),
//   	CreateDashboardExecutiveSummaryWithQ: jsii.String("createDashboardExecutiveSummaryWithQ"),
//   	CreateDatabricksDataSource: jsii.String("createDatabricksDataSource"),
//   	CreateDb2DataSource: jsii.String("createDb2DataSource"),
//   	CreateDenodoDataSource: jsii.String("createDenodoDataSource"),
//   	CreateDocumentDbDataSource: jsii.String("createDocumentDbDataSource"),
//   	CreateDremioDataSource: jsii.String("createDremioDataSource"),
//   	CreateDynamoDbDataSource: jsii.String("createDynamoDbDataSource"),
//   	CreateExasolDataSource: jsii.String("createExasolDataSource"),
//   	CreateFileDataSource: jsii.String("createFileDataSource"),
//   	CreateGitHubDataSource: jsii.String("createGitHubDataSource"),
//   	CreateGoogleAnalyticsDataSource: jsii.String("createGoogleAnalyticsDataSource"),
//   	CreateGoogleBigQueryDataSource: jsii.String("createGoogleBigQueryDataSource"),
//   	CreateGoogleSheetsDataSource: jsii.String("createGoogleSheetsDataSource"),
//   	CreateImpalaDataSource: jsii.String("createImpalaDataSource"),
//   	CreateJiraDataSource: jsii.String("createJiraDataSource"),
//   	CreateMariaDbDataSource: jsii.String("createMariaDbDataSource"),
//   	CreateMongoAtlasDataSource: jsii.String("createMongoAtlasDataSource"),
//   	CreateMongoDbDataSource: jsii.String("createMongoDbDataSource"),
//   	CreateMySqlDataSource: jsii.String("createMySqlDataSource"),
//   	CreateOpenSearchDataSource: jsii.String("createOpenSearchDataSource"),
//   	CreateOracleDataSource: jsii.String("createOracleDataSource"),
//   	CreatePayPalDataSource: jsii.String("createPayPalDataSource"),
//   	CreatePostgreSqlDataSource: jsii.String("createPostgreSqlDataSource"),
//   	CreatePrestoDataSource: jsii.String("createPrestoDataSource"),
//   	CreateRadiantDataSource: jsii.String("createRadiantDataSource"),
//   	CreateRdsDataSource: jsii.String("createRdsDataSource"),
//   	CreateRedshiftAutoDiscoveredDataSource: jsii.String("createRedshiftAutoDiscoveredDataSource"),
//   	CreateRedshiftManualDataSource: jsii.String("createRedshiftManualDataSource"),
//   	CreateS3AnalyticsDataSource: jsii.String("createS3AnalyticsDataSource"),
//   	CreateS3DataSource: jsii.String("createS3DataSource"),
//   	CreateS3TablesDataSource: jsii.String("createS3TablesDataSource"),
//   	CreateSalesforceDataSource: jsii.String("createSalesforceDataSource"),
//   	CreateSapHanaDataSource: jsii.String("createSapHanaDataSource"),
//   	CreateServiceNowDataSource: jsii.String("createServiceNowDataSource"),
//   	CreateSharedFolders: jsii.String("createSharedFolders"),
//   	CreateSnowflakeDataSource: jsii.String("createSnowflakeDataSource"),
//   	CreateSpaces: jsii.String("createSpaces"),
//   	CreateSparkDataSource: jsii.String("createSparkDataSource"),
//   	CreateSpiceDataset: jsii.String("createSpiceDataset"),
//   	CreateSqlServerDataSource: jsii.String("createSqlServerDataSource"),
//   	CreateSquareDataSource: jsii.String("createSquareDataSource"),
//   	CreateStarburstDataSource: jsii.String("createStarburstDataSource"),
//   	CreateTeradataDataSource: jsii.String("createTeradataDataSource"),
//   	CreateTimestreamDataSource: jsii.String("createTimestreamDataSource"),
//   	CreateTrinoDataSource: jsii.String("createTrinoDataSource"),
//   	CreateTwitterDataSource: jsii.String("createTwitterDataSource"),
//   	Dashboard: jsii.String("dashboard"),
//   	DatabricksDataSource: jsii.String("databricksDataSource"),
//   	Db2DataSource: jsii.String("db2DataSource"),
//   	DenodoDataSource: jsii.String("denodoDataSource"),
//   	DocumentDbDataSource: jsii.String("documentDbDataSource"),
//   	DremioDataSource: jsii.String("dremioDataSource"),
//   	DynamoDbDataSource: jsii.String("dynamoDbDataSource"),
//   	EditVisualWithQ: jsii.String("editVisualWithQ"),
//   	ExasolDataSource: jsii.String("exasolDataSource"),
//   	ExportToCsv: jsii.String("exportToCsv"),
//   	ExportToCsvInScheduledReports: jsii.String("exportToCsvInScheduledReports"),
//   	ExportToExcel: jsii.String("exportToExcel"),
//   	ExportToExcelInScheduledReports: jsii.String("exportToExcelInScheduledReports"),
//   	ExportToPdf: jsii.String("exportToPdf"),
//   	ExportToPdfInScheduledReports: jsii.String("exportToPdfInScheduledReports"),
//   	Extension: jsii.String("extension"),
//   	FactSetAction: jsii.String("factSetAction"),
//   	FileDataSource: jsii.String("fileDataSource"),
//   	Flow: jsii.String("flow"),
//   	GenericHttpAction: jsii.String("genericHttpAction"),
//   	GithubAction: jsii.String("githubAction"),
//   	GitHubDataSource: jsii.String("gitHubDataSource"),
//   	GoogleAnalyticsDataSource: jsii.String("googleAnalyticsDataSource"),
//   	GoogleBigQueryDataSource: jsii.String("googleBigQueryDataSource"),
//   	GoogleCalendarAction: jsii.String("googleCalendarAction"),
//   	GoogleDriveKnowledgeBase: jsii.String("googleDriveKnowledgeBase"),
//   	GoogleSheetsDataSource: jsii.String("googleSheetsDataSource"),
//   	HubspotAction: jsii.String("hubspotAction"),
//   	HuggingFaceAction: jsii.String("huggingFaceAction"),
//   	IdcKnowledgeBase: jsii.String("idcKnowledgeBase"),
//   	ImpalaDataSource: jsii.String("impalaDataSource"),
//   	InboundEmailTrigger: jsii.String("inboundEmailTrigger"),
//   	IncludeContentInScheduledReportsEmail: jsii.String("includeContentInScheduledReportsEmail"),
//   	IntercomAction: jsii.String("intercomAction"),
//   	InvokeAppsAiInference: jsii.String("invokeAppsAiInference"),
//   	JiraAction: jsii.String("jiraAction"),
//   	JiraDataSource: jsii.String("jiraDataSource"),
//   	KnowledgeBase: jsii.String("knowledgeBase"),
//   	LinearAction: jsii.String("linearAction"),
//   	ManageSharedFolders: jsii.String("manageSharedFolders"),
//   	MariaDbDataSource: jsii.String("mariaDbDataSource"),
//   	McpAction: jsii.String("mcpAction"),
//   	MondayAction: jsii.String("mondayAction"),
//   	MongoAtlasDataSource: jsii.String("mongoAtlasDataSource"),
//   	MongoDbDataSource: jsii.String("mongoDbDataSource"),
//   	MsExchangeAction: jsii.String("msExchangeAction"),
//   	MsTeamsAction: jsii.String("msTeamsAction"),
//   	MySqlDataSource: jsii.String("mySqlDataSource"),
//   	NewRelicAction: jsii.String("newRelicAction"),
//   	NotionAction: jsii.String("notionAction"),
//   	OneDriveAction: jsii.String("oneDriveAction"),
//   	OneDriveKnowledgeBase: jsii.String("oneDriveKnowledgeBase"),
//   	OpenApiAction: jsii.String("openApiAction"),
//   	OpenSearchDataSource: jsii.String("openSearchDataSource"),
//   	OracleDataSource: jsii.String("oracleDataSource"),
//   	PagerDutyAction: jsii.String("pagerDutyAction"),
//   	PayPalDataSource: jsii.String("payPalDataSource"),
//   	PerformFlowUiTask: jsii.String("performFlowUiTask"),
//   	PostgreSqlDataSource: jsii.String("postgreSqlDataSource"),
//   	PrestoDataSource: jsii.String("prestoDataSource"),
//   	PrintReports: jsii.String("printReports"),
//   	PublishWithoutApproval: jsii.String("publishWithoutApproval"),
//   	QBusinessKnowledgeBase: jsii.String("qBusinessKnowledgeBase"),
//   	QuickEventTrigger: jsii.String("quickEventTrigger"),
//   	RadiantDataSource: jsii.String("radiantDataSource"),
//   	RdsDataSource: jsii.String("rdsDataSource"),
//   	RedshiftAutoDiscoveredDataSource: jsii.String("redshiftAutoDiscoveredDataSource"),
//   	RedshiftManualDataSource: jsii.String("redshiftManualDataSource"),
//   	RenameSharedFolders: jsii.String("renameSharedFolders"),
//   	Research: jsii.String("research"),
//   	S3AnalyticsDataSource: jsii.String("s3AnalyticsDataSource"),
//   	S3DataSource: jsii.String("s3DataSource"),
//   	S3KnowledgeBase: jsii.String("s3KnowledgeBase"),
//   	S3TablesDataSource: jsii.String("s3TablesDataSource"),
//   	SalesforceAction: jsii.String("salesforceAction"),
//   	SalesforceDataSource: jsii.String("salesforceDataSource"),
//   	SandPGlobalEnergyAction: jsii.String("sandPGlobalEnergyAction"),
//   	SandPgmiAction: jsii.String("sandPgmiAction"),
//   	SapBillOfMaterialAction: jsii.String("sapBillOfMaterialAction"),
//   	SapBusinessPartnerAction: jsii.String("sapBusinessPartnerAction"),
//   	SapHanaDataSource: jsii.String("sapHanaDataSource"),
//   	SapMaterialStockAction: jsii.String("sapMaterialStockAction"),
//   	SapPhysicalInventoryAction: jsii.String("sapPhysicalInventoryAction"),
//   	SapProductMasterDataAction: jsii.String("sapProductMasterDataAction"),
//   	ScheduleTrigger: jsii.String("scheduleTrigger"),
//   	ServiceNowAction: jsii.String("serviceNowAction"),
//   	ServiceNowDataSource: jsii.String("serviceNowDataSource"),
//   	ShareAdobeAnalyticsDataSource: jsii.String("shareAdobeAnalyticsDataSource"),
//   	ShareAmazonBedrockArsAction: jsii.String("shareAmazonBedrockArsAction"),
//   	ShareAmazonBedrockFsAction: jsii.String("shareAmazonBedrockFsAction"),
//   	ShareAmazonBedrockKrsAction: jsii.String("shareAmazonBedrockKrsAction"),
//   	ShareAmazonSThreeAction: jsii.String("shareAmazonSThreeAction"),
//   	ShareAnalyses: jsii.String("shareAnalyses"),
//   	ShareApps: jsii.String("shareApps"),
//   	ShareAsanaAction: jsii.String("shareAsanaAction"),
//   	ShareAthenaDataSource: jsii.String("shareAthenaDataSource"),
//   	ShareAuroraDataSource: jsii.String("shareAuroraDataSource"),
//   	ShareBambooHrAction: jsii.String("shareBambooHrAction"),
//   	ShareBedrockManagedKnowledgeBase: jsii.String("shareBedrockManagedKnowledgeBase"),
//   	ShareBoxAgentAction: jsii.String("shareBoxAgentAction"),
//   	ShareBoxKnowledgeBase: jsii.String("shareBoxKnowledgeBase"),
//   	ShareCanvaAgentAction: jsii.String("shareCanvaAgentAction"),
//   	ShareChatAgents: jsii.String("shareChatAgents"),
//   	ShareComprehendAction: jsii.String("shareComprehendAction"),
//   	ShareComprehendMedicalAction: jsii.String("shareComprehendMedicalAction"),
//   	ShareConfluenceAction: jsii.String("shareConfluenceAction"),
//   	ShareConfluenceKnowledgeBase: jsii.String("shareConfluenceKnowledgeBase"),
//   	ShareDashboards: jsii.String("shareDashboards"),
//   	ShareDatabricksDataSource: jsii.String("shareDatabricksDataSource"),
//   	ShareDatasets: jsii.String("shareDatasets"),
//   	ShareDataSources: jsii.String("shareDataSources"),
//   	ShareDb2DataSource: jsii.String("shareDb2DataSource"),
//   	ShareDenodoDataSource: jsii.String("shareDenodoDataSource"),
//   	ShareDocumentDbDataSource: jsii.String("shareDocumentDbDataSource"),
//   	ShareDremioDataSource: jsii.String("shareDremioDataSource"),
//   	ShareDynamoDbDataSource: jsii.String("shareDynamoDbDataSource"),
//   	ShareExasolDataSource: jsii.String("shareExasolDataSource"),
//   	ShareFactSetAction: jsii.String("shareFactSetAction"),
//   	ShareFileDataSource: jsii.String("shareFileDataSource"),
//   	ShareGenericHttpAction: jsii.String("shareGenericHttpAction"),
//   	ShareGithubAction: jsii.String("shareGithubAction"),
//   	ShareGitHubDataSource: jsii.String("shareGitHubDataSource"),
//   	ShareGoogleAnalyticsDataSource: jsii.String("shareGoogleAnalyticsDataSource"),
//   	ShareGoogleBigQueryDataSource: jsii.String("shareGoogleBigQueryDataSource"),
//   	ShareGoogleCalendarAction: jsii.String("shareGoogleCalendarAction"),
//   	ShareGoogleDriveKnowledgeBase: jsii.String("shareGoogleDriveKnowledgeBase"),
//   	ShareGoogleSheetsDataSource: jsii.String("shareGoogleSheetsDataSource"),
//   	ShareHubspotAction: jsii.String("shareHubspotAction"),
//   	ShareHuggingFaceAction: jsii.String("shareHuggingFaceAction"),
//   	ShareIdcKnowledgeBase: jsii.String("shareIdcKnowledgeBase"),
//   	ShareImpalaDataSource: jsii.String("shareImpalaDataSource"),
//   	ShareIntercomAction: jsii.String("shareIntercomAction"),
//   	ShareJiraAction: jsii.String("shareJiraAction"),
//   	ShareJiraDataSource: jsii.String("shareJiraDataSource"),
//   	ShareKnowledgeBases: jsii.String("shareKnowledgeBases"),
//   	ShareLinearAction: jsii.String("shareLinearAction"),
//   	ShareMariaDbDataSource: jsii.String("shareMariaDbDataSource"),
//   	ShareMcpAction: jsii.String("shareMcpAction"),
//   	ShareMondayAction: jsii.String("shareMondayAction"),
//   	ShareMongoAtlasDataSource: jsii.String("shareMongoAtlasDataSource"),
//   	ShareMongoDbDataSource: jsii.String("shareMongoDbDataSource"),
//   	ShareMsExchangeAction: jsii.String("shareMsExchangeAction"),
//   	ShareMsTeamsAction: jsii.String("shareMsTeamsAction"),
//   	ShareMySqlDataSource: jsii.String("shareMySqlDataSource"),
//   	ShareNewRelicAction: jsii.String("shareNewRelicAction"),
//   	ShareNotionAction: jsii.String("shareNotionAction"),
//   	ShareOneDriveAction: jsii.String("shareOneDriveAction"),
//   	ShareOneDriveKnowledgeBase: jsii.String("shareOneDriveKnowledgeBase"),
//   	ShareOpenApiAction: jsii.String("shareOpenApiAction"),
//   	ShareOpenSearchDataSource: jsii.String("shareOpenSearchDataSource"),
//   	ShareOracleDataSource: jsii.String("shareOracleDataSource"),
//   	SharePagerDutyAction: jsii.String("sharePagerDutyAction"),
//   	SharePayPalDataSource: jsii.String("sharePayPalDataSource"),
//   	SharePointAction: jsii.String("sharePointAction"),
//   	SharePointKnowledgeBase: jsii.String("sharePointKnowledgeBase"),
//   	SharePostgreSqlDataSource: jsii.String("sharePostgreSqlDataSource"),
//   	SharePrestoDataSource: jsii.String("sharePrestoDataSource"),
//   	ShareQBusinessKnowledgeBase: jsii.String("shareQBusinessKnowledgeBase"),
//   	ShareRadiantDataSource: jsii.String("shareRadiantDataSource"),
//   	ShareRdsDataSource: jsii.String("shareRdsDataSource"),
//   	ShareRedshiftAutoDiscoveredDataSource: jsii.String("shareRedshiftAutoDiscoveredDataSource"),
//   	ShareRedshiftManualDataSource: jsii.String("shareRedshiftManualDataSource"),
//   	ShareS3AnalyticsDataSource: jsii.String("shareS3AnalyticsDataSource"),
//   	ShareS3DataSource: jsii.String("shareS3DataSource"),
//   	ShareS3KnowledgeBase: jsii.String("shareS3KnowledgeBase"),
//   	ShareS3TablesDataSource: jsii.String("shareS3TablesDataSource"),
//   	ShareSalesforceAction: jsii.String("shareSalesforceAction"),
//   	ShareSalesforceDataSource: jsii.String("shareSalesforceDataSource"),
//   	ShareSandPGlobalEnergyAction: jsii.String("shareSandPGlobalEnergyAction"),
//   	ShareSandPgmiAction: jsii.String("shareSandPgmiAction"),
//   	ShareSapBillOfMaterialAction: jsii.String("shareSapBillOfMaterialAction"),
//   	ShareSapBusinessPartnerAction: jsii.String("shareSapBusinessPartnerAction"),
//   	ShareSapHanaDataSource: jsii.String("shareSapHanaDataSource"),
//   	ShareSapMaterialStockAction: jsii.String("shareSapMaterialStockAction"),
//   	ShareSapPhysicalInventoryAction: jsii.String("shareSapPhysicalInventoryAction"),
//   	ShareSapProductMasterDataAction: jsii.String("shareSapProductMasterDataAction"),
//   	ShareServiceNowAction: jsii.String("shareServiceNowAction"),
//   	ShareServiceNowDataSource: jsii.String("shareServiceNowDataSource"),
//   	ShareSharePointAction: jsii.String("shareSharePointAction"),
//   	ShareSharePointKnowledgeBase: jsii.String("shareSharePointKnowledgeBase"),
//   	ShareSlackAction: jsii.String("shareSlackAction"),
//   	ShareSmartsheetAction: jsii.String("shareSmartsheetAction"),
//   	ShareSnowflakeDataSource: jsii.String("shareSnowflakeDataSource"),
//   	ShareSpaces: jsii.String("shareSpaces"),
//   	ShareSparkDataSource: jsii.String("shareSparkDataSource"),
//   	ShareSqlServerDataSource: jsii.String("shareSqlServerDataSource"),
//   	ShareSquareDataSource: jsii.String("shareSquareDataSource"),
//   	ShareStarburstDataSource: jsii.String("shareStarburstDataSource"),
//   	ShareTeradataDataSource: jsii.String("shareTeradataDataSource"),
//   	ShareTextractAction: jsii.String("shareTextractAction"),
//   	ShareTimestreamDataSource: jsii.String("shareTimestreamDataSource"),
//   	ShareTrinoDataSource: jsii.String("shareTrinoDataSource"),
//   	ShareTwitterDataSource: jsii.String("shareTwitterDataSource"),
//   	ShareWebCrawlerKnowledgeBase: jsii.String("shareWebCrawlerKnowledgeBase"),
//   	ShareZendeskAction: jsii.String("shareZendeskAction"),
//   	SlackAction: jsii.String("slackAction"),
//   	SmartsheetAction: jsii.String("smartsheetAction"),
//   	SnowflakeDataSource: jsii.String("snowflakeDataSource"),
//   	Space: jsii.String("space"),
//   	SparkDataSource: jsii.String("sparkDataSource"),
//   	SqlServerDataSource: jsii.String("sqlServerDataSource"),
//   	SquareDataSource: jsii.String("squareDataSource"),
//   	StarburstDataSource: jsii.String("starburstDataSource"),
//   	SubscribeDashboardEmailReports: jsii.String("subscribeDashboardEmailReports"),
//   	TeradataDataSource: jsii.String("teradataDataSource"),
//   	TextractAction: jsii.String("textractAction"),
//   	TimestreamDataSource: jsii.String("timestreamDataSource"),
//   	Topic: jsii.String("topic"),
//   	Trigger: jsii.String("trigger"),
//   	TrinoDataSource: jsii.String("trinoDataSource"),
//   	TwitterDataSource: jsii.String("twitterDataSource"),
//   	UpdateAdobeAnalyticsDataSource: jsii.String("updateAdobeAnalyticsDataSource"),
//   	UpdateAthenaDataSource: jsii.String("updateAthenaDataSource"),
//   	UpdateAuroraDataSource: jsii.String("updateAuroraDataSource"),
//   	UpdateDatabricksDataSource: jsii.String("updateDatabricksDataSource"),
//   	UpdateDb2DataSource: jsii.String("updateDb2DataSource"),
//   	UpdateDenodoDataSource: jsii.String("updateDenodoDataSource"),
//   	UpdateDocumentDbDataSource: jsii.String("updateDocumentDbDataSource"),
//   	UpdateDremioDataSource: jsii.String("updateDremioDataSource"),
//   	UpdateDynamoDbDataSource: jsii.String("updateDynamoDbDataSource"),
//   	UpdateExasolDataSource: jsii.String("updateExasolDataSource"),
//   	UpdateFileDataSource: jsii.String("updateFileDataSource"),
//   	UpdateGitHubDataSource: jsii.String("updateGitHubDataSource"),
//   	UpdateGoogleAnalyticsDataSource: jsii.String("updateGoogleAnalyticsDataSource"),
//   	UpdateGoogleBigQueryDataSource: jsii.String("updateGoogleBigQueryDataSource"),
//   	UpdateGoogleSheetsDataSource: jsii.String("updateGoogleSheetsDataSource"),
//   	UpdateImpalaDataSource: jsii.String("updateImpalaDataSource"),
//   	UpdateJiraDataSource: jsii.String("updateJiraDataSource"),
//   	UpdateMariaDbDataSource: jsii.String("updateMariaDbDataSource"),
//   	UpdateMongoAtlasDataSource: jsii.String("updateMongoAtlasDataSource"),
//   	UpdateMongoDbDataSource: jsii.String("updateMongoDbDataSource"),
//   	UpdateMySqlDataSource: jsii.String("updateMySqlDataSource"),
//   	UpdateOpenSearchDataSource: jsii.String("updateOpenSearchDataSource"),
//   	UpdateOracleDataSource: jsii.String("updateOracleDataSource"),
//   	UpdatePayPalDataSource: jsii.String("updatePayPalDataSource"),
//   	UpdatePostgreSqlDataSource: jsii.String("updatePostgreSqlDataSource"),
//   	UpdatePrestoDataSource: jsii.String("updatePrestoDataSource"),
//   	UpdateRadiantDataSource: jsii.String("updateRadiantDataSource"),
//   	UpdateRdsDataSource: jsii.String("updateRdsDataSource"),
//   	UpdateRedshiftAutoDiscoveredDataSource: jsii.String("updateRedshiftAutoDiscoveredDataSource"),
//   	UpdateRedshiftManualDataSource: jsii.String("updateRedshiftManualDataSource"),
//   	UpdateS3AnalyticsDataSource: jsii.String("updateS3AnalyticsDataSource"),
//   	UpdateS3DataSource: jsii.String("updateS3DataSource"),
//   	UpdateS3TablesDataSource: jsii.String("updateS3TablesDataSource"),
//   	UpdateSalesforceDataSource: jsii.String("updateSalesforceDataSource"),
//   	UpdateSapHanaDataSource: jsii.String("updateSapHanaDataSource"),
//   	UpdateServiceNowDataSource: jsii.String("updateServiceNowDataSource"),
//   	UpdateSnowflakeDataSource: jsii.String("updateSnowflakeDataSource"),
//   	UpdateSparkDataSource: jsii.String("updateSparkDataSource"),
//   	UpdateSqlServerDataSource: jsii.String("updateSqlServerDataSource"),
//   	UpdateSquareDataSource: jsii.String("updateSquareDataSource"),
//   	UpdateStarburstDataSource: jsii.String("updateStarburstDataSource"),
//   	UpdateTeradataDataSource: jsii.String("updateTeradataDataSource"),
//   	UpdateTimestreamDataSource: jsii.String("updateTimestreamDataSource"),
//   	UpdateTrinoDataSource: jsii.String("updateTrinoDataSource"),
//   	UpdateTwitterDataSource: jsii.String("updateTwitterDataSource"),
//   	UseAgentWebSearch: jsii.String("useAgentWebSearch"),
//   	UseAmazonBedrockArsAction: jsii.String("useAmazonBedrockArsAction"),
//   	UseAmazonBedrockFsAction: jsii.String("useAmazonBedrockFsAction"),
//   	UseAmazonBedrockKrsAction: jsii.String("useAmazonBedrockKrsAction"),
//   	UseAmazonSThreeAction: jsii.String("useAmazonSThreeAction"),
//   	UseAsanaAction: jsii.String("useAsanaAction"),
//   	UseBambooHrAction: jsii.String("useBambooHrAction"),
//   	UseBedrockManagedKnowledgeBase: jsii.String("useBedrockManagedKnowledgeBase"),
//   	UseBedrockModels: jsii.String("useBedrockModels"),
//   	UseBoxAgentAction: jsii.String("useBoxAgentAction"),
//   	UseBoxKnowledgeBase: jsii.String("useBoxKnowledgeBase"),
//   	UseBrowserExtension: jsii.String("useBrowserExtension"),
//   	UseCanvaAgentAction: jsii.String("useCanvaAgentAction"),
//   	UseComprehendAction: jsii.String("useComprehendAction"),
//   	UseComprehendMedicalAction: jsii.String("useComprehendMedicalAction"),
//   	UseConfluenceAction: jsii.String("useConfluenceAction"),
//   	UseConfluenceKnowledgeBase: jsii.String("useConfluenceKnowledgeBase"),
//   	UseExcelAddInExtension: jsii.String("useExcelAddInExtension"),
//   	UseFactSetAction: jsii.String("useFactSetAction"),
//   	UseGenericHttpAction: jsii.String("useGenericHttpAction"),
//   	UseGithubAction: jsii.String("useGithubAction"),
//   	UseGoogleCalendarAction: jsii.String("useGoogleCalendarAction"),
//   	UseGoogleDriveKnowledgeBase: jsii.String("useGoogleDriveKnowledgeBase"),
//   	UseHubspotAction: jsii.String("useHubspotAction"),
//   	UseHuggingFaceAction: jsii.String("useHuggingFaceAction"),
//   	UseIdcKnowledgeBase: jsii.String("useIdcKnowledgeBase"),
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
//   	UseOneDriveKnowledgeBase: jsii.String("useOneDriveKnowledgeBase"),
//   	UseOpenApiAction: jsii.String("useOpenApiAction"),
//   	UseOutlookAddInExtension: jsii.String("useOutlookAddInExtension"),
//   	UsePagerDutyAction: jsii.String("usePagerDutyAction"),
//   	UsePowerpointAddInExtension: jsii.String("usePowerpointAddInExtension"),
//   	UseQBusinessKnowledgeBase: jsii.String("useQBusinessKnowledgeBase"),
//   	UseS3KnowledgeBase: jsii.String("useS3KnowledgeBase"),
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
//   	UseSharePointKnowledgeBase: jsii.String("useSharePointKnowledgeBase"),
//   	UseSlackAction: jsii.String("useSlackAction"),
//   	UseSmartsheetAction: jsii.String("useSmartsheetAction"),
//   	UseTextractAction: jsii.String("useTextractAction"),
//   	UseWebCrawlerKnowledgeBase: jsii.String("useWebCrawlerKnowledgeBase"),
//   	UseWordAddInExtension: jsii.String("useWordAddInExtension"),
//   	UseZendeskAction: jsii.String("useZendeskAction"),
//   	ViewAccountSpiceCapacity: jsii.String("viewAccountSpiceCapacity"),
//   	WebCrawlerKnowledgeBase: jsii.String("webCrawlerKnowledgeBase"),
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-adobeanalyticsdatasource
	//
	AdobeAnalyticsDataSource *string `field:"optional" json:"adobeAnalyticsDataSource" yaml:"adobeAnalyticsDataSource"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-athenadatasource
	//
	AthenaDataSource *string `field:"optional" json:"athenaDataSource" yaml:"athenaDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-auroradatasource
	//
	AuroraDataSource *string `field:"optional" json:"auroraDataSource" yaml:"auroraDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-automate
	//
	Automate *string `field:"optional" json:"automate" yaml:"automate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-bamboohraction
	//
	BambooHrAction *string `field:"optional" json:"bambooHrAction" yaml:"bambooHrAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-bedrockmanagedknowledgebase
	//
	BedrockManagedKnowledgeBase *string `field:"optional" json:"bedrockManagedKnowledgeBase" yaml:"bedrockManagedKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-boxagentaction
	//
	BoxAgentAction *string `field:"optional" json:"boxAgentAction" yaml:"boxAgentAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-boxknowledgebase
	//
	BoxKnowledgeBase *string `field:"optional" json:"boxKnowledgeBase" yaml:"boxKnowledgeBase"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-confluenceknowledgebase
	//
	ConfluenceKnowledgeBase *string `field:"optional" json:"confluenceKnowledgeBase" yaml:"confluenceKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createadobeanalyticsdatasource
	//
	CreateAdobeAnalyticsDataSource *string `field:"optional" json:"createAdobeAnalyticsDataSource" yaml:"createAdobeAnalyticsDataSource"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatebedrockmanagedknowledgebase
	//
	CreateAndUpdateBedrockManagedKnowledgeBase *string `field:"optional" json:"createAndUpdateBedrockManagedKnowledgeBase" yaml:"createAndUpdateBedrockManagedKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateboxagentaction
	//
	CreateAndUpdateBoxAgentAction *string `field:"optional" json:"createAndUpdateBoxAgentAction" yaml:"createAndUpdateBoxAgentAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateboxknowledgebase
	//
	CreateAndUpdateBoxKnowledgeBase *string `field:"optional" json:"createAndUpdateBoxKnowledgeBase" yaml:"createAndUpdateBoxKnowledgeBase"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateconfluenceknowledgebase
	//
	CreateAndUpdateConfluenceKnowledgeBase *string `field:"optional" json:"createAndUpdateConfluenceKnowledgeBase" yaml:"createAndUpdateConfluenceKnowledgeBase"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdategoogledriveknowledgebase
	//
	CreateAndUpdateGoogleDriveKnowledgeBase *string `field:"optional" json:"createAndUpdateGoogleDriveKnowledgeBase" yaml:"createAndUpdateGoogleDriveKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatehubspotaction
	//
	CreateAndUpdateHubspotAction *string `field:"optional" json:"createAndUpdateHubspotAction" yaml:"createAndUpdateHubspotAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatehuggingfaceaction
	//
	CreateAndUpdateHuggingFaceAction *string `field:"optional" json:"createAndUpdateHuggingFaceAction" yaml:"createAndUpdateHuggingFaceAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateidcknowledgebase
	//
	CreateAndUpdateIdcKnowledgeBase *string `field:"optional" json:"createAndUpdateIdcKnowledgeBase" yaml:"createAndUpdateIdcKnowledgeBase"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateonedriveknowledgebase
	//
	CreateAndUpdateOneDriveKnowledgeBase *string `field:"optional" json:"createAndUpdateOneDriveKnowledgeBase" yaml:"createAndUpdateOneDriveKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateopenapiaction
	//
	CreateAndUpdateOpenApiAction *string `field:"optional" json:"createAndUpdateOpenApiAction" yaml:"createAndUpdateOpenApiAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatepagerdutyaction
	//
	CreateAndUpdatePagerDutyAction *string `field:"optional" json:"createAndUpdatePagerDutyAction" yaml:"createAndUpdatePagerDutyAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdateqbusinessknowledgebase
	//
	CreateAndUpdateQBusinessKnowledgeBase *string `field:"optional" json:"createAndUpdateQBusinessKnowledgeBase" yaml:"createAndUpdateQBusinessKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdates3knowledgebase
	//
	CreateAndUpdateS3KnowledgeBase *string `field:"optional" json:"createAndUpdateS3KnowledgeBase" yaml:"createAndUpdateS3KnowledgeBase"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatesharepointknowledgebase
	//
	CreateAndUpdateSharePointKnowledgeBase *string `field:"optional" json:"createAndUpdateSharePointKnowledgeBase" yaml:"createAndUpdateSharePointKnowledgeBase"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatewebcrawlerknowledgebase
	//
	CreateAndUpdateWebCrawlerKnowledgeBase *string `field:"optional" json:"createAndUpdateWebCrawlerKnowledgeBase" yaml:"createAndUpdateWebCrawlerKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatezendeskaction
	//
	CreateAndUpdateZendeskAction *string `field:"optional" json:"createAndUpdateZendeskAction" yaml:"createAndUpdateZendeskAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createathenadatasource
	//
	CreateAthenaDataSource *string `field:"optional" json:"createAthenaDataSource" yaml:"createAthenaDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createauroradatasource
	//
	CreateAuroraDataSource *string `field:"optional" json:"createAuroraDataSource" yaml:"createAuroraDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createchatagents
	//
	CreateChatAgents *string `field:"optional" json:"createChatAgents" yaml:"createChatAgents"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createdashboardexecutivesummarywithq
	//
	CreateDashboardExecutiveSummaryWithQ *string `field:"optional" json:"createDashboardExecutiveSummaryWithQ" yaml:"createDashboardExecutiveSummaryWithQ"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createdatabricksdatasource
	//
	CreateDatabricksDataSource *string `field:"optional" json:"createDatabricksDataSource" yaml:"createDatabricksDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createdb2datasource
	//
	CreateDb2DataSource *string `field:"optional" json:"createDb2DataSource" yaml:"createDb2DataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createdenododatasource
	//
	CreateDenodoDataSource *string `field:"optional" json:"createDenodoDataSource" yaml:"createDenodoDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createdocumentdbdatasource
	//
	CreateDocumentDbDataSource *string `field:"optional" json:"createDocumentDbDataSource" yaml:"createDocumentDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createdremiodatasource
	//
	CreateDremioDataSource *string `field:"optional" json:"createDremioDataSource" yaml:"createDremioDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createdynamodbdatasource
	//
	CreateDynamoDbDataSource *string `field:"optional" json:"createDynamoDbDataSource" yaml:"createDynamoDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createexasoldatasource
	//
	CreateExasolDataSource *string `field:"optional" json:"createExasolDataSource" yaml:"createExasolDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createfiledatasource
	//
	CreateFileDataSource *string `field:"optional" json:"createFileDataSource" yaml:"createFileDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-creategithubdatasource
	//
	CreateGitHubDataSource *string `field:"optional" json:"createGitHubDataSource" yaml:"createGitHubDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-creategoogleanalyticsdatasource
	//
	CreateGoogleAnalyticsDataSource *string `field:"optional" json:"createGoogleAnalyticsDataSource" yaml:"createGoogleAnalyticsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-creategooglebigquerydatasource
	//
	CreateGoogleBigQueryDataSource *string `field:"optional" json:"createGoogleBigQueryDataSource" yaml:"createGoogleBigQueryDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-creategooglesheetsdatasource
	//
	CreateGoogleSheetsDataSource *string `field:"optional" json:"createGoogleSheetsDataSource" yaml:"createGoogleSheetsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createimpaladatasource
	//
	CreateImpalaDataSource *string `field:"optional" json:"createImpalaDataSource" yaml:"createImpalaDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createjiradatasource
	//
	CreateJiraDataSource *string `field:"optional" json:"createJiraDataSource" yaml:"createJiraDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createmariadbdatasource
	//
	CreateMariaDbDataSource *string `field:"optional" json:"createMariaDbDataSource" yaml:"createMariaDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createmongoatlasdatasource
	//
	CreateMongoAtlasDataSource *string `field:"optional" json:"createMongoAtlasDataSource" yaml:"createMongoAtlasDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createmongodbdatasource
	//
	CreateMongoDbDataSource *string `field:"optional" json:"createMongoDbDataSource" yaml:"createMongoDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createmysqldatasource
	//
	CreateMySqlDataSource *string `field:"optional" json:"createMySqlDataSource" yaml:"createMySqlDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createopensearchdatasource
	//
	CreateOpenSearchDataSource *string `field:"optional" json:"createOpenSearchDataSource" yaml:"createOpenSearchDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createoracledatasource
	//
	CreateOracleDataSource *string `field:"optional" json:"createOracleDataSource" yaml:"createOracleDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createpaypaldatasource
	//
	CreatePayPalDataSource *string `field:"optional" json:"createPayPalDataSource" yaml:"createPayPalDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createpostgresqldatasource
	//
	CreatePostgreSqlDataSource *string `field:"optional" json:"createPostgreSqlDataSource" yaml:"createPostgreSqlDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createprestodatasource
	//
	CreatePrestoDataSource *string `field:"optional" json:"createPrestoDataSource" yaml:"createPrestoDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createradiantdatasource
	//
	CreateRadiantDataSource *string `field:"optional" json:"createRadiantDataSource" yaml:"createRadiantDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createrdsdatasource
	//
	CreateRdsDataSource *string `field:"optional" json:"createRdsDataSource" yaml:"createRdsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createredshiftautodiscovereddatasource
	//
	CreateRedshiftAutoDiscoveredDataSource *string `field:"optional" json:"createRedshiftAutoDiscoveredDataSource" yaml:"createRedshiftAutoDiscoveredDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createredshiftmanualdatasource
	//
	CreateRedshiftManualDataSource *string `field:"optional" json:"createRedshiftManualDataSource" yaml:"createRedshiftManualDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-creates3analyticsdatasource
	//
	CreateS3AnalyticsDataSource *string `field:"optional" json:"createS3AnalyticsDataSource" yaml:"createS3AnalyticsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-creates3datasource
	//
	CreateS3DataSource *string `field:"optional" json:"createS3DataSource" yaml:"createS3DataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-creates3tablesdatasource
	//
	CreateS3TablesDataSource *string `field:"optional" json:"createS3TablesDataSource" yaml:"createS3TablesDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createsalesforcedatasource
	//
	CreateSalesforceDataSource *string `field:"optional" json:"createSalesforceDataSource" yaml:"createSalesforceDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createsaphanadatasource
	//
	CreateSapHanaDataSource *string `field:"optional" json:"createSapHanaDataSource" yaml:"createSapHanaDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createservicenowdatasource
	//
	CreateServiceNowDataSource *string `field:"optional" json:"createServiceNowDataSource" yaml:"createServiceNowDataSource"`
	// The ability to create shared folders.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createsharedfolders
	//
	CreateSharedFolders *string `field:"optional" json:"createSharedFolders" yaml:"createSharedFolders"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createsnowflakedatasource
	//
	CreateSnowflakeDataSource *string `field:"optional" json:"createSnowflakeDataSource" yaml:"createSnowflakeDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createspaces
	//
	CreateSpaces *string `field:"optional" json:"createSpaces" yaml:"createSpaces"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createsparkdatasource
	//
	CreateSparkDataSource *string `field:"optional" json:"createSparkDataSource" yaml:"createSparkDataSource"`
	// The ability to create a SPICE dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createspicedataset
	//
	CreateSpiceDataset *string `field:"optional" json:"createSpiceDataset" yaml:"createSpiceDataset"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createsqlserverdatasource
	//
	CreateSqlServerDataSource *string `field:"optional" json:"createSqlServerDataSource" yaml:"createSqlServerDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createsquaredatasource
	//
	CreateSquareDataSource *string `field:"optional" json:"createSquareDataSource" yaml:"createSquareDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createstarburstdatasource
	//
	CreateStarburstDataSource *string `field:"optional" json:"createStarburstDataSource" yaml:"createStarburstDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createteradatadatasource
	//
	CreateTeradataDataSource *string `field:"optional" json:"createTeradataDataSource" yaml:"createTeradataDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createtimestreamdatasource
	//
	CreateTimestreamDataSource *string `field:"optional" json:"createTimestreamDataSource" yaml:"createTimestreamDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createtrinodatasource
	//
	CreateTrinoDataSource *string `field:"optional" json:"createTrinoDataSource" yaml:"createTrinoDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createtwitterdatasource
	//
	CreateTwitterDataSource *string `field:"optional" json:"createTwitterDataSource" yaml:"createTwitterDataSource"`
	// The ability to perform dashboard-related actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-dashboard
	//
	Dashboard *string `field:"optional" json:"dashboard" yaml:"dashboard"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-databricksdatasource
	//
	DatabricksDataSource *string `field:"optional" json:"databricksDataSource" yaml:"databricksDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-db2datasource
	//
	Db2DataSource *string `field:"optional" json:"db2DataSource" yaml:"db2DataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-denododatasource
	//
	DenodoDataSource *string `field:"optional" json:"denodoDataSource" yaml:"denodoDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-documentdbdatasource
	//
	DocumentDbDataSource *string `field:"optional" json:"documentDbDataSource" yaml:"documentDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-dremiodatasource
	//
	DremioDataSource *string `field:"optional" json:"dremioDataSource" yaml:"dremioDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-dynamodbdatasource
	//
	DynamoDbDataSource *string `field:"optional" json:"dynamoDbDataSource" yaml:"dynamoDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-editvisualwithq
	//
	EditVisualWithQ *string `field:"optional" json:"editVisualWithQ" yaml:"editVisualWithQ"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-exasoldatasource
	//
	ExasolDataSource *string `field:"optional" json:"exasolDataSource" yaml:"exasolDataSource"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-filedatasource
	//
	FileDataSource *string `field:"optional" json:"fileDataSource" yaml:"fileDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-flow
	//
	Flow *string `field:"optional" json:"flow" yaml:"flow"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-generichttpaction
	//
	GenericHttpAction *string `field:"optional" json:"genericHttpAction" yaml:"genericHttpAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-githubaction
	//
	GithubAction *string `field:"optional" json:"githubAction" yaml:"githubAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-githubdatasource
	//
	GitHubDataSource *string `field:"optional" json:"gitHubDataSource" yaml:"gitHubDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-googleanalyticsdatasource
	//
	GoogleAnalyticsDataSource *string `field:"optional" json:"googleAnalyticsDataSource" yaml:"googleAnalyticsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-googlebigquerydatasource
	//
	GoogleBigQueryDataSource *string `field:"optional" json:"googleBigQueryDataSource" yaml:"googleBigQueryDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-googlecalendaraction
	//
	GoogleCalendarAction *string `field:"optional" json:"googleCalendarAction" yaml:"googleCalendarAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-googledriveknowledgebase
	//
	GoogleDriveKnowledgeBase *string `field:"optional" json:"googleDriveKnowledgeBase" yaml:"googleDriveKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-googlesheetsdatasource
	//
	GoogleSheetsDataSource *string `field:"optional" json:"googleSheetsDataSource" yaml:"googleSheetsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-hubspotaction
	//
	HubspotAction *string `field:"optional" json:"hubspotAction" yaml:"hubspotAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-huggingfaceaction
	//
	HuggingFaceAction *string `field:"optional" json:"huggingFaceAction" yaml:"huggingFaceAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-idcknowledgebase
	//
	IdcKnowledgeBase *string `field:"optional" json:"idcKnowledgeBase" yaml:"idcKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-impaladatasource
	//
	ImpalaDataSource *string `field:"optional" json:"impalaDataSource" yaml:"impalaDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-inboundemailtrigger
	//
	InboundEmailTrigger *string `field:"optional" json:"inboundEmailTrigger" yaml:"inboundEmailTrigger"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-jiradatasource
	//
	JiraDataSource *string `field:"optional" json:"jiraDataSource" yaml:"jiraDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-knowledgebase
	//
	KnowledgeBase *string `field:"optional" json:"knowledgeBase" yaml:"knowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-linearaction
	//
	LinearAction *string `field:"optional" json:"linearAction" yaml:"linearAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-managesharedfolders
	//
	ManageSharedFolders *string `field:"optional" json:"manageSharedFolders" yaml:"manageSharedFolders"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-mariadbdatasource
	//
	MariaDbDataSource *string `field:"optional" json:"mariaDbDataSource" yaml:"mariaDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-mcpaction
	//
	McpAction *string `field:"optional" json:"mcpAction" yaml:"mcpAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-mondayaction
	//
	MondayAction *string `field:"optional" json:"mondayAction" yaml:"mondayAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-mongoatlasdatasource
	//
	MongoAtlasDataSource *string `field:"optional" json:"mongoAtlasDataSource" yaml:"mongoAtlasDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-mongodbdatasource
	//
	MongoDbDataSource *string `field:"optional" json:"mongoDbDataSource" yaml:"mongoDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-msexchangeaction
	//
	MsExchangeAction *string `field:"optional" json:"msExchangeAction" yaml:"msExchangeAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-msteamsaction
	//
	MsTeamsAction *string `field:"optional" json:"msTeamsAction" yaml:"msTeamsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-mysqldatasource
	//
	MySqlDataSource *string `field:"optional" json:"mySqlDataSource" yaml:"mySqlDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-newrelicaction
	//
	NewRelicAction *string `field:"optional" json:"newRelicAction" yaml:"newRelicAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-notionaction
	//
	NotionAction *string `field:"optional" json:"notionAction" yaml:"notionAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-onedriveaction
	//
	OneDriveAction *string `field:"optional" json:"oneDriveAction" yaml:"oneDriveAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-onedriveknowledgebase
	//
	OneDriveKnowledgeBase *string `field:"optional" json:"oneDriveKnowledgeBase" yaml:"oneDriveKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-openapiaction
	//
	OpenApiAction *string `field:"optional" json:"openApiAction" yaml:"openApiAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-opensearchdatasource
	//
	OpenSearchDataSource *string `field:"optional" json:"openSearchDataSource" yaml:"openSearchDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-oracledatasource
	//
	OracleDataSource *string `field:"optional" json:"oracleDataSource" yaml:"oracleDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-pagerdutyaction
	//
	PagerDutyAction *string `field:"optional" json:"pagerDutyAction" yaml:"pagerDutyAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-paypaldatasource
	//
	PayPalDataSource *string `field:"optional" json:"payPalDataSource" yaml:"payPalDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-performflowuitask
	//
	PerformFlowUiTask *string `field:"optional" json:"performFlowUiTask" yaml:"performFlowUiTask"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-postgresqldatasource
	//
	PostgreSqlDataSource *string `field:"optional" json:"postgreSqlDataSource" yaml:"postgreSqlDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-prestodatasource
	//
	PrestoDataSource *string `field:"optional" json:"prestoDataSource" yaml:"prestoDataSource"`
	// The ability to print reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-printreports
	//
	PrintReports *string `field:"optional" json:"printReports" yaml:"printReports"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-publishwithoutapproval
	//
	PublishWithoutApproval *string `field:"optional" json:"publishWithoutApproval" yaml:"publishWithoutApproval"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-qbusinessknowledgebase
	//
	QBusinessKnowledgeBase *string `field:"optional" json:"qBusinessKnowledgeBase" yaml:"qBusinessKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-quickeventtrigger
	//
	QuickEventTrigger *string `field:"optional" json:"quickEventTrigger" yaml:"quickEventTrigger"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-radiantdatasource
	//
	RadiantDataSource *string `field:"optional" json:"radiantDataSource" yaml:"radiantDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-rdsdatasource
	//
	RdsDataSource *string `field:"optional" json:"rdsDataSource" yaml:"rdsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-redshiftautodiscovereddatasource
	//
	RedshiftAutoDiscoveredDataSource *string `field:"optional" json:"redshiftAutoDiscoveredDataSource" yaml:"redshiftAutoDiscoveredDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-redshiftmanualdatasource
	//
	RedshiftManualDataSource *string `field:"optional" json:"redshiftManualDataSource" yaml:"redshiftManualDataSource"`
	// The ability to rename shared folders.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-renamesharedfolders
	//
	RenameSharedFolders *string `field:"optional" json:"renameSharedFolders" yaml:"renameSharedFolders"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-research
	//
	Research *string `field:"optional" json:"research" yaml:"research"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-s3analyticsdatasource
	//
	S3AnalyticsDataSource *string `field:"optional" json:"s3AnalyticsDataSource" yaml:"s3AnalyticsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-s3datasource
	//
	S3DataSource *string `field:"optional" json:"s3DataSource" yaml:"s3DataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-s3knowledgebase
	//
	S3KnowledgeBase *string `field:"optional" json:"s3KnowledgeBase" yaml:"s3KnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-s3tablesdatasource
	//
	S3TablesDataSource *string `field:"optional" json:"s3TablesDataSource" yaml:"s3TablesDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-salesforceaction
	//
	SalesforceAction *string `field:"optional" json:"salesforceAction" yaml:"salesforceAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-salesforcedatasource
	//
	SalesforceDataSource *string `field:"optional" json:"salesforceDataSource" yaml:"salesforceDataSource"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-saphanadatasource
	//
	SapHanaDataSource *string `field:"optional" json:"sapHanaDataSource" yaml:"sapHanaDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sapmaterialstockaction
	//
	SapMaterialStockAction *string `field:"optional" json:"sapMaterialStockAction" yaml:"sapMaterialStockAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sapphysicalinventoryaction
	//
	SapPhysicalInventoryAction *string `field:"optional" json:"sapPhysicalInventoryAction" yaml:"sapPhysicalInventoryAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sapproductmasterdataaction
	//
	SapProductMasterDataAction *string `field:"optional" json:"sapProductMasterDataAction" yaml:"sapProductMasterDataAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-scheduletrigger
	//
	ScheduleTrigger *string `field:"optional" json:"scheduleTrigger" yaml:"scheduleTrigger"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-servicenowaction
	//
	ServiceNowAction *string `field:"optional" json:"serviceNowAction" yaml:"serviceNowAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-servicenowdatasource
	//
	ServiceNowDataSource *string `field:"optional" json:"serviceNowDataSource" yaml:"serviceNowDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareadobeanalyticsdatasource
	//
	ShareAdobeAnalyticsDataSource *string `field:"optional" json:"shareAdobeAnalyticsDataSource" yaml:"shareAdobeAnalyticsDataSource"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareathenadatasource
	//
	ShareAthenaDataSource *string `field:"optional" json:"shareAthenaDataSource" yaml:"shareAthenaDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareauroradatasource
	//
	ShareAuroraDataSource *string `field:"optional" json:"shareAuroraDataSource" yaml:"shareAuroraDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharebamboohraction
	//
	ShareBambooHrAction *string `field:"optional" json:"shareBambooHrAction" yaml:"shareBambooHrAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharebedrockmanagedknowledgebase
	//
	ShareBedrockManagedKnowledgeBase *string `field:"optional" json:"shareBedrockManagedKnowledgeBase" yaml:"shareBedrockManagedKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareboxagentaction
	//
	ShareBoxAgentAction *string `field:"optional" json:"shareBoxAgentAction" yaml:"shareBoxAgentAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareboxknowledgebase
	//
	ShareBoxKnowledgeBase *string `field:"optional" json:"shareBoxKnowledgeBase" yaml:"shareBoxKnowledgeBase"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareconfluenceknowledgebase
	//
	ShareConfluenceKnowledgeBase *string `field:"optional" json:"shareConfluenceKnowledgeBase" yaml:"shareConfluenceKnowledgeBase"`
	// The ability to share dashboards.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharedashboards
	//
	ShareDashboards *string `field:"optional" json:"shareDashboards" yaml:"shareDashboards"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharedatabricksdatasource
	//
	ShareDatabricksDataSource *string `field:"optional" json:"shareDatabricksDataSource" yaml:"shareDatabricksDataSource"`
	// The ability to share datasets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharedatasets
	//
	ShareDatasets *string `field:"optional" json:"shareDatasets" yaml:"shareDatasets"`
	// The ability to share data sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharedatasources
	//
	ShareDataSources *string `field:"optional" json:"shareDataSources" yaml:"shareDataSources"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharedb2datasource
	//
	ShareDb2DataSource *string `field:"optional" json:"shareDb2DataSource" yaml:"shareDb2DataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharedenododatasource
	//
	ShareDenodoDataSource *string `field:"optional" json:"shareDenodoDataSource" yaml:"shareDenodoDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharedocumentdbdatasource
	//
	ShareDocumentDbDataSource *string `field:"optional" json:"shareDocumentDbDataSource" yaml:"shareDocumentDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharedremiodatasource
	//
	ShareDremioDataSource *string `field:"optional" json:"shareDremioDataSource" yaml:"shareDremioDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharedynamodbdatasource
	//
	ShareDynamoDbDataSource *string `field:"optional" json:"shareDynamoDbDataSource" yaml:"shareDynamoDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareexasoldatasource
	//
	ShareExasolDataSource *string `field:"optional" json:"shareExasolDataSource" yaml:"shareExasolDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharefactsetaction
	//
	ShareFactSetAction *string `field:"optional" json:"shareFactSetAction" yaml:"shareFactSetAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharefiledatasource
	//
	ShareFileDataSource *string `field:"optional" json:"shareFileDataSource" yaml:"shareFileDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharegenerichttpaction
	//
	ShareGenericHttpAction *string `field:"optional" json:"shareGenericHttpAction" yaml:"shareGenericHttpAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharegithubaction
	//
	ShareGithubAction *string `field:"optional" json:"shareGithubAction" yaml:"shareGithubAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharegithubdatasource
	//
	ShareGitHubDataSource *string `field:"optional" json:"shareGitHubDataSource" yaml:"shareGitHubDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharegoogleanalyticsdatasource
	//
	ShareGoogleAnalyticsDataSource *string `field:"optional" json:"shareGoogleAnalyticsDataSource" yaml:"shareGoogleAnalyticsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharegooglebigquerydatasource
	//
	ShareGoogleBigQueryDataSource *string `field:"optional" json:"shareGoogleBigQueryDataSource" yaml:"shareGoogleBigQueryDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharegooglecalendaraction
	//
	ShareGoogleCalendarAction *string `field:"optional" json:"shareGoogleCalendarAction" yaml:"shareGoogleCalendarAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharegoogledriveknowledgebase
	//
	ShareGoogleDriveKnowledgeBase *string `field:"optional" json:"shareGoogleDriveKnowledgeBase" yaml:"shareGoogleDriveKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharegooglesheetsdatasource
	//
	ShareGoogleSheetsDataSource *string `field:"optional" json:"shareGoogleSheetsDataSource" yaml:"shareGoogleSheetsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharehubspotaction
	//
	ShareHubspotAction *string `field:"optional" json:"shareHubspotAction" yaml:"shareHubspotAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharehuggingfaceaction
	//
	ShareHuggingFaceAction *string `field:"optional" json:"shareHuggingFaceAction" yaml:"shareHuggingFaceAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareidcknowledgebase
	//
	ShareIdcKnowledgeBase *string `field:"optional" json:"shareIdcKnowledgeBase" yaml:"shareIdcKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareimpaladatasource
	//
	ShareImpalaDataSource *string `field:"optional" json:"shareImpalaDataSource" yaml:"shareImpalaDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareintercomaction
	//
	ShareIntercomAction *string `field:"optional" json:"shareIntercomAction" yaml:"shareIntercomAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharejiraaction
	//
	ShareJiraAction *string `field:"optional" json:"shareJiraAction" yaml:"shareJiraAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharejiradatasource
	//
	ShareJiraDataSource *string `field:"optional" json:"shareJiraDataSource" yaml:"shareJiraDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareknowledgebases
	//
	ShareKnowledgeBases *string `field:"optional" json:"shareKnowledgeBases" yaml:"shareKnowledgeBases"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharelinearaction
	//
	ShareLinearAction *string `field:"optional" json:"shareLinearAction" yaml:"shareLinearAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharemariadbdatasource
	//
	ShareMariaDbDataSource *string `field:"optional" json:"shareMariaDbDataSource" yaml:"shareMariaDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharemcpaction
	//
	ShareMcpAction *string `field:"optional" json:"shareMcpAction" yaml:"shareMcpAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharemondayaction
	//
	ShareMondayAction *string `field:"optional" json:"shareMondayAction" yaml:"shareMondayAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharemongoatlasdatasource
	//
	ShareMongoAtlasDataSource *string `field:"optional" json:"shareMongoAtlasDataSource" yaml:"shareMongoAtlasDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharemongodbdatasource
	//
	ShareMongoDbDataSource *string `field:"optional" json:"shareMongoDbDataSource" yaml:"shareMongoDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharemsexchangeaction
	//
	ShareMsExchangeAction *string `field:"optional" json:"shareMsExchangeAction" yaml:"shareMsExchangeAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharemsteamsaction
	//
	ShareMsTeamsAction *string `field:"optional" json:"shareMsTeamsAction" yaml:"shareMsTeamsAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharemysqldatasource
	//
	ShareMySqlDataSource *string `field:"optional" json:"shareMySqlDataSource" yaml:"shareMySqlDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharenewrelicaction
	//
	ShareNewRelicAction *string `field:"optional" json:"shareNewRelicAction" yaml:"shareNewRelicAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharenotionaction
	//
	ShareNotionAction *string `field:"optional" json:"shareNotionAction" yaml:"shareNotionAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareonedriveaction
	//
	ShareOneDriveAction *string `field:"optional" json:"shareOneDriveAction" yaml:"shareOneDriveAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareonedriveknowledgebase
	//
	ShareOneDriveKnowledgeBase *string `field:"optional" json:"shareOneDriveKnowledgeBase" yaml:"shareOneDriveKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareopenapiaction
	//
	ShareOpenApiAction *string `field:"optional" json:"shareOpenApiAction" yaml:"shareOpenApiAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareopensearchdatasource
	//
	ShareOpenSearchDataSource *string `field:"optional" json:"shareOpenSearchDataSource" yaml:"shareOpenSearchDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareoracledatasource
	//
	ShareOracleDataSource *string `field:"optional" json:"shareOracleDataSource" yaml:"shareOracleDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharepagerdutyaction
	//
	SharePagerDutyAction *string `field:"optional" json:"sharePagerDutyAction" yaml:"sharePagerDutyAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharepaypaldatasource
	//
	SharePayPalDataSource *string `field:"optional" json:"sharePayPalDataSource" yaml:"sharePayPalDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharepointaction
	//
	SharePointAction *string `field:"optional" json:"sharePointAction" yaml:"sharePointAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharepointknowledgebase
	//
	SharePointKnowledgeBase *string `field:"optional" json:"sharePointKnowledgeBase" yaml:"sharePointKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharepostgresqldatasource
	//
	SharePostgreSqlDataSource *string `field:"optional" json:"sharePostgreSqlDataSource" yaml:"sharePostgreSqlDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareprestodatasource
	//
	SharePrestoDataSource *string `field:"optional" json:"sharePrestoDataSource" yaml:"sharePrestoDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareqbusinessknowledgebase
	//
	ShareQBusinessKnowledgeBase *string `field:"optional" json:"shareQBusinessKnowledgeBase" yaml:"shareQBusinessKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareradiantdatasource
	//
	ShareRadiantDataSource *string `field:"optional" json:"shareRadiantDataSource" yaml:"shareRadiantDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharerdsdatasource
	//
	ShareRdsDataSource *string `field:"optional" json:"shareRdsDataSource" yaml:"shareRdsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareredshiftautodiscovereddatasource
	//
	ShareRedshiftAutoDiscoveredDataSource *string `field:"optional" json:"shareRedshiftAutoDiscoveredDataSource" yaml:"shareRedshiftAutoDiscoveredDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareredshiftmanualdatasource
	//
	ShareRedshiftManualDataSource *string `field:"optional" json:"shareRedshiftManualDataSource" yaml:"shareRedshiftManualDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shares3analyticsdatasource
	//
	ShareS3AnalyticsDataSource *string `field:"optional" json:"shareS3AnalyticsDataSource" yaml:"shareS3AnalyticsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shares3datasource
	//
	ShareS3DataSource *string `field:"optional" json:"shareS3DataSource" yaml:"shareS3DataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shares3knowledgebase
	//
	ShareS3KnowledgeBase *string `field:"optional" json:"shareS3KnowledgeBase" yaml:"shareS3KnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shares3tablesdatasource
	//
	ShareS3TablesDataSource *string `field:"optional" json:"shareS3TablesDataSource" yaml:"shareS3TablesDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesalesforceaction
	//
	ShareSalesforceAction *string `field:"optional" json:"shareSalesforceAction" yaml:"shareSalesforceAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesalesforcedatasource
	//
	ShareSalesforceDataSource *string `field:"optional" json:"shareSalesforceDataSource" yaml:"shareSalesforceDataSource"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesaphanadatasource
	//
	ShareSapHanaDataSource *string `field:"optional" json:"shareSapHanaDataSource" yaml:"shareSapHanaDataSource"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareservicenowdatasource
	//
	ShareServiceNowDataSource *string `field:"optional" json:"shareServiceNowDataSource" yaml:"shareServiceNowDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesharepointaction
	//
	ShareSharePointAction *string `field:"optional" json:"shareSharePointAction" yaml:"shareSharePointAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesharepointknowledgebase
	//
	ShareSharePointKnowledgeBase *string `field:"optional" json:"shareSharePointKnowledgeBase" yaml:"shareSharePointKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareslackaction
	//
	ShareSlackAction *string `field:"optional" json:"shareSlackAction" yaml:"shareSlackAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesmartsheetaction
	//
	ShareSmartsheetAction *string `field:"optional" json:"shareSmartsheetAction" yaml:"shareSmartsheetAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesnowflakedatasource
	//
	ShareSnowflakeDataSource *string `field:"optional" json:"shareSnowflakeDataSource" yaml:"shareSnowflakeDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharespaces
	//
	ShareSpaces *string `field:"optional" json:"shareSpaces" yaml:"shareSpaces"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesparkdatasource
	//
	ShareSparkDataSource *string `field:"optional" json:"shareSparkDataSource" yaml:"shareSparkDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesqlserverdatasource
	//
	ShareSqlServerDataSource *string `field:"optional" json:"shareSqlServerDataSource" yaml:"shareSqlServerDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharesquaredatasource
	//
	ShareSquareDataSource *string `field:"optional" json:"shareSquareDataSource" yaml:"shareSquareDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharestarburstdatasource
	//
	ShareStarburstDataSource *string `field:"optional" json:"shareStarburstDataSource" yaml:"shareStarburstDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareteradatadatasource
	//
	ShareTeradataDataSource *string `field:"optional" json:"shareTeradataDataSource" yaml:"shareTeradataDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharetextractaction
	//
	ShareTextractAction *string `field:"optional" json:"shareTextractAction" yaml:"shareTextractAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharetimestreamdatasource
	//
	ShareTimestreamDataSource *string `field:"optional" json:"shareTimestreamDataSource" yaml:"shareTimestreamDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharetrinodatasource
	//
	ShareTrinoDataSource *string `field:"optional" json:"shareTrinoDataSource" yaml:"shareTrinoDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharetwitterdatasource
	//
	ShareTwitterDataSource *string `field:"optional" json:"shareTwitterDataSource" yaml:"shareTwitterDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharewebcrawlerknowledgebase
	//
	ShareWebCrawlerKnowledgeBase *string `field:"optional" json:"shareWebCrawlerKnowledgeBase" yaml:"shareWebCrawlerKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharezendeskaction
	//
	ShareZendeskAction *string `field:"optional" json:"shareZendeskAction" yaml:"shareZendeskAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-slackaction
	//
	SlackAction *string `field:"optional" json:"slackAction" yaml:"slackAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-smartsheetaction
	//
	SmartsheetAction *string `field:"optional" json:"smartsheetAction" yaml:"smartsheetAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-snowflakedatasource
	//
	SnowflakeDataSource *string `field:"optional" json:"snowflakeDataSource" yaml:"snowflakeDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-space
	//
	Space *string `field:"optional" json:"space" yaml:"space"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sparkdatasource
	//
	SparkDataSource *string `field:"optional" json:"sparkDataSource" yaml:"sparkDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sqlserverdatasource
	//
	SqlServerDataSource *string `field:"optional" json:"sqlServerDataSource" yaml:"sqlServerDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-squaredatasource
	//
	SquareDataSource *string `field:"optional" json:"squareDataSource" yaml:"squareDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-starburstdatasource
	//
	StarburstDataSource *string `field:"optional" json:"starburstDataSource" yaml:"starburstDataSource"`
	// The ability to subscribe to email reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-subscribedashboardemailreports
	//
	SubscribeDashboardEmailReports *string `field:"optional" json:"subscribeDashboardEmailReports" yaml:"subscribeDashboardEmailReports"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-teradatadatasource
	//
	TeradataDataSource *string `field:"optional" json:"teradataDataSource" yaml:"teradataDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-textractaction
	//
	TextractAction *string `field:"optional" json:"textractAction" yaml:"textractAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-timestreamdatasource
	//
	TimestreamDataSource *string `field:"optional" json:"timestreamDataSource" yaml:"timestreamDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-topic
	//
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-trigger
	//
	Trigger *string `field:"optional" json:"trigger" yaml:"trigger"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-trinodatasource
	//
	TrinoDataSource *string `field:"optional" json:"trinoDataSource" yaml:"trinoDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-twitterdatasource
	//
	TwitterDataSource *string `field:"optional" json:"twitterDataSource" yaml:"twitterDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updateadobeanalyticsdatasource
	//
	UpdateAdobeAnalyticsDataSource *string `field:"optional" json:"updateAdobeAnalyticsDataSource" yaml:"updateAdobeAnalyticsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updateathenadatasource
	//
	UpdateAthenaDataSource *string `field:"optional" json:"updateAthenaDataSource" yaml:"updateAthenaDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updateauroradatasource
	//
	UpdateAuroraDataSource *string `field:"optional" json:"updateAuroraDataSource" yaml:"updateAuroraDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatedatabricksdatasource
	//
	UpdateDatabricksDataSource *string `field:"optional" json:"updateDatabricksDataSource" yaml:"updateDatabricksDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatedb2datasource
	//
	UpdateDb2DataSource *string `field:"optional" json:"updateDb2DataSource" yaml:"updateDb2DataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatedenododatasource
	//
	UpdateDenodoDataSource *string `field:"optional" json:"updateDenodoDataSource" yaml:"updateDenodoDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatedocumentdbdatasource
	//
	UpdateDocumentDbDataSource *string `field:"optional" json:"updateDocumentDbDataSource" yaml:"updateDocumentDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatedremiodatasource
	//
	UpdateDremioDataSource *string `field:"optional" json:"updateDremioDataSource" yaml:"updateDremioDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatedynamodbdatasource
	//
	UpdateDynamoDbDataSource *string `field:"optional" json:"updateDynamoDbDataSource" yaml:"updateDynamoDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updateexasoldatasource
	//
	UpdateExasolDataSource *string `field:"optional" json:"updateExasolDataSource" yaml:"updateExasolDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatefiledatasource
	//
	UpdateFileDataSource *string `field:"optional" json:"updateFileDataSource" yaml:"updateFileDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updategithubdatasource
	//
	UpdateGitHubDataSource *string `field:"optional" json:"updateGitHubDataSource" yaml:"updateGitHubDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updategoogleanalyticsdatasource
	//
	UpdateGoogleAnalyticsDataSource *string `field:"optional" json:"updateGoogleAnalyticsDataSource" yaml:"updateGoogleAnalyticsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updategooglebigquerydatasource
	//
	UpdateGoogleBigQueryDataSource *string `field:"optional" json:"updateGoogleBigQueryDataSource" yaml:"updateGoogleBigQueryDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updategooglesheetsdatasource
	//
	UpdateGoogleSheetsDataSource *string `field:"optional" json:"updateGoogleSheetsDataSource" yaml:"updateGoogleSheetsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updateimpaladatasource
	//
	UpdateImpalaDataSource *string `field:"optional" json:"updateImpalaDataSource" yaml:"updateImpalaDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatejiradatasource
	//
	UpdateJiraDataSource *string `field:"optional" json:"updateJiraDataSource" yaml:"updateJiraDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatemariadbdatasource
	//
	UpdateMariaDbDataSource *string `field:"optional" json:"updateMariaDbDataSource" yaml:"updateMariaDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatemongoatlasdatasource
	//
	UpdateMongoAtlasDataSource *string `field:"optional" json:"updateMongoAtlasDataSource" yaml:"updateMongoAtlasDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatemongodbdatasource
	//
	UpdateMongoDbDataSource *string `field:"optional" json:"updateMongoDbDataSource" yaml:"updateMongoDbDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatemysqldatasource
	//
	UpdateMySqlDataSource *string `field:"optional" json:"updateMySqlDataSource" yaml:"updateMySqlDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updateopensearchdatasource
	//
	UpdateOpenSearchDataSource *string `field:"optional" json:"updateOpenSearchDataSource" yaml:"updateOpenSearchDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updateoracledatasource
	//
	UpdateOracleDataSource *string `field:"optional" json:"updateOracleDataSource" yaml:"updateOracleDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatepaypaldatasource
	//
	UpdatePayPalDataSource *string `field:"optional" json:"updatePayPalDataSource" yaml:"updatePayPalDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatepostgresqldatasource
	//
	UpdatePostgreSqlDataSource *string `field:"optional" json:"updatePostgreSqlDataSource" yaml:"updatePostgreSqlDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updateprestodatasource
	//
	UpdatePrestoDataSource *string `field:"optional" json:"updatePrestoDataSource" yaml:"updatePrestoDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updateradiantdatasource
	//
	UpdateRadiantDataSource *string `field:"optional" json:"updateRadiantDataSource" yaml:"updateRadiantDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updaterdsdatasource
	//
	UpdateRdsDataSource *string `field:"optional" json:"updateRdsDataSource" yaml:"updateRdsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updateredshiftautodiscovereddatasource
	//
	UpdateRedshiftAutoDiscoveredDataSource *string `field:"optional" json:"updateRedshiftAutoDiscoveredDataSource" yaml:"updateRedshiftAutoDiscoveredDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updateredshiftmanualdatasource
	//
	UpdateRedshiftManualDataSource *string `field:"optional" json:"updateRedshiftManualDataSource" yaml:"updateRedshiftManualDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updates3analyticsdatasource
	//
	UpdateS3AnalyticsDataSource *string `field:"optional" json:"updateS3AnalyticsDataSource" yaml:"updateS3AnalyticsDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updates3datasource
	//
	UpdateS3DataSource *string `field:"optional" json:"updateS3DataSource" yaml:"updateS3DataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updates3tablesdatasource
	//
	UpdateS3TablesDataSource *string `field:"optional" json:"updateS3TablesDataSource" yaml:"updateS3TablesDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatesalesforcedatasource
	//
	UpdateSalesforceDataSource *string `field:"optional" json:"updateSalesforceDataSource" yaml:"updateSalesforceDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatesaphanadatasource
	//
	UpdateSapHanaDataSource *string `field:"optional" json:"updateSapHanaDataSource" yaml:"updateSapHanaDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updateservicenowdatasource
	//
	UpdateServiceNowDataSource *string `field:"optional" json:"updateServiceNowDataSource" yaml:"updateServiceNowDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatesnowflakedatasource
	//
	UpdateSnowflakeDataSource *string `field:"optional" json:"updateSnowflakeDataSource" yaml:"updateSnowflakeDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatesparkdatasource
	//
	UpdateSparkDataSource *string `field:"optional" json:"updateSparkDataSource" yaml:"updateSparkDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatesqlserverdatasource
	//
	UpdateSqlServerDataSource *string `field:"optional" json:"updateSqlServerDataSource" yaml:"updateSqlServerDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatesquaredatasource
	//
	UpdateSquareDataSource *string `field:"optional" json:"updateSquareDataSource" yaml:"updateSquareDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatestarburstdatasource
	//
	UpdateStarburstDataSource *string `field:"optional" json:"updateStarburstDataSource" yaml:"updateStarburstDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updateteradatadatasource
	//
	UpdateTeradataDataSource *string `field:"optional" json:"updateTeradataDataSource" yaml:"updateTeradataDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatetimestreamdatasource
	//
	UpdateTimestreamDataSource *string `field:"optional" json:"updateTimestreamDataSource" yaml:"updateTimestreamDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatetrinodatasource
	//
	UpdateTrinoDataSource *string `field:"optional" json:"updateTrinoDataSource" yaml:"updateTrinoDataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-updatetwitterdatasource
	//
	UpdateTwitterDataSource *string `field:"optional" json:"updateTwitterDataSource" yaml:"updateTwitterDataSource"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usebedrockmanagedknowledgebase
	//
	UseBedrockManagedKnowledgeBase *string `field:"optional" json:"useBedrockManagedKnowledgeBase" yaml:"useBedrockManagedKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usebedrockmodels
	//
	UseBedrockModels *string `field:"optional" json:"useBedrockModels" yaml:"useBedrockModels"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useboxagentaction
	//
	UseBoxAgentAction *string `field:"optional" json:"useBoxAgentAction" yaml:"useBoxAgentAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useboxknowledgebase
	//
	UseBoxKnowledgeBase *string `field:"optional" json:"useBoxKnowledgeBase" yaml:"useBoxKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usebrowserextension
	//
	UseBrowserExtension *string `field:"optional" json:"useBrowserExtension" yaml:"useBrowserExtension"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useconfluenceknowledgebase
	//
	UseConfluenceKnowledgeBase *string `field:"optional" json:"useConfluenceKnowledgeBase" yaml:"useConfluenceKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useexceladdinextension
	//
	UseExcelAddInExtension *string `field:"optional" json:"useExcelAddInExtension" yaml:"useExcelAddInExtension"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usegoogledriveknowledgebase
	//
	UseGoogleDriveKnowledgeBase *string `field:"optional" json:"useGoogleDriveKnowledgeBase" yaml:"useGoogleDriveKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usehubspotaction
	//
	UseHubspotAction *string `field:"optional" json:"useHubspotAction" yaml:"useHubspotAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usehuggingfaceaction
	//
	UseHuggingFaceAction *string `field:"optional" json:"useHuggingFaceAction" yaml:"useHuggingFaceAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useidcknowledgebase
	//
	UseIdcKnowledgeBase *string `field:"optional" json:"useIdcKnowledgeBase" yaml:"useIdcKnowledgeBase"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useonedriveknowledgebase
	//
	UseOneDriveKnowledgeBase *string `field:"optional" json:"useOneDriveKnowledgeBase" yaml:"useOneDriveKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useopenapiaction
	//
	UseOpenApiAction *string `field:"optional" json:"useOpenApiAction" yaml:"useOpenApiAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useoutlookaddinextension
	//
	UseOutlookAddInExtension *string `field:"optional" json:"useOutlookAddInExtension" yaml:"useOutlookAddInExtension"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usepagerdutyaction
	//
	UsePagerDutyAction *string `field:"optional" json:"usePagerDutyAction" yaml:"usePagerDutyAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usepowerpointaddinextension
	//
	UsePowerpointAddInExtension *string `field:"optional" json:"usePowerpointAddInExtension" yaml:"usePowerpointAddInExtension"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useqbusinessknowledgebase
	//
	UseQBusinessKnowledgeBase *string `field:"optional" json:"useQBusinessKnowledgeBase" yaml:"useQBusinessKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-uses3knowledgebase
	//
	UseS3KnowledgeBase *string `field:"optional" json:"useS3KnowledgeBase" yaml:"useS3KnowledgeBase"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usesharepointknowledgebase
	//
	UseSharePointKnowledgeBase *string `field:"optional" json:"useSharePointKnowledgeBase" yaml:"useSharePointKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-useslackaction
	//
	UseSlackAction *string `field:"optional" json:"useSlackAction" yaml:"useSlackAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usesmartsheetaction
	//
	UseSmartsheetAction *string `field:"optional" json:"useSmartsheetAction" yaml:"useSmartsheetAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usetextractaction
	//
	UseTextractAction *string `field:"optional" json:"useTextractAction" yaml:"useTextractAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usewebcrawlerknowledgebase
	//
	UseWebCrawlerKnowledgeBase *string `field:"optional" json:"useWebCrawlerKnowledgeBase" yaml:"useWebCrawlerKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usewordaddinextension
	//
	UseWordAddInExtension *string `field:"optional" json:"useWordAddInExtension" yaml:"useWordAddInExtension"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-usezendeskaction
	//
	UseZendeskAction *string `field:"optional" json:"useZendeskAction" yaml:"useZendeskAction"`
	// The ability to view account SPICE capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-viewaccountspicecapacity
	//
	ViewAccountSpiceCapacity *string `field:"optional" json:"viewAccountSpiceCapacity" yaml:"viewAccountSpiceCapacity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-webcrawlerknowledgebase
	//
	WebCrawlerKnowledgeBase *string `field:"optional" json:"webCrawlerKnowledgeBase" yaml:"webCrawlerKnowledgeBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-zendeskaction
	//
	ZendeskAction *string `field:"optional" json:"zendeskAction" yaml:"zendeskAction"`
}

