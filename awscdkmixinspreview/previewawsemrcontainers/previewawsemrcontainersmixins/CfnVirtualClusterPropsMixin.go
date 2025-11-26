package previewawsemrcontainersmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsemrcontainers/previewawsemrcontainersmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::EMRContainers::VirtualCluster` resource specifies a virtual cluster.
//
// A virtual cluster is a managed entity on Amazon EMR on EKS. You can create, describe, list, and delete virtual clusters. They do not consume any additional resources in your system. A single virtual cluster maps to a single Kubernetes namespace. Given this relationship, you can model virtual clusters the same way you model Kubernetes namespaces to meet your requirements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVirtualClusterPropsMixin := awscdkmixinspreview.Mixins.NewCfnVirtualClusterPropsMixin(&CfnVirtualClusterMixinProps{
//   	ContainerProvider: &ContainerProviderProperty{
//   		Id: jsii.String("id"),
//   		Info: &ContainerInfoProperty{
//   			EksInfo: &EksInfoProperty{
//   				Namespace: jsii.String("namespace"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	Name: jsii.String("name"),
//   	SecurityConfigurationId: jsii.String("securityConfigurationId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-virtualcluster.html
//
type CfnVirtualClusterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnVirtualClusterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVirtualClusterPropsMixin
type jsiiProxy_CfnVirtualClusterPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnVirtualClusterPropsMixin) Props() *CfnVirtualClusterMixinProps {
	var returns *CfnVirtualClusterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualClusterPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EMRContainers::VirtualCluster`.
func NewCfnVirtualClusterPropsMixin(props *CfnVirtualClusterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnVirtualClusterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVirtualClusterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVirtualClusterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnVirtualClusterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EMRContainers::VirtualCluster`.
func NewCfnVirtualClusterPropsMixin_Override(c CfnVirtualClusterPropsMixin, props *CfnVirtualClusterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnVirtualClusterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnVirtualClusterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVirtualClusterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnVirtualClusterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVirtualClusterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnVirtualClusterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVirtualClusterPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnVirtualClusterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

