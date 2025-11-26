package previewawsappmeshmixins


// An object that represents a listener for a virtual gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualGatewayListenerProperty := &VirtualGatewayListenerProperty{
//   	ConnectionPool: &VirtualGatewayConnectionPoolProperty{
//   		Grpc: &VirtualGatewayGrpcConnectionPoolProperty{
//   			MaxRequests: jsii.Number(123),
//   		},
//   		Http: &VirtualGatewayHttpConnectionPoolProperty{
//   			MaxConnections: jsii.Number(123),
//   			MaxPendingRequests: jsii.Number(123),
//   		},
//   		Http2: &VirtualGatewayHttp2ConnectionPoolProperty{
//   			MaxRequests: jsii.Number(123),
//   		},
//   	},
//   	HealthCheck: &VirtualGatewayHealthCheckPolicyProperty{
//   		HealthyThreshold: jsii.Number(123),
//   		IntervalMillis: jsii.Number(123),
//   		Path: jsii.String("path"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		TimeoutMillis: jsii.Number(123),
//   		UnhealthyThreshold: jsii.Number(123),
//   	},
//   	PortMapping: &VirtualGatewayPortMappingProperty{
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
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
//   		Validation: &VirtualGatewayListenerTlsValidationContextProperty{
//   			SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   				Match: &SubjectAlternativeNameMatchersProperty{
//   					Exact: []*string{
//   						jsii.String("exact"),
//   					},
//   				},
//   			},
//   			Trust: &VirtualGatewayListenerTlsValidationContextTrustProperty{
//   				File: &VirtualGatewayTlsValidationContextFileTrustProperty{
//   					CertificateChain: jsii.String("certificateChain"),
//   				},
//   				Sds: &VirtualGatewayTlsValidationContextSdsTrustProperty{
//   					SecretName: jsii.String("secretName"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistener.html
//
type CfnVirtualGatewayPropsMixin_VirtualGatewayListenerProperty struct {
	// The connection pool information for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistener.html#cfn-appmesh-virtualgateway-virtualgatewaylistener-connectionpool
	//
	ConnectionPool interface{} `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistener.html#cfn-appmesh-virtualgateway-virtualgatewaylistener-healthcheck
	//
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The port mapping information for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistener.html#cfn-appmesh-virtualgateway-virtualgatewaylistener-portmapping
	//
	PortMapping interface{} `field:"optional" json:"portMapping" yaml:"portMapping"`
	// A reference to an object that represents the Transport Layer Security (TLS) properties for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistener.html#cfn-appmesh-virtualgateway-virtualgatewaylistener-tls
	//
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

