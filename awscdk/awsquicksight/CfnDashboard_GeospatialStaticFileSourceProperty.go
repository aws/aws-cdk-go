package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialStaticFileSourceProperty := &GeospatialStaticFileSourceProperty{
//   	StaticFileId: jsii.String("staticFileId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialstaticfilesource.html
//
type CfnDashboard_GeospatialStaticFileSourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialstaticfilesource.html#cfn-quicksight-dashboard-geospatialstaticfilesource-staticfileid
	//
	StaticFileId *string `field:"required" json:"staticFileId" yaml:"staticFileId"`
}

