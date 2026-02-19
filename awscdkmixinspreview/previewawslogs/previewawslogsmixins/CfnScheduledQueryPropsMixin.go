package previewawslogsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs/previewawslogsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new Scheduled Query that allows you to define a Logs Insights query that will run on a schedule and configure actions to take with the query results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnScheduledQueryPropsMixin := awscdkmixinspreview.Mixins.NewCfnScheduledQueryPropsMixin(&CfnScheduledQueryMixinProps{
//   	Description: jsii.String("description"),
//   	DestinationConfiguration: &DestinationConfigurationProperty{
//   		S3Configuration: &S3ConfigurationProperty{
//   			DestinationIdentifier: jsii.String("destinationIdentifier"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	LogGroupIdentifiers: []*string{
//   		jsii.String("logGroupIdentifiers"),
//   	},
//   	Name: jsii.String("name"),
//   	QueryLanguage: jsii.String("queryLanguage"),
//   	QueryString: jsii.String("queryString"),
//   	ScheduleEndTime: jsii.Number(123),
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   	ScheduleStartTime: jsii.Number(123),
//   	StartTimeOffset: jsii.Number(123),
//   	State: jsii.String("state"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Timezone: jsii.String("timezone"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html
//
type CfnScheduledQueryPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnScheduledQueryMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnScheduledQueryPropsMixin
type jsiiProxy_CfnScheduledQueryPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnScheduledQueryPropsMixin) Props() *CfnScheduledQueryMixinProps {
	var returns *CfnScheduledQueryMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQueryPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Logs::ScheduledQuery`.
func NewCfnScheduledQueryPropsMixin(props *CfnScheduledQueryMixinProps, options *mixins.CfnPropertyMixinOptions) CfnScheduledQueryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnScheduledQueryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnScheduledQueryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnScheduledQueryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Logs::ScheduledQuery`.
func NewCfnScheduledQueryPropsMixin_Override(c CfnScheduledQueryPropsMixin, props *CfnScheduledQueryMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnScheduledQueryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnScheduledQueryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScheduledQueryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnScheduledQueryPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScheduledQueryPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnScheduledQueryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScheduledQueryPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnScheduledQueryPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

