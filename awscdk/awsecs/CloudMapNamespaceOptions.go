package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery"
)

// The options for creating an AWS Cloud Map namespace.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var containerOptions containerDefinitionOptions
//
//
//   container := taskDefinition.AddContainer(jsii.String("MyContainer"), containerOptions)
//
//   container.AddPortMappings(&PortMapping{
//   	Name: jsii.String("api"),
//   	ContainerPort: jsii.Number(8080),
//   })
//
//   cluster.AddDefaultCloudMapNamespace(&CloudMapNamespaceOptions{
//   	Name: jsii.String("local"),
//   })
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	ServiceConnectConfiguration: &ServiceConnectProps{
//   		Services: []serviceConnectService{
//   			&serviceConnectService{
//   				PortMappingName: jsii.String("api"),
//   				DnsName: jsii.String("http-api"),
//   				Port: jsii.Number(80),
//   			},
//   		},
//   	},
//   })
//
type CloudMapNamespaceOptions struct {
	// The name of the namespace, such as example.com.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of CloudMap Namespace to create.
	// Default: PrivateDns.
	//
	Type awsservicediscovery.NamespaceType `field:"optional" json:"type" yaml:"type"`
	// This property specifies whether to set the provided namespace as the service connect default in the cluster properties.
	// Default: false.
	//
	UseForServiceConnect *bool `field:"optional" json:"useForServiceConnect" yaml:"useForServiceConnect"`
	// The VPC to associate the namespace with.
	//
	// This property is required for private DNS namespaces.
	// Default: VPC of the cluster for Private DNS Namespace, otherwise none.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

