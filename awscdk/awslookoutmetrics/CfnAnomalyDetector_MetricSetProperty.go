package awslookoutmetrics


// Contains information about a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricSetProperty := &MetricSetProperty{
//   	MetricList: []interface{}{
//   		&MetricProperty{
//   			AggregationFunction: jsii.String("aggregationFunction"),
//   			MetricName: jsii.String("metricName"),
//
//   			// the properties below are optional
//   			Namespace: jsii.String("namespace"),
//   		},
//   	},
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
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			HistoricalDataPathList: []*string{
//   				jsii.String("historicalDataPathList"),
//   			},
//   			TemplatedPathList: []*string{
//   				jsii.String("templatedPathList"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	DimensionList: []*string{
//   		jsii.String("dimensionList"),
//   	},
//   	MetricSetDescription: jsii.String("metricSetDescription"),
//   	MetricSetFrequency: jsii.String("metricSetFrequency"),
//   	Offset: jsii.Number(123),
//   	TimestampColumn: &TimestampColumnProperty{
//   		ColumnFormat: jsii.String("columnFormat"),
//   		ColumnName: jsii.String("columnName"),
//   	},
//   	Timezone: jsii.String("timezone"),
//   }
//
type CfnAnomalyDetector_MetricSetProperty struct {
	// A list of metrics that the dataset will contain.
	MetricList interface{} `field:"required" json:"metricList" yaml:"metricList"`
	// The name of the dataset.
	MetricSetName *string `field:"required" json:"metricSetName" yaml:"metricSetName"`
	// Contains information about how the source data should be interpreted.
	MetricSource interface{} `field:"required" json:"metricSource" yaml:"metricSource"`
	// A list of the fields you want to treat as dimensions.
	DimensionList *[]*string `field:"optional" json:"dimensionList" yaml:"dimensionList"`
	// A description of the dataset you are creating.
	MetricSetDescription *string `field:"optional" json:"metricSetDescription" yaml:"metricSetDescription"`
	// The frequency with which the source data will be analyzed for anomalies.
	MetricSetFrequency *string `field:"optional" json:"metricSetFrequency" yaml:"metricSetFrequency"`
	// After an interval ends, the amount of seconds that the detector waits before importing data.
	//
	// Offset is only supported for S3, Redshift, Athena and datasources.
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Contains information about the column used for tracking time in your source data.
	TimestampColumn interface{} `field:"optional" json:"timestampColumn" yaml:"timestampColumn"`
	// The time zone in which your source data was recorded.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

