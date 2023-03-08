package awslookoutmetrics


// Contains information about how the source data should be interpreted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricSourceProperty := &MetricSourceProperty{
//   	AppFlowConfig: &AppFlowConfigProperty{
//   		FlowName: jsii.String("flowName"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	CloudwatchConfig: &CloudwatchConfigProperty{
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	RdsSourceConfig: &RDSSourceConfigProperty{
//   		DatabaseHost: jsii.String("databaseHost"),
//   		DatabaseName: jsii.String("databaseName"),
//   		DatabasePort: jsii.Number(123),
//   		DbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   		RoleArn: jsii.String("roleArn"),
//   		SecretManagerArn: jsii.String("secretManagerArn"),
//   		TableName: jsii.String("tableName"),
//   		VpcConfiguration: &VpcConfigurationProperty{
//   			SecurityGroupIdList: []*string{
//   				jsii.String("securityGroupIdList"),
//   			},
//   			SubnetIdList: []*string{
//   				jsii.String("subnetIdList"),
//   			},
//   		},
//   	},
//   	RedshiftSourceConfig: &RedshiftSourceConfigProperty{
//   		ClusterIdentifier: jsii.String("clusterIdentifier"),
//   		DatabaseHost: jsii.String("databaseHost"),
//   		DatabaseName: jsii.String("databaseName"),
//   		DatabasePort: jsii.Number(123),
//   		RoleArn: jsii.String("roleArn"),
//   		SecretManagerArn: jsii.String("secretManagerArn"),
//   		TableName: jsii.String("tableName"),
//   		VpcConfiguration: &VpcConfigurationProperty{
//   			SecurityGroupIdList: []*string{
//   				jsii.String("securityGroupIdList"),
//   			},
//   			SubnetIdList: []*string{
//   				jsii.String("subnetIdList"),
//   			},
//   		},
//   	},
//   	S3SourceConfig: &S3SourceConfigProperty{
//   		FileFormatDescriptor: &FileFormatDescriptorProperty{
//   			CsvFormatDescriptor: &CsvFormatDescriptorProperty{
//   				Charset: jsii.String("charset"),
//   				ContainsHeader: jsii.Boolean(false),
//   				Delimiter: jsii.String("delimiter"),
//   				FileCompression: jsii.String("fileCompression"),
//   				HeaderList: []*string{
//   					jsii.String("headerList"),
//   				},
//   				QuoteSymbol: jsii.String("quoteSymbol"),
//   			},
//   			JsonFormatDescriptor: &JsonFormatDescriptorProperty{
//   				Charset: jsii.String("charset"),
//   				FileCompression: jsii.String("fileCompression"),
//   			},
//   		},
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		HistoricalDataPathList: []*string{
//   			jsii.String("historicalDataPathList"),
//   		},
//   		TemplatedPathList: []*string{
//   			jsii.String("templatedPathList"),
//   		},
//   	},
//   }
//
type CfnAnomalyDetector_MetricSourceProperty struct {
	// Details about an AppFlow datasource.
	AppFlowConfig interface{} `field:"optional" json:"appFlowConfig" yaml:"appFlowConfig"`
	// Details about an Amazon CloudWatch monitoring datasource.
	CloudwatchConfig interface{} `field:"optional" json:"cloudwatchConfig" yaml:"cloudwatchConfig"`
	// Details about an Amazon Relational Database Service (RDS) datasource.
	RdsSourceConfig interface{} `field:"optional" json:"rdsSourceConfig" yaml:"rdsSourceConfig"`
	// Details about an Amazon Redshift database datasource.
	RedshiftSourceConfig interface{} `field:"optional" json:"redshiftSourceConfig" yaml:"redshiftSourceConfig"`
	// Contains information about the configuration of the S3 bucket that contains source files.
	S3SourceConfig interface{} `field:"optional" json:"s3SourceConfig" yaml:"s3SourceConfig"`
}

