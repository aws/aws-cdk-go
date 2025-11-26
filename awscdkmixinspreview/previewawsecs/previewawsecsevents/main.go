package previewawsecsevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents",
		reflect.TypeOf((*ClusterEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "awsAPICallViaCloudTrailPattern", GoMethod: "AwsAPICallViaCloudTrailPattern"},
			_jsii_.MemberMethod{JsiiMethod: "eCSContainerInstanceStateChangePattern", GoMethod: "ECSContainerInstanceStateChangePattern"},
			_jsii_.MemberMethod{JsiiMethod: "eCSServiceActionPattern", GoMethod: "ECSServiceActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "eCSTaskStateChangePattern", GoMethod: "ECSTaskStateChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ClusterEvents_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.Attributes",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.AwsvpcConfiguration",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_AwsvpcConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.NetworkConfiguration",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_NetworkConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.Overrides",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_Overrides)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.Overrides1",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_Overrides1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.Overrides1Item",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_Overrides1Item)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.OverridesItem",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_OverridesItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.RequestParametersItem",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_RequestParametersItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.RequestParametersItem1",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_RequestParametersItem1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.ResponseElements",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_ResponseElements)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.ResponseElementsItem",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_ResponseElementsItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.ResponseElementsItemItem",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_ResponseElementsItemItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.ResponseElementsItemItem1",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_ResponseElementsItemItem1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.ResponseElementsItemItemItem",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_ResponseElementsItemItemItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.SessionContext",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_SessionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.SessionIssuer",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_SessionIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail.UserIdentity",
		reflect.TypeOf((*ClusterEvents_AWSAPICallViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSContainerInstanceStateChange",
		reflect.TypeOf((*ClusterEvents_ECSContainerInstanceStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ClusterEvents_ECSContainerInstanceStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSContainerInstanceStateChange.AttachmentDetails",
		reflect.TypeOf((*ClusterEvents_ECSContainerInstanceStateChange_AttachmentDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSContainerInstanceStateChange.AttributesDetails",
		reflect.TypeOf((*ClusterEvents_ECSContainerInstanceStateChange_AttributesDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSContainerInstanceStateChange.DetailsItems",
		reflect.TypeOf((*ClusterEvents_ECSContainerInstanceStateChange_DetailsItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSContainerInstanceStateChange.ECSContainerInstanceStateChangeProps",
		reflect.TypeOf((*ClusterEvents_ECSContainerInstanceStateChange_ECSContainerInstanceStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSContainerInstanceStateChange.ResourceDetails",
		reflect.TypeOf((*ClusterEvents_ECSContainerInstanceStateChange_ResourceDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSContainerInstanceStateChange.VersionInfo",
		reflect.TypeOf((*ClusterEvents_ECSContainerInstanceStateChange_VersionInfo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSServiceAction",
		reflect.TypeOf((*ClusterEvents_ECSServiceAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ClusterEvents_ECSServiceAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSServiceAction.ECSServiceActionProps",
		reflect.TypeOf((*ClusterEvents_ECSServiceAction_ECSServiceActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSTaskStateChange",
		reflect.TypeOf((*ClusterEvents_ECSTaskStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ClusterEvents_ECSTaskStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSTaskStateChange.AttachmentDetails",
		reflect.TypeOf((*ClusterEvents_ECSTaskStateChange_AttachmentDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSTaskStateChange.AttributesDetails",
		reflect.TypeOf((*ClusterEvents_ECSTaskStateChange_AttributesDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSTaskStateChange.ContainerDetails",
		reflect.TypeOf((*ClusterEvents_ECSTaskStateChange_ContainerDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSTaskStateChange.Details",
		reflect.TypeOf((*ClusterEvents_ECSTaskStateChange_Details)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSTaskStateChange.ECSTaskStateChangeProps",
		reflect.TypeOf((*ClusterEvents_ECSTaskStateChange_ECSTaskStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSTaskStateChange.NetworkBindingDetails",
		reflect.TypeOf((*ClusterEvents_ECSTaskStateChange_NetworkBindingDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSTaskStateChange.NetworkInterfaceDetails",
		reflect.TypeOf((*ClusterEvents_ECSTaskStateChange_NetworkInterfaceDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSTaskStateChange.Overrides",
		reflect.TypeOf((*ClusterEvents_ECSTaskStateChange_Overrides)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSTaskStateChange.OverridesItem",
		reflect.TypeOf((*ClusterEvents_ECSTaskStateChange_OverridesItem)(nil)).Elem(),
	)
}
