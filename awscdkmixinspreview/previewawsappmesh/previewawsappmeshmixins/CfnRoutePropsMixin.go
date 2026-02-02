package previewawsappmeshmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappmesh/previewawsappmeshmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a route that is associated with a virtual router.
//
// You can route several different protocols and define a retry policy for a route. Traffic can be routed to one or more virtual nodes.
//
// For more information about routes, see [Routes](https://docs.aws.amazon.com/app-mesh/latest/userguide/routes.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRoutePropsMixin := awscdkmixinspreview.Mixins.NewCfnRoutePropsMixin(&CfnRouteMixinProps{
//   	MeshName: jsii.String("meshName"),
//   	MeshOwner: jsii.String("meshOwner"),
//   	RouteName: jsii.String("routeName"),
//   	Spec: &RouteSpecProperty{
//   		GrpcRoute: &GrpcRouteProperty{
//   			Action: &GrpcRouteActionProperty{
//   				WeightedTargets: []interface{}{
//   					&WeightedTargetProperty{
//   						Port: jsii.Number(123),
//   						VirtualNode: jsii.String("virtualNode"),
//   						Weight: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Match: &GrpcRouteMatchProperty{
//   				Metadata: []interface{}{
//   					&GrpcRouteMetadataProperty{
//   						Invert: jsii.Boolean(false),
//   						Match: &GrpcRouteMetadataMatchMethodProperty{
//   							Exact: jsii.String("exact"),
//   							Prefix: jsii.String("prefix"),
//   							Range: &MatchRangeProperty{
//   								End: jsii.Number(123),
//   								Start: jsii.Number(123),
//   							},
//   							Regex: jsii.String("regex"),
//   							Suffix: jsii.String("suffix"),
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				MethodName: jsii.String("methodName"),
//   				Port: jsii.Number(123),
//   				ServiceName: jsii.String("serviceName"),
//   			},
//   			RetryPolicy: &GrpcRetryPolicyProperty{
//   				GrpcRetryEvents: []*string{
//   					jsii.String("grpcRetryEvents"),
//   				},
//   				HttpRetryEvents: []*string{
//   					jsii.String("httpRetryEvents"),
//   				},
//   				MaxRetries: jsii.Number(123),
//   				PerRetryTimeout: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				TcpRetryEvents: []*string{
//   					jsii.String("tcpRetryEvents"),
//   				},
//   			},
//   			Timeout: &GrpcTimeoutProperty{
//   				Idle: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				PerRequest: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Http2Route: &HttpRouteProperty{
//   			Action: &HttpRouteActionProperty{
//   				WeightedTargets: []interface{}{
//   					&WeightedTargetProperty{
//   						Port: jsii.Number(123),
//   						VirtualNode: jsii.String("virtualNode"),
//   						Weight: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Match: &HttpRouteMatchProperty{
//   				Headers: []interface{}{
//   					&HttpRouteHeaderProperty{
//   						Invert: jsii.Boolean(false),
//   						Match: &HeaderMatchMethodProperty{
//   							Exact: jsii.String("exact"),
//   							Prefix: jsii.String("prefix"),
//   							Range: &MatchRangeProperty{
//   								End: jsii.Number(123),
//   								Start: jsii.Number(123),
//   							},
//   							Regex: jsii.String("regex"),
//   							Suffix: jsii.String("suffix"),
//   						},
//   						Name: jsii.String("name"),
//   					},
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
//   				Scheme: jsii.String("scheme"),
//   			},
//   			RetryPolicy: &HttpRetryPolicyProperty{
//   				HttpRetryEvents: []*string{
//   					jsii.String("httpRetryEvents"),
//   				},
//   				MaxRetries: jsii.Number(123),
//   				PerRetryTimeout: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				TcpRetryEvents: []*string{
//   					jsii.String("tcpRetryEvents"),
//   				},
//   			},
//   			Timeout: &HttpTimeoutProperty{
//   				Idle: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				PerRequest: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   		},
//   		HttpRoute: &HttpRouteProperty{
//   			Action: &HttpRouteActionProperty{
//   				WeightedTargets: []interface{}{
//   					&WeightedTargetProperty{
//   						Port: jsii.Number(123),
//   						VirtualNode: jsii.String("virtualNode"),
//   						Weight: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Match: &HttpRouteMatchProperty{
//   				Headers: []interface{}{
//   					&HttpRouteHeaderProperty{
//   						Invert: jsii.Boolean(false),
//   						Match: &HeaderMatchMethodProperty{
//   							Exact: jsii.String("exact"),
//   							Prefix: jsii.String("prefix"),
//   							Range: &MatchRangeProperty{
//   								End: jsii.Number(123),
//   								Start: jsii.Number(123),
//   							},
//   							Regex: jsii.String("regex"),
//   							Suffix: jsii.String("suffix"),
//   						},
//   						Name: jsii.String("name"),
//   					},
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
//   				Scheme: jsii.String("scheme"),
//   			},
//   			RetryPolicy: &HttpRetryPolicyProperty{
//   				HttpRetryEvents: []*string{
//   					jsii.String("httpRetryEvents"),
//   				},
//   				MaxRetries: jsii.Number(123),
//   				PerRetryTimeout: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				TcpRetryEvents: []*string{
//   					jsii.String("tcpRetryEvents"),
//   				},
//   			},
//   			Timeout: &HttpTimeoutProperty{
//   				Idle: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				PerRequest: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Priority: jsii.Number(123),
//   		TcpRoute: &TcpRouteProperty{
//   			Action: &TcpRouteActionProperty{
//   				WeightedTargets: []interface{}{
//   					&WeightedTargetProperty{
//   						Port: jsii.Number(123),
//   						VirtualNode: jsii.String("virtualNode"),
//   						Weight: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Match: &TcpRouteMatchProperty{
//   				Port: jsii.Number(123),
//   			},
//   			Timeout: &TcpTimeoutProperty{
//   				Idle: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
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
//   	VirtualRouterName: jsii.String("virtualRouterName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-route.html
//
type CfnRoutePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRouteMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRoutePropsMixin
type jsiiProxy_CfnRoutePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRoutePropsMixin) Props() *CfnRouteMixinProps {
	var returns *CfnRouteMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoutePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppMesh::Route`.
func NewCfnRoutePropsMixin(props *CfnRouteMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRoutePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRoutePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRoutePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnRoutePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppMesh::Route`.
func NewCfnRoutePropsMixin_Override(c CfnRoutePropsMixin, props *CfnRouteMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnRoutePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRoutePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRoutePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnRoutePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRoutePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnRoutePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRoutePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRoutePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

