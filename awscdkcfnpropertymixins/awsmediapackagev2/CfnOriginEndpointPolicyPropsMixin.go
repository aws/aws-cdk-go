package awsmediapackagev2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmediapackagev2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies the configuration parameters of a policy associated with a MediaPackage V2 origin endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var policy interface{}
//
//   cfnOriginEndpointPolicyPropsMixin := awscdkcfnpropertymixins.Aws_mediapackagev2.NewCfnOriginEndpointPolicyPropsMixin(&CfnOriginEndpointPolicyMixinProps{
//   	CdnAuthConfiguration: &CdnAuthConfigurationProperty{
//   		CdnIdentifierSecretArns: []*string{
//   			jsii.String("cdnIdentifierSecretArns"),
//   		},
//   		SecretsRoleArn: jsii.String("secretsRoleArn"),
//   	},
//   	ChannelGroupName: jsii.String("channelGroupName"),
//   	ChannelName: jsii.String("channelName"),
//   	OriginEndpointName: jsii.String("originEndpointName"),
//   	Policy: policy,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpointpolicy.html
//
type CfnOriginEndpointPolicyPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnOriginEndpointPolicyMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOriginEndpointPolicyPropsMixin
type jsiiProxy_CfnOriginEndpointPolicyPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnOriginEndpointPolicyPropsMixin) Props() *CfnOriginEndpointPolicyMixinProps {
	var returns *CfnOriginEndpointPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpointPolicyPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaPackageV2::OriginEndpointPolicy`.
func NewCfnOriginEndpointPolicyPropsMixin(props *CfnOriginEndpointPolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnOriginEndpointPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOriginEndpointPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOriginEndpointPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediapackagev2.CfnOriginEndpointPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaPackageV2::OriginEndpointPolicy`.
func NewCfnOriginEndpointPolicyPropsMixin_Override(c CfnOriginEndpointPolicyPropsMixin, props *CfnOriginEndpointPolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediapackagev2.CfnOriginEndpointPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnOriginEndpointPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOriginEndpointPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mediapackagev2.CfnOriginEndpointPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOriginEndpointPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_mediapackagev2.CfnOriginEndpointPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOriginEndpointPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnOriginEndpointPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

