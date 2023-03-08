package awsappmesh


// Basic configuration properties for a VirtualNode.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var mesh mesh
//   vpc := ec2.NewVpc(this, jsii.String("vpc"))
//   namespace := cloudmap.NewPrivateDnsNamespace(this, jsii.String("test-namespace"), &PrivateDnsNamespaceProps{
//   	Vpc: Vpc,
//   	Name: jsii.String("domain.local"),
//   })
//   service := namespace.CreateService(jsii.String("Svc"))
//   node := mesh.addVirtualNode(jsii.String("virtual-node"), &VirtualNodeBaseProps{
//   	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
//   	Listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
//   			Port: jsii.Number(8081),
//   			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
//   				HealthyThreshold: jsii.Number(3),
//   				Interval: cdk.Duration_Seconds(jsii.Number(5)),
//   				 // minimum
//   				Path: jsii.String("/health-check-path"),
//   				Timeout: cdk.Duration_*Seconds(jsii.Number(2)),
//   				 // minimum
//   				UnhealthyThreshold: jsii.Number(2),
//   			}),
//   		}),
//   	},
//   	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
//   })
//
type VirtualNodeBaseProps struct {
	// Access Logging Configuration for the virtual node.
	AccessLog AccessLog `field:"optional" json:"accessLog" yaml:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	BackendDefaults *BackendDefaults `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// Virtual Services that this is node expected to send outbound traffic to.
	Backends *[]Backend `field:"optional" json:"backends" yaml:"backends"`
	// Initial listener for the virtual node.
	Listeners *[]VirtualNodeListener `field:"optional" json:"listeners" yaml:"listeners"`
	// Defines how upstream clients will discover this VirtualNode.
	ServiceDiscovery ServiceDiscovery `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
	// The name of the VirtualNode.
	VirtualNodeName *string `field:"optional" json:"virtualNodeName" yaml:"virtualNodeName"`
}

