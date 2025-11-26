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
		"@aws-cdk/mixins-preview.aws_networkmanager.events.CoreNetworkEvents.NetworkManagerPolicyUpdate",
		reflect.TypeOf((*CoreNetworkEvents_NetworkManagerPolicyUpdate)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CoreNetworkEvents_NetworkManagerPolicyUpdate{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.CoreNetworkEvents.NetworkManagerPolicyUpdate.NetworkManagerPolicyUpdateProps",
		reflect.TypeOf((*CoreNetworkEvents_NetworkManagerPolicyUpdate_NetworkManagerPolicyUpdateProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.CoreNetworkEvents.NetworkManagerSegmentUpdate",
		reflect.TypeOf((*CoreNetworkEvents_NetworkManagerSegmentUpdate)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CoreNetworkEvents_NetworkManagerSegmentUpdate{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.CoreNetworkEvents.NetworkManagerSegmentUpdate.NetworkManagerSegmentUpdateProps",
		reflect.TypeOf((*CoreNetworkEvents_NetworkManagerSegmentUpdate_NetworkManagerSegmentUpdateProps)(nil)).Elem(),
	)
}
