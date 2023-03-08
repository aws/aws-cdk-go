package awsappmesh


// Basic configuration properties for a VirtualNode.
//
// Example:
//   var mesh mesh
//   vpc := ec2.NewVpc(this, jsii.String("vpc"))
//   namespace := cloudmap.NewPrivateDnsNamespace(this, jsii.String("test-namespace"), &privateDnsNamespaceProps{
//   	vpc: vpc,
//   	name: jsii.String("domain.local"),
//   })
//   service := namespace.createService(jsii.String("Svc"))
//   node := mesh.addVirtualNode(jsii.String("virtual-node"), &virtualNodeBaseProps{
//   	serviceDiscovery: appmesh.serviceDiscovery.cloudMap(service),
//   	listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener.http(&httpVirtualNodeListenerOptions{
//   			port: jsii.Number(8081),
//   			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
//   				healthyThreshold: jsii.Number(3),
//   				interval: cdk.duration.seconds(jsii.Number(5)),
//   				 // minimum
//   				path: jsii.String("/health-check-path"),
//   				timeout: cdk.*duration.seconds(jsii.Number(2)),
//   				 // minimum
//   				unhealthyThreshold: jsii.Number(2),
//   			}),
//   		}),
//   	},
//   	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
//   })
//
// Experimental.
type VirtualNodeBaseProps struct {
	// Access Logging Configuration for the virtual node.
	// Experimental.
	AccessLog AccessLog `field:"optional" json:"accessLog" yaml:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	// Experimental.
	BackendDefaults *BackendDefaults `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// Virtual Services that this is node expected to send outbound traffic to.
	// Experimental.
	Backends *[]Backend `field:"optional" json:"backends" yaml:"backends"`
	// Initial listener for the virtual node.
	// Experimental.
	Listeners *[]VirtualNodeListener `field:"optional" json:"listeners" yaml:"listeners"`
	// Defines how upstream clients will discover this VirtualNode.
	// Experimental.
	ServiceDiscovery ServiceDiscovery `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
	// The name of the VirtualNode.
	// Experimental.
	VirtualNodeName *string `field:"optional" json:"virtualNodeName" yaml:"virtualNodeName"`
}

