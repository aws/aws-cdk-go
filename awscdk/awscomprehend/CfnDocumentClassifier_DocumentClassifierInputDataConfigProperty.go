package awscomprehend


// The input properties for training a document classifier.
//
// For more information on how the input file is formatted, see [Preparing training data](https://docs.aws.amazon.com/comprehend/latest/dg/prep-classifier-data.html) in the Comprehend Developer Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentClassifierInputDataConfigProperty := &DocumentClassifierInputDataConfigProperty{
//   	AugmentedManifests: []interface{}{
//   		&AugmentedManifestsListItemProperty{
//   			AttributeNames: []*string{
//   				jsii.String("attributeNames"),
//   			},
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			Split: jsii.String("split"),
//   		},
//   	},
//   	DataFormat: jsii.String("dataFormat"),
//   	DocumentReaderConfig: &DocumentReaderConfigProperty{
//   		DocumentReadAction: jsii.String("documentReadAction"),
//
//   		// the properties below are optional
//   		DocumentReadMode: jsii.String("documentReadMode"),
//   		FeatureTypes: []*string{
//   			jsii.String("featureTypes"),
//   		},
//   	},
//   	Documents: &DocumentClassifierDocumentsProperty{
//   		S3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		TestS3Uri: jsii.String("testS3Uri"),
//   	},
//   	DocumentType: jsii.String("documentType"),
//   	LabelDelimiter: jsii.String("labelDelimiter"),
//   	S3Uri: jsii.String("s3Uri"),
//   	TestS3Uri: jsii.String("testS3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentclassifierinputdataconfig.html
//
type CfnDocumentClassifier_DocumentClassifierInputDataConfigProperty struct {
	// A list of augmented manifest files that provide training data for your custom model.
	//
	// An augmented manifest file is a labeled dataset that is produced by Amazon SageMaker Ground Truth.
	//
	// This parameter is required if you set `DataFormat` to `AUGMENTED_MANIFEST` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentclassifierinputdataconfig.html#cfn-comprehend-documentclassifier-documentclassifierinputdataconfig-augmentedmanifests
	//
	AugmentedManifests interface{} `field:"optional" json:"augmentedManifests" yaml:"augmentedManifests"`
	// The format of your training data:.
	//
	// - `COMPREHEND_CSV` : A two-column CSV file, where labels are provided in the first column, and documents are provided in the second. If you use this value, you must provide the `S3Uri` parameter in your request.
	// - `AUGMENTED_MANIFEST` : A labeled dataset that is produced by Amazon SageMaker Ground Truth. This file is in JSON lines format. Each line is a complete JSON object that contains a training document and its associated labels.
	//
	// If you use this value, you must provide the `AugmentedManifests` parameter in your request.
	//
	// If you don't specify a value, Amazon Comprehend uses `COMPREHEND_CSV` as the default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentclassifierinputdataconfig.html#cfn-comprehend-documentclassifier-documentclassifierinputdataconfig-dataformat
	//
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentclassifierinputdataconfig.html#cfn-comprehend-documentclassifier-documentclassifierinputdataconfig-documentreaderconfig
	//
	DocumentReaderConfig interface{} `field:"optional" json:"documentReaderConfig" yaml:"documentReaderConfig"`
	// The S3 location of the training documents.
	//
	// This parameter is required in a request to create a native document model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentclassifierinputdataconfig.html#cfn-comprehend-documentclassifier-documentclassifierinputdataconfig-documents
	//
	Documents interface{} `field:"optional" json:"documents" yaml:"documents"`
	// The type of input documents for training the model.
	//
	// Provide plain-text documents to create a plain-text model, and provide semi-structured documents to create a native document model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentclassifierinputdataconfig.html#cfn-comprehend-documentclassifier-documentclassifierinputdataconfig-documenttype
	//
	DocumentType *string `field:"optional" json:"documentType" yaml:"documentType"`
	// Indicates the delimiter used to separate each label for training a multi-label classifier.
	//
	// The default delimiter between labels is a pipe (|). You can use a different character as a delimiter (if it's an allowed character) by specifying it under Delimiter for labels. If the training documents use a delimiter other than the default or the delimiter you specify, the labels on that line will be combined to make a single unique label, such as LABELLABELLABEL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentclassifierinputdataconfig.html#cfn-comprehend-documentclassifier-documentclassifierinputdataconfig-labeldelimiter
	//
	LabelDelimiter *string `field:"optional" json:"labelDelimiter" yaml:"labelDelimiter"`
	// The Amazon S3 URI for the input data.
	//
	// The S3 bucket must be in the same Region as the API endpoint that you are calling. The URI can point to a single input file or it can provide the prefix for a collection of input files.
	//
	// For example, if you use the URI `S3://bucketName/prefix` , if the prefix is a single file, Amazon Comprehend uses that file as input. If more than one file begins with the prefix, Amazon Comprehend uses all of them as input.
	//
	// This parameter is required if you set `DataFormat` to `COMPREHEND_CSV` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentclassifierinputdataconfig.html#cfn-comprehend-documentclassifier-documentclassifierinputdataconfig-s3uri
	//
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
	// This specifies the Amazon S3 location that contains the test annotations for the document classifier.
	//
	// The URI must be in the same AWS Region as the API endpoint that you are calling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentclassifierinputdataconfig.html#cfn-comprehend-documentclassifier-documentclassifierinputdataconfig-tests3uri
	//
	TestS3Uri *string `field:"optional" json:"testS3Uri" yaml:"testS3Uri"`
}

