package awssagemaker


// Configuration specifying how to treat different headers.
//
// If no headers are specified Amazon SageMaker will by default base64 encode when capturing the data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captureContentTypeHeaderProperty := &CaptureContentTypeHeaderProperty{
//   	CsvContentTypes: []*string{
//   		jsii.String("csvContentTypes"),
//   	},
//   	JsonContentTypes: []*string{
//   		jsii.String("jsonContentTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-capturecontenttypeheader.html
//
type CfnInferenceExperiment_CaptureContentTypeHeaderProperty struct {
	// The list of all content type headers that Amazon SageMaker will treat as CSV and capture accordingly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-capturecontenttypeheader.html#cfn-sagemaker-inferenceexperiment-capturecontenttypeheader-csvcontenttypes
	//
	CsvContentTypes *[]*string `field:"optional" json:"csvContentTypes" yaml:"csvContentTypes"`
	// The list of all content type headers that SageMaker will treat as JSON and capture accordingly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-capturecontenttypeheader.html#cfn-sagemaker-inferenceexperiment-capturecontenttypeheader-jsoncontenttypes
	//
	JsonContentTypes *[]*string `field:"optional" json:"jsonContentTypes" yaml:"jsonContentTypes"`
}

