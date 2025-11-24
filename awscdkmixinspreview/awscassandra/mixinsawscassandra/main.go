package mixinsawscassandra

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnKeyspaceMixinProps",
		reflect.TypeOf((*CfnKeyspaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnKeyspacePropsMixin",
		reflect.TypeOf((*CfnKeyspacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnKeyspacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnKeyspacePropsMixin.ReplicationSpecificationProperty",
		reflect.TypeOf((*CfnKeyspacePropsMixin_ReplicationSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTableMixinProps",
		reflect.TypeOf((*CfnTableMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin",
		reflect.TypeOf((*CfnTablePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTablePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin.AutoScalingSettingProperty",
		reflect.TypeOf((*CfnTablePropsMixin_AutoScalingSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin.AutoScalingSpecificationProperty",
		reflect.TypeOf((*CfnTablePropsMixin_AutoScalingSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin.BillingModeProperty",
		reflect.TypeOf((*CfnTablePropsMixin_BillingModeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin.CdcSpecificationProperty",
		reflect.TypeOf((*CfnTablePropsMixin_CdcSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin.ClusteringKeyColumnProperty",
		reflect.TypeOf((*CfnTablePropsMixin_ClusteringKeyColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin.ColumnProperty",
		reflect.TypeOf((*CfnTablePropsMixin_ColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin.EncryptionSpecificationProperty",
		reflect.TypeOf((*CfnTablePropsMixin_EncryptionSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin.ProvisionedThroughputProperty",
		reflect.TypeOf((*CfnTablePropsMixin_ProvisionedThroughputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin.ReplicaSpecificationProperty",
		reflect.TypeOf((*CfnTablePropsMixin_ReplicaSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin.ScalingPolicyProperty",
		reflect.TypeOf((*CfnTablePropsMixin_ScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin.TargetTrackingScalingPolicyConfigurationProperty",
		reflect.TypeOf((*CfnTablePropsMixin_TargetTrackingScalingPolicyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTypeMixinProps",
		reflect.TypeOf((*CfnTypeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTypePropsMixin",
		reflect.TypeOf((*CfnTypePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTypePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTypePropsMixin.FieldProperty",
		reflect.TypeOf((*CfnTypePropsMixin_FieldProperty)(nil)).Elem(),
	)
}
