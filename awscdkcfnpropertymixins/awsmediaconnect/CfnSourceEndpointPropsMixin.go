package awsmediaconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a source endpoint in AWS Elemental MediaConnect that can be connected to an input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnSourceEndpointPropsMixin := awscdkcfnpropertymixins.Aws_mediaconnect.NewCfnSourceEndpointPropsMixin(&CfnSourceEndpointMixinProps{
//   	Configuration: &SourceEndpointConfigurationProperty{
//   		RouterOutput: &RouterOutputSourceEndpointConfigurationProperty{
//   			RouterOutputArn: jsii.String("routerOutputArn"),
//   		},
//   		Standard: &StandardSourceEndpointConfigurationProperty{
//   			Name: jsii.String("name"),
//   			Protocol: jsii.String("protocol"),
//   			ProtocolConfiguration: &StandardSourceEndpointProtocolConfigurationProperty{
//   				Rist: &RistStandardSourceEndpointConfigurationProperty{
//   					DestinationIngestPort: jsii.Number(123),
//   					RecoveryLatencyMilliseconds: jsii.Number(123),
//   				},
//   				Rtp: &RtpStandardSourceEndpointConfigurationProperty{
//   					DestinationIngestPort: jsii.Number(123),
//   					ForwardErrorCorrectionState: jsii.String("forwardErrorCorrectionState"),
//   				},
//   				SrtCaller: &SrtCallerStandardSourceEndpointConfigurationProperty{
//   					DestinationIngestPort: jsii.Number(123),
//   					EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   				},
//   				SrtListener: &SrtListenerStandardSourceEndpointConfigurationProperty{
//   					EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					SourceAddress: jsii.String("sourceAddress"),
//   					SourceListenerPort: jsii.Number(123),
//   					StreamId: jsii.String("streamId"),
//   				},
//   			},
//   			SourceNetworkingConfiguration: &StandardSourceEndpointNetworkingConfigurationProperty{
//   				Public: &PublicSourceEndpointNetworkingConfigurationProperty{
//   					SourceAllowRules: []interface{}{
//   						&PublicSourceEndpointNetworkingConfigurationAllowRuleProperty{
//   							Cidr: jsii.String("cidr"),
//   						},
//   					},
//   				},
//   				Vpc: &VpcSourceEndpointNetworkingConfigurationProperty{
//   					SecurityGroupIds: []*string{
//   						jsii.String("securityGroupIds"),
//   					},
//   					SubnetId: jsii.String("subnetId"),
//   				},
//   			},
//   			SourceNetworkingType: jsii.String("sourceNetworkingType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-sourceendpoint.html
//
type CfnSourceEndpointPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnSourceEndpointMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSourceEndpointPropsMixin
type jsiiProxy_CfnSourceEndpointPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnSourceEndpointPropsMixin) Props() *CfnSourceEndpointMixinProps {
	var returns *CfnSourceEndpointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceEndpointPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaConnect::SourceEndpoint`.
func NewCfnSourceEndpointPropsMixin(props *CfnSourceEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnSourceEndpointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSourceEndpointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSourceEndpointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaConnect::SourceEndpoint`.
func NewCfnSourceEndpointPropsMixin_Override(c CfnSourceEndpointPropsMixin, props *CfnSourceEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnSourceEndpointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSourceEndpointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSourceEndpointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSourceEndpointPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSourceEndpointPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

