package awsroute53resolver


// In a [CreateResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverRule.html) request, an array of the IPs that you want to forward DNS queries to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetAddressProperty := &TargetAddressProperty{
//   	Ip: jsii.String("ip"),
//
//   	// the properties below are optional
//   	Port: jsii.String("port"),
//   }
//
type CfnResolverRule_TargetAddressProperty struct {
	// One IPv4 address that you want to forward DNS queries to.
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// The port at `Ip` that you want to forward DNS queries to.
	Port *string `field:"optional" json:"port" yaml:"port"`
}

