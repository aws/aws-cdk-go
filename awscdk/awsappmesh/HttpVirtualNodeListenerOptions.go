package awsappmesh


// Represent the HTTP Node Listener property.
//
// Example:
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
//   				Interval: awscdk.Duration_Seconds(jsii.Number(5)),
//   				 // minimum
//   				Path: jsii.String("/health-check-path"),
//   				Timeout: awscdk.Duration_*Seconds(jsii.Number(2)),
//   				 // minimum
//   				UnhealthyThreshold: jsii.Number(2),
//   			}),
//   		}),
//   	},
//   	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
//   })
//
type HttpVirtualNodeListenerOptions struct {
	// Connection pool for http listeners.
	// Default: - None.
	//
	ConnectionPool *HttpConnectionPool `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	// Default: - no healthcheck.
	//
	HealthCheck HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	// Default: - none.
	//
	OutlierDetection *OutlierDetection `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// Port to listen for connections on.
	// Default: - 8080.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Timeout for HTTP protocol.
	// Default: - None.
	//
	Timeout *HttpTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	// Default: - none.
	//
	Tls *ListenerTlsOptions `field:"optional" json:"tls" yaml:"tls"`
}

