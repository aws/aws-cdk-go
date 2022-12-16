package awslookoutmetrics


// Contains information about a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricSetProperty := &metricSetProperty{
//   	metricList: []interface{}{
//   		&metricProperty{
//   			aggregationFunction: jsii.String("aggregationFunction"),
//   			metricName: jsii.String("metricName"),
//
//   			// the properties below are optional
//   			namespace: jsii.String("namespace"),
//   		},
//   	},
//   	metricSetName: jsii.String("metricSetName"),
//   	metricSource: &metricSourceProperty{
//   		appFlowConfig: &appFlowConfigProperty{
//   			flowName: jsii.String("flowName"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		cloudwatchConfig: &cloudwatchConfigProperty{
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		rdsSourceConfig: &rDSSourceConfigProperty{
//   			databaseHost: jsii.String("databaseHost"),
//   			databaseName: jsii.String("databaseName"),
//   			databasePort: jsii.Number(123),
//   			dbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   			roleArn: jsii.String("roleArn"),
//   			secretManagerArn: jsii.String("secretManagerArn"),
//   			tableName: jsii.String("tableName"),
//   			vpcConfiguration: &vpcConfigurationProperty{
//   				securityGroupIdList: []*string{
//   					jsii.String("securityGroupIdList"),
//   				},
//   				subnetIdList: []*string{
//   					jsii.String("subnetIdList"),
//   				},
//   			},
//   		},
//   		redshiftSourceConfig: &redshiftSourceConfigProperty{
//   			clusterIdentifier: jsii.String("clusterIdentifier"),
//   			databaseHost: jsii.String("databaseHost"),
//   			databaseName: jsii.String("databaseName"),
//   			databasePort: jsii.Number(123),
//   			roleArn: jsii.String("roleArn"),
//   			secretManagerArn: jsii.String("secretManagerArn"),
//   			tableName: jsii.String("tableName"),
//   			vpcConfiguration: &vpcConfigurationProperty{
//   				securityGroupIdList: []*string{
//   					jsii.String("securityGroupIdList"),
//   				},
//   				subnetIdList: []*string{
//   					jsii.String("subnetIdList"),
//   				},
//   			},
//   		},
//   		s3SourceConfig: &s3SourceConfigProperty{
//   			fileFormatDescriptor: &fileFormatDescriptorProperty{
//   				csvFormatDescriptor: &csvFormatDescriptorProperty{
//   					charset: jsii.String("charset"),
//   					containsHeader: jsii.Boolean(false),
//   					delimiter: jsii.String("delimiter"),
//   					fileCompression: jsii.String("fileCompression"),
//   					headerList: []*string{
//   						jsii.String("headerList"),
//   					},
//   					quoteSymbol: jsii.String("quoteSymbol"),
//   				},
//   				jsonFormatDescriptor: &jsonFormatDescriptorProperty{
//   					charset: jsii.String("charset"),
//   					fileCompression: jsii.String("fileCompression"),
//   				},
//   			},
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			historicalDataPathList: []*string{
//   				jsii.String("historicalDataPathList"),
//   			},
//   			templatedPathList: []*string{
//   				jsii.String("templatedPathList"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	dimensionList: []*string{
//   		jsii.String("dimensionList"),
//   	},
//   	metricSetDescription: jsii.String("metricSetDescription"),
//   	metricSetFrequency: jsii.String("metricSetFrequency"),
//   	offset: jsii.Number(123),
//   	timestampColumn: &timestampColumnProperty{
//   		columnFormat: jsii.String("columnFormat"),
//   		columnName: jsii.String("columnName"),
//   	},
//   	timezone: jsii.String("timezone"),
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
	// Offset is only supported for S3 and Redshift datasources.
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Contains information about the column used for tracking time in your source data.
	TimestampColumn interface{} `field:"optional" json:"timestampColumn" yaml:"timestampColumn"`
	// The time zone in which your source data was recorded.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

