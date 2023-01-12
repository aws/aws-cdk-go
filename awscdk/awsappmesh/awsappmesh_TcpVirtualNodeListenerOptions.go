package awsappmesh


// Represent the TCP Node Listener property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var healthCheck healthCheck
//   var mutualTlsValidationTrust mutualTlsValidationTrust
//   var subjectAlternativeNames subjectAlternativeNames
//   var tlsCertificate tlsCertificate
//
//   tcpVirtualNodeListenerOptions := &tcpVirtualNodeListenerOptions{
//   	connectionPool: &tcpConnectionPool{
//   		maxConnections: jsii.Number(123),
//   	},
//   	healthCheck: healthCheck,
//   	outlierDetection: &outlierDetection{
//   		baseEjectionDuration: cdk.duration.minutes(jsii.Number(30)),
//   		interval: cdk.*duration.minutes(jsii.Number(30)),
//   		maxEjectionPercent: jsii.Number(123),
//   		maxServerErrors: jsii.Number(123),
//   	},
//   	port: jsii.Number(123),
//   	timeout: &tcpTimeout{
//   		idle: cdk.*duration.minutes(jsii.Number(30)),
//   	},
//   	tls: &listenerTlsOptions{
//   		certificate: tlsCertificate,
//   		mode: awscdk.Aws_appmesh.tlsMode_STRICT,
//
//   		// the properties below are optional
//   		mutualTlsValidation: &mutualTlsValidation{
//   			trust: mutualTlsValidationTrust,
//
//   			// the properties below are optional
//   			subjectAlternativeNames: subjectAlternativeNames,
//   		},
//   	},
//   }
//
type TcpVirtualNodeListenerOptions struct {
	// Connection pool for http listeners.
	ConnectionPool *TcpConnectionPool `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	HealthCheck HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	OutlierDetection *OutlierDetection `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// Port to listen for connections on.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Timeout for TCP protocol.
	Timeout *TcpTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	Tls *ListenerTlsOptions `field:"optional" json:"tls" yaml:"tls"`
}

