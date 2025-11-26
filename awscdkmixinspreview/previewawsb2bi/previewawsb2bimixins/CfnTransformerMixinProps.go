package previewawsb2bimixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTransformerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransformerMixinProps := &CfnTransformerMixinProps{
//   	EdiType: &EdiTypeProperty{
//   		X12Details: &X12DetailsProperty{
//   			TransactionSet: jsii.String("transactionSet"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	FileFormat: jsii.String("fileFormat"),
//   	InputConversion: &InputConversionProperty{
//   		AdvancedOptions: &AdvancedOptionsProperty{
//   			X12: &X12AdvancedOptionsProperty{
//   				SplitOptions: &X12SplitOptionsProperty{
//   					SplitBy: jsii.String("splitBy"),
//   				},
//   				ValidationOptions: &X12ValidationOptionsProperty{
//   					ValidationRules: []interface{}{
//   						&X12ValidationRuleProperty{
//   							CodeListValidationRule: &X12CodeListValidationRuleProperty{
//   								CodesToAdd: []*string{
//   									jsii.String("codesToAdd"),
//   								},
//   								CodesToRemove: []*string{
//   									jsii.String("codesToRemove"),
//   								},
//   								ElementId: jsii.String("elementId"),
//   							},
//   							ElementLengthValidationRule: &X12ElementLengthValidationRuleProperty{
//   								ElementId: jsii.String("elementId"),
//   								MaxLength: jsii.Number(123),
//   								MinLength: jsii.Number(123),
//   							},
//   							ElementRequirementValidationRule: &X12ElementRequirementValidationRuleProperty{
//   								ElementPosition: jsii.String("elementPosition"),
//   								Requirement: jsii.String("requirement"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		FormatOptions: &FormatOptionsProperty{
//   			X12: &X12DetailsProperty{
//   				TransactionSet: jsii.String("transactionSet"),
//   				Version: jsii.String("version"),
//   			},
//   		},
//   		FromFormat: jsii.String("fromFormat"),
//   	},
//   	Mapping: &MappingProperty{
//   		Template: jsii.String("template"),
//   		TemplateLanguage: jsii.String("templateLanguage"),
//   	},
//   	MappingTemplate: jsii.String("mappingTemplate"),
//   	Name: jsii.String("name"),
//   	OutputConversion: &OutputConversionProperty{
//   		AdvancedOptions: &AdvancedOptionsProperty{
//   			X12: &X12AdvancedOptionsProperty{
//   				SplitOptions: &X12SplitOptionsProperty{
//   					SplitBy: jsii.String("splitBy"),
//   				},
//   				ValidationOptions: &X12ValidationOptionsProperty{
//   					ValidationRules: []interface{}{
//   						&X12ValidationRuleProperty{
//   							CodeListValidationRule: &X12CodeListValidationRuleProperty{
//   								CodesToAdd: []*string{
//   									jsii.String("codesToAdd"),
//   								},
//   								CodesToRemove: []*string{
//   									jsii.String("codesToRemove"),
//   								},
//   								ElementId: jsii.String("elementId"),
//   							},
//   							ElementLengthValidationRule: &X12ElementLengthValidationRuleProperty{
//   								ElementId: jsii.String("elementId"),
//   								MaxLength: jsii.Number(123),
//   								MinLength: jsii.Number(123),
//   							},
//   							ElementRequirementValidationRule: &X12ElementRequirementValidationRuleProperty{
//   								ElementPosition: jsii.String("elementPosition"),
//   								Requirement: jsii.String("requirement"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		FormatOptions: &FormatOptionsProperty{
//   			X12: &X12DetailsProperty{
//   				TransactionSet: jsii.String("transactionSet"),
//   				Version: jsii.String("version"),
//   			},
//   		},
//   		ToFormat: jsii.String("toFormat"),
//   	},
//   	SampleDocument: jsii.String("sampleDocument"),
//   	SampleDocuments: &SampleDocumentsProperty{
//   		BucketName: jsii.String("bucketName"),
//   		Keys: []interface{}{
//   			&SampleDocumentKeysProperty{
//   				Input: jsii.String("input"),
//   				Output: jsii.String("output"),
//   			},
//   		},
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html
//
type CfnTransformerMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-editype
	//
	// Deprecated: this property has been deprecated.
	EdiType interface{} `field:"optional" json:"ediType" yaml:"ediType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-fileformat
	//
	// Deprecated: this property has been deprecated.
	FileFormat *string `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// Returns a structure that contains the format options for the transformation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-inputconversion
	//
	InputConversion interface{} `field:"optional" json:"inputConversion" yaml:"inputConversion"`
	// Returns the structure that contains the mapping template and its language (either XSLT or JSONATA).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-mapping
	//
	Mapping interface{} `field:"optional" json:"mapping" yaml:"mapping"`
	// This shape is deprecated: This is a legacy trait.
	//
	// Please use input-conversion or output-conversion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-mappingtemplate
	//
	// Deprecated: this property has been deprecated.
	MappingTemplate *string `field:"optional" json:"mappingTemplate" yaml:"mappingTemplate"`
	// Returns the descriptive name for the transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Returns the `OutputConversion` object, which contains the format options for the outbound transformation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-outputconversion
	//
	OutputConversion interface{} `field:"optional" json:"outputConversion" yaml:"outputConversion"`
	// This shape is deprecated: This is a legacy trait.
	//
	// Please use input-conversion or output-conversion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-sampledocument
	//
	// Deprecated: this property has been deprecated.
	SampleDocument *string `field:"optional" json:"sampleDocument" yaml:"sampleDocument"`
	// Returns a structure that contains the Amazon S3 bucket and an array of the corresponding keys used to identify the location for your sample documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-sampledocuments
	//
	SampleDocuments interface{} `field:"optional" json:"sampleDocuments" yaml:"sampleDocuments"`
	// Returns the state of the newly created transformer.
	//
	// The transformer can be either `active` or `inactive` . For the transformer to be used in a capability, its status must `active` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// A key-value pair for a specific transformer.
	//
	// Tags are metadata that you can use to search for and group capabilities for various purposes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

