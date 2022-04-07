package awscdkservicecatalogalpha

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
			_jsii_.MemberMethod{JsiiMethod: "associateTagOptions", GoMethod: "AssociateTagOptions"},
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
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
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
			_jsii_.MemberMethod{JsiiMethod: "setLocalLaunchRole", GoMethod: "SetLocalLaunchRole"},
			_jsii_.MemberMethod{JsiiMethod: "setLocalLaunchRoleName", GoMethod: "SetLocalLaunchRoleName"},
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
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateTagOptions", GoMethod: "AssociateTagOptions"},
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
			_jsii_.MemberMethod{JsiiMethod: "setLocalLaunchRole", GoMethod: "SetLocalLaunchRole"},
			_jsii_.MemberMethod{JsiiMethod: "setLocalLaunchRoleName", GoMethod: "SetLocalLaunchRoleName"},
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
			_jsii_.MemberMethod{JsiiMethod: "associateTagOptions", GoMethod: "AssociateTagOptions"},
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
	_jsii_.RegisterClass(
		"@aws-cdk/aws-servicecatalog-alpha.ProductStack",
		reflect.TypeOf((*ProductStack)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "account", GoGetter: "Account"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addTransform", GoMethod: "AddTransform"},
			_jsii_.MemberMethod{JsiiMethod: "allocateLogicalId", GoMethod: "AllocateLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "artifactId", GoGetter: "ArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "bundlingRequired", GoGetter: "BundlingRequired"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberMethod{JsiiMethod: "exportValue", GoMethod: "ExportValue"},
			_jsii_.MemberMethod{JsiiMethod: "formatArn", GoMethod: "FormatArn"},
			_jsii_.MemberMethod{JsiiMethod: "getLogicalId", GoMethod: "GetLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "nested", GoGetter: "Nested"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackParent", GoGetter: "NestedStackParent"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackResource", GoGetter: "NestedStackResource"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationArns", GoGetter: "NotificationArns"},
			_jsii_.MemberProperty{JsiiProperty: "partition", GoGetter: "Partition"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "regionalFact", GoMethod: "RegionalFact"},
			_jsii_.MemberMethod{JsiiMethod: "renameLogicalId", GoMethod: "RenameLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "reportMissingContextKey", GoMethod: "ReportMissingContextKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "splitArn", GoMethod: "SplitArn"},
			_jsii_.MemberProperty{JsiiProperty: "stackId", GoGetter: "StackId"},
			_jsii_.MemberProperty{JsiiProperty: "stackName", GoGetter: "StackName"},
			_jsii_.MemberProperty{JsiiProperty: "synthesizer", GoGetter: "Synthesizer"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "templateFile", GoGetter: "TemplateFile"},
			_jsii_.MemberProperty{JsiiProperty: "templateOptions", GoGetter: "TemplateOptions"},
			_jsii_.MemberProperty{JsiiProperty: "terminationProtection", GoGetter: "TerminationProtection"},
			_jsii_.MemberMethod{JsiiMethod: "toJsonString", GoMethod: "ToJsonString"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "urlSuffix", GoGetter: "UrlSuffix"},
		},
		func() interface{} {
			j := jsiiProxy_ProductStack{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkStack)
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
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
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
			j := jsiiProxy_TagOptions{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-servicecatalog-alpha.TagOptionsProps",
		reflect.TypeOf((*TagOptionsProps)(nil)).Elem(),
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
