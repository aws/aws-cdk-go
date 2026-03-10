package previewawsappmeshevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_appmesh.events.AWSServiceEventViaCloudTrail",
		reflect.TypeOf((*AWSServiceEventViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AWSServiceEventViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_appmesh.events.AWSServiceEventViaCloudTrail.AWSServiceEventViaCloudTrailProps",
		reflect.TypeOf((*AWSServiceEventViaCloudTrail_AWSServiceEventViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_appmesh.events.AWSServiceEventViaCloudTrail.ServiceEventDetails",
		reflect.TypeOf((*AWSServiceEventViaCloudTrail_ServiceEventDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_appmesh.events.AWSServiceEventViaCloudTrail.UserIdentity",
		reflect.TypeOf((*AWSServiceEventViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
}
