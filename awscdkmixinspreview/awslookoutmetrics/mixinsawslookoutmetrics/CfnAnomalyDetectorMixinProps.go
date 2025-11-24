package mixinsawslookoutmetrics


// Properties for CfnAnomalyDetectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAnomalyDetectorMixinProps := &CfnAnomalyDetectorMixinProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-anomalydetector.html
//
type CfnAnomalyDetectorMixinProps struct {
	// Contains information about the configuration of the anomaly detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-anomalydetector.html#cfn-lookoutmetrics-anomalydetector-anomalydetectorconfig
	//
	AnomalyDetectorConfig interface{} `field:"optional" json:"anomalyDetectorConfig" yaml:"anomalyDetectorConfig"`
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
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The detector's dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-anomalydetector.html#cfn-lookoutmetrics-anomalydetector-metricsetlist
	//
	MetricSetList interface{} `field:"optional" json:"metricSetList" yaml:"metricSetList"`
}

