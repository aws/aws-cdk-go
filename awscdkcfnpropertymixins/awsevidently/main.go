package awsevidently

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnExperimentMixinProps",
		reflect.TypeOf((*CfnExperimentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnExperimentPropsMixin",
		reflect.TypeOf((*CfnExperimentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnExperimentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnExperimentPropsMixin.MetricGoalObjectProperty",
		reflect.TypeOf((*CfnExperimentPropsMixin_MetricGoalObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnExperimentPropsMixin.OnlineAbConfigObjectProperty",
		reflect.TypeOf((*CfnExperimentPropsMixin_OnlineAbConfigObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnExperimentPropsMixin.RunningStatusObjectProperty",
		reflect.TypeOf((*CfnExperimentPropsMixin_RunningStatusObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnExperimentPropsMixin.TreatmentObjectProperty",
		reflect.TypeOf((*CfnExperimentPropsMixin_TreatmentObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnExperimentPropsMixin.TreatmentToWeightProperty",
		reflect.TypeOf((*CfnExperimentPropsMixin_TreatmentToWeightProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnFeatureMixinProps",
		reflect.TypeOf((*CfnFeatureMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnFeaturePropsMixin",
		reflect.TypeOf((*CfnFeaturePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFeaturePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnFeaturePropsMixin.EntityOverrideProperty",
		reflect.TypeOf((*CfnFeaturePropsMixin_EntityOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnFeaturePropsMixin.VariationObjectProperty",
		reflect.TypeOf((*CfnFeaturePropsMixin_VariationObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnLaunchMixinProps",
		reflect.TypeOf((*CfnLaunchMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnLaunchPropsMixin",
		reflect.TypeOf((*CfnLaunchPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLaunchPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnLaunchPropsMixin.ExecutionStatusObjectProperty",
		reflect.TypeOf((*CfnLaunchPropsMixin_ExecutionStatusObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnLaunchPropsMixin.GroupToWeightProperty",
		reflect.TypeOf((*CfnLaunchPropsMixin_GroupToWeightProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnLaunchPropsMixin.LaunchGroupObjectProperty",
		reflect.TypeOf((*CfnLaunchPropsMixin_LaunchGroupObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnLaunchPropsMixin.MetricDefinitionObjectProperty",
		reflect.TypeOf((*CfnLaunchPropsMixin_MetricDefinitionObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnLaunchPropsMixin.SegmentOverrideProperty",
		reflect.TypeOf((*CfnLaunchPropsMixin_SegmentOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnLaunchPropsMixin.StepConfigProperty",
		reflect.TypeOf((*CfnLaunchPropsMixin_StepConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnProjectMixinProps",
		reflect.TypeOf((*CfnProjectMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnProjectPropsMixin",
		reflect.TypeOf((*CfnProjectPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnProjectPropsMixin.AppConfigResourceObjectProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_AppConfigResourceObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnProjectPropsMixin.DataDeliveryObjectProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_DataDeliveryObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnProjectPropsMixin.S3DestinationProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_S3DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnSegmentMixinProps",
		reflect.TypeOf((*CfnSegmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_evidently.CfnSegmentPropsMixin",
		reflect.TypeOf((*CfnSegmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSegmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
