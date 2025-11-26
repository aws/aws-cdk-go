package previewawsec2mixins


// Specifies a Classic Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   classicLoadBalancerProperty := &ClassicLoadBalancerProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-classicloadbalancer.html
//
type CfnSpotFleetPropsMixin_ClassicLoadBalancerProperty struct {
	// The name of the load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-classicloadbalancer.html#cfn-ec2-spotfleet-classicloadbalancer-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

