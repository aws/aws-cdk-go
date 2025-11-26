package previewawsquicksightmixins


// Geospatial column group that denotes a hierarchy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geoSpatialColumnGroupProperty := &GeoSpatialColumnGroupProperty{
//   	Columns: []*string{
//   		jsii.String("columns"),
//   	},
//   	CountryCode: jsii.String("countryCode"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-geospatialcolumngroup.html
//
type CfnDataSetPropsMixin_GeoSpatialColumnGroupProperty struct {
	// Columns in this hierarchy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-geospatialcolumngroup.html#cfn-quicksight-dataset-geospatialcolumngroup-columns
	//
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// Country code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-geospatialcolumngroup.html#cfn-quicksight-dataset-geospatialcolumngroup-countrycode
	//
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// A display name for the hierarchy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-geospatialcolumngroup.html#cfn-quicksight-dataset-geospatialcolumngroup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

