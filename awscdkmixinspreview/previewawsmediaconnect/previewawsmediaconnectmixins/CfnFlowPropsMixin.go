package previewawsmediaconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmediaconnect/previewawsmediaconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::MediaConnect::Flow` resource defines a connection between one or more video sources and one or more outputs.
//
// For each flow, you specify the transport protocol to use, encryption information, and details for any outputs or entitlements that you want. AWS Elemental MediaConnect returns an ingest endpoint where you can send your live video as a single unicast stream. The service replicates and distributes the video to every output that you specify, whether inside or outside the AWS Cloud. You can also set up entitlements on a flow to allow other AWS accounts to access your content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var automatic interface{}
//
//   cfnFlowPropsMixin := awscdkmixinspreview.Mixins.NewCfnFlowPropsMixin(&CfnFlowMixinProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	FlowSize: jsii.String("flowSize"),
//   	Maintenance: &MaintenanceProperty{
//   		MaintenanceDay: jsii.String("maintenanceDay"),
//   		MaintenanceStartHour: jsii.String("maintenanceStartHour"),
//   	},
//   	MediaStreams: []interface{}{
//   		&MediaStreamProperty{
//   			Attributes: &MediaStreamAttributesProperty{
//   				Fmtp: &FmtpProperty{
//   					ChannelOrder: jsii.String("channelOrder"),
//   					Colorimetry: jsii.String("colorimetry"),
//   					ExactFramerate: jsii.String("exactFramerate"),
//   					Par: jsii.String("par"),
//   					Range: jsii.String("range"),
//   					ScanMode: jsii.String("scanMode"),
//   					Tcs: jsii.String("tcs"),
//   				},
//   				Lang: jsii.String("lang"),
//   			},
//   			ClockRate: jsii.Number(123),
//   			Description: jsii.String("description"),
//   			Fmt: jsii.Number(123),
//   			MediaStreamId: jsii.Number(123),
//   			MediaStreamName: jsii.String("mediaStreamName"),
//   			MediaStreamType: jsii.String("mediaStreamType"),
//   			VideoFormat: jsii.String("videoFormat"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	NdiConfig: &NdiConfigProperty{
//   		MachineName: jsii.String("machineName"),
//   		NdiDiscoveryServers: []interface{}{
//   			&NdiDiscoveryServerConfigProperty{
//   				DiscoveryServerAddress: jsii.String("discoveryServerAddress"),
//   				DiscoveryServerPort: jsii.Number(123),
//   				VpcInterfaceAdapter: jsii.String("vpcInterfaceAdapter"),
//   			},
//   		},
//   		NdiState: jsii.String("ndiState"),
//   	},
//   	Source: &SourceProperty{
//   		Decryption: &EncryptionProperty{
//   			Algorithm: jsii.String("algorithm"),
//   			ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   			DeviceId: jsii.String("deviceId"),
//   			KeyType: jsii.String("keyType"),
//   			Region: jsii.String("region"),
//   			ResourceId: jsii.String("resourceId"),
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   			Url: jsii.String("url"),
//   		},
//   		Description: jsii.String("description"),
//   		EntitlementArn: jsii.String("entitlementArn"),
//   		GatewayBridgeSource: &GatewayBridgeSourceProperty{
//   			BridgeArn: jsii.String("bridgeArn"),
//   			VpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   				VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   			},
//   		},
//   		IngestIp: jsii.String("ingestIp"),
//   		IngestPort: jsii.Number(123),
//   		MaxBitrate: jsii.Number(123),
//   		MaxLatency: jsii.Number(123),
//   		MaxSyncBuffer: jsii.Number(123),
//   		MediaStreamSourceConfigurations: []interface{}{
//   			&MediaStreamSourceConfigurationProperty{
//   				EncodingName: jsii.String("encodingName"),
//   				InputConfigurations: []interface{}{
//   					&InputConfigurationProperty{
//   						InputPort: jsii.Number(123),
//   						Interface: &InterfaceProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   				MediaStreamName: jsii.String("mediaStreamName"),
//   			},
//   		},
//   		MinLatency: jsii.Number(123),
//   		Name: jsii.String("name"),
//   		Protocol: jsii.String("protocol"),
//   		RouterIntegrationState: jsii.String("routerIntegrationState"),
//   		RouterIntegrationTransitDecryption: &FlowTransitEncryptionProperty{
//   			EncryptionKeyConfiguration: &FlowTransitEncryptionKeyConfigurationProperty{
//   				Automatic: automatic,
//   				SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   			},
//   			EncryptionKeyType: jsii.String("encryptionKeyType"),
//   		},
//   		SenderControlPort: jsii.Number(123),
//   		SenderIpAddress: jsii.String("senderIpAddress"),
//   		SourceArn: jsii.String("sourceArn"),
//   		SourceIngestPort: jsii.String("sourceIngestPort"),
//   		SourceListenerAddress: jsii.String("sourceListenerAddress"),
//   		SourceListenerPort: jsii.Number(123),
//   		StreamId: jsii.String("streamId"),
//   		VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   		WhitelistCidr: jsii.String("whitelistCidr"),
//   	},
//   	SourceFailoverConfig: &FailoverConfigProperty{
//   		FailoverMode: jsii.String("failoverMode"),
//   		RecoveryWindow: jsii.Number(123),
//   		SourcePriority: &SourcePriorityProperty{
//   			PrimarySource: jsii.String("primarySource"),
//   		},
//   		State: jsii.String("state"),
//   	},
//   	SourceMonitoringConfig: &SourceMonitoringConfigProperty{
//   		AudioMonitoringSettings: []interface{}{
//   			&AudioMonitoringSettingProperty{
//   				SilentAudio: &SilentAudioProperty{
//   					State: jsii.String("state"),
//   					ThresholdSeconds: jsii.Number(123),
//   				},
//   			},
//   		},
//   		ContentQualityAnalysisState: jsii.String("contentQualityAnalysisState"),
//   		ThumbnailState: jsii.String("thumbnailState"),
//   		VideoMonitoringSettings: []interface{}{
//   			&VideoMonitoringSettingProperty{
//   				BlackFrames: &BlackFramesProperty{
//   					State: jsii.String("state"),
//   					ThresholdSeconds: jsii.Number(123),
//   				},
//   				FrozenFrames: &FrozenFramesProperty{
//   					State: jsii.String("state"),
//   					ThresholdSeconds: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	VpcInterfaces: []interface{}{
//   		&VpcInterfaceProperty{
//   			Name: jsii.String("name"),
//   			NetworkInterfaceIds: []*string{
//   				jsii.String("networkInterfaceIds"),
//   			},
//   			NetworkInterfaceType: jsii.String("networkInterfaceType"),
//   			RoleArn: jsii.String("roleArn"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html
//
type CfnFlowPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFlowMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFlowPropsMixin
type jsiiProxy_CfnFlowPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFlowPropsMixin) Props() *CfnFlowMixinProps {
	var returns *CfnFlowMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaConnect::Flow`.
func NewCfnFlowPropsMixin(props *CfnFlowMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFlowPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFlowPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFlowPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnFlowPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaConnect::Flow`.
func NewCfnFlowPropsMixin_Override(c CfnFlowPropsMixin, props *CfnFlowMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnFlowPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFlowPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlowPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnFlowPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlowPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnFlowPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFlowPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFlowPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

