package awsecs


// The Service Connect configuration of your Amazon ECS service.
//
// The configuration for this service to discover and connect to services, and be discovered by, and connected from, other services within a namespace.
//
// Tasks that run in a namespace can use short names to connect to services in the namespace. Tasks can connect to services across all of the clusters in the namespace. Tasks connect through a managed proxy container that collects logs and metrics for increased visibility. Only the tasks that Amazon ECS services create are supported with Service Connect. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceConnectConfigurationProperty := &serviceConnectConfigurationProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	logConfiguration: &logConfigurationProperty{
//   		logDriver: jsii.String("logDriver"),
//   		options: map[string]*string{
//   			"optionsKey": jsii.String("options"),
//   		},
//   		secretOptions: []interface{}{
//   			&secretProperty{
//   				name: jsii.String("name"),
//   				valueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   	},
//   	namespace: jsii.String("namespace"),
//   	services: []interface{}{
//   		&serviceConnectServiceProperty{
//   			portName: jsii.String("portName"),
//
//   			// the properties below are optional
//   			clientAliases: []interface{}{
//   				&serviceConnectClientAliasProperty{
//   					port: jsii.Number(123),
//
//   					// the properties below are optional
//   					dnsName: jsii.String("dnsName"),
//   				},
//   			},
//   			discoveryName: jsii.String("discoveryName"),
//   			ingressPortOverride: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnService_ServiceConnectConfigurationProperty struct {
	// Specifies whether to use Service Connect with this service.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// `CfnService.ServiceConnectConfigurationProperty.LogConfiguration`.
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The namespace name or full Amazon Resource Name (ARN) of the AWS Cloud Map namespace for use with Service Connect.
	//
	// The namespace must be in the same AWS Region as the Amazon ECS service and cluster. The type of namespace doesn't affect Service Connect. For more information about AWS Cloud Map , see [Working with Services](https://docs.aws.amazon.com/) in the *AWS Cloud Map Developer Guide* .
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The list of Service Connect service objects.
	//
	// These are names and aliases (also known as endpoints) that are used by other Amazon ECS services to connect to this service.
	//
	// This field is not required for a "client" Amazon ECS service that's a member of a namespace only to connect to other services within the namespace. An example of this would be a frontend application that accepts incoming requests from either a load balancer that's attached to the service or by other means.
	//
	// An object selects a port from the task definition, assigns a name for the AWS Cloud Map service, and a list of aliases (endpoints) and ports for client applications to refer to this service.
	Services interface{} `field:"optional" json:"services" yaml:"services"`
}

