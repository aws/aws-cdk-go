package awsquicksight


// Describes the properties for a solid color.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   geospatialSolidColorProperty := &GeospatialSolidColorProperty{
//   	Color: jsii.String("color"),
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialsolidcolor.html
//
type CfnTemplatePropsMixin_GeospatialSolidColorProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialsolidcolor.html#cfn-quicksight-template-geospatialsolidcolor-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Defines view state of the color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialsolidcolor.html#cfn-quicksight-template-geospatialsolidcolor-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

