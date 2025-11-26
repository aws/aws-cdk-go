package previewawslookoutmetricsmixins


// Contains information about a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metricSetProperty := &MetricSetProperty{
//   	DimensionList: []*string{
//   		jsii.String("dimensionList"),
//   	},
//   	MetricList: []interface{}{
//   		&MetricProperty{
//   			AggregationFunction: jsii.String("aggregationFunction"),
//   			MetricName: jsii.String("metricName"),
//   			Namespace: jsii.String("namespace"),
//   		},
//   	},
//   	MetricSetDescription: jsii.String("metricSetDescription"),
//   	MetricSetFrequency: jsii.String("metricSetFrequency"),
//   	MetricSetName: jsii.String("metricSetName"),
//   	MetricSource: &MetricSourceProperty{
//   		AppFlowConfig: &AppFlowConfigProperty{
//   			FlowName: jsii.String("flowName"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		CloudwatchConfig: &CloudwatchConfigProperty{
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		RdsSourceConfig: &RDSSourceConfigProperty{
//   			DatabaseHost: jsii.String("databaseHost"),
//   			DatabaseName: jsii.String("databaseName"),
//   			DatabasePort: jsii.Number(123),
//   			DbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   			RoleArn: jsii.String("roleArn"),
//   			SecretManagerArn: jsii.String("secretManagerArn"),
//   			TableName: jsii.String("tableName"),
//   			VpcConfiguration: &VpcConfigurationProperty{
//   				SecurityGroupIdList: []*string{
//   					jsii.String("securityGroupIdList"),
//   				},
//   				SubnetIdList: []*string{
//   					jsii.String("subnetIdList"),
//   				},
//   			},
//   		},
//   		RedshiftSourceConfig: &RedshiftSourceConfigProperty{
//   			ClusterIdentifier: jsii.String("clusterIdentifier"),
//   			DatabaseHost: jsii.String("databaseHost"),
//   			DatabaseName: jsii.String("databaseName"),
//   			DatabasePort: jsii.Number(123),
//   			RoleArn: jsii.String("roleArn"),
//   			SecretManagerArn: jsii.String("secretManagerArn"),
//   			TableName: jsii.String("tableName"),
//   			VpcConfiguration: &VpcConfigurationProperty{
//   				SecurityGroupIdList: []*string{
//   					jsii.String("securityGroupIdList"),
//   				},
//   				SubnetIdList: []*string{
//   					jsii.String("subnetIdList"),
//   				},
//   			},
//   		},
//   		S3SourceConfig: &S3SourceConfigProperty{
//   			FileFormatDescriptor: &FileFormatDescriptorProperty{
//   				CsvFormatDescriptor: &CsvFormatDescriptorProperty{
//   					Charset: jsii.String("charset"),
//   					ContainsHeader: jsii.Boolean(false),
//   					Delimiter: jsii.String("delimiter"),
//   					FileCompression: jsii.String("fileCompression"),
//   					HeaderList: []*string{
//   						jsii.String("headerList"),
//   					},
//   					QuoteSymbol: jsii.String("quoteSymbol"),
//   				},
//   				JsonFormatDescriptor: &JsonFormatDescriptorProperty{
//   					Charset: jsii.String("charset"),
//   					FileCompression: jsii.String("fileCompression"),
//   				},
//   			},
//   			HistoricalDataPathList: []*string{
//   				jsii.String("historicalDataPathList"),
//   			},
//   			RoleArn: jsii.String("roleArn"),
//   			TemplatedPathList: []*string{
//   				jsii.String("templatedPathList"),
//   			},
//   		},
//   	},
//   	Offset: jsii.Number(123),
//   	TimestampColumn: &TimestampColumnProperty{
//   		ColumnFormat: jsii.String("columnFormat"),
//   		ColumnName: jsii.String("columnName"),
//   	},
//   	Timezone: jsii.String("timezone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricset.html
//
type CfnAnomalyDetectorPropsMixin_MetricSetProperty struct {
	// A list of the fields you want to treat as dimensions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricset.html#cfn-lookoutmetrics-anomalydetector-metricset-dimensionlist
	//
	DimensionList *[]*string `field:"optional" json:"dimensionList" yaml:"dimensionList"`
	// A list of metrics that the dataset will contain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricset.html#cfn-lookoutmetrics-anomalydetector-metricset-metriclist
	//
	MetricList interface{} `field:"optional" json:"metricList" yaml:"metricList"`
	// A description of the dataset you are creating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricset.html#cfn-lookoutmetrics-anomalydetector-metricset-metricsetdescription
	//
	MetricSetDescription *string `field:"optional" json:"metricSetDescription" yaml:"metricSetDescription"`
	// The frequency with which the source data will be analyzed for anomalies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricset.html#cfn-lookoutmetrics-anomalydetector-metricset-metricsetfrequency
	//
	MetricSetFrequency *string `field:"optional" json:"metricSetFrequency" yaml:"metricSetFrequency"`
	// The name of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricset.html#cfn-lookoutmetrics-anomalydetector-metricset-metricsetname
	//
	MetricSetName *string `field:"optional" json:"metricSetName" yaml:"metricSetName"`
	// Contains information about how the source data should be interpreted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricset.html#cfn-lookoutmetrics-anomalydetector-metricset-metricsource
	//
	MetricSource interface{} `field:"optional" json:"metricSource" yaml:"metricSource"`
	// After an interval ends, the amount of seconds that the detector waits before importing data.
	//
	// Offset is only supported for S3, Redshift, Athena and datasources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricset.html#cfn-lookoutmetrics-anomalydetector-metricset-offset
	//
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Contains information about the column used for tracking time in your source data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricset.html#cfn-lookoutmetrics-anomalydetector-metricset-timestampcolumn
	//
	TimestampColumn interface{} `field:"optional" json:"timestampColumn" yaml:"timestampColumn"`
	// The time zone in which your source data was recorded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricset.html#cfn-lookoutmetrics-anomalydetector-metricset-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

