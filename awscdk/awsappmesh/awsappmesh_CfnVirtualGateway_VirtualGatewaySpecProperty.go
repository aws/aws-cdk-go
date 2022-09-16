package awsappmesh


// An object that represents the specification of a service mesh resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewaySpecProperty := &virtualGatewaySpecProperty{
//   	listeners: []interface{}{
//   		&virtualGatewayListenerProperty{
//   			portMapping: &virtualGatewayPortMappingProperty{
//   				port: jsii.Number(123),
//   				protocol: jsii.String("protocol"),
//   			},
//
//   			// the properties below are optional
//   			connectionPool: &virtualGatewayConnectionPoolProperty{
//   				grpc: &virtualGatewayGrpcConnectionPoolProperty{
//   					maxRequests: jsii.Number(123),
//   				},
//   				http: &virtualGatewayHttpConnectionPoolProperty{
//   					maxConnections: jsii.Number(123),
//
//   					// the properties below are optional
//   					maxPendingRequests: jsii.Number(123),
//   				},
//   				http2: &virtualGatewayHttp2ConnectionPoolProperty{
//   					maxRequests: jsii.Number(123),
//   				},
//   			},
//   			healthCheck: &virtualGatewayHealthCheckPolicyProperty{
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
//   			tls: &virtualGatewayListenerTlsProperty{
//   				certificate: &virtualGatewayListenerTlsCertificateProperty{
//   					acm: &virtualGatewayListenerTlsAcmCertificateProperty{
//   						certificateArn: jsii.String("certificateArn"),
//   					},
//   					file: &virtualGatewayListenerTlsFileCertificateProperty{
//   						certificateChain: jsii.String("certificateChain"),
//   						privateKey: jsii.String("privateKey"),
//   					},
//   					sds: &virtualGatewayListenerTlsSdsCertificateProperty{
//   						secretName: jsii.String("secretName"),
//   					},
//   				},
//   				mode: jsii.String("mode"),
//
//   				// the properties below are optional
//   				validation: &virtualGatewayListenerTlsValidationContextProperty{
//   					trust: &virtualGatewayListenerTlsValidationContextTrustProperty{
//   						file: &virtualGatewayTlsValidationContextFileTrustProperty{
//   							certificateChain: jsii.String("certificateChain"),
//   						},
//   						sds: &virtualGatewayTlsValidationContextSdsTrustProperty{
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
//
//   	// the properties below are optional
//   	backendDefaults: &virtualGatewayBackendDefaultsProperty{
//   		clientPolicy: &virtualGatewayClientPolicyProperty{
//   			tls: &virtualGatewayClientPolicyTlsProperty{
//   				validation: &virtualGatewayTlsValidationContextProperty{
//   					trust: &virtualGatewayTlsValidationContextTrustProperty{
//   						acm: &virtualGatewayTlsValidationContextAcmTrustProperty{
//   							certificateAuthorityArns: []*string{
//   								jsii.String("certificateAuthorityArns"),
//   							},
//   						},
//   						file: &virtualGatewayTlsValidationContextFileTrustProperty{
//   							certificateChain: jsii.String("certificateChain"),
//   						},
//   						sds: &virtualGatewayTlsValidationContextSdsTrustProperty{
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
//   				certificate: &virtualGatewayClientTlsCertificateProperty{
//   					file: &virtualGatewayListenerTlsFileCertificateProperty{
//   						certificateChain: jsii.String("certificateChain"),
//   						privateKey: jsii.String("privateKey"),
//   					},
//   					sds: &virtualGatewayListenerTlsSdsCertificateProperty{
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
//   	logging: &virtualGatewayLoggingProperty{
//   		accessLog: &virtualGatewayAccessLogProperty{
//   			file: &virtualGatewayFileAccessLogProperty{
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
//   }
//
type CfnVirtualGateway_VirtualGatewaySpecProperty struct {
	// The listeners that the mesh endpoint is expected to receive inbound traffic from.
	//
	// You can specify one listener.
	Listeners interface{} `field:"required" json:"listeners" yaml:"listeners"`
	// A reference to an object that represents the defaults for backends.
	BackendDefaults interface{} `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// An object that represents logging information.
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
}

