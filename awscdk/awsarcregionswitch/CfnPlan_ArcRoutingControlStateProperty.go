package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   arcRoutingControlStateProperty := &ArcRoutingControlStateProperty{
//   	RoutingControlArn: jsii.String("routingControlArn"),
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-arcroutingcontrolstate.html
//
type CfnPlan_ArcRoutingControlStateProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-arcroutingcontrolstate.html#cfn-arcregionswitch-plan-arcroutingcontrolstate-routingcontrolarn
	//
	RoutingControlArn *string `field:"required" json:"routingControlArn" yaml:"routingControlArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-arcroutingcontrolstate.html#cfn-arcregionswitch-plan-arcroutingcontrolstate-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
}

