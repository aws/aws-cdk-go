package previewawsbedrockagentcoremixins


// The S3 location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3LocationProperty := &S3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-s3location.html
//
type CfnBrowserCustomPropsMixin_S3LocationProperty struct {
	// The S3 location bucket name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-s3location.html#cfn-bedrockagentcore-browsercustom-s3location-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 location object prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-s3location.html#cfn-bedrockagentcore-browsercustom-s3location-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

