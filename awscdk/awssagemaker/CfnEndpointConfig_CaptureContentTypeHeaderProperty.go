package awssagemaker


// Specifies the JSON and CSV content types of the data that the endpoint captures.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-capturecontenttypeheader.html
//
type CfnEndpointConfig_CaptureContentTypeHeaderProperty struct {
	// A list of the CSV content types of the data that the endpoint captures.
	//
	// For the endpoint to capture the data, you must also specify the content type when you invoke the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-capturecontenttypeheader.html#cfn-sagemaker-endpointconfig-capturecontenttypeheader-csvcontenttypes
	//
	CsvContentTypes *[]*string `field:"optional" json:"csvContentTypes" yaml:"csvContentTypes"`
	// A list of the JSON content types of the data that the endpoint captures.
	//
	// For the endpoint to capture the data, you must also specify the content type when you invoke the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-capturecontenttypeheader.html#cfn-sagemaker-endpointconfig-capturecontenttypeheader-jsoncontenttypes
	//
	JsonContentTypes *[]*string `field:"optional" json:"jsonContentTypes" yaml:"jsonContentTypes"`
}

