package previewawsneptunemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsneptune/previewawsneptunemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// `AWS::Neptune::DBParameterGroup` creates a new DB parameter group.
//
// This type can be declared in a template and referenced in the `DBParameterGroupName` parameter of `AWS::Neptune::DBInstance` .
//
// > Applying a parameter group to a DB instance might require the instance to reboot, resulting in a database outage for the duration of the reboot. > If you provide a custom `DBParameterGroup` that you associate with `DBInstance` , it is best to specify an `EngineVersion` property in `DBCluster` . That `EngineVersion` needs to be compatible with the value of the `Family` property in the `DBParameterGroup` .
//
// A DB parameter group is initially created with the default parameters for the database engine used by the DB instance. To provide custom values for any of the parameters, you must modify the group after creating it using *ModifyDBParameterGroup* . Once you've created a DB parameter group, you need to associate it with your DB instance using *ModifyDBInstance* . When you associate a new DB parameter group with a running DB instance, you need to reboot the DB instance without failover for the new DB parameter group and associated settings to take effect.
//
// > After you create a DB parameter group, you should wait at least 5 minutes before creating your first DB instance that uses that DB parameter group as the default parameter group. This allows Amazon Neptune to fully complete the create action before the parameter group is used as the default for a new DB instance. This is especially important for parameters that are critical when creating the default database for a DB instance, such as the character set for the default database defined by the `character_set_database` parameter. You can use the *Parameter Groups* option of the Amazon Neptune console or the *DescribeDBParameters* command to verify that your DB parameter group has been created or modified.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var parameters interface{}
//
//   cfnDBParameterGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnDBParameterGroupPropsMixin(&CfnDBParameterGroupMixinProps{
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
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbparametergroup.html
//
type CfnDBParameterGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDBParameterGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDBParameterGroupPropsMixin
type jsiiProxy_CfnDBParameterGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDBParameterGroupPropsMixin) Props() *CfnDBParameterGroupMixinProps {
	var returns *CfnDBParameterGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBParameterGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Neptune::DBParameterGroup`.
func NewCfnDBParameterGroupPropsMixin(props *CfnDBParameterGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDBParameterGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDBParameterGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDBParameterGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_neptune.mixins.CfnDBParameterGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Neptune::DBParameterGroup`.
func NewCfnDBParameterGroupPropsMixin_Override(c CfnDBParameterGroupPropsMixin, props *CfnDBParameterGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_neptune.mixins.CfnDBParameterGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDBParameterGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBParameterGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_neptune.mixins.CfnDBParameterGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBParameterGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_neptune.mixins.CfnDBParameterGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDBParameterGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDBParameterGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

