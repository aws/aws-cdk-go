package previewawsemrserverlessmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnApplicationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var configurationObjectProperty_ ConfigurationObjectProperty
//
//   cfnApplicationMixinProps := &CfnApplicationMixinProps{
//   	Architecture: jsii.String("architecture"),
//   	AutoStartConfiguration: &AutoStartConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	AutoStopConfiguration: &AutoStopConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		IdleTimeoutMinutes: jsii.Number(123),
//   	},
//   	IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   		IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   	},
//   	ImageConfiguration: &ImageConfigurationInputProperty{
//   		ImageUri: jsii.String("imageUri"),
//   	},
//   	InitialCapacity: []interface{}{
//   		&InitialCapacityConfigKeyValuePairProperty{
//   			Key: jsii.String("key"),
//   			Value: &InitialCapacityConfigProperty{
//   				WorkerConfiguration: &WorkerConfigurationProperty{
//   					Cpu: jsii.String("cpu"),
//   					Disk: jsii.String("disk"),
//   					DiskType: jsii.String("diskType"),
//   					Memory: jsii.String("memory"),
//   				},
//   				WorkerCount: jsii.Number(123),
//   			},
//   		},
//   	},
//   	InteractiveConfiguration: &InteractiveConfigurationProperty{
//   		LivyEndpointEnabled: jsii.Boolean(false),
//   		StudioEnabled: jsii.Boolean(false),
//   	},
//   	MaximumCapacity: &MaximumAllowedResourcesProperty{
//   		Cpu: jsii.String("cpu"),
//   		Disk: jsii.String("disk"),
//   		Memory: jsii.String("memory"),
//   	},
//   	MonitoringConfiguration: &MonitoringConfigurationProperty{
//   		CloudWatchLoggingConfiguration: &CloudWatchLoggingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   			LogTypeMap: []interface{}{
//   				&LogTypeMapKeyValuePairProperty{
//   					Key: jsii.String("key"),
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		ManagedPersistenceMonitoringConfiguration: &ManagedPersistenceMonitoringConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   		},
//   		PrometheusMonitoringConfiguration: &PrometheusMonitoringConfigurationProperty{
//   			RemoteWriteUrl: jsii.String("remoteWriteUrl"),
//   		},
//   		S3MonitoringConfiguration: &S3MonitoringConfigurationProperty{
//   			EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   			LogUri: jsii.String("logUri"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	ReleaseLabel: jsii.String("releaseLabel"),
//   	RuntimeConfiguration: []interface{}{
//   		&ConfigurationObjectProperty{
//   			Classification: jsii.String("classification"),
//   			Configurations: []interface{}{
//   				configurationObjectProperty_,
//   			},
//   			Properties: map[string]*string{
//   				"propertiesKey": jsii.String("properties"),
//   			},
//   		},
//   	},
//   	SchedulerConfiguration: &SchedulerConfigurationProperty{
//   		MaxConcurrentRuns: jsii.Number(123),
//   		QueueTimeoutMinutes: jsii.Number(123),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	WorkerTypeSpecifications: map[string]interface{}{
//   		"workerTypeSpecificationsKey": &WorkerTypeSpecificationInputProperty{
//   			"imageConfiguration": &ImageConfigurationInputProperty{
//   				"imageUri": jsii.String("imageUri"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html
//
type CfnApplicationMixinProps struct {
	// The CPU architecture of an application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-architecture
	//
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// The configuration for an application to automatically start on job submission.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-autostartconfiguration
	//
	AutoStartConfiguration interface{} `field:"optional" json:"autoStartConfiguration" yaml:"autoStartConfiguration"`
	// The configuration for an application to automatically stop after a certain amount of time being idle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-autostopconfiguration
	//
	AutoStopConfiguration interface{} `field:"optional" json:"autoStopConfiguration" yaml:"autoStopConfiguration"`
	// The IAM Identity Center configuration applied to enable trusted identity propagation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-identitycenterconfiguration
	//
	IdentityCenterConfiguration interface{} `field:"optional" json:"identityCenterConfiguration" yaml:"identityCenterConfiguration"`
	// The image configuration applied to all worker types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-imageconfiguration
	//
	ImageConfiguration interface{} `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
	// The initial capacity of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-initialcapacity
	//
	InitialCapacity interface{} `field:"optional" json:"initialCapacity" yaml:"initialCapacity"`
	// The interactive configuration object that enables the interactive use cases for an application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-interactiveconfiguration
	//
	InteractiveConfiguration interface{} `field:"optional" json:"interactiveConfiguration" yaml:"interactiveConfiguration"`
	// The maximum capacity of the application.
	//
	// This is cumulative across all workers at any given point in time during the lifespan of the application is created. No new resources will be created once any one of the defined limits is hit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-maximumcapacity
	//
	MaximumCapacity interface{} `field:"optional" json:"maximumCapacity" yaml:"maximumCapacity"`
	// A configuration specification to be used when provisioning an application.
	//
	// A configuration consists of a classification, properties, and optional nested configurations. A classification refers to an application-specific configuration file. Properties are the settings you want to change in that file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-monitoringconfiguration
	//
	MonitoringConfiguration interface{} `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
	// The name of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The network configuration for customer VPC connectivity for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The EMR release associated with the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-releaselabel
	//
	ReleaseLabel *string `field:"optional" json:"releaseLabel" yaml:"releaseLabel"`
	// The [Configuration](https://docs.aws.amazon.com/emr-serverless/latest/APIReference/API_Configuration.html) specifications of an application. Each configuration consists of a classification and properties. You use this parameter when creating or updating an application. To see the runtimeConfiguration object of an application, run the [GetApplication](https://docs.aws.amazon.com/emr-serverless/latest/APIReference/API_GetApplication.html) API operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-runtimeconfiguration
	//
	RuntimeConfiguration interface{} `field:"optional" json:"runtimeConfiguration" yaml:"runtimeConfiguration"`
	// The scheduler configuration for batch and streaming jobs running on this application.
	//
	// Supported with release labels emr-7.0.0 and above.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-schedulerconfiguration
	//
	SchedulerConfiguration interface{} `field:"optional" json:"schedulerConfiguration" yaml:"schedulerConfiguration"`
	// The tags assigned to the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of application, such as Spark or Hive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The specification applied to each worker type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-workertypespecifications
	//
	WorkerTypeSpecifications interface{} `field:"optional" json:"workerTypeSpecifications" yaml:"workerTypeSpecifications"`
}

