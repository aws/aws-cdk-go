package awsappflow


// When you use Amazon S3 as the source, the configuration format that you provide the flow input data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3InputFormatConfigProperty := &S3InputFormatConfigProperty{
//   	S3InputFileType: jsii.String("s3InputFileType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-s3inputformatconfig.html
//
type CfnFlow_S3InputFormatConfigProperty struct {
	// The file type that Amazon AppFlow gets from your Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-s3inputformatconfig.html#cfn-appflow-flow-s3inputformatconfig-s3inputfiletype
	//
	S3InputFileType *string `field:"optional" json:"s3InputFileType" yaml:"s3InputFileType"`
}

