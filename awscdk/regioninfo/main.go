package regioninfo

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.region_info.Default",
		reflect.TypeOf((*Default)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Default{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.region_info.Fact",
		reflect.TypeOf((*Fact)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Fact{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.region_info.FactName",
		reflect.TypeOf((*FactName)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_FactName{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.region_info.IFact",
		reflect.TypeOf((*IFact)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IFact{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.region_info.RegionInfo",
		reflect.TypeOf((*RegionInfo)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "adotLambdaLayerArn", GoMethod: "AdotLambdaLayerArn"},
			_jsii_.MemberProperty{JsiiProperty: "appMeshRepositoryAccount", GoGetter: "AppMeshRepositoryAccount"},
			_jsii_.MemberProperty{JsiiProperty: "cdkMetadataResourceAvailable", GoGetter: "CdkMetadataResourceAvailable"},
			_jsii_.MemberMethod{JsiiMethod: "cloudwatchLambdaInsightsArn", GoMethod: "CloudwatchLambdaInsightsArn"},
			_jsii_.MemberProperty{JsiiProperty: "dlcRepositoryAccount", GoGetter: "DlcRepositoryAccount"},
			_jsii_.MemberProperty{JsiiProperty: "domainSuffix", GoGetter: "DomainSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "ebsEnvEndpointHostedZoneId", GoGetter: "EbsEnvEndpointHostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "elbv2Account", GoGetter: "Elbv2Account"},
			_jsii_.MemberProperty{JsiiProperty: "firehoseCidrBlock", GoGetter: "FirehoseCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "isOptInRegion", GoGetter: "IsOptInRegion"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "partition", GoGetter: "Partition"},
			_jsii_.MemberProperty{JsiiProperty: "s3StaticWebsiteEndpoint", GoGetter: "S3StaticWebsiteEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "s3StaticWebsiteHostedZoneId", GoGetter: "S3StaticWebsiteHostedZoneId"},
			_jsii_.MemberMethod{JsiiMethod: "servicePrincipal", GoMethod: "ServicePrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointServiceNamePrefix", GoGetter: "VpcEndpointServiceNamePrefix"},
		},
		func() interface{} {
			return &jsiiProxy_RegionInfo{}
		},
	)
}
