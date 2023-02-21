package awslookoutmetrics


// Contains information about how the source data should be interpreted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricSourceProperty := &metricSourceProperty{
//   	appFlowConfig: &appFlowConfigProperty{
//   		flowName: jsii.String("flowName"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	cloudwatchConfig: &cloudwatchConfigProperty{
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	rdsSourceConfig: &rDSSourceConfigProperty{
//   		databaseHost: jsii.String("databaseHost"),
//   		databaseName: jsii.String("databaseName"),
//   		databasePort: jsii.Number(123),
//   		dbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   		roleArn: jsii.String("roleArn"),
//   		secretManagerArn: jsii.String("secretManagerArn"),
//   		tableName: jsii.String("tableName"),
//   		vpcConfiguration: &vpcConfigurationProperty{
//   			securityGroupIdList: []*string{
//   				jsii.String("securityGroupIdList"),
//   			},
//   			subnetIdList: []*string{
//   				jsii.String("subnetIdList"),
//   			},
//   		},
//   	},
//   	redshiftSourceConfig: &redshiftSourceConfigProperty{
//   		clusterIdentifier: jsii.String("clusterIdentifier"),
//   		databaseHost: jsii.String("databaseHost"),
//   		databaseName: jsii.String("databaseName"),
//   		databasePort: jsii.Number(123),
//   		roleArn: jsii.String("roleArn"),
//   		secretManagerArn: jsii.String("secretManagerArn"),
//   		tableName: jsii.String("tableName"),
//   		vpcConfiguration: &vpcConfigurationProperty{
//   			securityGroupIdList: []*string{
//   				jsii.String("securityGroupIdList"),
//   			},
//   			subnetIdList: []*string{
//   				jsii.String("subnetIdList"),
//   			},
//   		},
//   	},
//   	s3SourceConfig: &s3SourceConfigProperty{
//   		fileFormatDescriptor: &fileFormatDescriptorProperty{
//   			csvFormatDescriptor: &csvFormatDescriptorProperty{
//   				charset: jsii.String("charset"),
//   				containsHeader: jsii.Boolean(false),
//   				delimiter: jsii.String("delimiter"),
//   				fileCompression: jsii.String("fileCompression"),
//   				headerList: []*string{
//   					jsii.String("headerList"),
//   				},
//   				quoteSymbol: jsii.String("quoteSymbol"),
//   			},
//   			jsonFormatDescriptor: &jsonFormatDescriptorProperty{
//   				charset: jsii.String("charset"),
//   				fileCompression: jsii.String("fileCompression"),
//   			},
//   		},
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		historicalDataPathList: []*string{
//   			jsii.String("historicalDataPathList"),
//   		},
//   		templatedPathList: []*string{
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

