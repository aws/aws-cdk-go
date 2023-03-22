package awssagemaker


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
type CfnInferenceExperiment_DataStorageConfigProperty struct {
	// `CfnInferenceExperiment.DataStorageConfigProperty.Destination`.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// `CfnInferenceExperiment.DataStorageConfigProperty.ContentType`.
	ContentType interface{} `field:"optional" json:"contentType" yaml:"contentType"`
	// `CfnInferenceExperiment.DataStorageConfigProperty.KmsKey`.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

