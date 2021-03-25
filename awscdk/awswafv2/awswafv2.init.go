package awswafv2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_wafv2.CfnIPSet",
		reflect.TypeOf((*CfnIPSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "addresses", GoGetter: "Addresses"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddressVersion", GoGetter: "IpAddressVersion"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIPSet{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnIPSetProps",
		reflect.TypeOf((*CfnIPSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_wafv2.CfnRegexPatternSet",
		reflect.TypeOf((*CfnRegexPatternSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberProperty{JsiiProperty: "regularExpressionList", GoGetter: "RegularExpressionList"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRegexPatternSet{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRegexPatternSetProps",
		reflect.TypeOf((*CfnRegexPatternSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_wafv2.CfnRuleGroup",
		reflect.TypeOf((*CfnRuleGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "capacity", GoGetter: "Capacity"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "visibilityConfig", GoGetter: "VisibilityConfig"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRuleGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.AndStatementOneProperty",
		reflect.TypeOf((*CfnRuleGroup_AndStatementOneProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.AndStatementTwoProperty",
		reflect.TypeOf((*CfnRuleGroup_AndStatementTwoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.ByteMatchStatementProperty",
		reflect.TypeOf((*CfnRuleGroup_ByteMatchStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.FieldToMatchProperty",
		reflect.TypeOf((*CfnRuleGroup_FieldToMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.ForwardedIPConfigurationProperty",
		reflect.TypeOf((*CfnRuleGroup_ForwardedIPConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.GeoMatchStatementProperty",
		reflect.TypeOf((*CfnRuleGroup_GeoMatchStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_wafv2.CfnRuleGroup.IPSetForwardedIPConfigurationProperty",
		reflect.TypeOf((*CfnRuleGroup_IPSetForwardedIPConfigurationProperty)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fallbackBehavior", GoGetter: "FallbackBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "position", GoGetter: "Position"},
		},
		func() interface{} {
			return &jsiiProxy_CfnRuleGroup_IPSetForwardedIPConfigurationProperty{}
		},
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_wafv2.CfnRuleGroup.IPSetReferenceStatementProperty",
		reflect.TypeOf((*CfnRuleGroup_IPSetReferenceStatementProperty)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "ipSetForwardedIpConfig", GoGetter: "IpSetForwardedIpConfig"},
		},
		func() interface{} {
			return &jsiiProxy_CfnRuleGroup_IPSetReferenceStatementProperty{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.NotStatementOneProperty",
		reflect.TypeOf((*CfnRuleGroup_NotStatementOneProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.NotStatementTwoProperty",
		reflect.TypeOf((*CfnRuleGroup_NotStatementTwoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.OrStatementOneProperty",
		reflect.TypeOf((*CfnRuleGroup_OrStatementOneProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.OrStatementTwoProperty",
		reflect.TypeOf((*CfnRuleGroup_OrStatementTwoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.RateBasedStatementOneProperty",
		reflect.TypeOf((*CfnRuleGroup_RateBasedStatementOneProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.RateBasedStatementTwoProperty",
		reflect.TypeOf((*CfnRuleGroup_RateBasedStatementTwoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.RegexPatternSetReferenceStatementProperty",
		reflect.TypeOf((*CfnRuleGroup_RegexPatternSetReferenceStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.RuleActionProperty",
		reflect.TypeOf((*CfnRuleGroup_RuleActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.RuleProperty",
		reflect.TypeOf((*CfnRuleGroup_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.SizeConstraintStatementProperty",
		reflect.TypeOf((*CfnRuleGroup_SizeConstraintStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.SqliMatchStatementProperty",
		reflect.TypeOf((*CfnRuleGroup_SqliMatchStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.StatementOneProperty",
		reflect.TypeOf((*CfnRuleGroup_StatementOneProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.StatementThreeProperty",
		reflect.TypeOf((*CfnRuleGroup_StatementThreeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.StatementTwoProperty",
		reflect.TypeOf((*CfnRuleGroup_StatementTwoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.TextTransformationProperty",
		reflect.TypeOf((*CfnRuleGroup_TextTransformationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.VisibilityConfigProperty",
		reflect.TypeOf((*CfnRuleGroup_VisibilityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroup.XssMatchStatementProperty",
		reflect.TypeOf((*CfnRuleGroup_XssMatchStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnRuleGroupProps",
		reflect.TypeOf((*CfnRuleGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_wafv2.CfnWebACL",
		reflect.TypeOf((*CfnWebACL)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCapacity", GoGetter: "AttrCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAction", GoGetter: "DefaultAction"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "visibilityConfig", GoGetter: "VisibilityConfig"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWebACL{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.AndStatementOneProperty",
		reflect.TypeOf((*CfnWebACL_AndStatementOneProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.AndStatementTwoProperty",
		reflect.TypeOf((*CfnWebACL_AndStatementTwoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.ByteMatchStatementProperty",
		reflect.TypeOf((*CfnWebACL_ByteMatchStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.DefaultActionProperty",
		reflect.TypeOf((*CfnWebACL_DefaultActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.ExcludedRuleProperty",
		reflect.TypeOf((*CfnWebACL_ExcludedRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.FieldToMatchProperty",
		reflect.TypeOf((*CfnWebACL_FieldToMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.ForwardedIPConfigurationProperty",
		reflect.TypeOf((*CfnWebACL_ForwardedIPConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.GeoMatchStatementProperty",
		reflect.TypeOf((*CfnWebACL_GeoMatchStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_wafv2.CfnWebACL.IPSetForwardedIPConfigurationProperty",
		reflect.TypeOf((*CfnWebACL_IPSetForwardedIPConfigurationProperty)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fallbackBehavior", GoGetter: "FallbackBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "position", GoGetter: "Position"},
		},
		func() interface{} {
			return &jsiiProxy_CfnWebACL_IPSetForwardedIPConfigurationProperty{}
		},
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_wafv2.CfnWebACL.IPSetReferenceStatementProperty",
		reflect.TypeOf((*CfnWebACL_IPSetReferenceStatementProperty)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "ipSetForwardedIpConfig", GoGetter: "IpSetForwardedIpConfig"},
		},
		func() interface{} {
			return &jsiiProxy_CfnWebACL_IPSetReferenceStatementProperty{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.ManagedRuleGroupStatementProperty",
		reflect.TypeOf((*CfnWebACL_ManagedRuleGroupStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.NotStatementOneProperty",
		reflect.TypeOf((*CfnWebACL_NotStatementOneProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.NotStatementTwoProperty",
		reflect.TypeOf((*CfnWebACL_NotStatementTwoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.OrStatementOneProperty",
		reflect.TypeOf((*CfnWebACL_OrStatementOneProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.OrStatementTwoProperty",
		reflect.TypeOf((*CfnWebACL_OrStatementTwoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.OverrideActionProperty",
		reflect.TypeOf((*CfnWebACL_OverrideActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.RateBasedStatementOneProperty",
		reflect.TypeOf((*CfnWebACL_RateBasedStatementOneProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.RateBasedStatementTwoProperty",
		reflect.TypeOf((*CfnWebACL_RateBasedStatementTwoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.RegexPatternSetReferenceStatementProperty",
		reflect.TypeOf((*CfnWebACL_RegexPatternSetReferenceStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.RuleActionProperty",
		reflect.TypeOf((*CfnWebACL_RuleActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.RuleGroupReferenceStatementProperty",
		reflect.TypeOf((*CfnWebACL_RuleGroupReferenceStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.RuleProperty",
		reflect.TypeOf((*CfnWebACL_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.SizeConstraintStatementProperty",
		reflect.TypeOf((*CfnWebACL_SizeConstraintStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.SqliMatchStatementProperty",
		reflect.TypeOf((*CfnWebACL_SqliMatchStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.StatementOneProperty",
		reflect.TypeOf((*CfnWebACL_StatementOneProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.StatementThreeProperty",
		reflect.TypeOf((*CfnWebACL_StatementThreeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.StatementTwoProperty",
		reflect.TypeOf((*CfnWebACL_StatementTwoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.TextTransformationProperty",
		reflect.TypeOf((*CfnWebACL_TextTransformationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.VisibilityConfigProperty",
		reflect.TypeOf((*CfnWebACL_VisibilityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACL.XssMatchStatementProperty",
		reflect.TypeOf((*CfnWebACL_XssMatchStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_wafv2.CfnWebACLAssociation",
		reflect.TypeOf((*CfnWebACLAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
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
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "resourceArn", GoGetter: "ResourceArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "webAclArn", GoGetter: "WebAclArn"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWebACLAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACLAssociationProps",
		reflect.TypeOf((*CfnWebACLAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_wafv2.CfnWebACLProps",
		reflect.TypeOf((*CfnWebACLProps)(nil)).Elem(),
	)
}
