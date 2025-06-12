package awsecs


// Port mappings allow containers to access ports on the host container instance to send or receive traffic.
//
// Example:
//   var container containerDefinition
//
//
//   container.AddPortMappings(&PortMapping{
//   	ContainerPort: ecs.*containerDefinition_CONTAINER_PORT_USE_RANGE(),
//   	ContainerPortRange: jsii.String("8080-8081"),
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
	//
	// If you want to expose a port range, you must specify `CONTAINER_PORT_USE_RANGE` as container port.
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
	// The protocol used by Service Connect.
	//
	// Valid values are AppProtocol.http, AppProtocol.http2, and
	// AppProtocol.grpc. The protocol determines what telemetry will be shown in the ECS Console for
	// Service Connect services using this port mapping.
	//
	// This field may only be set when the task definition uses Bridge or Awsvpc network modes.
	// Default: - no app protocol.
	//
	AppProtocol AppProtocol `field:"optional" json:"appProtocol" yaml:"appProtocol"`
	// The port number range on the container that's bound to the dynamically mapped host port range.
	//
	// The following rules apply when you specify a `containerPortRange`:
	//
	// - You must specify `CONTAINER_PORT_USE_RANGE` as `containerPort`
	// - You must use either the `bridge` network mode or the `awsvpc` network mode.
	// - The container instance must have at least version 1.67.0 of the container agent and at least version 1.67.0-1 of the `ecs-init` package
	// - You can specify a maximum of 100 port ranges per container.
	// - A port can only be included in one port mapping per container.
	// - You cannot specify overlapping port ranges.
	// - The first port in the range must be less than last port in the range.
	//
	// If you want to expose a single port, you must not set a range.
	ContainerPortRange *string `field:"optional" json:"containerPortRange" yaml:"containerPortRange"`
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
	// Default: - no port mapping name.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The protocol used for the port mapping.
	//
	// Valid values are Protocol.TCP and Protocol.UDP.
	// Default: TCP.
	//
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
}

