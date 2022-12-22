package awsemrserverless

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &cfnApplicationProps{
//   	releaseLabel: jsii.String("releaseLabel"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	architecture: jsii.String("architecture"),
//   	autoStartConfiguration: &autoStartConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	autoStopConfiguration: &autoStopConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   		idleTimeoutMinutes: jsii.Number(123),
//   	},
//   	initialCapacity: []interface{}{
//   		&initialCapacityConfigKeyValuePairProperty{
//   			key: jsii.String("key"),
//   			value: &initialCapacityConfigProperty{
//   				workerConfiguration: &workerConfigurationProperty{
//   					cpu: jsii.String("cpu"),
//   					memory: jsii.String("memory"),
//
//   					// the properties below are optional
//   					disk: jsii.String("disk"),
//   				},
//   				workerCount: jsii.Number(123),
//   			},
//   		},
//   	},
//   	maximumCapacity: &maximumAllowedResourcesProperty{
//   		cpu: jsii.String("cpu"),
//   		memory: jsii.String("memory"),
//
//   		// the properties below are optional
//   		disk: jsii.String("disk"),
//   	},
//   	name: jsii.String("name"),
//   	networkConfiguration: &networkConfigurationProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnApplicationProps struct {
	// The EMR release version associated with the application.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 64
	//
	// *Pattern* : `^[A-Za-z0-9._/-]+$`
	ReleaseLabel *string `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// The type of application, such as Spark or Hive.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `AWS::EMRServerless::Application.Architecture`.
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// The configuration for an application to automatically start on job submission.
	AutoStartConfiguration interface{} `field:"optional" json:"autoStartConfiguration" yaml:"autoStartConfiguration"`
	// The configuration for an application to automatically stop after a certain amount of time being idle.
	AutoStopConfiguration interface{} `field:"optional" json:"autoStopConfiguration" yaml:"autoStopConfiguration"`
	// The initial capacity of the application.
	InitialCapacity interface{} `field:"optional" json:"initialCapacity" yaml:"initialCapacity"`
	// The maximum capacity of the application.
	//
	// This is cumulative across all workers at any given point in time during the lifespan of the application is created. No new resources will be created once any one of the defined limits is hit.
	MaximumCapacity interface{} `field:"optional" json:"maximumCapacity" yaml:"maximumCapacity"`
	// The name of the application.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 64
	//
	// *Pattern* : `^[A-Za-z0-9._\\/#-]+$`
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The network configuration for customer VPC connectivity for the application.
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The tags assigned to the application.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

