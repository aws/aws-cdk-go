package awskinesisanalytics


// Specifies either the application code, or the location of the application code, for a Flink-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeContentProperty := &codeContentProperty{
//   	s3ContentLocation: &s3ContentLocationProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		fileKey: jsii.String("fileKey"),
//
//   		// the properties below are optional
//   		objectVersion: jsii.String("objectVersion"),
//   	},
//   	textContent: jsii.String("textContent"),
//   	zipFileContent: jsii.String("zipFileContent"),
//   }
//
type CfnApplicationV2_CodeContentProperty struct {
	// Information about the Amazon S3 bucket that contains the application code.
	S3ContentLocation interface{} `field:"optional" json:"s3ContentLocation" yaml:"s3ContentLocation"`
	// The text-format code for a Flink-based Kinesis Data Analytics application.
	TextContent *string `field:"optional" json:"textContent" yaml:"textContent"`
	// The zip-format code for a Flink-based Kinesis Data Analytics application.
	ZipFileContent *string `field:"optional" json:"zipFileContent" yaml:"zipFileContent"`
}

