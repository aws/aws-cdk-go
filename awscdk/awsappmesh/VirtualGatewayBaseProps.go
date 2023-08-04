package awsappmesh


// Basic configuration properties for a VirtualGateway.
//
// Example:
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
//   				Interval: awscdk.Duration_Seconds(jsii.Number(10)),
//   			}),
//   		}),
//   	},
//   })
//
type VirtualGatewayBaseProps struct {
	// Access Logging Configuration for the VirtualGateway.
	// Default: - no access logging.
	//
	AccessLog AccessLog `field:"optional" json:"accessLog" yaml:"accessLog"`
	// Default Configuration Virtual Node uses to communicate with Virtual Service.
	// Default: - No Config.
	//
	BackendDefaults *BackendDefaults `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// Listeners for the VirtualGateway.
	//
	// Only one is supported.
	// Default: - Single HTTP listener on port 8080.
	//
	Listeners *[]VirtualGatewayListener `field:"optional" json:"listeners" yaml:"listeners"`
	// Name of the VirtualGateway.
	// Default: - A name is automatically determined.
	//
	VirtualGatewayName *string `field:"optional" json:"virtualGatewayName" yaml:"virtualGatewayName"`
}

