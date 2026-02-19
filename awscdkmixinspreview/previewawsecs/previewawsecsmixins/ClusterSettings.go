package previewawsecsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsecs/previewawsecsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Applies one or more cluster settings to an ECS cluster.
//
// If a setting with the same name already exists, its value is replaced.
//
// Example:
//   import ecs "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//
//   cluster := ecs.NewCfnCluster(scope, jsii.String("Cluster"))
//   awscdkmixinspreview.Mixins_Of(cluster).Apply(awscdkmixinspreview.NewClusterSettings([]ClusterSettingsProperty{
//   	&ClusterSettingsProperty{
//   		Name: jsii.String("containerInsights"),
//   		Value: jsii.String("enhanced"),
//   	},
//   }))
//
// Experimental.
type ClusterSettings interface {
	core.Mixin
	// Applies the mixin functionality to the target construct.
	// Experimental.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	// Experimental.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for ClusterSettings
type jsiiProxy_ClusterSettings struct {
	internal.Type__coreMixin
}

// Experimental.
func NewClusterSettings(settings *[]*awsecs.CfnCluster_ClusterSettingsProperty) ClusterSettings {
	_init_.Initialize()

	if err := validateNewClusterSettingsParameters(settings); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClusterSettings{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.ClusterSettings",
		[]interface{}{settings},
		&j,
	)

	return &j
}

// Experimental.
func NewClusterSettings_Override(c ClusterSettings, settings *[]*awsecs.CfnCluster_ClusterSettingsProperty) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.ClusterSettings",
		[]interface{}{settings},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func ClusterSettings_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateClusterSettings_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.ClusterSettings",
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

