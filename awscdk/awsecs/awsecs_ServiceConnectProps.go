package awsecs


// Interface for Service Connect configuration.
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
type ServiceConnectProps struct {
	// The log driver configuration to use for the Service Connect agent logs.
	LogDriver LogDriver `field:"optional" json:"logDriver" yaml:"logDriver"`
	// The cloudmap namespace to register this service into.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The list of Services, including a port mapping, terse client alias, and optional intermediate DNS name.
	//
	// This property may be left blank if the current ECS service does not need to advertise any ports via Service Connect.
	Services *[]*ServiceConnectService `field:"optional" json:"services" yaml:"services"`
}

