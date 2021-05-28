package awsappmesh

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute",
		reflect.TypeOf((*CfnGatewayRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
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
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
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
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.GrpcGatewayRouteProperty",
		reflect.TypeOf((*CfnGatewayRoute_GrpcGatewayRouteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.HttpGatewayRouteActionProperty",
		reflect.TypeOf((*CfnGatewayRoute_HttpGatewayRouteActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.HttpGatewayRouteMatchProperty",
		reflect.TypeOf((*CfnGatewayRoute_HttpGatewayRouteMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnGatewayRoute.HttpGatewayRouteProperty",
		reflect.TypeOf((*CfnGatewayRoute_HttpGatewayRouteProperty)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
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
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "routeName", GoGetter: "RouteName"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
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
		"aws-cdk-lib.aws_appmesh.CfnRoute.RouteSpecProperty",
		reflect.TypeOf((*CfnRoute_RouteSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appmesh.CfnRoute.TcpRouteActionProperty",
		reflect.TypeOf((*CfnRoute_TcpRouteActionProperty)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
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
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
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
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
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
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
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
}
