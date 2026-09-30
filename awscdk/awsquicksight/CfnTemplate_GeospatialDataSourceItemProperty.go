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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialdatasourceitem.html
//
type CfnTemplate_GeospatialDataSourceItemProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialdatasourceitem.html#cfn-quicksight-template-geospatialdatasourceitem-staticfiledatasource
	//
	StaticFileDataSource interface{} `field:"optional" json:"staticFileDataSource" yaml:"staticFileDataSource"`
}

