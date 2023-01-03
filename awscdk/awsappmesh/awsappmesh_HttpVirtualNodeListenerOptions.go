package awsappmesh


// Represent the HTTP Node Listener property.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
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
type HttpVirtualNodeListenerOptions struct {
	// Connection pool for http listeners.
	ConnectionPool *HttpConnectionPool `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	HealthCheck HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	OutlierDetection *OutlierDetection `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// Port to listen for connections on.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Timeout for HTTP protocol.
	Timeout *HttpTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	Tls *ListenerTlsOptions `field:"optional" json:"tls" yaml:"tls"`
}

