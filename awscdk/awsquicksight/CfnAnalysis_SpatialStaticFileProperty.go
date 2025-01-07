package awsquicksight


// A static file that contains the geospatial data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spatialStaticFileProperty := &SpatialStaticFileProperty{
//   	StaticFileId: jsii.String("staticFileId"),
//
//   	// the properties below are optional
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-spatialstaticfile.html
//
type CfnAnalysis_SpatialStaticFileProperty struct {
	// The ID of the spatial static file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-spatialstaticfile.html#cfn-quicksight-analysis-spatialstaticfile-staticfileid
	//
	StaticFileId *string `field:"required" json:"staticFileId" yaml:"staticFileId"`
	// The source of the spatial static file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-spatialstaticfile.html#cfn-quicksight-analysis-spatialstaticfile-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

