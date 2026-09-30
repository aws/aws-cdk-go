package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   imageStaticFileProperty := &ImageStaticFileProperty{
//   	Source: &StaticFileSourceProperty{
//   		S3Options: &StaticFileS3SourceOptionsProperty{
//   			BucketName: jsii.String("bucketName"),
//   			ObjectKey: jsii.String("objectKey"),
//   			Region: jsii.String("region"),
//   		},
//   		UrlOptions: &StaticFileUrlSourceOptionsProperty{
//   			Url: jsii.String("url"),
//   		},
//   	},
//   	StaticFileId: jsii.String("staticFileId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagestaticfile.html
//
type CfnTemplatePropsMixin_ImageStaticFileProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagestaticfile.html#cfn-quicksight-template-imagestaticfile-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagestaticfile.html#cfn-quicksight-template-imagestaticfile-staticfileid
	//
	StaticFileId *string `field:"optional" json:"staticFileId" yaml:"staticFileId"`
}

