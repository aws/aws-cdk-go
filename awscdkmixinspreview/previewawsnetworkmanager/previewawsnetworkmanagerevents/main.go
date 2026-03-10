package previewawsnetworkmanagerevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.CoreNetworkEvents",
		reflect.TypeOf((*CoreNetworkEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "networkManagerPolicyUpdatePattern", GoMethod: "NetworkManagerPolicyUpdatePattern"},
			_jsii_.MemberMethod{JsiiMethod: "networkManagerSegmentUpdatePattern", GoMethod: "NetworkManagerSegmentUpdatePattern"},
		},
		func() interface{} {
			return &jsiiProxy_CoreNetworkEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.NetworkManagerPolicyUpdate",
		reflect.TypeOf((*NetworkManagerPolicyUpdate)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_NetworkManagerPolicyUpdate{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.NetworkManagerPolicyUpdate.NetworkManagerPolicyUpdateProps",
		reflect.TypeOf((*NetworkManagerPolicyUpdate_NetworkManagerPolicyUpdateProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.NetworkManagerSegmentUpdate",
		reflect.TypeOf((*NetworkManagerSegmentUpdate)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_NetworkManagerSegmentUpdate{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.NetworkManagerSegmentUpdate.NetworkManagerSegmentUpdateProps",
		reflect.TypeOf((*NetworkManagerSegmentUpdate_NetworkManagerSegmentUpdateProps)(nil)).Elem(),
	)
}
