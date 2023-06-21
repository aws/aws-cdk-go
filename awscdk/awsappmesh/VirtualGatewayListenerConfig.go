package awsappmesh


// Properties for a VirtualGateway listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayListenerConfig := &VirtualGatewayListenerConfig{
//   	Listener: &VirtualGatewayListenerProperty{
//   		PortMapping: &VirtualGatewayPortMappingProperty{
//   			Port: jsii.Number(123),
//   			Protocol: jsii.String("protocol"),
//   		},
//
//   		// the properties below are optional
//   		ConnectionPool: &VirtualGatewayConnectionPoolProperty{
//   			Grpc: &VirtualGatewayGrpcConnectionPoolProperty{
//   				MaxRequests: jsii.Number(123),
//   			},
//   			Http: &VirtualGatewayHttpConnectionPoolProperty{
//   				MaxConnections: jsii.Number(123),
//
//   				// the properties below are optional
//   				MaxPendingRequests: jsii.Number(123),
//   			},
//   			Http2: &VirtualGatewayHttp2ConnectionPoolProperty{
//   				MaxRequests: jsii.Number(123),
//   			},
//   		},
//   		HealthCheck: &VirtualGatewayHealthCheckPolicyProperty{
//   			HealthyThreshold: jsii.Number(123),
//   			IntervalMillis: jsii.Number(123),
//   			Protocol: jsii.String("protocol"),
//   			TimeoutMillis: jsii.Number(123),
//   			UnhealthyThreshold: jsii.Number(123),
//
//   			// the properties below are optional
//   			Path: jsii.String("path"),
//   			Port: jsii.Number(123),
//   		},
//   		Tls: &VirtualGatewayListenerTlsProperty{
//   			Certificate: &VirtualGatewayListenerTlsCertificateProperty{
//   				Acm: &VirtualGatewayListenerTlsAcmCertificateProperty{
//   					CertificateArn: jsii.String("certificateArn"),
//   				},
//   				File: &VirtualGatewayListenerTlsFileCertificateProperty{
//   					CertificateChain: jsii.String("certificateChain"),
//   					PrivateKey: jsii.String("privateKey"),
//   				},
//   				Sds: &VirtualGatewayListenerTlsSdsCertificateProperty{
//   					SecretName: jsii.String("secretName"),
//   				},
//   			},
//   			Mode: jsii.String("mode"),
//
//   			// the properties below are optional
//   			Validation: &VirtualGatewayListenerTlsValidationContextProperty{
//   				Trust: &VirtualGatewayListenerTlsValidationContextTrustProperty{
//   					File: &VirtualGatewayTlsValidationContextFileTrustProperty{
//   						CertificateChain: jsii.String("certificateChain"),
//   					},
//   					Sds: &VirtualGatewayTlsValidationContextSdsTrustProperty{
//   						SecretName: jsii.String("secretName"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   					Match: &SubjectAlternativeNameMatchersProperty{
//   						Exact: []*string{
//   							jsii.String("exact"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type VirtualGatewayListenerConfig struct {
	// Single listener config for a VirtualGateway.
	Listener *CfnVirtualGateway_VirtualGatewayListenerProperty `field:"required" json:"listener" yaml:"listener"`
}

