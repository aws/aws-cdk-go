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
//   cfnVirtualGatewayProps := &cfnVirtualGatewayProps{
//   	meshName: jsii.String("meshName"),
//   	spec: &virtualGatewaySpecProperty{
//   		listeners: []interface{}{
//   			&virtualGatewayListenerProperty{
//   				portMapping: &virtualGatewayPortMappingProperty{
//   					port: jsii.Number(123),
//   					protocol: jsii.String("protocol"),
//   				},
//
//   				// the properties below are optional
//   				connectionPool: &virtualGatewayConnectionPoolProperty{
//   					grpc: &virtualGatewayGrpcConnectionPoolProperty{
//   						maxRequests: jsii.Number(123),
//   					},
//   					http: &virtualGatewayHttpConnectionPoolProperty{
//   						maxConnections: jsii.Number(123),
//
//   						// the properties below are optional
//   						maxPendingRequests: jsii.Number(123),
//   					},
//   					http2: &virtualGatewayHttp2ConnectionPoolProperty{
//   						maxRequests: jsii.Number(123),
//   					},
//   				},
//   				healthCheck: &virtualGatewayHealthCheckPolicyProperty{
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
//   				tls: &virtualGatewayListenerTlsProperty{
//   					certificate: &virtualGatewayListenerTlsCertificateProperty{
//   						acm: &virtualGatewayListenerTlsAcmCertificateProperty{
//   							certificateArn: jsii.String("certificateArn"),
//   						},
//   						file: &virtualGatewayListenerTlsFileCertificateProperty{
//   							certificateChain: jsii.String("certificateChain"),
//   							privateKey: jsii.String("privateKey"),
//   						},
//   						sds: &virtualGatewayListenerTlsSdsCertificateProperty{
//   							secretName: jsii.String("secretName"),
//   						},
//   					},
//   					mode: jsii.String("mode"),
//
//   					// the properties below are optional
//   					validation: &virtualGatewayListenerTlsValidationContextProperty{
//   						trust: &virtualGatewayListenerTlsValidationContextTrustProperty{
//   							file: &virtualGatewayTlsValidationContextFileTrustProperty{
//   								certificateChain: jsii.String("certificateChain"),
//   							},
//   							sds: &virtualGatewayTlsValidationContextSdsTrustProperty{
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
//
//   		// the properties below are optional
//   		backendDefaults: &virtualGatewayBackendDefaultsProperty{
//   			clientPolicy: &virtualGatewayClientPolicyProperty{
//   				tls: &virtualGatewayClientPolicyTlsProperty{
//   					validation: &virtualGatewayTlsValidationContextProperty{
//   						trust: &virtualGatewayTlsValidationContextTrustProperty{
//   							acm: &virtualGatewayTlsValidationContextAcmTrustProperty{
//   								certificateAuthorityArns: []*string{
//   									jsii.String("certificateAuthorityArns"),
//   								},
//   							},
//   							file: &virtualGatewayTlsValidationContextFileTrustProperty{
//   								certificateChain: jsii.String("certificateChain"),
//   							},
//   							sds: &virtualGatewayTlsValidationContextSdsTrustProperty{
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
//   					certificate: &virtualGatewayClientTlsCertificateProperty{
//   						file: &virtualGatewayListenerTlsFileCertificateProperty{
//   							certificateChain: jsii.String("certificateChain"),
//   							privateKey: jsii.String("privateKey"),
//   						},
//   						sds: &virtualGatewayListenerTlsSdsCertificateProperty{
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
//   		logging: &virtualGatewayLoggingProperty{
//   			accessLog: &virtualGatewayAccessLogProperty{
//   				file: &virtualGatewayFileAccessLogProperty{
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
//   	virtualGatewayName: jsii.String("virtualGatewayName"),
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

