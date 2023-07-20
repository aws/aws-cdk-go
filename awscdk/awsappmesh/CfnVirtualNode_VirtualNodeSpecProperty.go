package awsappmesh


// An object that represents the specification of a virtual node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualNodeSpecProperty := &VirtualNodeSpecProperty{
//   	BackendDefaults: &BackendDefaultsProperty{
//   		ClientPolicy: &ClientPolicyProperty{
//   			Tls: &ClientPolicyTlsProperty{
//   				Validation: &TlsValidationContextProperty{
//   					Trust: &TlsValidationContextTrustProperty{
//   						Acm: &TlsValidationContextAcmTrustProperty{
//   							CertificateAuthorityArns: []*string{
//   								jsii.String("certificateAuthorityArns"),
//   							},
//   						},
//   						File: &TlsValidationContextFileTrustProperty{
//   							CertificateChain: jsii.String("certificateChain"),
//   						},
//   						Sds: &TlsValidationContextSdsTrustProperty{
//   							SecretName: jsii.String("secretName"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   						Match: &SubjectAlternativeNameMatchersProperty{
//   							Exact: []*string{
//   								jsii.String("exact"),
//   							},
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				Certificate: &ClientTlsCertificateProperty{
//   					File: &ListenerTlsFileCertificateProperty{
//   						CertificateChain: jsii.String("certificateChain"),
//   						PrivateKey: jsii.String("privateKey"),
//   					},
//   					Sds: &ListenerTlsSdsCertificateProperty{
//   						SecretName: jsii.String("secretName"),
//   					},
//   				},
//   				Enforce: jsii.Boolean(false),
//   				Ports: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	Backends: []interface{}{
//   		&BackendProperty{
//   			VirtualService: &VirtualServiceBackendProperty{
//   				VirtualServiceName: jsii.String("virtualServiceName"),
//
//   				// the properties below are optional
//   				ClientPolicy: &ClientPolicyProperty{
//   					Tls: &ClientPolicyTlsProperty{
//   						Validation: &TlsValidationContextProperty{
//   							Trust: &TlsValidationContextTrustProperty{
//   								Acm: &TlsValidationContextAcmTrustProperty{
//   									CertificateAuthorityArns: []*string{
//   										jsii.String("certificateAuthorityArns"),
//   									},
//   								},
//   								File: &TlsValidationContextFileTrustProperty{
//   									CertificateChain: jsii.String("certificateChain"),
//   								},
//   								Sds: &TlsValidationContextSdsTrustProperty{
//   									SecretName: jsii.String("secretName"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   								Match: &SubjectAlternativeNameMatchersProperty{
//   									Exact: []*string{
//   										jsii.String("exact"),
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						Certificate: &ClientTlsCertificateProperty{
//   							File: &ListenerTlsFileCertificateProperty{
//   								CertificateChain: jsii.String("certificateChain"),
//   								PrivateKey: jsii.String("privateKey"),
//   							},
//   							Sds: &ListenerTlsSdsCertificateProperty{
//   								SecretName: jsii.String("secretName"),
//   							},
//   						},
//   						Enforce: jsii.Boolean(false),
//   						Ports: []interface{}{
//   							jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Listeners: []interface{}{
//   		&ListenerProperty{
//   			PortMapping: &PortMappingProperty{
//   				Port: jsii.Number(123),
//   				Protocol: jsii.String("protocol"),
//   			},
//
//   			// the properties below are optional
//   			ConnectionPool: &VirtualNodeConnectionPoolProperty{
//   				Grpc: &VirtualNodeGrpcConnectionPoolProperty{
//   					MaxRequests: jsii.Number(123),
//   				},
//   				Http: &VirtualNodeHttpConnectionPoolProperty{
//   					MaxConnections: jsii.Number(123),
//
//   					// the properties below are optional
//   					MaxPendingRequests: jsii.Number(123),
//   				},
//   				Http2: &VirtualNodeHttp2ConnectionPoolProperty{
//   					MaxRequests: jsii.Number(123),
//   				},
//   				Tcp: &VirtualNodeTcpConnectionPoolProperty{
//   					MaxConnections: jsii.Number(123),
//   				},
//   			},
//   			HealthCheck: &HealthCheckProperty{
//   				HealthyThreshold: jsii.Number(123),
//   				IntervalMillis: jsii.Number(123),
//   				Protocol: jsii.String("protocol"),
//   				TimeoutMillis: jsii.Number(123),
//   				UnhealthyThreshold: jsii.Number(123),
//
//   				// the properties below are optional
//   				Path: jsii.String("path"),
//   				Port: jsii.Number(123),
//   			},
//   			OutlierDetection: &OutlierDetectionProperty{
//   				BaseEjectionDuration: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				Interval: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				MaxEjectionPercent: jsii.Number(123),
//   				MaxServerErrors: jsii.Number(123),
//   			},
//   			Timeout: &ListenerTimeoutProperty{
//   				Grpc: &GrpcTimeoutProperty{
//   					Idle: &DurationProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.Number(123),
//   					},
//   					PerRequest: &DurationProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.Number(123),
//   					},
//   				},
//   				Http: &HttpTimeoutProperty{
//   					Idle: &DurationProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.Number(123),
//   					},
//   					PerRequest: &DurationProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.Number(123),
//   					},
//   				},
//   				Http2: &HttpTimeoutProperty{
//   					Idle: &DurationProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.Number(123),
//   					},
//   					PerRequest: &DurationProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.Number(123),
//   					},
//   				},
//   				Tcp: &TcpTimeoutProperty{
//   					Idle: &DurationProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Tls: &ListenerTlsProperty{
//   				Certificate: &ListenerTlsCertificateProperty{
//   					Acm: &ListenerTlsAcmCertificateProperty{
//   						CertificateArn: jsii.String("certificateArn"),
//   					},
//   					File: &ListenerTlsFileCertificateProperty{
//   						CertificateChain: jsii.String("certificateChain"),
//   						PrivateKey: jsii.String("privateKey"),
//   					},
//   					Sds: &ListenerTlsSdsCertificateProperty{
//   						SecretName: jsii.String("secretName"),
//   					},
//   				},
//   				Mode: jsii.String("mode"),
//
//   				// the properties below are optional
//   				Validation: &ListenerTlsValidationContextProperty{
//   					Trust: &ListenerTlsValidationContextTrustProperty{
//   						File: &TlsValidationContextFileTrustProperty{
//   							CertificateChain: jsii.String("certificateChain"),
//   						},
//   						Sds: &TlsValidationContextSdsTrustProperty{
//   							SecretName: jsii.String("secretName"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   						Match: &SubjectAlternativeNameMatchersProperty{
//   							Exact: []*string{
//   								jsii.String("exact"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Logging: &LoggingProperty{
//   		AccessLog: &AccessLogProperty{
//   			File: &FileAccessLogProperty{
//   				Path: jsii.String("path"),
//
//   				// the properties below are optional
//   				Format: &LoggingFormatProperty{
//   					Json: []interface{}{
//   						&JsonFormatRefProperty{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Text: jsii.String("text"),
//   				},
//   			},
//   		},
//   	},
//   	ServiceDiscovery: &ServiceDiscoveryProperty{
//   		AwsCloudMap: &AwsCloudMapServiceDiscoveryProperty{
//   			NamespaceName: jsii.String("namespaceName"),
//   			ServiceName: jsii.String("serviceName"),
//
//   			// the properties below are optional
//   			Attributes: []interface{}{
//   				&AwsCloudMapInstanceAttributeProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			IpPreference: jsii.String("ipPreference"),
//   		},
//   		Dns: &DnsServiceDiscoveryProperty{
//   			Hostname: jsii.String("hostname"),
//
//   			// the properties below are optional
//   			IpPreference: jsii.String("ipPreference"),
//   			ResponseType: jsii.String("responseType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualnodespec.html
//
type CfnVirtualNode_VirtualNodeSpecProperty struct {
	// A reference to an object that represents the defaults for backends.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualnodespec.html#cfn-appmesh-virtualnode-virtualnodespec-backenddefaults
	//
	BackendDefaults interface{} `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// The backends that the virtual node is expected to send outbound traffic to.
	//
	// > App Mesh doesn't validate the existence of those virtual services specified in backends. This is to prevent a cyclic dependency between virtual nodes and virtual services creation. Make sure the virtual service name is correct. The virtual service can be created afterwards if it doesn't already exist.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualnodespec.html#cfn-appmesh-virtualnode-virtualnodespec-backends
	//
	Backends interface{} `field:"optional" json:"backends" yaml:"backends"`
	// The listener that the virtual node is expected to receive inbound traffic from.
	//
	// You can specify one listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualnodespec.html#cfn-appmesh-virtualnode-virtualnodespec-listeners
	//
	Listeners interface{} `field:"optional" json:"listeners" yaml:"listeners"`
	// The inbound and outbound access logging information for the virtual node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualnodespec.html#cfn-appmesh-virtualnode-virtualnodespec-logging
	//
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// The service discovery information for the virtual node.
	//
	// If your virtual node does not expect ingress traffic, you can omit this parameter. If you specify a `listener` , then you must specify service discovery information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualnodespec.html#cfn-appmesh-virtualnode-virtualnodespec-servicediscovery
	//
	ServiceDiscovery interface{} `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
}

