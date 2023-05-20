package awsec2


// Describes the load balancer options when creating an AWS Verified Access endpoint using the `load-balancer` type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerOptionsProperty := &LoadBalancerOptionsProperty{
//   	LoadBalancerArn: jsii.String("loadBalancerArn"),
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnVerifiedAccessEndpoint_LoadBalancerOptionsProperty struct {
	// The ARN of the load balancer.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// The IP port number.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The IP protocol.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The IDs of the subnets.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

