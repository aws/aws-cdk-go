package awsappmesh


// Represent the HTTP2 Node Listener property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var healthCheck HealthCheck
//   var mutualTlsValidationTrust MutualTlsValidationTrust
//   var subjectAlternativeNames SubjectAlternativeNames
//   var tlsCertificate TlsCertificate
//
//   http2VirtualNodeListenerOptions := &Http2VirtualNodeListenerOptions{
//   	ConnectionPool: &Http2ConnectionPool{
//   		MaxRequests: jsii.Number(123),
//   	},
//   	HealthCheck: healthCheck,
//   	OutlierDetection: &OutlierDetection{
//   		BaseEjectionDuration: cdk.Duration_Minutes(jsii.Number(30)),
//   		Interval: cdk.Duration_*Minutes(jsii.Number(30)),
//   		MaxEjectionPercent: jsii.Number(123),
//   		MaxServerErrors: jsii.Number(123),
//   	},
//   	Port: jsii.Number(123),
//   	Timeout: &HttpTimeout{
//   		Idle: cdk.Duration_*Minutes(jsii.Number(30)),
//   		PerRequest: cdk.Duration_*Minutes(jsii.Number(30)),
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
type Http2VirtualNodeListenerOptions struct {
	// Connection pool for http2 listeners.
	// Default: - None.
	//
	ConnectionPool *Http2ConnectionPool `field:"optional" json:"connectionPool" yaml:"connectionPool"`
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

