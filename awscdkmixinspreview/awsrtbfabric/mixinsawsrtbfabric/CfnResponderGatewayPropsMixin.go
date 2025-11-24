package mixinsawsrtbfabric

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsrtbfabric/mixinsawsrtbfabric/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a responder gateway.
//
// > A domain name or managed endpoint is required.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResponderGatewayPropsMixin := awscdkmixinspreview.Mixins.NewCfnResponderGatewayPropsMixin(&CfnResponderGatewayMixinProps{
//   	Description: jsii.String("description"),
//   	DomainName: jsii.String("domainName"),
//   	ManagedEndpointConfiguration: &ManagedEndpointConfigurationProperty{
//   		AutoScalingGroupsConfiguration: &AutoScalingGroupsConfigurationProperty{
//   			AutoScalingGroupNameList: []*string{
//   				jsii.String("autoScalingGroupNameList"),
//   			},
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		EksEndpointsConfiguration: &EksEndpointsConfigurationProperty{
//   			ClusterApiServerCaCertificateChain: jsii.String("clusterApiServerCaCertificateChain"),
//   			ClusterApiServerEndpointUri: jsii.String("clusterApiServerEndpointUri"),
//   			ClusterName: jsii.String("clusterName"),
//   			EndpointsResourceName: jsii.String("endpointsResourceName"),
//   			EndpointsResourceNamespace: jsii.String("endpointsResourceNamespace"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrustStoreConfiguration: &TrustStoreConfigurationProperty{
//   		CertificateAuthorityCertificates: []*string{
//   			jsii.String("certificateAuthorityCertificates"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-respondergateway.html
//
type CfnResponderGatewayPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnResponderGatewayMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResponderGatewayPropsMixin
type jsiiProxy_CfnResponderGatewayPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnResponderGatewayPropsMixin) Props() *CfnResponderGatewayMixinProps {
	var returns *CfnResponderGatewayMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponderGatewayPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RTBFabric::ResponderGateway`.
func NewCfnResponderGatewayPropsMixin(props *CfnResponderGatewayMixinProps, options *mixins.CfnPropertyMixinOptions) CfnResponderGatewayPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResponderGatewayPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResponderGatewayPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnResponderGatewayPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RTBFabric::ResponderGateway`.
func NewCfnResponderGatewayPropsMixin_Override(c CfnResponderGatewayPropsMixin, props *CfnResponderGatewayMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnResponderGatewayPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnResponderGatewayPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResponderGatewayPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnResponderGatewayPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResponderGatewayPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnResponderGatewayPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResponderGatewayPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResponderGatewayPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

