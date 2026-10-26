package awslicensemanagerusersubscriptions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpoint",
		reflect.TypeOf((*CfnLicenseServerEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addResourceDependency", GoMethod: "AddResourceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreationTime", GoGetter: "AttrCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrLicenseServerEndpointArn", GoGetter: "AttrLicenseServerEndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrLicenseServerEndpointId", GoGetter: "AttrLicenseServerEndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "attrLicenseServerEndpointProvisioningStatus", GoGetter: "AttrLicenseServerEndpointProvisioningStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrLicenseServers", GoGetter: "AttrLicenseServers"},
			_jsii_.MemberProperty{JsiiProperty: "attrServerEndpoint", GoGetter: "AttrServerEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "attrServerType", GoGetter: "AttrServerType"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderArn", GoGetter: "IdentityProviderArn"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "licenseServerEndpointRef", GoGetter: "LicenseServerEndpointRef"},
			_jsii_.MemberProperty{JsiiProperty: "licenseServerSettings", GoGetter: "LicenseServerSettings"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "removeResourceDependency", GoMethod: "RemoveResourceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLicenseServerEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawslicensemanagerusersubscriptionsILicenseServerEndpointRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpoint.CredentialsProviderProperty",
		reflect.TypeOf((*CfnLicenseServerEndpoint_CredentialsProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpoint.LicenseServerProperty",
		reflect.TypeOf((*CfnLicenseServerEndpoint_LicenseServerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpoint.LicenseServerSettingsProperty",
		reflect.TypeOf((*CfnLicenseServerEndpoint_LicenseServerSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpoint.RdsSalSettingsProperty",
		reflect.TypeOf((*CfnLicenseServerEndpoint_RdsSalSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpoint.SecretsManagerCredentialsProviderProperty",
		reflect.TypeOf((*CfnLicenseServerEndpoint_SecretsManagerCredentialsProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpoint.ServerEndpointProperty",
		reflect.TypeOf((*CfnLicenseServerEndpoint_ServerEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpoint.ServerSettingsProperty",
		reflect.TypeOf((*CfnLicenseServerEndpoint_ServerSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpointProps",
		reflect.TypeOf((*CfnLicenseServerEndpointProps)(nil)).Elem(),
	)
}
