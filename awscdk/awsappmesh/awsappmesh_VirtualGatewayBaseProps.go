package awsappmesh


// Basic configuration properties for a VirtualGateway.
//
// Example:
//   var mesh mesh
//
//
//   gateway := mesh.addVirtualGateway(jsii.String("gateway"), &virtualGatewayBaseProps{
//   	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
//   	virtualGatewayName: jsii.String("virtualGateway"),
//   	listeners: []virtualGatewayListener{
//   		appmesh.*virtualGatewayListener.http(&httpGatewayListenerOptions{
//   			port: jsii.Number(443),
//   			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
//   				interval: cdk.duration.seconds(jsii.Number(10)),
//   			}),
//   		}),
//   	},
//   })
//
// Experimental.
type VirtualGatewayBaseProps struct {
	// Access Logging Configuration for the VirtualGateway.
	// Experimental.
	AccessLog AccessLog `field:"optional" json:"accessLog" yaml:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	// Experimental.
	BackendDefaults *BackendDefaults `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// Listeners for the VirtualGateway.
	//
	// Only one is supported.
	// Experimental.
	Listeners *[]VirtualGatewayListener `field:"optional" json:"listeners" yaml:"listeners"`
	// Name of the VirtualGateway.
	// Experimental.
	VirtualGatewayName *string `field:"optional" json:"virtualGatewayName" yaml:"virtualGatewayName"`
}

