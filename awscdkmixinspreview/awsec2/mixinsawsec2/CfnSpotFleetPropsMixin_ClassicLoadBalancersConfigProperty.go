package mixinsawsec2


// Specifies the Classic Load Balancers to attach to a Spot Fleet.
//
// Spot Fleet registers the running Spot Instances with these Classic Load Balancers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   classicLoadBalancersConfigProperty := &ClassicLoadBalancersConfigProperty{
//   	ClassicLoadBalancers: []interface{}{
//   		&ClassicLoadBalancerProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-classicloadbalancersconfig.html
//
type CfnSpotFleetPropsMixin_ClassicLoadBalancersConfigProperty struct {
	// One or more Classic Load Balancers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-classicloadbalancersconfig.html#cfn-ec2-spotfleet-classicloadbalancersconfig-classicloadbalancers
	//
	ClassicLoadBalancers interface{} `field:"optional" json:"classicLoadBalancers" yaml:"classicLoadBalancers"`
}

