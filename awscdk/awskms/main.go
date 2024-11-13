package awskms

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kms.Alias",
		reflect.TypeOf((*Alias)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAlias", GoMethod: "AddAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "aliasArn", GoGetter: "AliasArn"},
			_jsii_.MemberProperty{JsiiProperty: "aliasName", GoGetter: "AliasName"},
			_jsii_.MemberProperty{JsiiProperty: "aliasTargetKey", GoGetter: "AliasTargetKey"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantDecrypt", GoMethod: "GrantDecrypt"},
			_jsii_.MemberMethod{JsiiMethod: "grantEncrypt", GoMethod: "GrantEncrypt"},
			_jsii_.MemberMethod{JsiiMethod: "grantEncryptDecrypt", GoMethod: "GrantEncryptDecrypt"},
			_jsii_.MemberMethod{JsiiMethod: "grantGenerateMac", GoMethod: "GrantGenerateMac"},
			_jsii_.MemberMethod{JsiiMethod: "grantVerifyMac", GoMethod: "GrantVerifyMac"},
			_jsii_.MemberProperty{JsiiProperty: "keyArn", GoGetter: "KeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "keyId", GoGetter: "KeyId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Alias{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAlias)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kms.AliasAttributes",
		reflect.TypeOf((*AliasAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kms.AliasProps",
		reflect.TypeOf((*AliasProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kms.CfnAlias",
		reflect.TypeOf((*CfnAlias)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "aliasName", GoGetter: "AliasName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
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
			_jsii_.MemberProperty{JsiiProperty: "targetKeyId", GoGetter: "TargetKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAlias{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kms.CfnAliasProps",
		reflect.TypeOf((*CfnAliasProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kms.CfnKey",
		reflect.TypeOf((*CfnKey)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrKeyId", GoGetter: "AttrKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "bypassPolicyLockoutSafetyCheck", GoGetter: "BypassPolicyLockoutSafetyCheck"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enableKeyRotation", GoGetter: "EnableKeyRotation"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "keyPolicy", GoGetter: "KeyPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "keySpec", GoGetter: "KeySpec"},
			_jsii_.MemberProperty{JsiiProperty: "keyUsage", GoGetter: "KeyUsage"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "multiRegion", GoGetter: "MultiRegion"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pendingWindowInDays", GoGetter: "PendingWindowInDays"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "rotationPeriodInDays", GoGetter: "RotationPeriodInDays"},
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
			j := jsiiProxy_CfnKey{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kms.CfnKeyProps",
		reflect.TypeOf((*CfnKeyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kms.CfnReplicaKey",
		reflect.TypeOf((*CfnReplicaKey)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrKeyId", GoGetter: "AttrKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "keyPolicy", GoGetter: "KeyPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pendingWindowInDays", GoGetter: "PendingWindowInDays"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyArn", GoGetter: "PrimaryKeyArn"},
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
			j := jsiiProxy_CfnReplicaKey{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kms.CfnReplicaKeyProps",
		reflect.TypeOf((*CfnReplicaKeyProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_kms.IAlias",
		reflect.TypeOf((*IAlias)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAlias", GoMethod: "AddAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "aliasName", GoGetter: "AliasName"},
			_jsii_.MemberProperty{JsiiProperty: "aliasTargetKey", GoGetter: "AliasTargetKey"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantDecrypt", GoMethod: "GrantDecrypt"},
			_jsii_.MemberMethod{JsiiMethod: "grantEncrypt", GoMethod: "GrantEncrypt"},
			_jsii_.MemberMethod{JsiiMethod: "grantEncryptDecrypt", GoMethod: "GrantEncryptDecrypt"},
			_jsii_.MemberMethod{JsiiMethod: "grantGenerateMac", GoMethod: "GrantGenerateMac"},
			_jsii_.MemberMethod{JsiiMethod: "grantVerifyMac", GoMethod: "GrantVerifyMac"},
			_jsii_.MemberProperty{JsiiProperty: "keyArn", GoGetter: "KeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "keyId", GoGetter: "KeyId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IAlias{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IKey)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_kms.IKey",
		reflect.TypeOf((*IKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAlias", GoMethod: "AddAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantDecrypt", GoMethod: "GrantDecrypt"},
			_jsii_.MemberMethod{JsiiMethod: "grantEncrypt", GoMethod: "GrantEncrypt"},
			_jsii_.MemberMethod{JsiiMethod: "grantEncryptDecrypt", GoMethod: "GrantEncryptDecrypt"},
			_jsii_.MemberMethod{JsiiMethod: "grantGenerateMac", GoMethod: "GrantGenerateMac"},
			_jsii_.MemberMethod{JsiiMethod: "grantVerifyMac", GoMethod: "GrantVerifyMac"},
			_jsii_.MemberProperty{JsiiProperty: "keyArn", GoGetter: "KeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "keyId", GoGetter: "KeyId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IKey{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kms.Key",
		reflect.TypeOf((*Key)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAlias", GoMethod: "AddAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantAdmin", GoMethod: "GrantAdmin"},
			_jsii_.MemberMethod{JsiiMethod: "grantDecrypt", GoMethod: "GrantDecrypt"},
			_jsii_.MemberMethod{JsiiMethod: "grantEncrypt", GoMethod: "GrantEncrypt"},
			_jsii_.MemberMethod{JsiiMethod: "grantEncryptDecrypt", GoMethod: "GrantEncryptDecrypt"},
			_jsii_.MemberMethod{JsiiMethod: "grantGenerateMac", GoMethod: "GrantGenerateMac"},
			_jsii_.MemberMethod{JsiiMethod: "grantVerifyMac", GoMethod: "GrantVerifyMac"},
			_jsii_.MemberProperty{JsiiProperty: "keyArn", GoGetter: "KeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "keyId", GoGetter: "KeyId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustAccountIdentities", GoGetter: "TrustAccountIdentities"},
		},
		func() interface{} {
			j := jsiiProxy_Key{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IKey)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kms.KeyLookupOptions",
		reflect.TypeOf((*KeyLookupOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kms.KeyProps",
		reflect.TypeOf((*KeyProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_kms.KeySpec",
		reflect.TypeOf((*KeySpec)(nil)).Elem(),
		map[string]interface{}{
			"SYMMETRIC_DEFAULT": KeySpec_SYMMETRIC_DEFAULT,
			"RSA_2048": KeySpec_RSA_2048,
			"RSA_3072": KeySpec_RSA_3072,
			"RSA_4096": KeySpec_RSA_4096,
			"ECC_NIST_P256": KeySpec_ECC_NIST_P256,
			"ECC_NIST_P384": KeySpec_ECC_NIST_P384,
			"ECC_NIST_P521": KeySpec_ECC_NIST_P521,
			"ECC_SECG_P256K1": KeySpec_ECC_SECG_P256K1,
			"HMAC_224": KeySpec_HMAC_224,
			"HMAC_256": KeySpec_HMAC_256,
			"HMAC_384": KeySpec_HMAC_384,
			"HMAC_512": KeySpec_HMAC_512,
			"SM2": KeySpec_SM2,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_kms.KeyUsage",
		reflect.TypeOf((*KeyUsage)(nil)).Elem(),
		map[string]interface{}{
			"ENCRYPT_DECRYPT": KeyUsage_ENCRYPT_DECRYPT,
			"SIGN_VERIFY": KeyUsage_SIGN_VERIFY,
			"GENERATE_VERIFY_MAC": KeyUsage_GENERATE_VERIFY_MAC,
			"KEY_AGREEMENT": KeyUsage_KEY_AGREEMENT,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kms.ViaServicePrincipal",
		reflect.TypeOf((*ViaServicePrincipal)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToAssumeRolePolicy", GoMethod: "AddToAssumeRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addToPolicy", GoMethod: "AddToPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addToPrincipalPolicy", GoMethod: "AddToPrincipalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRoleAction", GoGetter: "AssumeRoleAction"},
			_jsii_.MemberMethod{JsiiMethod: "dedupeString", GoMethod: "DedupeString"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "policyFragment", GoGetter: "PolicyFragment"},
			_jsii_.MemberProperty{JsiiProperty: "principalAccount", GoGetter: "PrincipalAccount"},
			_jsii_.MemberMethod{JsiiMethod: "toJSON", GoMethod: "ToJSON"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "withConditions", GoMethod: "WithConditions"},
			_jsii_.MemberMethod{JsiiMethod: "withSessionTags", GoMethod: "WithSessionTags"},
		},
		func() interface{} {
			j := jsiiProxy_ViaServicePrincipal{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamPrincipalBase)
			return &j
		},
	)
}
