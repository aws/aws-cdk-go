package awsmediaconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a destination endpoint in AWS Elemental MediaConnect that can be connected to a router output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDestinationEndpointPropsMixin := awscdkcfnpropertymixins.Aws_mediaconnect.NewCfnDestinationEndpointPropsMixin(&CfnDestinationEndpointMixinProps{
//   	Configuration: &DestinationEndpointConfigurationProperty{
//   		RouterInput: &RouterInputDestinationEndpointConfigurationProperty{
//   			RouterInputArn: jsii.String("routerInputArn"),
//   		},
//   		Standard: &StandardDestinationEndpointConfigurationProperty{
//   			DestinationNetworkingConfiguration: &StandardDestinationEndpointNetworkingConfigurationProperty{
//   				Public: &PublicDestinationEndpointNetworkingConfigurationProperty{
//   					DestinationAllowRules: []interface{}{
//   						&PublicDestinationEndpointNetworkingConfigurationAllowRuleProperty{
//   							Cidr: jsii.String("cidr"),
//   						},
//   					},
//   				},
//   				Vpc: &VpcDestinationEndpointNetworkingConfigurationProperty{
//   					SecurityGroupIds: []*string{
//   						jsii.String("securityGroupIds"),
//   					},
//   					SubnetId: jsii.String("subnetId"),
//   				},
//   			},
//   			DestinationNetworkingType: jsii.String("destinationNetworkingType"),
//   			Name: jsii.String("name"),
//   			Protocol: jsii.String("protocol"),
//   			ProtocolConfiguration: &StandardDestinationEndpointProtocolConfigurationProperty{
//   				Rist: &RistStandardDestinationEndpointConfigurationProperty{
//   					DestinationAddress: jsii.String("destinationAddress"),
//   					DestinationIngestPort: jsii.Number(123),
//   				},
//   				Rtp: &RtpStandardDestinationEndpointConfigurationProperty{
//   					DestinationAddress: jsii.String("destinationAddress"),
//   					DestinationIngestPort: jsii.Number(123),
//   					ForwardErrorCorrectionState: jsii.String("forwardErrorCorrectionState"),
//   				},
//   				SrtCaller: &SrtCallerStandardDestinationEndpointConfigurationProperty{
//   					EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					SourceListenerPort: jsii.Number(123),
//   				},
//   				SrtListener: &SrtListenerStandardDestinationEndpointConfigurationProperty{
//   					DestinationAddress: jsii.String("destinationAddress"),
//   					DestinationIngestPort: jsii.Number(123),
//   					EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					StreamId: jsii.String("streamId"),
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	RegionName: jsii.String("regionName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-destinationendpoint.html
//
type CfnDestinationEndpointPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDestinationEndpointMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDestinationEndpointPropsMixin
type jsiiProxy_CfnDestinationEndpointPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDestinationEndpointPropsMixin) Props() *CfnDestinationEndpointMixinProps {
	var returns *CfnDestinationEndpointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDestinationEndpointPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaConnect::DestinationEndpoint`.
func NewCfnDestinationEndpointPropsMixin(props *CfnDestinationEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDestinationEndpointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDestinationEndpointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDestinationEndpointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaConnect::DestinationEndpoint`.
func NewCfnDestinationEndpointPropsMixin_Override(c CfnDestinationEndpointPropsMixin, props *CfnDestinationEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDestinationEndpointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDestinationEndpointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDestinationEndpointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDestinationEndpointPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDestinationEndpointPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

