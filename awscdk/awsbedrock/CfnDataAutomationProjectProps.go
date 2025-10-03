package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataAutomationProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataAutomationProjectProps := &CfnDataAutomationProjectProps{
//   	ProjectName: jsii.String("projectName"),
//
//   	// the properties below are optional
//   	CustomOutputConfiguration: &CustomOutputConfigurationProperty{
//   		Blueprints: []interface{}{
//   			&BlueprintItemProperty{
//   				BlueprintArn: jsii.String("blueprintArn"),
//
//   				// the properties below are optional
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
//   			ModalityProcessing: &ModalityProcessingConfigurationProperty{
//   				State: jsii.String("state"),
//   			},
//   		},
//   		Document: &DocumentOverrideConfigurationProperty{
//   			ModalityProcessing: &ModalityProcessingConfigurationProperty{
//   				State: jsii.String("state"),
//   			},
//   			Splitter: &SplitterConfigurationProperty{
//   				State: jsii.String("state"),
//   			},
//   		},
//   		Image: &ImageOverrideConfigurationProperty{
//   			ModalityProcessing: &ModalityProcessingConfigurationProperty{
//   				State: jsii.String("state"),
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
//   		},
//   	},
//   	ProjectDescription: jsii.String("projectDescription"),
//   	StandardOutputConfiguration: &StandardOutputConfigurationProperty{
//   		Audio: &AudioStandardOutputConfigurationProperty{
//   			Extraction: &AudioStandardExtractionProperty{
//   				Category: &AudioExtractionCategoryProperty{
//   					State: jsii.String("state"),
//
//   					// the properties below are optional
//   					Types: []*string{
//   						jsii.String("types"),
//   					},
//   				},
//   			},
//   			GenerativeField: &AudioStandardGenerativeFieldProperty{
//   				State: jsii.String("state"),
//
//   				// the properties below are optional
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
//
//   					// the properties below are optional
//   					Types: []*string{
//   						jsii.String("types"),
//   					},
//   				},
//   			},
//   			GenerativeField: &ImageStandardGenerativeFieldProperty{
//   				State: jsii.String("state"),
//
//   				// the properties below are optional
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
//
//   					// the properties below are optional
//   					Types: []*string{
//   						jsii.String("types"),
//   					},
//   				},
//   			},
//   			GenerativeField: &VideoStandardGenerativeFieldProperty{
//   				State: jsii.String("state"),
//
//   				// the properties below are optional
//   				Types: []*string{
//   					jsii.String("types"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationproject.html
//
type CfnDataAutomationProjectProps struct {
	// The project's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationproject.html#cfn-bedrock-dataautomationproject-projectname
	//
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// Blueprints to apply to objects processed by the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationproject.html#cfn-bedrock-dataautomationproject-customoutputconfiguration
	//
	CustomOutputConfiguration interface{} `field:"optional" json:"customOutputConfiguration" yaml:"customOutputConfiguration"`
	// The AWS KMS encryption context to use for encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationproject.html#cfn-bedrock-dataautomationproject-kmsencryptioncontext
	//
	KmsEncryptionContext interface{} `field:"optional" json:"kmsEncryptionContext" yaml:"kmsEncryptionContext"`
	// The AWS KMS key to use for encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationproject.html#cfn-bedrock-dataautomationproject-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Additional settings for the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationproject.html#cfn-bedrock-dataautomationproject-overrideconfiguration
	//
	OverrideConfiguration interface{} `field:"optional" json:"overrideConfiguration" yaml:"overrideConfiguration"`
	// The project's description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationproject.html#cfn-bedrock-dataautomationproject-projectdescription
	//
	ProjectDescription *string `field:"optional" json:"projectDescription" yaml:"projectDescription"`
	// The project's standard output configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationproject.html#cfn-bedrock-dataautomationproject-standardoutputconfiguration
	//
	StandardOutputConfiguration interface{} `field:"optional" json:"standardOutputConfiguration" yaml:"standardOutputConfiguration"`
	// List of Tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationproject.html#cfn-bedrock-dataautomationproject-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

