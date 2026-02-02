package previewawscloudtrailmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudtrail/previewawscloudtrailmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a trail that specifies the settings for delivery of log data to an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTrailPropsMixin := awscdkmixinspreview.Mixins.NewCfnTrailPropsMixin(&CfnTrailMixinProps{
//   	AdvancedEventSelectors: []interface{}{
//   		&AdvancedEventSelectorProperty{
//   			FieldSelectors: []interface{}{
//   				&AdvancedFieldSelectorProperty{
//   					EndsWith: []*string{
//   						jsii.String("endsWith"),
//   					},
//   					EqualTo: []*string{
//   						jsii.String("equalTo"),
//   					},
//   					Field: jsii.String("field"),
//   					NotEndsWith: []*string{
//   						jsii.String("notEndsWith"),
//   					},
//   					NotEquals: []*string{
//   						jsii.String("notEquals"),
//   					},
//   					NotStartsWith: []*string{
//   						jsii.String("notStartsWith"),
//   					},
//   					StartsWith: []*string{
//   						jsii.String("startsWith"),
//   					},
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	AggregationConfigurations: []interface{}{
//   		&AggregationConfigurationProperty{
//   			EventCategory: jsii.String("eventCategory"),
//   			Templates: []*string{
//   				jsii.String("templates"),
//   			},
//   		},
//   	},
//   	CloudWatchLogsLogGroupArn: jsii.String("cloudWatchLogsLogGroupArn"),
//   	CloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   	EnableLogFileValidation: jsii.Boolean(false),
//   	EventSelectors: []interface{}{
//   		&EventSelectorProperty{
//   			DataResources: []interface{}{
//   				&DataResourceProperty{
//   					Type: jsii.String("type"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			ExcludeManagementEventSources: []*string{
//   				jsii.String("excludeManagementEventSources"),
//   			},
//   			IncludeManagementEvents: jsii.Boolean(false),
//   			ReadWriteType: jsii.String("readWriteType"),
//   		},
//   	},
//   	IncludeGlobalServiceEvents: jsii.Boolean(false),
//   	InsightSelectors: []interface{}{
//   		&InsightSelectorProperty{
//   			EventCategories: []*string{
//   				jsii.String("eventCategories"),
//   			},
//   			InsightType: jsii.String("insightType"),
//   		},
//   	},
//   	IsLogging: jsii.Boolean(false),
//   	IsMultiRegionTrail: jsii.Boolean(false),
//   	IsOrganizationTrail: jsii.Boolean(false),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	S3BucketName: jsii.String("s3BucketName"),
//   	S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	SnsTopicName: jsii.String("snsTopicName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrailName: jsii.String("trailName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html
//
type CfnTrailPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTrailMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTrailPropsMixin
type jsiiProxy_CfnTrailPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTrailPropsMixin) Props() *CfnTrailMixinProps {
	var returns *CfnTrailMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrailPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudTrail::Trail`.
func NewCfnTrailPropsMixin(props *CfnTrailMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTrailPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTrailPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTrailPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudtrail.mixins.CfnTrailPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudTrail::Trail`.
func NewCfnTrailPropsMixin_Override(c CfnTrailPropsMixin, props *CfnTrailMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudtrail.mixins.CfnTrailPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTrailPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTrailPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudtrail.mixins.CfnTrailPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTrailPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudtrail.mixins.CfnTrailPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTrailPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTrailPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

