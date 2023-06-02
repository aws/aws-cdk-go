package awsappmesh


// An object that represents a listener for a virtual node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerProperty := &ListenerProperty{
//   	PortMapping: &PortMappingProperty{
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   	},
//
//   	// the properties below are optional
//   	ConnectionPool: &VirtualNodeConnectionPoolProperty{
//   		Grpc: &VirtualNodeGrpcConnectionPoolProperty{
//   			MaxRequests: jsii.Number(123),
//   		},
//   		Http: &VirtualNodeHttpConnectionPoolProperty{
//   			MaxConnections: jsii.Number(123),
//
//   			// the properties below are optional
//   			MaxPendingRequests: jsii.Number(123),
//   		},
//   		Http2: &VirtualNodeHttp2ConnectionPoolProperty{
//   			MaxRequests: jsii.Number(123),
//   		},
//   		Tcp: &VirtualNodeTcpConnectionPoolProperty{
//   			MaxConnections: jsii.Number(123),
//   		},
//   	},
//   	HealthCheck: &HealthCheckProperty{
//   		HealthyThreshold: jsii.Number(123),
//   		IntervalMillis: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		TimeoutMillis: jsii.Number(123),
//   		UnhealthyThreshold: jsii.Number(123),
//
//   		// the properties below are optional
//   		Path: jsii.String("path"),
//   		Port: jsii.Number(123),
//   	},
//   	OutlierDetection: &OutlierDetectionProperty{
//   		BaseEjectionDuration: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   		Interval: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   		MaxEjectionPercent: jsii.Number(123),
//   		MaxServerErrors: jsii.Number(123),
//   	},
//   	Timeout: &ListenerTimeoutProperty{
//   		Grpc: &GrpcTimeoutProperty{
//   			Idle: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			PerRequest: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   		Http: &HttpTimeoutProperty{
//   			Idle: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			PerRequest: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   		Http2: &HttpTimeoutProperty{
//   			Idle: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			PerRequest: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   		Tcp: &TcpTimeoutProperty{
//   			Idle: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Tls: &ListenerTlsProperty{
//   		Certificate: &ListenerTlsCertificateProperty{
//   			Acm: &ListenerTlsAcmCertificateProperty{
//   				CertificateArn: jsii.String("certificateArn"),
//   			},
//   			File: &ListenerTlsFileCertificateProperty{
//   				CertificateChain: jsii.String("certificateChain"),
//   				PrivateKey: jsii.String("privateKey"),
//   			},
//   			Sds: &ListenerTlsSdsCertificateProperty{
//   				SecretName: jsii.String("secretName"),
//   			},
//   		},
//   		Mode: jsii.String("mode"),
//
//   		// the properties below are optional
//   		Validation: &ListenerTlsValidationContextProperty{
//   			Trust: &ListenerTlsValidationContextTrustProperty{
//   				File: &TlsValidationContextFileTrustProperty{
//   					CertificateChain: jsii.String("certificateChain"),
//   				},
//   				Sds: &TlsValidationContextSdsTrustProperty{
//   					SecretName: jsii.String("secretName"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   				Match: &SubjectAlternativeNameMatchersProperty{
//   					Exact: []*string{
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

