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
//   serviceConnectConfigurationProperty := &ServiceConnectConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	LogConfiguration: &LogConfigurationProperty{
//   		LogDriver: jsii.String("logDriver"),
//   		Options: map[string]*string{
//   			"optionsKey": jsii.String("options"),
//   		},
//   		SecretOptions: []interface{}{
//   			&SecretProperty{
//   				Name: jsii.String("name"),
//   				ValueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   	},
//   	Namespace: jsii.String("namespace"),
//   	Services: []interface{}{
//   		&ServiceConnectServiceProperty{
//   			PortName: jsii.String("portName"),
//
//   			// the properties below are optional
//   			ClientAliases: []interface{}{
//   				&ServiceConnectClientAliasProperty{
//   					Port: jsii.Number(123),
//
//   					// the properties below are optional
//   					DnsName: jsii.String("dnsName"),
//   				},
//   			},
//   			DiscoveryName: jsii.String("discoveryName"),
//   			IngressPortOverride: jsii.Number(123),
//   			Timeout: &TimeoutConfigurationProperty{
//   				IdleTimeoutSeconds: jsii.Number(123),
//   				PerRequestTimeoutSeconds: jsii.Number(123),
//   			},
//   			Tls: &ServiceConnectTlsConfigurationProperty{
//   				IssuerCertificateAuthority: &ServiceConnectTlsCertificateAuthorityProperty{
//   					AwsPcaAuthorityArn: jsii.String("awsPcaAuthorityArn"),
//   				},
//
//   				// the properties below are optional
//   				KmsKey: jsii.String("kmsKey"),
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnectconfiguration.html
//
type CfnService_ServiceConnectConfigurationProperty struct {
	// Specifies whether to use Service Connect with this service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnectconfiguration.html#cfn-ecs-service-serviceconnectconfiguration-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The log configuration for the container.
	//
	// This parameter maps to `LogConfig` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--log-driver` option to [`docker run`](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/commandline/run/) .
	//
	// By default, containers use the same logging driver that the Docker daemon uses. However, the container might use a different logging driver than the Docker daemon by specifying a log driver configuration in the container definition. For more information about the options for different supported log drivers, see [Configure logging drivers](https://docs.aws.amazon.com/https://docs.docker.com/engine/admin/logging/overview/) in the Docker documentation.
	//
	// Understand the following when specifying a log configuration for your containers.
	//
	// - Amazon ECS currently supports a subset of the logging drivers available to the Docker daemon. Additional log drivers may be available in future releases of the Amazon ECS container agent.
	//
	// For tasks on AWS Fargate , the supported log drivers are `awslogs` , `splunk` , and `awsfirelens` .
	//
	// For tasks hosted on Amazon EC2 instances, the supported log drivers are `awslogs` , `fluentd` , `gelf` , `json-file` , `journald` , `logentries` , `syslog` , `splunk` , and `awsfirelens` .
	// - This parameter requires version 1.18 of the Docker Remote API or greater on your container instance.
	// - For tasks that are hosted on Amazon EC2 instances, the Amazon ECS container agent must register the available logging drivers with the `ECS_AVAILABLE_LOGGING_DRIVERS` environment variable before containers placed on that instance can use these log configuration options. For more information, see [Amazon ECS container agent configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html) in the *Amazon Elastic Container Service Developer Guide* .
	// - For tasks that are on AWS Fargate , because you don't have access to the underlying infrastructure your tasks are hosted on, any additional software needed must be installed outside of the task. For example, the Fluentd output aggregators or a remote host running Logstash to send Gelf logs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnectconfiguration.html#cfn-ecs-service-serviceconnectconfiguration-logconfiguration
	//
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The namespace name or full Amazon Resource Name (ARN) of the AWS Cloud Map namespace for use with Service Connect.
	//
	// The namespace must be in the same AWS Region as the Amazon ECS service and cluster. The type of namespace doesn't affect Service Connect. For more information about AWS Cloud Map , see [Working with Services](https://docs.aws.amazon.com/cloud-map/latest/dg/working-with-services.html) in the *AWS Cloud Map Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnectconfiguration.html#cfn-ecs-service-serviceconnectconfiguration-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The list of Service Connect service objects.
	//
	// These are names and aliases (also known as endpoints) that are used by other Amazon ECS services to connect to this service.
	//
	// This field is not required for a "client" Amazon ECS service that's a member of a namespace only to connect to other services within the namespace. An example of this would be a frontend application that accepts incoming requests from either a load balancer that's attached to the service or by other means.
	//
	// An object selects a port from the task definition, assigns a name for the AWS Cloud Map service, and a list of aliases (endpoints) and ports for client applications to refer to this service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnectconfiguration.html#cfn-ecs-service-serviceconnectconfiguration-services
	//
	Services interface{} `field:"optional" json:"services" yaml:"services"`
}

