package previewawsroute53mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsroute53/previewawsroute53mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Route53::HealthCheck` resource is a Route 53 resource type that contains settings for a Route 53 health check.
//
// For information about associating health checks with records, see [HealthCheckId](https://docs.aws.amazon.com/Route53/latest/APIReference/API_ResourceRecordSet.html#Route53-Type-ResourceRecordSet-HealthCheckId) in [ChangeResourceRecordSets](https://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html) .
//
// > You can't create a health check with simple routing.
//
// *ELB Load Balancers*
//
// If you're registering EC2 instances with an Elastic Load Balancing (ELB) load balancer, do not create Amazon Route 53 health checks for the EC2 instances. When you register an EC2 instance with a load balancer, you configure settings for an ELB health check, which performs a similar function to a Route 53 health check.
//
// *Private Hosted Zones*
//
// You can associate health checks with failover records in a private hosted zone. Note the following:
//
// - Route 53 health checkers are outside the VPC. To check the health of an endpoint within a VPC by IP address, you must assign a public IP address to the instance in the VPC.
// - You can configure a health checker to check the health of an external resource that the instance relies on, such as a database server.
// - You can create a CloudWatch metric, associate an alarm with the metric, and then create a health check that is based on the state of the alarm. For example, you might create a CloudWatch metric that checks the status of the Amazon EC2 `StatusCheckFailed` metric, add an alarm to the metric, and then create a health check that is based on the state of the alarm. For information about creating CloudWatch metrics and alarms by using the CloudWatch console, see the [Amazon CloudWatch User Guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/WhatIsCloudWatch.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHealthCheckPropsMixin := awscdkmixinspreview.Mixins.NewCfnHealthCheckPropsMixin(&CfnHealthCheckMixinProps{
//   	HealthCheckConfig: &HealthCheckConfigProperty{
//   		AlarmIdentifier: &AlarmIdentifierProperty{
//   			Name: jsii.String("name"),
//   			Region: jsii.String("region"),
//   		},
//   		ChildHealthChecks: []*string{
//   			jsii.String("childHealthChecks"),
//   		},
//   		EnableSni: jsii.Boolean(false),
//   		FailureThreshold: jsii.Number(123),
//   		FullyQualifiedDomainName: jsii.String("fullyQualifiedDomainName"),
//   		HealthThreshold: jsii.Number(123),
//   		InsufficientDataHealthStatus: jsii.String("insufficientDataHealthStatus"),
//   		Inverted: jsii.Boolean(false),
//   		IpAddress: jsii.String("ipAddress"),
//   		MeasureLatency: jsii.Boolean(false),
//   		Port: jsii.Number(123),
//   		Regions: []*string{
//   			jsii.String("regions"),
//   		},
//   		RequestInterval: jsii.Number(123),
//   		ResourcePath: jsii.String("resourcePath"),
//   		RoutingControlArn: jsii.String("routingControlArn"),
//   		SearchString: jsii.String("searchString"),
//   		Type: jsii.String("type"),
//   	},
//   	HealthCheckTags: []HealthCheckTagProperty{
//   		&HealthCheckTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html
//
type CfnHealthCheckPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnHealthCheckMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnHealthCheckPropsMixin
type jsiiProxy_CfnHealthCheckPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnHealthCheckPropsMixin) Props() *CfnHealthCheckMixinProps {
	var returns *CfnHealthCheckMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHealthCheckPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53::HealthCheck`.
func NewCfnHealthCheckPropsMixin(props *CfnHealthCheckMixinProps, options *mixins.CfnPropertyMixinOptions) CfnHealthCheckPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnHealthCheckPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnHealthCheckPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnHealthCheckPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53::HealthCheck`.
func NewCfnHealthCheckPropsMixin_Override(c CfnHealthCheckPropsMixin, props *CfnHealthCheckMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnHealthCheckPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnHealthCheckPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnHealthCheckPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnHealthCheckPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHealthCheckPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnHealthCheckPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHealthCheckPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnHealthCheckPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

