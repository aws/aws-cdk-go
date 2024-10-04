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
	// Returns the details for the EDI standard that is being used for the transformer.
	//
	// Currently, only X12 is supported. X12 is a set of standards and corresponding messages that define specific business documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-editype
	//
	// Deprecated: this property has been deprecated.
	EdiType interface{} `field:"optional" json:"ediType" yaml:"ediType"`
	// Returns that the currently supported file formats for EDI transformations are `JSON` and `XML` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-fileformat
	//
	// Deprecated: this property has been deprecated.
	FileFormat *string `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-inputconversion
	//
	InputConversion interface{} `field:"optional" json:"inputConversion" yaml:"inputConversion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-mapping
	//
	Mapping interface{} `field:"optional" json:"mapping" yaml:"mapping"`
	// Returns a sample EDI document that is used by a transformer as a guide for processing the EDI data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-mappingtemplate
	//
	// Deprecated: this property has been deprecated.
	MappingTemplate *string `field:"optional" json:"mappingTemplate" yaml:"mappingTemplate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-outputconversion
	//
	OutputConversion interface{} `field:"optional" json:"outputConversion" yaml:"outputConversion"`
	// Returns a sample EDI document that is used by a transformer as a guide for processing the EDI data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-sampledocument
	//
	// Deprecated: this property has been deprecated.
	SampleDocument *string `field:"optional" json:"sampleDocument" yaml:"sampleDocument"`
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

