package awselasticloadbalancingv2


// Result of attaching a target to load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var targetJson interface{}
//
//   loadBalancerTargetProps := &loadBalancerTargetProps{
//   	targetType: awscdk.Aws_elasticloadbalancingv2.targetType_INSTANCE,
//
//   	// the properties below are optional
//   	targetJson: targetJson,
//   }
//
type LoadBalancerTargetProps struct {
	// What kind of target this is.
	TargetType TargetType `field:"required" json:"targetType" yaml:"targetType"`
	// JSON representing the target's direct addition to the TargetGroup list.
	//
	// May be omitted if the target is going to register itself later.
	TargetJson interface{} `field:"optional" json:"targetJson" yaml:"targetJson"`
}

