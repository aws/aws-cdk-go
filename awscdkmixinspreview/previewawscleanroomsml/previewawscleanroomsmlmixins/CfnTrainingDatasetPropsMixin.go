package previewawscleanroomsmlmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscleanroomsml/previewawscleanroomsmlmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines the information necessary to create a training dataset.
//
// In Clean Rooms ML, the `TrainingDataset` is metadata that points to a Glue table, which is read only during `AudienceModel` creation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTrainingDatasetPropsMixin := awscdkmixinspreview.Mixins.NewCfnTrainingDatasetPropsMixin(&CfnTrainingDatasetMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrainingData: []interface{}{
//   		&DatasetProperty{
//   			InputConfig: &DatasetInputConfigProperty{
//   				DataSource: &DataSourceProperty{
//   					GlueDataSource: &GlueDataSourceProperty{
//   						CatalogId: jsii.String("catalogId"),
//   						DatabaseName: jsii.String("databaseName"),
//   						TableName: jsii.String("tableName"),
//   					},
//   				},
//   				Schema: []interface{}{
//   					&ColumnSchemaProperty{
//   						ColumnName: jsii.String("columnName"),
//   						ColumnTypes: []*string{
//   							jsii.String("columnTypes"),
//   						},
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-trainingdataset.html
//
type CfnTrainingDatasetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTrainingDatasetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTrainingDatasetPropsMixin
type jsiiProxy_CfnTrainingDatasetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTrainingDatasetPropsMixin) Props() *CfnTrainingDatasetMixinProps {
	var returns *CfnTrainingDatasetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrainingDatasetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CleanRoomsML::TrainingDataset`.
func NewCfnTrainingDatasetPropsMixin(props *CfnTrainingDatasetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTrainingDatasetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTrainingDatasetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTrainingDatasetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanroomsml.mixins.CfnTrainingDatasetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CleanRoomsML::TrainingDataset`.
func NewCfnTrainingDatasetPropsMixin_Override(c CfnTrainingDatasetPropsMixin, props *CfnTrainingDatasetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanroomsml.mixins.CfnTrainingDatasetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTrainingDatasetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTrainingDatasetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cleanroomsml.mixins.CfnTrainingDatasetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTrainingDatasetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cleanroomsml.mixins.CfnTrainingDatasetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTrainingDatasetPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTrainingDatasetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

