package awsquicksight


// Geospatial column group that denotes a hierarchy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geoSpatialColumnGroupProperty := &GeoSpatialColumnGroupProperty{
//   	Columns: []*string{
//   		jsii.String("columns"),
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CountryCode: jsii.String("countryCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-geospatialcolumngroup.html
//
type CfnDataSet_GeoSpatialColumnGroupProperty struct {
	// Columns in this hierarchy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-geospatialcolumngroup.html#cfn-quicksight-dataset-geospatialcolumngroup-columns
	//
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
	// A display name for the hierarchy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-geospatialcolumngroup.html#cfn-quicksight-dataset-geospatialcolumngroup-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Country code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-geospatialcolumngroup.html#cfn-quicksight-dataset-geospatialcolumngroup-countrycode
	//
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
}

