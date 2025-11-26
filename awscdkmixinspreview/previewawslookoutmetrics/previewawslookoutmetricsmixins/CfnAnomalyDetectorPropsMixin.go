package previewawslookoutmetricsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslookoutmetrics/previewawslookoutmetricsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// > End of support notice: On Oct 9, 2025, AWS will end support for Amazon Lookout for Metrics.
//
// After Oct 9, 2025, you will no longer be able to access the Amazon Lookout for Metrics console or Amazon Lookout for Metrics resources. For more information, see [Amazon Lookout for Metrics end of support](https://docs.aws.amazon.com//blogs/machine-learning/transitioning-off-amazon-lookout-for-metrics/) .
//
// The `AWS::LookoutMetrics::AnomalyDetector` type creates an anomaly detector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAnomalyDetectorPropsMixin := awscdkmixinspreview.Mixins.NewCfnAnomalyDetectorPropsMixin(&CfnAnomalyDetectorMixinProps{
//   	AnomalyDetectorConfig: &AnomalyDetectorConfigProperty{
//   		AnomalyDetectorFrequency: jsii.String("anomalyDetectorFrequency"),
//   	},
//   	AnomalyDetectorDescription: jsii.String("anomalyDetectorDescription"),
//   	AnomalyDetectorName: jsii.String("anomalyDetectorName"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	MetricSetList: []interface{}{
//   		&MetricSetProperty{
//   			DimensionList: []*string{
//   				jsii.String("dimensionList"),
//   			},
//   			MetricList: []interface{}{
//   				&MetricProperty{
//   					AggregationFunction: jsii.String("aggregationFunction"),
//   					MetricName: jsii.String("metricName"),
//   					Namespace: jsii.String("namespace"),
//   				},
//   			},
//   			MetricSetDescription: jsii.String("metricSetDescription"),
//   			MetricSetFrequency: jsii.String("metricSetFrequency"),
//   			MetricSetName: jsii.String("metricSetName"),
//   			MetricSource: &MetricSourceProperty{
//   				AppFlowConfig: &AppFlowConfigProperty{
//   					FlowName: jsii.String("flowName"),
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				CloudwatchConfig: &CloudwatchConfigProperty{
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				RdsSourceConfig: &RDSSourceConfigProperty{
//   					DatabaseHost: jsii.String("databaseHost"),
//   					DatabaseName: jsii.String("databaseName"),
//   					DatabasePort: jsii.Number(123),
//   					DbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   					RoleArn: jsii.String("roleArn"),
//   					SecretManagerArn: jsii.String("secretManagerArn"),
//   					TableName: jsii.String("tableName"),
//   					VpcConfiguration: &VpcConfigurationProperty{
//   						SecurityGroupIdList: []*string{
//   							jsii.String("securityGroupIdList"),
//   						},
//   						SubnetIdList: []*string{
//   							jsii.String("subnetIdList"),
//   						},
//   					},
//   				},
//   				RedshiftSourceConfig: &RedshiftSourceConfigProperty{
//   					ClusterIdentifier: jsii.String("clusterIdentifier"),
//   					DatabaseHost: jsii.String("databaseHost"),
//   					DatabaseName: jsii.String("databaseName"),
//   					DatabasePort: jsii.Number(123),
//   					RoleArn: jsii.String("roleArn"),
//   					SecretManagerArn: jsii.String("secretManagerArn"),
//   					TableName: jsii.String("tableName"),
//   					VpcConfiguration: &VpcConfigurationProperty{
//   						SecurityGroupIdList: []*string{
//   							jsii.String("securityGroupIdList"),
//   						},
//   						SubnetIdList: []*string{
//   							jsii.String("subnetIdList"),
//   						},
//   					},
//   				},
//   				S3SourceConfig: &S3SourceConfigProperty{
//   					FileFormatDescriptor: &FileFormatDescriptorProperty{
//   						CsvFormatDescriptor: &CsvFormatDescriptorProperty{
//   							Charset: jsii.String("charset"),
//   							ContainsHeader: jsii.Boolean(false),
//   							Delimiter: jsii.String("delimiter"),
//   							FileCompression: jsii.String("fileCompression"),
//   							HeaderList: []*string{
//   								jsii.String("headerList"),
//   							},
//   							QuoteSymbol: jsii.String("quoteSymbol"),
//   						},
//   						JsonFormatDescriptor: &JsonFormatDescriptorProperty{
//   							Charset: jsii.String("charset"),
//   							FileCompression: jsii.String("fileCompression"),
//   						},
//   					},
//   					HistoricalDataPathList: []*string{
//   						jsii.String("historicalDataPathList"),
//   					},
//   					RoleArn: jsii.String("roleArn"),
//   					TemplatedPathList: []*string{
//   						jsii.String("templatedPathList"),
//   					},
//   				},
//   			},
//   			Offset: jsii.Number(123),
//   			TimestampColumn: &TimestampColumnProperty{
//   				ColumnFormat: jsii.String("columnFormat"),
//   				ColumnName: jsii.String("columnName"),
//   			},
//   			Timezone: jsii.String("timezone"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-anomalydetector.html
//
type CfnAnomalyDetectorPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAnomalyDetectorMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAnomalyDetectorPropsMixin
type jsiiProxy_CfnAnomalyDetectorPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAnomalyDetectorPropsMixin) Props() *CfnAnomalyDetectorMixinProps {
	var returns *CfnAnomalyDetectorMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetectorPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::LookoutMetrics::AnomalyDetector`.
func NewCfnAnomalyDetectorPropsMixin(props *CfnAnomalyDetectorMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAnomalyDetectorPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAnomalyDetectorPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAnomalyDetectorPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::LookoutMetrics::AnomalyDetector`.
func NewCfnAnomalyDetectorPropsMixin_Override(c CfnAnomalyDetectorPropsMixin, props *CfnAnomalyDetectorMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAnomalyDetectorPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAnomalyDetectorPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAnomalyDetectorPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAnomalyDetectorPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAnomalyDetectorPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

