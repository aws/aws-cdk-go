// The CDK Construct Library for EC2 Image Builder
package awsimagebuilderalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.HttpTokens",
		reflect.TypeOf((*HttpTokens)(nil)).Elem(),
		map[string]interface{}{
			"OPTIONAL": HttpTokens_OPTIONAL,
			"REQUIRED": HttpTokens_REQUIRED,
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-imagebuilder-alpha.IInfrastructureConfiguration",
		reflect.TypeOf((*IInfrastructureConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureConfigurationArn", GoGetter: "InfrastructureConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureConfigurationName", GoGetter: "InfrastructureConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IInfrastructureConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.InfrastructureConfiguration",
		reflect.TypeOf((*InfrastructureConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureConfigurationArn", GoGetter: "InfrastructureConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureConfigurationName", GoGetter: "InfrastructureConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfile", GoGetter: "InstanceProfile"},
			_jsii_.MemberProperty{JsiiProperty: "logBucket", GoGetter: "LogBucket"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_InfrastructureConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInfrastructureConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.InfrastructureConfigurationLogging",
		reflect.TypeOf((*InfrastructureConfigurationLogging)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.InfrastructureConfigurationProps",
		reflect.TypeOf((*InfrastructureConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.Tenancy",
		reflect.TypeOf((*Tenancy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": Tenancy_DEFAULT,
			"DEDICATED": Tenancy_DEDICATED,
			"HOST": Tenancy_HOST,
		},
	)
}
