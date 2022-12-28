package awsapprunner

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceProps := &cfnServiceProps{
//   	sourceConfiguration: &sourceConfigurationProperty{
//   		authenticationConfiguration: &authenticationConfigurationProperty{
//   			accessRoleArn: jsii.String("accessRoleArn"),
//   			connectionArn: jsii.String("connectionArn"),
//   		},
//   		autoDeploymentsEnabled: jsii.Boolean(false),
//   		codeRepository: &codeRepositoryProperty{
//   			repositoryUrl: jsii.String("repositoryUrl"),
//   			sourceCodeVersion: &sourceCodeVersionProperty{
//   				type: jsii.String("type"),
//   				value: jsii.String("value"),
//   			},
//
//   			// the properties below are optional
//   			codeConfiguration: &codeConfigurationProperty{
//   				configurationSource: jsii.String("configurationSource"),
//
//   				// the properties below are optional
//   				codeConfigurationValues: &codeConfigurationValuesProperty{
//   					runtime: jsii.String("runtime"),
//
//   					// the properties below are optional
//   					buildCommand: jsii.String("buildCommand"),
//   					port: jsii.String("port"),
//   					runtimeEnvironmentVariables: []interface{}{
//   						&keyValuePairProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					startCommand: jsii.String("startCommand"),
//   				},
//   			},
//   		},
//   		imageRepository: &imageRepositoryProperty{
//   			imageIdentifier: jsii.String("imageIdentifier"),
//   			imageRepositoryType: jsii.String("imageRepositoryType"),
//
//   			// the properties below are optional
//   			imageConfiguration: &imageConfigurationProperty{
//   				port: jsii.String("port"),
//   				runtimeEnvironmentVariables: []interface{}{
//   					&keyValuePairProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				startCommand: jsii.String("startCommand"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	autoScalingConfigurationArn: jsii.String("autoScalingConfigurationArn"),
//   	encryptionConfiguration: &encryptionConfigurationProperty{
//   		kmsKey: jsii.String("kmsKey"),
//   	},
//   	healthCheckConfiguration: &healthCheckConfigurationProperty{
//   		healthyThreshold: jsii.Number(123),
//   		interval: jsii.Number(123),
//   		path: jsii.String("path"),
//   		protocol: jsii.String("protocol"),
//   		timeout: jsii.Number(123),
//   		unhealthyThreshold: jsii.Number(123),
//   	},
//   	instanceConfiguration: &instanceConfigurationProperty{
//   		cpu: jsii.String("cpu"),
//   		instanceRoleArn: jsii.String("instanceRoleArn"),
//   		memory: jsii.String("memory"),
//   	},
//   	networkConfiguration: &networkConfigurationProperty{
//   		egressConfiguration: &egressConfigurationProperty{
//   			egressType: jsii.String("egressType"),
//
//   			// the properties below are optional
//   			vpcConnectorArn: jsii.String("vpcConnectorArn"),
//   		},
//   		ingressConfiguration: &ingressConfigurationProperty{
//   			isPubliclyAccessible: jsii.Boolean(false),
//   		},
//   	},
//   	observabilityConfiguration: &serviceObservabilityConfigurationProperty{
//   		observabilityEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		observabilityConfigurationArn: jsii.String("observabilityConfigurationArn"),
//   	},
//   	serviceName: jsii.String("serviceName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnServiceProps struct {
	// The source to deploy to the App Runner service.
	//
	// It can be a code or an image repository.
	SourceConfiguration interface{} `field:"required" json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// The Amazon Resource Name (ARN) of an App Runner automatic scaling configuration resource that you want to associate with your service.
	//
	// If not provided, App Runner associates the latest revision of a default auto scaling configuration.
	//
	// Specify an ARN with a name and a revision number to associate that revision. For example: `arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability/3`
	//
	// Specify just the name to associate the latest revision. For example: `arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability`
	AutoScalingConfigurationArn *string `field:"optional" json:"autoScalingConfigurationArn" yaml:"autoScalingConfigurationArn"`
	// An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs.
	//
	// By default, App Runner uses an AWS managed key .
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The settings for the health check that AWS App Runner performs to monitor the health of the App Runner service.
	HealthCheckConfiguration interface{} `field:"optional" json:"healthCheckConfiguration" yaml:"healthCheckConfiguration"`
	// The runtime configuration of instances (scaling units) of your service.
	InstanceConfiguration interface{} `field:"optional" json:"instanceConfiguration" yaml:"instanceConfiguration"`
	// Configuration settings related to network traffic of the web application that the App Runner service runs.
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The observability configuration of your service.
	ObservabilityConfiguration interface{} `field:"optional" json:"observabilityConfiguration" yaml:"observabilityConfiguration"`
	// A name for the App Runner service.
	//
	// It must be unique across all the running App Runner services in your AWS account in the AWS Region .
	//
	// If you don't specify a name, AWS CloudFormation generates a name for your service.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// An optional list of metadata items that you can associate with the App Runner service resource.
	//
	// A tag is a key-value pair.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

