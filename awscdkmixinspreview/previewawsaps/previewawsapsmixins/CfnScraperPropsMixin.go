package previewawsapsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsaps/previewawsapsmixins/internal"
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
//   cfnScraperPropsMixin := awscdkmixinspreview.Mixins.NewCfnScraperPropsMixin(&CfnScraperMixinProps{
//   	Alias: jsii.String("alias"),
//   	Destination: &DestinationProperty{
//   		AmpConfiguration: &AmpConfigurationProperty{
//   			WorkspaceArn: jsii.String("workspaceArn"),
//   		},
//   	},
//   	RoleConfiguration: &RoleConfigurationProperty{
//   		SourceRoleArn: jsii.String("sourceRoleArn"),
//   		TargetRoleArn: jsii.String("targetRoleArn"),
//   	},
//   	ScrapeConfiguration: &ScrapeConfigurationProperty{
//   		ConfigurationBlob: jsii.String("configurationBlob"),
//   	},
//   	ScraperLoggingConfiguration: &ScraperLoggingConfigurationProperty{
//   		LoggingDestination: &ScraperLoggingDestinationProperty{
//   			CloudWatchLogs: &CloudWatchLogDestinationProperty{
//   				LogGroupArn: jsii.String("logGroupArn"),
//   			},
//   		},
//   		ScraperComponents: []interface{}{
//   			&ScraperComponentProperty{
//   				Config: &ComponentConfigProperty{
//   					Options: map[string]*string{
//   						"optionsKey": jsii.String("options"),
//   					},
//   				},
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	Source: &SourceProperty{
//   		EksConfiguration: &EksConfigurationProperty{
//   			ClusterArn: jsii.String("clusterArn"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   		VpcConfiguration: &VpcConfigurationProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html
//
type CfnScraperPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnScraperMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnScraperPropsMixin
type jsiiProxy_CfnScraperPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnScraperPropsMixin) Props() *CfnScraperMixinProps {
	var returns *CfnScraperMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScraperPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::APS::Scraper`.
func NewCfnScraperPropsMixin(props *CfnScraperMixinProps, options *mixins.CfnPropertyMixinOptions) CfnScraperPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnScraperPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnScraperPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::APS::Scraper`.
func NewCfnScraperPropsMixin_Override(c CfnScraperPropsMixin, props *CfnScraperMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnScraperPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScraperPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScraperPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScraperPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnScraperPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

