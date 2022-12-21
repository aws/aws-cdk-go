package awsappmesh

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.AccessLog",
		reflect.TypeOf((*AccessLog)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_AccessLog{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.AccessLogConfig",
		reflect.TypeOf((*AccessLogConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.Backend",
		reflect.TypeOf((*Backend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_Backend{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.BackendConfig",
		reflect.TypeOf((*BackendConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.BackendDefaults",
		reflect.TypeOf((*BackendDefaults)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute",
		reflect.TypeOf((*CfnGatewayRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrGatewayRouteName", GoGetter: "AttrGatewayRouteName"},
			_jsii_.MemberProperty{JsiiProperty: "attrMeshName", GoGetter: "AttrMeshName"},
			_jsii_.MemberProperty{JsiiProperty: "attrMeshOwner", GoGetter: "AttrMeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceOwner", GoGetter: "AttrResourceOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrUid", GoGetter: "AttrUid"},
			_jsii_.MemberProperty{JsiiProperty: "attrVirtualGatewayName", GoGetter: "AttrVirtualGatewayName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayRouteName", GoGetter: "GatewayRouteName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwner", GoGetter: "MeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "virtualGatewayName", GoGetter: "VirtualGatewayName"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGatewayRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.GatewayRouteHostnameMatchProperty",
		reflect.TypeOf((*CfnGatewayRoute_GatewayRouteHostnameMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.GatewayRouteHostnameRewriteProperty",
		reflect.TypeOf((*CfnGatewayRoute_GatewayRouteHostnameRewriteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.GatewayRouteMetadataMatchProperty",
		reflect.TypeOf((*CfnGatewayRoute_GatewayRouteMetadataMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.GatewayRouteRangeMatchProperty",
		reflect.TypeOf((*CfnGatewayRoute_GatewayRouteRangeMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.GatewayRouteSpecProperty",
		reflect.TypeOf((*CfnGatewayRoute_GatewayRouteSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.GatewayRouteTargetProperty",
		reflect.TypeOf((*CfnGatewayRoute_GatewayRouteTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.GatewayRouteVirtualServiceProperty",
		reflect.TypeOf((*CfnGatewayRoute_GatewayRouteVirtualServiceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.GrpcGatewayRouteActionProperty",
		reflect.TypeOf((*CfnGatewayRoute_GrpcGatewayRouteActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.GrpcGatewayRouteMatchProperty",
		reflect.TypeOf((*CfnGatewayRoute_GrpcGatewayRouteMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.GrpcGatewayRouteMetadataProperty",
		reflect.TypeOf((*CfnGatewayRoute_GrpcGatewayRouteMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.GrpcGatewayRouteProperty",
		reflect.TypeOf((*CfnGatewayRoute_GrpcGatewayRouteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.GrpcGatewayRouteRewriteProperty",
		reflect.TypeOf((*CfnGatewayRoute_GrpcGatewayRouteRewriteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.HttpGatewayRouteActionProperty",
		reflect.TypeOf((*CfnGatewayRoute_HttpGatewayRouteActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.HttpGatewayRouteHeaderMatchProperty",
		reflect.TypeOf((*CfnGatewayRoute_HttpGatewayRouteHeaderMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.HttpGatewayRouteHeaderProperty",
		reflect.TypeOf((*CfnGatewayRoute_HttpGatewayRouteHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.HttpGatewayRouteMatchProperty",
		reflect.TypeOf((*CfnGatewayRoute_HttpGatewayRouteMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.HttpGatewayRoutePathRewriteProperty",
		reflect.TypeOf((*CfnGatewayRoute_HttpGatewayRoutePathRewriteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.HttpGatewayRoutePrefixRewriteProperty",
		reflect.TypeOf((*CfnGatewayRoute_HttpGatewayRoutePrefixRewriteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.HttpGatewayRouteProperty",
		reflect.TypeOf((*CfnGatewayRoute_HttpGatewayRouteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.HttpGatewayRouteRewriteProperty",
		reflect.TypeOf((*CfnGatewayRoute_HttpGatewayRouteRewriteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.HttpPathMatchProperty",
		reflect.TypeOf((*CfnGatewayRoute_HttpPathMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.HttpQueryParameterMatchProperty",
		reflect.TypeOf((*CfnGatewayRoute_HttpQueryParameterMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.QueryParameterProperty",
		reflect.TypeOf((*CfnGatewayRoute_QueryParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRouteProps",
		reflect.TypeOf((*CfnGatewayRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.CfnMesh",
		reflect.TypeOf((*CfnMesh)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrMeshName", GoGetter: "AttrMeshName"},
			_jsii_.MemberProperty{JsiiProperty: "attrMeshOwner", GoGetter: "AttrMeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceOwner", GoGetter: "AttrResourceOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrUid", GoGetter: "AttrUid"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMesh{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnMesh.EgressFilterProperty",
		reflect.TypeOf((*CfnMesh_EgressFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnMesh.MeshServiceDiscoveryProperty",
		reflect.TypeOf((*CfnMesh_MeshServiceDiscoveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnMesh.MeshSpecProperty",
		reflect.TypeOf((*CfnMesh_MeshSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnMeshProps",
		reflect.TypeOf((*CfnMeshProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.CfnRoute",
		reflect.TypeOf((*CfnRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrMeshName", GoGetter: "AttrMeshName"},
			_jsii_.MemberProperty{JsiiProperty: "attrMeshOwner", GoGetter: "AttrMeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceOwner", GoGetter: "AttrResourceOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrRouteName", GoGetter: "AttrRouteName"},
			_jsii_.MemberProperty{JsiiProperty: "attrUid", GoGetter: "AttrUid"},
			_jsii_.MemberProperty{JsiiProperty: "attrVirtualRouterName", GoGetter: "AttrVirtualRouterName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwner", GoGetter: "MeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "routeName", GoGetter: "RouteName"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouterName", GoGetter: "VirtualRouterName"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.DurationProperty",
		reflect.TypeOf((*CfnRoute_DurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.GrpcRetryPolicyProperty",
		reflect.TypeOf((*CfnRoute_GrpcRetryPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.GrpcRouteActionProperty",
		reflect.TypeOf((*CfnRoute_GrpcRouteActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.GrpcRouteMatchProperty",
		reflect.TypeOf((*CfnRoute_GrpcRouteMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.GrpcRouteMetadataMatchMethodProperty",
		reflect.TypeOf((*CfnRoute_GrpcRouteMetadataMatchMethodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.GrpcRouteMetadataProperty",
		reflect.TypeOf((*CfnRoute_GrpcRouteMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.GrpcRouteProperty",
		reflect.TypeOf((*CfnRoute_GrpcRouteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.GrpcTimeoutProperty",
		reflect.TypeOf((*CfnRoute_GrpcTimeoutProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.HeaderMatchMethodProperty",
		reflect.TypeOf((*CfnRoute_HeaderMatchMethodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.HttpPathMatchProperty",
		reflect.TypeOf((*CfnRoute_HttpPathMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.HttpQueryParameterMatchProperty",
		reflect.TypeOf((*CfnRoute_HttpQueryParameterMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.HttpRetryPolicyProperty",
		reflect.TypeOf((*CfnRoute_HttpRetryPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.HttpRouteActionProperty",
		reflect.TypeOf((*CfnRoute_HttpRouteActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.HttpRouteHeaderProperty",
		reflect.TypeOf((*CfnRoute_HttpRouteHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.HttpRouteMatchProperty",
		reflect.TypeOf((*CfnRoute_HttpRouteMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.HttpRouteProperty",
		reflect.TypeOf((*CfnRoute_HttpRouteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.HttpTimeoutProperty",
		reflect.TypeOf((*CfnRoute_HttpTimeoutProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.MatchRangeProperty",
		reflect.TypeOf((*CfnRoute_MatchRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.QueryParameterProperty",
		reflect.TypeOf((*CfnRoute_QueryParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.RouteSpecProperty",
		reflect.TypeOf((*CfnRoute_RouteSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.TcpRouteActionProperty",
		reflect.TypeOf((*CfnRoute_TcpRouteActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.TcpRouteMatchProperty",
		reflect.TypeOf((*CfnRoute_TcpRouteMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.TcpRouteProperty",
		reflect.TypeOf((*CfnRoute_TcpRouteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.TcpTimeoutProperty",
		reflect.TypeOf((*CfnRoute_TcpTimeoutProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.WeightedTargetProperty",
		reflect.TypeOf((*CfnRoute_WeightedTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRouteProps",
		reflect.TypeOf((*CfnRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway",
		reflect.TypeOf((*CfnVirtualGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrMeshName", GoGetter: "AttrMeshName"},
			_jsii_.MemberProperty{JsiiProperty: "attrMeshOwner", GoGetter: "AttrMeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceOwner", GoGetter: "AttrResourceOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrUid", GoGetter: "AttrUid"},
			_jsii_.MemberProperty{JsiiProperty: "attrVirtualGatewayName", GoGetter: "AttrVirtualGatewayName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwner", GoGetter: "MeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "virtualGatewayName", GoGetter: "VirtualGatewayName"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVirtualGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.JsonFormatRefProperty",
		reflect.TypeOf((*CfnVirtualGateway_JsonFormatRefProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.LoggingFormatProperty",
		reflect.TypeOf((*CfnVirtualGateway_LoggingFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.SubjectAlternativeNameMatchersProperty",
		reflect.TypeOf((*CfnVirtualGateway_SubjectAlternativeNameMatchersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.SubjectAlternativeNamesProperty",
		reflect.TypeOf((*CfnVirtualGateway_SubjectAlternativeNamesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayAccessLogProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayAccessLogProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayBackendDefaultsProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayBackendDefaultsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayClientPolicyProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayClientPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayClientPolicyTlsProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayClientPolicyTlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayClientTlsCertificateProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayClientTlsCertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayConnectionPoolProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayConnectionPoolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayFileAccessLogProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayFileAccessLogProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayGrpcConnectionPoolProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayGrpcConnectionPoolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayHealthCheckPolicyProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayHealthCheckPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayHttp2ConnectionPoolProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayHttp2ConnectionPoolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayHttpConnectionPoolProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayHttpConnectionPoolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayListenerProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayListenerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayListenerTlsAcmCertificateProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayListenerTlsAcmCertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayListenerTlsCertificateProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayListenerTlsCertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayListenerTlsFileCertificateProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayListenerTlsFileCertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayListenerTlsProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayListenerTlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayListenerTlsSdsCertificateProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayListenerTlsSdsCertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayListenerTlsValidationContextProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayListenerTlsValidationContextProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayListenerTlsValidationContextTrustProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayListenerTlsValidationContextTrustProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayLoggingProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayLoggingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayPortMappingProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayPortMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewaySpecProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewaySpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayTlsValidationContextAcmTrustProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayTlsValidationContextAcmTrustProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayTlsValidationContextFileTrustProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayTlsValidationContextFileTrustProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayTlsValidationContextProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayTlsValidationContextProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayTlsValidationContextSdsTrustProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayTlsValidationContextSdsTrustProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGateway.VirtualGatewayTlsValidationContextTrustProperty",
		reflect.TypeOf((*CfnVirtualGateway_VirtualGatewayTlsValidationContextTrustProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualGatewayProps",
		reflect.TypeOf((*CfnVirtualGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		reflect.TypeOf((*CfnVirtualNode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrMeshName", GoGetter: "AttrMeshName"},
			_jsii_.MemberProperty{JsiiProperty: "attrMeshOwner", GoGetter: "AttrMeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceOwner", GoGetter: "AttrResourceOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrUid", GoGetter: "AttrUid"},
			_jsii_.MemberProperty{JsiiProperty: "attrVirtualNodeName", GoGetter: "AttrVirtualNodeName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwner", GoGetter: "MeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNodeName", GoGetter: "VirtualNodeName"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVirtualNode{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.AccessLogProperty",
		reflect.TypeOf((*CfnVirtualNode_AccessLogProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.AwsCloudMapInstanceAttributeProperty",
		reflect.TypeOf((*CfnVirtualNode_AwsCloudMapInstanceAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.AwsCloudMapServiceDiscoveryProperty",
		reflect.TypeOf((*CfnVirtualNode_AwsCloudMapServiceDiscoveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.BackendDefaultsProperty",
		reflect.TypeOf((*CfnVirtualNode_BackendDefaultsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.BackendProperty",
		reflect.TypeOf((*CfnVirtualNode_BackendProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.ClientPolicyProperty",
		reflect.TypeOf((*CfnVirtualNode_ClientPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.ClientPolicyTlsProperty",
		reflect.TypeOf((*CfnVirtualNode_ClientPolicyTlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.ClientTlsCertificateProperty",
		reflect.TypeOf((*CfnVirtualNode_ClientTlsCertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.DnsServiceDiscoveryProperty",
		reflect.TypeOf((*CfnVirtualNode_DnsServiceDiscoveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.DurationProperty",
		reflect.TypeOf((*CfnVirtualNode_DurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.FileAccessLogProperty",
		reflect.TypeOf((*CfnVirtualNode_FileAccessLogProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.GrpcTimeoutProperty",
		reflect.TypeOf((*CfnVirtualNode_GrpcTimeoutProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.HealthCheckProperty",
		reflect.TypeOf((*CfnVirtualNode_HealthCheckProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.HttpTimeoutProperty",
		reflect.TypeOf((*CfnVirtualNode_HttpTimeoutProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.JsonFormatRefProperty",
		reflect.TypeOf((*CfnVirtualNode_JsonFormatRefProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.ListenerProperty",
		reflect.TypeOf((*CfnVirtualNode_ListenerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.ListenerTimeoutProperty",
		reflect.TypeOf((*CfnVirtualNode_ListenerTimeoutProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.ListenerTlsAcmCertificateProperty",
		reflect.TypeOf((*CfnVirtualNode_ListenerTlsAcmCertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.ListenerTlsCertificateProperty",
		reflect.TypeOf((*CfnVirtualNode_ListenerTlsCertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.ListenerTlsFileCertificateProperty",
		reflect.TypeOf((*CfnVirtualNode_ListenerTlsFileCertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.ListenerTlsProperty",
		reflect.TypeOf((*CfnVirtualNode_ListenerTlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.ListenerTlsSdsCertificateProperty",
		reflect.TypeOf((*CfnVirtualNode_ListenerTlsSdsCertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.ListenerTlsValidationContextProperty",
		reflect.TypeOf((*CfnVirtualNode_ListenerTlsValidationContextProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.ListenerTlsValidationContextTrustProperty",
		reflect.TypeOf((*CfnVirtualNode_ListenerTlsValidationContextTrustProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.LoggingFormatProperty",
		reflect.TypeOf((*CfnVirtualNode_LoggingFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.LoggingProperty",
		reflect.TypeOf((*CfnVirtualNode_LoggingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.OutlierDetectionProperty",
		reflect.TypeOf((*CfnVirtualNode_OutlierDetectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.PortMappingProperty",
		reflect.TypeOf((*CfnVirtualNode_PortMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.ServiceDiscoveryProperty",
		reflect.TypeOf((*CfnVirtualNode_ServiceDiscoveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.SubjectAlternativeNameMatchersProperty",
		reflect.TypeOf((*CfnVirtualNode_SubjectAlternativeNameMatchersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.SubjectAlternativeNamesProperty",
		reflect.TypeOf((*CfnVirtualNode_SubjectAlternativeNamesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.TcpTimeoutProperty",
		reflect.TypeOf((*CfnVirtualNode_TcpTimeoutProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.TlsValidationContextAcmTrustProperty",
		reflect.TypeOf((*CfnVirtualNode_TlsValidationContextAcmTrustProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.TlsValidationContextFileTrustProperty",
		reflect.TypeOf((*CfnVirtualNode_TlsValidationContextFileTrustProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.TlsValidationContextProperty",
		reflect.TypeOf((*CfnVirtualNode_TlsValidationContextProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.TlsValidationContextSdsTrustProperty",
		reflect.TypeOf((*CfnVirtualNode_TlsValidationContextSdsTrustProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.TlsValidationContextTrustProperty",
		reflect.TypeOf((*CfnVirtualNode_TlsValidationContextTrustProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.VirtualNodeConnectionPoolProperty",
		reflect.TypeOf((*CfnVirtualNode_VirtualNodeConnectionPoolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.VirtualNodeGrpcConnectionPoolProperty",
		reflect.TypeOf((*CfnVirtualNode_VirtualNodeGrpcConnectionPoolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.VirtualNodeHttp2ConnectionPoolProperty",
		reflect.TypeOf((*CfnVirtualNode_VirtualNodeHttp2ConnectionPoolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.VirtualNodeHttpConnectionPoolProperty",
		reflect.TypeOf((*CfnVirtualNode_VirtualNodeHttpConnectionPoolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.VirtualNodeSpecProperty",
		reflect.TypeOf((*CfnVirtualNode_VirtualNodeSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.VirtualNodeTcpConnectionPoolProperty",
		reflect.TypeOf((*CfnVirtualNode_VirtualNodeTcpConnectionPoolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode.VirtualServiceBackendProperty",
		reflect.TypeOf((*CfnVirtualNode_VirtualServiceBackendProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNodeProps",
		reflect.TypeOf((*CfnVirtualNodeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.CfnVirtualRouter",
		reflect.TypeOf((*CfnVirtualRouter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrMeshName", GoGetter: "AttrMeshName"},
			_jsii_.MemberProperty{JsiiProperty: "attrMeshOwner", GoGetter: "AttrMeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceOwner", GoGetter: "AttrResourceOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrUid", GoGetter: "AttrUid"},
			_jsii_.MemberProperty{JsiiProperty: "attrVirtualRouterName", GoGetter: "AttrVirtualRouterName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwner", GoGetter: "MeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouterName", GoGetter: "VirtualRouterName"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVirtualRouter{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualRouter.PortMappingProperty",
		reflect.TypeOf((*CfnVirtualRouter_PortMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualRouter.VirtualRouterListenerProperty",
		reflect.TypeOf((*CfnVirtualRouter_VirtualRouterListenerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualRouter.VirtualRouterSpecProperty",
		reflect.TypeOf((*CfnVirtualRouter_VirtualRouterSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualRouterProps",
		reflect.TypeOf((*CfnVirtualRouterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.CfnVirtualService",
		reflect.TypeOf((*CfnVirtualService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrMeshName", GoGetter: "AttrMeshName"},
			_jsii_.MemberProperty{JsiiProperty: "attrMeshOwner", GoGetter: "AttrMeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceOwner", GoGetter: "AttrResourceOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrUid", GoGetter: "AttrUid"},
			_jsii_.MemberProperty{JsiiProperty: "attrVirtualServiceName", GoGetter: "AttrVirtualServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwner", GoGetter: "MeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceName", GoGetter: "VirtualServiceName"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVirtualService{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualService.VirtualNodeServiceProviderProperty",
		reflect.TypeOf((*CfnVirtualService_VirtualNodeServiceProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualService.VirtualRouterServiceProviderProperty",
		reflect.TypeOf((*CfnVirtualService_VirtualRouterServiceProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualService.VirtualServiceProviderProperty",
		reflect.TypeOf((*CfnVirtualService_VirtualServiceProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualService.VirtualServiceSpecProperty",
		reflect.TypeOf((*CfnVirtualService_VirtualServiceSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnVirtualServiceProps",
		reflect.TypeOf((*CfnVirtualServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CommonGatewayRouteSpecOptions",
		reflect.TypeOf((*CommonGatewayRouteSpecOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appmesh.DnsResponseType",
		reflect.TypeOf((*DnsResponseType)(nil)).Elem(),
		map[string]interface{}{
			"LOAD_BALANCER": DnsResponseType_LOAD_BALANCER,
			"ENDPOINTS": DnsResponseType_ENDPOINTS,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.GatewayRoute",
		reflect.TypeOf((*GatewayRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayRouteArn", GoGetter: "GatewayRouteArn"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayRouteName", GoGetter: "GatewayRouteName"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "virtualGateway", GoGetter: "VirtualGateway"},
		},
		func() interface{} {
			j := jsiiProxy_GatewayRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGatewayRoute)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.GatewayRouteAttributes",
		reflect.TypeOf((*GatewayRouteAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.GatewayRouteBaseProps",
		reflect.TypeOf((*GatewayRouteBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.GatewayRouteHostnameMatch",
		reflect.TypeOf((*GatewayRouteHostnameMatch)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_GatewayRouteHostnameMatch{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.GatewayRouteHostnameMatchConfig",
		reflect.TypeOf((*GatewayRouteHostnameMatchConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.GatewayRouteProps",
		reflect.TypeOf((*GatewayRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.GatewayRouteSpec",
		reflect.TypeOf((*GatewayRouteSpec)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_GatewayRouteSpec{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.GatewayRouteSpecConfig",
		reflect.TypeOf((*GatewayRouteSpecConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.GrpcConnectionPool",
		reflect.TypeOf((*GrpcConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.GrpcGatewayListenerOptions",
		reflect.TypeOf((*GrpcGatewayListenerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.GrpcGatewayRouteMatch",
		reflect.TypeOf((*GrpcGatewayRouteMatch)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.GrpcGatewayRouteSpecOptions",
		reflect.TypeOf((*GrpcGatewayRouteSpecOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.GrpcHealthCheckOptions",
		reflect.TypeOf((*GrpcHealthCheckOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appmesh.GrpcRetryEvent",
		reflect.TypeOf((*GrpcRetryEvent)(nil)).Elem(),
		map[string]interface{}{
			"CANCELLED": GrpcRetryEvent_CANCELLED,
			"DEADLINE_EXCEEDED": GrpcRetryEvent_DEADLINE_EXCEEDED,
			"INTERNAL_ERROR": GrpcRetryEvent_INTERNAL_ERROR,
			"RESOURCE_EXHAUSTED": GrpcRetryEvent_RESOURCE_EXHAUSTED,
			"UNAVAILABLE": GrpcRetryEvent_UNAVAILABLE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.GrpcRetryPolicy",
		reflect.TypeOf((*GrpcRetryPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.GrpcRouteMatch",
		reflect.TypeOf((*GrpcRouteMatch)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.GrpcRouteSpecOptions",
		reflect.TypeOf((*GrpcRouteSpecOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.GrpcTimeout",
		reflect.TypeOf((*GrpcTimeout)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.GrpcVirtualNodeListenerOptions",
		reflect.TypeOf((*GrpcVirtualNodeListenerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		reflect.TypeOf((*HeaderMatch)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_HeaderMatch{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.HeaderMatchConfig",
		reflect.TypeOf((*HeaderMatchConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.HealthCheck",
		reflect.TypeOf((*HealthCheck)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_HealthCheck{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.HealthCheckBindOptions",
		reflect.TypeOf((*HealthCheckBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.HealthCheckConfig",
		reflect.TypeOf((*HealthCheckConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.Http2ConnectionPool",
		reflect.TypeOf((*Http2ConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.Http2GatewayListenerOptions",
		reflect.TypeOf((*Http2GatewayListenerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.Http2VirtualNodeListenerOptions",
		reflect.TypeOf((*Http2VirtualNodeListenerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.HttpConnectionPool",
		reflect.TypeOf((*HttpConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.HttpGatewayListenerOptions",
		reflect.TypeOf((*HttpGatewayListenerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.HttpGatewayRouteMatch",
		reflect.TypeOf((*HttpGatewayRouteMatch)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.HttpGatewayRoutePathMatch",
		reflect.TypeOf((*HttpGatewayRoutePathMatch)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_HttpGatewayRoutePathMatch{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.HttpGatewayRoutePathMatchConfig",
		reflect.TypeOf((*HttpGatewayRoutePathMatchConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.HttpGatewayRouteSpecOptions",
		reflect.TypeOf((*HttpGatewayRouteSpecOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.HttpHealthCheckOptions",
		reflect.TypeOf((*HttpHealthCheckOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appmesh.HttpRetryEvent",
		reflect.TypeOf((*HttpRetryEvent)(nil)).Elem(),
		map[string]interface{}{
			"SERVER_ERROR": HttpRetryEvent_SERVER_ERROR,
			"GATEWAY_ERROR": HttpRetryEvent_GATEWAY_ERROR,
			"CLIENT_ERROR": HttpRetryEvent_CLIENT_ERROR,
			"STREAM_ERROR": HttpRetryEvent_STREAM_ERROR,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.HttpRetryPolicy",
		reflect.TypeOf((*HttpRetryPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.HttpRouteMatch",
		reflect.TypeOf((*HttpRouteMatch)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appmesh.HttpRouteMethod",
		reflect.TypeOf((*HttpRouteMethod)(nil)).Elem(),
		map[string]interface{}{
			"GET": HttpRouteMethod_GET,
			"HEAD": HttpRouteMethod_HEAD,
			"POST": HttpRouteMethod_POST,
			"PUT": HttpRouteMethod_PUT,
			"DELETE": HttpRouteMethod_DELETE,
			"CONNECT": HttpRouteMethod_CONNECT,
			"OPTIONS": HttpRouteMethod_OPTIONS,
			"TRACE": HttpRouteMethod_TRACE,
			"PATCH": HttpRouteMethod_PATCH,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.HttpRoutePathMatch",
		reflect.TypeOf((*HttpRoutePathMatch)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_HttpRoutePathMatch{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.HttpRoutePathMatchConfig",
		reflect.TypeOf((*HttpRoutePathMatchConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appmesh.HttpRouteProtocol",
		reflect.TypeOf((*HttpRouteProtocol)(nil)).Elem(),
		map[string]interface{}{
			"HTTP": HttpRouteProtocol_HTTP,
			"HTTPS": HttpRouteProtocol_HTTPS,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.HttpRouteSpecOptions",
		reflect.TypeOf((*HttpRouteSpecOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.HttpTimeout",
		reflect.TypeOf((*HttpTimeout)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.HttpVirtualNodeListenerOptions",
		reflect.TypeOf((*HttpVirtualNodeListenerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appmesh.IGatewayRoute",
		reflect.TypeOf((*IGatewayRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayRouteArn", GoGetter: "GatewayRouteArn"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayRouteName", GoGetter: "GatewayRouteName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "virtualGateway", GoGetter: "VirtualGateway"},
		},
		func() interface{} {
			j := jsiiProxy_IGatewayRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appmesh.IMesh",
		reflect.TypeOf((*IMesh)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addVirtualGateway", GoMethod: "AddVirtualGateway"},
			_jsii_.MemberMethod{JsiiMethod: "addVirtualNode", GoMethod: "AddVirtualNode"},
			_jsii_.MemberMethod{JsiiMethod: "addVirtualRouter", GoMethod: "AddVirtualRouter"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "meshArn", GoGetter: "MeshArn"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IMesh{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appmesh.IRoute",
		reflect.TypeOf((*IRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "routeArn", GoGetter: "RouteArn"},
			_jsii_.MemberProperty{JsiiProperty: "routeName", GoGetter: "RouteName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouter", GoGetter: "VirtualRouter"},
		},
		func() interface{} {
			j := jsiiProxy_IRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appmesh.IVirtualGateway",
		reflect.TypeOf((*IVirtualGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addGatewayRoute", GoMethod: "AddGatewayRoute"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grantStreamAggregatedResources", GoMethod: "GrantStreamAggregatedResources"},
			_jsii_.MemberProperty{JsiiProperty: "mesh", GoGetter: "Mesh"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "virtualGatewayArn", GoGetter: "VirtualGatewayArn"},
			_jsii_.MemberProperty{JsiiProperty: "virtualGatewayName", GoGetter: "VirtualGatewayName"},
		},
		func() interface{} {
			j := jsiiProxy_IVirtualGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appmesh.IVirtualNode",
		reflect.TypeOf((*IVirtualNode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grantStreamAggregatedResources", GoMethod: "GrantStreamAggregatedResources"},
			_jsii_.MemberProperty{JsiiProperty: "mesh", GoGetter: "Mesh"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNodeArn", GoGetter: "VirtualNodeArn"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNodeName", GoGetter: "VirtualNodeName"},
		},
		func() interface{} {
			j := jsiiProxy_IVirtualNode{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appmesh.IVirtualRouter",
		reflect.TypeOf((*IVirtualRouter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "mesh", GoGetter: "Mesh"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouterArn", GoGetter: "VirtualRouterArn"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouterName", GoGetter: "VirtualRouterName"},
		},
		func() interface{} {
			j := jsiiProxy_IVirtualRouter{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appmesh.IVirtualService",
		reflect.TypeOf((*IVirtualService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "mesh", GoGetter: "Mesh"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceArn", GoGetter: "VirtualServiceArn"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceName", GoGetter: "VirtualServiceName"},
		},
		func() interface{} {
			j := jsiiProxy_IVirtualService{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appmesh.IpPreference",
		reflect.TypeOf((*IpPreference)(nil)).Elem(),
		map[string]interface{}{
			"IPV4_ONLY": IpPreference_IPV4_ONLY,
			"IPV4_PREFERRED": IpPreference_IPV4_PREFERRED,
			"IPV6_ONLY": IpPreference_IPV6_ONLY,
			"IPV6_PREFERRED": IpPreference_IPV6_PREFERRED,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.ListenerTlsOptions",
		reflect.TypeOf((*ListenerTlsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.Mesh",
		reflect.TypeOf((*Mesh)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addVirtualGateway", GoMethod: "AddVirtualGateway"},
			_jsii_.MemberMethod{JsiiMethod: "addVirtualNode", GoMethod: "AddVirtualNode"},
			_jsii_.MemberMethod{JsiiMethod: "addVirtualRouter", GoMethod: "AddVirtualRouter"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "meshArn", GoGetter: "MeshArn"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Mesh{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMesh)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appmesh.MeshFilterType",
		reflect.TypeOf((*MeshFilterType)(nil)).Elem(),
		map[string]interface{}{
			"ALLOW_ALL": MeshFilterType_ALLOW_ALL,
			"DROP_ALL": MeshFilterType_DROP_ALL,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.MeshProps",
		reflect.TypeOf((*MeshProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.MeshServiceDiscovery",
		reflect.TypeOf((*MeshServiceDiscovery)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.MutualTlsCertificate",
		reflect.TypeOf((*MutualTlsCertificate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "differentiator", GoGetter: "Differentiator"},
		},
		func() interface{} {
			j := jsiiProxy_MutualTlsCertificate{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TlsCertificate)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.MutualTlsValidation",
		reflect.TypeOf((*MutualTlsValidation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.MutualTlsValidationTrust",
		reflect.TypeOf((*MutualTlsValidationTrust)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "differentiator", GoGetter: "Differentiator"},
		},
		func() interface{} {
			j := jsiiProxy_MutualTlsValidationTrust{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TlsValidationTrust)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.OutlierDetection",
		reflect.TypeOf((*OutlierDetection)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.QueryParameterMatch",
		reflect.TypeOf((*QueryParameterMatch)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_QueryParameterMatch{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.QueryParameterMatchConfig",
		reflect.TypeOf((*QueryParameterMatchConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.Route",
		reflect.TypeOf((*Route)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "routeArn", GoGetter: "RouteArn"},
			_jsii_.MemberProperty{JsiiProperty: "routeName", GoGetter: "RouteName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouter", GoGetter: "VirtualRouter"},
		},
		func() interface{} {
			j := jsiiProxy_Route{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRoute)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.RouteAttributes",
		reflect.TypeOf((*RouteAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.RouteBaseProps",
		reflect.TypeOf((*RouteBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.RouteProps",
		reflect.TypeOf((*RouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.RouteSpec",
		reflect.TypeOf((*RouteSpec)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_RouteSpec{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.RouteSpecConfig",
		reflect.TypeOf((*RouteSpecConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.RouteSpecOptionsBase",
		reflect.TypeOf((*RouteSpecOptionsBase)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.ServiceDiscovery",
		reflect.TypeOf((*ServiceDiscovery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ServiceDiscovery{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.ServiceDiscoveryConfig",
		reflect.TypeOf((*ServiceDiscoveryConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.SubjectAlternativeNames",
		reflect.TypeOf((*SubjectAlternativeNames)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_SubjectAlternativeNames{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.SubjectAlternativeNamesMatcherConfig",
		reflect.TypeOf((*SubjectAlternativeNamesMatcherConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.TcpConnectionPool",
		reflect.TypeOf((*TcpConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.TcpHealthCheckOptions",
		reflect.TypeOf((*TcpHealthCheckOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appmesh.TcpRetryEvent",
		reflect.TypeOf((*TcpRetryEvent)(nil)).Elem(),
		map[string]interface{}{
			"CONNECTION_ERROR": TcpRetryEvent_CONNECTION_ERROR,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.TcpRouteSpecOptions",
		reflect.TypeOf((*TcpRouteSpecOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.TcpTimeout",
		reflect.TypeOf((*TcpTimeout)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.TcpVirtualNodeListenerOptions",
		reflect.TypeOf((*TcpVirtualNodeListenerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.TlsCertificate",
		reflect.TypeOf((*TlsCertificate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_TlsCertificate{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.TlsCertificateConfig",
		reflect.TypeOf((*TlsCertificateConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.TlsClientPolicy",
		reflect.TypeOf((*TlsClientPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appmesh.TlsMode",
		reflect.TypeOf((*TlsMode)(nil)).Elem(),
		map[string]interface{}{
			"STRICT": TlsMode_STRICT,
			"PERMISSIVE": TlsMode_PERMISSIVE,
			"DISABLED": TlsMode_DISABLED,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.TlsValidation",
		reflect.TypeOf((*TlsValidation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.TlsValidationTrust",
		reflect.TypeOf((*TlsValidationTrust)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_TlsValidationTrust{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.TlsValidationTrustConfig",
		reflect.TypeOf((*TlsValidationTrustConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.VirtualGateway",
		reflect.TypeOf((*VirtualGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addGatewayRoute", GoMethod: "AddGatewayRoute"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantStreamAggregatedResources", GoMethod: "GrantStreamAggregatedResources"},
			_jsii_.MemberProperty{JsiiProperty: "listeners", GoGetter: "Listeners"},
			_jsii_.MemberProperty{JsiiProperty: "mesh", GoGetter: "Mesh"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "virtualGatewayArn", GoGetter: "VirtualGatewayArn"},
			_jsii_.MemberProperty{JsiiProperty: "virtualGatewayName", GoGetter: "VirtualGatewayName"},
		},
		func() interface{} {
			j := jsiiProxy_VirtualGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVirtualGateway)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualGatewayAttributes",
		reflect.TypeOf((*VirtualGatewayAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualGatewayBaseProps",
		reflect.TypeOf((*VirtualGatewayBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.VirtualGatewayListener",
		reflect.TypeOf((*VirtualGatewayListener)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_VirtualGatewayListener{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualGatewayListenerConfig",
		reflect.TypeOf((*VirtualGatewayListenerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualGatewayProps",
		reflect.TypeOf((*VirtualGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.VirtualNode",
		reflect.TypeOf((*VirtualNode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBackend", GoMethod: "AddBackend"},
			_jsii_.MemberMethod{JsiiMethod: "addListener", GoMethod: "AddListener"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantStreamAggregatedResources", GoMethod: "GrantStreamAggregatedResources"},
			_jsii_.MemberProperty{JsiiProperty: "mesh", GoGetter: "Mesh"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNodeArn", GoGetter: "VirtualNodeArn"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNodeName", GoGetter: "VirtualNodeName"},
		},
		func() interface{} {
			j := jsiiProxy_VirtualNode{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVirtualNode)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualNodeAttributes",
		reflect.TypeOf((*VirtualNodeAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualNodeBaseProps",
		reflect.TypeOf((*VirtualNodeBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.VirtualNodeListener",
		reflect.TypeOf((*VirtualNodeListener)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_VirtualNodeListener{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualNodeListenerConfig",
		reflect.TypeOf((*VirtualNodeListenerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualNodeProps",
		reflect.TypeOf((*VirtualNodeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.VirtualRouter",
		reflect.TypeOf((*VirtualRouter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mesh", GoGetter: "Mesh"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouterArn", GoGetter: "VirtualRouterArn"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouterName", GoGetter: "VirtualRouterName"},
		},
		func() interface{} {
			j := jsiiProxy_VirtualRouter{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVirtualRouter)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualRouterAttributes",
		reflect.TypeOf((*VirtualRouterAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualRouterBaseProps",
		reflect.TypeOf((*VirtualRouterBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.VirtualRouterListener",
		reflect.TypeOf((*VirtualRouterListener)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_VirtualRouterListener{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualRouterListenerConfig",
		reflect.TypeOf((*VirtualRouterListenerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualRouterProps",
		reflect.TypeOf((*VirtualRouterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.VirtualService",
		reflect.TypeOf((*VirtualService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mesh", GoGetter: "Mesh"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceArn", GoGetter: "VirtualServiceArn"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceName", GoGetter: "VirtualServiceName"},
		},
		func() interface{} {
			j := jsiiProxy_VirtualService{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVirtualService)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualServiceAttributes",
		reflect.TypeOf((*VirtualServiceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualServiceBackendOptions",
		reflect.TypeOf((*VirtualServiceBackendOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualServiceProps",
		reflect.TypeOf((*VirtualServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.VirtualServiceProvider",
		reflect.TypeOf((*VirtualServiceProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_VirtualServiceProvider{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.VirtualServiceProviderConfig",
		reflect.TypeOf((*VirtualServiceProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.WeightedTarget",
		reflect.TypeOf((*WeightedTarget)(nil)).Elem(),
	)
}
