package awsappmesh


// Properties for a VirtualNode listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualNodeListenerConfig := &virtualNodeListenerConfig{
//   	listener: &listenerProperty{
//   		portMapping: &portMappingProperty{
//   			port: jsii.Number(123),
//   			protocol: jsii.String("protocol"),
//   		},
//
//   		// the properties below are optional
//   		connectionPool: &virtualNodeConnectionPoolProperty{
//   			grpc: &virtualNodeGrpcConnectionPoolProperty{
//   				maxRequests: jsii.Number(123),
//   			},
//   			http: &virtualNodeHttpConnectionPoolProperty{
//   				maxConnections: jsii.Number(123),
//
//   				// the properties below are optional
//   				maxPendingRequests: jsii.Number(123),
//   			},
//   			http2: &virtualNodeHttp2ConnectionPoolProperty{
//   				maxRequests: jsii.Number(123),
//   			},
//   			tcp: &virtualNodeTcpConnectionPoolProperty{
//   				maxConnections: jsii.Number(123),
//   			},
//   		},
//   		healthCheck: &healthCheckProperty{
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
//   		outlierDetection: &outlierDetectionProperty{
//   			baseEjectionDuration: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//   			interval: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//   			maxEjectionPercent: jsii.Number(123),
//   			maxServerErrors: jsii.Number(123),
//   		},
//   		timeout: &listenerTimeoutProperty{
//   			grpc: &grpcTimeoutProperty{
//   				idle: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   				perRequest: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   			},
//   			http: &httpTimeoutProperty{
//   				idle: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   				perRequest: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   			},
//   			http2: &httpTimeoutProperty{
//   				idle: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   				perRequest: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   			},
//   			tcp: &tcpTimeoutProperty{
//   				idle: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   			},
//   		},
//   		tls: &listenerTlsProperty{
//   			certificate: &listenerTlsCertificateProperty{
//   				acm: &listenerTlsAcmCertificateProperty{
//   					certificateArn: jsii.String("certificateArn"),
//   				},
//   				file: &listenerTlsFileCertificateProperty{
//   					certificateChain: jsii.String("certificateChain"),
//   					privateKey: jsii.String("privateKey"),
//   				},
//   				sds: &listenerTlsSdsCertificateProperty{
//   					secretName: jsii.String("secretName"),
//   				},
//   			},
//   			mode: jsii.String("mode"),
//
//   			// the properties below are optional
//   			validation: &listenerTlsValidationContextProperty{
//   				trust: &listenerTlsValidationContextTrustProperty{
//   					file: &tlsValidationContextFileTrustProperty{
//   						certificateChain: jsii.String("certificateChain"),
//   					},
//   					sds: &tlsValidationContextSdsTrustProperty{
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
type VirtualNodeListenerConfig struct {
	// Single listener config for a VirtualNode.
	Listener *CfnVirtualNode_ListenerProperty `field:"required" json:"listener" yaml:"listener"`
}

