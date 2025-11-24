package mixinsawsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnVirtualGatewayPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVirtualGatewayMixinProps := &CfnVirtualGatewayMixinProps{
//   	MeshName: jsii.String("meshName"),
//   	MeshOwner: jsii.String("meshOwner"),
//   	Spec: &VirtualGatewaySpecProperty{
//   		BackendDefaults: &VirtualGatewayBackendDefaultsProperty{
//   			ClientPolicy: &VirtualGatewayClientPolicyProperty{
//   				Tls: &VirtualGatewayClientPolicyTlsProperty{
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
//   					Validation: &VirtualGatewayTlsValidationContextProperty{
//   						SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   							Match: &SubjectAlternativeNameMatchersProperty{
//   								Exact: []*string{
//   									jsii.String("exact"),
//   								},
//   							},
//   						},
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
//   					},
//   				},
//   			},
//   		},
//   		Listeners: []interface{}{
//   			&VirtualGatewayListenerProperty{
//   				ConnectionPool: &VirtualGatewayConnectionPoolProperty{
//   					Grpc: &VirtualGatewayGrpcConnectionPoolProperty{
//   						MaxRequests: jsii.Number(123),
//   					},
//   					Http: &VirtualGatewayHttpConnectionPoolProperty{
//   						MaxConnections: jsii.Number(123),
//   						MaxPendingRequests: jsii.Number(123),
//   					},
//   					Http2: &VirtualGatewayHttp2ConnectionPoolProperty{
//   						MaxRequests: jsii.Number(123),
//   					},
//   				},
//   				HealthCheck: &VirtualGatewayHealthCheckPolicyProperty{
//   					HealthyThreshold: jsii.Number(123),
//   					IntervalMillis: jsii.Number(123),
//   					Path: jsii.String("path"),
//   					Port: jsii.Number(123),
//   					Protocol: jsii.String("protocol"),
//   					TimeoutMillis: jsii.Number(123),
//   					UnhealthyThreshold: jsii.Number(123),
//   				},
//   				PortMapping: &VirtualGatewayPortMappingProperty{
//   					Port: jsii.Number(123),
//   					Protocol: jsii.String("protocol"),
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
//   					Validation: &VirtualGatewayListenerTlsValidationContextProperty{
//   						SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   							Match: &SubjectAlternativeNameMatchersProperty{
//   								Exact: []*string{
//   									jsii.String("exact"),
//   								},
//   							},
//   						},
//   						Trust: &VirtualGatewayListenerTlsValidationContextTrustProperty{
//   							File: &VirtualGatewayTlsValidationContextFileTrustProperty{
//   								CertificateChain: jsii.String("certificateChain"),
//   							},
//   							Sds: &VirtualGatewayTlsValidationContextSdsTrustProperty{
//   								SecretName: jsii.String("secretName"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		Logging: &VirtualGatewayLoggingProperty{
//   			AccessLog: &VirtualGatewayAccessLogProperty{
//   				File: &VirtualGatewayFileAccessLogProperty{
//   					Format: &LoggingFormatProperty{
//   						Json: []interface{}{
//   							&JsonFormatRefProperty{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						Text: jsii.String("text"),
//   					},
//   					Path: jsii.String("path"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VirtualGatewayName: jsii.String("virtualGatewayName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualgateway.html
//
type CfnVirtualGatewayMixinProps struct {
	// The name of the service mesh that the virtual gateway resides in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualgateway.html#cfn-appmesh-virtualgateway-meshname
	//
	MeshName *string `field:"optional" json:"meshName" yaml:"meshName"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then it's the ID of the account that shared the mesh with your account. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualgateway.html#cfn-appmesh-virtualgateway-meshowner
	//
	MeshOwner *string `field:"optional" json:"meshOwner" yaml:"meshOwner"`
	// The specifications of the virtual gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualgateway.html#cfn-appmesh-virtualgateway-spec
	//
	Spec interface{} `field:"optional" json:"spec" yaml:"spec"`
	// Optional metadata that you can apply to the virtual gateway to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualgateway.html#cfn-appmesh-virtualgateway-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the virtual gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualgateway.html#cfn-appmesh-virtualgateway-virtualgatewayname
	//
	VirtualGatewayName *string `field:"optional" json:"virtualGatewayName" yaml:"virtualGatewayName"`
}

