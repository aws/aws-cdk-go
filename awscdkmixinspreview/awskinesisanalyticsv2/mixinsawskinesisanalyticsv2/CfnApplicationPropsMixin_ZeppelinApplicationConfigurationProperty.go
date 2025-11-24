package mixinsawskinesisanalyticsv2


// The configuration of a Kinesis Data Analytics Studio notebook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   			MavenReference: &MavenReferenceProperty{
//   				ArtifactId: jsii.String("artifactId"),
//   				GroupId: jsii.String("groupId"),
//   				Version: jsii.String("version"),
//   			},
//   			S3ContentLocation: &S3ContentLocationProperty{
//   				BucketArn: jsii.String("bucketArn"),
//   				FileKey: jsii.String("fileKey"),
//   				ObjectVersion: jsii.String("objectVersion"),
//   			},
//   		},
//   	},
//   	DeployAsApplicationConfiguration: &DeployAsApplicationConfigurationProperty{
//   		S3ContentLocation: &S3ContentBaseLocationProperty{
//   			BasePath: jsii.String("basePath"),
//   			BucketArn: jsii.String("bucketArn"),
//   		},
//   	},
//   	MonitoringConfiguration: &ZeppelinMonitoringConfigurationProperty{
//   		LogLevel: jsii.String("logLevel"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-zeppelinapplicationconfiguration.html
//
type CfnApplicationPropsMixin_ZeppelinApplicationConfigurationProperty struct {
	// The Amazon Glue Data Catalog that you use in queries in a Kinesis Data Analytics Studio notebook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-zeppelinapplicationconfiguration.html#cfn-kinesisanalyticsv2-application-zeppelinapplicationconfiguration-catalogconfiguration
	//
	CatalogConfiguration interface{} `field:"optional" json:"catalogConfiguration" yaml:"catalogConfiguration"`
	// A list of `CustomArtifactConfiguration` objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-zeppelinapplicationconfiguration.html#cfn-kinesisanalyticsv2-application-zeppelinapplicationconfiguration-customartifactsconfiguration
	//
	CustomArtifactsConfiguration interface{} `field:"optional" json:"customArtifactsConfiguration" yaml:"customArtifactsConfiguration"`
	// The information required to deploy a Kinesis Data Analytics Studio notebook as an application with durable state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-zeppelinapplicationconfiguration.html#cfn-kinesisanalyticsv2-application-zeppelinapplicationconfiguration-deployasapplicationconfiguration
	//
	DeployAsApplicationConfiguration interface{} `field:"optional" json:"deployAsApplicationConfiguration" yaml:"deployAsApplicationConfiguration"`
	// The monitoring configuration of a Kinesis Data Analytics Studio notebook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-zeppelinapplicationconfiguration.html#cfn-kinesisanalyticsv2-application-zeppelinapplicationconfiguration-monitoringconfiguration
	//
	MonitoringConfiguration interface{} `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
}

