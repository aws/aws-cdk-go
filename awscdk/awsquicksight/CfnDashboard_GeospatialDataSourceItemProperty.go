package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialDataSourceItemProperty := &GeospatialDataSourceItemProperty{
//   	StaticFileDataSource: &GeospatialStaticFileSourceProperty{
//   		StaticFileId: jsii.String("staticFileId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialdatasourceitem.html
//
type CfnDashboard_GeospatialDataSourceItemProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialdatasourceitem.html#cfn-quicksight-dashboard-geospatialdatasourceitem-staticfiledatasource
	//
	StaticFileDataSource interface{} `field:"optional" json:"staticFileDataSource" yaml:"staticFileDataSource"`
}

