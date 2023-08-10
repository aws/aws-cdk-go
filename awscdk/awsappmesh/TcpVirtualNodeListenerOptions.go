package awsappmesh


// Represent the TCP Node Listener property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
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
//   		BaseEjectionDuration: cdk.Duration_Minutes(jsii.Number(30)),
//   		Interval: cdk.Duration_*Minutes(jsii.Number(30)),
//   		MaxEjectionPercent: jsii.Number(123),
//   		MaxServerErrors: jsii.Number(123),
//   	},
//   	Port: jsii.Number(123),
//   	Timeout: &TcpTimeout{
//   		Idle: cdk.Duration_*Minutes(jsii.Number(30)),
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

