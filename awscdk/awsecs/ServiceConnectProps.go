package awsecs


// Interface for Service Connect configuration.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//
//
//   customService := ecs.NewFargateService(this, jsii.String("CustomizedService"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	ServiceConnectConfiguration: &ServiceConnectProps{
//   		LogDriver: ecs.LogDrivers_AwsLogs(&AwsLogDriverProps{
//   			StreamPrefix: jsii.String("sc-traffic"),
//   		}),
//   		Services: []serviceConnectService{
//   			&serviceConnectService{
//   				PortMappingName: jsii.String("api"),
//   				DnsName: jsii.String("customized-api"),
//   				Port: jsii.Number(80),
//   				IngressPortOverride: jsii.Number(20040),
//   				DiscoveryName: jsii.String("custom"),
//   			},
//   		},
//   	},
//   })
//
type ServiceConnectProps struct {
	// The log driver configuration to use for the Service Connect agent logs.
	// Default: - none.
	//
	LogDriver LogDriver `field:"optional" json:"logDriver" yaml:"logDriver"`
	// The cloudmap namespace to register this service into.
	// Default: the cloudmap namespace specified on the cluster.
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The list of Services, including a port mapping, terse client alias, and optional intermediate DNS name.
	//
	// This property may be left blank if the current ECS service does not need to advertise any ports via Service Connect.
	// Default: none.
	//
	Services *[]*ServiceConnectService `field:"optional" json:"services" yaml:"services"`
}

