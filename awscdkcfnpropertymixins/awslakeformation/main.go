package awslakeformation

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataCellsFilterMixinProps",
		reflect.TypeOf((*CfnDataCellsFilterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataCellsFilterPropsMixin",
		reflect.TypeOf((*CfnDataCellsFilterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataCellsFilterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataCellsFilterPropsMixin.ColumnWildcardProperty",
		reflect.TypeOf((*CfnDataCellsFilterPropsMixin_ColumnWildcardProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataCellsFilterPropsMixin.RowFilterProperty",
		reflect.TypeOf((*CfnDataCellsFilterPropsMixin_RowFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataLakeSettingsMixinProps",
		reflect.TypeOf((*CfnDataLakeSettingsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataLakeSettingsPropsMixin",
		reflect.TypeOf((*CfnDataLakeSettingsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataLakeSettingsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataLakeSettingsPropsMixin.DataLakePrincipalProperty",
		reflect.TypeOf((*CfnDataLakeSettingsPropsMixin_DataLakePrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataLakeSettingsPropsMixin.PrincipalPermissionsProperty",
		reflect.TypeOf((*CfnDataLakeSettingsPropsMixin_PrincipalPermissionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPermissionsMixinProps",
		reflect.TypeOf((*CfnPermissionsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPermissionsPropsMixin",
		reflect.TypeOf((*CfnPermissionsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPermissionsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPermissionsPropsMixin.ColumnWildcardProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_ColumnWildcardProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPermissionsPropsMixin.DataLakePrincipalProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_DataLakePrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPermissionsPropsMixin.DataLocationResourceProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_DataLocationResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPermissionsPropsMixin.DatabaseResourceProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_DatabaseResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPermissionsPropsMixin.ResourceProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_ResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPermissionsPropsMixin.TableResourceProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_TableResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPermissionsPropsMixin.TableWildcardProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_TableWildcardProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPermissionsPropsMixin.TableWithColumnsResourceProperty",
		reflect.TypeOf((*CfnPermissionsPropsMixin_TableWithColumnsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPrincipalPermissionsMixinProps",
		reflect.TypeOf((*CfnPrincipalPermissionsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPrincipalPermissionsPropsMixin",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPrincipalPermissionsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPrincipalPermissionsPropsMixin.ColumnWildcardProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_ColumnWildcardProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPrincipalPermissionsPropsMixin.DataCellsFilterResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_DataCellsFilterResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPrincipalPermissionsPropsMixin.DataLakePrincipalProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_DataLakePrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPrincipalPermissionsPropsMixin.DataLocationResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_DataLocationResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPrincipalPermissionsPropsMixin.DatabaseResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_DatabaseResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPrincipalPermissionsPropsMixin.LFTagKeyResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_LFTagKeyResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPrincipalPermissionsPropsMixin.LFTagPolicyResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_LFTagPolicyResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPrincipalPermissionsPropsMixin.LFTagProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_LFTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPrincipalPermissionsPropsMixin.ResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_ResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPrincipalPermissionsPropsMixin.TableResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_TableResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnPrincipalPermissionsPropsMixin.TableWithColumnsResourceProperty",
		reflect.TypeOf((*CfnPrincipalPermissionsPropsMixin_TableWithColumnsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnResourceMixinProps",
		reflect.TypeOf((*CfnResourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnResourcePropsMixin",
		reflect.TypeOf((*CfnResourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnTagAssociationMixinProps",
		reflect.TypeOf((*CfnTagAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnTagAssociationPropsMixin",
		reflect.TypeOf((*CfnTagAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTagAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnTagAssociationPropsMixin.DatabaseResourceProperty",
		reflect.TypeOf((*CfnTagAssociationPropsMixin_DatabaseResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnTagAssociationPropsMixin.LFTagPairProperty",
		reflect.TypeOf((*CfnTagAssociationPropsMixin_LFTagPairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnTagAssociationPropsMixin.ResourceProperty",
		reflect.TypeOf((*CfnTagAssociationPropsMixin_ResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnTagAssociationPropsMixin.TableResourceProperty",
		reflect.TypeOf((*CfnTagAssociationPropsMixin_TableResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnTagAssociationPropsMixin.TableWithColumnsResourceProperty",
		reflect.TypeOf((*CfnTagAssociationPropsMixin_TableWithColumnsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnTagMixinProps",
		reflect.TypeOf((*CfnTagMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnTagPropsMixin",
		reflect.TypeOf((*CfnTagPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTagPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
