package interfacesawssecurityhub

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_securityhub.AggregatorV2Reference",
		reflect.TypeOf((*AggregatorV2Reference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_securityhub.AutomationRuleReference",
		reflect.TypeOf((*AutomationRuleReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_securityhub.AutomationRuleV2Reference",
		reflect.TypeOf((*AutomationRuleV2Reference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_securityhub.ConfigurationPolicyReference",
		reflect.TypeOf((*ConfigurationPolicyReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_securityhub.ConnectorV2Reference",
		reflect.TypeOf((*ConnectorV2Reference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_securityhub.DelegatedAdminReference",
		reflect.TypeOf((*DelegatedAdminReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_securityhub.FindingAggregatorReference",
		reflect.TypeOf((*FindingAggregatorReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_securityhub.HubReference",
		reflect.TypeOf((*HubReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_securityhub.HubV2Reference",
		reflect.TypeOf((*HubV2Reference)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_securityhub.IAggregatorV2Ref",
		reflect.TypeOf((*IAggregatorV2Ref)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aggregatorV2Ref", GoGetter: "AggregatorV2Ref"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IAggregatorV2Ref{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_securityhub.IAutomationRuleRef",
		reflect.TypeOf((*IAutomationRuleRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "automationRuleRef", GoGetter: "AutomationRuleRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IAutomationRuleRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_securityhub.IAutomationRuleV2Ref",
		reflect.TypeOf((*IAutomationRuleV2Ref)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "automationRuleV2Ref", GoGetter: "AutomationRuleV2Ref"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IAutomationRuleV2Ref{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_securityhub.IConfigurationPolicyRef",
		reflect.TypeOf((*IConfigurationPolicyRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "configurationPolicyRef", GoGetter: "ConfigurationPolicyRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IConfigurationPolicyRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_securityhub.IConnectorV2Ref",
		reflect.TypeOf((*IConnectorV2Ref)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connectorV2Ref", GoGetter: "ConnectorV2Ref"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IConnectorV2Ref{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_securityhub.IDelegatedAdminRef",
		reflect.TypeOf((*IDelegatedAdminRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "delegatedAdminRef", GoGetter: "DelegatedAdminRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IDelegatedAdminRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_securityhub.IFindingAggregatorRef",
		reflect.TypeOf((*IFindingAggregatorRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "findingAggregatorRef", GoGetter: "FindingAggregatorRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IFindingAggregatorRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_securityhub.IHubRef",
		reflect.TypeOf((*IHubRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "hubRef", GoGetter: "HubRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IHubRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_securityhub.IHubV2Ref",
		reflect.TypeOf((*IHubV2Ref)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "hubV2Ref", GoGetter: "HubV2Ref"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IHubV2Ref{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_securityhub.IInsightRef",
		reflect.TypeOf((*IInsightRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "insightRef", GoGetter: "InsightRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IInsightRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_securityhub.IOrganizationConfigurationRef",
		reflect.TypeOf((*IOrganizationConfigurationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "organizationConfigurationRef", GoGetter: "OrganizationConfigurationRef"},
		},
		func() interface{} {
			j := jsiiProxy_IOrganizationConfigurationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_securityhub.IPolicyAssociationRef",
		reflect.TypeOf((*IPolicyAssociationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "policyAssociationRef", GoGetter: "PolicyAssociationRef"},
		},
		func() interface{} {
			j := jsiiProxy_IPolicyAssociationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_securityhub.IProductSubscriptionRef",
		reflect.TypeOf((*IProductSubscriptionRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "productSubscriptionRef", GoGetter: "ProductSubscriptionRef"},
		},
		func() interface{} {
			j := jsiiProxy_IProductSubscriptionRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_securityhub.ISecurityControlRef",
		reflect.TypeOf((*ISecurityControlRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "securityControlRef", GoGetter: "SecurityControlRef"},
		},
		func() interface{} {
			j := jsiiProxy_ISecurityControlRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_securityhub.IStandardRef",
		reflect.TypeOf((*IStandardRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "standardRef", GoGetter: "StandardRef"},
		},
		func() interface{} {
			j := jsiiProxy_IStandardRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_securityhub.InsightReference",
		reflect.TypeOf((*InsightReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_securityhub.OrganizationConfigurationReference",
		reflect.TypeOf((*OrganizationConfigurationReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_securityhub.PolicyAssociationReference",
		reflect.TypeOf((*PolicyAssociationReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_securityhub.ProductSubscriptionReference",
		reflect.TypeOf((*ProductSubscriptionReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_securityhub.SecurityControlReference",
		reflect.TypeOf((*SecurityControlReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_securityhub.StandardReference",
		reflect.TypeOf((*StandardReference)(nil)).Elem(),
	)
}
