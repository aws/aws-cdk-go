package previewawsappmeshmixins


// An object that represents the specification of a virtual node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualNodeSpecProperty := &VirtualNodeSpecProperty{
//   	BackendDefaults: &BackendDefaultsProperty{
//   		ClientPolicy: &ClientPolicyProperty{
//   			Tls: &ClientPolicyTlsProperty{
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
//   				Validation: &TlsValidationContextProperty{
//   					SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   						Match: &SubjectAlternativeNameMatchersProperty{
//   							Exact: []*string{
//   								jsii.String("exact"),
//   							},
//   						},
//   					},
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
//   				},
//   			},
//   		},
//   	},
//   	Backends: []interface{}{
//   		&BackendProperty{
//   			VirtualService: &VirtualServiceBackendProperty{
//   				ClientPolicy: &ClientPolicyProperty{
//   					Tls: &ClientPolicyTlsProperty{
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
//   						Validation: &TlsValidationContextProperty{
//   							SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   								Match: &SubjectAlternativeNameMatchersProperty{
//   									Exact: []*string{
//   										jsii.String("exact"),
//   									},
//   								},
//   							},
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
//   						},
//   					},
//   				},
//   				VirtualServiceName: jsii.String("virtualServiceName"),
//   			},
//   		},
//   	},
//   	Listeners: []interface{}{
//   		&ListenerProperty{
//   			ConnectionPool: &VirtualNodeConnectionPoolProperty{
//   				Grpc: &VirtualNodeGrpcConnectionPoolProperty{
//   					MaxRequests: jsii.Number(123),
//   				},
//   				Http: &VirtualNodeHttpConnectionPoolProperty{
//   					MaxConnections: jsii.Number(123),
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
//   				Path: jsii.String("path"),
//   				Port: jsii.Number(123),
//   				Protocol: jsii.String("protocol"),
//   				TimeoutMillis: jsii.Number(123),
//   				UnhealthyThreshold: jsii.Number(123),
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
//   			PortMapping: &PortMappingProperty{
//   				Port: jsii.Number(123),
//   				Protocol: jsii.String("protocol"),
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
//   				Validation: &ListenerTlsValidationContextProperty{
//   					SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   						Match: &SubjectAlternativeNameMatchersProperty{
//   							Exact: []*string{
//   								jsii.String("exact"),
//   							},
//   						},
//   					},
//   					Trust: &ListenerTlsValidationContextTrustProperty{
//   						File: &TlsValidationContextFileTrustProperty{
//   							CertificateChain: jsii.String("certificateChain"),
//   						},
//   						Sds: &TlsValidationContextSdsTrustProperty{
//   							SecretName: jsii.String("secretName"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Logging: &LoggingProperty{
//   		AccessLog: &AccessLogProperty{
//   			File: &FileAccessLogProperty{
//   				Format: &LoggingFormatProperty{
//   					Json: []interface{}{
//   						&JsonFormatRefProperty{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Text: jsii.String("text"),
//   				},
//   				Path: jsii.String("path"),
//   			},
//   		},
//   	},
//   	ServiceDiscovery: &ServiceDiscoveryProperty{
//   		AwsCloudMap: &AwsCloudMapServiceDiscoveryProperty{
//   			Attributes: []interface{}{
//   				&AwsCloudMapInstanceAttributeProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			IpPreference: jsii.String("ipPreference"),
//   			NamespaceName: jsii.String("namespaceName"),
//   			ServiceName: jsii.String("serviceName"),
//   		},
//   		Dns: &DnsServiceDiscoveryProperty{
//   			Hostname: jsii.String("hostname"),
//   			IpPreference: jsii.String("ipPreference"),
//   			ResponseType: jsii.String("responseType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualnodespec.html
//
type CfnVirtualNodePropsMixin_VirtualNodeSpecProperty struct {
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

