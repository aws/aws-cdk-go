package previewawsmediapackagemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmediapackage/previewawsmediapackagemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a packaging configuration in a packaging group.
//
// The packaging configuration represents a single delivery point for an asset. It determines the format and setting for the egressing content. Specify only one package format per configuration, such as `HlsPackage` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPackagingConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnPackagingConfigurationPropsMixin(&CfnPackagingConfigurationMixinProps{
//   	CmafPackage: &CmafPackageProperty{
//   		Encryption: &CmafEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   					PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
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
//   				IncludeIframeOnlyStream: jsii.Boolean(false),
//   				ManifestName: jsii.String("manifestName"),
//   				ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   				RepeatExtXKey: jsii.Boolean(false),
//   				StreamSelection: &StreamSelectionProperty{
//   					MaxVideoBitsPerSecond: jsii.Number(123),
//   					MinVideoBitsPerSecond: jsii.Number(123),
//   					StreamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//   		IncludeEncoderConfigurationInSegments: jsii.Boolean(false),
//   		SegmentDurationSeconds: jsii.Number(123),
//   	},
//   	DashPackage: &DashPackageProperty{
//   		DashManifests: []interface{}{
//   			&DashManifestProperty{
//   				ManifestLayout: jsii.String("manifestLayout"),
//   				ManifestName: jsii.String("manifestName"),
//   				MinBufferTimeSeconds: jsii.Number(123),
//   				Profile: jsii.String("profile"),
//   				ScteMarkersSource: jsii.String("scteMarkersSource"),
//   				StreamSelection: &StreamSelectionProperty{
//   					MaxVideoBitsPerSecond: jsii.Number(123),
//   					MinVideoBitsPerSecond: jsii.Number(123),
//   					StreamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//   		Encryption: &DashEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   					PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//   			},
//   		},
//   		IncludeEncoderConfigurationInSegments: jsii.Boolean(false),
//   		IncludeIframeOnlyStream: jsii.Boolean(false),
//   		PeriodTriggers: []*string{
//   			jsii.String("periodTriggers"),
//   		},
//   		SegmentDurationSeconds: jsii.Number(123),
//   		SegmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   	},
//   	HlsPackage: &HlsPackageProperty{
//   		Encryption: &HlsEncryptionProperty{
//   			ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   			EncryptionMethod: jsii.String("encryptionMethod"),
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   					PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
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
//   				IncludeIframeOnlyStream: jsii.Boolean(false),
//   				ManifestName: jsii.String("manifestName"),
//   				ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   				RepeatExtXKey: jsii.Boolean(false),
//   				StreamSelection: &StreamSelectionProperty{
//   					MaxVideoBitsPerSecond: jsii.Number(123),
//   					MinVideoBitsPerSecond: jsii.Number(123),
//   					StreamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//   		IncludeDvbSubtitles: jsii.Boolean(false),
//   		SegmentDurationSeconds: jsii.Number(123),
//   		UseAudioRenditionGroup: jsii.Boolean(false),
//   	},
//   	Id: jsii.String("id"),
//   	MssPackage: &MssPackageProperty{
//   		Encryption: &MssEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   					PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//   			},
//   		},
//   		MssManifests: []interface{}{
//   			&MssManifestProperty{
//   				ManifestName: jsii.String("manifestName"),
//   				StreamSelection: &StreamSelectionProperty{
//   					MaxVideoBitsPerSecond: jsii.Number(123),
//   					MinVideoBitsPerSecond: jsii.Number(123),
//   					StreamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//   		SegmentDurationSeconds: jsii.Number(123),
//   	},
//   	PackagingGroupId: jsii.String("packagingGroupId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html
//
type CfnPackagingConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPackagingConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPackagingConfigurationPropsMixin
type jsiiProxy_CfnPackagingConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPackagingConfigurationPropsMixin) Props() *CfnPackagingConfigurationMixinProps {
	var returns *CfnPackagingConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaPackage::PackagingConfiguration`.
func NewCfnPackagingConfigurationPropsMixin(props *CfnPackagingConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPackagingConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPackagingConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPackagingConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaPackage::PackagingConfiguration`.
func NewCfnPackagingConfigurationPropsMixin_Override(c CfnPackagingConfigurationPropsMixin, props *CfnPackagingConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPackagingConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPackagingConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPackagingConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPackagingConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPackagingConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

