package awsnovaact

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsnovaact/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::NovaAct::WorkflowDefinition Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnWorkflowDefinitionPropsMixin := awscdkcfnpropertymixins.Aws_novaact.NewCfnWorkflowDefinitionPropsMixin(&CfnWorkflowDefinitionMixinProps{
//   	Description: jsii.String("description"),
//   	ExportConfig: &WorkflowExportConfigProperty{
//   		S3BucketName: jsii.String("s3BucketName"),
//   		S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	},
//   	Name: jsii.String("name"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-novaact-workflowdefinition.html
//
type CfnWorkflowDefinitionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnWorkflowDefinitionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkflowDefinitionPropsMixin
type jsiiProxy_CfnWorkflowDefinitionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnWorkflowDefinitionPropsMixin) Props() *CfnWorkflowDefinitionMixinProps {
	var returns *CfnWorkflowDefinitionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflowDefinitionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NovaAct::WorkflowDefinition`.
func NewCfnWorkflowDefinitionPropsMixin(props *CfnWorkflowDefinitionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnWorkflowDefinitionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkflowDefinitionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkflowDefinitionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_novaact.CfnWorkflowDefinitionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NovaAct::WorkflowDefinition`.
func NewCfnWorkflowDefinitionPropsMixin_Override(c CfnWorkflowDefinitionPropsMixin, props *CfnWorkflowDefinitionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_novaact.CfnWorkflowDefinitionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnWorkflowDefinitionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkflowDefinitionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_novaact.CfnWorkflowDefinitionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkflowDefinitionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_novaact.CfnWorkflowDefinitionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkflowDefinitionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnWorkflowDefinitionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

