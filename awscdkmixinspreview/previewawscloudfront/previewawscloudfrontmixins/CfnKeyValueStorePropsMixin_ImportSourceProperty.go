package previewawscloudfrontmixins


// The import source for the key value store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   importSourceProperty := &ImportSourceProperty{
//   	SourceArn: jsii.String("sourceArn"),
//   	SourceType: jsii.String("sourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-keyvaluestore-importsource.html
//
type CfnKeyValueStorePropsMixin_ImportSourceProperty struct {
	// The Amazon Resource Name (ARN) of the import source for the key value store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-keyvaluestore-importsource.html#cfn-cloudfront-keyvaluestore-importsource-sourcearn
	//
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
	// The source type of the import source for the key value store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-keyvaluestore-importsource.html#cfn-cloudfront-keyvaluestore-importsource-sourcetype
	//
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

