package awsappmesh


// An object that represents the specification of a virtual node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualNodeSpecProperty := &virtualNodeSpecProperty{
//   	backendDefaults: &backendDefaultsProperty{
//   		clientPolicy: &clientPolicyProperty{
//   			tls: &clientPolicyTlsProperty{
//   				validation: &tlsValidationContextProperty{
//   					trust: &tlsValidationContextTrustProperty{
//   						acm: &tlsValidationContextAcmTrustProperty{
//   							certificateAuthorityArns: []*string{
//   								jsii.String("certificateAuthorityArns"),
//   							},
//   						},
//   						file: &tlsValidationContextFileTrustProperty{
//   							certificateChain: jsii.String("certificateChain"),
//   						},
//   						sds: &tlsValidationContextSdsTrustProperty{
//   							secretName: jsii.String("secretName"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   						match: &subjectAlternativeNameMatchersProperty{
//   							exact: []*string{
//   								jsii.String("exact"),
//   							},
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				certificate: &clientTlsCertificateProperty{
//   					file: &listenerTlsFileCertificateProperty{
//   						certificateChain: jsii.String("certificateChain"),
//   						privateKey: jsii.String("privateKey"),
//   					},
//   					sds: &listenerTlsSdsCertificateProperty{
//   						secretName: jsii.String("secretName"),
//   					},
//   				},
//   				enforce: jsii.Boolean(false),
//   				ports: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	backends: []interface{}{
//   		&backendProperty{
//   			virtualService: &virtualServiceBackendProperty{
//   				virtualServiceName: jsii.String("virtualServiceName"),
//
//   				// the properties below are optional
//   				clientPolicy: &clientPolicyProperty{
//   					tls: &clientPolicyTlsProperty{
//   						validation: &tlsValidationContextProperty{
//   							trust: &tlsValidationContextTrustProperty{
//   								acm: &tlsValidationContextAcmTrustProperty{
//   									certificateAuthorityArns: []*string{
//   										jsii.String("certificateAuthorityArns"),
//   									},
//   								},
//   								file: &tlsValidationContextFileTrustProperty{
//   									certificateChain: jsii.String("certificateChain"),
//   								},
//   								sds: &tlsValidationContextSdsTrustProperty{
//   									secretName: jsii.String("secretName"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   								match: &subjectAlternativeNameMatchersProperty{
//   									exact: []*string{
//   										jsii.String("exact"),
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						certificate: &clientTlsCertificateProperty{
//   							file: &listenerTlsFileCertificateProperty{
//   								certificateChain: jsii.String("certificateChain"),
//   								privateKey: jsii.String("privateKey"),
//   							},
//   							sds: &listenerTlsSdsCertificateProperty{
//   								secretName: jsii.String("secretName"),
//   							},
//   						},
//   						enforce: jsii.Boolean(false),
//   						ports: []interface{}{
//   							jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	listeners: []interface{}{
//   		&listenerProperty{
//   			portMapping: &portMappingProperty{
//   				port: jsii.Number(123),
//   				protocol: jsii.String("protocol"),
//   			},
//
//   			// the properties below are optional
//   			connectionPool: &virtualNodeConnectionPoolProperty{
//   				grpc: &virtualNodeGrpcConnectionPoolProperty{
//   					maxRequests: jsii.Number(123),
//   				},
//   				http: &virtualNodeHttpConnectionPoolProperty{
//   					maxConnections: jsii.Number(123),
//
//   					// the properties below are optional
//   					maxPendingRequests: jsii.Number(123),
//   				},
//   				http2: &virtualNodeHttp2ConnectionPoolProperty{
//   					maxRequests: jsii.Number(123),
//   				},
//   				tcp: &virtualNodeTcpConnectionPoolProperty{
//   					maxConnections: jsii.Number(123),
//   				},
//   			},
//   			healthCheck: &healthCheckProperty{
//   				healthyThreshold: jsii.Number(123),
//   				intervalMillis: jsii.Number(123),
//   				protocol: jsii.String("protocol"),
//   				timeoutMillis: jsii.Number(123),
//   				unhealthyThreshold: jsii.Number(123),
//
//   				// the properties below are optional
//   				path: jsii.String("path"),
//   				port: jsii.Number(123),
//   			},
//   			outlierDetection: &outlierDetectionProperty{
//   				baseEjectionDuration: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   				interval: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   				maxEjectionPercent: jsii.Number(123),
//   				maxServerErrors: jsii.Number(123),
//   			},
//   			timeout: &listenerTimeoutProperty{
//   				grpc: &grpcTimeoutProperty{
//   					idle: &durationProperty{
//   						unit: jsii.String("unit"),
//   						value: jsii.Number(123),
//   					},
//   					perRequest: &durationProperty{
//   						unit: jsii.String("unit"),
//   						value: jsii.Number(123),
//   					},
//   				},
//   				http: &httpTimeoutProperty{
//   					idle: &durationProperty{
//   						unit: jsii.String("unit"),
//   						value: jsii.Number(123),
//   					},
//   					perRequest: &durationProperty{
//   						unit: jsii.String("unit"),
//   						value: jsii.Number(123),
//   					},
//   				},
//   				http2: &httpTimeoutProperty{
//   					idle: &durationProperty{
//   						unit: jsii.String("unit"),
//   						value: jsii.Number(123),
//   					},
//   					perRequest: &durationProperty{
//   						unit: jsii.String("unit"),
//   						value: jsii.Number(123),
//   					},
//   				},
//   				tcp: &tcpTimeoutProperty{
//   					idle: &durationProperty{
//   						unit: jsii.String("unit"),
//   						value: jsii.Number(123),
//   					},
//   				},
//   			},
//   			tls: &listenerTlsProperty{
//   				certificate: &listenerTlsCertificateProperty{
//   					acm: &listenerTlsAcmCertificateProperty{
//   						certificateArn: jsii.String("certificateArn"),
//   					},
//   					file: &listenerTlsFileCertificateProperty{
//   						certificateChain: jsii.String("certificateChain"),
//   						privateKey: jsii.String("privateKey"),
//   					},
//   					sds: &listenerTlsSdsCertificateProperty{
//   						secretName: jsii.String("secretName"),
//   					},
//   				},
//   				mode: jsii.String("mode"),
//
//   				// the properties below are optional
//   				validation: &listenerTlsValidationContextProperty{
//   					trust: &listenerTlsValidationContextTrustProperty{
//   						file: &tlsValidationContextFileTrustProperty{
//   							certificateChain: jsii.String("certificateChain"),
//   						},
//   						sds: &tlsValidationContextSdsTrustProperty{
//   							secretName: jsii.String("secretName"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   						match: &subjectAlternativeNameMatchersProperty{
//   							exact: []*string{
//   								jsii.String("exact"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	logging: &loggingProperty{
//   		accessLog: &accessLogProperty{
//   			file: &fileAccessLogProperty{
//   				path: jsii.String("path"),
//
//   				// the properties below are optional
//   				format: &loggingFormatProperty{
//   					json: []interface{}{
//   						&jsonFormatRefProperty{
//   							key: jsii.String("key"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					text: jsii.String("text"),
//   				},
//   			},
//   		},
//   	},
//   	serviceDiscovery: &serviceDiscoveryProperty{
//   		awsCloudMap: &awsCloudMapServiceDiscoveryProperty{
//   			namespaceName: jsii.String("namespaceName"),
//   			serviceName: jsii.String("serviceName"),
//
//   			// the properties below are optional
//   			attributes: []interface{}{
//   				&awsCloudMapInstanceAttributeProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			ipPreference: jsii.String("ipPreference"),
//   		},
//   		dns: &dnsServiceDiscoveryProperty{
//   			hostname: jsii.String("hostname"),
//
//   			// the properties below are optional
//   			ipPreference: jsii.String("ipPreference"),
//   			responseType: jsii.String("responseType"),
//   		},
//   	},
//   }
//
type CfnVirtualNode_VirtualNodeSpecProperty struct {
	// A reference to an object that represents the defaults for backends.
	BackendDefaults interface{} `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// The backends that the virtual node is expected to send outbound traffic to.
	Backends interface{} `field:"optional" json:"backends" yaml:"backends"`
	// The listener that the virtual node is expected to receive inbound traffic from.
	//
	// You can specify one listener.
	Listeners interface{} `field:"optional" json:"listeners" yaml:"listeners"`
	// The inbound and outbound access logging information for the virtual node.
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// The service discovery information for the virtual node.
	//
	// If your virtual node does not expect ingress traffic, you can omit this parameter. If you specify a `listener` , then you must specify service discovery information.
	ServiceDiscovery interface{} `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
}

