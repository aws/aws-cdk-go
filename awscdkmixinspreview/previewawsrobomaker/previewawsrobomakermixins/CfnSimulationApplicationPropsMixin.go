package previewawsrobomakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsrobomaker/previewawsrobomakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::RoboMaker::SimulationApplication` resource creates an AWS RoboMaker simulation application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSimulationApplicationPropsMixin := awscdkmixinspreview.Mixins.NewCfnSimulationApplicationPropsMixin(&CfnSimulationApplicationMixinProps{
//   	CurrentRevisionId: jsii.String("currentRevisionId"),
//   	Environment: jsii.String("environment"),
//   	Name: jsii.String("name"),
//   	RenderingEngine: &RenderingEngineProperty{
//   		Name: jsii.String("name"),
//   		Version: jsii.String("version"),
//   	},
//   	RobotSoftwareSuite: &RobotSoftwareSuiteProperty{
//   		Name: jsii.String("name"),
//   		Version: jsii.String("version"),
//   	},
//   	SimulationSoftwareSuite: &SimulationSoftwareSuiteProperty{
//   		Name: jsii.String("name"),
//   		Version: jsii.String("version"),
//   	},
//   	Sources: []interface{}{
//   		&SourceConfigProperty{
//   			Architecture: jsii.String("architecture"),
//   			S3Bucket: jsii.String("s3Bucket"),
//   			S3Key: jsii.String("s3Key"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-simulationapplication.html
//
type CfnSimulationApplicationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSimulationApplicationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSimulationApplicationPropsMixin
type jsiiProxy_CfnSimulationApplicationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSimulationApplicationPropsMixin) Props() *CfnSimulationApplicationMixinProps {
	var returns *CfnSimulationApplicationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimulationApplicationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RoboMaker::SimulationApplication`.
func NewCfnSimulationApplicationPropsMixin(props *CfnSimulationApplicationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSimulationApplicationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSimulationApplicationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSimulationApplicationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_robomaker.mixins.CfnSimulationApplicationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RoboMaker::SimulationApplication`.
func NewCfnSimulationApplicationPropsMixin_Override(c CfnSimulationApplicationPropsMixin, props *CfnSimulationApplicationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_robomaker.mixins.CfnSimulationApplicationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSimulationApplicationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSimulationApplicationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_robomaker.mixins.CfnSimulationApplicationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSimulationApplicationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_robomaker.mixins.CfnSimulationApplicationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSimulationApplicationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSimulationApplicationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

