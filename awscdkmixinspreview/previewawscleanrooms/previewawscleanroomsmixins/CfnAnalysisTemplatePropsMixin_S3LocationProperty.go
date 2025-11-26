package previewawscleanroomsmixins


// The S3 location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3LocationProperty := &S3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-s3location.html
//
type CfnAnalysisTemplatePropsMixin_S3LocationProperty struct {
	// The bucket name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-s3location.html#cfn-cleanrooms-analysistemplate-s3location-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The object key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-s3location.html#cfn-cleanrooms-analysistemplate-s3location-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
}

