package awslookoutmetrics


// Properties for defining a `CfnAnomalyDetector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnomalyDetectorProps := &cfnAnomalyDetectorProps{
//   	anomalyDetectorConfig: &anomalyDetectorConfigProperty{
//   		anomalyDetectorFrequency: jsii.String("anomalyDetectorFrequency"),
//   	},
//   	metricSetList: []interface{}{
//   		&metricSetProperty{
//   			metricList: []interface{}{
//   				&metricProperty{
//   					aggregationFunction: jsii.String("aggregationFunction"),
//   					metricName: jsii.String("metricName"),
//
//   					// the properties below are optional
//   					namespace: jsii.String("namespace"),
//   				},
//   			},
//   			metricSetName: jsii.String("metricSetName"),
//   			metricSource: &metricSourceProperty{
//   				appFlowConfig: &appFlowConfigProperty{
//   					flowName: jsii.String("flowName"),
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				cloudwatchConfig: &cloudwatchConfigProperty{
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				rdsSourceConfig: &rDSSourceConfigProperty{
//   					databaseHost: jsii.String("databaseHost"),
//   					databaseName: jsii.String("databaseName"),
//   					databasePort: jsii.Number(123),
//   					dbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   					roleArn: jsii.String("roleArn"),
//   					secretManagerArn: jsii.String("secretManagerArn"),
//   					tableName: jsii.String("tableName"),
//   					vpcConfiguration: &vpcConfigurationProperty{
//   						securityGroupIdList: []*string{
//   							jsii.String("securityGroupIdList"),
//   						},
//   						subnetIdList: []*string{
//   							jsii.String("subnetIdList"),
//   						},
//   					},
//   				},
//   				redshiftSourceConfig: &redshiftSourceConfigProperty{
//   					clusterIdentifier: jsii.String("clusterIdentifier"),
//   					databaseHost: jsii.String("databaseHost"),
//   					databaseName: jsii.String("databaseName"),
//   					databasePort: jsii.Number(123),
//   					roleArn: jsii.String("roleArn"),
//   					secretManagerArn: jsii.String("secretManagerArn"),
//   					tableName: jsii.String("tableName"),
//   					vpcConfiguration: &vpcConfigurationProperty{
//   						securityGroupIdList: []*string{
//   							jsii.String("securityGroupIdList"),
//   						},
//   						subnetIdList: []*string{
//   							jsii.String("subnetIdList"),
//   						},
//   					},
//   				},
//   				s3SourceConfig: &s3SourceConfigProperty{
//   					fileFormatDescriptor: &fileFormatDescriptorProperty{
//   						csvFormatDescriptor: &csvFormatDescriptorProperty{
//   							charset: jsii.String("charset"),
//   							containsHeader: jsii.Boolean(false),
//   							delimiter: jsii.String("delimiter"),
//   							fileCompression: jsii.String("fileCompression"),
//   							headerList: []*string{
//   								jsii.String("headerList"),
//   							},
//   							quoteSymbol: jsii.String("quoteSymbol"),
//   						},
//   						jsonFormatDescriptor: &jsonFormatDescriptorProperty{
//   							charset: jsii.String("charset"),
//   							fileCompression: jsii.String("fileCompression"),
//   						},
//   					},
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					historicalDataPathList: []*string{
//   						jsii.String("historicalDataPathList"),
//   					},
//   					templatedPathList: []*string{
//   						jsii.String("templatedPathList"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			dimensionList: []*string{
//   				jsii.String("dimensionList"),
//   			},
//   			metricSetDescription: jsii.String("metricSetDescription"),
//   			metricSetFrequency: jsii.String("metricSetFrequency"),
//   			offset: jsii.Number(123),
//   			timestampColumn: &timestampColumnProperty{
//   				columnFormat: jsii.String("columnFormat"),
//   				columnName: jsii.String("columnName"),
//   			},
//   			timezone: jsii.String("timezone"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	anomalyDetectorDescription: jsii.String("anomalyDetectorDescription"),
//   	anomalyDetectorName: jsii.String("anomalyDetectorName"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
type CfnAnomalyDetectorProps struct {
	// Contains information about the configuration of the anomaly detector.
	AnomalyDetectorConfig interface{} `field:"required" json:"anomalyDetectorConfig" yaml:"anomalyDetectorConfig"`
	// The detector's dataset.
	MetricSetList interface{} `field:"required" json:"metricSetList" yaml:"metricSetList"`
	// A description of the detector.
	AnomalyDetectorDescription *string `field:"optional" json:"anomalyDetectorDescription" yaml:"anomalyDetectorDescription"`
	// The name of the detector.
	AnomalyDetectorName *string `field:"optional" json:"anomalyDetectorName" yaml:"anomalyDetectorName"`
	// The ARN of the KMS key to use to encrypt your data.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

