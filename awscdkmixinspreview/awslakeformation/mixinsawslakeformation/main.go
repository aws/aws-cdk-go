package mixinsawslakeformation

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnDataCellsFilterMixinProps",
		reflect.TypeOf((*CfnDataCellsFilterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnDataCellsFilterPropsMixin",
		reflect.TypeOf((*CfnDataCellsFilterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataCellsFilterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnDataCellsFilterPropsMixin.ColumnWildcardProperty",
		reflect.TypeOf((*CfnDataCellsFilterPropsMixin_ColumnWildcardProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnDataCellsFilterPropsMixin.RowFilterProperty",
		reflect.TypeOf((*CfnDataCellsFilterPropsMixin_RowFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnDataLakeSettingsMixinProps",
		reflect.TypeOf((*CfnDataLakeSettingsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnDataLakeSettingsPropsMixin",
		reflect.TypeOf((*CfnDataLakeSettingsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataLakeSettingsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnDataLakeSettingsPropsMixin.DataLakePrincipalProperty",
		reflect.TypeOf((*CfnDataLakeSettingsPropsMixin_DataLakePrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnDataLakeSettingsPropsMixin.PrincipalPermissionsProperty",
		reflect.TypeOf((*CfnDataLakeSettingsPropsMixin_PrincipalPermissionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPermissionsMixinProps",
		reflect.TypeOf((*CfnPermissionsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPermissionsPropsMixin",
		reflect.TypeOf((*CfnPermissionsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPermissionsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPermissionsPropsMixin.ColumnWildcardProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_ColumnWildcardProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPermissionsPropsMixin.DataLakePrincipalProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_DataLakePrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPermissionsPropsMixin.DataLocationResourceProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_DataLocationResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPermissionsPropsMixin.DatabaseResourceProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_DatabaseResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPermissionsPropsMixin.ResourceProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_ResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPermissionsPropsMixin.TableResourceProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_TableResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPermissionsPropsMixin.TableWildcardProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_TableWildcardProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPermissionsPropsMixin.TableWithColumnsResourceProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_TableWithColumnsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsMixinProps",
		reflect.TypeOf((*CfnPrincipalPermissionsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPrincipalPermissionsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin.ColumnWildcardProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_ColumnWildcardProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin.DataCellsFilterResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_DataCellsFilterResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin.DataLakePrincipalProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_DataLakePrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin.DataLocationResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_DataLocationResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin.DatabaseResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_DatabaseResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin.LFTagKeyResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_LFTagKeyResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin.LFTagPolicyResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_LFTagPolicyResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin.LFTagProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_LFTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin.ResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_ResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin.TableResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_TableResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin.TableWithColumnsResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_TableWithColumnsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnResourceMixinProps",
		reflect.TypeOf((*CfnResourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnResourcePropsMixin",
		reflect.TypeOf((*CfnResourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnTagAssociationMixinProps",
		reflect.TypeOf((*CfnTagAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnTagAssociationPropsMixin",
		reflect.TypeOf((*CfnTagAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTagAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnTagAssociationPropsMixin.DatabaseResourceProperty",
		reflect.TypeOf((*CfnTagAssociationPropsMixin_DatabaseResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnTagAssociationPropsMixin.LFTagPairProperty",
		reflect.TypeOf((*CfnTagAssociationPropsMixin_LFTagPairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnTagAssociationPropsMixin.ResourceProperty",
		reflect.TypeOf((*CfnTagAssociationPropsMixin_ResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnTagAssociationPropsMixin.TableResourceProperty",
		reflect.TypeOf((*CfnTagAssociationPropsMixin_TableResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnTagAssociationPropsMixin.TableWithColumnsResourceProperty",
		reflect.TypeOf((*CfnTagAssociationPropsMixin_TableWithColumnsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnTagMixinProps",
		reflect.TypeOf((*CfnTagMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnTagPropsMixin",
		reflect.TypeOf((*CfnTagPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTagPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
