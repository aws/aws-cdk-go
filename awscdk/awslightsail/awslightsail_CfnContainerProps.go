package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnContainer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContainerProps := &cfnContainerProps{
//   	power: jsii.String("power"),
//   	scale: jsii.Number(123),
//   	serviceName: jsii.String("serviceName"),
//
//   	// the properties below are optional
//   	containerServiceDeployment: &containerServiceDeploymentProperty{
//   		containers: []interface{}{
//   			&containerProperty{
//   				command: []*string{
//   					jsii.String("command"),
//   				},
//   				containerName: jsii.String("containerName"),
//   				environment: []interface{}{
//   					&environmentVariableProperty{
//   						value: jsii.String("value"),
//   						variable: jsii.String("variable"),
//   					},
//   				},
//   				image: jsii.String("image"),
//   				ports: []interface{}{
//   					&portInfoProperty{
//   						port: jsii.String("port"),
//   						protocol: jsii.String("protocol"),
//   					},
//   				},
//   			},
//   		},
//   		publicEndpoint: &publicEndpointProperty{
//   			containerName: jsii.String("containerName"),
//   			containerPort: jsii.Number(123),
//   			healthCheckConfig: &healthCheckConfigProperty{
//   				healthyThreshold: jsii.Number(123),
//   				intervalSeconds: jsii.Number(123),
//   				path: jsii.String("path"),
//   				successCodes: jsii.String("successCodes"),
//   				timeoutSeconds: jsii.Number(123),
//   				unhealthyThreshold: jsii.Number(123),
//   			},
//   		},
//   	},
//   	isDisabled: jsii.Boolean(false),
//   	publicDomainNames: []interface{}{
//   		&publicDomainNameProperty{
//   			certificateName: jsii.String("certificateName"),
//   			domainNames: []*string{
//   				jsii.String("domainNames"),
//   			},
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
type CfnContainerProps struct {
	// The power specification of the container service.
	//
	// The power specifies the amount of RAM, the number of vCPUs, and the base price of the container service.
	Power *string `field:"required" json:"power" yaml:"power"`
	// The scale specification of the container service.
	//
	// The scale specifies the allocated compute nodes of the container service.
	Scale *float64 `field:"required" json:"scale" yaml:"scale"`
	// The name of the container service.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// An object that describes the current container deployment of the container service.
	ContainerServiceDeployment interface{} `field:"optional" json:"containerServiceDeployment" yaml:"containerServiceDeployment"`
	// A Boolean value indicating whether the container service is disabled.
	IsDisabled interface{} `field:"optional" json:"isDisabled" yaml:"isDisabled"`
	// The public domain name of the container service, such as `example.com` and `www.example.com` .
	//
	// You can specify up to four public domain names for a container service. The domain names that you specify are used when you create a deployment with a container that is configured as the public endpoint of your container service.
	//
	// If you don't specify public domain names, then you can use the default domain of the container service.
	//
	// > You must create and validate an SSL/TLS certificate before you can use public domain names with your container service. Use the [AWS::Lightsail::Certificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-certificate.html) resource to create a certificate for the public domain names that you want to use with your container service.
	PublicDomainNames interface{} `field:"optional" json:"publicDomainNames" yaml:"publicDomainNames"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

