package previewawsappmeshmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappmesh/previewawsappmeshmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a virtual node within a service mesh.
//
// A virtual node acts as a logical pointer to a particular task group, such as an Amazon ECS service or a Kubernetes deployment. When you create a virtual node, you can specify the service discovery information for your task group, and whether the proxy running in a task group will communicate with other proxies using Transport Layer Security (TLS).
//
// You define a `listener` for any inbound traffic that your virtual node expects. Any virtual service that your virtual node expects to communicate to is specified as a `backend` .
//
// The response metadata for your new virtual node contains the `arn` that is associated with the virtual node. Set this value to the full ARN; for example, `arn:aws:appmesh:us-west-2:123456789012:myMesh/default/virtualNode/myApp` ) as the `APPMESH_RESOURCE_ARN` environment variable for your task group's Envoy proxy container in your task definition or pod spec. This is then mapped to the `node.id` and `node.cluster` Envoy parameters.
//
// > By default, App Mesh uses the name of the resource you specified in `APPMESH_RESOURCE_ARN` when Envoy is referring to itself in metrics and traces. You can override this behavior by setting the `APPMESH_RESOURCE_CLUSTER` environment variable with your own name.
//
// For more information about virtual nodes, see [Virtual nodes](https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_nodes.html) . You must be using `1.15.0` or later of the Envoy image when setting these variables. For more information aboutApp Mesh Envoy variables, see [Envoy image](https://docs.aws.amazon.com/app-mesh/latest/userguide/envoy.html) in the AWS App Mesh User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVirtualNodePropsMixin := awscdkmixinspreview.Mixins.NewCfnVirtualNodePropsMixin(&CfnVirtualNodeMixinProps{
//   	MeshName: jsii.String("meshName"),
//   	MeshOwner: jsii.String("meshOwner"),
//   	Spec: &VirtualNodeSpecProperty{
//   		BackendDefaults: &BackendDefaultsProperty{
//   			ClientPolicy: &ClientPolicyProperty{
//   				Tls: &ClientPolicyTlsProperty{
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
//   					Validation: &TlsValidationContextProperty{
//   						SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   							Match: &SubjectAlternativeNameMatchersProperty{
//   								Exact: []*string{
//   									jsii.String("exact"),
//   								},
//   							},
//   						},
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
//   					},
//   				},
//   			},
//   		},
//   		Backends: []interface{}{
//   			&BackendProperty{
//   				VirtualService: &VirtualServiceBackendProperty{
//   					ClientPolicy: &ClientPolicyProperty{
//   						Tls: &ClientPolicyTlsProperty{
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
//   							Validation: &TlsValidationContextProperty{
//   								SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   									Match: &SubjectAlternativeNameMatchersProperty{
//   										Exact: []*string{
//   											jsii.String("exact"),
//   										},
//   									},
//   								},
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
//   							},
//   						},
//   					},
//   					VirtualServiceName: jsii.String("virtualServiceName"),
//   				},
//   			},
//   		},
//   		Listeners: []interface{}{
//   			&ListenerProperty{
//   				ConnectionPool: &VirtualNodeConnectionPoolProperty{
//   					Grpc: &VirtualNodeGrpcConnectionPoolProperty{
//   						MaxRequests: jsii.Number(123),
//   					},
//   					Http: &VirtualNodeHttpConnectionPoolProperty{
//   						MaxConnections: jsii.Number(123),
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
//   					Path: jsii.String("path"),
//   					Port: jsii.Number(123),
//   					Protocol: jsii.String("protocol"),
//   					TimeoutMillis: jsii.Number(123),
//   					UnhealthyThreshold: jsii.Number(123),
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
//   				PortMapping: &PortMappingProperty{
//   					Port: jsii.Number(123),
//   					Protocol: jsii.String("protocol"),
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
//   					Validation: &ListenerTlsValidationContextProperty{
//   						SubjectAlternativeNames: &SubjectAlternativeNamesProperty{
//   							Match: &SubjectAlternativeNameMatchersProperty{
//   								Exact: []*string{
//   									jsii.String("exact"),
//   								},
//   							},
//   						},
//   						Trust: &ListenerTlsValidationContextTrustProperty{
//   							File: &TlsValidationContextFileTrustProperty{
//   								CertificateChain: jsii.String("certificateChain"),
//   							},
//   							Sds: &TlsValidationContextSdsTrustProperty{
//   								SecretName: jsii.String("secretName"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		Logging: &LoggingProperty{
//   			AccessLog: &AccessLogProperty{
//   				File: &FileAccessLogProperty{
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
//   		ServiceDiscovery: &ServiceDiscoveryProperty{
//   			AwsCloudMap: &AwsCloudMapServiceDiscoveryProperty{
//   				Attributes: []interface{}{
//   					&AwsCloudMapInstanceAttributeProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				IpPreference: jsii.String("ipPreference"),
//   				NamespaceName: jsii.String("namespaceName"),
//   				ServiceName: jsii.String("serviceName"),
//   			},
//   			Dns: &DnsServiceDiscoveryProperty{
//   				Hostname: jsii.String("hostname"),
//   				IpPreference: jsii.String("ipPreference"),
//   				ResponseType: jsii.String("responseType"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VirtualNodeName: jsii.String("virtualNodeName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualnode.html
//
type CfnVirtualNodePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnVirtualNodeMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVirtualNodePropsMixin
type jsiiProxy_CfnVirtualNodePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnVirtualNodePropsMixin) Props() *CfnVirtualNodeMixinProps {
	var returns *CfnVirtualNodeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNodePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppMesh::VirtualNode`.
func NewCfnVirtualNodePropsMixin(props *CfnVirtualNodeMixinProps, options *mixins.CfnPropertyMixinOptions) CfnVirtualNodePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVirtualNodePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVirtualNodePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnVirtualNodePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppMesh::VirtualNode`.
func NewCfnVirtualNodePropsMixin_Override(c CfnVirtualNodePropsMixin, props *CfnVirtualNodeMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnVirtualNodePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnVirtualNodePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVirtualNodePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnVirtualNodePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVirtualNodePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnVirtualNodePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVirtualNodePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnVirtualNodePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

