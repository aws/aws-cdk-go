package awsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a custom permissions profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnCustomPermissionsPropsMixin := awscdkcfnpropertymixins.Aws_quicksight.NewCfnCustomPermissionsPropsMixin(&CfnCustomPermissionsMixinProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	Capabilities: &CapabilitiesProperty{
//   		Action: jsii.String("action"),
//   		AddOrRunAnomalyDetectionForAnalyses: jsii.String("addOrRunAnomalyDetectionForAnalyses"),
//   		AmazonBedrockArsAction: jsii.String("amazonBedrockArsAction"),
//   		AmazonBedrockFsAction: jsii.String("amazonBedrockFsAction"),
//   		AmazonBedrockKrsAction: jsii.String("amazonBedrockKrsAction"),
//   		AmazonSThreeAction: jsii.String("amazonSThreeAction"),
//   		Analysis: jsii.String("analysis"),
//   		ApproveFlowShareRequests: jsii.String("approveFlowShareRequests"),
//   		AsanaAction: jsii.String("asanaAction"),
//   		Automate: jsii.String("automate"),
//   		BambooHrAction: jsii.String("bambooHrAction"),
//   		BoxAgentAction: jsii.String("boxAgentAction"),
//   		BuildCalculatedFieldWithQ: jsii.String("buildCalculatedFieldWithQ"),
//   		CanvaAgentAction: jsii.String("canvaAgentAction"),
//   		ChatAgent: jsii.String("chatAgent"),
//   		ComprehendAction: jsii.String("comprehendAction"),
//   		ComprehendMedicalAction: jsii.String("comprehendMedicalAction"),
//   		ConfluenceAction: jsii.String("confluenceAction"),
//   		CreateAndUpdateAmazonBedrockArsAction: jsii.String("createAndUpdateAmazonBedrockArsAction"),
//   		CreateAndUpdateAmazonBedrockFsAction: jsii.String("createAndUpdateAmazonBedrockFsAction"),
//   		CreateAndUpdateAmazonBedrockKrsAction: jsii.String("createAndUpdateAmazonBedrockKrsAction"),
//   		CreateAndUpdateAmazonSThreeAction: jsii.String("createAndUpdateAmazonSThreeAction"),
//   		CreateAndUpdateAsanaAction: jsii.String("createAndUpdateAsanaAction"),
//   		CreateAndUpdateBambooHrAction: jsii.String("createAndUpdateBambooHrAction"),
//   		CreateAndUpdateBoxAgentAction: jsii.String("createAndUpdateBoxAgentAction"),
//   		CreateAndUpdateCanvaAgentAction: jsii.String("createAndUpdateCanvaAgentAction"),
//   		CreateAndUpdateComprehendAction: jsii.String("createAndUpdateComprehendAction"),
//   		CreateAndUpdateComprehendMedicalAction: jsii.String("createAndUpdateComprehendMedicalAction"),
//   		CreateAndUpdateConfluenceAction: jsii.String("createAndUpdateConfluenceAction"),
//   		CreateAndUpdateDashboardEmailReports: jsii.String("createAndUpdateDashboardEmailReports"),
//   		CreateAndUpdateDatasets: jsii.String("createAndUpdateDatasets"),
//   		CreateAndUpdateDataSources: jsii.String("createAndUpdateDataSources"),
//   		CreateAndUpdateFactSetAction: jsii.String("createAndUpdateFactSetAction"),
//   		CreateAndUpdateGenericHttpAction: jsii.String("createAndUpdateGenericHttpAction"),
//   		CreateAndUpdateGithubAction: jsii.String("createAndUpdateGithubAction"),
//   		CreateAndUpdateGoogleCalendarAction: jsii.String("createAndUpdateGoogleCalendarAction"),
//   		CreateAndUpdateHubspotAction: jsii.String("createAndUpdateHubspotAction"),
//   		CreateAndUpdateHuggingFaceAction: jsii.String("createAndUpdateHuggingFaceAction"),
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
//   		CreateAndUpdateOpenApiAction: jsii.String("createAndUpdateOpenApiAction"),
//   		CreateAndUpdatePagerDutyAction: jsii.String("createAndUpdatePagerDutyAction"),
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
//   		CreateAndUpdateSlackAction: jsii.String("createAndUpdateSlackAction"),
//   		CreateAndUpdateSmartsheetAction: jsii.String("createAndUpdateSmartsheetAction"),
//   		CreateAndUpdateTextractAction: jsii.String("createAndUpdateTextractAction"),
//   		CreateAndUpdateThemes: jsii.String("createAndUpdateThemes"),
//   		CreateAndUpdateThresholdAlerts: jsii.String("createAndUpdateThresholdAlerts"),
//   		CreateAndUpdateZendeskAction: jsii.String("createAndUpdateZendeskAction"),
//   		CreateChatAgents: jsii.String("createChatAgents"),
//   		CreateDashboardExecutiveSummaryWithQ: jsii.String("createDashboardExecutiveSummaryWithQ"),
//   		CreateSharedFolders: jsii.String("createSharedFolders"),
//   		CreateSpiceDataset: jsii.String("createSpiceDataset"),
//   		Dashboard: jsii.String("dashboard"),
//   		EditVisualWithQ: jsii.String("editVisualWithQ"),
//   		ExportToCsv: jsii.String("exportToCsv"),
//   		ExportToCsvInScheduledReports: jsii.String("exportToCsvInScheduledReports"),
//   		ExportToExcel: jsii.String("exportToExcel"),
//   		ExportToExcelInScheduledReports: jsii.String("exportToExcelInScheduledReports"),
//   		ExportToPdf: jsii.String("exportToPdf"),
//   		ExportToPdfInScheduledReports: jsii.String("exportToPdfInScheduledReports"),
//   		Extension: jsii.String("extension"),
//   		FactSetAction: jsii.String("factSetAction"),
//   		Flow: jsii.String("flow"),
//   		GenericHttpAction: jsii.String("genericHttpAction"),
//   		GithubAction: jsii.String("githubAction"),
//   		GoogleCalendarAction: jsii.String("googleCalendarAction"),
//   		HubspotAction: jsii.String("hubspotAction"),
//   		HuggingFaceAction: jsii.String("huggingFaceAction"),
//   		IncludeContentInScheduledReportsEmail: jsii.String("includeContentInScheduledReportsEmail"),
//   		IntercomAction: jsii.String("intercomAction"),
//   		JiraAction: jsii.String("jiraAction"),
//   		KnowledgeBase: jsii.String("knowledgeBase"),
//   		LinearAction: jsii.String("linearAction"),
//   		ManageSharedFolders: jsii.String("manageSharedFolders"),
//   		McpAction: jsii.String("mcpAction"),
//   		MondayAction: jsii.String("mondayAction"),
//   		MsExchangeAction: jsii.String("msExchangeAction"),
//   		MsTeamsAction: jsii.String("msTeamsAction"),
//   		NewRelicAction: jsii.String("newRelicAction"),
//   		NotionAction: jsii.String("notionAction"),
//   		OneDriveAction: jsii.String("oneDriveAction"),
//   		OpenApiAction: jsii.String("openApiAction"),
//   		PagerDutyAction: jsii.String("pagerDutyAction"),
//   		PerformFlowUiTask: jsii.String("performFlowUiTask"),
//   		PrintReports: jsii.String("printReports"),
//   		PublishWithoutApproval: jsii.String("publishWithoutApproval"),
//   		RenameSharedFolders: jsii.String("renameSharedFolders"),
//   		Research: jsii.String("research"),
//   		SalesforceAction: jsii.String("salesforceAction"),
//   		SandPGlobalEnergyAction: jsii.String("sandPGlobalEnergyAction"),
//   		SandPgmiAction: jsii.String("sandPgmiAction"),
//   		SapBillOfMaterialAction: jsii.String("sapBillOfMaterialAction"),
//   		SapBusinessPartnerAction: jsii.String("sapBusinessPartnerAction"),
//   		SapMaterialStockAction: jsii.String("sapMaterialStockAction"),
//   		SapPhysicalInventoryAction: jsii.String("sapPhysicalInventoryAction"),
//   		SapProductMasterDataAction: jsii.String("sapProductMasterDataAction"),
//   		ServiceNowAction: jsii.String("serviceNowAction"),
//   		ShareAmazonBedrockArsAction: jsii.String("shareAmazonBedrockArsAction"),
//   		ShareAmazonBedrockFsAction: jsii.String("shareAmazonBedrockFsAction"),
//   		ShareAmazonBedrockKrsAction: jsii.String("shareAmazonBedrockKrsAction"),
//   		ShareAmazonSThreeAction: jsii.String("shareAmazonSThreeAction"),
//   		ShareAnalyses: jsii.String("shareAnalyses"),
//   		ShareAsanaAction: jsii.String("shareAsanaAction"),
//   		ShareBambooHrAction: jsii.String("shareBambooHrAction"),
//   		ShareBoxAgentAction: jsii.String("shareBoxAgentAction"),
//   		ShareCanvaAgentAction: jsii.String("shareCanvaAgentAction"),
//   		ShareComprehendAction: jsii.String("shareComprehendAction"),
//   		ShareComprehendMedicalAction: jsii.String("shareComprehendMedicalAction"),
//   		ShareConfluenceAction: jsii.String("shareConfluenceAction"),
//   		ShareDashboards: jsii.String("shareDashboards"),
//   		ShareDatasets: jsii.String("shareDatasets"),
//   		ShareDataSources: jsii.String("shareDataSources"),
//   		ShareFactSetAction: jsii.String("shareFactSetAction"),
//   		ShareGenericHttpAction: jsii.String("shareGenericHttpAction"),
//   		ShareGithubAction: jsii.String("shareGithubAction"),
//   		ShareGoogleCalendarAction: jsii.String("shareGoogleCalendarAction"),
//   		ShareHubspotAction: jsii.String("shareHubspotAction"),
//   		ShareHuggingFaceAction: jsii.String("shareHuggingFaceAction"),
//   		ShareIntercomAction: jsii.String("shareIntercomAction"),
//   		ShareJiraAction: jsii.String("shareJiraAction"),
//   		ShareKnowledgeBases: jsii.String("shareKnowledgeBases"),
//   		ShareLinearAction: jsii.String("shareLinearAction"),
//   		ShareMcpAction: jsii.String("shareMcpAction"),
//   		ShareMondayAction: jsii.String("shareMondayAction"),
//   		ShareMsExchangeAction: jsii.String("shareMsExchangeAction"),
//   		ShareMsTeamsAction: jsii.String("shareMsTeamsAction"),
//   		ShareNewRelicAction: jsii.String("shareNewRelicAction"),
//   		ShareNotionAction: jsii.String("shareNotionAction"),
//   		ShareOneDriveAction: jsii.String("shareOneDriveAction"),
//   		ShareOpenApiAction: jsii.String("shareOpenApiAction"),
//   		SharePagerDutyAction: jsii.String("sharePagerDutyAction"),
//   		SharePointAction: jsii.String("sharePointAction"),
//   		ShareSalesforceAction: jsii.String("shareSalesforceAction"),
//   		ShareSandPGlobalEnergyAction: jsii.String("shareSandPGlobalEnergyAction"),
//   		ShareSandPgmiAction: jsii.String("shareSandPgmiAction"),
//   		ShareSapBillOfMaterialAction: jsii.String("shareSapBillOfMaterialAction"),
//   		ShareSapBusinessPartnerAction: jsii.String("shareSapBusinessPartnerAction"),
//   		ShareSapMaterialStockAction: jsii.String("shareSapMaterialStockAction"),
//   		ShareSapPhysicalInventoryAction: jsii.String("shareSapPhysicalInventoryAction"),
//   		ShareSapProductMasterDataAction: jsii.String("shareSapProductMasterDataAction"),
//   		ShareServiceNowAction: jsii.String("shareServiceNowAction"),
//   		ShareSharePointAction: jsii.String("shareSharePointAction"),
//   		ShareSlackAction: jsii.String("shareSlackAction"),
//   		ShareSmartsheetAction: jsii.String("shareSmartsheetAction"),
//   		ShareTextractAction: jsii.String("shareTextractAction"),
//   		ShareZendeskAction: jsii.String("shareZendeskAction"),
//   		SlackAction: jsii.String("slackAction"),
//   		SmartsheetAction: jsii.String("smartsheetAction"),
//   		Space: jsii.String("space"),
//   		SubscribeDashboardEmailReports: jsii.String("subscribeDashboardEmailReports"),
//   		TextractAction: jsii.String("textractAction"),
//   		Topic: jsii.String("topic"),
//   		UseAgentWebSearch: jsii.String("useAgentWebSearch"),
//   		UseAmazonBedrockArsAction: jsii.String("useAmazonBedrockArsAction"),
//   		UseAmazonBedrockFsAction: jsii.String("useAmazonBedrockFsAction"),
//   		UseAmazonBedrockKrsAction: jsii.String("useAmazonBedrockKrsAction"),
//   		UseAmazonSThreeAction: jsii.String("useAmazonSThreeAction"),
//   		UseAsanaAction: jsii.String("useAsanaAction"),
//   		UseBambooHrAction: jsii.String("useBambooHrAction"),
//   		UseBedrockModels: jsii.String("useBedrockModels"),
//   		UseBoxAgentAction: jsii.String("useBoxAgentAction"),
//   		UseCanvaAgentAction: jsii.String("useCanvaAgentAction"),
//   		UseComprehendAction: jsii.String("useComprehendAction"),
//   		UseComprehendMedicalAction: jsii.String("useComprehendMedicalAction"),
//   		UseConfluenceAction: jsii.String("useConfluenceAction"),
//   		UseFactSetAction: jsii.String("useFactSetAction"),
//   		UseGenericHttpAction: jsii.String("useGenericHttpAction"),
//   		UseGithubAction: jsii.String("useGithubAction"),
//   		UseGoogleCalendarAction: jsii.String("useGoogleCalendarAction"),
//   		UseHubspotAction: jsii.String("useHubspotAction"),
//   		UseHuggingFaceAction: jsii.String("useHuggingFaceAction"),
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
//   		UseOpenApiAction: jsii.String("useOpenApiAction"),
//   		UsePagerDutyAction: jsii.String("usePagerDutyAction"),
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
//   		UseSlackAction: jsii.String("useSlackAction"),
//   		UseSmartsheetAction: jsii.String("useSmartsheetAction"),
//   		UseTextractAction: jsii.String("useTextractAction"),
//   		UseZendeskAction: jsii.String("useZendeskAction"),
//   		ViewAccountSpiceCapacity: jsii.String("viewAccountSpiceCapacity"),
//   		ZendeskAction: jsii.String("zendeskAction"),
//   	},
//   	CustomPermissionsName: jsii.String("customPermissionsName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-custompermissions.html
//
type CfnCustomPermissionsPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnCustomPermissionsMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCustomPermissionsPropsMixin
type jsiiProxy_CfnCustomPermissionsPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnCustomPermissionsPropsMixin) Props() *CfnCustomPermissionsMixinProps {
	var returns *CfnCustomPermissionsMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomPermissionsPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QuickSight::CustomPermissions`.
func NewCfnCustomPermissionsPropsMixin(props *CfnCustomPermissionsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnCustomPermissionsPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCustomPermissionsPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCustomPermissionsPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnCustomPermissionsPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QuickSight::CustomPermissions`.
func NewCfnCustomPermissionsPropsMixin_Override(c CfnCustomPermissionsPropsMixin, props *CfnCustomPermissionsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnCustomPermissionsPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnCustomPermissionsPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomPermissionsPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnCustomPermissionsPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCustomPermissionsPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnCustomPermissionsPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCustomPermissionsPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCustomPermissionsPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

