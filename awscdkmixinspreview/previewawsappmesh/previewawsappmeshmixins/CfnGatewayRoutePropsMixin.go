package previewawsappmeshmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappmesh/previewawsappmeshmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a gateway route.
//
// A gateway route is attached to a virtual gateway and routes traffic to an existing virtual service. If a route matches a request, it can distribute traffic to a target virtual service.
//
// For more information about gateway routes, see [Gateway routes](https://docs.aws.amazon.com/app-mesh/latest/userguide/gateway-routes.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGatewayRoutePropsMixin := awscdkmixinspreview.Mixins.NewCfnGatewayRoutePropsMixin(&CfnGatewayRouteMixinProps{
//   	GatewayRouteName: jsii.String("gatewayRouteName"),
//   	MeshName: jsii.String("meshName"),
//   	MeshOwner: jsii.String("meshOwner"),
//   	Spec: &GatewayRouteSpecProperty{
//   		GrpcRoute: &GrpcGatewayRouteProperty{
//   			Action: &GrpcGatewayRouteActionProperty{
//   				Rewrite: &GrpcGatewayRouteRewriteProperty{
//   					Hostname: &GatewayRouteHostnameRewriteProperty{
//   						DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   					},
//   				},
//   				Target: &GatewayRouteTargetProperty{
//   					Port: jsii.Number(123),
//   					VirtualService: &GatewayRouteVirtualServiceProperty{
//   						VirtualServiceName: jsii.String("virtualServiceName"),
//   					},
//   				},
//   			},
//   			Match: &GrpcGatewayRouteMatchProperty{
//   				Hostname: &GatewayRouteHostnameMatchProperty{
//   					Exact: jsii.String("exact"),
//   					Suffix: jsii.String("suffix"),
//   				},
//   				Metadata: []interface{}{
//   					&GrpcGatewayRouteMetadataProperty{
//   						Invert: jsii.Boolean(false),
//   						Match: &GatewayRouteMetadataMatchProperty{
//   							Exact: jsii.String("exact"),
//   							Prefix: jsii.String("prefix"),
//   							Range: &GatewayRouteRangeMatchProperty{
//   								End: jsii.Number(123),
//   								Start: jsii.Number(123),
//   							},
//   							Regex: jsii.String("regex"),
//   							Suffix: jsii.String("suffix"),
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Port: jsii.Number(123),
//   				ServiceName: jsii.String("serviceName"),
//   			},
//   		},
//   		Http2Route: &HttpGatewayRouteProperty{
//   			Action: &HttpGatewayRouteActionProperty{
//   				Rewrite: &HttpGatewayRouteRewriteProperty{
//   					Hostname: &GatewayRouteHostnameRewriteProperty{
//   						DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   					},
//   					Path: &HttpGatewayRoutePathRewriteProperty{
//   						Exact: jsii.String("exact"),
//   					},
//   					Prefix: &HttpGatewayRoutePrefixRewriteProperty{
//   						DefaultPrefix: jsii.String("defaultPrefix"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Target: &GatewayRouteTargetProperty{
//   					Port: jsii.Number(123),
//   					VirtualService: &GatewayRouteVirtualServiceProperty{
//   						VirtualServiceName: jsii.String("virtualServiceName"),
//   					},
//   				},
//   			},
//   			Match: &HttpGatewayRouteMatchProperty{
//   				Headers: []interface{}{
//   					&HttpGatewayRouteHeaderProperty{
//   						Invert: jsii.Boolean(false),
//   						Match: &HttpGatewayRouteHeaderMatchProperty{
//   							Exact: jsii.String("exact"),
//   							Prefix: jsii.String("prefix"),
//   							Range: &GatewayRouteRangeMatchProperty{
//   								End: jsii.Number(123),
//   								Start: jsii.Number(123),
//   							},
//   							Regex: jsii.String("regex"),
//   							Suffix: jsii.String("suffix"),
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Hostname: &GatewayRouteHostnameMatchProperty{
//   					Exact: jsii.String("exact"),
//   					Suffix: jsii.String("suffix"),
//   				},
//   				Method: jsii.String("method"),
//   				Path: &HttpPathMatchProperty{
//   					Exact: jsii.String("exact"),
//   					Regex: jsii.String("regex"),
//   				},
//   				Port: jsii.Number(123),
//   				Prefix: jsii.String("prefix"),
//   				QueryParameters: []interface{}{
//   					&QueryParameterProperty{
//   						Match: &HttpQueryParameterMatchProperty{
//   							Exact: jsii.String("exact"),
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   		HttpRoute: &HttpGatewayRouteProperty{
//   			Action: &HttpGatewayRouteActionProperty{
//   				Rewrite: &HttpGatewayRouteRewriteProperty{
//   					Hostname: &GatewayRouteHostnameRewriteProperty{
//   						DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   					},
//   					Path: &HttpGatewayRoutePathRewriteProperty{
//   						Exact: jsii.String("exact"),
//   					},
//   					Prefix: &HttpGatewayRoutePrefixRewriteProperty{
//   						DefaultPrefix: jsii.String("defaultPrefix"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Target: &GatewayRouteTargetProperty{
//   					Port: jsii.Number(123),
//   					VirtualService: &GatewayRouteVirtualServiceProperty{
//   						VirtualServiceName: jsii.String("virtualServiceName"),
//   					},
//   				},
//   			},
//   			Match: &HttpGatewayRouteMatchProperty{
//   				Headers: []interface{}{
//   					&HttpGatewayRouteHeaderProperty{
//   						Invert: jsii.Boolean(false),
//   						Match: &HttpGatewayRouteHeaderMatchProperty{
//   							Exact: jsii.String("exact"),
//   							Prefix: jsii.String("prefix"),
//   							Range: &GatewayRouteRangeMatchProperty{
//   								End: jsii.Number(123),
//   								Start: jsii.Number(123),
//   							},
//   							Regex: jsii.String("regex"),
//   							Suffix: jsii.String("suffix"),
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Hostname: &GatewayRouteHostnameMatchProperty{
//   					Exact: jsii.String("exact"),
//   					Suffix: jsii.String("suffix"),
//   				},
//   				Method: jsii.String("method"),
//   				Path: &HttpPathMatchProperty{
//   					Exact: jsii.String("exact"),
//   					Regex: jsii.String("regex"),
//   				},
//   				Port: jsii.Number(123),
//   				Prefix: jsii.String("prefix"),
//   				QueryParameters: []interface{}{
//   					&QueryParameterProperty{
//   						Match: &HttpQueryParameterMatchProperty{
//   							Exact: jsii.String("exact"),
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   		Priority: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-gatewayroute.html
//
type CfnGatewayRoutePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnGatewayRouteMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGatewayRoutePropsMixin
type jsiiProxy_CfnGatewayRoutePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnGatewayRoutePropsMixin) Props() *CfnGatewayRouteMixinProps {
	var returns *CfnGatewayRouteMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRoutePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppMesh::GatewayRoute`.
func NewCfnGatewayRoutePropsMixin(props *CfnGatewayRouteMixinProps, options *mixins.CfnPropertyMixinOptions) CfnGatewayRoutePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGatewayRoutePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGatewayRoutePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnGatewayRoutePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppMesh::GatewayRoute`.
func NewCfnGatewayRoutePropsMixin_Override(c CfnGatewayRoutePropsMixin, props *CfnGatewayRouteMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnGatewayRoutePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnGatewayRoutePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGatewayRoutePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnGatewayRoutePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGatewayRoutePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnGatewayRoutePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGatewayRoutePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGatewayRoutePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

