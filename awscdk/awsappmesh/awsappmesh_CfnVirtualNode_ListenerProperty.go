package awsappmesh


// An object that represents a listener for a virtual node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerProperty := &listenerProperty{
//   	portMapping: &portMappingProperty{
//   		port: jsii.Number(123),
//   		protocol: jsii.String("protocol"),
//   	},
//
//   	// the properties below are optional
//   	connectionPool: &virtualNodeConnectionPoolProperty{
//   		grpc: &virtualNodeGrpcConnectionPoolProperty{
//   			maxRequests: jsii.Number(123),
//   		},
//   		http: &virtualNodeHttpConnectionPoolProperty{
//   			maxConnections: jsii.Number(123),
//
//   			// the properties below are optional
//   			maxPendingRequests: jsii.Number(123),
//   		},
//   		http2: &virtualNodeHttp2ConnectionPoolProperty{
//   			maxRequests: jsii.Number(123),
//   		},
//   		tcp: &virtualNodeTcpConnectionPoolProperty{
//   			maxConnections: jsii.Number(123),
//   		},
//   	},
//   	healthCheck: &healthCheckProperty{
//   		healthyThreshold: jsii.Number(123),
//   		intervalMillis: jsii.Number(123),
//   		protocol: jsii.String("protocol"),
//   		timeoutMillis: jsii.Number(123),
//   		unhealthyThreshold: jsii.Number(123),
//
//   		// the properties below are optional
//   		path: jsii.String("path"),
//   		port: jsii.Number(123),
//   	},
//   	outlierDetection: &outlierDetectionProperty{
//   		baseEjectionDuration: &durationProperty{
//   			unit: jsii.String("unit"),
//   			value: jsii.Number(123),
//   		},
//   		interval: &durationProperty{
//   			unit: jsii.String("unit"),
//   			value: jsii.Number(123),
//   		},
//   		maxEjectionPercent: jsii.Number(123),
//   		maxServerErrors: jsii.Number(123),
//   	},
//   	timeout: &listenerTimeoutProperty{
//   		grpc: &grpcTimeoutProperty{
//   			idle: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//   			perRequest: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//   		},
//   		http: &httpTimeoutProperty{
//   			idle: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//   			perRequest: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//   		},
//   		http2: &httpTimeoutProperty{
//   			idle: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//   			perRequest: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//   		},
//   		tcp: &tcpTimeoutProperty{
//   			idle: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	tls: &listenerTlsProperty{
//   		certificate: &listenerTlsCertificateProperty{
//   			acm: &listenerTlsAcmCertificateProperty{
//   				certificateArn: jsii.String("certificateArn"),
//   			},
//   			file: &listenerTlsFileCertificateProperty{
//   				certificateChain: jsii.String("certificateChain"),
//   				privateKey: jsii.String("privateKey"),
//   			},
//   			sds: &listenerTlsSdsCertificateProperty{
//   				secretName: jsii.String("secretName"),
//   			},
//   		},
//   		mode: jsii.String("mode"),
//
//   		// the properties below are optional
//   		validation: &listenerTlsValidationContextProperty{
//   			trust: &listenerTlsValidationContextTrustProperty{
//   				file: &tlsValidationContextFileTrustProperty{
//   					certificateChain: jsii.String("certificateChain"),
//   				},
//   				sds: &tlsValidationContextSdsTrustProperty{
//   					secretName: jsii.String("secretName"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   				match: &subjectAlternativeNameMatchersProperty{
//   					exact: []*string{
//   						jsii.String("exact"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnVirtualNode_ListenerProperty struct {
	// The port mapping information for the listener.
	PortMapping interface{} `field:"required" json:"portMapping" yaml:"portMapping"`
	// The connection pool information for the listener.
	ConnectionPool interface{} `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The outlier detection information for the listener.
	OutlierDetection interface{} `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// An object that represents timeouts for different protocols.
	Timeout interface{} `field:"optional" json:"timeout" yaml:"timeout"`
	// A reference to an object that represents the Transport Layer Security (TLS) properties for a listener.
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

