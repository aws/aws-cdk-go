package mixinsawsmediaconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsmediaconnect/mixinsawsmediaconnect/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a router input in AWS Elemental MediaConnect that is used to ingest content to be transmitted to router outputs.
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
//   cfnRouterInputPropsMixin := awscdkmixinspreview.Mixins.NewCfnRouterInputPropsMixin(&CfnRouterInputMixinProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	Configuration: &RouterInputConfigurationProperty{
//   		Failover: &FailoverRouterInputConfigurationProperty{
//   			NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   			PrimarySourceIndex: jsii.Number(123),
//   			ProtocolConfigurations: []interface{}{
//   				&FailoverRouterInputProtocolConfigurationProperty{
//   					Rist: &RistRouterInputConfigurationProperty{
//   						Port: jsii.Number(123),
//   						RecoveryLatencyMilliseconds: jsii.Number(123),
//   					},
//   					Rtp: &RtpRouterInputConfigurationProperty{
//   						ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   						Port: jsii.Number(123),
//   					},
//   					SrtCaller: &SrtCallerRouterInputConfigurationProperty{
//   						DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   							EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   								RoleArn: jsii.String("roleArn"),
//   								SecretArn: jsii.String("secretArn"),
//   							},
//   						},
//   						MinimumLatencyMilliseconds: jsii.Number(123),
//   						SourceAddress: jsii.String("sourceAddress"),
//   						SourcePort: jsii.Number(123),
//   						StreamId: jsii.String("streamId"),
//   					},
//   					SrtListener: &SrtListenerRouterInputConfigurationProperty{
//   						DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   							EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   								RoleArn: jsii.String("roleArn"),
//   								SecretArn: jsii.String("secretArn"),
//   							},
//   						},
//   						MinimumLatencyMilliseconds: jsii.Number(123),
//   						Port: jsii.Number(123),
//   					},
//   				},
//   			},
//   			SourcePriorityMode: jsii.String("sourcePriorityMode"),
//   		},
//   		MediaConnectFlow: &MediaConnectFlowRouterInputConfigurationProperty{
//   			FlowArn: jsii.String("flowArn"),
//   			FlowOutputArn: jsii.String("flowOutputArn"),
//   			SourceTransitDecryption: &FlowTransitEncryptionProperty{
//   				EncryptionKeyConfiguration: &FlowTransitEncryptionKeyConfigurationProperty{
//   					Automatic: automatic,
//   					SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   				EncryptionKeyType: jsii.String("encryptionKeyType"),
//   			},
//   		},
//   		Merge: &MergeRouterInputConfigurationProperty{
//   			MergeRecoveryWindowMilliseconds: jsii.Number(123),
//   			NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   			ProtocolConfigurations: []interface{}{
//   				&MergeRouterInputProtocolConfigurationProperty{
//   					Rist: &RistRouterInputConfigurationProperty{
//   						Port: jsii.Number(123),
//   						RecoveryLatencyMilliseconds: jsii.Number(123),
//   					},
//   					Rtp: &RtpRouterInputConfigurationProperty{
//   						ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   						Port: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		Standard: &StandardRouterInputConfigurationProperty{
//   			NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   			Protocol: jsii.String("protocol"),
//   			ProtocolConfiguration: &RouterInputProtocolConfigurationProperty{
//   				Rist: &RistRouterInputConfigurationProperty{
//   					Port: jsii.Number(123),
//   					RecoveryLatencyMilliseconds: jsii.Number(123),
//   				},
//   				Rtp: &RtpRouterInputConfigurationProperty{
//   					ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   					Port: jsii.Number(123),
//   				},
//   				SrtCaller: &SrtCallerRouterInputConfigurationProperty{
//   					DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   						EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   							RoleArn: jsii.String("roleArn"),
//   							SecretArn: jsii.String("secretArn"),
//   						},
//   					},
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					SourceAddress: jsii.String("sourceAddress"),
//   					SourcePort: jsii.Number(123),
//   					StreamId: jsii.String("streamId"),
//   				},
//   				SrtListener: &SrtListenerRouterInputConfigurationProperty{
//   					DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
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
//   	TransitEncryption: &RouterInputTransitEncryptionProperty{
//   		EncryptionKeyConfiguration: &RouterInputTransitEncryptionKeyConfigurationProperty{
//   			Automatic: automatic,
//   			SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//   		EncryptionKeyType: jsii.String("encryptionKeyType"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html
//
type CfnRouterInputPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRouterInputMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRouterInputPropsMixin
type jsiiProxy_CfnRouterInputPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRouterInputPropsMixin) Props() *CfnRouterInputMixinProps {
	var returns *CfnRouterInputMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouterInputPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaConnect::RouterInput`.
func NewCfnRouterInputPropsMixin(props *CfnRouterInputMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRouterInputPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRouterInputPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRouterInputPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnRouterInputPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaConnect::RouterInput`.
func NewCfnRouterInputPropsMixin_Override(c CfnRouterInputPropsMixin, props *CfnRouterInputMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnRouterInputPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRouterInputPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRouterInputPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnRouterInputPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRouterInputPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnRouterInputPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRouterInputPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRouterInputPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

