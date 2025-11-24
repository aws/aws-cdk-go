package mixinsawsdatazone

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsdatazone/mixinsawsdatazone/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DataZone::DataSource` resource specifies an Amazon DataZone data source that is used to import technical metadata of assets (data) from the source databases or data warehouses into Amazon DataZone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataSourcePropsMixin := awscdkmixinspreview.Mixins.NewCfnDataSourcePropsMixin(&CfnDataSourceMixinProps{
//   	AssetFormsInput: []interface{}{
//   		&FormInputProperty{
//   			Content: jsii.String("content"),
//   			FormName: jsii.String("formName"),
//   			TypeIdentifier: jsii.String("typeIdentifier"),
//   			TypeRevision: jsii.String("typeRevision"),
//   		},
//   	},
//   	Configuration: &DataSourceConfigurationInputProperty{
//   		GlueRunConfiguration: &GlueRunConfigurationInputProperty{
//   			AutoImportDataQualityResult: jsii.Boolean(false),
//   			CatalogName: jsii.String("catalogName"),
//   			DataAccessRole: jsii.String("dataAccessRole"),
//   			RelationalFilterConfigurations: []interface{}{
//   				&RelationalFilterConfigurationProperty{
//   					DatabaseName: jsii.String("databaseName"),
//   					FilterExpressions: []interface{}{
//   						&FilterExpressionProperty{
//   							Expression: jsii.String("expression"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					SchemaName: jsii.String("schemaName"),
//   				},
//   			},
//   		},
//   		RedshiftRunConfiguration: &RedshiftRunConfigurationInputProperty{
//   			DataAccessRole: jsii.String("dataAccessRole"),
//   			RedshiftCredentialConfiguration: &RedshiftCredentialConfigurationProperty{
//   				SecretManagerArn: jsii.String("secretManagerArn"),
//   			},
//   			RedshiftStorage: &RedshiftStorageProperty{
//   				RedshiftClusterSource: &RedshiftClusterStorageProperty{
//   					ClusterName: jsii.String("clusterName"),
//   				},
//   				RedshiftServerlessSource: &RedshiftServerlessStorageProperty{
//   					WorkgroupName: jsii.String("workgroupName"),
//   				},
//   			},
//   			RelationalFilterConfigurations: []interface{}{
//   				&RelationalFilterConfigurationProperty{
//   					DatabaseName: jsii.String("databaseName"),
//   					FilterExpressions: []interface{}{
//   						&FilterExpressionProperty{
//   							Expression: jsii.String("expression"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					SchemaName: jsii.String("schemaName"),
//   				},
//   			},
//   		},
//   		SageMakerRunConfiguration: &SageMakerRunConfigurationInputProperty{
//   			TrackingAssets: map[string][]*string{
//   				"trackingAssetsKey": []*string{
//   					jsii.String("trackingAssets"),
//   				},
//   			},
//   		},
//   	},
//   	ConnectionIdentifier: jsii.String("connectionIdentifier"),
//   	Description: jsii.String("description"),
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	EnableSetting: jsii.String("enableSetting"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	Name: jsii.String("name"),
//   	ProjectIdentifier: jsii.String("projectIdentifier"),
//   	PublishOnImport: jsii.Boolean(false),
//   	Recommendation: &RecommendationConfigurationProperty{
//   		EnableBusinessNameGeneration: jsii.Boolean(false),
//   	},
//   	Schedule: &ScheduleConfigurationProperty{
//   		Schedule: jsii.String("schedule"),
//   		Timezone: jsii.String("timezone"),
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html
//
type CfnDataSourcePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDataSourceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataSourcePropsMixin
type jsiiProxy_CfnDataSourcePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDataSourcePropsMixin) Props() *CfnDataSourceMixinProps {
	var returns *CfnDataSourceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSourcePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataZone::DataSource`.
func NewCfnDataSourcePropsMixin(props *CfnDataSourceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDataSourcePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataSourcePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataSourcePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataZone::DataSource`.
func NewCfnDataSourcePropsMixin_Override(c CfnDataSourcePropsMixin, props *CfnDataSourceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDataSourcePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSourcePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataSourcePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnDataSourcePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataSourcePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDataSourcePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

