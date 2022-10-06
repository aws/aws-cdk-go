package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnVirtualNode`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVirtualNodeProps := &cfnVirtualNodeProps{
//   	meshName: jsii.String("meshName"),
//   	spec: &virtualNodeSpecProperty{
//   		backendDefaults: &backendDefaultsProperty{
//   			clientPolicy: &clientPolicyProperty{
//   				tls: &clientPolicyTlsProperty{
//   					validation: &tlsValidationContextProperty{
//   						trust: &tlsValidationContextTrustProperty{
//   							acm: &tlsValidationContextAcmTrustProperty{
//   								certificateAuthorityArns: []*string{
//   									jsii.String("certificateAuthorityArns"),
//   								},
//   							},
//   							file: &tlsValidationContextFileTrustProperty{
//   								certificateChain: jsii.String("certificateChain"),
//   							},
//   							sds: &tlsValidationContextSdsTrustProperty{
//   								secretName: jsii.String("secretName"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   							match: &subjectAlternativeNameMatchersProperty{
//   								exact: []*string{
//   									jsii.String("exact"),
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					certificate: &clientTlsCertificateProperty{
//   						file: &listenerTlsFileCertificateProperty{
//   							certificateChain: jsii.String("certificateChain"),
//   							privateKey: jsii.String("privateKey"),
//   						},
//   						sds: &listenerTlsSdsCertificateProperty{
//   							secretName: jsii.String("secretName"),
//   						},
//   					},
//   					enforce: jsii.Boolean(false),
//   					ports: []interface{}{
//   						jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		backends: []interface{}{
//   			&backendProperty{
//   				virtualService: &virtualServiceBackendProperty{
//   					virtualServiceName: jsii.String("virtualServiceName"),
//
//   					// the properties below are optional
//   					clientPolicy: &clientPolicyProperty{
//   						tls: &clientPolicyTlsProperty{
//   							validation: &tlsValidationContextProperty{
//   								trust: &tlsValidationContextTrustProperty{
//   									acm: &tlsValidationContextAcmTrustProperty{
//   										certificateAuthorityArns: []*string{
//   											jsii.String("certificateAuthorityArns"),
//   										},
//   									},
//   									file: &tlsValidationContextFileTrustProperty{
//   										certificateChain: jsii.String("certificateChain"),
//   									},
//   									sds: &tlsValidationContextSdsTrustProperty{
//   										secretName: jsii.String("secretName"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   									match: &subjectAlternativeNameMatchersProperty{
//   										exact: []*string{
//   											jsii.String("exact"),
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							certificate: &clientTlsCertificateProperty{
//   								file: &listenerTlsFileCertificateProperty{
//   									certificateChain: jsii.String("certificateChain"),
//   									privateKey: jsii.String("privateKey"),
//   								},
//   								sds: &listenerTlsSdsCertificateProperty{
//   									secretName: jsii.String("secretName"),
//   								},
//   							},
//   							enforce: jsii.Boolean(false),
//   							ports: []interface{}{
//   								jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		listeners: []interface{}{
//   			&listenerProperty{
//   				portMapping: &portMappingProperty{
//   					port: jsii.Number(123),
//   					protocol: jsii.String("protocol"),
//   				},
//
//   				// the properties below are optional
//   				connectionPool: &virtualNodeConnectionPoolProperty{
//   					grpc: &virtualNodeGrpcConnectionPoolProperty{
//   						maxRequests: jsii.Number(123),
//   					},
//   					http: &virtualNodeHttpConnectionPoolProperty{
//   						maxConnections: jsii.Number(123),
//
//   						// the properties below are optional
//   						maxPendingRequests: jsii.Number(123),
//   					},
//   					http2: &virtualNodeHttp2ConnectionPoolProperty{
//   						maxRequests: jsii.Number(123),
//   					},
//   					tcp: &virtualNodeTcpConnectionPoolProperty{
//   						maxConnections: jsii.Number(123),
//   					},
//   				},
//   				healthCheck: &healthCheckProperty{
//   					healthyThreshold: jsii.Number(123),
//   					intervalMillis: jsii.Number(123),
//   					protocol: jsii.String("protocol"),
//   					timeoutMillis: jsii.Number(123),
//   					unhealthyThreshold: jsii.Number(123),
//
//   					// the properties below are optional
//   					path: jsii.String("path"),
//   					port: jsii.Number(123),
//   				},
//   				outlierDetection: &outlierDetectionProperty{
//   					baseEjectionDuration: &durationProperty{
//   						unit: jsii.String("unit"),
//   						value: jsii.Number(123),
//   					},
//   					interval: &durationProperty{
//   						unit: jsii.String("unit"),
//   						value: jsii.Number(123),
//   					},
//   					maxEjectionPercent: jsii.Number(123),
//   					maxServerErrors: jsii.Number(123),
//   				},
//   				timeout: &listenerTimeoutProperty{
//   					grpc: &grpcTimeoutProperty{
//   						idle: &durationProperty{
//   							unit: jsii.String("unit"),
//   							value: jsii.Number(123),
//   						},
//   						perRequest: &durationProperty{
//   							unit: jsii.String("unit"),
//   							value: jsii.Number(123),
//   						},
//   					},
//   					http: &httpTimeoutProperty{
//   						idle: &durationProperty{
//   							unit: jsii.String("unit"),
//   							value: jsii.Number(123),
//   						},
//   						perRequest: &durationProperty{
//   							unit: jsii.String("unit"),
//   							value: jsii.Number(123),
//   						},
//   					},
//   					http2: &httpTimeoutProperty{
//   						idle: &durationProperty{
//   							unit: jsii.String("unit"),
//   							value: jsii.Number(123),
//   						},
//   						perRequest: &durationProperty{
//   							unit: jsii.String("unit"),
//   							value: jsii.Number(123),
//   						},
//   					},
//   					tcp: &tcpTimeoutProperty{
//   						idle: &durationProperty{
//   							unit: jsii.String("unit"),
//   							value: jsii.Number(123),
//   						},
//   					},
//   				},
//   				tls: &listenerTlsProperty{
//   					certificate: &listenerTlsCertificateProperty{
//   						acm: &listenerTlsAcmCertificateProperty{
//   							certificateArn: jsii.String("certificateArn"),
//   						},
//   						file: &listenerTlsFileCertificateProperty{
//   							certificateChain: jsii.String("certificateChain"),
//   							privateKey: jsii.String("privateKey"),
//   						},
//   						sds: &listenerTlsSdsCertificateProperty{
//   							secretName: jsii.String("secretName"),
//   						},
//   					},
//   					mode: jsii.String("mode"),
//
//   					// the properties below are optional
//   					validation: &listenerTlsValidationContextProperty{
//   						trust: &listenerTlsValidationContextTrustProperty{
//   							file: &tlsValidationContextFileTrustProperty{
//   								certificateChain: jsii.String("certificateChain"),
//   							},
//   							sds: &tlsValidationContextSdsTrustProperty{
//   								secretName: jsii.String("secretName"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   							match: &subjectAlternativeNameMatchersProperty{
//   								exact: []*string{
//   									jsii.String("exact"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		logging: &loggingProperty{
//   			accessLog: &accessLogProperty{
//   				file: &fileAccessLogProperty{
//   					path: jsii.String("path"),
//
//   					// the properties below are optional
//   					format: &loggingFormatProperty{
//   						json: []interface{}{
//   							&jsonFormatRefProperty{
//   								key: jsii.String("key"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   						text: jsii.String("text"),
//   					},
//   				},
//   			},
//   		},
//   		serviceDiscovery: &serviceDiscoveryProperty{
//   			awsCloudMap: &awsCloudMapServiceDiscoveryProperty{
//   				namespaceName: jsii.String("namespaceName"),
//   				serviceName: jsii.String("serviceName"),
//
//   				// the properties below are optional
//   				attributes: []interface{}{
//   					&awsCloudMapInstanceAttributeProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				ipPreference: jsii.String("ipPreference"),
//   			},
//   			dns: &dnsServiceDiscoveryProperty{
//   				hostname: jsii.String("hostname"),
//
//   				// the properties below are optional
//   				ipPreference: jsii.String("ipPreference"),
//   				responseType: jsii.String("responseType"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	meshOwner: jsii.String("meshOwner"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	virtualNodeName: jsii.String("virtualNodeName"),
//   }
//
type CfnVirtualNodeProps struct {
	// The name of the service mesh to create the virtual node in.
	MeshName *string `field:"required" json:"meshName" yaml:"meshName"`
	// The virtual node specification to apply.
	Spec interface{} `field:"required" json:"spec" yaml:"spec"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then the account that you specify must share the mesh with your account before you can create the resource in the service mesh. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	MeshOwner *string `field:"optional" json:"meshOwner" yaml:"meshOwner"`
	// Optional metadata that you can apply to the virtual node to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name to use for the virtual node.
	VirtualNodeName *string `field:"optional" json:"virtualNodeName" yaml:"virtualNodeName"`
}

