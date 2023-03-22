package awssagemaker


// Configuration specifying how to treat different headers.
//
// If no headers are specified SageMaker will by default base64 encode when capturing the data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captureContentTypeHeaderProperty := &captureContentTypeHeaderProperty{
//   	csvContentTypes: []*string{
//   		jsii.String("csvContentTypes"),
//   	},
//   	jsonContentTypes: []*string{
//   		jsii.String("jsonContentTypes"),
//   	},
//   }
//
type CfnInferenceExperiment_CaptureContentTypeHeaderProperty struct {
	// The list of all content type headers that SageMaker will treat as CSV and capture accordingly.
	CsvContentTypes *[]*string `field:"optional" json:"csvContentTypes" yaml:"csvContentTypes"`
	// The list of all content type headers that SageMaker will treat as JSON and capture accordingly.
	JsonContentTypes *[]*string `field:"optional" json:"jsonContentTypes" yaml:"jsonContentTypes"`
}

