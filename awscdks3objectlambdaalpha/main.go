// The CDK Construct Library for AWS::S3ObjectLambda
package awscdks3objectlambdaalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-s3objectlambda-alpha.AccessPoint",
		reflect.TypeOf((*AccessPoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessPointArn", GoGetter: "AccessPointArn"},
			_jsii_.MemberProperty{JsiiProperty: "accessPointCreationDate", GoGetter: "AccessPointCreationDate"},
			_jsii_.MemberProperty{JsiiProperty: "accessPointName", GoGetter: "AccessPointName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "regionalDomainName", GoGetter: "RegionalDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "virtualHostedUrlForObject", GoMethod: "VirtualHostedUrlForObject"},
		},
		func() interface{} {
			j := jsiiProxy_AccessPoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAccessPoint)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-s3objectlambda-alpha.AccessPointAttributes",
		reflect.TypeOf((*AccessPointAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-s3objectlambda-alpha.AccessPointProps",
		reflect.TypeOf((*AccessPointProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-s3objectlambda-alpha.IAccessPoint",
		reflect.TypeOf((*IAccessPoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessPointArn", GoGetter: "AccessPointArn"},
			_jsii_.MemberProperty{JsiiProperty: "accessPointCreationDate", GoGetter: "AccessPointCreationDate"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "regionalDomainName", GoGetter: "RegionalDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "virtualHostedUrlForObject", GoMethod: "VirtualHostedUrlForObject"},
		},
		func() interface{} {
			j := jsiiProxy_IAccessPoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
}
