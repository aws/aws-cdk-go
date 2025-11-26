package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawseks/previewawseksmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an access entry.
//
// An access entry allows an IAM principal to access your cluster. Access entries can replace the need to maintain entries in the `aws-auth` `ConfigMap` for authentication. You have the following options for authorizing an IAM principal to access Kubernetes objects on your cluster: Kubernetes role-based access control (RBAC), Amazon EKS, or both. Kubernetes RBAC authorization requires you to create and manage Kubernetes `Role` , `ClusterRole` , `RoleBinding` , and `ClusterRoleBinding` objects, in addition to managing access entries. If you use Amazon EKS authorization exclusively, you don't need to create and manage Kubernetes `Role` , `ClusterRole` , `RoleBinding` , and `ClusterRoleBinding` objects.
//
// For more information about access entries, see [Access entries](https://docs.aws.amazon.com/eks/latest/userguide/access-entries.html) in the *Amazon EKS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAccessEntryPropsMixin := awscdkmixinspreview.Mixins.NewCfnAccessEntryPropsMixin(&CfnAccessEntryMixinProps{
//   	AccessPolicies: []interface{}{
//   		&AccessPolicyProperty{
//   			AccessScope: &AccessScopeProperty{
//   				Namespaces: []*string{
//   					jsii.String("namespaces"),
//   				},
//   				Type: jsii.String("type"),
//   			},
//   			PolicyArn: jsii.String("policyArn"),
//   		},
//   	},
//   	ClusterName: jsii.String("clusterName"),
//   	KubernetesGroups: []*string{
//   		jsii.String("kubernetesGroups"),
//   	},
//   	PrincipalArn: jsii.String("principalArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	Username: jsii.String("username"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-accessentry.html
//
type CfnAccessEntryPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAccessEntryMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAccessEntryPropsMixin
type jsiiProxy_CfnAccessEntryPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAccessEntryPropsMixin) Props() *CfnAccessEntryMixinProps {
	var returns *CfnAccessEntryMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessEntryPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EKS::AccessEntry`.
func NewCfnAccessEntryPropsMixin(props *CfnAccessEntryMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAccessEntryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAccessEntryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAccessEntryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAccessEntryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EKS::AccessEntry`.
func NewCfnAccessEntryPropsMixin_Override(c CfnAccessEntryPropsMixin, props *CfnAccessEntryMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAccessEntryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAccessEntryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAccessEntryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAccessEntryPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccessEntryPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAccessEntryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAccessEntryPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAccessEntryPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

