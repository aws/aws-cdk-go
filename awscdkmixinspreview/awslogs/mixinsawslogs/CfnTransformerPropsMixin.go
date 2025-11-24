package mixinsawslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awslogs/mixinsawslogs/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates or updates a *log transformer* for a single log group.
//
// You use log transformers to transform log events into a different format, making them easier for you to process and analyze. You can also transform logs from different sources into standardized formats that contains relevant, source-specific information.
//
// After you have created a transformer, CloudWatch Logs performs the transformations at the time of log ingestion. You can then refer to the transformed versions of the logs during operations such as querying with CloudWatch Logs Insights or creating metric filters or subscription filers.
//
// You can also use a transformer to copy metadata from metadata keys into the log events themselves. This metadata can include log group name, log stream name, account ID and Region.
//
// A transformer for a log group is a series of processors, where each processor applies one type of transformation to the log events ingested into this log group. The processors work one after another, in the order that you list them, like a pipeline. For more information about the available processors to use in a transformer, see [Processors that you can use](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-Processors) .
//
// Having log events in standardized format enables visibility across your applications for your log analysis, reporting, and alarming needs. CloudWatch Logs provides transformation for common log types with out-of-the-box transformation templates for major AWS log sources such as VPC flow logs, Lambda, and Amazon RDS. You can use pre-built transformation templates or create custom transformation policies.
//
// You can create transformers only for the log groups in the Standard log class.
//
// You can also set up a transformer at the account level. For more information, see [PutAccountPolicy](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutAccountPolicy.html) . If there is both a log-group level transformer created with `PutTransformer` and an account-level transformer that could apply to the same log group, the log group uses only the log-group level transformer. It ignores the account-level transformer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransformerPropsMixin := awscdkmixinspreview.Mixins.NewCfnTransformerPropsMixin(&CfnTransformerMixinProps{
//   	LogGroupIdentifier: jsii.String("logGroupIdentifier"),
//   	TransformerConfig: []interface{}{
//   		&ProcessorProperty{
//   			AddKeys: &AddKeysProperty{
//   				Entries: []interface{}{
//   					&AddKeyEntryProperty{
//   						Key: jsii.String("key"),
//   						OverwriteIfExists: jsii.Boolean(false),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			CopyValue: &CopyValueProperty{
//   				Entries: []interface{}{
//   					&CopyValueEntryProperty{
//   						OverwriteIfExists: jsii.Boolean(false),
//   						Source: jsii.String("source"),
//   						Target: jsii.String("target"),
//   					},
//   				},
//   			},
//   			Csv: &CsvProperty{
//   				Columns: []*string{
//   					jsii.String("columns"),
//   				},
//   				Delimiter: jsii.String("delimiter"),
//   				QuoteCharacter: jsii.String("quoteCharacter"),
//   				Source: jsii.String("source"),
//   			},
//   			DateTimeConverter: &DateTimeConverterProperty{
//   				Locale: jsii.String("locale"),
//   				MatchPatterns: []*string{
//   					jsii.String("matchPatterns"),
//   				},
//   				Source: jsii.String("source"),
//   				SourceTimezone: jsii.String("sourceTimezone"),
//   				Target: jsii.String("target"),
//   				TargetFormat: jsii.String("targetFormat"),
//   				TargetTimezone: jsii.String("targetTimezone"),
//   			},
//   			DeleteKeys: &DeleteKeysProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   			Grok: &GrokProperty{
//   				Match: jsii.String("match"),
//   				Source: jsii.String("source"),
//   			},
//   			ListToMap: &ListToMapProperty{
//   				Flatten: jsii.Boolean(false),
//   				FlattenedElement: jsii.String("flattenedElement"),
//   				Key: jsii.String("key"),
//   				Source: jsii.String("source"),
//   				Target: jsii.String("target"),
//   				ValueKey: jsii.String("valueKey"),
//   			},
//   			LowerCaseString: &LowerCaseStringProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   			MoveKeys: &MoveKeysProperty{
//   				Entries: []interface{}{
//   					&MoveKeyEntryProperty{
//   						OverwriteIfExists: jsii.Boolean(false),
//   						Source: jsii.String("source"),
//   						Target: jsii.String("target"),
//   					},
//   				},
//   			},
//   			ParseCloudfront: &ParseCloudfrontProperty{
//   				Source: jsii.String("source"),
//   			},
//   			ParseJson: &ParseJSONProperty{
//   				Destination: jsii.String("destination"),
//   				Source: jsii.String("source"),
//   			},
//   			ParseKeyValue: &ParseKeyValueProperty{
//   				Destination: jsii.String("destination"),
//   				FieldDelimiter: jsii.String("fieldDelimiter"),
//   				KeyPrefix: jsii.String("keyPrefix"),
//   				KeyValueDelimiter: jsii.String("keyValueDelimiter"),
//   				NonMatchValue: jsii.String("nonMatchValue"),
//   				OverwriteIfExists: jsii.Boolean(false),
//   				Source: jsii.String("source"),
//   			},
//   			ParsePostgres: &ParsePostgresProperty{
//   				Source: jsii.String("source"),
//   			},
//   			ParseRoute53: &ParseRoute53Property{
//   				Source: jsii.String("source"),
//   			},
//   			ParseToOcsf: &ParseToOCSFProperty{
//   				EventSource: jsii.String("eventSource"),
//   				OcsfVersion: jsii.String("ocsfVersion"),
//   				Source: jsii.String("source"),
//   			},
//   			ParseVpc: &ParseVPCProperty{
//   				Source: jsii.String("source"),
//   			},
//   			ParseWaf: &ParseWAFProperty{
//   				Source: jsii.String("source"),
//   			},
//   			RenameKeys: &RenameKeysProperty{
//   				Entries: []interface{}{
//   					&RenameKeyEntryProperty{
//   						Key: jsii.String("key"),
//   						OverwriteIfExists: jsii.Boolean(false),
//   						RenameTo: jsii.String("renameTo"),
//   					},
//   				},
//   			},
//   			SplitString: &SplitStringProperty{
//   				Entries: []interface{}{
//   					&SplitStringEntryProperty{
//   						Delimiter: jsii.String("delimiter"),
//   						Source: jsii.String("source"),
//   					},
//   				},
//   			},
//   			SubstituteString: &SubstituteStringProperty{
//   				Entries: []interface{}{
//   					&SubstituteStringEntryProperty{
//   						From: jsii.String("from"),
//   						Source: jsii.String("source"),
//   						To: jsii.String("to"),
//   					},
//   				},
//   			},
//   			TrimString: &TrimStringProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   			TypeConverter: &TypeConverterProperty{
//   				Entries: []interface{}{
//   					&TypeConverterEntryProperty{
//   						Key: jsii.String("key"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			UpperCaseString: &UpperCaseStringProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-transformer.html
//
type CfnTransformerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTransformerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTransformerPropsMixin
type jsiiProxy_CfnTransformerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTransformerPropsMixin) Props() *CfnTransformerMixinProps {
	var returns *CfnTransformerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Logs::Transformer`.
func NewCfnTransformerPropsMixin(props *CfnTransformerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTransformerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTransformerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransformerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Logs::Transformer`.
func NewCfnTransformerPropsMixin_Override(c CfnTransformerPropsMixin, props *CfnTransformerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTransformerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransformerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTransformerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnTransformerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransformerPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTransformerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

