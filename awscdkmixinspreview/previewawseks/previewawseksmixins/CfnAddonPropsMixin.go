package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawseks/previewawseksmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon EKS add-on.
//
// Amazon EKS add-ons help to automate the provisioning and lifecycle management of common operational software for Amazon EKS clusters. For more information, see [Amazon EKS add-ons](https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html) in the *Amazon EKS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAddonPropsMixin := awscdkmixinspreview.Mixins.NewCfnAddonPropsMixin(&CfnAddonMixinProps{
//   	AddonName: jsii.String("addonName"),
//   	AddonVersion: jsii.String("addonVersion"),
//   	ClusterName: jsii.String("clusterName"),
//   	ConfigurationValues: jsii.String("configurationValues"),
//   	NamespaceConfig: &NamespaceConfigProperty{
//   		Namespace: jsii.String("namespace"),
//   	},
//   	PodIdentityAssociations: []interface{}{
//   		&PodIdentityAssociationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			ServiceAccount: jsii.String("serviceAccount"),
//   		},
//   	},
//   	PreserveOnDelete: jsii.Boolean(false),
//   	ResolveConflicts: jsii.String("resolveConflicts"),
//   	ServiceAccountRoleArn: jsii.String("serviceAccountRoleArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html
//
type CfnAddonPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAddonMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAddonPropsMixin
type jsiiProxy_CfnAddonPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAddonPropsMixin) Props() *CfnAddonMixinProps {
	var returns *CfnAddonMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddonPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EKS::Addon`.
func NewCfnAddonPropsMixin(props *CfnAddonMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAddonPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAddonPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAddonPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAddonPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EKS::Addon`.
func NewCfnAddonPropsMixin_Override(c CfnAddonPropsMixin, props *CfnAddonMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAddonPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAddonPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAddonPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAddonPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAddonPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAddonPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAddonPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAddonPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

