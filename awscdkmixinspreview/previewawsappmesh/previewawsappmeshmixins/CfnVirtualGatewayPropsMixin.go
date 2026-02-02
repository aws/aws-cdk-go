package previewawsappmeshmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappmesh/previewawsappmeshmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a virtual gateway.
//
// A virtual gateway allows resources outside your mesh to communicate to resources that are inside your mesh. The virtual gateway represents an Envoy proxy running in an Amazon ECS task, in a Kubernetes service, or on an Amazon EC2 instance. Unlike a virtual node, which represents an Envoy running with an application, a virtual gateway represents Envoy deployed by itself.
//
// For more information about virtual gateways, see [Virtual gateways](https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_gateways.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVirtualGatewayPropsMixin := awscdkmixinspreview.Mixins.NewCfnVirtualGatewayPropsMixin(&CfnVirtualGatewayMixinProps{
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualgateway.html
//
type CfnVirtualGatewayPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnVirtualGatewayMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVirtualGatewayPropsMixin
type jsiiProxy_CfnVirtualGatewayPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnVirtualGatewayPropsMixin) Props() *CfnVirtualGatewayMixinProps {
	var returns *CfnVirtualGatewayMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualGatewayPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppMesh::VirtualGateway`.
func NewCfnVirtualGatewayPropsMixin(props *CfnVirtualGatewayMixinProps, options *mixins.CfnPropertyMixinOptions) CfnVirtualGatewayPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVirtualGatewayPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVirtualGatewayPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnVirtualGatewayPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppMesh::VirtualGateway`.
func NewCfnVirtualGatewayPropsMixin_Override(c CfnVirtualGatewayPropsMixin, props *CfnVirtualGatewayMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnVirtualGatewayPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnVirtualGatewayPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVirtualGatewayPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnVirtualGatewayPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVirtualGatewayPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnVirtualGatewayPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVirtualGatewayPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnVirtualGatewayPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

