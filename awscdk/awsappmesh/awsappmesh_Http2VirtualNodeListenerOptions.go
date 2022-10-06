package awsappmesh


// Represent the HTTP2 Node Listener prorperty.
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
//   http2VirtualNodeListenerOptions := &http2VirtualNodeListenerOptions{
//   	connectionPool: &http2ConnectionPool{
//   		maxRequests: jsii.Number(123),
//   	},
//   	healthCheck: healthCheck,
//   	outlierDetection: &outlierDetection{
//   		baseEjectionDuration: duration,
//   		interval: duration,
//   		maxEjectionPercent: jsii.Number(123),
//   		maxServerErrors: jsii.Number(123),
//   	},
//   	port: jsii.Number(123),
//   	timeout: &httpTimeout{
//   		idle: duration,
//   		perRequest: duration,
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
// Experimental.
type Http2VirtualNodeListenerOptions struct {
	// Connection pool for http2 listeners.
	// Experimental.
	ConnectionPool *Http2ConnectionPool `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	// Experimental.
	HealthCheck HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	// Experimental.
	OutlierDetection *OutlierDetection `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// Port to listen for connections on.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Timeout for HTTP protocol.
	// Experimental.
	Timeout *HttpTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	// Experimental.
	Tls *ListenerTlsOptions `field:"optional" json:"tls" yaml:"tls"`
}

