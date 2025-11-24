package mixinsawskinesisanalyticsv2


// Specifies either the application code, or the location of the application code, for a Managed Service for Apache Flink application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeContentProperty := &CodeContentProperty{
//   	S3ContentLocation: &S3ContentLocationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		FileKey: jsii.String("fileKey"),
//   		ObjectVersion: jsii.String("objectVersion"),
//   	},
//   	TextContent: jsii.String("textContent"),
//   	ZipFileContent: jsii.String("zipFileContent"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-codecontent.html
//
type CfnApplicationPropsMixin_CodeContentProperty struct {
	// Information about the Amazon S3 bucket that contains the application code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-codecontent.html#cfn-kinesisanalyticsv2-application-codecontent-s3contentlocation
	//
	S3ContentLocation interface{} `field:"optional" json:"s3ContentLocation" yaml:"s3ContentLocation"`
	// The text-format code for a Managed Service for Apache Flink application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-codecontent.html#cfn-kinesisanalyticsv2-application-codecontent-textcontent
	//
	TextContent *string `field:"optional" json:"textContent" yaml:"textContent"`
	// The zip-format code for a Managed Service for Apache Flink application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-codecontent.html#cfn-kinesisanalyticsv2-application-codecontent-zipfilecontent
	//
	ZipFileContent *string `field:"optional" json:"zipFileContent" yaml:"zipFileContent"`
}

