package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVirtualNode`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVirtualNodeProps := &CfnVirtualNodeProps{
//   	MeshName: jsii.String("meshName"),
//   	Spec: &VirtualNodeSpecProperty{
//   		BackendDefaults: &BackendDefaultsProperty{
//   			ClientPolicy: &ClientPolicyProperty{
//   				Tls: &ClientPolicyTlsProperty{
//   					Validation: &TlsValidationContextProperty{
//   						Trust: &TlsValidationContextTrustProperty{
//   							Acm: &TlsValidationContextAcmTrustProperty{
//   								CertificateAuthorityArns: []*string{
//   									jsii.String("certificateAuthorityArns"),
//   								},
//   							},
//   							File: &TlsValidationContextFileTrustProperty{
//   								CertificateChain: jsii.String("certificateChain"),
//   							},
//   							Sds: &TlsValidationContextSdsTrustProperty{
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
//   					Certificate: &ClientTlsCertificateProperty{
//   						File: &ListenerTlsFileCertificateProperty{
//   							CertificateChain: jsii.String("certificateChain"),
//   							PrivateKey: jsii.String("privateKey"),
//   						},
//   						Sds: &ListenerTlsSdsCertificateProperty{
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
//   		Backends: []interface{}{
//   			&BackendProperty{
//   				VirtualService: &VirtualServiceBackendProperty{
//   					VirtualServiceName: jsii.String("virtualServiceName"),
//
//   					// the properties below are optional
//   					ClientPolicy: &ClientPolicyProperty{
//   						Tls: &ClientPolicyTlsProperty{
//   							Validation: &TlsValidationContextProperty{
//   								Trust: &TlsValidationContextTrustProperty{
//   									Acm: &TlsValidationContextAcmTrustProperty{
//   										CertificateAuthorityArns: []*string{
//   											jsii.String("certificateAuthorityArns"),
//   										},
//   									},
//   									File: &TlsValidationContextFileTrustProperty{
//   										CertificateChain: jsii.String("certificateChain"),
//   									},
//   									Sds: &TlsValidationContextSdsTrustProperty{
//   										SecretName: jsii.String("secretName"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   									Match: &SubjectAlternativeNameMatchersProperty{
//   										Exact: []*string{
//   											jsii.String("exact"),
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							Certificate: &ClientTlsCertificateProperty{
//   								File: &ListenerTlsFileCertificateProperty{
//   									CertificateChain: jsii.String("certificateChain"),
//   									PrivateKey: jsii.String("privateKey"),
//   								},
//   								Sds: &ListenerTlsSdsCertificateProperty{
//   									SecretName: jsii.String("secretName"),
//   								},
//   							},
//   							Enforce: jsii.Boolean(false),
//   							Ports: []interface{}{
//   								jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		Listeners: []interface{}{
//   			&ListenerProperty{
//   				PortMapping: &PortMappingProperty{
//   					Port: jsii.Number(123),
//   					Protocol: jsii.String("protocol"),
//   				},
//
//   				// the properties below are optional
//   				ConnectionPool: &VirtualNodeConnectionPoolProperty{
//   					Grpc: &VirtualNodeGrpcConnectionPoolProperty{
//   						MaxRequests: jsii.Number(123),
//   					},
//   					Http: &VirtualNodeHttpConnectionPoolProperty{
//   						MaxConnections: jsii.Number(123),
//
//   						// the properties below are optional
//   						MaxPendingRequests: jsii.Number(123),
//   					},
//   					Http2: &VirtualNodeHttp2ConnectionPoolProperty{
//   						MaxRequests: jsii.Number(123),
//   					},
//   					Tcp: &VirtualNodeTcpConnectionPoolProperty{
//   						MaxConnections: jsii.Number(123),
//   					},
//   				},
//   				HealthCheck: &HealthCheckProperty{
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
//   				OutlierDetection: &OutlierDetectionProperty{
//   					BaseEjectionDuration: &DurationProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.Number(123),
//   					},
//   					Interval: &DurationProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.Number(123),
//   					},
//   					MaxEjectionPercent: jsii.Number(123),
//   					MaxServerErrors: jsii.Number(123),
//   				},
//   				Timeout: &ListenerTimeoutProperty{
//   					Grpc: &GrpcTimeoutProperty{
//   						Idle: &DurationProperty{
//   							Unit: jsii.String("unit"),
//   							Value: jsii.Number(123),
//   						},
//   						PerRequest: &DurationProperty{
//   							Unit: jsii.String("unit"),
//   							Value: jsii.Number(123),
//   						},
//   					},
//   					Http: &HttpTimeoutProperty{
//   						Idle: &DurationProperty{
//   							Unit: jsii.String("unit"),
//   							Value: jsii.Number(123),
//   						},
//   						PerRequest: &DurationProperty{
//   							Unit: jsii.String("unit"),
//   							Value: jsii.Number(123),
//   						},
//   					},
//   					Http2: &HttpTimeoutProperty{
//   						Idle: &DurationProperty{
//   							Unit: jsii.String("unit"),
//   							Value: jsii.Number(123),
//   						},
//   						PerRequest: &DurationProperty{
//   							Unit: jsii.String("unit"),
//   							Value: jsii.Number(123),
//   						},
//   					},
//   					Tcp: &TcpTimeoutProperty{
//   						Idle: &DurationProperty{
//   							Unit: jsii.String("unit"),
//   							Value: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Tls: &ListenerTlsProperty{
//   					Certificate: &ListenerTlsCertificateProperty{
//   						Acm: &ListenerTlsAcmCertificateProperty{
//   							CertificateArn: jsii.String("certificateArn"),
//   						},
//   						File: &ListenerTlsFileCertificateProperty{
//   							CertificateChain: jsii.String("certificateChain"),
//   							PrivateKey: jsii.String("privateKey"),
//   						},
//   						Sds: &ListenerTlsSdsCertificateProperty{
//   							SecretName: jsii.String("secretName"),
//   						},
//   					},
//   					Mode: jsii.String("mode"),
//
//   					// the properties below are optional
//   					Validation: &ListenerTlsValidationContextProperty{
//   						Trust: &ListenerTlsValidationContextTrustProperty{
//   							File: &TlsValidationContextFileTrustProperty{
//   								CertificateChain: jsii.String("certificateChain"),
//   							},
//   							Sds: &TlsValidationContextSdsTrustProperty{
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
//   		Logging: &LoggingProperty{
//   			AccessLog: &AccessLogProperty{
//   				File: &FileAccessLogProperty{
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
//   		ServiceDiscovery: &ServiceDiscoveryProperty{
//   			AwsCloudMap: &AwsCloudMapServiceDiscoveryProperty{
//   				NamespaceName: jsii.String("namespaceName"),
//   				ServiceName: jsii.String("serviceName"),
//
//   				// the properties below are optional
//   				Attributes: []interface{}{
//   					&AwsCloudMapInstanceAttributeProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				IpPreference: jsii.String("ipPreference"),
//   			},
//   			Dns: &DnsServiceDiscoveryProperty{
//   				Hostname: jsii.String("hostname"),
//
//   				// the properties below are optional
//   				IpPreference: jsii.String("ipPreference"),
//   				ResponseType: jsii.String("responseType"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	MeshOwner: jsii.String("meshOwner"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VirtualNodeName: jsii.String("virtualNodeName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualnode.html
//
type CfnVirtualNodeProps struct {
	// The name of the service mesh to create the virtual node in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualnode.html#cfn-appmesh-virtualnode-meshname
	//
	MeshName *string `field:"required" json:"meshName" yaml:"meshName"`
	// The virtual node specification to apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualnode.html#cfn-appmesh-virtualnode-spec
	//
	Spec interface{} `field:"required" json:"spec" yaml:"spec"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then the account that you specify must share the mesh with your account before you can create the resource in the service mesh. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualnode.html#cfn-appmesh-virtualnode-meshowner
	//
	MeshOwner *string `field:"optional" json:"meshOwner" yaml:"meshOwner"`
	// Optional metadata that you can apply to the virtual node to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualnode.html#cfn-appmesh-virtualnode-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name to use for the virtual node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualnode.html#cfn-appmesh-virtualnode-virtualnodename
	//
	VirtualNodeName *string `field:"optional" json:"virtualNodeName" yaml:"virtualNodeName"`
}

