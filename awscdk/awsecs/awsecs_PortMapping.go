package awsecs


// Port mappings allow containers to access ports on the host container instance to send or receive traffic.
//
// Example:
//   var taskDefinition taskDefinition
//   var cluster cluster
//
//
//   // Add a container to the task definition
//   specificContainer := taskDefinition.addContainer(jsii.String("Container"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("/aws/aws-example-app")),
//   	memoryLimitMiB: jsii.Number(2048),
//   })
//
//   // Add a port mapping
//   specificContainer.addPortMappings(&portMapping{
//   	containerPort: jsii.Number(7600),
//   	protocol: ecs.protocol_TCP,
//   })
//
//   ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	cloudMapOptions: &cloudMapOptions{
//   		// Create SRV records - useful for bridge networking
//   		dnsRecordType: cloudmap.dnsRecordType_SRV,
//   		// Targets port TCP port 7600 `specificContainer`
//   		container: specificContainer,
//   		containerPort: jsii.Number(7600),
//   	},
//   })
//
type PortMapping struct {
	// The port number on the container that is bound to the user-specified or automatically assigned host port.
	//
	// If you are using containers in a task with the awsvpc or host network mode, exposed ports should be specified using containerPort.
	// If you are using containers in a task with the bridge network mode and you specify a container port and not a host port,
	// your container automatically receives a host port in the ephemeral port range.
	//
	// For more information, see hostPort.
	// Port mappings that are automatically assigned in this way do not count toward the 100 reserved ports limit of a container instance.
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
	// The protocol used by Service Connect.
	//
	// Valid values are AppProtocol.http, AppProtocol.http2, and
	// AppProtocol.grpc. The protocol determines what telemetry will be shown in the ECS Console for
	// Service Connect services using this port mapping.
	//
	// This field may only be set when the task definition uses Bridge or Awsvpc network modes.
	AppProtocol AppProtocol `field:"optional" json:"appProtocol" yaml:"appProtocol"`
	// The port number on the container instance to reserve for your container.
	//
	// If you are using containers in a task with the awsvpc or host network mode,
	// the hostPort can either be left blank or set to the same value as the containerPort.
	//
	// If you are using containers in a task with the bridge network mode,
	// you can specify a non-reserved host port for your container port mapping, or
	// you can omit the hostPort (or set it to 0) while specifying a containerPort and
	// your container automatically receives a port in the ephemeral port range for
	// your container instance operating system and Docker version.
	HostPort *float64 `field:"optional" json:"hostPort" yaml:"hostPort"`
	// The name to give the port mapping.
	//
	// Name is required in order to use the port mapping with ECS Service Connect.
	// This field may only be set when the task definition uses Bridge or Awsvpc network modes.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The protocol used for the port mapping.
	//
	// Valid values are Protocol.TCP and Protocol.UDP.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
}

