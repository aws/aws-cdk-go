package awsquicksight


// The static file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   staticFileProperty := &StaticFileProperty{
//   	ImageStaticFile: &ImageStaticFileProperty{
//   		StaticFileId: jsii.String("staticFileId"),
//
//   		// the properties below are optional
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
//   	},
//   	SpatialStaticFile: &SpatialStaticFileProperty{
//   		StaticFileId: jsii.String("staticFileId"),
//
//   		// the properties below are optional
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
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-staticfile.html
//
type CfnDashboard_StaticFileProperty struct {
	// The image static file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-staticfile.html#cfn-quicksight-dashboard-staticfile-imagestaticfile
	//
	ImageStaticFile interface{} `field:"optional" json:"imageStaticFile" yaml:"imageStaticFile"`
	// The spacial static file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-staticfile.html#cfn-quicksight-dashboard-staticfile-spatialstaticfile
	//
	SpatialStaticFile interface{} `field:"optional" json:"spatialStaticFile" yaml:"spatialStaticFile"`
}

