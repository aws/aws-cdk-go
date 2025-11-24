package mixinsawsmediaconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsmediaconnect/mixinsawsmediaconnect/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a router input in AWS Elemental MediaConnect that can be used to egress content transmitted from router inputs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var automatic interface{}
//   var default_ interface{}
//
//   cfnRouterOutputPropsMixin := awscdkmixinspreview.Mixins.NewCfnRouterOutputPropsMixin(&CfnRouterOutputMixinProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	Configuration: &RouterOutputConfigurationProperty{
//   		MediaConnectFlow: &MediaConnectFlowRouterOutputConfigurationProperty{
//   			DestinationTransitEncryption: &FlowTransitEncryptionProperty{
//   				EncryptionKeyConfiguration: &FlowTransitEncryptionKeyConfigurationProperty{
//   					Automatic: automatic,
//   					SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   				EncryptionKeyType: jsii.String("encryptionKeyType"),
//   			},
//   			FlowArn: jsii.String("flowArn"),
//   			FlowSourceArn: jsii.String("flowSourceArn"),
//   		},
//   		MediaLiveInput: &MediaLiveInputRouterOutputConfigurationProperty{
//   			DestinationTransitEncryption: &MediaLiveTransitEncryptionProperty{
//   				EncryptionKeyConfiguration: &MediaLiveTransitEncryptionKeyConfigurationProperty{
//   					Automatic: automatic,
//   					SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   				EncryptionKeyType: jsii.String("encryptionKeyType"),
//   			},
//   			MediaLiveInputArn: jsii.String("mediaLiveInputArn"),
//   			MediaLivePipelineId: jsii.String("mediaLivePipelineId"),
//   		},
//   		Standard: &StandardRouterOutputConfigurationProperty{
//   			NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   			Protocol: jsii.String("protocol"),
//   			ProtocolConfiguration: &RouterOutputProtocolConfigurationProperty{
//   				Rist: &RistRouterOutputConfigurationProperty{
//   					DestinationAddress: jsii.String("destinationAddress"),
//   					DestinationPort: jsii.Number(123),
//   				},
//   				Rtp: &RtpRouterOutputConfigurationProperty{
//   					DestinationAddress: jsii.String("destinationAddress"),
//   					DestinationPort: jsii.Number(123),
//   					ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   				},
//   				SrtCaller: &SrtCallerRouterOutputConfigurationProperty{
//   					DestinationAddress: jsii.String("destinationAddress"),
//   					DestinationPort: jsii.Number(123),
//   					EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
//   						EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   							RoleArn: jsii.String("roleArn"),
//   							SecretArn: jsii.String("secretArn"),
//   						},
//   					},
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					StreamId: jsii.String("streamId"),
//   				},
//   				SrtListener: &SrtListenerRouterOutputConfigurationProperty{
//   					EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
//   						EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   							RoleArn: jsii.String("roleArn"),
//   							SecretArn: jsii.String("secretArn"),
//   						},
//   					},
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					Port: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	MaintenanceConfiguration: &MaintenanceConfigurationProperty{
//   		Default: default_,
//   		PreferredDayTime: &PreferredDayTimeMaintenanceConfigurationProperty{
//   			Day: jsii.String("day"),
//   			Time: jsii.String("time"),
//   		},
//   	},
//   	MaximumBitrate: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	RegionName: jsii.String("regionName"),
//   	RoutingScope: jsii.String("routingScope"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tier: jsii.String("tier"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html
//
type CfnRouterOutputPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRouterOutputMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRouterOutputPropsMixin
type jsiiProxy_CfnRouterOutputPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRouterOutputPropsMixin) Props() *CfnRouterOutputMixinProps {
	var returns *CfnRouterOutputMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouterOutputPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaConnect::RouterOutput`.
func NewCfnRouterOutputPropsMixin(props *CfnRouterOutputMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRouterOutputPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRouterOutputPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRouterOutputPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnRouterOutputPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaConnect::RouterOutput`.
func NewCfnRouterOutputPropsMixin_Override(c CfnRouterOutputPropsMixin, props *CfnRouterOutputMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnRouterOutputPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRouterOutputPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRouterOutputPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnRouterOutputPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRouterOutputPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnRouterOutputPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRouterOutputPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRouterOutputPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

