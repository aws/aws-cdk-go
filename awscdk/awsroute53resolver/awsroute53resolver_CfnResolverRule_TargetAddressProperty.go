package awsroute53resolver


// In a [CreateResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverRule.html) request, an array of the IPs that you want to forward DNS queries to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetAddressProperty := &targetAddressProperty{
//   	ip: jsii.String("ip"),
//
//   	// the properties below are optional
//   	port: jsii.String("port"),
//   }
//
type CfnResolverRule_TargetAddressProperty struct {
	// One IP address that you want to forward DNS queries to.
	//
	// You can specify only IPv4 addresses.
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// The port at `Ip` that you want to forward DNS queries to.
	Port *string `field:"optional" json:"port" yaml:"port"`
}

