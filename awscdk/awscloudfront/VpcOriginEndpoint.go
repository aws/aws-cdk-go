package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// Represents the VPC origin endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationLoadBalancer ApplicationLoadBalancer
//
//   vpcOriginEndpoint := awscdk.Aws_cloudfront.VpcOriginEndpoint_ApplicationLoadBalancer(applicationLoadBalancer)
//
type VpcOriginEndpoint interface {
	// The domain name of the CloudFront VPC origin endpoint configuration.
	// Default: - No domain name configured.
	//
	DomainName() *string
	// The ARN of the CloudFront VPC origin endpoint configuration.
	EndpointArn() *string
}

// The jsii proxy struct for VpcOriginEndpoint
type jsiiProxy_VpcOriginEndpoint struct {
	_ byte // padding
}

func (j *jsiiProxy_VpcOriginEndpoint) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcOriginEndpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}


func NewVpcOriginEndpoint_Override(v VpcOriginEndpoint) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.VpcOriginEndpoint",
		nil, // no parameters
		v,
	)
}

// A VPC origin endpoint from an Application Load Balancer.
func VpcOriginEndpoint_ApplicationLoadBalancer(alb awselasticloadbalancingv2.IApplicationLoadBalancer) VpcOriginEndpoint {
	_init_.Initialize()

	if err := validateVpcOriginEndpoint_ApplicationLoadBalancerParameters(alb); err != nil {
		panic(err)
	}
	var returns VpcOriginEndpoint

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.VpcOriginEndpoint",
		"applicationLoadBalancer",
		[]interface{}{alb},
		&returns,
	)

	return returns
}

// A VPC origin endpoint from an EC2 instance.
func VpcOriginEndpoint_Ec2Instance(instance awsec2.IInstance) VpcOriginEndpoint {
	_init_.Initialize()

	if err := validateVpcOriginEndpoint_Ec2InstanceParameters(instance); err != nil {
		panic(err)
	}
	var returns VpcOriginEndpoint

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.VpcOriginEndpoint",
		"ec2Instance",
		[]interface{}{instance},
		&returns,
	)

	return returns
}

// A VPC origin endpoint from an Network Load Balancer.
func VpcOriginEndpoint_NetworkLoadBalancer(nlb awselasticloadbalancingv2.INetworkLoadBalancer) VpcOriginEndpoint {
	_init_.Initialize()

	if err := validateVpcOriginEndpoint_NetworkLoadBalancerParameters(nlb); err != nil {
		panic(err)
	}
	var returns VpcOriginEndpoint

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.VpcOriginEndpoint",
		"networkLoadBalancer",
		[]interface{}{nlb},
		&returns,
	)

	return returns
}

