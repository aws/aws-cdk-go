package mixinsawsinternetmonitor

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsinternetmonitor/mixinsawsinternetmonitor/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::InternetMonitor::Monitor` resource is an Internet Monitor resource type that contains information about how you create a monitor in Amazon CloudWatch Internet Monitor.
//
// A monitor in Internet Monitor provides visibility into performance and availability between your applications hosted on AWS and your end users, using a traffic profile that it creates based on the application resources that you add: Virtual Private Clouds (VPCs), Amazon CloudFront distributions, or WorkSpaces directories.
//
// Internet Monitor also alerts you to internet issues that impact your application in the city-networks (geographies and networks) where your end users use it. With Internet Monitor, you can quickly pinpoint the locations and providers that are affected, so that you can address the issue.
//
// For more information, see [Using Amazon CloudWatch Internet Monitor](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-InternetMonitor.html) in the *Amazon CloudWatch User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMonitorPropsMixin := awscdkmixinspreview.Mixins.NewCfnMonitorPropsMixin(&CfnMonitorMixinProps{
//   	HealthEventsConfig: &HealthEventsConfigProperty{
//   		AvailabilityLocalHealthEventsConfig: &LocalHealthEventsConfigProperty{
//   			HealthScoreThreshold: jsii.Number(123),
//   			MinTrafficImpact: jsii.Number(123),
//   			Status: jsii.String("status"),
//   		},
//   		AvailabilityScoreThreshold: jsii.Number(123),
//   		PerformanceLocalHealthEventsConfig: &LocalHealthEventsConfigProperty{
//   			HealthScoreThreshold: jsii.Number(123),
//   			MinTrafficImpact: jsii.Number(123),
//   			Status: jsii.String("status"),
//   		},
//   		PerformanceScoreThreshold: jsii.Number(123),
//   	},
//   	IncludeLinkedAccounts: jsii.Boolean(false),
//   	InternetMeasurementsLogDelivery: &InternetMeasurementsLogDeliveryProperty{
//   		S3Config: &S3ConfigProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			LogDeliveryStatus: jsii.String("logDeliveryStatus"),
//   		},
//   	},
//   	LinkedAccountId: jsii.String("linkedAccountId"),
//   	MaxCityNetworksToMonitor: jsii.Number(123),
//   	MonitorName: jsii.String("monitorName"),
//   	Resources: []*string{
//   		jsii.String("resources"),
//   	},
//   	ResourcesToAdd: []*string{
//   		jsii.String("resourcesToAdd"),
//   	},
//   	ResourcesToRemove: []*string{
//   		jsii.String("resourcesToRemove"),
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrafficPercentageToMonitor: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-internetmonitor-monitor.html
//
type CfnMonitorPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMonitorMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMonitorPropsMixin
type jsiiProxy_CfnMonitorPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMonitorPropsMixin) Props() *CfnMonitorMixinProps {
	var returns *CfnMonitorMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitorPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::InternetMonitor::Monitor`.
func NewCfnMonitorPropsMixin(props *CfnMonitorMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMonitorPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMonitorPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMonitorPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_internetmonitor.mixins.CfnMonitorPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::InternetMonitor::Monitor`.
func NewCfnMonitorPropsMixin_Override(c CfnMonitorPropsMixin, props *CfnMonitorMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_internetmonitor.mixins.CfnMonitorPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMonitorPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMonitorPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_internetmonitor.mixins.CfnMonitorPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMonitorPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_internetmonitor.mixins.CfnMonitorPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMonitorPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMonitorPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

