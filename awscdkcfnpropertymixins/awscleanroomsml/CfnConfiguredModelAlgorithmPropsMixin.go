package awscleanroomsml

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscleanroomsml/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::CleanRoomsML::ConfiguredModelAlgorithm Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnConfiguredModelAlgorithmPropsMixin := awscdkcfnpropertymixins.Aws_cleanroomsml.NewCfnConfiguredModelAlgorithmPropsMixin(&CfnConfiguredModelAlgorithmMixinProps{
//   	Description: jsii.String("description"),
//   	InferenceContainerConfig: &InferenceContainerConfigProperty{
//   		ImageUri: jsii.String("imageUri"),
//   	},
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrainingContainerConfig: &ContainerConfigProperty{
//   		Arguments: []*string{
//   			jsii.String("arguments"),
//   		},
//   		Entrypoint: []*string{
//   			jsii.String("entrypoint"),
//   		},
//   		ImageUri: jsii.String("imageUri"),
//   		MetricDefinitions: []interface{}{
//   			&MetricDefinitionProperty{
//   				Name: jsii.String("name"),
//   				Regex: jsii.String("regex"),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithm.html
//
type CfnConfiguredModelAlgorithmPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnConfiguredModelAlgorithmMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConfiguredModelAlgorithmPropsMixin
type jsiiProxy_CfnConfiguredModelAlgorithmPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnConfiguredModelAlgorithmPropsMixin) Props() *CfnConfiguredModelAlgorithmMixinProps {
	var returns *CfnConfiguredModelAlgorithmMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguredModelAlgorithmPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CleanRoomsML::ConfiguredModelAlgorithm`.
func NewCfnConfiguredModelAlgorithmPropsMixin(props *CfnConfiguredModelAlgorithmMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnConfiguredModelAlgorithmPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfiguredModelAlgorithmPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfiguredModelAlgorithmPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cleanroomsml.CfnConfiguredModelAlgorithmPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CleanRoomsML::ConfiguredModelAlgorithm`.
func NewCfnConfiguredModelAlgorithmPropsMixin_Override(c CfnConfiguredModelAlgorithmPropsMixin, props *CfnConfiguredModelAlgorithmMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cleanroomsml.CfnConfiguredModelAlgorithmPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnConfiguredModelAlgorithmPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConfiguredModelAlgorithmPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cleanroomsml.CfnConfiguredModelAlgorithmPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfiguredModelAlgorithmPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cleanroomsml.CfnConfiguredModelAlgorithmPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfiguredModelAlgorithmPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnConfiguredModelAlgorithmPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

