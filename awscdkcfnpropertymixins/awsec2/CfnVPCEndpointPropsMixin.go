package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a VPC endpoint.
//
// A VPC endpoint provides a private connection between your VPC and an endpoint service. You can use an endpoint service provided by AWS , an AWS Marketplace Partner, or another AWS accounts in your organization. For more information, see the [AWS PrivateLink User Guide](https://docs.aws.amazon.com/vpc/latest/privatelink/) .
//
// An endpoint of type `Interface` establishes connections between the subnets in your VPC and an AWS service , your own service, or a service hosted by another AWS account . With an interface VPC endpoint, you specify the subnets in which to create the endpoint and the security groups to associate with the endpoint network interfaces.
//
// An endpoint of type `gateway` serves as a target for a route in your route table for traffic destined for Amazon S3 or DynamoDB . You can specify an endpoint policy for the endpoint, which controls access to the service from your VPC. You can also specify the VPC route tables that use the endpoint. For more information about connectivity to Amazon S3 , see [Why can't I connect to an S3 bucket using a gateway VPC endpoint?](https://docs.aws.amazon.com/premiumsupport/knowledge-center/connect-s3-vpc-endpoint)
//
// An endpoint of type `GatewayLoadBalancer` provides private connectivity between your VPC and virtual appliances from a service provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var policyDocument interface{}
//
//   cfnVPCEndpointPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnVPCEndpointPropsMixin(&CfnVPCEndpointMixinProps{
//   	DnsOptions: &DnsOptionsSpecificationProperty{
//   		DnsRecordIpType: jsii.String("dnsRecordIpType"),
//   		PrivateDnsOnlyForInboundResolverEndpoint: jsii.String("privateDnsOnlyForInboundResolverEndpoint"),
//   		PrivateDnsPreference: jsii.String("privateDnsPreference"),
//   		PrivateDnsSpecifiedDomains: []*string{
//   			jsii.String("privateDnsSpecifiedDomains"),
//   		},
//   	},
//   	IpAddressType: jsii.String("ipAddressType"),
//   	PolicyDocument: policyDocument,
//   	PrivateDnsEnabled: jsii.Boolean(false),
//   	ResourceConfigurationArn: jsii.String("resourceConfigurationArn"),
//   	RouteTableIds: []interface{}{
//   		jsii.String("routeTableIds"),
//   	},
//   	SecurityGroupIds: []interface{}{
//   		jsii.String("securityGroupIds"),
//   	},
//   	ServiceName: jsii.String("serviceName"),
//   	ServiceNetworkArn: jsii.String("serviceNetworkArn"),
//   	ServiceRegion: jsii.String("serviceRegion"),
//   	SubnetIds: []interface{}{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcEndpointType: jsii.String("vpcEndpointType"),
//   	VpcId: jsii.String("vpcId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html
//
type CfnVPCEndpointPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnVPCEndpointMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVPCEndpointPropsMixin
type jsiiProxy_CfnVPCEndpointPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnVPCEndpointPropsMixin) Props() *CfnVPCEndpointMixinProps {
	var returns *CfnVPCEndpointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEndpointPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::VPCEndpoint`.
func NewCfnVPCEndpointPropsMixin(props *CfnVPCEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnVPCEndpointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVPCEndpointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVPCEndpointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnVPCEndpointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::VPCEndpoint`.
func NewCfnVPCEndpointPropsMixin_Override(c CfnVPCEndpointPropsMixin, props *CfnVPCEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnVPCEndpointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnVPCEndpointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVPCEndpointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnVPCEndpointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVPCEndpointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnVPCEndpointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVPCEndpointPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnVPCEndpointPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

