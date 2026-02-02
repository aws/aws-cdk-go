package previewawsmediapackagemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmediapackage/previewawsmediapackagemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create an endpoint on an AWS Elemental MediaPackage channel.
//
// An endpoint represents a single delivery point of a channel, and defines content output handling through various components, such as packaging protocols, DRM and encryption integration, and more.
//
// After it's created, an endpoint provides a fixed public URL. This URL remains the same throughout the lifetime of the endpoint, regardless of any failures or upgrades that might occur. Integrate the URL with a downstream CDN (such as Amazon CloudFront) or playback device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOriginEndpointPropsMixin := awscdkmixinspreview.Mixins.NewCfnOriginEndpointPropsMixin(&CfnOriginEndpointMixinProps{
//   	Authorization: &AuthorizationProperty{
//   		CdnIdentifierSecret: jsii.String("cdnIdentifierSecret"),
//   		SecretsRoleArn: jsii.String("secretsRoleArn"),
//   	},
//   	ChannelId: jsii.String("channelId"),
//   	CmafPackage: &CmafPackageProperty{
//   		Encryption: &CmafEncryptionProperty{
//   			ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   			EncryptionMethod: jsii.String("encryptionMethod"),
//   			KeyRotationIntervalSeconds: jsii.Number(123),
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				CertificateArn: jsii.String("certificateArn"),
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   					PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   				ResourceId: jsii.String("resourceId"),
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//   			},
//   		},
//   		HlsManifests: []interface{}{
//   			&HlsManifestProperty{
//   				AdMarkers: jsii.String("adMarkers"),
//   				AdsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   				AdTriggers: []*string{
//   					jsii.String("adTriggers"),
//   				},
//   				Id: jsii.String("id"),
//   				IncludeIframeOnlyStream: jsii.Boolean(false),
//   				ManifestName: jsii.String("manifestName"),
//   				PlaylistType: jsii.String("playlistType"),
//   				PlaylistWindowSeconds: jsii.Number(123),
//   				ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   				Url: jsii.String("url"),
//   			},
//   		},
//   		SegmentDurationSeconds: jsii.Number(123),
//   		SegmentPrefix: jsii.String("segmentPrefix"),
//   		StreamSelection: &StreamSelectionProperty{
//   			MaxVideoBitsPerSecond: jsii.Number(123),
//   			MinVideoBitsPerSecond: jsii.Number(123),
//   			StreamOrder: jsii.String("streamOrder"),
//   		},
//   	},
//   	DashPackage: &DashPackageProperty{
//   		AdsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   		AdTriggers: []*string{
//   			jsii.String("adTriggers"),
//   		},
//   		Encryption: &DashEncryptionProperty{
//   			KeyRotationIntervalSeconds: jsii.Number(123),
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				CertificateArn: jsii.String("certificateArn"),
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   					PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   				ResourceId: jsii.String("resourceId"),
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//   			},
//   		},
//   		IncludeIframeOnlyStream: jsii.Boolean(false),
//   		ManifestLayout: jsii.String("manifestLayout"),
//   		ManifestWindowSeconds: jsii.Number(123),
//   		MinBufferTimeSeconds: jsii.Number(123),
//   		MinUpdatePeriodSeconds: jsii.Number(123),
//   		PeriodTriggers: []*string{
//   			jsii.String("periodTriggers"),
//   		},
//   		Profile: jsii.String("profile"),
//   		SegmentDurationSeconds: jsii.Number(123),
//   		SegmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   		StreamSelection: &StreamSelectionProperty{
//   			MaxVideoBitsPerSecond: jsii.Number(123),
//   			MinVideoBitsPerSecond: jsii.Number(123),
//   			StreamOrder: jsii.String("streamOrder"),
//   		},
//   		SuggestedPresentationDelaySeconds: jsii.Number(123),
//   		UtcTiming: jsii.String("utcTiming"),
//   		UtcTimingUri: jsii.String("utcTimingUri"),
//   	},
//   	Description: jsii.String("description"),
//   	HlsPackage: &HlsPackageProperty{
//   		AdMarkers: jsii.String("adMarkers"),
//   		AdsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   		AdTriggers: []*string{
//   			jsii.String("adTriggers"),
//   		},
//   		Encryption: &HlsEncryptionProperty{
//   			ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   			EncryptionMethod: jsii.String("encryptionMethod"),
//   			KeyRotationIntervalSeconds: jsii.Number(123),
//   			RepeatExtXKey: jsii.Boolean(false),
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				CertificateArn: jsii.String("certificateArn"),
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   					PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   				ResourceId: jsii.String("resourceId"),
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//   			},
//   		},
//   		IncludeDvbSubtitles: jsii.Boolean(false),
//   		IncludeIframeOnlyStream: jsii.Boolean(false),
//   		PlaylistType: jsii.String("playlistType"),
//   		PlaylistWindowSeconds: jsii.Number(123),
//   		ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   		SegmentDurationSeconds: jsii.Number(123),
//   		StreamSelection: &StreamSelectionProperty{
//   			MaxVideoBitsPerSecond: jsii.Number(123),
//   			MinVideoBitsPerSecond: jsii.Number(123),
//   			StreamOrder: jsii.String("streamOrder"),
//   		},
//   		UseAudioRenditionGroup: jsii.Boolean(false),
//   	},
//   	Id: jsii.String("id"),
//   	ManifestName: jsii.String("manifestName"),
//   	MssPackage: &MssPackageProperty{
//   		Encryption: &MssEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				CertificateArn: jsii.String("certificateArn"),
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   					PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   				ResourceId: jsii.String("resourceId"),
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//   			},
//   		},
//   		ManifestWindowSeconds: jsii.Number(123),
//   		SegmentDurationSeconds: jsii.Number(123),
//   		StreamSelection: &StreamSelectionProperty{
//   			MaxVideoBitsPerSecond: jsii.Number(123),
//   			MinVideoBitsPerSecond: jsii.Number(123),
//   			StreamOrder: jsii.String("streamOrder"),
//   		},
//   	},
//   	Origination: jsii.String("origination"),
//   	StartoverWindowSeconds: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeDelaySeconds: jsii.Number(123),
//   	Whitelist: []*string{
//   		jsii.String("whitelist"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-originendpoint.html
//
type CfnOriginEndpointPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnOriginEndpointMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOriginEndpointPropsMixin
type jsiiProxy_CfnOriginEndpointPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnOriginEndpointPropsMixin) Props() *CfnOriginEndpointMixinProps {
	var returns *CfnOriginEndpointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpointPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaPackage::OriginEndpoint`.
func NewCfnOriginEndpointPropsMixin(props *CfnOriginEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) CfnOriginEndpointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOriginEndpointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOriginEndpointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaPackage::OriginEndpoint`.
func NewCfnOriginEndpointPropsMixin_Override(c CfnOriginEndpointPropsMixin, props *CfnOriginEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnOriginEndpointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOriginEndpointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOriginEndpointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOriginEndpointPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnOriginEndpointPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

