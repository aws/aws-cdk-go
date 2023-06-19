package awsappmesh


// Represent the TCP Node Listener prorperty.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var healthCheck healthCheck
//   var mutualTlsValidationTrust mutualTlsValidationTrust
//   var subjectAlternativeNames subjectAlternativeNames
//   var tlsCertificate tlsCertificate
//
//   tcpVirtualNodeListenerOptions := &TcpVirtualNodeListenerOptions{
//   	ConnectionPool: &TcpConnectionPool{
//   		MaxConnections: jsii.Number(123),
//   	},
//   	HealthCheck: healthCheck,
//   	OutlierDetection: &OutlierDetection{
//   		BaseEjectionDuration: duration,
//   		Interval: duration,
//   		MaxEjectionPercent: jsii.Number(123),
//   		MaxServerErrors: jsii.Number(123),
//   	},
//   	Port: jsii.Number(123),
//   	Timeout: &TcpTimeout{
//   		Idle: duration,
//   	},
//   	Tls: &ListenerTlsOptions{
//   		Certificate: tlsCertificate,
//   		Mode: awscdk.Aws_appmesh.TlsMode_STRICT,
//
//   		// the properties below are optional
//   		MutualTlsValidation: &MutualTlsValidation{
//   			Trust: mutualTlsValidationTrust,
//
//   			// the properties below are optional
//   			SubjectAlternativeNames: subjectAlternativeNames,
//   		},
//   	},
//   }
//
// Experimental.
type TcpVirtualNodeListenerOptions struct {
	// Connection pool for http listeners.
	// Experimental.
	ConnectionPool *TcpConnectionPool `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	// Experimental.
	HealthCheck HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	// Experimental.
	OutlierDetection *OutlierDetection `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// Port to listen for connections on.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Timeout for TCP protocol.
	// Experimental.
	Timeout *TcpTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	// Experimental.
	Tls *ListenerTlsOptions `field:"optional" json:"tls" yaml:"tls"`
}

