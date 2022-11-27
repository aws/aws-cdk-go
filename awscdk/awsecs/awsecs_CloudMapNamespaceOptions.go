package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery"
)

// The options for creating an AWS Cloud Map namespace.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var container containerDefinition
//
//
//   container.addPortMappings(&portMapping{
//   	name: jsii.String("api"),
//   	containerPort: jsii.Number(8080),
//   })
//
//   taskDefinition.addContainer(container)
//
//   cluster.addDefaultCloudMapNamespace(&cloudMapNamespaceOptions{
//   	name: jsii.String("local"),
//   })
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	serviceConnectConfiguration: &serviceConnectProps{
//   		services: []serviceConnectService{
//   			&serviceConnectService{
//   				portMappingName: jsii.String("api"),
//   				dnsName: jsii.String("http-api"),
//   				port: jsii.Number(80),
//   			},
//   		},
//   	},
//   })
//
type CloudMapNamespaceOptions struct {
	// The name of the namespace, such as example.com.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of CloudMap Namespace to create.
	Type awsservicediscovery.NamespaceType `field:"optional" json:"type" yaml:"type"`
	// This property specifies whether to set the provided namespace as the service connect default in the cluster properties.
	UseForServiceConnect *bool `field:"optional" json:"useForServiceConnect" yaml:"useForServiceConnect"`
	// The VPC to associate the namespace with.
	//
	// This property is required for private DNS namespaces.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

