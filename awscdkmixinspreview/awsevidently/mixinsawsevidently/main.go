package mixinsawsevidently

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnExperimentMixinProps",
		reflect.TypeOf((*CfnExperimentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnExperimentPropsMixin",
		reflect.TypeOf((*CfnExperimentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnExperimentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnExperimentPropsMixin.MetricGoalObjectProperty",
		reflect.TypeOf((*CfnExperimentPropsMixin_MetricGoalObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnExperimentPropsMixin.OnlineAbConfigObjectProperty",
		reflect.TypeOf((*CfnExperimentPropsMixin_OnlineAbConfigObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnExperimentPropsMixin.RunningStatusObjectProperty",
		reflect.TypeOf((*CfnExperimentPropsMixin_RunningStatusObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnExperimentPropsMixin.TreatmentObjectProperty",
		reflect.TypeOf((*CfnExperimentPropsMixin_TreatmentObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnExperimentPropsMixin.TreatmentToWeightProperty",
		reflect.TypeOf((*CfnExperimentPropsMixin_TreatmentToWeightProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnFeatureMixinProps",
		reflect.TypeOf((*CfnFeatureMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnFeaturePropsMixin",
		reflect.TypeOf((*CfnFeaturePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFeaturePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnFeaturePropsMixin.EntityOverrideProperty",
		reflect.TypeOf((*CfnFeaturePropsMixin_EntityOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnFeaturePropsMixin.VariationObjectProperty",
		reflect.TypeOf((*CfnFeaturePropsMixin_VariationObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnLaunchMixinProps",
		reflect.TypeOf((*CfnLaunchMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnLaunchPropsMixin",
		reflect.TypeOf((*CfnLaunchPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLaunchPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnLaunchPropsMixin.ExecutionStatusObjectProperty",
		reflect.TypeOf((*CfnLaunchPropsMixin_ExecutionStatusObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnLaunchPropsMixin.GroupToWeightProperty",
		reflect.TypeOf((*CfnLaunchPropsMixin_GroupToWeightProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnLaunchPropsMixin.LaunchGroupObjectProperty",
		reflect.TypeOf((*CfnLaunchPropsMixin_LaunchGroupObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnLaunchPropsMixin.MetricDefinitionObjectProperty",
		reflect.TypeOf((*CfnLaunchPropsMixin_MetricDefinitionObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnLaunchPropsMixin.SegmentOverrideProperty",
		reflect.TypeOf((*CfnLaunchPropsMixin_SegmentOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnLaunchPropsMixin.StepConfigProperty",
		reflect.TypeOf((*CfnLaunchPropsMixin_StepConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnProjectMixinProps",
		reflect.TypeOf((*CfnProjectMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnProjectPropsMixin",
		reflect.TypeOf((*CfnProjectPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnProjectPropsMixin.AppConfigResourceObjectProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_AppConfigResourceObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnProjectPropsMixin.DataDeliveryObjectProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_DataDeliveryObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnProjectPropsMixin.S3DestinationProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_S3DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnSegmentMixinProps",
		reflect.TypeOf((*CfnSegmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnSegmentPropsMixin",
		reflect.TypeOf((*CfnSegmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSegmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
