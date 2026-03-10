package previewawsfisevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_fis.events.FISExperimentStateChange",
		reflect.TypeOf((*FISExperimentStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_FISExperimentStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fis.events.FISExperimentStateChange.FISExperimentStateChangeProps",
		reflect.TypeOf((*FISExperimentStateChange_FISExperimentStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fis.events.FISExperimentStateChange.NewState",
		reflect.TypeOf((*FISExperimentStateChange_NewState)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fis.events.FISExperimentStateChange.OldState",
		reflect.TypeOf((*FISExperimentStateChange_OldState)(nil)).Elem(),
	)
}
