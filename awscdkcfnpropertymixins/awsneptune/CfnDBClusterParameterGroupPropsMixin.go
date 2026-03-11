package awsneptune

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsneptune/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Neptune::DBClusterParameterGroup` resource creates a new Amazon Neptune DB cluster parameter group.
//
// > Applying a parameter group to a DB cluster might require instances to reboot, resulting in a database outage while the instances reboot. > If you provide a custom `DBClusterParameterGroup` that you associate with the `DBCluster` , it is best to specify an `EngineVersion` property in the `DBCluster` . That `EngineVersion` needs to be compatible with the value of the `Family` property in the `DBClusterParameterGroup` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var parameters interface{}
//
//   cfnDBClusterParameterGroupPropsMixin := awscdkcfnpropertymixins.Aws_neptune.NewCfnDBClusterParameterGroupPropsMixin(&CfnDBClusterParameterGroupMixinProps{
//   	Description: jsii.String("description"),
//   	Family: jsii.String("family"),
//   	Name: jsii.String("name"),
//   	Parameters: parameters,
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbclusterparametergroup.html
//
type CfnDBClusterParameterGroupPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDBClusterParameterGroupMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDBClusterParameterGroupPropsMixin
type jsiiProxy_CfnDBClusterParameterGroupPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDBClusterParameterGroupPropsMixin) Props() *CfnDBClusterParameterGroupMixinProps {
	var returns *CfnDBClusterParameterGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroupPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Neptune::DBClusterParameterGroup`.
func NewCfnDBClusterParameterGroupPropsMixin(props *CfnDBClusterParameterGroupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDBClusterParameterGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDBClusterParameterGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDBClusterParameterGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_neptune.CfnDBClusterParameterGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Neptune::DBClusterParameterGroup`.
func NewCfnDBClusterParameterGroupPropsMixin_Override(c CfnDBClusterParameterGroupPropsMixin, props *CfnDBClusterParameterGroupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_neptune.CfnDBClusterParameterGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDBClusterParameterGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBClusterParameterGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_neptune.CfnDBClusterParameterGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBClusterParameterGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_neptune.CfnDBClusterParameterGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDBClusterParameterGroupPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDBClusterParameterGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

