package awscloudfront

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.AccessLevel",
		reflect.TypeOf((*AccessLevel)(nil)).Elem(),
		map[string]interface{}{
			"READ": AccessLevel_READ,
			"READ_VERSIONED": AccessLevel_READ_VERSIONED,
			"LIST": AccessLevel_LIST,
			"WRITE": AccessLevel_WRITE,
			"DELETE": AccessLevel_DELETE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.AddBehaviorOptions",
		reflect.TypeOf((*AddBehaviorOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.AllowedMethods",
		reflect.TypeOf((*AllowedMethods)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "methods", GoGetter: "Methods"},
		},
		func() interface{} {
			return &jsiiProxy_AllowedMethods{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.AnycastIpListReference",
		reflect.TypeOf((*AnycastIpListReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.AssetImportSource",
		reflect.TypeOf((*AssetImportSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
		},
		func() interface{} {
			j := jsiiProxy_AssetImportSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ImportSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.Behavior",
		reflect.TypeOf((*Behavior)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.BehaviorOptions",
		reflect.TypeOf((*BehaviorOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CacheCookieBehavior",
		reflect.TypeOf((*CacheCookieBehavior)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "behavior", GoGetter: "Behavior"},
			_jsii_.MemberProperty{JsiiProperty: "cookies", GoGetter: "Cookies"},
		},
		func() interface{} {
			return &jsiiProxy_CacheCookieBehavior{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CacheHeaderBehavior",
		reflect.TypeOf((*CacheHeaderBehavior)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "behavior", GoGetter: "Behavior"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
		},
		func() interface{} {
			return &jsiiProxy_CacheHeaderBehavior{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CachePolicy",
		reflect.TypeOf((*CachePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyId", GoGetter: "CachePolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyRef", GoGetter: "CachePolicyRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CachePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICachePolicy)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CachePolicyProps",
		reflect.TypeOf((*CachePolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CachePolicyReference",
		reflect.TypeOf((*CachePolicyReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CacheQueryStringBehavior",
		reflect.TypeOf((*CacheQueryStringBehavior)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "behavior", GoGetter: "Behavior"},
			_jsii_.MemberProperty{JsiiProperty: "queryStrings", GoGetter: "QueryStrings"},
		},
		func() interface{} {
			return &jsiiProxy_CacheQueryStringBehavior{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CachedMethods",
		reflect.TypeOf((*CachedMethods)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "methods", GoGetter: "Methods"},
		},
		func() interface{} {
			return &jsiiProxy_CachedMethods{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnAnycastIpList",
		reflect.TypeOf((*CfnAnycastIpList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "anycastIpListRef", GoGetter: "AnycastIpListRef"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAnycastIpList", GoGetter: "AttrAnycastIpList"},
			_jsii_.MemberProperty{JsiiProperty: "attrETag", GoGetter: "AttrETag"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddressType", GoGetter: "IpAddressType"},
			_jsii_.MemberProperty{JsiiProperty: "ipCount", GoGetter: "IpCount"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAnycastIpList{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAnycastIpListRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnAnycastIpList.AnycastIpListProperty",
		reflect.TypeOf((*CfnAnycastIpList_AnycastIpListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnAnycastIpList.TagsProperty",
		reflect.TypeOf((*CfnAnycastIpList_TagsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnAnycastIpListProps",
		reflect.TypeOf((*CfnAnycastIpListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnCachePolicy",
		reflect.TypeOf((*CfnCachePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastModifiedTime", GoGetter: "AttrLastModifiedTime"},
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyConfig", GoGetter: "CachePolicyConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyRef", GoGetter: "CachePolicyRef"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCachePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICachePolicyRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnCachePolicy.CachePolicyConfigProperty",
		reflect.TypeOf((*CfnCachePolicy_CachePolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnCachePolicy.CookiesConfigProperty",
		reflect.TypeOf((*CfnCachePolicy_CookiesConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnCachePolicy.HeadersConfigProperty",
		reflect.TypeOf((*CfnCachePolicy_HeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnCachePolicy.ParametersInCacheKeyAndForwardedToOriginProperty",
		reflect.TypeOf((*CfnCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnCachePolicy.QueryStringsConfigProperty",
		reflect.TypeOf((*CfnCachePolicy_QueryStringsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnCachePolicyProps",
		reflect.TypeOf((*CfnCachePolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnCloudFrontOriginAccessIdentity",
		reflect.TypeOf((*CfnCloudFrontOriginAccessIdentity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrS3CanonicalUserId", GoGetter: "AttrS3CanonicalUserId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontOriginAccessIdentityConfig", GoGetter: "CloudFrontOriginAccessIdentityConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontOriginAccessIdentityRef", GoGetter: "CloudFrontOriginAccessIdentityRef"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudFrontOriginAccessIdentity{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICloudFrontOriginAccessIdentityRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnCloudFrontOriginAccessIdentity.CloudFrontOriginAccessIdentityConfigProperty",
		reflect.TypeOf((*CfnCloudFrontOriginAccessIdentity_CloudFrontOriginAccessIdentityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnCloudFrontOriginAccessIdentityProps",
		reflect.TypeOf((*CfnCloudFrontOriginAccessIdentityProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnConnectionGroup",
		reflect.TypeOf((*CfnConnectionGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "anycastIpListId", GoGetter: "AnycastIpListId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedTime", GoGetter: "AttrCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrETag", GoGetter: "AttrETag"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrIsDefault", GoGetter: "AttrIsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastModifiedTime", GoGetter: "AttrLastModifiedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrRoutingEndpoint", GoGetter: "AttrRoutingEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "connectionGroupRef", GoGetter: "ConnectionGroupRef"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Enabled", GoGetter: "Ipv6Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectionGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConnectionGroupRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnConnectionGroupProps",
		reflect.TypeOf((*CfnConnectionGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnContinuousDeploymentPolicy",
		reflect.TypeOf((*CfnContinuousDeploymentPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastModifiedTime", GoGetter: "AttrLastModifiedTime"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "continuousDeploymentPolicyConfig", GoGetter: "ContinuousDeploymentPolicyConfig"},
			_jsii_.MemberProperty{JsiiProperty: "continuousDeploymentPolicyRef", GoGetter: "ContinuousDeploymentPolicyRef"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnContinuousDeploymentPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IContinuousDeploymentPolicyRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnContinuousDeploymentPolicy.ContinuousDeploymentPolicyConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnContinuousDeploymentPolicy.SessionStickinessConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicy_SessionStickinessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnContinuousDeploymentPolicy.SingleHeaderConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicy_SingleHeaderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnContinuousDeploymentPolicy.SingleHeaderPolicyConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicy_SingleHeaderPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnContinuousDeploymentPolicy.SingleWeightConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicy_SingleWeightConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnContinuousDeploymentPolicy.SingleWeightPolicyConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicy_SingleWeightPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnContinuousDeploymentPolicy.TrafficConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicy_TrafficConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnContinuousDeploymentPolicyProps",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution",
		reflect.TypeOf((*CfnDistribution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrDomainName", GoGetter: "AttrDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "distributionConfig", GoGetter: "DistributionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "distributionRef", GoGetter: "DistributionRef"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDistribution{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDistributionRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.CacheBehaviorProperty",
		reflect.TypeOf((*CfnDistribution_CacheBehaviorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.CookiesProperty",
		reflect.TypeOf((*CfnDistribution_CookiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.CustomErrorResponseProperty",
		reflect.TypeOf((*CfnDistribution_CustomErrorResponseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.CustomOriginConfigProperty",
		reflect.TypeOf((*CfnDistribution_CustomOriginConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.DefaultCacheBehaviorProperty",
		reflect.TypeOf((*CfnDistribution_DefaultCacheBehaviorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.DefinitionProperty",
		reflect.TypeOf((*CfnDistribution_DefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.DistributionConfigProperty",
		reflect.TypeOf((*CfnDistribution_DistributionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.ForwardedValuesProperty",
		reflect.TypeOf((*CfnDistribution_ForwardedValuesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.FunctionAssociationProperty",
		reflect.TypeOf((*CfnDistribution_FunctionAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.GeoRestrictionProperty",
		reflect.TypeOf((*CfnDistribution_GeoRestrictionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.GrpcConfigProperty",
		reflect.TypeOf((*CfnDistribution_GrpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.LambdaFunctionAssociationProperty",
		reflect.TypeOf((*CfnDistribution_LambdaFunctionAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.LegacyCustomOriginProperty",
		reflect.TypeOf((*CfnDistribution_LegacyCustomOriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.LegacyS3OriginProperty",
		reflect.TypeOf((*CfnDistribution_LegacyS3OriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.LoggingProperty",
		reflect.TypeOf((*CfnDistribution_LoggingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.OriginCustomHeaderProperty",
		reflect.TypeOf((*CfnDistribution_OriginCustomHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.OriginGroupFailoverCriteriaProperty",
		reflect.TypeOf((*CfnDistribution_OriginGroupFailoverCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.OriginGroupMemberProperty",
		reflect.TypeOf((*CfnDistribution_OriginGroupMemberProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.OriginGroupMembersProperty",
		reflect.TypeOf((*CfnDistribution_OriginGroupMembersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.OriginGroupProperty",
		reflect.TypeOf((*CfnDistribution_OriginGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.OriginGroupsProperty",
		reflect.TypeOf((*CfnDistribution_OriginGroupsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.OriginProperty",
		reflect.TypeOf((*CfnDistribution_OriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.OriginShieldProperty",
		reflect.TypeOf((*CfnDistribution_OriginShieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.ParameterDefinitionProperty",
		reflect.TypeOf((*CfnDistribution_ParameterDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.RestrictionsProperty",
		reflect.TypeOf((*CfnDistribution_RestrictionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.S3OriginConfigProperty",
		reflect.TypeOf((*CfnDistribution_S3OriginConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.StatusCodesProperty",
		reflect.TypeOf((*CfnDistribution_StatusCodesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.StringSchemaProperty",
		reflect.TypeOf((*CfnDistribution_StringSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.TenantConfigProperty",
		reflect.TypeOf((*CfnDistribution_TenantConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.ViewerCertificateProperty",
		reflect.TypeOf((*CfnDistribution_ViewerCertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution.VpcOriginConfigProperty",
		reflect.TypeOf((*CfnDistribution_VpcOriginConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistributionProps",
		reflect.TypeOf((*CfnDistributionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnDistributionTenant",
		reflect.TypeOf((*CfnDistributionTenant)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedTime", GoGetter: "AttrCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrDomainResults", GoGetter: "AttrDomainResults"},
			_jsii_.MemberProperty{JsiiProperty: "attrETag", GoGetter: "AttrETag"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastModifiedTime", GoGetter: "AttrLastModifiedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "connectionGroupId", GoGetter: "ConnectionGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customizations", GoGetter: "Customizations"},
			_jsii_.MemberProperty{JsiiProperty: "distributionId", GoGetter: "DistributionId"},
			_jsii_.MemberProperty{JsiiProperty: "distributionTenantRef", GoGetter: "DistributionTenantRef"},
			_jsii_.MemberProperty{JsiiProperty: "domains", GoGetter: "Domains"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "managedCertificateRequest", GoGetter: "ManagedCertificateRequest"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDistributionTenant{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDistributionTenantRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistributionTenant.CertificateProperty",
		reflect.TypeOf((*CfnDistributionTenant_CertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistributionTenant.CustomizationsProperty",
		reflect.TypeOf((*CfnDistributionTenant_CustomizationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistributionTenant.DomainResultProperty",
		reflect.TypeOf((*CfnDistributionTenant_DomainResultProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistributionTenant.GeoRestrictionCustomizationProperty",
		reflect.TypeOf((*CfnDistributionTenant_GeoRestrictionCustomizationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistributionTenant.ManagedCertificateRequestProperty",
		reflect.TypeOf((*CfnDistributionTenant_ManagedCertificateRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistributionTenant.ParameterProperty",
		reflect.TypeOf((*CfnDistributionTenant_ParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistributionTenant.WebAclCustomizationProperty",
		reflect.TypeOf((*CfnDistributionTenant_WebAclCustomizationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnDistributionTenantProps",
		reflect.TypeOf((*CfnDistributionTenantProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnFunction",
		reflect.TypeOf((*CfnFunction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrFunctionArn", GoGetter: "AttrFunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrFunctionMetadataFunctionArn", GoGetter: "AttrFunctionMetadataFunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrStage", GoGetter: "AttrStage"},
			_jsii_.MemberProperty{JsiiProperty: "autoPublish", GoGetter: "AutoPublish"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "functionCode", GoGetter: "FunctionCode"},
			_jsii_.MemberProperty{JsiiProperty: "functionConfig", GoGetter: "FunctionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "functionMetadata", GoGetter: "FunctionMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "functionRef", GoGetter: "FunctionRef"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFunction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFunctionRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnFunction.FunctionConfigProperty",
		reflect.TypeOf((*CfnFunction_FunctionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnFunction.FunctionMetadataProperty",
		reflect.TypeOf((*CfnFunction_FunctionMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnFunction.KeyValueStoreAssociationProperty",
		reflect.TypeOf((*CfnFunction_KeyValueStoreAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnFunctionProps",
		reflect.TypeOf((*CfnFunctionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnKeyGroup",
		reflect.TypeOf((*CfnKeyGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastModifiedTime", GoGetter: "AttrLastModifiedTime"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "keyGroupConfig", GoGetter: "KeyGroupConfig"},
			_jsii_.MemberProperty{JsiiProperty: "keyGroupRef", GoGetter: "KeyGroupRef"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnKeyGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IKeyGroupRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnKeyGroup.KeyGroupConfigProperty",
		reflect.TypeOf((*CfnKeyGroup_KeyGroupConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnKeyGroupProps",
		reflect.TypeOf((*CfnKeyGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnKeyValueStore",
		reflect.TypeOf((*CfnKeyValueStore)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "importSource", GoGetter: "ImportSource"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "keyValueStoreRef", GoGetter: "KeyValueStoreRef"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnKeyValueStore{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IKeyValueStoreRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnKeyValueStore.ImportSourceProperty",
		reflect.TypeOf((*CfnKeyValueStore_ImportSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnKeyValueStoreProps",
		reflect.TypeOf((*CfnKeyValueStoreProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnMonitoringSubscription",
		reflect.TypeOf((*CfnMonitoringSubscription)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "distributionId", GoGetter: "DistributionId"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringSubscription", GoGetter: "MonitoringSubscription"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringSubscriptionRef", GoGetter: "MonitoringSubscriptionRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMonitoringSubscription{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMonitoringSubscriptionRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnMonitoringSubscription.MonitoringSubscriptionProperty",
		reflect.TypeOf((*CfnMonitoringSubscription_MonitoringSubscriptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnMonitoringSubscription.RealtimeMetricsSubscriptionConfigProperty",
		reflect.TypeOf((*CfnMonitoringSubscription_RealtimeMetricsSubscriptionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnMonitoringSubscriptionProps",
		reflect.TypeOf((*CfnMonitoringSubscriptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnOriginAccessControl",
		reflect.TypeOf((*CfnOriginAccessControl)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessControlConfig", GoGetter: "OriginAccessControlConfig"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessControlRef", GoGetter: "OriginAccessControlRef"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOriginAccessControl{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOriginAccessControlRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnOriginAccessControl.OriginAccessControlConfigProperty",
		reflect.TypeOf((*CfnOriginAccessControl_OriginAccessControlConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnOriginAccessControlProps",
		reflect.TypeOf((*CfnOriginAccessControlProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnOriginRequestPolicy",
		reflect.TypeOf((*CfnOriginRequestPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastModifiedTime", GoGetter: "AttrLastModifiedTime"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyConfig", GoGetter: "OriginRequestPolicyConfig"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyRef", GoGetter: "OriginRequestPolicyRef"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOriginRequestPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOriginRequestPolicyRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnOriginRequestPolicy.CookiesConfigProperty",
		reflect.TypeOf((*CfnOriginRequestPolicy_CookiesConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnOriginRequestPolicy.HeadersConfigProperty",
		reflect.TypeOf((*CfnOriginRequestPolicy_HeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnOriginRequestPolicy.OriginRequestPolicyConfigProperty",
		reflect.TypeOf((*CfnOriginRequestPolicy_OriginRequestPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnOriginRequestPolicy.QueryStringsConfigProperty",
		reflect.TypeOf((*CfnOriginRequestPolicy_QueryStringsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnOriginRequestPolicyProps",
		reflect.TypeOf((*CfnOriginRequestPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnPublicKey",
		reflect.TypeOf((*CfnPublicKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedTime", GoGetter: "AttrCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "publicKeyConfig", GoGetter: "PublicKeyConfig"},
			_jsii_.MemberProperty{JsiiProperty: "publicKeyRef", GoGetter: "PublicKeyRef"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPublicKey{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPublicKeyRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnPublicKey.PublicKeyConfigProperty",
		reflect.TypeOf((*CfnPublicKey_PublicKeyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnPublicKeyProps",
		reflect.TypeOf((*CfnPublicKeyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnRealtimeLogConfig",
		reflect.TypeOf((*CfnRealtimeLogConfig)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endPoints", GoGetter: "EndPoints"},
			_jsii_.MemberProperty{JsiiProperty: "fields", GoGetter: "Fields"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigRef", GoGetter: "RealtimeLogConfigRef"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "samplingRate", GoGetter: "SamplingRate"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRealtimeLogConfig{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRealtimeLogConfigRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnRealtimeLogConfig.EndPointProperty",
		reflect.TypeOf((*CfnRealtimeLogConfig_EndPointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnRealtimeLogConfig.KinesisStreamConfigProperty",
		reflect.TypeOf((*CfnRealtimeLogConfig_KinesisStreamConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnRealtimeLogConfigProps",
		reflect.TypeOf((*CfnRealtimeLogConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy",
		reflect.TypeOf((*CfnResponseHeadersPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastModifiedTime", GoGetter: "AttrLastModifiedTime"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyConfig", GoGetter: "ResponseHeadersPolicyConfig"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyRef", GoGetter: "ResponseHeadersPolicyRef"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResponseHeadersPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResponseHeadersPolicyRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.AccessControlAllowHeadersProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_AccessControlAllowHeadersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.AccessControlAllowMethodsProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_AccessControlAllowMethodsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.AccessControlAllowOriginsProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_AccessControlAllowOriginsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.AccessControlExposeHeadersProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_AccessControlExposeHeadersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.ContentSecurityPolicyProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_ContentSecurityPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.ContentTypeOptionsProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_ContentTypeOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.CorsConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_CorsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.CustomHeaderProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_CustomHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.CustomHeadersConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_CustomHeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.FrameOptionsProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_FrameOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.ReferrerPolicyProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_ReferrerPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.RemoveHeaderProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_RemoveHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.RemoveHeadersConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_RemoveHeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.ResponseHeadersPolicyConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_ResponseHeadersPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.SecurityHeadersConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_SecurityHeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.ServerTimingHeadersConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_ServerTimingHeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.StrictTransportSecurityProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_StrictTransportSecurityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy.XSSProtectionProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicy_XSSProtectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicyProps",
		reflect.TypeOf((*CfnResponseHeadersPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnStreamingDistribution",
		reflect.TypeOf((*CfnStreamingDistribution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrDomainName", GoGetter: "AttrDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "streamingDistributionConfig", GoGetter: "StreamingDistributionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "streamingDistributionRef", GoGetter: "StreamingDistributionRef"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStreamingDistribution{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IStreamingDistributionRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnStreamingDistribution.LoggingProperty",
		reflect.TypeOf((*CfnStreamingDistribution_LoggingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnStreamingDistribution.S3OriginProperty",
		reflect.TypeOf((*CfnStreamingDistribution_S3OriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnStreamingDistribution.StreamingDistributionConfigProperty",
		reflect.TypeOf((*CfnStreamingDistribution_StreamingDistributionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnStreamingDistribution.TrustedSignersProperty",
		reflect.TypeOf((*CfnStreamingDistribution_TrustedSignersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnStreamingDistributionProps",
		reflect.TypeOf((*CfnStreamingDistributionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CfnVpcOrigin",
		reflect.TypeOf((*CfnVpcOrigin)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedTime", GoGetter: "AttrCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastModifiedTime", GoGetter: "AttrLastModifiedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcOriginEndpointConfig", GoGetter: "VpcOriginEndpointConfig"},
			_jsii_.MemberProperty{JsiiProperty: "vpcOriginRef", GoGetter: "VpcOriginRef"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVpcOrigin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVpcOriginRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnVpcOrigin.VpcOriginEndpointConfigProperty",
		reflect.TypeOf((*CfnVpcOrigin_VpcOriginEndpointConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CfnVpcOriginProps",
		reflect.TypeOf((*CfnVpcOriginProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.CloudFrontAllowedCachedMethods",
		reflect.TypeOf((*CloudFrontAllowedCachedMethods)(nil)).Elem(),
		map[string]interface{}{
			"GET_HEAD": CloudFrontAllowedCachedMethods_GET_HEAD,
			"GET_HEAD_OPTIONS": CloudFrontAllowedCachedMethods_GET_HEAD_OPTIONS,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.CloudFrontAllowedMethods",
		reflect.TypeOf((*CloudFrontAllowedMethods)(nil)).Elem(),
		map[string]interface{}{
			"GET_HEAD": CloudFrontAllowedMethods_GET_HEAD,
			"GET_HEAD_OPTIONS": CloudFrontAllowedMethods_GET_HEAD_OPTIONS,
			"ALL": CloudFrontAllowedMethods_ALL,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CloudFrontOriginAccessIdentityReference",
		reflect.TypeOf((*CloudFrontOriginAccessIdentityReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.CloudFrontWebDistribution",
		reflect.TypeOf((*CloudFrontWebDistribution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "distributionArn", GoGetter: "DistributionArn"},
			_jsii_.MemberProperty{JsiiProperty: "distributionDomainName", GoGetter: "DistributionDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "distributionId", GoGetter: "DistributionId"},
			_jsii_.MemberProperty{JsiiProperty: "distributionRef", GoGetter: "DistributionRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantCreateInvalidation", GoMethod: "GrantCreateInvalidation"},
			_jsii_.MemberProperty{JsiiProperty: "loggingBucket", GoGetter: "LoggingBucket"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudFrontWebDistribution{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDistribution)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CloudFrontWebDistributionAttributes",
		reflect.TypeOf((*CloudFrontWebDistributionAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CloudFrontWebDistributionProps",
		reflect.TypeOf((*CloudFrontWebDistributionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ConnectionGroupReference",
		reflect.TypeOf((*ConnectionGroupReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ContinuousDeploymentPolicyReference",
		reflect.TypeOf((*ContinuousDeploymentPolicyReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.CustomOriginConfig",
		reflect.TypeOf((*CustomOriginConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.Distribution",
		reflect.TypeOf((*Distribution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBehavior", GoMethod: "AddBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "attachWebAclId", GoMethod: "AttachWebAclId"},
			_jsii_.MemberProperty{JsiiProperty: "distributionArn", GoGetter: "DistributionArn"},
			_jsii_.MemberProperty{JsiiProperty: "distributionDomainName", GoGetter: "DistributionDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "distributionId", GoGetter: "DistributionId"},
			_jsii_.MemberProperty{JsiiProperty: "distributionRef", GoGetter: "DistributionRef"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantCreateInvalidation", GoMethod: "GrantCreateInvalidation"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metric401ErrorRate", GoMethod: "Metric401ErrorRate"},
			_jsii_.MemberMethod{JsiiMethod: "metric403ErrorRate", GoMethod: "Metric403ErrorRate"},
			_jsii_.MemberMethod{JsiiMethod: "metric404ErrorRate", GoMethod: "Metric404ErrorRate"},
			_jsii_.MemberMethod{JsiiMethod: "metric4xxErrorRate", GoMethod: "Metric4xxErrorRate"},
			_jsii_.MemberMethod{JsiiMethod: "metric502ErrorRate", GoMethod: "Metric502ErrorRate"},
			_jsii_.MemberMethod{JsiiMethod: "metric503ErrorRate", GoMethod: "Metric503ErrorRate"},
			_jsii_.MemberMethod{JsiiMethod: "metric504ErrorRate", GoMethod: "Metric504ErrorRate"},
			_jsii_.MemberMethod{JsiiMethod: "metric5xxErrorRate", GoMethod: "Metric5xxErrorRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricBytesDownloaded", GoMethod: "MetricBytesDownloaded"},
			_jsii_.MemberMethod{JsiiMethod: "metricBytesUploaded", GoMethod: "MetricBytesUploaded"},
			_jsii_.MemberMethod{JsiiMethod: "metricCacheHitRate", GoMethod: "MetricCacheHitRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricOriginLatency", GoMethod: "MetricOriginLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricRequests", GoMethod: "MetricRequests"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalErrorRate", GoMethod: "MetricTotalErrorRate"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Distribution{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDistribution)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.DistributionAttributes",
		reflect.TypeOf((*DistributionAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.DistributionProps",
		reflect.TypeOf((*DistributionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.DistributionReference",
		reflect.TypeOf((*DistributionReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.DistributionTenantReference",
		reflect.TypeOf((*DistributionTenantReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.EdgeLambda",
		reflect.TypeOf((*EdgeLambda)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.Endpoint",
		reflect.TypeOf((*Endpoint)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Endpoint{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ErrorResponse",
		reflect.TypeOf((*ErrorResponse)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.FailoverStatusCode",
		reflect.TypeOf((*FailoverStatusCode)(nil)).Elem(),
		map[string]interface{}{
			"FORBIDDEN": FailoverStatusCode_FORBIDDEN,
			"NOT_FOUND": FailoverStatusCode_NOT_FOUND,
			"INTERNAL_SERVER_ERROR": FailoverStatusCode_INTERNAL_SERVER_ERROR,
			"BAD_GATEWAY": FailoverStatusCode_BAD_GATEWAY,
			"SERVICE_UNAVAILABLE": FailoverStatusCode_SERVICE_UNAVAILABLE,
			"GATEWAY_TIMEOUT": FailoverStatusCode_GATEWAY_TIMEOUT,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.FileCodeOptions",
		reflect.TypeOf((*FileCodeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.Function",
		reflect.TypeOf((*Function)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "functionArn", GoGetter: "FunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
			_jsii_.MemberProperty{JsiiProperty: "functionRef", GoGetter: "FunctionRef"},
			_jsii_.MemberProperty{JsiiProperty: "functionRuntime", GoGetter: "FunctionRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "functionStage", GoGetter: "FunctionStage"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Function{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFunction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.FunctionAssociation",
		reflect.TypeOf((*FunctionAssociation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.FunctionAttributes",
		reflect.TypeOf((*FunctionAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.FunctionCode",
		reflect.TypeOf((*FunctionCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
		},
		func() interface{} {
			return &jsiiProxy_FunctionCode{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.FunctionEventType",
		reflect.TypeOf((*FunctionEventType)(nil)).Elem(),
		map[string]interface{}{
			"VIEWER_REQUEST": FunctionEventType_VIEWER_REQUEST,
			"VIEWER_RESPONSE": FunctionEventType_VIEWER_RESPONSE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.FunctionProps",
		reflect.TypeOf((*FunctionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.FunctionReference",
		reflect.TypeOf((*FunctionReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.FunctionRuntime",
		reflect.TypeOf((*FunctionRuntime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_FunctionRuntime{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.FunctionUrlOriginAccessControl",
		reflect.TypeOf((*FunctionUrlOriginAccessControl)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessControlId", GoGetter: "OriginAccessControlId"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FunctionUrlOriginAccessControl{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOriginAccessControl)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.FunctionUrlOriginAccessControlProps",
		reflect.TypeOf((*FunctionUrlOriginAccessControlProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.GeoRestriction",
		reflect.TypeOf((*GeoRestriction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "locations", GoGetter: "Locations"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionType", GoGetter: "RestrictionType"},
		},
		func() interface{} {
			return &jsiiProxy_GeoRestriction{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.HeadersFrameOption",
		reflect.TypeOf((*HeadersFrameOption)(nil)).Elem(),
		map[string]interface{}{
			"DENY": HeadersFrameOption_DENY,
			"SAMEORIGIN": HeadersFrameOption_SAMEORIGIN,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.HeadersReferrerPolicy",
		reflect.TypeOf((*HeadersReferrerPolicy)(nil)).Elem(),
		map[string]interface{}{
			"NO_REFERRER": HeadersReferrerPolicy_NO_REFERRER,
			"NO_REFERRER_WHEN_DOWNGRADE": HeadersReferrerPolicy_NO_REFERRER_WHEN_DOWNGRADE,
			"ORIGIN": HeadersReferrerPolicy_ORIGIN,
			"ORIGIN_WHEN_CROSS_ORIGIN": HeadersReferrerPolicy_ORIGIN_WHEN_CROSS_ORIGIN,
			"SAME_ORIGIN": HeadersReferrerPolicy_SAME_ORIGIN,
			"STRICT_ORIGIN": HeadersReferrerPolicy_STRICT_ORIGIN,
			"STRICT_ORIGIN_WHEN_CROSS_ORIGIN": HeadersReferrerPolicy_STRICT_ORIGIN_WHEN_CROSS_ORIGIN,
			"UNSAFE_URL": HeadersReferrerPolicy_UNSAFE_URL,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.HttpVersion",
		reflect.TypeOf((*HttpVersion)(nil)).Elem(),
		map[string]interface{}{
			"HTTP1_1": HttpVersion_HTTP1_1,
			"HTTP2": HttpVersion_HTTP2,
			"HTTP2_AND_3": HttpVersion_HTTP2_AND_3,
			"HTTP3": HttpVersion_HTTP3,
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IAnycastIpListRef",
		reflect.TypeOf((*IAnycastIpListRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "anycastIpListRef", GoGetter: "AnycastIpListRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IAnycastIpListRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.ICachePolicy",
		reflect.TypeOf((*ICachePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyId", GoGetter: "CachePolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyRef", GoGetter: "CachePolicyRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_ICachePolicy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICachePolicyRef)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.ICachePolicyRef",
		reflect.TypeOf((*ICachePolicyRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyRef", GoGetter: "CachePolicyRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_ICachePolicyRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.ICloudFrontOriginAccessIdentityRef",
		reflect.TypeOf((*ICloudFrontOriginAccessIdentityRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontOriginAccessIdentityRef", GoGetter: "CloudFrontOriginAccessIdentityRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_ICloudFrontOriginAccessIdentityRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IConnectionGroupRef",
		reflect.TypeOf((*IConnectionGroupRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connectionGroupRef", GoGetter: "ConnectionGroupRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IConnectionGroupRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IContinuousDeploymentPolicyRef",
		reflect.TypeOf((*IContinuousDeploymentPolicyRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "continuousDeploymentPolicyRef", GoGetter: "ContinuousDeploymentPolicyRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IContinuousDeploymentPolicyRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IDistribution",
		reflect.TypeOf((*IDistribution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "distributionArn", GoGetter: "DistributionArn"},
			_jsii_.MemberProperty{JsiiProperty: "distributionDomainName", GoGetter: "DistributionDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "distributionId", GoGetter: "DistributionId"},
			_jsii_.MemberProperty{JsiiProperty: "distributionRef", GoGetter: "DistributionRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantCreateInvalidation", GoMethod: "GrantCreateInvalidation"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDistribution{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDistributionRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IDistributionRef",
		reflect.TypeOf((*IDistributionRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "distributionRef", GoGetter: "DistributionRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IDistributionRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IDistributionTenantRef",
		reflect.TypeOf((*IDistributionTenantRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "distributionTenantRef", GoGetter: "DistributionTenantRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IDistributionTenantRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IFunction",
		reflect.TypeOf((*IFunction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "functionArn", GoGetter: "FunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
			_jsii_.MemberProperty{JsiiProperty: "functionRef", GoGetter: "FunctionRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IFunction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFunctionRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IFunctionRef",
		reflect.TypeOf((*IFunctionRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "functionRef", GoGetter: "FunctionRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IFunctionRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IKeyGroup",
		reflect.TypeOf((*IKeyGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "keyGroupId", GoGetter: "KeyGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "keyGroupRef", GoGetter: "KeyGroupRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IKeyGroup{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IKeyGroupRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IKeyGroupRef",
		reflect.TypeOf((*IKeyGroupRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "keyGroupRef", GoGetter: "KeyGroupRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IKeyGroupRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IKeyValueStore",
		reflect.TypeOf((*IKeyValueStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "keyValueStoreArn", GoGetter: "KeyValueStoreArn"},
			_jsii_.MemberProperty{JsiiProperty: "keyValueStoreId", GoGetter: "KeyValueStoreId"},
			_jsii_.MemberProperty{JsiiProperty: "keyValueStoreRef", GoGetter: "KeyValueStoreRef"},
			_jsii_.MemberProperty{JsiiProperty: "keyValueStoreStatus", GoGetter: "KeyValueStoreStatus"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IKeyValueStore{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IKeyValueStoreRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IKeyValueStoreRef",
		reflect.TypeOf((*IKeyValueStoreRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "keyValueStoreRef", GoGetter: "KeyValueStoreRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IKeyValueStoreRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IMonitoringSubscriptionRef",
		reflect.TypeOf((*IMonitoringSubscriptionRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "monitoringSubscriptionRef", GoGetter: "MonitoringSubscriptionRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IMonitoringSubscriptionRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IOrigin",
		reflect.TypeOf((*IOrigin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IOrigin{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IOriginAccessControl",
		reflect.TypeOf((*IOriginAccessControl)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessControlId", GoGetter: "OriginAccessControlId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IOriginAccessControl{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IOriginAccessControlRef",
		reflect.TypeOf((*IOriginAccessControlRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessControlRef", GoGetter: "OriginAccessControlRef"},
		},
		func() interface{} {
			j := jsiiProxy_IOriginAccessControlRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IOriginAccessIdentity",
		reflect.TypeOf((*IOriginAccessIdentity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontOriginAccessIdentityRef", GoGetter: "CloudFrontOriginAccessIdentityRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessIdentityId", GoGetter: "OriginAccessIdentityId"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessIdentityName", GoGetter: "OriginAccessIdentityName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IOriginAccessIdentity{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICloudFrontOriginAccessIdentityRef)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IOriginRequestPolicy",
		reflect.TypeOf((*IOriginRequestPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyId", GoGetter: "OriginRequestPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyRef", GoGetter: "OriginRequestPolicyRef"},
		},
		func() interface{} {
			j := jsiiProxy_IOriginRequestPolicy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOriginRequestPolicyRef)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IOriginRequestPolicyRef",
		reflect.TypeOf((*IOriginRequestPolicyRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyRef", GoGetter: "OriginRequestPolicyRef"},
		},
		func() interface{} {
			j := jsiiProxy_IOriginRequestPolicyRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IPublicKey",
		reflect.TypeOf((*IPublicKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "publicKeyId", GoGetter: "PublicKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "publicKeyRef", GoGetter: "PublicKeyRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IPublicKey{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPublicKeyRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IPublicKeyRef",
		reflect.TypeOf((*IPublicKeyRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "publicKeyRef", GoGetter: "PublicKeyRef"},
		},
		func() interface{} {
			j := jsiiProxy_IPublicKeyRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IRealtimeLogConfig",
		reflect.TypeOf((*IRealtimeLogConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigArn", GoGetter: "RealtimeLogConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigName", GoGetter: "RealtimeLogConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigRef", GoGetter: "RealtimeLogConfigRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IRealtimeLogConfig{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRealtimeLogConfigRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IRealtimeLogConfigRef",
		reflect.TypeOf((*IRealtimeLogConfigRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigRef", GoGetter: "RealtimeLogConfigRef"},
		},
		func() interface{} {
			j := jsiiProxy_IRealtimeLogConfigRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IResponseHeadersPolicy",
		reflect.TypeOf((*IResponseHeadersPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyId", GoGetter: "ResponseHeadersPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyRef", GoGetter: "ResponseHeadersPolicyRef"},
		},
		func() interface{} {
			j := jsiiProxy_IResponseHeadersPolicy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResponseHeadersPolicyRef)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IResponseHeadersPolicyRef",
		reflect.TypeOf((*IResponseHeadersPolicyRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyRef", GoGetter: "ResponseHeadersPolicyRef"},
		},
		func() interface{} {
			j := jsiiProxy_IResponseHeadersPolicyRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IStreamingDistributionRef",
		reflect.TypeOf((*IStreamingDistributionRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "streamingDistributionRef", GoGetter: "StreamingDistributionRef"},
		},
		func() interface{} {
			j := jsiiProxy_IStreamingDistributionRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IVpcOrigin",
		reflect.TypeOf((*IVpcOrigin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vpcOriginArn", GoGetter: "VpcOriginArn"},
			_jsii_.MemberProperty{JsiiProperty: "vpcOriginId", GoGetter: "VpcOriginId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcOriginRef", GoGetter: "VpcOriginRef"},
		},
		func() interface{} {
			j := jsiiProxy_IVpcOrigin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVpcOriginRef)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudfront.IVpcOriginRef",
		reflect.TypeOf((*IVpcOriginRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "vpcOriginRef", GoGetter: "VpcOriginRef"},
		},
		func() interface{} {
			j := jsiiProxy_IVpcOriginRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.ImportSource",
		reflect.TypeOf((*ImportSource)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ImportSource{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.InlineImportSource",
		reflect.TypeOf((*InlineImportSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "data", GoGetter: "Data"},
		},
		func() interface{} {
			j := jsiiProxy_InlineImportSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ImportSource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.KeyGroup",
		reflect.TypeOf((*KeyGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyGroupId", GoGetter: "KeyGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "keyGroupRef", GoGetter: "KeyGroupRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KeyGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IKeyGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.KeyGroupProps",
		reflect.TypeOf((*KeyGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.KeyGroupReference",
		reflect.TypeOf((*KeyGroupReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.KeyValueStore",
		reflect.TypeOf((*KeyValueStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyValueStoreArn", GoGetter: "KeyValueStoreArn"},
			_jsii_.MemberProperty{JsiiProperty: "keyValueStoreId", GoGetter: "KeyValueStoreId"},
			_jsii_.MemberProperty{JsiiProperty: "keyValueStoreRef", GoGetter: "KeyValueStoreRef"},
			_jsii_.MemberProperty{JsiiProperty: "keyValueStoreStatus", GoGetter: "KeyValueStoreStatus"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KeyValueStore{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IKeyValueStore)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.KeyValueStoreProps",
		reflect.TypeOf((*KeyValueStoreProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.KeyValueStoreReference",
		reflect.TypeOf((*KeyValueStoreReference)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.LambdaEdgeEventType",
		reflect.TypeOf((*LambdaEdgeEventType)(nil)).Elem(),
		map[string]interface{}{
			"ORIGIN_REQUEST": LambdaEdgeEventType_ORIGIN_REQUEST,
			"ORIGIN_RESPONSE": LambdaEdgeEventType_ORIGIN_RESPONSE,
			"VIEWER_REQUEST": LambdaEdgeEventType_VIEWER_REQUEST,
			"VIEWER_RESPONSE": LambdaEdgeEventType_VIEWER_RESPONSE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.LambdaFunctionAssociation",
		reflect.TypeOf((*LambdaFunctionAssociation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.LoggingConfiguration",
		reflect.TypeOf((*LoggingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.MonitoringSubscriptionReference",
		reflect.TypeOf((*MonitoringSubscriptionReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.OriginAccessControlBaseProps",
		reflect.TypeOf((*OriginAccessControlBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.OriginAccessControlOriginType",
		reflect.TypeOf((*OriginAccessControlOriginType)(nil)).Elem(),
		map[string]interface{}{
			"S3": OriginAccessControlOriginType_S3,
			"LAMBDA": OriginAccessControlOriginType_LAMBDA,
			"MEDIASTORE": OriginAccessControlOriginType_MEDIASTORE,
			"MEDIAPACKAGEV2": OriginAccessControlOriginType_MEDIAPACKAGEV2,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.OriginAccessControlReference",
		reflect.TypeOf((*OriginAccessControlReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.OriginAccessIdentity",
		reflect.TypeOf((*OriginAccessIdentity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "arn", GoMethod: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontOriginAccessIdentityRef", GoGetter: "CloudFrontOriginAccessIdentityRef"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontOriginAccessIdentityS3CanonicalUserId", GoGetter: "CloudFrontOriginAccessIdentityS3CanonicalUserId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessIdentityId", GoGetter: "OriginAccessIdentityId"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessIdentityName", GoGetter: "OriginAccessIdentityName"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OriginAccessIdentity{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOriginAccessIdentity)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.OriginAccessIdentityProps",
		reflect.TypeOf((*OriginAccessIdentityProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.OriginBase",
		reflect.TypeOf((*OriginBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "renderCustomOriginConfig", GoMethod: "RenderCustomOriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "renderS3OriginConfig", GoMethod: "RenderS3OriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "renderVpcOriginConfig", GoMethod: "RenderVpcOriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "validateResponseCompletionTimeoutWithReadTimeout", GoMethod: "ValidateResponseCompletionTimeoutWithReadTimeout"},
		},
		func() interface{} {
			j := jsiiProxy_OriginBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOrigin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.OriginBindConfig",
		reflect.TypeOf((*OriginBindConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.OriginBindOptions",
		reflect.TypeOf((*OriginBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.OriginFailoverConfig",
		reflect.TypeOf((*OriginFailoverConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.OriginIpAddressType",
		reflect.TypeOf((*OriginIpAddressType)(nil)).Elem(),
		map[string]interface{}{
			"IPV4": OriginIpAddressType_IPV4,
			"IPV6": OriginIpAddressType_IPV6,
			"DUALSTACK": OriginIpAddressType_DUALSTACK,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.OriginOptions",
		reflect.TypeOf((*OriginOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.OriginProps",
		reflect.TypeOf((*OriginProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.OriginProtocolPolicy",
		reflect.TypeOf((*OriginProtocolPolicy)(nil)).Elem(),
		map[string]interface{}{
			"HTTP_ONLY": OriginProtocolPolicy_HTTP_ONLY,
			"MATCH_VIEWER": OriginProtocolPolicy_MATCH_VIEWER,
			"HTTPS_ONLY": OriginProtocolPolicy_HTTPS_ONLY,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.OriginRequestCookieBehavior",
		reflect.TypeOf((*OriginRequestCookieBehavior)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "behavior", GoGetter: "Behavior"},
			_jsii_.MemberProperty{JsiiProperty: "cookies", GoGetter: "Cookies"},
		},
		func() interface{} {
			return &jsiiProxy_OriginRequestCookieBehavior{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.OriginRequestHeaderBehavior",
		reflect.TypeOf((*OriginRequestHeaderBehavior)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "behavior", GoGetter: "Behavior"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
		},
		func() interface{} {
			return &jsiiProxy_OriginRequestHeaderBehavior{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.OriginRequestPolicy",
		reflect.TypeOf((*OriginRequestPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyId", GoGetter: "OriginRequestPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyRef", GoGetter: "OriginRequestPolicyRef"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OriginRequestPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOriginRequestPolicy)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.OriginRequestPolicyProps",
		reflect.TypeOf((*OriginRequestPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.OriginRequestPolicyReference",
		reflect.TypeOf((*OriginRequestPolicyReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.OriginRequestQueryStringBehavior",
		reflect.TypeOf((*OriginRequestQueryStringBehavior)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "behavior", GoGetter: "Behavior"},
			_jsii_.MemberProperty{JsiiProperty: "queryStrings", GoGetter: "QueryStrings"},
		},
		func() interface{} {
			return &jsiiProxy_OriginRequestQueryStringBehavior{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.OriginSelectionCriteria",
		reflect.TypeOf((*OriginSelectionCriteria)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": OriginSelectionCriteria_DEFAULT,
			"MEDIA_QUALITY_BASED": OriginSelectionCriteria_MEDIA_QUALITY_BASED,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.OriginSslPolicy",
		reflect.TypeOf((*OriginSslPolicy)(nil)).Elem(),
		map[string]interface{}{
			"SSL_V3": OriginSslPolicy_SSL_V3,
			"TLS_V1": OriginSslPolicy_TLS_V1,
			"TLS_V1_1": OriginSslPolicy_TLS_V1_1,
			"TLS_V1_2": OriginSslPolicy_TLS_V1_2,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.PriceClass",
		reflect.TypeOf((*PriceClass)(nil)).Elem(),
		map[string]interface{}{
			"PRICE_CLASS_100": PriceClass_PRICE_CLASS_100,
			"PRICE_CLASS_200": PriceClass_PRICE_CLASS_200,
			"PRICE_CLASS_ALL": PriceClass_PRICE_CLASS_ALL,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.PublicKey",
		reflect.TypeOf((*PublicKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "publicKeyId", GoGetter: "PublicKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "publicKeyRef", GoGetter: "PublicKeyRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PublicKey{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPublicKey)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.PublicKeyProps",
		reflect.TypeOf((*PublicKeyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.PublicKeyReference",
		reflect.TypeOf((*PublicKeyReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.RealtimeLogConfig",
		reflect.TypeOf((*RealtimeLogConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigArn", GoGetter: "RealtimeLogConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigName", GoGetter: "RealtimeLogConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigRef", GoGetter: "RealtimeLogConfigRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RealtimeLogConfig{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRealtimeLogConfig)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.RealtimeLogConfigProps",
		reflect.TypeOf((*RealtimeLogConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.RealtimeLogConfigReference",
		reflect.TypeOf((*RealtimeLogConfigReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ResponseCustomHeader",
		reflect.TypeOf((*ResponseCustomHeader)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ResponseCustomHeadersBehavior",
		reflect.TypeOf((*ResponseCustomHeadersBehavior)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersContentSecurityPolicy",
		reflect.TypeOf((*ResponseHeadersContentSecurityPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersContentTypeOptions",
		reflect.TypeOf((*ResponseHeadersContentTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersCorsBehavior",
		reflect.TypeOf((*ResponseHeadersCorsBehavior)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersFrameOptions",
		reflect.TypeOf((*ResponseHeadersFrameOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersPolicy",
		reflect.TypeOf((*ResponseHeadersPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyId", GoGetter: "ResponseHeadersPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyRef", GoGetter: "ResponseHeadersPolicyRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ResponseHeadersPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResponseHeadersPolicy)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersPolicyProps",
		reflect.TypeOf((*ResponseHeadersPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersPolicyReference",
		reflect.TypeOf((*ResponseHeadersPolicyReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersReferrerPolicy",
		reflect.TypeOf((*ResponseHeadersReferrerPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersStrictTransportSecurity",
		reflect.TypeOf((*ResponseHeadersStrictTransportSecurity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersXSSProtection",
		reflect.TypeOf((*ResponseHeadersXSSProtection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ResponseSecurityHeadersBehavior",
		reflect.TypeOf((*ResponseSecurityHeadersBehavior)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.S3ImportSource",
		reflect.TypeOf((*S3ImportSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
		},
		func() interface{} {
			j := jsiiProxy_S3ImportSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ImportSource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.S3OriginAccessControl",
		reflect.TypeOf((*S3OriginAccessControl)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessControlId", GoGetter: "OriginAccessControlId"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_S3OriginAccessControl{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOriginAccessControl)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.S3OriginAccessControlProps",
		reflect.TypeOf((*S3OriginAccessControlProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.S3OriginConfig",
		reflect.TypeOf((*S3OriginConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.SSLMethod",
		reflect.TypeOf((*SSLMethod)(nil)).Elem(),
		map[string]interface{}{
			"SNI": SSLMethod_SNI,
			"VIP": SSLMethod_VIP,
			"STATIC_IP": SSLMethod_STATIC_IP,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.SecurityPolicyProtocol",
		reflect.TypeOf((*SecurityPolicyProtocol)(nil)).Elem(),
		map[string]interface{}{
			"SSL_V3": SecurityPolicyProtocol_SSL_V3,
			"TLS_V1": SecurityPolicyProtocol_TLS_V1,
			"TLS_V1_2016": SecurityPolicyProtocol_TLS_V1_2016,
			"TLS_V1_1_2016": SecurityPolicyProtocol_TLS_V1_1_2016,
			"TLS_V1_2_2018": SecurityPolicyProtocol_TLS_V1_2_2018,
			"TLS_V1_2_2019": SecurityPolicyProtocol_TLS_V1_2_2019,
			"TLS_V1_2_2021": SecurityPolicyProtocol_TLS_V1_2_2021,
			"TLS_V1_2_2025": SecurityPolicyProtocol_TLS_V1_2_2025,
			"TLS_V1_3_2025": SecurityPolicyProtocol_TLS_V1_3_2025,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.Signing",
		reflect.TypeOf((*Signing)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "behavior", GoGetter: "Behavior"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
		},
		func() interface{} {
			return &jsiiProxy_Signing{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.SigningBehavior",
		reflect.TypeOf((*SigningBehavior)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": SigningBehavior_ALWAYS,
			"NEVER": SigningBehavior_NEVER,
			"NO_OVERRIDE": SigningBehavior_NO_OVERRIDE,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.SigningProtocol",
		reflect.TypeOf((*SigningProtocol)(nil)).Elem(),
		map[string]interface{}{
			"SIGV4": SigningProtocol_SIGV4,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.SourceConfiguration",
		reflect.TypeOf((*SourceConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.StreamingDistributionReference",
		reflect.TypeOf((*StreamingDistributionReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.ViewerCertificate",
		reflect.TypeOf((*ViewerCertificate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aliases", GoGetter: "Aliases"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			return &jsiiProxy_ViewerCertificate{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.ViewerCertificateOptions",
		reflect.TypeOf((*ViewerCertificateOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudfront.ViewerProtocolPolicy",
		reflect.TypeOf((*ViewerProtocolPolicy)(nil)).Elem(),
		map[string]interface{}{
			"HTTPS_ONLY": ViewerProtocolPolicy_HTTPS_ONLY,
			"REDIRECT_TO_HTTPS": ViewerProtocolPolicy_REDIRECT_TO_HTTPS,
			"ALLOW_ALL": ViewerProtocolPolicy_ALLOW_ALL,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.VpcOrigin",
		reflect.TypeOf((*VpcOrigin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcOriginArn", GoGetter: "VpcOriginArn"},
			_jsii_.MemberProperty{JsiiProperty: "vpcOriginId", GoGetter: "VpcOriginId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcOriginRef", GoGetter: "VpcOriginRef"},
		},
		func() interface{} {
			j := jsiiProxy_VpcOrigin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVpcOrigin)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.VpcOriginAttributes",
		reflect.TypeOf((*VpcOriginAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront.VpcOriginEndpoint",
		reflect.TypeOf((*VpcOriginEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "endpointArn", GoGetter: "EndpointArn"},
		},
		func() interface{} {
			return &jsiiProxy_VpcOriginEndpoint{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.VpcOriginOptions",
		reflect.TypeOf((*VpcOriginOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.VpcOriginProps",
		reflect.TypeOf((*VpcOriginProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront.VpcOriginReference",
		reflect.TypeOf((*VpcOriginReference)(nil)).Elem(),
	)
}
