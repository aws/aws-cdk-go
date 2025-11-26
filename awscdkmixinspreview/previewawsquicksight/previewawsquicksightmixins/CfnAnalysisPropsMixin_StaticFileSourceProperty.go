package previewawsquicksightmixins


// The source of the static file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   staticFileSourceProperty := &StaticFileSourceProperty{
//   	S3Options: &StaticFileS3SourceOptionsProperty{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKey: jsii.String("objectKey"),
//   		Region: jsii.String("region"),
//   	},
//   	UrlOptions: &StaticFileUrlSourceOptionsProperty{
//   		Url: jsii.String("url"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-staticfilesource.html
//
type CfnAnalysisPropsMixin_StaticFileSourceProperty struct {
	// The structure that contains the Amazon S3 location to download the static file from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-staticfilesource.html#cfn-quicksight-analysis-staticfilesource-s3options
	//
	S3Options interface{} `field:"optional" json:"s3Options" yaml:"s3Options"`
	// The structure that contains the URL to download the static file from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-staticfilesource.html#cfn-quicksight-analysis-staticfilesource-urloptions
	//
	UrlOptions interface{} `field:"optional" json:"urlOptions" yaml:"urlOptions"`
}

