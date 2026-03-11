package awspersonalize

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awspersonalize/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an empty dataset and adds it to the specified dataset group.
//
// Use [CreateDatasetImportJob](https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetImportJob.html) to import your training data to a dataset.
//
// There are 5 types of datasets:
//
// - Item interactions
// - Items
// - Users
// - Action interactions (you can't use CloudFormation to create an Action interactions dataset)
// - Actions (you can't use CloudFormation to create an Actions dataset)
//
// Each dataset type has an associated schema with required field types. Only the `Item interactions` dataset is required in order to train a model (also referred to as creating a solution).
//
// A dataset can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
// - DELETE PENDING > DELETE IN_PROGRESS
//
// To get the status of the dataset, call [DescribeDataset](https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDataset.html) .
//
// **Related APIs** - [CreateDatasetGroup](https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetGroup.html)
// - [ListDatasets](https://docs.aws.amazon.com/personalize/latest/dg/API_ListDatasets.html)
// - [DescribeDataset](https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDataset.html)
// - [DeleteDataset](https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteDataset.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataSource interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnDatasetPropsMixin := awscdkcfnpropertymixins.Aws_personalize.NewCfnDatasetPropsMixin(&CfnDatasetMixinProps{
//   	DatasetGroupArn: jsii.String("datasetGroupArn"),
//   	DatasetImportJob: &DatasetImportJobProperty{
//   		DatasetArn: jsii.String("datasetArn"),
//   		DatasetImportJobArn: jsii.String("datasetImportJobArn"),
//   		DataSource: dataSource,
//   		JobName: jsii.String("jobName"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	DatasetType: jsii.String("datasetType"),
//   	Name: jsii.String("name"),
//   	SchemaArn: jsii.String("schemaArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-dataset.html
//
type CfnDatasetPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDatasetMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDatasetPropsMixin
type jsiiProxy_CfnDatasetPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDatasetPropsMixin) Props() *CfnDatasetMixinProps {
	var returns *CfnDatasetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Personalize::Dataset`.
func NewCfnDatasetPropsMixin(props *CfnDatasetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDatasetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDatasetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDatasetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnDatasetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Personalize::Dataset`.
func NewCfnDatasetPropsMixin_Override(c CfnDatasetPropsMixin, props *CfnDatasetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnDatasetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDatasetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDatasetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnDatasetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDatasetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnDatasetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDatasetPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDatasetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

