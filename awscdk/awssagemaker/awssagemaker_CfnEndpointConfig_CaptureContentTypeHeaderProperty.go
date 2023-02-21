package awssagemaker


// Specifies the JSON and CSV content types of the data that the endpoint captures.
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
type CfnEndpointConfig_CaptureContentTypeHeaderProperty struct {
	// A list of the CSV content types of the data that the endpoint captures.
	//
	// For the endpoint to capture the data, you must also specify the content type when you invoke the endpoint.
	CsvContentTypes *[]*string `field:"optional" json:"csvContentTypes" yaml:"csvContentTypes"`
	// A list of the JSON content types of the data that the endpoint captures.
	//
	// For the endpoint to capture the data, you must also specify the content type when you invoke the endpoint.
	JsonContentTypes *[]*string `field:"optional" json:"jsonContentTypes" yaml:"jsonContentTypes"`
}

