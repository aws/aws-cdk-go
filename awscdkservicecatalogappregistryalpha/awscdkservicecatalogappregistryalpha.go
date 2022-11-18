package awscdkservicecatalogappregistryalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.Application",
		reflect.TypeOf((*Application)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateAllStacksInScope", GoMethod: "AssociateAllStacksInScope"},
			_jsii_.MemberMethod{JsiiMethod: "associateApplicationWithStack", GoMethod: "AssociateApplicationWithStack"},
			_jsii_.MemberMethod{JsiiMethod: "associateAttributeGroup", GoMethod: "AssociateAttributeGroup"},
			_jsii_.MemberMethod{JsiiMethod: "associateStack", GoMethod: "AssociateStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "generateUniqueHash", GoMethod: "GenerateUniqueHash"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "shareApplication", GoMethod: "ShareApplication"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Application{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApplication)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.ApplicationAssociator",
		reflect.TypeOf((*ApplicationAssociator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "appRegistryApplication", GoMethod: "AppRegistryApplication"},
			_jsii_.MemberMethod{JsiiMethod: "associateStage", GoMethod: "AssociateStage"},
			_jsii_.MemberMethod{JsiiMethod: "isStageAssociated", GoMethod: "IsStageAssociated"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationAssociator{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.ApplicationAssociatorProps",
		reflect.TypeOf((*ApplicationAssociatorProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.ApplicationProps",
		reflect.TypeOf((*ApplicationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.AttributeGroup",
		reflect.TypeOf((*AttributeGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attributeGroupArn", GoGetter: "AttributeGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "attributeGroupId", GoGetter: "AttributeGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getAttributeGroupSharePermissionARN", GoMethod: "GetAttributeGroupSharePermissionARN"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "shareAttributeGroup", GoMethod: "ShareAttributeGroup"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AttributeGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAttributeGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.AttributeGroupProps",
		reflect.TypeOf((*AttributeGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.BindTargetApplicationResult",
		reflect.TypeOf((*BindTargetApplicationResult)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.CreateTargetApplicationOptions",
		reflect.TypeOf((*CreateTargetApplicationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.ExistingTargetApplicationOptions",
		reflect.TypeOf((*ExistingTargetApplicationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.IApplication",
		reflect.TypeOf((*IApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateAllStacksInScope", GoMethod: "AssociateAllStacksInScope"},
			_jsii_.MemberMethod{JsiiMethod: "associateApplicationWithStack", GoMethod: "AssociateApplicationWithStack"},
			_jsii_.MemberMethod{JsiiMethod: "associateAttributeGroup", GoMethod: "AssociateAttributeGroup"},
			_jsii_.MemberMethod{JsiiMethod: "associateStack", GoMethod: "AssociateStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "shareApplication", GoMethod: "ShareApplication"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IApplication{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.IAttributeGroup",
		reflect.TypeOf((*IAttributeGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attributeGroupArn", GoGetter: "AttributeGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "attributeGroupId", GoGetter: "AttributeGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "shareAttributeGroup", GoMethod: "ShareAttributeGroup"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IAttributeGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.ShareOptions",
		reflect.TypeOf((*ShareOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.SharePermission",
		reflect.TypeOf((*SharePermission)(nil)).Elem(),
		map[string]interface{}{
			"READ_ONLY": SharePermission_READ_ONLY,
			"ALLOW_ACCESS": SharePermission_ALLOW_ACCESS,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.TargetApplication",
		reflect.TypeOf((*TargetApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_TargetApplication{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.TargetApplicationCommonOptions",
		reflect.TypeOf((*TargetApplicationCommonOptions)(nil)).Elem(),
	)
}
