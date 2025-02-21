package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnContainer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContainerProps := &CfnContainerProps{
//   	Power: jsii.String("power"),
//   	Scale: jsii.Number(123),
//   	ServiceName: jsii.String("serviceName"),
//
//   	// the properties below are optional
//   	ContainerServiceDeployment: &ContainerServiceDeploymentProperty{
//   		Containers: []interface{}{
//   			&ContainerProperty{
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				ContainerName: jsii.String("containerName"),
//   				Environment: []interface{}{
//   					&EnvironmentVariableProperty{
//   						Value: jsii.String("value"),
//   						Variable: jsii.String("variable"),
//   					},
//   				},
//   				Image: jsii.String("image"),
//   				Ports: []interface{}{
//   					&PortInfoProperty{
//   						Port: jsii.String("port"),
//   						Protocol: jsii.String("protocol"),
//   					},
//   				},
//   			},
//   		},
//   		PublicEndpoint: &PublicEndpointProperty{
//   			ContainerName: jsii.String("containerName"),
//   			ContainerPort: jsii.Number(123),
//   			HealthCheckConfig: &HealthCheckConfigProperty{
//   				HealthyThreshold: jsii.Number(123),
//   				IntervalSeconds: jsii.Number(123),
//   				Path: jsii.String("path"),
//   				SuccessCodes: jsii.String("successCodes"),
//   				TimeoutSeconds: jsii.Number(123),
//   				UnhealthyThreshold: jsii.Number(123),
//   			},
//   		},
//   	},
//   	IsDisabled: jsii.Boolean(false),
//   	PrivateRegistryAccess: &PrivateRegistryAccessProperty{
//   		EcrImagePullerRole: &EcrImagePullerRoleProperty{
//   			IsActive: jsii.Boolean(false),
//   			PrincipalArn: jsii.String("principalArn"),
//   		},
//   	},
//   	PublicDomainNames: []interface{}{
//   		&PublicDomainNameProperty{
//   			CertificateName: jsii.String("certificateName"),
//   			DomainNames: []*string{
//   				jsii.String("domainNames"),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html
//
type CfnContainerProps struct {
	// The power specification of the container service.
	//
	// The power specifies the amount of RAM, the number of vCPUs, and the base price of the container service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html#cfn-lightsail-container-power
	//
	Power *string `field:"required" json:"power" yaml:"power"`
	// The scale specification of the container service.
	//
	// The scale specifies the allocated compute nodes of the container service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html#cfn-lightsail-container-scale
	//
	Scale *float64 `field:"required" json:"scale" yaml:"scale"`
	// The name of the container service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html#cfn-lightsail-container-servicename
	//
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// An object that describes the current container deployment of the container service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html#cfn-lightsail-container-containerservicedeployment
	//
	ContainerServiceDeployment interface{} `field:"optional" json:"containerServiceDeployment" yaml:"containerServiceDeployment"`
	// A Boolean value indicating whether the container service is disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html#cfn-lightsail-container-isdisabled
	//
	IsDisabled interface{} `field:"optional" json:"isDisabled" yaml:"isDisabled"`
	// An object that describes the configuration for the container service to access private container image repositories, such as Amazon Elastic Container Registry ( Amazon ECR ) private repositories.
	//
	// For more information, see [Configuring access to an Amazon ECR private repository for an Amazon Lightsail container service](https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-container-service-ecr-private-repo-access) in the *Amazon Lightsail Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html#cfn-lightsail-container-privateregistryaccess
	//
	PrivateRegistryAccess interface{} `field:"optional" json:"privateRegistryAccess" yaml:"privateRegistryAccess"`
	// The public domain name of the container service, such as `example.com` and `www.example.com` .
	//
	// You can specify up to four public domain names for a container service. The domain names that you specify are used when you create a deployment with a container that is configured as the public endpoint of your container service.
	//
	// If you don't specify public domain names, then you can use the default domain of the container service.
	//
	// > You must create and validate an SSL/TLS certificate before you can use public domain names with your container service. Use the [AWS::Lightsail::Certificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-certificate.html) resource to create a certificate for the public domain names that you want to use with your container service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html#cfn-lightsail-container-publicdomainnames
	//
	PublicDomainNames interface{} `field:"optional" json:"publicDomainNames" yaml:"publicDomainNames"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html#cfn-lightsail-container-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

