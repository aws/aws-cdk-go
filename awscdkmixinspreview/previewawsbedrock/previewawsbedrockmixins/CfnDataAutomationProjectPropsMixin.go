package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbedrock/previewawsbedrockmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A data automation project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataAutomationProjectPropsMixin := awscdkmixinspreview.Mixins.NewCfnDataAutomationProjectPropsMixin(&CfnDataAutomationProjectMixinProps{
//   	CustomOutputConfiguration: &CustomOutputConfigurationProperty{
//   		Blueprints: []interface{}{
//   			&BlueprintItemProperty{
//   				BlueprintArn: jsii.String("blueprintArn"),
//   				BlueprintStage: jsii.String("blueprintStage"),
//   				BlueprintVersion: jsii.String("blueprintVersion"),
//   			},
//   		},
//   	},
//   	KmsEncryptionContext: map[string]*string{
//   		"kmsEncryptionContextKey": jsii.String("kmsEncryptionContext"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	OverrideConfiguration: &OverrideConfigurationProperty{
//   		Audio: &AudioOverrideConfigurationProperty{
//   			LanguageConfiguration: &AudioLanguageConfigurationProperty{
//   				GenerativeOutputLanguage: jsii.String("generativeOutputLanguage"),
//   				IdentifyMultipleLanguages: jsii.Boolean(false),
//   				InputLanguages: []*string{
//   					jsii.String("inputLanguages"),
//   				},
//   			},
//   			ModalityProcessing: &ModalityProcessingConfigurationProperty{
//   				State: jsii.String("state"),
//   			},
//   			SensitiveDataConfiguration: &SensitiveDataConfigurationProperty{
//   				DetectionMode: jsii.String("detectionMode"),
//   				DetectionScope: []*string{
//   					jsii.String("detectionScope"),
//   				},
//   				PiiEntitiesConfiguration: &PIIEntitiesConfigurationProperty{
//   					PiiEntityTypes: []*string{
//   						jsii.String("piiEntityTypes"),
//   					},
//   					RedactionMaskMode: jsii.String("redactionMaskMode"),
//   				},
//   			},
//   		},
//   		Document: &DocumentOverrideConfigurationProperty{
//   			ModalityProcessing: &ModalityProcessingConfigurationProperty{
//   				State: jsii.String("state"),
//   			},
//   			SensitiveDataConfiguration: &SensitiveDataConfigurationProperty{
//   				DetectionMode: jsii.String("detectionMode"),
//   				DetectionScope: []*string{
//   					jsii.String("detectionScope"),
//   				},
//   				PiiEntitiesConfiguration: &PIIEntitiesConfigurationProperty{
//   					PiiEntityTypes: []*string{
//   						jsii.String("piiEntityTypes"),
//   					},
//   					RedactionMaskMode: jsii.String("redactionMaskMode"),
//   				},
//   			},
//   			Splitter: &SplitterConfigurationProperty{
//   				State: jsii.String("state"),
//   			},
//   		},
//   		Image: &ImageOverrideConfigurationProperty{
//   			ModalityProcessing: &ModalityProcessingConfigurationProperty{
//   				State: jsii.String("state"),
//   			},
//   			SensitiveDataConfiguration: &SensitiveDataConfigurationProperty{
//   				DetectionMode: jsii.String("detectionMode"),
//   				DetectionScope: []*string{
//   					jsii.String("detectionScope"),
//   				},
//   				PiiEntitiesConfiguration: &PIIEntitiesConfigurationProperty{
//   					PiiEntityTypes: []*string{
//   						jsii.String("piiEntityTypes"),
//   					},
//   					RedactionMaskMode: jsii.String("redactionMaskMode"),
//   				},
//   			},
//   		},
//   		ModalityRouting: &ModalityRoutingConfigurationProperty{
//   			Jpeg: jsii.String("jpeg"),
//   			Mov: jsii.String("mov"),
//   			Mp4: jsii.String("mp4"),
//   			Png: jsii.String("png"),
//   		},
//   		Video: &VideoOverrideConfigurationProperty{
//   			ModalityProcessing: &ModalityProcessingConfigurationProperty{
//   				State: jsii.String("state"),
//   			},
//   			SensitiveDataConfiguration: &SensitiveDataConfigurationProperty{
//   				DetectionMode: jsii.String("detectionMode"),
//   				DetectionScope: []*string{
//   					jsii.String("detectionScope"),
//   				},
//   				PiiEntitiesConfiguration: &PIIEntitiesConfigurationProperty{
//   					PiiEntityTypes: []*string{
//   						jsii.String("piiEntityTypes"),
//   					},
//   					RedactionMaskMode: jsii.String("redactionMaskMode"),
//   				},
//   			},
//   		},
//   	},
//   	ProjectDescription: jsii.String("projectDescription"),
//   	ProjectName: jsii.String("projectName"),
//   	ProjectType: jsii.String("projectType"),
//   	StandardOutputConfiguration: &StandardOutputConfigurationProperty{
//   		Audio: &AudioStandardOutputConfigurationProperty{
//   			Extraction: &AudioStandardExtractionProperty{
//   				Category: &AudioExtractionCategoryProperty{
//   					State: jsii.String("state"),
//   					TypeConfiguration: &AudioExtractionCategoryTypeConfigurationProperty{
//   						Transcript: &TranscriptConfigurationProperty{
//   							ChannelLabeling: &ChannelLabelingConfigurationProperty{
//   								State: jsii.String("state"),
//   							},
//   							SpeakerLabeling: &SpeakerLabelingConfigurationProperty{
//   								State: jsii.String("state"),
//   							},
//   						},
//   					},
//   					Types: []*string{
//   						jsii.String("types"),
//   					},
//   				},
//   			},
//   			GenerativeField: &AudioStandardGenerativeFieldProperty{
//   				State: jsii.String("state"),
//   				Types: []*string{
//   					jsii.String("types"),
//   				},
//   			},
//   		},
//   		Document: &DocumentStandardOutputConfigurationProperty{
//   			Extraction: &DocumentStandardExtractionProperty{
//   				BoundingBox: &DocumentBoundingBoxProperty{
//   					State: jsii.String("state"),
//   				},
//   				Granularity: &DocumentExtractionGranularityProperty{
//   					Types: []*string{
//   						jsii.String("types"),
//   					},
//   				},
//   			},
//   			GenerativeField: &DocumentStandardGenerativeFieldProperty{
//   				State: jsii.String("state"),
//   			},
//   			OutputFormat: &DocumentOutputFormatProperty{
//   				AdditionalFileFormat: &DocumentOutputAdditionalFileFormatProperty{
//   					State: jsii.String("state"),
//   				},
//   				TextFormat: &DocumentOutputTextFormatProperty{
//   					Types: []*string{
//   						jsii.String("types"),
//   					},
//   				},
//   			},
//   		},
//   		Image: &ImageStandardOutputConfigurationProperty{
//   			Extraction: &ImageStandardExtractionProperty{
//   				BoundingBox: &ImageBoundingBoxProperty{
//   					State: jsii.String("state"),
//   				},
//   				Category: &ImageExtractionCategoryProperty{
//   					State: jsii.String("state"),
//   					Types: []*string{
//   						jsii.String("types"),
//   					},
//   				},
//   			},
//   			GenerativeField: &ImageStandardGenerativeFieldProperty{
//   				State: jsii.String("state"),
//   				Types: []*string{
//   					jsii.String("types"),
//   				},
//   			},
//   		},
//   		Video: &VideoStandardOutputConfigurationProperty{
//   			Extraction: &VideoStandardExtractionProperty{
//   				BoundingBox: &VideoBoundingBoxProperty{
//   					State: jsii.String("state"),
//   				},
//   				Category: &VideoExtractionCategoryProperty{
//   					State: jsii.String("state"),
//   					Types: []*string{
//   						jsii.String("types"),
//   					},
//   				},
//   			},
//   			GenerativeField: &VideoStandardGenerativeFieldProperty{
//   				State: jsii.String("state"),
//   				Types: []*string{
//   					jsii.String("types"),
//   				},
//   			},
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationproject.html
//
type CfnDataAutomationProjectPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDataAutomationProjectMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataAutomationProjectPropsMixin
type jsiiProxy_CfnDataAutomationProjectPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDataAutomationProjectPropsMixin) Props() *CfnDataAutomationProjectMixinProps {
	var returns *CfnDataAutomationProjectMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataAutomationProjectPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Bedrock::DataAutomationProject`.
func NewCfnDataAutomationProjectPropsMixin(props *CfnDataAutomationProjectMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDataAutomationProjectPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataAutomationProjectPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataAutomationProjectPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnDataAutomationProjectPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Bedrock::DataAutomationProject`.
func NewCfnDataAutomationProjectPropsMixin_Override(c CfnDataAutomationProjectPropsMixin, props *CfnDataAutomationProjectMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnDataAutomationProjectPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDataAutomationProjectPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataAutomationProjectPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnDataAutomationProjectPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataAutomationProjectPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnDataAutomationProjectPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataAutomationProjectPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDataAutomationProjectPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

