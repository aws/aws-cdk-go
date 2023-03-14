package awskinesisanalyticsv2


// The configuration of a Kinesis Data Analytics Studio notebook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   zeppelinApplicationConfigurationProperty := &zeppelinApplicationConfigurationProperty{
//   	catalogConfiguration: &catalogConfigurationProperty{
//   		glueDataCatalogConfiguration: &glueDataCatalogConfigurationProperty{
//   			databaseArn: jsii.String("databaseArn"),
//   		},
//   	},
//   	customArtifactsConfiguration: []interface{}{
//   		&customArtifactConfigurationProperty{
//   			artifactType: jsii.String("artifactType"),
//
//   			// the properties below are optional
//   			mavenReference: &mavenReferenceProperty{
//   				artifactId: jsii.String("artifactId"),
//   				groupId: jsii.String("groupId"),
//   				version: jsii.String("version"),
//   			},
//   			s3ContentLocation: &s3ContentLocationProperty{
//   				bucketArn: jsii.String("bucketArn"),
//   				fileKey: jsii.String("fileKey"),
//
//   				// the properties below are optional
//   				objectVersion: jsii.String("objectVersion"),
//   			},
//   		},
//   	},
//   	deployAsApplicationConfiguration: &deployAsApplicationConfigurationProperty{
//   		s3ContentLocation: &s3ContentBaseLocationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//
//   			// the properties below are optional
//   			basePath: jsii.String("basePath"),
//   		},
//   	},
//   	monitoringConfiguration: &zeppelinMonitoringConfigurationProperty{
//   		logLevel: jsii.String("logLevel"),
//   	},
//   }
//
type CfnApplication_ZeppelinApplicationConfigurationProperty struct {
	// The Amazon Glue Data Catalog that you use in queries in a Kinesis Data Analytics Studio notebook.
	CatalogConfiguration interface{} `field:"optional" json:"catalogConfiguration" yaml:"catalogConfiguration"`
	// A list of `CustomArtifactConfiguration` objects.
	CustomArtifactsConfiguration interface{} `field:"optional" json:"customArtifactsConfiguration" yaml:"customArtifactsConfiguration"`
	// The information required to deploy a Kinesis Data Analytics Studio notebook as an application with durable state.
	DeployAsApplicationConfiguration interface{} `field:"optional" json:"deployAsApplicationConfiguration" yaml:"deployAsApplicationConfiguration"`
	// The monitoring configuration of a Kinesis Data Analytics Studio notebook.
	MonitoringConfiguration interface{} `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
}

