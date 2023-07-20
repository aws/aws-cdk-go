package awsemrserverless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	ReleaseLabel: jsii.String("releaseLabel"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Architecture: jsii.String("architecture"),
//   	AutoStartConfiguration: &AutoStartConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	AutoStopConfiguration: &AutoStopConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		IdleTimeoutMinutes: jsii.Number(123),
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
//   					Memory: jsii.String("memory"),
//
//   					// the properties below are optional
//   					Disk: jsii.String("disk"),
//   				},
//   				WorkerCount: jsii.Number(123),
//   			},
//   		},
//   	},
//   	MaximumCapacity: &MaximumAllowedResourcesProperty{
//   		Cpu: jsii.String("cpu"),
//   		Memory: jsii.String("memory"),
//
//   		// the properties below are optional
//   		Disk: jsii.String("disk"),
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
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
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
type CfnApplicationProps struct {
	// The EMR release version associated with the application.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 64
	//
	// *Pattern* : `^[A-Za-z0-9._/-]+$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-releaselabel
	//
	ReleaseLabel *string `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// The type of application, such as Spark or Hive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The CPU architecture type of the application.
	//
	// Allowed values: `X86_64` or `ARM64`.
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
	// The image configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-imageconfiguration
	//
	ImageConfiguration interface{} `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
	// The initial capacity of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-initialcapacity
	//
	InitialCapacity interface{} `field:"optional" json:"initialCapacity" yaml:"initialCapacity"`
	// The maximum capacity of the application.
	//
	// This is cumulative across all workers at any given point in time during the lifespan of the application is created. No new resources will be created once any one of the defined limits is hit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-maximumcapacity
	//
	MaximumCapacity interface{} `field:"optional" json:"maximumCapacity" yaml:"maximumCapacity"`
	// The name of the application.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 64
	//
	// *Pattern* : `^[A-Za-z0-9._\\/#-]+$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The network configuration for customer VPC connectivity for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The tags assigned to the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html#cfn-emrserverless-application-workertypespecifications
	//
	WorkerTypeSpecifications interface{} `field:"optional" json:"workerTypeSpecifications" yaml:"workerTypeSpecifications"`
}

