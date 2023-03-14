package awsquicksight


// Geospatial column group that denotes a hierarchy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geoSpatialColumnGroupProperty := &geoSpatialColumnGroupProperty{
//   	columns: []*string{
//   		jsii.String("columns"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	countryCode: jsii.String("countryCode"),
//   }
//
type CfnDataSet_GeoSpatialColumnGroupProperty struct {
	// Columns in this hierarchy.
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
	// A display name for the hierarchy.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Country code.
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
}

