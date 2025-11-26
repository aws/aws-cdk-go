package previewawsbedrockmixins


// The project's standard output configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   standardOutputConfigurationProperty := &StandardOutputConfigurationProperty{
//   	Audio: &AudioStandardOutputConfigurationProperty{
//   		Extraction: &AudioStandardExtractionProperty{
//   			Category: &AudioExtractionCategoryProperty{
//   				State: jsii.String("state"),
//   				TypeConfiguration: &AudioExtractionCategoryTypeConfigurationProperty{
//   					Transcript: &TranscriptConfigurationProperty{
//   						ChannelLabeling: &ChannelLabelingConfigurationProperty{
//   							State: jsii.String("state"),
//   						},
//   						SpeakerLabeling: &SpeakerLabelingConfigurationProperty{
//   							State: jsii.String("state"),
//   						},
//   					},
//   				},
//   				Types: []*string{
//   					jsii.String("types"),
//   				},
//   			},
//   		},
//   		GenerativeField: &AudioStandardGenerativeFieldProperty{
//   			State: jsii.String("state"),
//   			Types: []*string{
//   				jsii.String("types"),
//   			},
//   		},
//   	},
//   	Document: &DocumentStandardOutputConfigurationProperty{
//   		Extraction: &DocumentStandardExtractionProperty{
//   			BoundingBox: &DocumentBoundingBoxProperty{
//   				State: jsii.String("state"),
//   			},
//   			Granularity: &DocumentExtractionGranularityProperty{
//   				Types: []*string{
//   					jsii.String("types"),
//   				},
//   			},
//   		},
//   		GenerativeField: &DocumentStandardGenerativeFieldProperty{
//   			State: jsii.String("state"),
//   		},
//   		OutputFormat: &DocumentOutputFormatProperty{
//   			AdditionalFileFormat: &DocumentOutputAdditionalFileFormatProperty{
//   				State: jsii.String("state"),
//   			},
//   			TextFormat: &DocumentOutputTextFormatProperty{
//   				Types: []*string{
//   					jsii.String("types"),
//   				},
//   			},
//   		},
//   	},
//   	Image: &ImageStandardOutputConfigurationProperty{
//   		Extraction: &ImageStandardExtractionProperty{
//   			BoundingBox: &ImageBoundingBoxProperty{
//   				State: jsii.String("state"),
//   			},
//   			Category: &ImageExtractionCategoryProperty{
//   				State: jsii.String("state"),
//   				Types: []*string{
//   					jsii.String("types"),
//   				},
//   			},
//   		},
//   		GenerativeField: &ImageStandardGenerativeFieldProperty{
//   			State: jsii.String("state"),
//   			Types: []*string{
//   				jsii.String("types"),
//   			},
//   		},
//   	},
//   	Video: &VideoStandardOutputConfigurationProperty{
//   		Extraction: &VideoStandardExtractionProperty{
//   			BoundingBox: &VideoBoundingBoxProperty{
//   				State: jsii.String("state"),
//   			},
//   			Category: &VideoExtractionCategoryProperty{
//   				State: jsii.String("state"),
//   				Types: []*string{
//   					jsii.String("types"),
//   				},
//   			},
//   		},
//   		GenerativeField: &VideoStandardGenerativeFieldProperty{
//   			State: jsii.String("state"),
//   			Types: []*string{
//   				jsii.String("types"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-standardoutputconfiguration.html
//
type CfnDataAutomationProjectPropsMixin_StandardOutputConfigurationProperty struct {
	// Settings for processing audio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-standardoutputconfiguration.html#cfn-bedrock-dataautomationproject-standardoutputconfiguration-audio
	//
	Audio interface{} `field:"optional" json:"audio" yaml:"audio"`
	// Settings for processing documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-standardoutputconfiguration.html#cfn-bedrock-dataautomationproject-standardoutputconfiguration-document
	//
	Document interface{} `field:"optional" json:"document" yaml:"document"`
	// Settings for processing images.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-standardoutputconfiguration.html#cfn-bedrock-dataautomationproject-standardoutputconfiguration-image
	//
	Image interface{} `field:"optional" json:"image" yaml:"image"`
	// Settings for processing video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-standardoutputconfiguration.html#cfn-bedrock-dataautomationproject-standardoutputconfiguration-video
	//
	Video interface{} `field:"optional" json:"video" yaml:"video"`
}

