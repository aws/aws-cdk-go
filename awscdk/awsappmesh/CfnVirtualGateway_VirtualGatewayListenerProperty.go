package awsappmesh


// An object that represents a listener for a virtual gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayListenerProperty := &VirtualGatewayListenerProperty{
//   	PortMapping: &VirtualGatewayPortMappingProperty{
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   	},
//
//   	// the properties below are optional
//   	ConnectionPool: &VirtualGatewayConnectionPoolProperty{
//   		Grpc: &VirtualGatewayGrpcConnectionPoolProperty{
//   			MaxRequests: jsii.Number(123),
//   		},
//   		Http: &VirtualGatewayHttpConnectionPoolProperty{
//   			MaxConnections: jsii.Number(123),
//
//   			// the properties below are optional
//   			MaxPendingRequests: jsii.Number(123),
//   		},
//   		Http2: &VirtualGatewayHttp2ConnectionPoolProperty{
//   			MaxRequests: jsii.Number(123),
//   		},
//   	},
//   	HealthCheck: &VirtualGatewayHealthCheckPolicyProperty{
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
//   	Tls: &VirtualGatewayListenerTlsProperty{
//   		Certificate: &VirtualGatewayListenerTlsCertificateProperty{
//   			Acm: &VirtualGatewayListenerTlsAcmCertificateProperty{
//   				CertificateArn: jsii.String("certificateArn"),
//   			},
//   			File: &VirtualGatewayListenerTlsFileCertificateProperty{
//   				CertificateChain: jsii.String("certificateChain"),
//   				PrivateKey: jsii.String("privateKey"),
//   			},
//   			Sds: &VirtualGatewayListenerTlsSdsCertificateProperty{
//   				SecretName: jsii.String("secretName"),
//   			},
//   		},
//   		Mode: jsii.String("mode"),
//
//   		// the properties below are optional
//   		Validation: &VirtualGatewayListenerTlsValidationContextProperty{
//   			Trust: &VirtualGatewayListenerTlsValidationContextTrustProperty{
//   				File: &VirtualGatewayTlsValidationContextFileTrustProperty{
//   					CertificateChain: jsii.String("certificateChain"),
//   				},
//   				Sds: &VirtualGatewayTlsValidationContextSdsTrustProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistener.html
//
type CfnVirtualGateway_VirtualGatewayListenerProperty struct {
	// The port mapping information for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistener.html#cfn-appmesh-virtualgateway-virtualgatewaylistener-portmapping
	//
	PortMapping interface{} `field:"required" json:"portMapping" yaml:"portMapping"`
	// The connection pool information for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistener.html#cfn-appmesh-virtualgateway-virtualgatewaylistener-connectionpool
	//
	ConnectionPool interface{} `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistener.html#cfn-appmesh-virtualgateway-virtualgatewaylistener-healthcheck
	//
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// A reference to an object that represents the Transport Layer Security (TLS) properties for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistener.html#cfn-appmesh-virtualgateway-virtualgatewaylistener-tls
	//
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

