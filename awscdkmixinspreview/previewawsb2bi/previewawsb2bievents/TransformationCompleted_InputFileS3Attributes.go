package previewawsb2bievents


// Type definition for Input-file-s3-attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputFileS3Attributes := &InputFileS3Attributes{
//   	Bucket: []*string{
//   		jsii.String("bucket"),
//   	},
//   	ObjectKey: []*string{
//   		jsii.String("objectKey"),
//   	},
//   	ObjectSizeBytes: []*string{
//   		jsii.String("objectSizeBytes"),
//   	},
//   }
//
// Experimental.
type TransformationCompleted_InputFileS3Attributes struct {
	// bucket property.
	//
	// Specify an array of string values to match this event if the actual value of bucket is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Bucket *[]*string `field:"optional" json:"bucket" yaml:"bucket"`
	// object-key property.
	//
	// Specify an array of string values to match this event if the actual value of object-key is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ObjectKey *[]*string `field:"optional" json:"objectKey" yaml:"objectKey"`
	// object-size-bytes property.
	//
	// Specify an array of string values to match this event if the actual value of object-size-bytes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ObjectSizeBytes *[]*string `field:"optional" json:"objectSizeBytes" yaml:"objectSizeBytes"`
}

