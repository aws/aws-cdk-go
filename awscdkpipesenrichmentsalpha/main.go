// The CDK Construct Library for Amazon EventBridge Pipes Enrichments
package awscdkpipesenrichmentsalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-enrichments-alpha.ApiDestinationEnrichment",
		reflect.TypeOf((*ApiDestinationEnrichment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "enrichmentArn", GoGetter: "EnrichmentArn"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
		},
		func() interface{} {
			j := jsiiProxy_ApiDestinationEnrichment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaIEnrichment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-enrichments-alpha.ApiDestinationEnrichmentProps",
		reflect.TypeOf((*ApiDestinationEnrichmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-enrichments-alpha.ApiGatewayEnrichment",
		reflect.TypeOf((*ApiGatewayEnrichment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "enrichmentArn", GoGetter: "EnrichmentArn"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
		},
		func() interface{} {
			j := jsiiProxy_ApiGatewayEnrichment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaIEnrichment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-enrichments-alpha.ApiGatewayEnrichmentProps",
		reflect.TypeOf((*ApiGatewayEnrichmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-enrichments-alpha.LambdaEnrichment",
		reflect.TypeOf((*LambdaEnrichment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "enrichmentArn", GoGetter: "EnrichmentArn"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaEnrichment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaIEnrichment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-enrichments-alpha.LambdaEnrichmentProps",
		reflect.TypeOf((*LambdaEnrichmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-enrichments-alpha.StepFunctionsEnrichment",
		reflect.TypeOf((*StepFunctionsEnrichment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "enrichmentArn", GoGetter: "EnrichmentArn"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
		},
		func() interface{} {
			j := jsiiProxy_StepFunctionsEnrichment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaIEnrichment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-enrichments-alpha.StepFunctionsEnrichmentProps",
		reflect.TypeOf((*StepFunctionsEnrichmentProps)(nil)).Elem(),
	)
}
