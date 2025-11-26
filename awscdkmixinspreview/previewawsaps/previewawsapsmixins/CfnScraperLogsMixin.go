package previewawsapsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsaps/previewawsapsmixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// A scraper is a fully-managed agentless collector that discovers and pulls metrics automatically.
//
// A scraper pulls metrics from Prometheus-compatible sources within an Amazon EKS cluster, and sends them to your Amazon Managed Service for Prometheus workspace. Scrapers are flexible. You can configure the scraper to control what metrics are collected, the frequency of collection, what transformations are applied to the metrics, and more.
//
// An IAM role will be created for you that Amazon Managed Service for Prometheus uses to access the metrics in your cluster. You must configure this role with a policy that allows it to scrape metrics from your cluster. For more information, see [Configuring your Amazon EKS cluster](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-collector-how-to.html#AMP-collector-eks-setup) in the *Amazon Managed Service for Prometheus User Guide* .
//
// The `scrapeConfiguration` parameter contains the YAML configuration for the scraper.
//
// > For more information about collectors, including what metrics are collected, and how to configure the scraper, see [Using an AWS managed collector](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-collector-how-to.html) in the *Amazon Managed Service for Prometheus User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnScraperLogsMixin := awscdkmixinspreview.Mixins.NewCfnScraperLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html
//
type CfnScraperLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnScraperLogsMixin
type jsiiProxy_CfnScraperLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnScraperLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScraperLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::APS::Scraper`.
func NewCfnScraperLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnScraperLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnScraperLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnScraperLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::APS::Scraper`.
func NewCfnScraperLogsMixin_Override(c CfnScraperLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnScraperLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScraperLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScraperLogsMixin_APPLICATION_LOGS() CfnScraperApplicationLogs {
	_init_.Initialize()
	var returns CfnScraperApplicationLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperLogsMixin",
		"APPLICATION_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScraperLogsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnScraperLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

