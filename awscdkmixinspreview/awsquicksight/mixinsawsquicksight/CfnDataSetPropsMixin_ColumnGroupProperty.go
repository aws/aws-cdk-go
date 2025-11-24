package mixinsawsquicksight


// Groupings of columns that work together in certain Quick Sight features.
//
// This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   columnGroupProperty := &ColumnGroupProperty{
//   	GeoSpatialColumnGroup: &GeoSpatialColumnGroupProperty{
//   		Columns: []*string{
//   			jsii.String("columns"),
//   		},
//   		CountryCode: jsii.String("countryCode"),
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columngroup.html
//
type CfnDataSetPropsMixin_ColumnGroupProperty struct {
	// Geospatial column group that denotes a hierarchy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columngroup.html#cfn-quicksight-dataset-columngroup-geospatialcolumngroup
	//
	GeoSpatialColumnGroup interface{} `field:"optional" json:"geoSpatialColumnGroup" yaml:"geoSpatialColumnGroup"`
}

