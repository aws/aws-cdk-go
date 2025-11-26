package previewawsqbusinessmixins


// Information required for Amazon Q Business to find a specific file in an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3Property := &S3Property{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-s3.html
//
type CfnPluginPropsMixin_S3Property struct {
	// The name of the S3 bucket that contains the file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-s3.html#cfn-qbusiness-plugin-s3-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The name of the file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-s3.html#cfn-qbusiness-plugin-s3-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
}

