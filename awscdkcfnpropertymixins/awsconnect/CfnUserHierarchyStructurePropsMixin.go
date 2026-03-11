package awsconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Contains information about a hierarchy structure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnUserHierarchyStructurePropsMixin := awscdkcfnpropertymixins.Aws_connect.NewCfnUserHierarchyStructurePropsMixin(&CfnUserHierarchyStructureMixinProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	UserHierarchyStructure: &UserHierarchyStructureProperty{
//   		LevelFive: &LevelFiveProperty{
//   			HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   			HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   			Name: jsii.String("name"),
//   		},
//   		LevelFour: &LevelFourProperty{
//   			HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   			HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   			Name: jsii.String("name"),
//   		},
//   		LevelOne: &LevelOneProperty{
//   			HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   			HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   			Name: jsii.String("name"),
//   		},
//   		LevelThree: &LevelThreeProperty{
//   			HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   			HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   			Name: jsii.String("name"),
//   		},
//   		LevelTwo: &LevelTwoProperty{
//   			HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   			HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-userhierarchystructure.html
//
type CfnUserHierarchyStructurePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnUserHierarchyStructureMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnUserHierarchyStructurePropsMixin
type jsiiProxy_CfnUserHierarchyStructurePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnUserHierarchyStructurePropsMixin) Props() *CfnUserHierarchyStructureMixinProps {
	var returns *CfnUserHierarchyStructureMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserHierarchyStructurePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::UserHierarchyStructure`.
func NewCfnUserHierarchyStructurePropsMixin(props *CfnUserHierarchyStructureMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnUserHierarchyStructurePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUserHierarchyStructurePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserHierarchyStructurePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnUserHierarchyStructurePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::UserHierarchyStructure`.
func NewCfnUserHierarchyStructurePropsMixin_Override(c CfnUserHierarchyStructurePropsMixin, props *CfnUserHierarchyStructureMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnUserHierarchyStructurePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnUserHierarchyStructurePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserHierarchyStructurePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnUserHierarchyStructurePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserHierarchyStructurePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnUserHierarchyStructurePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserHierarchyStructurePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnUserHierarchyStructurePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

