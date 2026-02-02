package previewawsmedialivemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmedialive/previewawsmedialivemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::MediaLive::Input resource is a MediaLive resource type that creates an input.
//
// A MediaLive input holds information that describes how the MediaLive channel is connected to the upstream system that is providing the source content that is to be transcoded.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   cfnInputPropsMixin := awscdkmixinspreview.Mixins.NewCfnInputPropsMixin(&CfnInputMixinProps{
//   	Destinations: []interface{}{
//   		&InputDestinationRequestProperty{
//   			Network: jsii.String("network"),
//   			NetworkRoutes: []interface{}{
//   				&InputRequestDestinationRouteProperty{
//   					Cidr: jsii.String("cidr"),
//   					Gateway: jsii.String("gateway"),
//   				},
//   			},
//   			StaticIpAddress: jsii.String("staticIpAddress"),
//   			StreamName: jsii.String("streamName"),
//   		},
//   	},
//   	InputDevices: []interface{}{
//   		&InputDeviceSettingsProperty{
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	InputNetworkLocation: jsii.String("inputNetworkLocation"),
//   	InputSecurityGroups: []*string{
//   		jsii.String("inputSecurityGroups"),
//   	},
//   	MediaConnectFlows: []interface{}{
//   		&MediaConnectFlowRequestProperty{
//   			FlowArn: jsii.String("flowArn"),
//   		},
//   	},
//   	MulticastSettings: &MulticastSettingsCreateRequestProperty{
//   		Sources: []interface{}{
//   			&MulticastSourceCreateRequestProperty{
//   				SourceIp: jsii.String("sourceIp"),
//   				Url: jsii.String("url"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	RouterSettings: &RouterSettingsProperty{
//   		Destinations: []interface{}{
//   			&RouterDestinationSettingsProperty{
//   				AvailabilityZoneName: jsii.String("availabilityZoneName"),
//   			},
//   		},
//   		EncryptionType: jsii.String("encryptionType"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	SdiSources: []*string{
//   		jsii.String("sdiSources"),
//   	},
//   	Smpte2110ReceiverGroupSettings: &Smpte2110ReceiverGroupSettingsProperty{
//   		Smpte2110ReceiverGroups: []interface{}{
//   			&Smpte2110ReceiverGroupProperty{
//   				SdpSettings: &Smpte2110ReceiverGroupSdpSettingsProperty{
//   					AncillarySdps: []interface{}{
//   						&InputSdpLocationProperty{
//   							MediaIndex: jsii.Number(123),
//   							SdpUrl: jsii.String("sdpUrl"),
//   						},
//   					},
//   					AudioSdps: []interface{}{
//   						&InputSdpLocationProperty{
//   							MediaIndex: jsii.Number(123),
//   							SdpUrl: jsii.String("sdpUrl"),
//   						},
//   					},
//   					VideoSdp: &InputSdpLocationProperty{
//   						MediaIndex: jsii.Number(123),
//   						SdpUrl: jsii.String("sdpUrl"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Sources: []interface{}{
//   		&InputSourceRequestProperty{
//   			PasswordParam: jsii.String("passwordParam"),
//   			Url: jsii.String("url"),
//   			Username: jsii.String("username"),
//   		},
//   	},
//   	SrtSettings: &SrtSettingsRequestProperty{
//   		SrtCallerSources: []interface{}{
//   			&SrtCallerSourceRequestProperty{
//   				Decryption: &SrtCallerDecryptionRequestProperty{
//   					Algorithm: jsii.String("algorithm"),
//   					PassphraseSecretArn: jsii.String("passphraseSecretArn"),
//   				},
//   				MinimumLatency: jsii.Number(123),
//   				SrtListenerAddress: jsii.String("srtListenerAddress"),
//   				SrtListenerPort: jsii.String("srtListenerPort"),
//   				StreamId: jsii.String("streamId"),
//   			},
//   		},
//   	},
//   	Tags: tags,
//   	Type: jsii.String("type"),
//   	Vpc: &InputVpcRequestProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html
//
type CfnInputPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnInputMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInputPropsMixin
type jsiiProxy_CfnInputPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnInputPropsMixin) Props() *CfnInputMixinProps {
	var returns *CfnInputMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaLive::Input`.
func NewCfnInputPropsMixin(props *CfnInputMixinProps, options *mixins.CfnPropertyMixinOptions) CfnInputPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInputPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInputPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnInputPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaLive::Input`.
func NewCfnInputPropsMixin_Override(c CfnInputPropsMixin, props *CfnInputMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnInputPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnInputPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInputPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnInputPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInputPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnInputPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInputPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnInputPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

