package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   neptuneUngracefulProperty := &NeptuneUngracefulProperty{
//   	Ungraceful: jsii.String("ungraceful"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-neptuneungraceful.html
//
type CfnPlanPropsMixin_NeptuneUngracefulProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-neptuneungraceful.html#cfn-arcregionswitch-plan-neptuneungraceful-ungraceful
	//
	Ungraceful *string `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

