package regioninfo

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// All standardized fact names.
//
// Example:
//   type myFact struct {
//   	region
//   	name
//   	value
//   }
//
//   regionInfo.fact.register(NewMyFact())
//
type FactName interface {
}

// The jsii proxy struct for FactName
type jsiiProxy_FactName struct {
	_ byte // padding
}

func NewFactName() FactName {
	_init_.Initialize()

	j := jsiiProxy_FactName{}

	_jsii_.Create(
		"aws-cdk-lib.region_info.FactName",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewFactName_Override(f FactName) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.region_info.FactName",
		nil, // no parameters
		f,
	)
}

// The ARN of CloudWatch Lambda Insights for a version (e.g. 1.0.98.0).
func FactName_CloudwatchLambdaInsightsVersion(version *string, arch *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.region_info.FactName",
		"cloudwatchLambdaInsightsVersion",
		[]interface{}{version, arch},
		&returns,
	)

	return returns
}

// The name of the regional service principal for a given service.
func FactName_ServicePrincipal(service *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.region_info.FactName",
		"servicePrincipal",
		[]interface{}{service},
		&returns,
	)

	return returns
}

func FactName_APPMESH_ECR_ACCOUNT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.region_info.FactName",
		"APPMESH_ECR_ACCOUNT",
		&returns,
	)
	return returns
}

func FactName_CDK_METADATA_RESOURCE_AVAILABLE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.region_info.FactName",
		"CDK_METADATA_RESOURCE_AVAILABLE",
		&returns,
	)
	return returns
}

func FactName_DLC_REPOSITORY_ACCOUNT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.region_info.FactName",
		"DLC_REPOSITORY_ACCOUNT",
		&returns,
	)
	return returns
}

func FactName_DOMAIN_SUFFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.region_info.FactName",
		"DOMAIN_SUFFIX",
		&returns,
	)
	return returns
}

func FactName_EBS_ENV_ENDPOINT_HOSTED_ZONE_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.region_info.FactName",
		"EBS_ENV_ENDPOINT_HOSTED_ZONE_ID",
		&returns,
	)
	return returns
}

func FactName_ELBV2_ACCOUNT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.region_info.FactName",
		"ELBV2_ACCOUNT",
		&returns,
	)
	return returns
}

func FactName_FIREHOSE_CIDR_BLOCK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.region_info.FactName",
		"FIREHOSE_CIDR_BLOCK",
		&returns,
	)
	return returns
}

func FactName_PARTITION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.region_info.FactName",
		"PARTITION",
		&returns,
	)
	return returns
}

func FactName_S3_STATIC_WEBSITE_ENDPOINT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.region_info.FactName",
		"S3_STATIC_WEBSITE_ENDPOINT",
		&returns,
	)
	return returns
}

func FactName_S3_STATIC_WEBSITE_ZONE_53_HOSTED_ZONE_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.region_info.FactName",
		"S3_STATIC_WEBSITE_ZONE_53_HOSTED_ZONE_ID",
		&returns,
	)
	return returns
}

func FactName_VPC_ENDPOINT_SERVICE_NAME_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.region_info.FactName",
		"VPC_ENDPOINT_SERVICE_NAME_PREFIX",
		&returns,
	)
	return returns
}

