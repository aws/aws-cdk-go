package awsbcmpricingcalculator


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   billIntervalProperty := &BillIntervalProperty{
//   	End: jsii.String("end"),
//   	Start: jsii.String("start"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmpricingcalculator-billscenario-billinterval.html
//
type CfnBillScenarioPropsMixin_BillIntervalProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmpricingcalculator-billscenario-billinterval.html#cfn-bcmpricingcalculator-billscenario-billinterval-end
	//
	End *string `field:"optional" json:"end" yaml:"end"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmpricingcalculator-billscenario-billinterval.html#cfn-bcmpricingcalculator-billscenario-billinterval-start
	//
	Start *string `field:"optional" json:"start" yaml:"start"`
}

