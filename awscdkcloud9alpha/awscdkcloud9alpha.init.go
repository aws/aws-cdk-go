package awscdkcloud9alpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-cloud9-alpha.CloneRepository",
		reflect.TypeOf((*CloneRepository)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "pathComponent", GoGetter: "PathComponent"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryUrl", GoGetter: "RepositoryUrl"},
		},
		func() interface{} {
			return &jsiiProxy_CloneRepository{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-cloud9-alpha.ConnectionType",
		reflect.TypeOf((*ConnectionType)(nil)).Elem(),
		map[string]interface{}{
			"CONNECT_SSH": ConnectionType_CONNECT_SSH,
			"CONNECT_SSM": ConnectionType_CONNECT_SSM,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-cloud9-alpha.Ec2Environment",
		reflect.TypeOf((*Ec2Environment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "ec2EnvironmentArn", GoGetter: "Ec2EnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "ec2EnvironmentName", GoGetter: "Ec2EnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "environmentId", GoGetter: "EnvironmentId"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ideUrl", GoGetter: "IdeUrl"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2Environment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEc2Environment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cloud9-alpha.Ec2EnvironmentProps",
		reflect.TypeOf((*Ec2EnvironmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-cloud9-alpha.IEc2Environment",
		reflect.TypeOf((*IEc2Environment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "ec2EnvironmentArn", GoGetter: "Ec2EnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "ec2EnvironmentName", GoGetter: "Ec2EnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IEc2Environment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-cloud9-alpha.ImageId",
		reflect.TypeOf((*ImageId)(nil)).Elem(),
		map[string]interface{}{
			"AMAZON_LINUX_2": ImageId_AMAZON_LINUX_2,
			"UBUNTU_18_04": ImageId_UBUNTU_18_04,
		},
	)
}
