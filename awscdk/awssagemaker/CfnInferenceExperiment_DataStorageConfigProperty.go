package awssagemaker


// The Amazon S3 location and configuration for storing inference request and response data.
//
// This is an optional parameter that you can use for data capture. For more information, see [Capture data](https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor-data-capture.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataStorageConfigProperty := &DataStorageConfigProperty{
//   	Destination: jsii.String("destination"),
//
//   	// the properties below are optional
//   	ContentType: &CaptureContentTypeHeaderProperty{
//   		CsvContentTypes: []*string{
//   			jsii.String("csvContentTypes"),
//   		},
//   		JsonContentTypes: []*string{
//   			jsii.String("jsonContentTypes"),
//   		},
//   	},
//   	KmsKey: jsii.String("kmsKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-datastorageconfig.html
//
type CfnInferenceExperiment_DataStorageConfigProperty struct {
	// The Amazon S3 bucket where the inference request and response data is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-datastorageconfig.html#cfn-sagemaker-inferenceexperiment-datastorageconfig-destination
	//
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Configuration specifying how to treat different headers.
	//
	// If no headers are specified SageMaker will by default base64 encode when capturing the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-datastorageconfig.html#cfn-sagemaker-inferenceexperiment-datastorageconfig-contenttype
	//
	ContentType interface{} `field:"optional" json:"contentType" yaml:"contentType"`
	// The AWS Key Management Service key that Amazon SageMaker uses to encrypt captured data at rest using Amazon S3 server-side encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-datastorageconfig.html#cfn-sagemaker-inferenceexperiment-datastorageconfig-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

