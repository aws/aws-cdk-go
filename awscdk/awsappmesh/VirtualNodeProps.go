package awsappmesh


// The properties used when creating a new VirtualNode.
//
// Example:
//   var mesh mesh
//   var service service
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
//   	Listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
//   			Port: jsii.Number(8080),
//   			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
//   				HealthyThreshold: jsii.Number(3),
//   				Interval: awscdk.Duration_Seconds(jsii.Number(5)),
//   				Path: jsii.String("/ping"),
//   				Timeout: awscdk.Duration_*Seconds(jsii.Number(2)),
//   				UnhealthyThreshold: jsii.Number(2),
//   			}),
//   			Timeout: &HttpTimeout{
//   				Idle: awscdk.Duration_*Seconds(jsii.Number(5)),
//   			},
//   		}),
//   	},
//   	BackendDefaults: &BackendDefaults{
//   		TlsClientPolicy: &TlsClientPolicy{
//   			Validation: &TlsValidation{
//   				Trust: appmesh.TlsValidationTrust_File(jsii.String("/keys/local_cert_chain.pem")),
//   			},
//   		},
//   	},
//   	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
//   })
//
//   cdk.Tags_Of(node).Add(jsii.String("Environment"), jsii.String("Dev"))
//
type VirtualNodeProps struct {
	// Access Logging Configuration for the virtual node.
	// Default: - No access logging.
	//
	AccessLog AccessLog `field:"optional" json:"accessLog" yaml:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	// Default: - No Config.
	//
	BackendDefaults *BackendDefaults `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// Virtual Services that this is node expected to send outbound traffic to.
	// Default: - No backends.
	//
	Backends *[]Backend `field:"optional" json:"backends" yaml:"backends"`
	// Initial listener for the virtual node.
	// Default: - No listeners.
	//
	Listeners *[]VirtualNodeListener `field:"optional" json:"listeners" yaml:"listeners"`
	// Defines how upstream clients will discover this VirtualNode.
	// Default: - No Service Discovery.
	//
	ServiceDiscovery ServiceDiscovery `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
	// The name of the VirtualNode.
	// Default: - A name is automatically determined.
	//
	VirtualNodeName *string `field:"optional" json:"virtualNodeName" yaml:"virtualNodeName"`
	// The Mesh which the VirtualNode belongs to.
	Mesh IMesh `field:"required" json:"mesh" yaml:"mesh"`
}

