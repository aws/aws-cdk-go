// The CDK Construct Library for Amazon EventBridge Pipes Sources
package awscdkpipessourcesalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-sources-alpha.SqsSource",
		reflect.TypeOf((*SqsSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "sourceArn", GoGetter: "SourceArn"},
		},
		func() interface{} {
			j := jsiiProxy_SqsSource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-sources-alpha.SqsSourceParameters",
		reflect.TypeOf((*SqsSourceParameters)(nil)).Elem(),
	)
}
