package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ssmSendCommandUngracefulProperty := &SsmSendCommandUngracefulProperty{
//   	Behavior: jsii.String("behavior"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmsendcommandungraceful.html
//
type CfnPlan_SsmSendCommandUngracefulProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmsendcommandungraceful.html#cfn-arcregionswitch-plan-ssmsendcommandungraceful-behavior
	//
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
}

