package regioninfo

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Information pertaining to an AWS region.
//
// Example:
//   // Get the information for "eu-west-1":
//   region := regionInfo.regionInfo.get(jsii.String("eu-west-1"))
//
//   // Access attributes:
//   region.s3StaticWebsiteEndpoint // s3-website-eu-west-1.amazonaws.com
//   region.servicePrincipal(jsii.String("logs.amazonaws.com"))
//
type RegionInfo interface {
	// The ID of the AWS account that owns the public ECR repository that contains the AWS App Mesh Envoy Proxy images in a given region.
	AppMeshRepositoryAccount() *string
	// Whether the `AWS::CDK::Metadata` CloudFormation Resource is available in this region or not.
	CdkMetadataResourceAvailable() *bool
	// The ID of the AWS account that owns the public ECR repository containing the AWS Deep Learning Containers images in this region.
	DlcRepositoryAccount() *string
	// The domain name suffix (e.g: amazonaws.com) for this region.
	DomainSuffix() *string
	// The hosted zone ID used by Route 53 to alias a EBS environment endpoint in this region (e.g: Z2O1EMRO9K5GLX).
	EbsEnvEndpointHostedZoneId() *string
	// The account ID for ELBv2 in this region.
	Elbv2Account() *string
	// The CIDR block used by Kinesis Data Firehose servers.
	FirehoseCidrBlock() *string
	Name() *string
	// The name of the ARN partition for this region (e.g: aws).
	Partition() *string
	// The endpoint used by S3 static website hosting in this region (e.g: s3-static-website-us-east-1.amazonaws.com).
	S3StaticWebsiteEndpoint() *string
	// The hosted zone ID used by Route 53 to alias a S3 static website in this region (e.g: Z2O1EMRO9K5GLX).
	S3StaticWebsiteHostedZoneId() *string
	// The prefix for VPC Endpoint Service names, cn.com.amazonaws.vpce for China regions, com.amazonaws.vpce otherwise.
	VpcEndpointServiceNamePrefix() *string
	// The ARN of the CloudWatch Lambda Insights extension, for the given version.
	CloudwatchLambdaInsightsArn(insightsVersion *string, architecture *string) *string
	// The name of the service principal for a given service in this region.
	ServicePrincipal(service *string) *string
}

// The jsii proxy struct for RegionInfo
type jsiiProxy_RegionInfo struct {
	_ byte // padding
}

func (j *jsiiProxy_RegionInfo) AppMeshRepositoryAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appMeshRepositoryAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegionInfo) CdkMetadataResourceAvailable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"cdkMetadataResourceAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegionInfo) DlcRepositoryAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dlcRepositoryAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegionInfo) DomainSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegionInfo) EbsEnvEndpointHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEnvEndpointHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegionInfo) Elbv2Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elbv2Account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegionInfo) FirehoseCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegionInfo) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegionInfo) Partition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegionInfo) S3StaticWebsiteEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3StaticWebsiteEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegionInfo) S3StaticWebsiteHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3StaticWebsiteHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegionInfo) VpcEndpointServiceNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointServiceNamePrefix",
		&returns,
	)
	return returns
}


// Obtain region info for a given region name.
func RegionInfo_Get(name *string) RegionInfo {
	_init_.Initialize()

	var returns RegionInfo

	_jsii_.StaticInvoke(
		"aws-cdk-lib.region_info.RegionInfo",
		"get",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Retrieves a collection of all fact values for all regions, limited to some partitions.
//
// Returns: a mapping with AWS region codes as the keys,
// and the fact in the given region as the value for that key.
func RegionInfo_LimitedRegionMap(factName *string, partitions *[]*string) *map[string]*string {
	_init_.Initialize()

	var returns *map[string]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.region_info.RegionInfo",
		"limitedRegionMap",
		[]interface{}{factName, partitions},
		&returns,
	)

	return returns
}

// Retrieves a collection of all fact values for all regions that fact is defined in.
//
// Returns: a mapping with AWS region codes as the keys,
// and the fact in the given region as the value for that key.
func RegionInfo_RegionMap(factName *string) *map[string]*string {
	_init_.Initialize()

	var returns *map[string]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.region_info.RegionInfo",
		"regionMap",
		[]interface{}{factName},
		&returns,
	)

	return returns
}

func RegionInfo_Regions() *[]RegionInfo {
	_init_.Initialize()
	var returns *[]RegionInfo
	_jsii_.StaticGet(
		"aws-cdk-lib.region_info.RegionInfo",
		"regions",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RegionInfo) CloudwatchLambdaInsightsArn(insightsVersion *string, architecture *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"cloudwatchLambdaInsightsArn",
		[]interface{}{insightsVersion, architecture},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegionInfo) ServicePrincipal(service *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"servicePrincipal",
		[]interface{}{service},
		&returns,
	)

	return returns
}

