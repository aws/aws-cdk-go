package awsb2bi

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTransformer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransformerProps := &CfnTransformerProps{
//   	Name: jsii.String("name"),
//   	Status: jsii.String("status"),
//
//   	// the properties below are optional
//   	EdiType: &EdiTypeProperty{
//   		X12Details: &X12DetailsProperty{
//   			TransactionSet: jsii.String("transactionSet"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	FileFormat: jsii.String("fileFormat"),
//   	InputConversion: &InputConversionProperty{
//   		FromFormat: jsii.String("fromFormat"),
//
//   		// the properties below are optional
//   		AdvancedOptions: &AdvancedOptionsProperty{
//   			X12: &X12AdvancedOptionsProperty{
//   				SplitOptions: &X12SplitOptionsProperty{
//   					SplitBy: jsii.String("splitBy"),
//   				},
//   			},
//   		},
//   		FormatOptions: &FormatOptionsProperty{
//   			X12: &X12DetailsProperty{
//   				TransactionSet: jsii.String("transactionSet"),
//   				Version: jsii.String("version"),
//   			},
//   		},
//   	},
//   	Mapping: &MappingProperty{
//   		TemplateLanguage: jsii.String("templateLanguage"),
//
//   		// the properties below are optional
//   		Template: jsii.String("template"),
//   	},
//   	MappingTemplate: jsii.String("mappingTemplate"),
//   	OutputConversion: &OutputConversionProperty{
//   		ToFormat: jsii.String("toFormat"),
//
//   		// the properties below are optional
//   		FormatOptions: &FormatOptionsProperty{
//   			X12: &X12DetailsProperty{
//   				TransactionSet: jsii.String("transactionSet"),
//   				Version: jsii.String("version"),
//   			},
//   		},
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
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html
//
type CfnTransformerProps struct {
	// Returns the descriptive name for the transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Returns the state of the newly created transformer.
	//
	// The transformer can be either `active` or `inactive` . For the transformer to be used in a capability, its status must `active` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
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
	// A key-value pair for a specific transformer.
	//
	// Tags are metadata that you can use to search for and group capabilities for various purposes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

