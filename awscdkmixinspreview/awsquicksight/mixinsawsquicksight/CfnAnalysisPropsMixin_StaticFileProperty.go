package mixinsawsquicksight


// The static file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   staticFileProperty := &StaticFileProperty{
//   	ImageStaticFile: &ImageStaticFileProperty{
//   		Source: &StaticFileSourceProperty{
//   			S3Options: &StaticFileS3SourceOptionsProperty{
//   				BucketName: jsii.String("bucketName"),
//   				ObjectKey: jsii.String("objectKey"),
//   				Region: jsii.String("region"),
//   			},
//   			UrlOptions: &StaticFileUrlSourceOptionsProperty{
//   				Url: jsii.String("url"),
//   			},
//   		},
//   		StaticFileId: jsii.String("staticFileId"),
//   	},
//   	SpatialStaticFile: &SpatialStaticFileProperty{
//   		Source: &StaticFileSourceProperty{
//   			S3Options: &StaticFileS3SourceOptionsProperty{
//   				BucketName: jsii.String("bucketName"),
//   				ObjectKey: jsii.String("objectKey"),
//   				Region: jsii.String("region"),
//   			},
//   			UrlOptions: &StaticFileUrlSourceOptionsProperty{
//   				Url: jsii.String("url"),
//   			},
//   		},
//   		StaticFileId: jsii.String("staticFileId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-staticfile.html
//
type CfnAnalysisPropsMixin_StaticFileProperty struct {
	// The image static file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-staticfile.html#cfn-quicksight-analysis-staticfile-imagestaticfile
	//
	ImageStaticFile interface{} `field:"optional" json:"imageStaticFile" yaml:"imageStaticFile"`
	// The spacial static file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-staticfile.html#cfn-quicksight-analysis-staticfile-spatialstaticfile
	//
	SpatialStaticFile interface{} `field:"optional" json:"spatialStaticFile" yaml:"spatialStaticFile"`
}

