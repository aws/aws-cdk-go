package regioninfo

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Provides default values for certain regional information points.
// Experimental.
type Default interface {
}

// The jsii proxy struct for Default
type jsiiProxy_Default struct {
	_ byte // padding
}

// Computes a "standard" AWS Service principal for a given service, region and suffix.
//
// This is useful for example when
// you need to compute a service principal name, but you do not have a synthesize-time region literal available (so
// all you have is `{ "Ref": "AWS::Region" }`). This way you get the same defaulting behavior that is normally used
// for built-in data.
// Experimental.
func Default_ServicePrincipal(service *string, region *string, urlSuffix *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.region_info.Default",
		"servicePrincipal",
		[]interface{}{service, region, urlSuffix},
		&returns,
	)

	return returns
}

func Default_VPC_ENDPOINT_SERVICE_NAME_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.region_info.Default",
		"VPC_ENDPOINT_SERVICE_NAME_PREFIX",
		&returns,
	)
	return returns
}

// A database of regional information.
// Experimental.
type Fact interface {
}

// The jsii proxy struct for Fact
type jsiiProxy_Fact struct {
	_ byte // padding
}

// Retrieves a fact from this Fact database.
//
// Returns: the fact value if it is known, and `undefined` otherwise.
// Experimental.
func Fact_Find(region *string, name *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.region_info.Fact",
		"find",
		[]interface{}{region, name},
		&returns,
	)

	return returns
}

// Registers a new fact in this Fact database.
// Experimental.
func Fact_Register(fact IFact, allowReplacing *bool) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.region_info.Fact",
		"register",
		[]interface{}{fact, allowReplacing},
	)
}

// Retrieve a fact from the Fact database.
//
// (retrieval will fail if the specified region or
// fact name does not exist.)
// Experimental.
func Fact_RequireFact(region *string, name *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.region_info.Fact",
		"requireFact",
		[]interface{}{region, name},
		&returns,
	)

	return returns
}

// Removes a fact from the database.
// Experimental.
func Fact_Unregister(region *string, name *string, value *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.region_info.Fact",
		"unregister",
		[]interface{}{region, name, value},
	)
}

func Fact_Regions() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.region_info.Fact",
		"regions",
		&returns,
	)
	return returns
}

// All standardized fact names.
// Experimental.
type FactName interface {
}

// The jsii proxy struct for FactName
type jsiiProxy_FactName struct {
	_ byte // padding
}

// Experimental.
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

// Experimental.
func NewFactName_Override(f FactName) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.region_info.FactName",
		nil, // no parameters
		f,
	)
}

// The name of the regional service principal for a given service.
// Experimental.
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

// A fact that can be registered about a particular region.
// Experimental.
type IFact interface {
	// The name of this fact.
	//
	// Standardized values are provided by the `Facts` class.
	// Experimental.
	Name() *string
	// The region for which this fact applies.
	// Experimental.
	Region() *string
	// The value of this fact.
	// Experimental.
	Value() *string
}

// The jsii proxy for IFact
type jsiiProxy_IFact struct {
	_ byte // padding
}

func (j *jsiiProxy_IFact) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFact) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFact) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

// Information pertaining to an AWS region.
// Experimental.
type RegionInfo interface {
	AppMeshRepositoryAccount() *string
	CdkMetadataResourceAvailable() *bool
	DlcRepositoryAccount() *string
	DomainSuffix() *string
	Elbv2Account() *string
	Name() *string
	Partition() *string
	S3StaticWebsiteEndpoint() *string
	S3StaticWebsiteHostedZoneId() *string
	VpcEndpointServiceNamePrefix() *string
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

func (j *jsiiProxy_RegionInfo) Elbv2Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elbv2Account",
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
// Experimental.
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

// Retrieves a collection of all fact values for all regions that fact is defined in.
//
// Returns: a mapping with AWS region codes as the keys,
// and the fact in the given region as the value for that key
// Experimental.
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

// The name of the service principal for a given service in this region.
// Experimental.
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

