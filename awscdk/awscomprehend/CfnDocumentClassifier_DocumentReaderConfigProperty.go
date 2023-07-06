package awscomprehend


// Provides configuration parameters to override the default actions for extracting text from PDF documents and image files.
//
// By default, Amazon Comprehend performs the following actions to extract text from files, based on the input file type:
//
// - *Word files* - Amazon Comprehend parser extracts the text.
// - *Digital PDF files* - Amazon Comprehend parser extracts the text.
// - *Image files and scanned PDF files* - Amazon Comprehend uses the Amazon Textract `DetectDocumentText` API to extract the text.
//
// `DocumentReaderConfig` does not apply to plain text files or Word files.
//
// For image files and PDF documents, you can override these default actions using the fields listed below. For more information, see [Setting text extraction options](https://docs.aws.amazon.com/comprehend/latest/dg/idp-set-textract-options.html) in the Comprehend Developer Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentReaderConfigProperty := &DocumentReaderConfigProperty{
//   	DocumentReadAction: jsii.String("documentReadAction"),
//
//   	// the properties below are optional
//   	DocumentReadMode: jsii.String("documentReadMode"),
//   	FeatureTypes: []*string{
//   		jsii.String("featureTypes"),
//   	},
//   }
//
type CfnDocumentClassifier_DocumentReaderConfigProperty struct {
	// This field defines the Amazon Textract API operation that Amazon Comprehend uses to extract text from PDF files and image files.
	//
	// Enter one of the following values:
	//
	// - `TEXTRACT_DETECT_DOCUMENT_TEXT` - The Amazon Comprehend service uses the `DetectDocumentText` API operation.
	// - `TEXTRACT_ANALYZE_DOCUMENT` - The Amazon Comprehend service uses the `AnalyzeDocument` API operation.
	DocumentReadAction *string `field:"required" json:"documentReadAction" yaml:"documentReadAction"`
	// Determines the text extraction actions for PDF files. Enter one of the following values:.
	//
	// - `SERVICE_DEFAULT` - use the Amazon Comprehend service defaults for PDF files.
	// - `FORCE_DOCUMENT_READ_ACTION` - Amazon Comprehend uses the Textract API specified by DocumentReadAction for all PDF files, including digital PDF files.
	DocumentReadMode *string `field:"optional" json:"documentReadMode" yaml:"documentReadMode"`
	// Specifies the type of Amazon Textract features to apply.
	//
	// If you chose `TEXTRACT_ANALYZE_DOCUMENT` as the read action, you must specify one or both of the following values:
	//
	// - `TABLES` - Returns information about any tables that are detected in the input document.
	// - `FORMS` - Returns information and the data from any forms that are detected in the input document.
	FeatureTypes *[]*string `field:"optional" json:"featureTypes" yaml:"featureTypes"`
}

