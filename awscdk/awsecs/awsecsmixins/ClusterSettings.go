package awsecsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/awsecsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Applies one or more cluster settings to an ECS cluster.
//
// If a setting with the same name already exists, its value is replaced.
//
// Example:
//   ecs.NewCfnCluster(this, jsii.String("Cluster")).With(mixins.NewClusterSettings([]ClusterSettingsProperty{
//   	&ClusterSettingsProperty{
//   		Name: jsii.String("containerInsights"),
//   		Value: jsii.String("enhanced"),
//   	},
//   }))
//
type ClusterSettings interface {
	awscdk.Mixin
	// Applies the mixin functionality to the target construct.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for ClusterSettings
type jsiiProxy_ClusterSettings struct {
	internal.Type__awscdkMixin
}

func NewClusterSettings(settings *[]*awsecs.CfnCluster_ClusterSettingsProperty) ClusterSettings {
	_init_.Initialize()

	if err := validateNewClusterSettingsParameters(settings); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClusterSettings{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.mixins.ClusterSettings",
		[]interface{}{settings},
		&j,
	)

	return &j
}

func NewClusterSettings_Override(c ClusterSettings, settings *[]*awsecs.CfnCluster_ClusterSettingsProperty) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.mixins.ClusterSettings",
		[]interface{}{settings},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func ClusterSettings_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateClusterSettings_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.mixins.ClusterSettings",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterSettings) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_ClusterSettings) Supports(construct constructs.IConstruct) *bool {
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

