package awsecs


// Port mappings allow containers to access ports on the host container instance to send or receive traffic.
//
// Example:
//   var container containerDefinition
//
//
//   container.AddPortMappings(&PortMapping{
//   	ContainerPort: jsii.Number(3000),
//   })
//
// Experimental.
type PortMapping struct {
	// The port number on the container that is bound to the user-specified or automatically assigned host port.
	//
	// If you are using containers in a task with the awsvpc or host network mode, exposed ports should be specified using containerPort.
	// If you are using containers in a task with the bridge network mode and you specify a container port and not a host port,
	// your container automatically receives a host port in the ephemeral port range.
	//
	// For more information, see hostPort.
	// Port mappings that are automatically assigned in this way do not count toward the 100 reserved ports limit of a container instance.
	// Experimental.
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
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
	// Experimental.
	HostPort *float64 `field:"optional" json:"hostPort" yaml:"hostPort"`
	// The protocol used for the port mapping.
	//
	// Valid values are Protocol.TCP and Protocol.UDP.
	// Experimental.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
}

