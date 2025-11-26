package previewawsapprunnermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnServicePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServiceMixinProps := &CfnServiceMixinProps{
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
//   			VpcConnectorArn: jsii.String("vpcConnectorArn"),
//   		},
//   		IngressConfiguration: &IngressConfigurationProperty{
//   			IsPubliclyAccessible: jsii.Boolean(false),
//   		},
//   		IpAddressType: jsii.String("ipAddressType"),
//   	},
//   	ObservabilityConfiguration: &ServiceObservabilityConfigurationProperty{
//   		ObservabilityConfigurationArn: jsii.String("observabilityConfigurationArn"),
//   		ObservabilityEnabled: jsii.Boolean(false),
//   	},
//   	ServiceName: jsii.String("serviceName"),
//   	SourceConfiguration: &SourceConfigurationProperty{
//   		AuthenticationConfiguration: &AuthenticationConfigurationProperty{
//   			AccessRoleArn: jsii.String("accessRoleArn"),
//   			ConnectionArn: jsii.String("connectionArn"),
//   		},
//   		AutoDeploymentsEnabled: jsii.Boolean(false),
//   		CodeRepository: &CodeRepositoryProperty{
//   			CodeConfiguration: &CodeConfigurationProperty{
//   				CodeConfigurationValues: &CodeConfigurationValuesProperty{
//   					BuildCommand: jsii.String("buildCommand"),
//   					Port: jsii.String("port"),
//   					Runtime: jsii.String("runtime"),
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
//   				ConfigurationSource: jsii.String("configurationSource"),
//   			},
//   			RepositoryUrl: jsii.String("repositoryUrl"),
//   			SourceCodeVersion: &SourceCodeVersionProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   			SourceDirectory: jsii.String("sourceDirectory"),
//   		},
//   		ImageRepository: &ImageRepositoryProperty{
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
//   			ImageIdentifier: jsii.String("imageIdentifier"),
//   			ImageRepositoryType: jsii.String("imageRepositoryType"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html
//
type CfnServiceMixinProps struct {
	// The Amazon Resource Name (ARN) of an App Runner automatic scaling configuration resource that you want to associate with your service.
	//
	// If not provided, App Runner associates the latest revision of a default auto scaling configuration.
	//
	// Specify an ARN with a name and a revision number to associate that revision. For example: `arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability/3`
	//
	// Specify just the name to associate the latest revision. For example: `arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-autoscalingconfigurationarn
	//
	AutoScalingConfigurationArn *string `field:"optional" json:"autoScalingConfigurationArn" yaml:"autoScalingConfigurationArn"`
	// An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs.
	//
	// By default, App Runner uses an AWS managed key .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The settings for the health check that AWS App Runner performs to monitor the health of the App Runner service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-healthcheckconfiguration
	//
	HealthCheckConfiguration interface{} `field:"optional" json:"healthCheckConfiguration" yaml:"healthCheckConfiguration"`
	// The runtime configuration of instances (scaling units) of your service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-instanceconfiguration
	//
	InstanceConfiguration interface{} `field:"optional" json:"instanceConfiguration" yaml:"instanceConfiguration"`
	// Configuration settings related to network traffic of the web application that the App Runner service runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The observability configuration of your service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-observabilityconfiguration
	//
	ObservabilityConfiguration interface{} `field:"optional" json:"observabilityConfiguration" yaml:"observabilityConfiguration"`
	// A name for the App Runner service.
	//
	// It must be unique across all the running App Runner services in your AWS account in the AWS Region .
	//
	// If you don't specify a name, CloudFormation generates a name for your service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-servicename
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The source to deploy to the App Runner service.
	//
	// It can be a code or an image repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-sourceconfiguration
	//
	SourceConfiguration interface{} `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// An optional list of metadata items that you can associate with the App Runner service resource.
	//
	// A tag is a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

