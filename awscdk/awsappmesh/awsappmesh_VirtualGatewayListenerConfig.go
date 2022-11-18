package awsappmesh


// Properties for a VirtualGateway listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayListenerConfig := &virtualGatewayListenerConfig{
//   	listener: &virtualGatewayListenerProperty{
//   		portMapping: &virtualGatewayPortMappingProperty{
//   			port: jsii.Number(123),
//   			protocol: jsii.String("protocol"),
//   		},
//
//   		// the properties below are optional
//   		connectionPool: &virtualGatewayConnectionPoolProperty{
//   			grpc: &virtualGatewayGrpcConnectionPoolProperty{
//   				maxRequests: jsii.Number(123),
//   			},
//   			http: &virtualGatewayHttpConnectionPoolProperty{
//   				maxConnections: jsii.Number(123),
//
//   				// the properties below are optional
//   				maxPendingRequests: jsii.Number(123),
//   			},
//   			http2: &virtualGatewayHttp2ConnectionPoolProperty{
//   				maxRequests: jsii.Number(123),
//   			},
//   		},
//   		healthCheck: &virtualGatewayHealthCheckPolicyProperty{
//   			healthyThreshold: jsii.Number(123),
//   			intervalMillis: jsii.Number(123),
//   			protocol: jsii.String("protocol"),
//   			timeoutMillis: jsii.Number(123),
//   			unhealthyThreshold: jsii.Number(123),
//
//   			// the properties below are optional
//   			path: jsii.String("path"),
//   			port: jsii.Number(123),
//   		},
//   		tls: &virtualGatewayListenerTlsProperty{
//   			certificate: &virtualGatewayListenerTlsCertificateProperty{
//   				acm: &virtualGatewayListenerTlsAcmCertificateProperty{
//   					certificateArn: jsii.String("certificateArn"),
//   				},
//   				file: &virtualGatewayListenerTlsFileCertificateProperty{
//   					certificateChain: jsii.String("certificateChain"),
//   					privateKey: jsii.String("privateKey"),
//   				},
//   				sds: &virtualGatewayListenerTlsSdsCertificateProperty{
//   					secretName: jsii.String("secretName"),
//   				},
//   			},
//   			mode: jsii.String("mode"),
//
//   			// the properties below are optional
//   			validation: &virtualGatewayListenerTlsValidationContextProperty{
//   				trust: &virtualGatewayListenerTlsValidationContextTrustProperty{
//   					file: &virtualGatewayTlsValidationContextFileTrustProperty{
//   						certificateChain: jsii.String("certificateChain"),
//   					},
//   					sds: &virtualGatewayTlsValidationContextSdsTrustProperty{
//   						secretName: jsii.String("secretName"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   					match: &subjectAlternativeNameMatchersProperty{
//   						exact: []*string{
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

