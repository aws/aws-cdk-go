package previewawsrummixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsrum/previewawsrummixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a CloudWatch RUM app monitor, which you can use to collect telemetry data from your application and send it to CloudWatch RUM.
//
// The data includes performance and reliability information such as page load time, client-side errors, and user behavior.
//
// After you create an app monitor, sign in to the CloudWatch RUM console to get the JavaScript code snippet to add to your web application. For more information, see [How do I find a code snippet that I've already generated?](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-find-code-snippet.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppMonitorPropsMixin := awscdkmixinspreview.Mixins.NewCfnAppMonitorPropsMixin(&CfnAppMonitorMixinProps{
//   	AppMonitorConfiguration: &AppMonitorConfigurationProperty{
//   		AllowCookies: jsii.Boolean(false),
//   		EnableXRay: jsii.Boolean(false),
//   		ExcludedPages: []*string{
//   			jsii.String("excludedPages"),
//   		},
//   		FavoritePages: []*string{
//   			jsii.String("favoritePages"),
//   		},
//   		GuestRoleArn: jsii.String("guestRoleArn"),
//   		IdentityPoolId: jsii.String("identityPoolId"),
//   		IncludedPages: []*string{
//   			jsii.String("includedPages"),
//   		},
//   		MetricDestinations: []interface{}{
//   			&MetricDestinationProperty{
//   				Destination: jsii.String("destination"),
//   				DestinationArn: jsii.String("destinationArn"),
//   				IamRoleArn: jsii.String("iamRoleArn"),
//   				MetricDefinitions: []interface{}{
//   					&MetricDefinitionProperty{
//   						DimensionKeys: map[string]*string{
//   							"dimensionKeysKey": jsii.String("dimensionKeys"),
//   						},
//   						EventPattern: jsii.String("eventPattern"),
//   						Name: jsii.String("name"),
//   						Namespace: jsii.String("namespace"),
//   						UnitLabel: jsii.String("unitLabel"),
//   						ValueKey: jsii.String("valueKey"),
//   					},
//   				},
//   			},
//   		},
//   		SessionSampleRate: jsii.Number(123),
//   		Telemetries: []*string{
//   			jsii.String("telemetries"),
//   		},
//   	},
//   	CustomEvents: &CustomEventsProperty{
//   		Status: jsii.String("status"),
//   	},
//   	CwLogEnabled: jsii.Boolean(false),
//   	DeobfuscationConfiguration: &DeobfuscationConfigurationProperty{
//   		JavaScriptSourceMaps: &JavaScriptSourceMapsProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	Domain: jsii.String("domain"),
//   	DomainList: []*string{
//   		jsii.String("domainList"),
//   	},
//   	Name: jsii.String("name"),
//   	Platform: jsii.String("platform"),
//   	ResourcePolicy: &ResourcePolicyProperty{
//   		PolicyDocument: jsii.String("policyDocument"),
//   		PolicyRevisionId: jsii.String("policyRevisionId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rum-appmonitor.html
//
type CfnAppMonitorPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAppMonitorMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAppMonitorPropsMixin
type jsiiProxy_CfnAppMonitorPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAppMonitorPropsMixin) Props() *CfnAppMonitorMixinProps {
	var returns *CfnAppMonitorMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitorPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RUM::AppMonitor`.
func NewCfnAppMonitorPropsMixin(props *CfnAppMonitorMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAppMonitorPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAppMonitorPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAppMonitorPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RUM::AppMonitor`.
func NewCfnAppMonitorPropsMixin_Override(c CfnAppMonitorPropsMixin, props *CfnAppMonitorMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAppMonitorPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAppMonitorPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAppMonitorPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAppMonitorPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAppMonitorPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

