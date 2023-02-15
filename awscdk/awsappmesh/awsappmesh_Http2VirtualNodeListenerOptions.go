package awsappmesh


// Represent the HTTP2 Node Listener property.
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
//   http2VirtualNodeListenerOptions := &http2VirtualNodeListenerOptions{
//   	connectionPool: &http2ConnectionPool{
//   		maxRequests: jsii.Number(123),
//   	},
//   	healthCheck: healthCheck,
//   	outlierDetection: &outlierDetection{
//   		baseEjectionDuration: cdk.duration.minutes(jsii.Number(30)),
//   		interval: cdk.*duration.minutes(jsii.Number(30)),
//   		maxEjectionPercent: jsii.Number(123),
//   		maxServerErrors: jsii.Number(123),
//   	},
//   	port: jsii.Number(123),
//   	timeout: &httpTimeout{
//   		idle: cdk.*duration.minutes(jsii.Number(30)),
//   		perRequest: cdk.*duration.minutes(jsii.Number(30)),
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
type Http2VirtualNodeListenerOptions struct {
	// Connection pool for http2 listeners.
	ConnectionPool *Http2ConnectionPool `field:"optional" json:"connectionPool" yaml:"connectionPool"`
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

