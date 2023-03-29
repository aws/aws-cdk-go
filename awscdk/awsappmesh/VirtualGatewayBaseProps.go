package awsappmesh


// Basic configuration properties for a VirtualGateway.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var mesh mesh
//
//
//   gateway := mesh.addVirtualGateway(jsii.String("gateway"), &VirtualGatewayBaseProps{
//   	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
//   	VirtualGatewayName: jsii.String("virtualGateway"),
//   	Listeners: []virtualGatewayListener{
//   		appmesh.*virtualGatewayListener_Http(&HttpGatewayListenerOptions{
//   			Port: jsii.Number(443),
//   			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
//   				Interval: cdk.Duration_Seconds(jsii.Number(10)),
//   			}),
//   		}),
//   	},
//   })
//
type VirtualGatewayBaseProps struct {
	// Access Logging Configuration for the VirtualGateway.
	AccessLog AccessLog `field:"optional" json:"accessLog" yaml:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	BackendDefaults *BackendDefaults `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// Listeners for the VirtualGateway.
	//
	// Only one is supported.
	Listeners *[]VirtualGatewayListener `field:"optional" json:"listeners" yaml:"listeners"`
	// Name of the VirtualGateway.
	VirtualGatewayName *string `field:"optional" json:"virtualGatewayName" yaml:"virtualGatewayName"`
}

