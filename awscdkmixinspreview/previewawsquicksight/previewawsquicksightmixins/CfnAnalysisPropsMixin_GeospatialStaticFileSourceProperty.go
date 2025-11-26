package previewawsquicksightmixins


// The source properties for a geospatial static file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialStaticFileSourceProperty := &GeospatialStaticFileSourceProperty{
//   	StaticFileId: jsii.String("staticFileId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialstaticfilesource.html
//
type CfnAnalysisPropsMixin_GeospatialStaticFileSourceProperty struct {
	// The ID of the static file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialstaticfilesource.html#cfn-quicksight-analysis-geospatialstaticfilesource-staticfileid
	//
	StaticFileId *string `field:"optional" json:"staticFileId" yaml:"staticFileId"`
}

