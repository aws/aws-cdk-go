package awsquicksight


// Groupings of columns that work together in certain Amazon QuickSight features.
//
// This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnGroupProperty := &columnGroupProperty{
//   	geoSpatialColumnGroup: &geoSpatialColumnGroupProperty{
//   		columns: []*string{
//   			jsii.String("columns"),
//   		},
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		countryCode: jsii.String("countryCode"),
//   	},
//   }
//
type CfnDataSet_ColumnGroupProperty struct {
	// Geospatial column group that denotes a hierarchy.
	GeoSpatialColumnGroup interface{} `field:"optional" json:"geoSpatialColumnGroup" yaml:"geoSpatialColumnGroup"`
}

