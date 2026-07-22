package awspersonalize

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awspersonalize/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Personalize::DataDeletionJob.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDataDeletionJobPropsMixin := awscdkcfnpropertymixins.Aws_personalize.NewCfnDataDeletionJobPropsMixin(&CfnDataDeletionJobMixinProps{
//   	DatasetGroupArn: jsii.String("datasetGroupArn"),
//   	DataSource: &DataSourceProperty{
//   		DataLocation: jsii.String("dataLocation"),
//   	},
//   	JobName: jsii.String("jobName"),
//   	RoleArn: jsii.String("roleArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datadeletionjob.html
//
type CfnDataDeletionJobPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDataDeletionJobMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataDeletionJobPropsMixin
type jsiiProxy_CfnDataDeletionJobPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDataDeletionJobPropsMixin) Props() *CfnDataDeletionJobMixinProps {
	var returns *CfnDataDeletionJobMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataDeletionJobPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Personalize::DataDeletionJob`.
func NewCfnDataDeletionJobPropsMixin(props *CfnDataDeletionJobMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDataDeletionJobPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataDeletionJobPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataDeletionJobPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnDataDeletionJobPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Personalize::DataDeletionJob`.
func NewCfnDataDeletionJobPropsMixin_Override(c CfnDataDeletionJobPropsMixin, props *CfnDataDeletionJobMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnDataDeletionJobPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDataDeletionJobPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataDeletionJobPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnDataDeletionJobPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataDeletionJobPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnDataDeletionJobPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataDeletionJobPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDataDeletionJobPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

