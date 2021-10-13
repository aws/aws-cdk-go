package awscdkawsservicecatalogappregistryalpha

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
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateAttributeGroup", GoMethod: "AssociateAttributeGroup"},
			_jsii_.MemberMethod{JsiiMethod: "associateStack", GoMethod: "AssociateStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "generateUniqueHash", GoMethod: "GenerateUniqueHash"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
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
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
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
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.IApplication",
		reflect.TypeOf((*IApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "associateAttributeGroup", GoMethod: "AssociateAttributeGroup"},
			_jsii_.MemberMethod{JsiiMethod: "associateStack", GoMethod: "AssociateStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
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
			_jsii_.MemberProperty{JsiiProperty: "attributeGroupArn", GoGetter: "AttributeGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "attributeGroupId", GoGetter: "AttributeGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IAttributeGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
}
