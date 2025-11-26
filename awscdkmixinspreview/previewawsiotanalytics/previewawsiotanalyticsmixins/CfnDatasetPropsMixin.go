package previewawsiotanalyticsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiotanalytics/previewawsiotanalyticsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::IoTAnalytics::Dataset resource stores data retrieved from a data store by applying a `queryAction` (an SQL query) or a `containerAction` (executing a containerized application).
//
// The data set can be populated manually by calling `CreateDatasetContent` or automatically according to a `trigger` you specify. For more information, see [How to Use AWS IoT Analytics](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how) in the *AWS IoT Analytics User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDatasetPropsMixin := awscdkmixinspreview.Mixins.NewCfnDatasetPropsMixin(&CfnDatasetMixinProps{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			ActionName: jsii.String("actionName"),
//   			ContainerAction: &ContainerActionProperty{
//   				ExecutionRoleArn: jsii.String("executionRoleArn"),
//   				Image: jsii.String("image"),
//   				ResourceConfiguration: &ResourceConfigurationProperty{
//   					ComputeType: jsii.String("computeType"),
//   					VolumeSizeInGb: jsii.Number(123),
//   				},
//   				Variables: []interface{}{
//   					&VariableProperty{
//   						DatasetContentVersionValue: &DatasetContentVersionValueProperty{
//   							DatasetName: jsii.String("datasetName"),
//   						},
//   						DoubleValue: jsii.Number(123),
//   						OutputFileUriValue: &OutputFileUriValueProperty{
//   							FileName: jsii.String("fileName"),
//   						},
//   						StringValue: jsii.String("stringValue"),
//   						VariableName: jsii.String("variableName"),
//   					},
//   				},
//   			},
//   			QueryAction: &QueryActionProperty{
//   				Filters: []interface{}{
//   					&FilterProperty{
//   						DeltaTime: &DeltaTimeProperty{
//   							OffsetSeconds: jsii.Number(123),
//   							TimeExpression: jsii.String("timeExpression"),
//   						},
//   					},
//   				},
//   				SqlQuery: jsii.String("sqlQuery"),
//   			},
//   		},
//   	},
//   	ContentDeliveryRules: []interface{}{
//   		&DatasetContentDeliveryRuleProperty{
//   			Destination: &DatasetContentDeliveryRuleDestinationProperty{
//   				IotEventsDestinationConfiguration: &IotEventsDestinationConfigurationProperty{
//   					InputName: jsii.String("inputName"),
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				S3DestinationConfiguration: &S3DestinationConfigurationProperty{
//   					Bucket: jsii.String("bucket"),
//   					GlueConfiguration: &GlueConfigurationProperty{
//   						DatabaseName: jsii.String("databaseName"),
//   						TableName: jsii.String("tableName"),
//   					},
//   					Key: jsii.String("key"),
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   			},
//   			EntryName: jsii.String("entryName"),
//   		},
//   	},
//   	DatasetName: jsii.String("datasetName"),
//   	LateDataRules: []interface{}{
//   		&LateDataRuleProperty{
//   			RuleConfiguration: &LateDataRuleConfigurationProperty{
//   				DeltaTimeSessionWindowConfiguration: &DeltaTimeSessionWindowConfigurationProperty{
//   					TimeoutInMinutes: jsii.Number(123),
//   				},
//   			},
//   			RuleName: jsii.String("ruleName"),
//   		},
//   	},
//   	RetentionPeriod: &RetentionPeriodProperty{
//   		NumberOfDays: jsii.Number(123),
//   		Unlimited: jsii.Boolean(false),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Triggers: []interface{}{
//   		&TriggerProperty{
//   			Schedule: &ScheduleProperty{
//   				ScheduleExpression: jsii.String("scheduleExpression"),
//   			},
//   			TriggeringDataset: &TriggeringDatasetProperty{
//   				DatasetName: jsii.String("datasetName"),
//   			},
//   		},
//   	},
//   	VersioningConfiguration: &VersioningConfigurationProperty{
//   		MaxVersions: jsii.Number(123),
//   		Unlimited: jsii.Boolean(false),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-dataset.html
//
type CfnDatasetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDatasetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDatasetPropsMixin
type jsiiProxy_CfnDatasetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDatasetPropsMixin) Props() *CfnDatasetMixinProps {
	var returns *CfnDatasetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTAnalytics::Dataset`.
func NewCfnDatasetPropsMixin(props *CfnDatasetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDatasetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDatasetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDatasetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTAnalytics::Dataset`.
func NewCfnDatasetPropsMixin_Override(c CfnDatasetPropsMixin, props *CfnDatasetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDatasetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDatasetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDatasetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatasetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDatasetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDatasetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

