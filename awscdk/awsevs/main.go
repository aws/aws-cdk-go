package awsevs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_evs.CfnEnvironment",
		reflect.TypeOf((*CfnEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrChecks", GoGetter: "AttrChecks"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrCredentials", GoGetter: "AttrCredentials"},
			_jsii_.MemberProperty{JsiiProperty: "attrEnvironmentArn", GoGetter: "AttrEnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrEnvironmentId", GoGetter: "AttrEnvironmentId"},
			_jsii_.MemberProperty{JsiiProperty: "attrEnvironmentState", GoGetter: "AttrEnvironmentState"},
			_jsii_.MemberProperty{JsiiProperty: "attrModifiedAt", GoGetter: "AttrModifiedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrStateDetails", GoGetter: "AttrStateDetails"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "connectivityInfo", GoGetter: "ConnectivityInfo"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "environmentName", GoGetter: "EnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "environmentRef", GoGetter: "EnvironmentRef"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "hosts", GoGetter: "Hosts"},
			_jsii_.MemberProperty{JsiiProperty: "initialVlans", GoGetter: "InitialVlans"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "licenseInfo", GoGetter: "LicenseInfo"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccessSecurityGroups", GoGetter: "ServiceAccessSecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccessSubnetId", GoGetter: "ServiceAccessSubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "siteId", GoGetter: "SiteId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "termsAccepted", GoGetter: "TermsAccepted"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vcfHostnames", GoGetter: "VcfHostnames"},
			_jsii_.MemberProperty{JsiiProperty: "vcfVersion", GoGetter: "VcfVersion"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsevsIEnvironmentRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_evs.CfnEnvironment.CheckProperty",
		reflect.TypeOf((*CfnEnvironment_CheckProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_evs.CfnEnvironment.ConnectivityInfoProperty",
		reflect.TypeOf((*CfnEnvironment_ConnectivityInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_evs.CfnEnvironment.HostInfoForCreateProperty",
		reflect.TypeOf((*CfnEnvironment_HostInfoForCreateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_evs.CfnEnvironment.InitialVlanInfoProperty",
		reflect.TypeOf((*CfnEnvironment_InitialVlanInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_evs.CfnEnvironment.InitialVlansProperty",
		reflect.TypeOf((*CfnEnvironment_InitialVlansProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_evs.CfnEnvironment.LicenseInfoProperty",
		reflect.TypeOf((*CfnEnvironment_LicenseInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_evs.CfnEnvironment.SecretProperty",
		reflect.TypeOf((*CfnEnvironment_SecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_evs.CfnEnvironment.ServiceAccessSecurityGroupsProperty",
		reflect.TypeOf((*CfnEnvironment_ServiceAccessSecurityGroupsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_evs.CfnEnvironment.VcfHostnamesProperty",
		reflect.TypeOf((*CfnEnvironment_VcfHostnamesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_evs.CfnEnvironmentProps",
		reflect.TypeOf((*CfnEnvironmentProps)(nil)).Elem(),
	)
}
