package previewawsquicksightmixins


// The data source properties for the geospatial data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialDataSourceItemProperty := &GeospatialDataSourceItemProperty{
//   	StaticFileDataSource: &GeospatialStaticFileSourceProperty{
//   		StaticFileId: jsii.String("staticFileId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialdatasourceitem.html
//
type CfnAnalysisPropsMixin_GeospatialDataSourceItemProperty struct {
	// The static file data source properties for the geospatial data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialdatasourceitem.html#cfn-quicksight-analysis-geospatialdatasourceitem-staticfiledatasource
	//
	StaticFileDataSource interface{} `field:"optional" json:"staticFileDataSource" yaml:"staticFileDataSource"`
}

