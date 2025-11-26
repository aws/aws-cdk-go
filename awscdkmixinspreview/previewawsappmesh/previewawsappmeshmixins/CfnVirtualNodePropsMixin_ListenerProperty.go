package previewawsappmeshmixins


// An object that represents a listener for a virtual node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   listenerProperty := &ListenerProperty{
//   	ConnectionPool: &VirtualNodeConnectionPoolProperty{
//   		Grpc: &VirtualNodeGrpcConnectionPoolProperty{
//   			MaxRequests: jsii.Number(123),
//   		},
//   		Http: &VirtualNodeHttpConnectionPoolProperty{
//   			MaxConnections: jsii.Number(123),
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
//   		Path: jsii.String("path"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		TimeoutMillis: jsii.Number(123),
//   		UnhealthyThreshold: jsii.Number(123),
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
//   	PortMapping: &PortMappingProperty{
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
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
//   		Validation: &ListenerTlsValidationContextProperty{
//   			SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   				Match: &SubjectAlternativeNameMatchersProperty{
//   					Exact: []*string{
//   						jsii.String("exact"),
//   					},
//   				},
//   			},
//   			Trust: &ListenerTlsValidationContextTrustProperty{
//   				File: &TlsValidationContextFileTrustProperty{
//   					CertificateChain: jsii.String("certificateChain"),
//   				},
//   				Sds: &TlsValidationContextSdsTrustProperty{
//   					SecretName: jsii.String("secretName"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listener.html
//
type CfnVirtualNodePropsMixin_ListenerProperty struct {
	// The connection pool information for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listener.html#cfn-appmesh-virtualnode-listener-connectionpool
	//
	ConnectionPool interface{} `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listener.html#cfn-appmesh-virtualnode-listener-healthcheck
	//
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The outlier detection information for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listener.html#cfn-appmesh-virtualnode-listener-outlierdetection
	//
	OutlierDetection interface{} `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// The port mapping information for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listener.html#cfn-appmesh-virtualnode-listener-portmapping
	//
	PortMapping interface{} `field:"optional" json:"portMapping" yaml:"portMapping"`
	// An object that represents timeouts for different protocols.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listener.html#cfn-appmesh-virtualnode-listener-timeout
	//
	Timeout interface{} `field:"optional" json:"timeout" yaml:"timeout"`
	// A reference to an object that represents the Transport Layer Security (TLS) properties for a listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listener.html#cfn-appmesh-virtualnode-listener-tls
	//
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

