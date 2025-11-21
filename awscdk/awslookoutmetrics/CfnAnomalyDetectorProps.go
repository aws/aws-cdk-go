package awslookoutmetrics


// Properties for defining a `CfnAnomalyDetector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnomalyDetectorProps := &CfnAnomalyDetectorProps{
//   	AnomalyDetectorConfig: &AnomalyDetectorConfigProperty{
//   		AnomalyDetectorFrequency: jsii.String("anomalyDetectorFrequency"),
//   	},
//   	MetricSetList: []interface{}{
//   		&MetricSetProperty{
//   			MetricList: []interface{}{
//   				&MetricProperty{
//   					AggregationFunction: jsii.String("aggregationFunction"),
//   					MetricName: jsii.String("metricName"),
//
//   					// the properties below are optional
//   					Namespace: jsii.String("namespace"),
//   				},
//   			},
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
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					HistoricalDataPathList: []*string{
//   						jsii.String("historicalDataPathList"),
//   					},
//   					TemplatedPathList: []*string{
//   						jsii.String("templatedPathList"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			DimensionList: []*string{
//   				jsii.String("dimensionList"),
//   			},
//   			MetricSetDescription: jsii.String("metricSetDescription"),
//   			MetricSetFrequency: jsii.String("metricSetFrequency"),
//   			Offset: jsii.Number(123),
//   			TimestampColumn: &TimestampColumnProperty{
//   				ColumnFormat: jsii.String("columnFormat"),
//   				ColumnName: jsii.String("columnName"),
//   			},
//   			Timezone: jsii.String("timezone"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AnomalyDetectorDescription: jsii.String("anomalyDetectorDescription"),
//   	AnomalyDetectorName: jsii.String("anomalyDetectorName"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-anomalydetector.html
//
type CfnAnomalyDetectorProps struct {
	// Contains information about the configuration of the anomaly detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-anomalydetector.html#cfn-lookoutmetrics-anomalydetector-anomalydetectorconfig
	//
	AnomalyDetectorConfig interface{} `field:"required" json:"anomalyDetectorConfig" yaml:"anomalyDetectorConfig"`
	// The detector's dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-anomalydetector.html#cfn-lookoutmetrics-anomalydetector-metricsetlist
	//
	MetricSetList interface{} `field:"required" json:"metricSetList" yaml:"metricSetList"`
	// A description of the detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-anomalydetector.html#cfn-lookoutmetrics-anomalydetector-anomalydetectordescription
	//
	AnomalyDetectorDescription *string `field:"optional" json:"anomalyDetectorDescription" yaml:"anomalyDetectorDescription"`
	// The name of the detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-anomalydetector.html#cfn-lookoutmetrics-anomalydetector-anomalydetectorname
	//
	AnomalyDetectorName *string `field:"optional" json:"anomalyDetectorName" yaml:"anomalyDetectorName"`
	// The ARN of the KMS key to use to encrypt your data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-anomalydetector.html#cfn-lookoutmetrics-anomalydetector-kmskeyarn
	//
	KmsKeyArn interface{} `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

