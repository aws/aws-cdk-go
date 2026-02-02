package previewawsobservabilityadminmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsobservabilityadmin/previewawsobservabilityadminmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a telemetry pipeline for processing and transforming telemetry data.
//
// The pipeline defines how data flows from sources through processors to destinations, enabling data transformation and delivering capabilities.
//
// **Using CloudWatch as a pipeline source** The following is an example of a `Body` property value for the `Configuration` object. { "Type": "AWS::ObservabilityAdmin::TelemetryPipelines", "Properties": { "Configuration": { "Body": "pipeline:\n source:\n cloudwatch_logs:\n log_event_metadata:\n data_source_name: \"my_data_source\"\n data_source_type: \"default\"\n aws:\n sts_role_arn: \"arn:aws:iam::123456789012:role/MyPipelineAccessRole\"\n processor:\n - parse_json: {}\n sink:\n - cloudwatch_logs:\n log_group: \"@original\"" } }
// } Type: AWS::ObservabilityAdmin::TelemetryPipelines
// Properties: Configuration: Body: | pipeline: source: cloudwatch_logs: log_event_metadata: data_source_name: "my_data_source" data_source_type: "default" aws: sts_role_arn: "arn:aws:iam::123456789012:role/MyPipelineAccessRole" processor: - parse_json: {} sink: - cloudwatch_logs: log_group: "@original" **Using Amazon S3 as a pipeline source** The following is an example of a `Body` property value for the `Configuration` object. { "Type": "AWS::ObservabilityAdmin::TelemetryPipelines", "Properties": { "Configuration": { "Body": "pipeline:\n source:\n s3:\n sqs:\n visibility_timeout: \"PT60S\"\n visibility_duplication_protection: true\n maximum_messages: 10\n queue_url: \"https://sqs.us-east-1.amazonaws.com/123456789012/my-sqs-queue\"\n notification_type: \"sqs\"\n codec:\n ndjson: {}\n aws:\n region: \"us-east-1\"\n sts_role_arn: \"arn:aws:iam::123456789012:role/MyAccessRole\"\n data_source_name: \"crowdstrike_falcon\"\n processor:\n - ocsf:\n version: \"1.5\"\n mapping_version: \"1.5.0\"\n schema:\n crowdstrike_falcon:\n sink:\n - cloudwatch_logs:\n log_group: \"my-log-group\"" } }
// } Type: AWS::ObservabilityAdmin::TelemetryPipelines
// Properties: Configuration: Body: | pipeline: source: s3: sqs: visibility_timeout: "PT60S" visibility_duplication_protection: true maximum_messages: 10 queue_url: "https://sqs.us-east-1.amazonaws.com/123456789012/my-sqs-queue" notification_type: "sqs" codec: ndjson: {} aws: region: "us-east-1" sts_role_arn: "arn:aws:iam::123456789012:role/MyAccessRole" data_source_name: "crowdstrike_falcon" processor: - ocsf: version: "1.5" mapping_version: "1.5.0" schema: crowdstrike_falcon: sink: - cloudwatch_logs: log_group: "my-log-group"
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTelemetryPipelinesPropsMixin := awscdkmixinspreview.Mixins.NewCfnTelemetryPipelinesPropsMixin(&CfnTelemetryPipelinesMixinProps{
//   	Configuration: &TelemetryPipelineConfigurationProperty{
//   		Body: jsii.String("body"),
//   	},
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetrypipelines.html
//
type CfnTelemetryPipelinesPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTelemetryPipelinesMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTelemetryPipelinesPropsMixin
type jsiiProxy_CfnTelemetryPipelinesPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTelemetryPipelinesPropsMixin) Props() *CfnTelemetryPipelinesMixinProps {
	var returns *CfnTelemetryPipelinesMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTelemetryPipelinesPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ObservabilityAdmin::TelemetryPipelines`.
func NewCfnTelemetryPipelinesPropsMixin(props *CfnTelemetryPipelinesMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTelemetryPipelinesPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTelemetryPipelinesPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTelemetryPipelinesPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryPipelinesPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ObservabilityAdmin::TelemetryPipelines`.
func NewCfnTelemetryPipelinesPropsMixin_Override(c CfnTelemetryPipelinesPropsMixin, props *CfnTelemetryPipelinesMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryPipelinesPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTelemetryPipelinesPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTelemetryPipelinesPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryPipelinesPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTelemetryPipelinesPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryPipelinesPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTelemetryPipelinesPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTelemetryPipelinesPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

