package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialCircleRadiusProperty := &GeospatialCircleRadiusProperty{
//   	Radius: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialcircleradius.html
//
type CfnTemplate_GeospatialCircleRadiusProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialcircleradius.html#cfn-quicksight-template-geospatialcircleradius-radius
	//
	Radius *float64 `field:"optional" json:"radius" yaml:"radius"`
}

