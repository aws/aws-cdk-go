package interfacesawssupportapp

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_supportapp.AccountAliasReference",
		reflect.TypeOf((*AccountAliasReference)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_supportapp.IAccountAliasRef",
		reflect.TypeOf((*IAccountAliasRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountAliasRef", GoGetter: "AccountAliasRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IAccountAliasRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_supportapp.ISlackChannelConfigurationRef",
		reflect.TypeOf((*ISlackChannelConfigurationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "slackChannelConfigurationRef", GoGetter: "SlackChannelConfigurationRef"},
		},
		func() interface{} {
			j := jsiiProxy_ISlackChannelConfigurationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_supportapp.ISlackWorkspaceConfigurationRef",
		reflect.TypeOf((*ISlackWorkspaceConfigurationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "slackWorkspaceConfigurationRef", GoGetter: "SlackWorkspaceConfigurationRef"},
		},
		func() interface{} {
			j := jsiiProxy_ISlackWorkspaceConfigurationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_supportapp.SlackChannelConfigurationReference",
		reflect.TypeOf((*SlackChannelConfigurationReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_supportapp.SlackWorkspaceConfigurationReference",
		reflect.TypeOf((*SlackWorkspaceConfigurationReference)(nil)).Elem(),
	)
}
