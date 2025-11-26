package previewawsquicksightmixins


// The structure that contains the Amazon S3 location to download the static file from.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   staticFileS3SourceOptionsProperty := &StaticFileS3SourceOptionsProperty{
//   	BucketName: jsii.String("bucketName"),
//   	ObjectKey: jsii.String("objectKey"),
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-staticfiles3sourceoptions.html
//
type CfnAnalysisPropsMixin_StaticFileS3SourceOptionsProperty struct {
	// The name of the Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-staticfiles3sourceoptions.html#cfn-quicksight-analysis-staticfiles3sourceoptions-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The identifier of the static file in the Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-staticfiles3sourceoptions.html#cfn-quicksight-analysis-staticfiles3sourceoptions-objectkey
	//
	ObjectKey *string `field:"optional" json:"objectKey" yaml:"objectKey"`
	// The Region of the Amazon S3 account that contains the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-staticfiles3sourceoptions.html#cfn-quicksight-analysis-staticfiles3sourceoptions-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

