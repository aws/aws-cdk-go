package mixinsawsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsquicksight/mixinsawsquicksight/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a custom permissions profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCustomPermissionsPropsMixin := awscdkmixinspreview.Mixins.NewCfnCustomPermissionsPropsMixin(&CfnCustomPermissionsMixinProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	Capabilities: &CapabilitiesProperty{
//   		AddOrRunAnomalyDetectionForAnalyses: jsii.String("addOrRunAnomalyDetectionForAnalyses"),
//   		Analysis: jsii.String("analysis"),
//   		CreateAndUpdateDashboardEmailReports: jsii.String("createAndUpdateDashboardEmailReports"),
//   		CreateAndUpdateDatasets: jsii.String("createAndUpdateDatasets"),
//   		CreateAndUpdateDataSources: jsii.String("createAndUpdateDataSources"),
//   		CreateAndUpdateThemes: jsii.String("createAndUpdateThemes"),
//   		CreateAndUpdateThresholdAlerts: jsii.String("createAndUpdateThresholdAlerts"),
//   		CreateSharedFolders: jsii.String("createSharedFolders"),
//   		CreateSpiceDataset: jsii.String("createSpiceDataset"),
//   		Dashboard: jsii.String("dashboard"),
//   		ExportToCsv: jsii.String("exportToCsv"),
//   		ExportToCsvInScheduledReports: jsii.String("exportToCsvInScheduledReports"),
//   		ExportToExcel: jsii.String("exportToExcel"),
//   		ExportToExcelInScheduledReports: jsii.String("exportToExcelInScheduledReports"),
//   		ExportToPdf: jsii.String("exportToPdf"),
//   		ExportToPdfInScheduledReports: jsii.String("exportToPdfInScheduledReports"),
//   		IncludeContentInScheduledReportsEmail: jsii.String("includeContentInScheduledReportsEmail"),
//   		PrintReports: jsii.String("printReports"),
//   		RenameSharedFolders: jsii.String("renameSharedFolders"),
//   		ShareAnalyses: jsii.String("shareAnalyses"),
//   		ShareDashboards: jsii.String("shareDashboards"),
//   		ShareDatasets: jsii.String("shareDatasets"),
//   		ShareDataSources: jsii.String("shareDataSources"),
//   		SubscribeDashboardEmailReports: jsii.String("subscribeDashboardEmailReports"),
//   		ViewAccountSpiceCapacity: jsii.String("viewAccountSpiceCapacity"),
//   	},
//   	CustomPermissionsName: jsii.String("customPermissionsName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-custompermissions.html
//
type CfnCustomPermissionsPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCustomPermissionsMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCustomPermissionsPropsMixin
type jsiiProxy_CfnCustomPermissionsPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
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

func (j *jsiiProxy_CfnCustomPermissionsPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QuickSight::CustomPermissions`.
func NewCfnCustomPermissionsPropsMixin(props *CfnCustomPermissionsMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCustomPermissionsPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCustomPermissionsPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCustomPermissionsPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnCustomPermissionsPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QuickSight::CustomPermissions`.
func NewCfnCustomPermissionsPropsMixin_Override(c CfnCustomPermissionsPropsMixin, props *CfnCustomPermissionsMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnCustomPermissionsPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCustomPermissionsPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomPermissionsPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnCustomPermissionsPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnCustomPermissionsPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCustomPermissionsPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
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

