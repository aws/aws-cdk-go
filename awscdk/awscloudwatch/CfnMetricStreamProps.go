package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMetricStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMetricStreamProps := &CfnMetricStreamProps{
//   	FirehoseArn: jsii.String("firehoseArn"),
//   	OutputFormat: jsii.String("outputFormat"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	ExcludeFilters: []interface{}{
//   		&MetricStreamFilterProperty{
//   			Namespace: jsii.String("namespace"),
//
//   			// the properties below are optional
//   			MetricNames: []*string{
//   				jsii.String("metricNames"),
//   			},
//   		},
//   	},
//   	IncludeFilters: []interface{}{
//   		&MetricStreamFilterProperty{
//   			Namespace: jsii.String("namespace"),
//
//   			// the properties below are optional
//   			MetricNames: []*string{
//   				jsii.String("metricNames"),
//   			},
//   		},
//   	},
//   	IncludeLinkedAccountsMetrics: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	StatisticsConfigurations: []interface{}{
//   		&MetricStreamStatisticsConfigurationProperty{
//   			AdditionalStatistics: []*string{
//   				jsii.String("additionalStatistics"),
//   			},
//   			IncludeMetrics: []interface{}{
//   				&MetricStreamStatisticsMetricProperty{
//   					MetricName: jsii.String("metricName"),
//   					Namespace: jsii.String("namespace"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricstream.html
//
type CfnMetricStreamProps struct {
	// The ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
	//
	// This Amazon Kinesis Firehose delivery stream must already exist and must be in the same account as the metric stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricstream.html#cfn-cloudwatch-metricstream-firehosearn
	//
	FirehoseArn *string `field:"required" json:"firehoseArn" yaml:"firehoseArn"`
	// The output format for the stream.
	//
	// Valid values are `json` , `opentelemetry1.0` and `opentelemetry0.7` For more information about metric stream output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html) .
	//
	// This parameter is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricstream.html#cfn-cloudwatch-metricstream-outputformat
	//
	OutputFormat *string `field:"required" json:"outputFormat" yaml:"outputFormat"`
	// The ARN of an IAM role that this metric stream will use to access Amazon Kinesis Firehose resources.
	//
	// This IAM role must already exist and must be in the same account as the metric stream. This IAM role must include the `firehose:PutRecord` and `firehose:PutRecordBatch` permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricstream.html#cfn-cloudwatch-metricstream-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces that you specify here.
	//
	// You cannot specify both `IncludeFilters` and `ExcludeFilters` in the same metric stream.
	//
	// When you modify the `IncludeFilters` or `ExcludeFilters` of an existing metric stream in any way, the metric stream is effectively restarted, so after such a change you will get only the datapoints that have a timestamp after the time of the update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricstream.html#cfn-cloudwatch-metricstream-excludefilters
	//
	ExcludeFilters interface{} `field:"optional" json:"excludeFilters" yaml:"excludeFilters"`
	// If you specify this parameter, the stream sends only the metrics from the metric namespaces that you specify here.
	//
	// You cannot specify both `IncludeFilters` and `ExcludeFilters` in the same metric stream.
	//
	// When you modify the `IncludeFilters` or `ExcludeFilters` of an existing metric stream in any way, the metric stream is effectively restarted, so after such a change you will get only the datapoints that have a timestamp after the time of the update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricstream.html#cfn-cloudwatch-metricstream-includefilters
	//
	IncludeFilters interface{} `field:"optional" json:"includeFilters" yaml:"includeFilters"`
	// If you are creating a metric stream in a monitoring account, specify `true` to include metrics from source accounts that are linked to this monitoring account, in the metric stream.
	//
	// The default is `false` .
	//
	// For more information about linking accounts, see [CloudWatch cross-account observability](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricstream.html#cfn-cloudwatch-metricstream-includelinkedaccountsmetrics
	//
	IncludeLinkedAccountsMetrics interface{} `field:"optional" json:"includeLinkedAccountsMetrics" yaml:"includeLinkedAccountsMetrics"`
	// If you are creating a new metric stream, this is the name for the new stream.
	//
	// The name must be different than the names of other metric streams in this account and Region.
	//
	// If you are updating a metric stream, specify the name of that stream here.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricstream.html#cfn-cloudwatch-metricstream-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// By default, a metric stream always sends the MAX, MIN, SUM, and SAMPLECOUNT statistics for each metric that is streamed.
	//
	// You can use this parameter to have the metric stream also send additional statistics in the stream. This array can have up to 100 members.
	//
	// For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's `OutputFormat` . If the `OutputFormat` is `json` , you can stream any additional statistic that is supported by CloudWatch , listed in [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html) . If the `OutputFormat` is OpenTelemetry, you can stream percentile statistics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricstream.html#cfn-cloudwatch-metricstream-statisticsconfigurations
	//
	StatisticsConfigurations interface{} `field:"optional" json:"statisticsConfigurations" yaml:"statisticsConfigurations"`
	// An array of key-value pairs to apply to the metric stream.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricstream.html#cfn-cloudwatch-metricstream-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

