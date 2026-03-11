package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AWS Verified Access endpoint specifies the application that AWS Verified Access provides access to.
//
// It must be attached to an AWS Verified Access group. An AWS Verified Access endpoint must also have an attached access policy before you attached it to a group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnVerifiedAccessEndpointPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnVerifiedAccessEndpointPropsMixin(&CfnVerifiedAccessEndpointMixinProps{
//   	ApplicationDomain: jsii.String("applicationDomain"),
//   	AttachmentType: jsii.String("attachmentType"),
//   	CidrOptions: &CidrOptionsProperty{
//   		Cidr: jsii.String("cidr"),
//   		PortRanges: []interface{}{
//   			&PortRangeProperty{
//   				FromPort: jsii.Number(123),
//   				ToPort: jsii.Number(123),
//   			},
//   		},
//   		Protocol: jsii.String("protocol"),
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	DomainCertificateArn: jsii.String("domainCertificateArn"),
//   	EndpointDomainPrefix: jsii.String("endpointDomainPrefix"),
//   	EndpointType: jsii.String("endpointType"),
//   	LoadBalancerOptions: &LoadBalancerOptionsProperty{
//   		LoadBalancerArn: jsii.String("loadBalancerArn"),
//   		Port: jsii.Number(123),
//   		PortRanges: []interface{}{
//   			&PortRangeProperty{
//   				FromPort: jsii.Number(123),
//   				ToPort: jsii.Number(123),
//   			},
//   		},
//   		Protocol: jsii.String("protocol"),
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	NetworkInterfaceOptions: &NetworkInterfaceOptionsProperty{
//   		NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   		Port: jsii.Number(123),
//   		PortRanges: []interface{}{
//   			&PortRangeProperty{
//   				FromPort: jsii.Number(123),
//   				ToPort: jsii.Number(123),
//   			},
//   		},
//   		Protocol: jsii.String("protocol"),
//   	},
//   	PolicyDocument: jsii.String("policyDocument"),
//   	PolicyEnabled: jsii.Boolean(false),
//   	RdsOptions: &RdsOptionsProperty{
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		RdsDbClusterArn: jsii.String("rdsDbClusterArn"),
//   		RdsDbInstanceArn: jsii.String("rdsDbInstanceArn"),
//   		RdsDbProxyArn: jsii.String("rdsDbProxyArn"),
//   		RdsEndpoint: jsii.String("rdsEndpoint"),
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SseSpecification: &SseSpecificationProperty{
//   		CustomerManagedKeyEnabled: jsii.Boolean(false),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VerifiedAccessGroupId: jsii.String("verifiedAccessGroupId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html
//
type CfnVerifiedAccessEndpointPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnVerifiedAccessEndpointMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVerifiedAccessEndpointPropsMixin
type jsiiProxy_CfnVerifiedAccessEndpointPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnVerifiedAccessEndpointPropsMixin) Props() *CfnVerifiedAccessEndpointMixinProps {
	var returns *CfnVerifiedAccessEndpointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpointPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::VerifiedAccessEndpoint`.
func NewCfnVerifiedAccessEndpointPropsMixin(props *CfnVerifiedAccessEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnVerifiedAccessEndpointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVerifiedAccessEndpointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVerifiedAccessEndpointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnVerifiedAccessEndpointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::VerifiedAccessEndpoint`.
func NewCfnVerifiedAccessEndpointPropsMixin_Override(c CfnVerifiedAccessEndpointPropsMixin, props *CfnVerifiedAccessEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnVerifiedAccessEndpointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnVerifiedAccessEndpointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVerifiedAccessEndpointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnVerifiedAccessEndpointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVerifiedAccessEndpointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnVerifiedAccessEndpointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVerifiedAccessEndpointPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnVerifiedAccessEndpointPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

