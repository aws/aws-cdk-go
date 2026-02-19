// The CDK Construct Library for AWS::Route53Resolver
package awscdkroute53resolveralpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-route53resolver-alpha.DnsBlockResponse",
		reflect.TypeOf((*DnsBlockResponse)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "blockOverrideDnsType", GoGetter: "BlockOverrideDnsType"},
			_jsii_.MemberProperty{JsiiProperty: "blockOverrideDomain", GoGetter: "BlockOverrideDomain"},
			_jsii_.MemberProperty{JsiiProperty: "blockOverrideTtl", GoGetter: "BlockOverrideTtl"},
			_jsii_.MemberProperty{JsiiProperty: "blockResponse", GoGetter: "BlockResponse"},
		},
		func() interface{} {
			return &jsiiProxy_DnsBlockResponse{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-route53resolver-alpha.DomainsConfig",
		reflect.TypeOf((*DomainsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomainList",
		reflect.TypeOf((*FirewallDomainList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "firewallDomainListArn", GoGetter: "FirewallDomainListArn"},
			_jsii_.MemberProperty{JsiiProperty: "firewallDomainListCreationTime", GoGetter: "FirewallDomainListCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "firewallDomainListCreatorRequestId", GoGetter: "FirewallDomainListCreatorRequestId"},
			_jsii_.MemberProperty{JsiiProperty: "firewallDomainListDomainCount", GoGetter: "FirewallDomainListDomainCount"},
			_jsii_.MemberProperty{JsiiProperty: "firewallDomainListId", GoGetter: "FirewallDomainListId"},
			_jsii_.MemberProperty{JsiiProperty: "firewallDomainListManagedOwnerName", GoGetter: "FirewallDomainListManagedOwnerName"},
			_jsii_.MemberProperty{JsiiProperty: "firewallDomainListModificationTime", GoGetter: "FirewallDomainListModificationTime"},
			_jsii_.MemberProperty{JsiiProperty: "firewallDomainListStatus", GoGetter: "FirewallDomainListStatus"},
			_jsii_.MemberProperty{JsiiProperty: "firewallDomainListStatusMessage", GoGetter: "FirewallDomainListStatusMessage"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_FirewallDomainList{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFirewallDomainList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomainListProps",
		reflect.TypeOf((*FirewallDomainListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomains",
		reflect.TypeOf((*FirewallDomains)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_FirewallDomains{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRule",
		reflect.TypeOf((*FirewallRule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleAction",
		reflect.TypeOf((*FirewallRuleAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberProperty{JsiiProperty: "blockResponse", GoGetter: "BlockResponse"},
		},
		func() interface{} {
			return &jsiiProxy_FirewallRuleAction{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroup",
		reflect.TypeOf((*FirewallRuleGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRule", GoMethod: "AddRule"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associate", GoMethod: "Associate"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupArn", GoGetter: "FirewallRuleGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupCreationTime", GoGetter: "FirewallRuleGroupCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupCreatorRequestId", GoGetter: "FirewallRuleGroupCreatorRequestId"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupId", GoGetter: "FirewallRuleGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupModificationTime", GoGetter: "FirewallRuleGroupModificationTime"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupOwnerId", GoGetter: "FirewallRuleGroupOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupRuleCount", GoGetter: "FirewallRuleGroupRuleCount"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupShareStatus", GoGetter: "FirewallRuleGroupShareStatus"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupStatus", GoGetter: "FirewallRuleGroupStatus"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupStatusMessage", GoGetter: "FirewallRuleGroupStatusMessage"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_FirewallRuleGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFirewallRuleGroup)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroupAssociation",
		reflect.TypeOf((*FirewallRuleGroupAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupAssociationArn", GoGetter: "FirewallRuleGroupAssociationArn"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupAssociationCreationTime", GoGetter: "FirewallRuleGroupAssociationCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupAssociationCreatorRequestId", GoGetter: "FirewallRuleGroupAssociationCreatorRequestId"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupAssociationId", GoGetter: "FirewallRuleGroupAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupAssociationManagedOwnerName", GoGetter: "FirewallRuleGroupAssociationManagedOwnerName"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupAssociationModificationTime", GoGetter: "FirewallRuleGroupAssociationModificationTime"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupAssociationStatus", GoGetter: "FirewallRuleGroupAssociationStatus"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupAssociationStatusMessage", GoGetter: "FirewallRuleGroupAssociationStatusMessage"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_FirewallRuleGroupAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroupAssociationOptions",
		reflect.TypeOf((*FirewallRuleGroupAssociationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroupAssociationProps",
		reflect.TypeOf((*FirewallRuleGroupAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroupProps",
		reflect.TypeOf((*FirewallRuleGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-route53resolver-alpha.IFirewallDomainList",
		reflect.TypeOf((*IFirewallDomainList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "firewallDomainListId", GoGetter: "FirewallDomainListId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IFirewallDomainList{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-route53resolver-alpha.IFirewallRuleGroup",
		reflect.TypeOf((*IFirewallRuleGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "firewallRuleGroupId", GoGetter: "FirewallRuleGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IFirewallRuleGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
}
