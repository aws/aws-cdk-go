package awsappmesh


// An object that represents the specification of a service mesh resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewaySpecProperty := &VirtualGatewaySpecProperty{
//   	Listeners: []interface{}{
//   		&VirtualGatewayListenerProperty{
//   			PortMapping: &VirtualGatewayPortMappingProperty{
//   				Port: jsii.Number(123),
//   				Protocol: jsii.String("protocol"),
//   			},
//
//   			// the properties below are optional
//   			ConnectionPool: &VirtualGatewayConnectionPoolProperty{
//   				Grpc: &VirtualGatewayGrpcConnectionPoolProperty{
//   					MaxRequests: jsii.Number(123),
//   				},
//   				Http: &VirtualGatewayHttpConnectionPoolProperty{
//   					MaxConnections: jsii.Number(123),
//
//   					// the properties below are optional
//   					MaxPendingRequests: jsii.Number(123),
//   				},
//   				Http2: &VirtualGatewayHttp2ConnectionPoolProperty{
//   					MaxRequests: jsii.Number(123),
//   				},
//   			},
//   			HealthCheck: &VirtualGatewayHealthCheckPolicyProperty{
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
//   			Tls: &VirtualGatewayListenerTlsProperty{
//   				Certificate: &VirtualGatewayListenerTlsCertificateProperty{
//   					Acm: &VirtualGatewayListenerTlsAcmCertificateProperty{
//   						CertificateArn: jsii.String("certificateArn"),
//   					},
//   					File: &VirtualGatewayListenerTlsFileCertificateProperty{
//   						CertificateChain: jsii.String("certificateChain"),
//   						PrivateKey: jsii.String("privateKey"),
//   					},
//   					Sds: &VirtualGatewayListenerTlsSdsCertificateProperty{
//   						SecretName: jsii.String("secretName"),
//   					},
//   				},
//   				Mode: jsii.String("mode"),
//
//   				// the properties below are optional
//   				Validation: &VirtualGatewayListenerTlsValidationContextProperty{
//   					Trust: &VirtualGatewayListenerTlsValidationContextTrustProperty{
//   						File: &VirtualGatewayTlsValidationContextFileTrustProperty{
//   							CertificateChain: jsii.String("certificateChain"),
//   						},
//   						Sds: &VirtualGatewayTlsValidationContextSdsTrustProperty{
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
//
//   	// the properties below are optional
//   	BackendDefaults: &VirtualGatewayBackendDefaultsProperty{
//   		ClientPolicy: &VirtualGatewayClientPolicyProperty{
//   			Tls: &VirtualGatewayClientPolicyTlsProperty{
//   				Validation: &VirtualGatewayTlsValidationContextProperty{
//   					Trust: &VirtualGatewayTlsValidationContextTrustProperty{
//   						Acm: &VirtualGatewayTlsValidationContextAcmTrustProperty{
//   							CertificateAuthorityArns: []*string{
//   								jsii.String("certificateAuthorityArns"),
//   							},
//   						},
//   						File: &VirtualGatewayTlsValidationContextFileTrustProperty{
//   							CertificateChain: jsii.String("certificateChain"),
//   						},
//   						Sds: &VirtualGatewayTlsValidationContextSdsTrustProperty{
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
//   				Certificate: &VirtualGatewayClientTlsCertificateProperty{
//   					File: &VirtualGatewayListenerTlsFileCertificateProperty{
//   						CertificateChain: jsii.String("certificateChain"),
//   						PrivateKey: jsii.String("privateKey"),
//   					},
//   					Sds: &VirtualGatewayListenerTlsSdsCertificateProperty{
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
//   	Logging: &VirtualGatewayLoggingProperty{
//   		AccessLog: &VirtualGatewayAccessLogProperty{
//   			File: &VirtualGatewayFileAccessLogProperty{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayspec.html
//
type CfnVirtualGateway_VirtualGatewaySpecProperty struct {
	// The listeners that the mesh endpoint is expected to receive inbound traffic from.
	//
	// You can specify one listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayspec.html#cfn-appmesh-virtualgateway-virtualgatewayspec-listeners
	//
	Listeners interface{} `field:"required" json:"listeners" yaml:"listeners"`
	// A reference to an object that represents the defaults for backends.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayspec.html#cfn-appmesh-virtualgateway-virtualgatewayspec-backenddefaults
	//
	BackendDefaults interface{} `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// An object that represents logging information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayspec.html#cfn-appmesh-virtualgateway-virtualgatewayspec-logging
	//
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
}

