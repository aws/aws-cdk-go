package awsquicksight


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-smallmultiplesaxisproperties.html
//
type CfnTemplate_SmallMultiplesAxisPropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-smallmultiplesaxisproperties.html#cfn-quicksight-template-smallmultiplesaxisproperties-placement
	//
	Placement *string `field:"optional" json:"placement" yaml:"placement"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-smallmultiplesaxisproperties.html#cfn-quicksight-template-smallmultiplesaxisproperties-scale
	//
	Scale *string `field:"optional" json:"scale" yaml:"scale"`
}

