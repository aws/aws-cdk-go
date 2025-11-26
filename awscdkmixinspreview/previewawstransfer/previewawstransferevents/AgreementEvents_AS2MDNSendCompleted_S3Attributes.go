package previewawstransferevents


// Type definition for S3-attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3Attributes := &S3Attributes{
//   	FileBucket: []*string{
//   		jsii.String("fileBucket"),
//   	},
//   	FileKey: []*string{
//   		jsii.String("fileKey"),
//   	},
//   	JsonBucket: []*string{
//   		jsii.String("jsonBucket"),
//   	},
//   	JsonKey: []*string{
//   		jsii.String("jsonKey"),
//   	},
//   	MdnBucket: []*string{
//   		jsii.String("mdnBucket"),
//   	},
//   	MdnKey: []*string{
//   		jsii.String("mdnKey"),
//   	},
//   }
//
// Experimental.
type AgreementEvents_AS2MDNSendCompleted_S3Attributes struct {
	// file-bucket property.
	//
	// Specify an array of string values to match this event if the actual value of file-bucket is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FileBucket *[]*string `field:"optional" json:"fileBucket" yaml:"fileBucket"`
	// file-key property.
	//
	// Specify an array of string values to match this event if the actual value of file-key is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FileKey *[]*string `field:"optional" json:"fileKey" yaml:"fileKey"`
	// json-bucket property.
	//
	// Specify an array of string values to match this event if the actual value of json-bucket is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JsonBucket *[]*string `field:"optional" json:"jsonBucket" yaml:"jsonBucket"`
	// json-key property.
	//
	// Specify an array of string values to match this event if the actual value of json-key is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JsonKey *[]*string `field:"optional" json:"jsonKey" yaml:"jsonKey"`
	// mdn-bucket property.
	//
	// Specify an array of string values to match this event if the actual value of mdn-bucket is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MdnBucket *[]*string `field:"optional" json:"mdnBucket" yaml:"mdnBucket"`
	// mdn-key property.
	//
	// Specify an array of string values to match this event if the actual value of mdn-key is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MdnKey *[]*string `field:"optional" json:"mdnKey" yaml:"mdnKey"`
}

