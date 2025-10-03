package awscloudfrontorigins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a distribution's VPC origin.
//
// Example:
//   // Creates a distribution from a Network Load Balancer
//   var vpc vpc
//
//   // Create a network load balancer in a VPC. 'internetFacing' can be 'false'.
//   nlb := elbv2.NewNetworkLoadBalancer(this, jsii.String("NLB"), &NetworkLoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(false),
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
//   	},
//   	SecurityGroups: []iSecurityGroup{
//   		ec2.NewSecurityGroup(this, jsii.String("NLB-SG"), &SecurityGroupProps{
//   			Vpc: *Vpc,
//   		}),
//   	},
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.VpcOrigin_WithNetworkLoadBalancer(nlb),
//   	},
//   })
//
type VpcOrigin interface {
	awscloudfront.OriginBase
	Props() *VpcOriginProps
	VpcOrigin() awscloudfront.IVpcOrigin
	SetVpcOrigin(val awscloudfront.IVpcOrigin)
	// Binds the origin to the associated Distribution.
	//
	// Can be used to grant permissions, create dependent resources, etc.
	Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
	RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty
	RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty
	RenderVpcOriginConfig() *awscloudfront.CfnDistribution_VpcOriginConfigProperty
}

// The jsii proxy struct for VpcOrigin
type jsiiProxy_VpcOrigin struct {
	internal.Type__awscloudfrontOriginBase
}

func (j *jsiiProxy_VpcOrigin) Props() *VpcOriginProps {
	var returns *VpcOriginProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcOrigin) VpcOrigin() awscloudfront.IVpcOrigin {
	var returns awscloudfront.IVpcOrigin
	_jsii_.Get(
		j,
		"vpcOrigin",
		&returns,
	)
	return returns
}


func NewVpcOrigin_Override(v VpcOrigin, domainName *string, props *VpcOriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.VpcOrigin",
		[]interface{}{domainName, props},
		v,
	)
}

func (j *jsiiProxy_VpcOrigin)SetVpcOrigin(val awscloudfront.IVpcOrigin) {
	_jsii_.Set(
		j,
		"vpcOrigin",
		val,
	)
}

// Create a VPC origin with an Application Load Balancer.
func VpcOrigin_WithApplicationLoadBalancer(alb awselasticloadbalancingv2.IApplicationLoadBalancer, props *VpcOriginWithEndpointProps) VpcOrigin {
	_init_.Initialize()

	if err := validateVpcOrigin_WithApplicationLoadBalancerParameters(alb, props); err != nil {
		panic(err)
	}
	var returns VpcOrigin

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront_origins.VpcOrigin",
		"withApplicationLoadBalancer",
		[]interface{}{alb, props},
		&returns,
	)

	return returns
}

// Create a VPC origin with an EC2 instance.
func VpcOrigin_WithEc2Instance(instance awsec2.IInstance, props *VpcOriginWithEndpointProps) VpcOrigin {
	_init_.Initialize()

	if err := validateVpcOrigin_WithEc2InstanceParameters(instance, props); err != nil {
		panic(err)
	}
	var returns VpcOrigin

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront_origins.VpcOrigin",
		"withEc2Instance",
		[]interface{}{instance, props},
		&returns,
	)

	return returns
}

// Create a VPC origin with a Network Load Balancer.
func VpcOrigin_WithNetworkLoadBalancer(nlb awselasticloadbalancingv2.INetworkLoadBalancer, props *VpcOriginWithEndpointProps) VpcOrigin {
	_init_.Initialize()

	if err := validateVpcOrigin_WithNetworkLoadBalancerParameters(nlb, props); err != nil {
		panic(err)
	}
	var returns VpcOrigin

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront_origins.VpcOrigin",
		"withNetworkLoadBalancer",
		[]interface{}{nlb, props},
		&returns,
	)

	return returns
}

// Create a VPC origin with an existing VPC origin resource.
func VpcOrigin_WithVpcOrigin(origin awscloudfront.IVpcOrigin, props *VpcOriginProps) VpcOrigin {
	_init_.Initialize()

	if err := validateVpcOrigin_WithVpcOriginParameters(origin, props); err != nil {
		panic(err)
	}
	var returns VpcOrigin

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront_origins.VpcOrigin",
		"withVpcOrigin",
		[]interface{}{origin, props},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcOrigin) Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	if err := v.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		v,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcOrigin) RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_CustomOriginConfigProperty

	_jsii_.Invoke(
		v,
		"renderCustomOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcOrigin) RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_S3OriginConfigProperty

	_jsii_.Invoke(
		v,
		"renderS3OriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcOrigin) RenderVpcOriginConfig() *awscloudfront.CfnDistribution_VpcOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_VpcOriginConfigProperty

	_jsii_.Invoke(
		v,
		"renderVpcOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

