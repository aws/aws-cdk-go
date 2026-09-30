package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   spatialStaticFileProperty := &SpatialStaticFileProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-spatialstaticfile.html
//
type CfnTemplatePropsMixin_SpatialStaticFileProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-spatialstaticfile.html#cfn-quicksight-template-spatialstaticfile-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-spatialstaticfile.html#cfn-quicksight-template-spatialstaticfile-staticfileid
	//
	StaticFileId *string `field:"optional" json:"staticFileId" yaml:"staticFileId"`
}

