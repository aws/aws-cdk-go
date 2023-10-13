package awsquicksight


// Configures the properties of a chart's axes that are used by small multiples panels.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   smallMultiplesAxisPropertiesProperty := &SmallMultiplesAxisPropertiesProperty{
//   	Placement: jsii.String("placement"),
//   	Scale: jsii.String("scale"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-smallmultiplesaxisproperties.html
//
type CfnDashboard_SmallMultiplesAxisPropertiesProperty struct {
	// Defines the placement of the axis.
	//
	// By default, axes are rendered `OUTSIDE` of the panels. Axes with `INDEPENDENT` scale are rendered `INSIDE` the panels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-smallmultiplesaxisproperties.html#cfn-quicksight-dashboard-smallmultiplesaxisproperties-placement
	//
	Placement *string `field:"optional" json:"placement" yaml:"placement"`
	// Determines whether scale of the axes are shared or independent.
	//
	// The default value is `SHARED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-smallmultiplesaxisproperties.html#cfn-quicksight-dashboard-smallmultiplesaxisproperties-scale
	//
	Scale *string `field:"optional" json:"scale" yaml:"scale"`
}

