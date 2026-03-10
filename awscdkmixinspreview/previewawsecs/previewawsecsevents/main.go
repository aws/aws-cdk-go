package previewawsecsevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.Attributes",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.AwsvpcConfiguration",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_AwsvpcConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.NetworkConfiguration",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_NetworkConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.Overrides",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_Overrides)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.Overrides1",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_Overrides1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.Overrides1Item",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_Overrides1Item)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.OverridesItem",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_OverridesItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.RequestParametersItem",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RequestParametersItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.RequestParametersItem1",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RequestParametersItem1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.ResponseElements",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_ResponseElements)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.ResponseElementsItem",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_ResponseElementsItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.ResponseElementsItemItem",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_ResponseElementsItemItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.ResponseElementsItemItem1",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_ResponseElementsItemItem1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.ResponseElementsItemItemItem",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_ResponseElementsItemItemItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.SessionContext",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_SessionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.SessionIssuer",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_SessionIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.AWSAPICallViaCloudTrail.UserIdentity",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents",
		reflect.TypeOf((*ClusterEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "awsAPICallViaCloudTrailPattern", GoMethod: "AwsAPICallViaCloudTrailPattern"},
			_jsii_.MemberMethod{JsiiMethod: "ecsContainerInstanceStateChangePattern", GoMethod: "EcsContainerInstanceStateChangePattern"},
			_jsii_.MemberMethod{JsiiMethod: "ecsServiceActionPattern", GoMethod: "EcsServiceActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "ecsTaskStateChangePattern", GoMethod: "EcsTaskStateChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSContainerInstanceStateChange",
		reflect.TypeOf((*ECSContainerInstanceStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ECSContainerInstanceStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSContainerInstanceStateChange.AttachmentDetails",
		reflect.TypeOf((*ECSContainerInstanceStateChange_AttachmentDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSContainerInstanceStateChange.AttributesDetails",
		reflect.TypeOf((*ECSContainerInstanceStateChange_AttributesDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSContainerInstanceStateChange.DetailsItems",
		reflect.TypeOf((*ECSContainerInstanceStateChange_DetailsItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSContainerInstanceStateChange.ECSContainerInstanceStateChangeProps",
		reflect.TypeOf((*ECSContainerInstanceStateChange_ECSContainerInstanceStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSContainerInstanceStateChange.ResourceDetails",
		reflect.TypeOf((*ECSContainerInstanceStateChange_ResourceDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSContainerInstanceStateChange.VersionInfo",
		reflect.TypeOf((*ECSContainerInstanceStateChange_VersionInfo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSServiceAction",
		reflect.TypeOf((*ECSServiceAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ECSServiceAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSServiceAction.ECSServiceActionProps",
		reflect.TypeOf((*ECSServiceAction_ECSServiceActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSTaskStateChange",
		reflect.TypeOf((*ECSTaskStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ECSTaskStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSTaskStateChange.AttachmentDetails",
		reflect.TypeOf((*ECSTaskStateChange_AttachmentDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSTaskStateChange.AttributesDetails",
		reflect.TypeOf((*ECSTaskStateChange_AttributesDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSTaskStateChange.ContainerDetails",
		reflect.TypeOf((*ECSTaskStateChange_ContainerDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSTaskStateChange.Details",
		reflect.TypeOf((*ECSTaskStateChange_Details)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSTaskStateChange.ECSTaskStateChangeProps",
		reflect.TypeOf((*ECSTaskStateChange_ECSTaskStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSTaskStateChange.NetworkBindingDetails",
		reflect.TypeOf((*ECSTaskStateChange_NetworkBindingDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSTaskStateChange.NetworkInterfaceDetails",
		reflect.TypeOf((*ECSTaskStateChange_NetworkInterfaceDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSTaskStateChange.Overrides",
		reflect.TypeOf((*ECSTaskStateChange_Overrides)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSTaskStateChange.OverridesItem",
		reflect.TypeOf((*ECSTaskStateChange_OverridesItem)(nil)).Elem(),
	)
}
