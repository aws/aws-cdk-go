package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialLineWidthProperty := &GeospatialLineWidthProperty{
//   	LineWidth: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatiallinewidth.html
//
type CfnTemplate_GeospatialLineWidthProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatiallinewidth.html#cfn-quicksight-template-geospatiallinewidth-linewidth
	//
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
}

