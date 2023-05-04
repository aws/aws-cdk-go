package awskinesisanalyticsv2


// The configuration of a Kinesis Data Analytics Studio notebook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   zeppelinApplicationConfigurationProperty := &ZeppelinApplicationConfigurationProperty{
//   	CatalogConfiguration: &CatalogConfigurationProperty{
//   		GlueDataCatalogConfiguration: &GlueDataCatalogConfigurationProperty{
//   			DatabaseArn: jsii.String("databaseArn"),
//   		},
//   	},
//   	CustomArtifactsConfiguration: []interface{}{
//   		&CustomArtifactConfigurationProperty{
//   			ArtifactType: jsii.String("artifactType"),
//
//   			// the properties below are optional
//   			MavenReference: &MavenReferenceProperty{
//   				ArtifactId: jsii.String("artifactId"),
//   				GroupId: jsii.String("groupId"),
//   				Version: jsii.String("version"),
//   			},
//   			S3ContentLocation: &S3ContentLocationProperty{
//   				BucketArn: jsii.String("bucketArn"),
//   				FileKey: jsii.String("fileKey"),
//
//   				// the properties below are optional
//   				ObjectVersion: jsii.String("objectVersion"),
//   			},
//   		},
//   	},
//   	DeployAsApplicationConfiguration: &DeployAsApplicationConfigurationProperty{
//   		S3ContentLocation: &S3ContentBaseLocationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//
//   			// the properties below are optional
//   			BasePath: jsii.String("basePath"),
//   		},
//   	},
//   	MonitoringConfiguration: &ZeppelinMonitoringConfigurationProperty{
//   		LogLevel: jsii.String("logLevel"),
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

