package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnVirtualGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVirtualGatewayProps := &CfnVirtualGatewayProps{
//   	MeshName: jsii.String("meshName"),
//   	Spec: &VirtualGatewaySpecProperty{
//   		Listeners: []interface{}{
//   			&VirtualGatewayListenerProperty{
//   				PortMapping: &VirtualGatewayPortMappingProperty{
//   					Port: jsii.Number(123),
//   					Protocol: jsii.String("protocol"),
//   				},
//
//   				// the properties below are optional
//   				ConnectionPool: &VirtualGatewayConnectionPoolProperty{
//   					Grpc: &VirtualGatewayGrpcConnectionPoolProperty{
//   						MaxRequests: jsii.Number(123),
//   					},
//   					Http: &VirtualGatewayHttpConnectionPoolProperty{
//   						MaxConnections: jsii.Number(123),
//
//   						// the properties below are optional
//   						MaxPendingRequests: jsii.Number(123),
//   					},
//   					Http2: &VirtualGatewayHttp2ConnectionPoolProperty{
//   						MaxRequests: jsii.Number(123),
//   					},
//   				},
//   				HealthCheck: &VirtualGatewayHealthCheckPolicyProperty{
//   					HealthyThreshold: jsii.Number(123),
//   					IntervalMillis: jsii.Number(123),
//   					Protocol: jsii.String("protocol"),
//   					TimeoutMillis: jsii.Number(123),
//   					UnhealthyThreshold: jsii.Number(123),
//
//   					// the properties below are optional
//   					Path: jsii.String("path"),
//   					Port: jsii.Number(123),
//   				},
//   				Tls: &VirtualGatewayListenerTlsProperty{
//   					Certificate: &VirtualGatewayListenerTlsCertificateProperty{
//   						Acm: &VirtualGatewayListenerTlsAcmCertificateProperty{
//   							CertificateArn: jsii.String("certificateArn"),
//   						},
//   						File: &VirtualGatewayListenerTlsFileCertificateProperty{
//   							CertificateChain: jsii.String("certificateChain"),
//   							PrivateKey: jsii.String("privateKey"),
//   						},
//   						Sds: &VirtualGatewayListenerTlsSdsCertificateProperty{
//   							SecretName: jsii.String("secretName"),
//   						},
//   					},
//   					Mode: jsii.String("mode"),
//
//   					// the properties below are optional
//   					Validation: &VirtualGatewayListenerTlsValidationContextProperty{
//   						Trust: &VirtualGatewayListenerTlsValidationContextTrustProperty{
//   							File: &VirtualGatewayTlsValidationContextFileTrustProperty{
//   								CertificateChain: jsii.String("certificateChain"),
//   							},
//   							Sds: &VirtualGatewayTlsValidationContextSdsTrustProperty{
//   								SecretName: jsii.String("secretName"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   							Match: &SubjectAlternativeNameMatchersProperty{
//   								Exact: []*string{
//   									jsii.String("exact"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		BackendDefaults: &VirtualGatewayBackendDefaultsProperty{
//   			ClientPolicy: &VirtualGatewayClientPolicyProperty{
//   				Tls: &VirtualGatewayClientPolicyTlsProperty{
//   					Validation: &VirtualGatewayTlsValidationContextProperty{
//   						Trust: &VirtualGatewayTlsValidationContextTrustProperty{
//   							Acm: &VirtualGatewayTlsValidationContextAcmTrustProperty{
//   								CertificateAuthorityArns: []*string{
//   									jsii.String("certificateAuthorityArns"),
//   								},
//   							},
//   							File: &VirtualGatewayTlsValidationContextFileTrustProperty{
//   								CertificateChain: jsii.String("certificateChain"),
//   							},
//   							Sds: &VirtualGatewayTlsValidationContextSdsTrustProperty{
//   								SecretName: jsii.String("secretName"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   							Match: &SubjectAlternativeNameMatchersProperty{
//   								Exact: []*string{
//   									jsii.String("exact"),
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					Certificate: &VirtualGatewayClientTlsCertificateProperty{
//   						File: &VirtualGatewayListenerTlsFileCertificateProperty{
//   							CertificateChain: jsii.String("certificateChain"),
//   							PrivateKey: jsii.String("privateKey"),
//   						},
//   						Sds: &VirtualGatewayListenerTlsSdsCertificateProperty{
//   							SecretName: jsii.String("secretName"),
//   						},
//   					},
//   					Enforce: jsii.Boolean(false),
//   					Ports: []interface{}{
//   						jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		Logging: &VirtualGatewayLoggingProperty{
//   			AccessLog: &VirtualGatewayAccessLogProperty{
//   				File: &VirtualGatewayFileAccessLogProperty{
//   					Path: jsii.String("path"),
//
//   					// the properties below are optional
//   					Format: &LoggingFormatProperty{
//   						Json: []interface{}{
//   							&JsonFormatRefProperty{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						Text: jsii.String("text"),
//   					},
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	MeshOwner: jsii.String("meshOwner"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VirtualGatewayName: jsii.String("virtualGatewayName"),
//   }
//
type CfnVirtualGatewayProps struct {
	// The name of the service mesh that the virtual gateway resides in.
	MeshName *string `field:"required" json:"meshName" yaml:"meshName"`
	// The specifications of the virtual gateway.
	Spec interface{} `field:"required" json:"spec" yaml:"spec"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then it's the ID of the account that shared the mesh with your account. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	MeshOwner *string `field:"optional" json:"meshOwner" yaml:"meshOwner"`
	// Optional metadata that you can apply to the virtual gateway to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the virtual gateway.
	VirtualGatewayName *string `field:"optional" json:"virtualGatewayName" yaml:"virtualGatewayName"`
}

