package awssagemaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource type definition for AWS::SageMaker::TrialComponent.
//
// A trial component is a stage of a machine learning trial, such as a preprocessing job, training job, or batch transform job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTrialComponentPropsMixin := awscdkcfnpropertymixins.Aws_sagemaker.NewCfnTrialComponentPropsMixin(&CfnTrialComponentMixinProps{
//   	DisplayName: jsii.String("displayName"),
//   	InputArtifacts: map[string]interface{}{
//   		"inputArtifactsKey": &TrialComponentArtifactProperty{
//   			"mediaType": jsii.String("mediaType"),
//   			"value": jsii.String("value"),
//   		},
//   	},
//   	MetadataProperties: &MetadataPropertiesProperty{
//   		CommitId: jsii.String("commitId"),
//   		GeneratedBy: jsii.String("generatedBy"),
//   		ProjectId: jsii.String("projectId"),
//   		Repository: jsii.String("repository"),
//   	},
//   	OutputArtifacts: map[string]interface{}{
//   		"outputArtifactsKey": &TrialComponentArtifactProperty{
//   			"mediaType": jsii.String("mediaType"),
//   			"value": jsii.String("value"),
//   		},
//   	},
//   	Parameters: map[string]interface{}{
//   		"parametersKey": &TrialComponentParameterValueProperty{
//   			"numberValue": jsii.Number(123),
//   			"stringValue": jsii.String("stringValue"),
//   		},
//   	},
//   	Status: &TrialComponentStatusProperty{
//   		Message: jsii.String("message"),
//   		PrimaryStatus: jsii.String("primaryStatus"),
//   	},
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrialComponentName: jsii.String("trialComponentName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-trialcomponent.html
//
type CfnTrialComponentPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTrialComponentMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTrialComponentPropsMixin
type jsiiProxy_CfnTrialComponentPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTrialComponentPropsMixin) Props() *CfnTrialComponentMixinProps {
	var returns *CfnTrialComponentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrialComponentPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::TrialComponent`.
func NewCfnTrialComponentPropsMixin(props *CfnTrialComponentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTrialComponentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTrialComponentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTrialComponentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnTrialComponentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::TrialComponent`.
func NewCfnTrialComponentPropsMixin_Override(c CfnTrialComponentPropsMixin, props *CfnTrialComponentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnTrialComponentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTrialComponentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTrialComponentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnTrialComponentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTrialComponentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnTrialComponentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTrialComponentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTrialComponentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

