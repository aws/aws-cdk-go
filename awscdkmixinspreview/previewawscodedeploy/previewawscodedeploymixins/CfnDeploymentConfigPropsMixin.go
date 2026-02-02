package previewawscodedeploymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscodedeploy/previewawscodedeploymixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CodeDeploy::DeploymentConfig` resource creates a set of deployment rules, deployment success conditions, and deployment failure conditions that AWS CodeDeploy uses during a deployment.
//
// The deployment configuration specifies the number or percentage of instances that must remain available at any time during a deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDeploymentConfigPropsMixin := awscdkmixinspreview.Mixins.NewCfnDeploymentConfigPropsMixin(&CfnDeploymentConfigMixinProps{
//   	ComputePlatform: jsii.String("computePlatform"),
//   	DeploymentConfigName: jsii.String("deploymentConfigName"),
//   	MinimumHealthyHosts: &MinimumHealthyHostsProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   	TrafficRoutingConfig: &TrafficRoutingConfigProperty{
//   		TimeBasedCanary: &TimeBasedCanaryProperty{
//   			CanaryInterval: jsii.Number(123),
//   			CanaryPercentage: jsii.Number(123),
//   		},
//   		TimeBasedLinear: &TimeBasedLinearProperty{
//   			LinearInterval: jsii.Number(123),
//   			LinearPercentage: jsii.Number(123),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	ZonalConfig: &ZonalConfigProperty{
//   		FirstZoneMonitorDurationInSeconds: jsii.Number(123),
//   		MinimumHealthyHostsPerZone: &MinimumHealthyHostsPerZoneProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.Number(123),
//   		},
//   		MonitorDurationInSeconds: jsii.Number(123),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html
//
type CfnDeploymentConfigPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDeploymentConfigMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDeploymentConfigPropsMixin
type jsiiProxy_CfnDeploymentConfigPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDeploymentConfigPropsMixin) Props() *CfnDeploymentConfigMixinProps {
	var returns *CfnDeploymentConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentConfigPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CodeDeploy::DeploymentConfig`.
func NewCfnDeploymentConfigPropsMixin(props *CfnDeploymentConfigMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDeploymentConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDeploymentConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDeploymentConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CodeDeploy::DeploymentConfig`.
func NewCfnDeploymentConfigPropsMixin_Override(c CfnDeploymentConfigPropsMixin, props *CfnDeploymentConfigMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDeploymentConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDeploymentConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeploymentConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDeploymentConfigPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDeploymentConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

