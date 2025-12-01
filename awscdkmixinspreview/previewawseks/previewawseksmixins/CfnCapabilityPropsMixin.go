package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawseks/previewawseksmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for EKS Capability.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityPropsMixin := awscdkmixinspreview.Mixins.NewCfnCapabilityPropsMixin(&CfnCapabilityMixinProps{
//   	CapabilityName: jsii.String("capabilityName"),
//   	ClusterName: jsii.String("clusterName"),
//   	Configuration: &CapabilityConfigurationProperty{
//   		ArgoCd: &ArgoCdProperty{
//   			AwsIdc: &AwsIdcProperty{
//   				IdcInstanceArn: jsii.String("idcInstanceArn"),
//   				IdcManagedApplicationArn: jsii.String("idcManagedApplicationArn"),
//   				IdcRegion: jsii.String("idcRegion"),
//   			},
//   			Namespace: jsii.String("namespace"),
//   			NetworkAccess: &NetworkAccessProperty{
//   				VpceIds: []*string{
//   					jsii.String("vpceIds"),
//   				},
//   			},
//   			RbacRoleMappings: []interface{}{
//   				&ArgoCdRoleMappingProperty{
//   					Identities: []interface{}{
//   						&SsoIdentityProperty{
//   							Id: jsii.String("id"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					Role: jsii.String("role"),
//   				},
//   			},
//   			ServerUrl: jsii.String("serverUrl"),
//   		},
//   	},
//   	DeletePropagationPolicy: jsii.String("deletePropagationPolicy"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html
//
type CfnCapabilityPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCapabilityMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCapabilityPropsMixin
type jsiiProxy_CfnCapabilityPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCapabilityPropsMixin) Props() *CfnCapabilityMixinProps {
	var returns *CfnCapabilityMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapabilityPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EKS::Capability`.
func NewCfnCapabilityPropsMixin(props *CfnCapabilityMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCapabilityPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCapabilityPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCapabilityPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EKS::Capability`.
func NewCfnCapabilityPropsMixin_Override(c CfnCapabilityPropsMixin, props *CfnCapabilityMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCapabilityPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCapabilityPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCapabilityPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCapabilityPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCapabilityPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

