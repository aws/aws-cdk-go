package awsapprunner

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceProps := &CfnServiceProps{
//   	SourceConfiguration: &SourceConfigurationProperty{
//   		AuthenticationConfiguration: &AuthenticationConfigurationProperty{
//   			AccessRoleArn: jsii.String("accessRoleArn"),
//   			ConnectionArn: jsii.String("connectionArn"),
//   		},
//   		AutoDeploymentsEnabled: jsii.Boolean(false),
//   		CodeRepository: &CodeRepositoryProperty{
//   			RepositoryUrl: jsii.String("repositoryUrl"),
//   			SourceCodeVersion: &SourceCodeVersionProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//
//   			// the properties below are optional
//   			CodeConfiguration: &CodeConfigurationProperty{
//   				ConfigurationSource: jsii.String("configurationSource"),
//
//   				// the properties below are optional
//   				CodeConfigurationValues: &CodeConfigurationValuesProperty{
//   					Runtime: jsii.String("runtime"),
//
//   					// the properties below are optional
//   					BuildCommand: jsii.String("buildCommand"),
//   					Port: jsii.String("port"),
//   					RuntimeEnvironmentSecrets: []interface{}{
//   						&KeyValuePairProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					RuntimeEnvironmentVariables: []interface{}{
//   						&KeyValuePairProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					StartCommand: jsii.String("startCommand"),
//   				},
//   			},
//   		},
//   		ImageRepository: &ImageRepositoryProperty{
//   			ImageIdentifier: jsii.String("imageIdentifier"),
//   			ImageRepositoryType: jsii.String("imageRepositoryType"),
//
//   			// the properties below are optional
//   			ImageConfiguration: &ImageConfigurationProperty{
//   				Port: jsii.String("port"),
//   				RuntimeEnvironmentSecrets: []interface{}{
//   					&KeyValuePairProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				RuntimeEnvironmentVariables: []interface{}{
//   					&KeyValuePairProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				StartCommand: jsii.String("startCommand"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	AutoScalingConfigurationArn: jsii.String("autoScalingConfigurationArn"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKey: jsii.String("kmsKey"),
//   	},
//   	HealthCheckConfiguration: &HealthCheckConfigurationProperty{
//   		HealthyThreshold: jsii.Number(123),
//   		Interval: jsii.Number(123),
//   		Path: jsii.String("path"),
//   		Protocol: jsii.String("protocol"),
//   		Timeout: jsii.Number(123),
//   		UnhealthyThreshold: jsii.Number(123),
//   	},
//   	InstanceConfiguration: &InstanceConfigurationProperty{
//   		Cpu: jsii.String("cpu"),
//   		InstanceRoleArn: jsii.String("instanceRoleArn"),
//   		Memory: jsii.String("memory"),
//   	},
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		EgressConfiguration: &EgressConfigurationProperty{
//   			EgressType: jsii.String("egressType"),
//
//   			// the properties below are optional
//   			VpcConnectorArn: jsii.String("vpcConnectorArn"),
//   		},
//   		IngressConfiguration: &IngressConfigurationProperty{
//   			IsPubliclyAccessible: jsii.Boolean(false),
//   		},
//   	},
//   	ObservabilityConfiguration: &ServiceObservabilityConfigurationProperty{
//   		ObservabilityEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		ObservabilityConfigurationArn: jsii.String("observabilityConfigurationArn"),
//   	},
//   	ServiceName: jsii.String("serviceName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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

