package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rdsUngracefulProperty := &RdsUngracefulProperty{
//   	Ungraceful: jsii.String("ungraceful"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdsungraceful.html
//
type CfnPlan_RdsUngracefulProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdsungraceful.html#cfn-arcregionswitch-plan-rdsungraceful-ungraceful
	//
	Ungraceful *string `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

