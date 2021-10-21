package awscdkawsservicecatalogalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-servicecatalog-alpha.CloudFormationProduct",
		reflect.TypeOf((*CloudFormationProduct)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "productArn", GoGetter: "ProductArn"},
			_jsii_.MemberProperty{JsiiProperty: "productId", GoGetter: "ProductId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudFormationProduct{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Product)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalog-alpha.CloudFormationProductProps",
		reflect.TypeOf((*CloudFormationProductProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalog-alpha.CloudFormationProductVersion",
		reflect.TypeOf((*CloudFormationProductVersion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalog-alpha.CloudFormationRuleConstraintOptions",
		reflect.TypeOf((*CloudFormationRuleConstraintOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-servicecatalog-alpha.CloudFormationTemplate",
		reflect.TypeOf((*CloudFormationTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_CloudFormationTemplate{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalog-alpha.CloudFormationTemplateConfig",
		reflect.TypeOf((*CloudFormationTemplateConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalog-alpha.CommonConstraintOptions",
		reflect.TypeOf((*CommonConstraintOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-servicecatalog-alpha.IPortfolio",
		reflect.TypeOf((*IPortfolio)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProduct", GoMethod: "AddProduct"},
			_jsii_.MemberMethod{JsiiMethod: "associateTagOptions", GoMethod: "AssociateTagOptions"},
			_jsii_.MemberMethod{JsiiMethod: "constrainCloudFormationParameters", GoMethod: "ConstrainCloudFormationParameters"},
			_jsii_.MemberMethod{JsiiMethod: "constrainTagUpdates", GoMethod: "ConstrainTagUpdates"},
			_jsii_.MemberMethod{JsiiMethod: "deployWithStackSets", GoMethod: "DeployWithStackSets"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "giveAccessToGroup", GoMethod: "GiveAccessToGroup"},
			_jsii_.MemberMethod{JsiiMethod: "giveAccessToRole", GoMethod: "GiveAccessToRole"},
			_jsii_.MemberMethod{JsiiMethod: "giveAccessToUser", GoMethod: "GiveAccessToUser"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnStackEvents", GoMethod: "NotifyOnStackEvents"},
			_jsii_.MemberProperty{JsiiProperty: "portfolioArn", GoGetter: "PortfolioArn"},
			_jsii_.MemberProperty{JsiiProperty: "portfolioId", GoGetter: "PortfolioId"},
			_jsii_.MemberMethod{JsiiMethod: "setLaunchRole", GoMethod: "SetLaunchRole"},
			_jsii_.MemberMethod{JsiiMethod: "shareWithAccount", GoMethod: "ShareWithAccount"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IPortfolio{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-servicecatalog-alpha.IProduct",
		reflect.TypeOf((*IProduct)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "productArn", GoGetter: "ProductArn"},
			_jsii_.MemberProperty{JsiiProperty: "productId", GoGetter: "ProductId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IProduct{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-servicecatalog-alpha.MessageLanguage",
		reflect.TypeOf((*MessageLanguage)(nil)).Elem(),
		map[string]interface{}{
			"EN": MessageLanguage_EN,
			"JP": MessageLanguage_JP,
			"ZH": MessageLanguage_ZH,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-servicecatalog-alpha.Portfolio",
		reflect.TypeOf((*Portfolio)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProduct", GoMethod: "AddProduct"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateTagOptions", GoMethod: "AssociateTagOptions"},
			_jsii_.MemberMethod{JsiiMethod: "constrainCloudFormationParameters", GoMethod: "ConstrainCloudFormationParameters"},
			_jsii_.MemberMethod{JsiiMethod: "constrainTagUpdates", GoMethod: "ConstrainTagUpdates"},
			_jsii_.MemberMethod{JsiiMethod: "deployWithStackSets", GoMethod: "DeployWithStackSets"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "generateUniqueHash", GoMethod: "GenerateUniqueHash"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "giveAccessToGroup", GoMethod: "GiveAccessToGroup"},
			_jsii_.MemberMethod{JsiiMethod: "giveAccessToRole", GoMethod: "GiveAccessToRole"},
			_jsii_.MemberMethod{JsiiMethod: "giveAccessToUser", GoMethod: "GiveAccessToUser"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnStackEvents", GoMethod: "NotifyOnStackEvents"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "portfolioArn", GoGetter: "PortfolioArn"},
			_jsii_.MemberProperty{JsiiProperty: "portfolioId", GoGetter: "PortfolioId"},
			_jsii_.MemberMethod{JsiiMethod: "setLaunchRole", GoMethod: "SetLaunchRole"},
			_jsii_.MemberMethod{JsiiMethod: "shareWithAccount", GoMethod: "ShareWithAccount"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Portfolio{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPortfolio)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalog-alpha.PortfolioProps",
		reflect.TypeOf((*PortfolioProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalog-alpha.PortfolioShareOptions",
		reflect.TypeOf((*PortfolioShareOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-servicecatalog-alpha.Product",
		reflect.TypeOf((*Product)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "productArn", GoGetter: "ProductArn"},
			_jsii_.MemberProperty{JsiiProperty: "productId", GoGetter: "ProductId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Product{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IProduct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalog-alpha.StackSetsConstraintOptions",
		reflect.TypeOf((*StackSetsConstraintOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-servicecatalog-alpha.TagOptions",
		reflect.TypeOf((*TagOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "tagOptionsMap", GoGetter: "TagOptionsMap"},
		},
		func() interface{} {
			return &jsiiProxy_TagOptions{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalog-alpha.TagUpdateConstraintOptions",
		reflect.TypeOf((*TagUpdateConstraintOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalog-alpha.TemplateRule",
		reflect.TypeOf((*TemplateRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalog-alpha.TemplateRuleAssertion",
		reflect.TypeOf((*TemplateRuleAssertion)(nil)).Elem(),
	)
}
