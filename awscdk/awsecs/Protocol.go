package awsecs


// Network protocol.
//
// Example:
//   var taskDefinition taskDefinition
//   var cluster cluster
//
//
//   // Add a container to the task definition
//   specificContainer := taskDefinition.AddContainer(jsii.String("Container"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("/aws/aws-example-app")),
//   	MemoryLimitMiB: jsii.Number(2048),
//   })
//
//   // Add a port mapping
//   specificContainer.AddPortMappings(&PortMapping{
//   	ContainerPort: jsii.Number(7600),
//   	Protocol: ecs.Protocol_TCP,
//   })
//
//   ecs.NewEc2Service(this, jsii.String("Service"), &Ec2ServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	CloudMapOptions: &CloudMapOptions{
//   		// Create SRV records - useful for bridge networking
//   		DnsRecordType: cloudmap.DnsRecordType_SRV,
//   		// Targets port TCP port 7600 `specificContainer`
//   		Container: specificContainer,
//   		ContainerPort: jsii.Number(7600),
//   	},
//   })
//
type Protocol string

const (
	// TCP.
	Protocol_TCP Protocol = "TCP"
	// UDP.
	Protocol_UDP Protocol = "UDP"
)

