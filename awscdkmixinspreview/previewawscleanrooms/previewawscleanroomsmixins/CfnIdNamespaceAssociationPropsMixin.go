package previewawscleanroomsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscleanrooms/previewawscleanroomsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Provides information to create the ID namespace association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdNamespaceAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnIdNamespaceAssociationPropsMixin(&CfnIdNamespaceAssociationMixinProps{
//   	Description: jsii.String("description"),
//   	IdMappingConfig: &IdMappingConfigProperty{
//   		AllowUseAsDimensionColumn: jsii.Boolean(false),
//   	},
//   	InputReferenceConfig: &IdNamespaceAssociationInputReferenceConfigProperty{
//   		InputReferenceArn: jsii.String("inputReferenceArn"),
//   		ManageResourcePolicies: jsii.Boolean(false),
//   	},
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-idnamespaceassociation.html
//
type CfnIdNamespaceAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIdNamespaceAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIdNamespaceAssociationPropsMixin
type jsiiProxy_CfnIdNamespaceAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIdNamespaceAssociationPropsMixin) Props() *CfnIdNamespaceAssociationMixinProps {
	var returns *CfnIdNamespaceAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdNamespaceAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CleanRooms::IdNamespaceAssociation`.
func NewCfnIdNamespaceAssociationPropsMixin(props *CfnIdNamespaceAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIdNamespaceAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIdNamespaceAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIdNamespaceAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnIdNamespaceAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CleanRooms::IdNamespaceAssociation`.
func NewCfnIdNamespaceAssociationPropsMixin_Override(c CfnIdNamespaceAssociationPropsMixin, props *CfnIdNamespaceAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnIdNamespaceAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIdNamespaceAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIdNamespaceAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnIdNamespaceAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIdNamespaceAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnIdNamespaceAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIdNamespaceAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnIdNamespaceAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

