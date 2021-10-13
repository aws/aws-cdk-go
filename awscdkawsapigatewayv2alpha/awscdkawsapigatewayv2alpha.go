// The CDK Construct Library for AWS::APIGatewayv2
package awscdkawsapigatewayv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkawsapigatewayv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdkawsapigatewayv2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Options for the Route with Integration resource.
// Experimental.
type AddRoutesOptions struct {
	// The integration to be configured on this route.
	// Experimental.
	Integration IHttpRouteIntegration `json:"integration"`
	// The path at which all of these routes are configured.
	// Experimental.
	Path *string `json:"path"`
	// The list of OIDC scopes to include in the authorization.
	//
	// These scopes will override the default authorization scopes on the gateway.
	// Set to [] to remove default scopes
	// Experimental.
	AuthorizationScopes *[]*string `json:"authorizationScopes"`
	// Authorizer to be associated to these routes.
	//
	// Use NoneAuthorizer to remove the default authorizer for the api
	// Experimental.
	Authorizer IHttpRouteAuthorizer `json:"authorizer"`
	// The HTTP methods to be configured.
	// Experimental.
	Methods *[]HttpMethod `json:"methods"`
}

// Create a new API mapping for API Gateway API endpoint.
// Experimental.
type ApiMapping interface {
	awscdk.Resource
	IApiMapping
	ApiMappingId() *string
	DomainName() IDomainName
	Env() *awscdk.ResourceEnvironment
	MappingKey() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for ApiMapping
type jsiiProxy_ApiMapping struct {
	internal.Type__awscdkResource
	jsiiProxy_IApiMapping
}

func (j *jsiiProxy_ApiMapping) ApiMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiMappingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiMapping) DomainName() IDomainName {
	var returns IDomainName
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiMapping) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiMapping) MappingKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiMapping) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiMapping) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewApiMapping(scope constructs.Construct, id *string, props *ApiMappingProps) ApiMapping {
	_init_.Initialize()

	j := jsiiProxy_ApiMapping{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.ApiMapping",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApiMapping_Override(a ApiMapping, scope constructs.Construct, id *string, props *ApiMappingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.ApiMapping",
		[]interface{}{scope, id, props},
		a,
	)
}

// import from API ID.
// Experimental.
func ApiMapping_FromApiMappingAttributes(scope constructs.Construct, id *string, attrs *ApiMappingAttributes) IApiMapping {
	_init_.Initialize()

	var returns IApiMapping

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.ApiMapping",
		"fromApiMappingAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.ApiMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ApiMapping_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.ApiMapping",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (a *jsiiProxy_ApiMapping) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (a *jsiiProxy_ApiMapping) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (a *jsiiProxy_ApiMapping) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (a *jsiiProxy_ApiMapping) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_ApiMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The attributes used to import existing ApiMapping.
// Experimental.
type ApiMappingAttributes struct {
	// The API mapping ID.
	// Experimental.
	ApiMappingId *string `json:"apiMappingId"`
}

// Properties used to create the ApiMapping resource.
// Experimental.
type ApiMappingProps struct {
	// The Api to which this mapping is applied.
	// Experimental.
	Api IApi `json:"api"`
	// custom domain name of the mapping target.
	// Experimental.
	DomainName IDomainName `json:"domainName"`
	// Api mapping key.
	//
	// The path where this stage should be mapped to on the domain
	// Experimental.
	ApiMappingKey *string `json:"apiMappingKey"`
	// stage for the ApiMapping resource required for WebSocket API defaults to default stage of an HTTP API.
	// Experimental.
	Stage IStage `json:"stage"`
}

// Payload format version for lambda authorizers.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html
//
// Experimental.
type AuthorizerPayloadVersion string

const (
	AuthorizerPayloadVersion_VERSION_1_0 AuthorizerPayloadVersion = "VERSION_1_0"
	AuthorizerPayloadVersion_VERSION_2_0 AuthorizerPayloadVersion = "VERSION_2_0"
)

// Options used when configuring multiple routes, at once.
//
// The options here are the ones that would be configured for all being set up.
// Experimental.
type BatchHttpRouteOptions struct {
	// The integration to be configured on this route.
	// Experimental.
	Integration IHttpRouteIntegration `json:"integration"`
}

// Supported CORS HTTP methods.
// Experimental.
type CorsHttpMethod string

const (
	CorsHttpMethod_ANY CorsHttpMethod = "ANY"
	CorsHttpMethod_DELETE CorsHttpMethod = "DELETE"
	CorsHttpMethod_GET CorsHttpMethod = "GET"
	CorsHttpMethod_HEAD CorsHttpMethod = "HEAD"
	CorsHttpMethod_OPTIONS CorsHttpMethod = "OPTIONS"
	CorsHttpMethod_PATCH CorsHttpMethod = "PATCH"
	CorsHttpMethod_POST CorsHttpMethod = "POST"
	CorsHttpMethod_PUT CorsHttpMethod = "PUT"
)

// Options for the CORS Configuration.
// Experimental.
type CorsPreflightOptions struct {
	// Specifies whether credentials are included in the CORS request.
	// Experimental.
	AllowCredentials *bool `json:"allowCredentials"`
	// Represents a collection of allowed headers.
	// Experimental.
	AllowHeaders *[]*string `json:"allowHeaders"`
	// Represents a collection of allowed HTTP methods.
	// Experimental.
	AllowMethods *[]CorsHttpMethod `json:"allowMethods"`
	// Represents a collection of allowed origins.
	// Experimental.
	AllowOrigins *[]*string `json:"allowOrigins"`
	// Represents a collection of exposed headers.
	// Experimental.
	ExposeHeaders *[]*string `json:"exposeHeaders"`
	// The duration that the browser should cache preflight request results.
	// Experimental.
	MaxAge awscdk.Duration `json:"maxAge"`
}

// Options for DomainMapping.
// Experimental.
type DomainMappingOptions struct {
	// The domain name for the mapping.
	// Experimental.
	DomainName IDomainName `json:"domainName"`
	// The API mapping key.
	//
	// Leave it undefined for the root path mapping.
	// Experimental.
	MappingKey *string `json:"mappingKey"`
}

// Custom domain resource for the API.
// Experimental.
type DomainName interface {
	awscdk.Resource
	IDomainName
	Env() *awscdk.ResourceEnvironment
	Name() *string
	Node() constructs.Node
	PhysicalName() *string
	RegionalDomainName() *string
	RegionalHostedZoneId() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for DomainName
type jsiiProxy_DomainName struct {
	internal.Type__awscdkResource
	jsiiProxy_IDomainName
}

func (j *jsiiProxy_DomainName) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) RegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) RegionalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDomainName(scope constructs.Construct, id *string, props *DomainNameProps) DomainName {
	_init_.Initialize()

	j := jsiiProxy_DomainName{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.DomainName",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDomainName_Override(d DomainName, scope constructs.Construct, id *string, props *DomainNameProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.DomainName",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import from attributes.
// Experimental.
func DomainName_FromDomainNameAttributes(scope constructs.Construct, id *string, attrs *DomainNameAttributes) IDomainName {
	_init_.Initialize()

	var returns IDomainName

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.DomainName",
		"fromDomainNameAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DomainName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.DomainName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func DomainName_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.DomainName",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (d *jsiiProxy_DomainName) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (d *jsiiProxy_DomainName) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (d *jsiiProxy_DomainName) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (d *jsiiProxy_DomainName) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DomainName) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// custom domain name attributes.
// Experimental.
type DomainNameAttributes struct {
	// domain name string.
	// Experimental.
	Name *string `json:"name"`
	// The domain name associated with the regional endpoint for this custom domain name.
	// Experimental.
	RegionalDomainName *string `json:"regionalDomainName"`
	// The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint.
	// Experimental.
	RegionalHostedZoneId *string `json:"regionalHostedZoneId"`
}

// properties used for creating the DomainName.
// Experimental.
type DomainNameProps struct {
	// The ACM certificate for this domain name.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `json:"certificate"`
	// The custom domain name.
	// Experimental.
	DomainName *string `json:"domainName"`
}

// Create a new API Gateway HTTP API endpoint.
// Experimental.
type HttpApi interface {
	awscdk.Resource
	IApi
	IHttpApi
	ApiEndpoint() *string
	ApiId() *string
	DefaultStage() IHttpStage
	DisableExecuteApiEndpoint() *bool
	Env() *awscdk.ResourceEnvironment
	HttpApiId() *string
	HttpApiName() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	Url() *string
	AddRoutes(options *AddRoutesOptions) *[]HttpRoute
	AddStage(id *string, options *HttpStageOptions) HttpStage
	AddVpcLink(options *VpcLinkProps) VpcLink
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDataProcessed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for HttpApi
type jsiiProxy_HttpApi struct {
	internal.Type__awscdkResource
	jsiiProxy_IApi
	jsiiProxy_IHttpApi
}

func (j *jsiiProxy_HttpApi) ApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) DefaultStage() IHttpStage {
	var returns IHttpStage
	_jsii_.Get(
		j,
		"defaultStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) DisableExecuteApiEndpoint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"disableExecuteApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) HttpApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) HttpApiName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpApiName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Experimental.
func NewHttpApi(scope constructs.Construct, id *string, props *HttpApiProps) HttpApi {
	_init_.Initialize()

	j := jsiiProxy_HttpApi{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpApi_Override(h HttpApi, scope constructs.Construct, id *string, props *HttpApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpApi",
		[]interface{}{scope, id, props},
		h,
	)
}

// Import an existing HTTP API into this CDK app.
// Experimental.
func HttpApi_FromHttpApiAttributes(scope constructs.Construct, id *string, attrs *HttpApiAttributes) IHttpApi {
	_init_.Initialize()

	var returns IHttpApi

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpApi",
		"fromHttpApiAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func HttpApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func HttpApi_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpApi",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add multiple routes that uses the same configuration.
//
// The routes all go to the same path, but for different
// methods.
// Experimental.
func (h *jsiiProxy_HttpApi) AddRoutes(options *AddRoutesOptions) *[]HttpRoute {
	var returns *[]HttpRoute

	_jsii_.Invoke(
		h,
		"addRoutes",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Add a new stage.
// Experimental.
func (h *jsiiProxy_HttpApi) AddStage(id *string, options *HttpStageOptions) HttpStage {
	var returns HttpStage

	_jsii_.Invoke(
		h,
		"addStage",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Add a new VpcLink.
// Experimental.
func (h *jsiiProxy_HttpApi) AddVpcLink(options *VpcLinkProps) VpcLink {
	var returns VpcLink

	_jsii_.Invoke(
		h,
		"addVpcLink",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (h *jsiiProxy_HttpApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		h,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (h *jsiiProxy_HttpApi) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (h *jsiiProxy_HttpApi) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (h *jsiiProxy_HttpApi) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Return the given named metric for this Api Gateway.
// Experimental.
func (h *jsiiProxy_HttpApi) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of client-side errors captured in a given period.
// Experimental.
func (h *jsiiProxy_HttpApi) MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricClientError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the total number API requests in a given period.
// Experimental.
func (h *jsiiProxy_HttpApi) MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the amount of data processed in bytes.
// Experimental.
func (h *jsiiProxy_HttpApi) MetricDataProcessed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricDataProcessed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
// Experimental.
func (h *jsiiProxy_HttpApi) MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricIntegrationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The time between when API Gateway receives a request from a client and when it returns a response to the client.
//
// The latency includes the integration latency and other API Gateway overhead.
// Experimental.
func (h *jsiiProxy_HttpApi) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of server-side errors captured in a given period.
// Experimental.
func (h *jsiiProxy_HttpApi) MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricServerError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (h *jsiiProxy_HttpApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes for importing an HttpApi into the CDK.
// Experimental.
type HttpApiAttributes struct {
	// The identifier of the HttpApi.
	// Experimental.
	HttpApiId *string `json:"httpApiId"`
	// The endpoint URL of the HttpApi.
	// Experimental.
	ApiEndpoint *string `json:"apiEndpoint"`
}

// Properties to initialize an instance of `HttpApi`.
// Experimental.
type HttpApiProps struct {
	// Name for the HTTP API resource.
	// Experimental.
	ApiName *string `json:"apiName"`
	// Specifies a CORS configuration for an API.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html
	//
	// Experimental.
	CorsPreflight *CorsPreflightOptions `json:"corsPreflight"`
	// Whether a default stage and deployment should be automatically created.
	// Experimental.
	CreateDefaultStage *bool `json:"createDefaultStage"`
	// Default OIDC scopes attached to all routes in the gateway, unless explicitly configured on the route.
	// Experimental.
	DefaultAuthorizationScopes *[]*string `json:"defaultAuthorizationScopes"`
	// Default Authorizer to applied to all routes in the gateway.
	// Experimental.
	DefaultAuthorizer IHttpRouteAuthorizer `json:"defaultAuthorizer"`
	// Configure a custom domain with the API mapping resource to the HTTP API.
	// Experimental.
	DefaultDomainMapping *DomainMappingOptions `json:"defaultDomainMapping"`
	// An integration that will be configured on the catch-all route ($default).
	// Experimental.
	DefaultIntegration IHttpRouteIntegration `json:"defaultIntegration"`
	// The description of the API.
	// Experimental.
	Description *string `json:"description"`
	// Specifies whether clients can invoke your API using the default endpoint.
	//
	// By default, clients can invoke your API with the default
	// `https://{api_id}.execute-api.{region}.amazonaws.com` endpoint. Enable
	// this if you would like clients to use your custom domain name.
	// Experimental.
	DisableExecuteApiEndpoint *bool `json:"disableExecuteApiEndpoint"`
}

// An authorizer for Http Apis.
// Experimental.
type HttpAuthorizer interface {
	awscdk.Resource
	IHttpAuthorizer
	AuthorizerId() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for HttpAuthorizer
type jsiiProxy_HttpAuthorizer struct {
	internal.Type__awscdkResource
	jsiiProxy_IHttpAuthorizer
}

func (j *jsiiProxy_HttpAuthorizer) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpAuthorizer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpAuthorizer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpAuthorizer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpAuthorizer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewHttpAuthorizer(scope constructs.Construct, id *string, props *HttpAuthorizerProps) HttpAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpAuthorizer{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpAuthorizer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpAuthorizer_Override(h HttpAuthorizer, scope constructs.Construct, id *string, props *HttpAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpAuthorizer",
		[]interface{}{scope, id, props},
		h,
	)
}

// Import an existing HTTP Authorizer into this CDK app.
// Experimental.
func HttpAuthorizer_FromHttpAuthorizerAttributes(scope constructs.Construct, id *string, attrs *HttpAuthorizerAttributes) IHttpRouteAuthorizer {
	_init_.Initialize()

	var returns IHttpRouteAuthorizer

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpAuthorizer",
		"fromHttpAuthorizerAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func HttpAuthorizer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpAuthorizer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func HttpAuthorizer_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpAuthorizer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (h *jsiiProxy_HttpAuthorizer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		h,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (h *jsiiProxy_HttpAuthorizer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (h *jsiiProxy_HttpAuthorizer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (h *jsiiProxy_HttpAuthorizer) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (h *jsiiProxy_HttpAuthorizer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Reference to an http authorizer.
// Experimental.
type HttpAuthorizerAttributes struct {
	// Id of the Authorizer.
	// Experimental.
	AuthorizerId *string `json:"authorizerId"`
	// Type of authorizer.
	//
	// Possible values are:
	// - JWT - JSON Web Token Authorizer
	// - CUSTOM - Lambda Authorizer
	// - NONE - No Authorization
	// Experimental.
	AuthorizerType *string `json:"authorizerType"`
}

// Properties to initialize an instance of `HttpAuthorizer`.
// Experimental.
type HttpAuthorizerProps struct {
	// HTTP Api to attach the authorizer to.
	// Experimental.
	HttpApi IHttpApi `json:"httpApi"`
	// The identity source for which authorization is requested.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-identitysource
	//
	// Experimental.
	IdentitySource *[]*string `json:"identitySource"`
	// The type of authorizer.
	// Experimental.
	Type HttpAuthorizerType `json:"type"`
	// Name of the authorizer.
	// Experimental.
	AuthorizerName *string `json:"authorizerName"`
	// The authorizer's Uniform Resource Identifier (URI).
	//
	// For REQUEST authorizers, this must be a well-formed Lambda function URI.
	// Experimental.
	AuthorizerUri *string `json:"authorizerUri"`
	// Specifies whether a Lambda authorizer returns a response in a simple format.
	//
	// If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
	// Experimental.
	EnableSimpleResponses *bool `json:"enableSimpleResponses"`
	// A list of the intended recipients of the JWT.
	//
	// A valid JWT must provide an aud that matches at least one entry in this list.
	// Experimental.
	JwtAudience *[]*string `json:"jwtAudience"`
	// The base domain of the identity provider that issues JWT.
	// Experimental.
	JwtIssuer *string `json:"jwtIssuer"`
	// Specifies the format of the payload sent to an HTTP API Lambda authorizer.
	// Experimental.
	PayloadFormatVersion AuthorizerPayloadVersion `json:"payloadFormatVersion"`
	// How long APIGateway should cache the results.
	//
	// Max 1 hour.
	// Experimental.
	ResultsCacheTtl awscdk.Duration `json:"resultsCacheTtl"`
}

// Supported Authorizer types.
// Experimental.
type HttpAuthorizerType string

const (
	HttpAuthorizerType_JWT HttpAuthorizerType = "JWT"
	HttpAuthorizerType_LAMBDA HttpAuthorizerType = "LAMBDA"
)

// Supported connection types.
// Experimental.
type HttpConnectionType string

const (
	HttpConnectionType_VPC_LINK HttpConnectionType = "VPC_LINK"
	HttpConnectionType_INTERNET HttpConnectionType = "INTERNET"
)

// The integration for an API route.
// Experimental.
type HttpIntegration interface {
	awscdk.Resource
	IHttpIntegration
	Env() *awscdk.ResourceEnvironment
	HttpApi() IHttpApi
	IntegrationId() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for HttpIntegration
type jsiiProxy_HttpIntegration struct {
	internal.Type__awscdkResource
	jsiiProxy_IHttpIntegration
}

func (j *jsiiProxy_HttpIntegration) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpIntegration) HttpApi() IHttpApi {
	var returns IHttpApi
	_jsii_.Get(
		j,
		"httpApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpIntegration) IntegrationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpIntegration) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpIntegration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewHttpIntegration(scope constructs.Construct, id *string, props *HttpIntegrationProps) HttpIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpIntegration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpIntegration_Override(h HttpIntegration, scope constructs.Construct, id *string, props *HttpIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpIntegration",
		[]interface{}{scope, id, props},
		h,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func HttpIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func HttpIntegration_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpIntegration",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (h *jsiiProxy_HttpIntegration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		h,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (h *jsiiProxy_HttpIntegration) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (h *jsiiProxy_HttpIntegration) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (h *jsiiProxy_HttpIntegration) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (h *jsiiProxy_HttpIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The integration properties.
// Experimental.
type HttpIntegrationProps struct {
	// The HTTP API to which this integration should be bound.
	// Experimental.
	HttpApi IHttpApi `json:"httpApi"`
	// Integration type.
	// Experimental.
	IntegrationType HttpIntegrationType `json:"integrationType"`
	// Integration URI.
	//
	// This will be the function ARN in the case of `HttpIntegrationType.LAMBDA_PROXY`,
	// or HTTP URL in the case of `HttpIntegrationType.HTTP_PROXY`.
	// Experimental.
	IntegrationUri *string `json:"integrationUri"`
	// The ID of the VPC link for a private integration.
	//
	// Supported only for HTTP APIs.
	// Experimental.
	ConnectionId *string `json:"connectionId"`
	// The type of the network connection to the integration endpoint.
	// Experimental.
	ConnectionType HttpConnectionType `json:"connectionType"`
	// The HTTP method to use when calling the underlying HTTP proxy.
	// Experimental.
	Method HttpMethod `json:"method"`
	// The version of the payload format.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	//
	// Experimental.
	PayloadFormatVersion PayloadFormatVersion `json:"payloadFormatVersion"`
	// Specifies the TLS configuration for a private integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `json:"secureServerName"`
}

// Supported integration types.
// Experimental.
type HttpIntegrationType string

const (
	HttpIntegrationType_LAMBDA_PROXY HttpIntegrationType = "LAMBDA_PROXY"
	HttpIntegrationType_HTTP_PROXY HttpIntegrationType = "HTTP_PROXY"
)

// Supported HTTP methods.
// Experimental.
type HttpMethod string

const (
	HttpMethod_ANY HttpMethod = "ANY"
	HttpMethod_DELETE HttpMethod = "DELETE"
	HttpMethod_GET HttpMethod = "GET"
	HttpMethod_HEAD HttpMethod = "HEAD"
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
	HttpMethod_PATCH HttpMethod = "PATCH"
	HttpMethod_POST HttpMethod = "POST"
	HttpMethod_PUT HttpMethod = "PUT"
)

// Explicitly configure no authorizers on specific HTTP API routes.
// Experimental.
type HttpNoneAuthorizer interface {
	IHttpRouteAuthorizer
	Bind(_arg *HttpRouteAuthorizerBindOptions) *HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpNoneAuthorizer
type jsiiProxy_HttpNoneAuthorizer struct {
	jsiiProxy_IHttpRouteAuthorizer
}

// Experimental.
func NewHttpNoneAuthorizer() HttpNoneAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpNoneAuthorizer{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpNoneAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewHttpNoneAuthorizer_Override(h HttpNoneAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpNoneAuthorizer",
		nil, // no parameters
		h,
	)
}

// Bind this authorizer to a specified Http route.
// Experimental.
func (h *jsiiProxy_HttpNoneAuthorizer) Bind(_arg *HttpRouteAuthorizerBindOptions) *HttpRouteAuthorizerConfig {
	var returns *HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

// Route class that creates the Route for API Gateway HTTP API.
// Experimental.
type HttpRoute interface {
	awscdk.Resource
	IHttpRoute
	Env() *awscdk.ResourceEnvironment
	HttpApi() IHttpApi
	Node() constructs.Node
	Path() *string
	PhysicalName() *string
	RouteId() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for HttpRoute
type jsiiProxy_HttpRoute struct {
	internal.Type__awscdkResource
	jsiiProxy_IHttpRoute
}

func (j *jsiiProxy_HttpRoute) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpRoute) HttpApi() IHttpApi {
	var returns IHttpApi
	_jsii_.Get(
		j,
		"httpApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpRoute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpRoute) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpRoute) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpRoute) RouteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpRoute) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewHttpRoute(scope constructs.Construct, id *string, props *HttpRouteProps) HttpRoute {
	_init_.Initialize()

	j := jsiiProxy_HttpRoute{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRoute",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpRoute_Override(h HttpRoute, scope constructs.Construct, id *string, props *HttpRouteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRoute",
		[]interface{}{scope, id, props},
		h,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func HttpRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRoute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func HttpRoute_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRoute",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (h *jsiiProxy_HttpRoute) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		h,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (h *jsiiProxy_HttpRoute) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (h *jsiiProxy_HttpRoute) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (h *jsiiProxy_HttpRoute) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (h *jsiiProxy_HttpRoute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Input to the bind() operation, that binds an authorizer to a route.
// Experimental.
type HttpRouteAuthorizerBindOptions struct {
	// The route to which the authorizer is being bound.
	// Experimental.
	Route IHttpRoute `json:"route"`
	// The scope for any constructs created as part of the bind.
	// Experimental.
	Scope constructs.Construct `json:"scope"`
}

// Results of binding an authorizer to an http route.
// Experimental.
type HttpRouteAuthorizerConfig struct {
	// The type of authorization.
	//
	// Possible values are:
	// - JWT - JSON Web Token Authorizer
	// - CUSTOM - Lambda Authorizer
	// - NONE - No Authorization
	// Experimental.
	AuthorizationType *string `json:"authorizationType"`
	// The list of OIDC scopes to include in the authorization.
	// Experimental.
	AuthorizationScopes *[]*string `json:"authorizationScopes"`
	// The authorizer id.
	// Experimental.
	AuthorizerId *string `json:"authorizerId"`
}

// Options to the HttpRouteIntegration during its bind operation.
// Experimental.
type HttpRouteIntegrationBindOptions struct {
	// The route to which this is being bound.
	// Experimental.
	Route IHttpRoute `json:"route"`
	// The current scope in which the bind is occurring.
	//
	// If the `HttpRouteIntegration` being bound creates additional constructs,
	// this will be used as their parent scope.
	// Experimental.
	Scope constructs.Construct `json:"scope"`
}

// Config returned back as a result of the bind.
// Experimental.
type HttpRouteIntegrationConfig struct {
	// Payload format version in the case of lambda proxy integration.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	//
	// Experimental.
	PayloadFormatVersion PayloadFormatVersion `json:"payloadFormatVersion"`
	// Integration type.
	// Experimental.
	Type HttpIntegrationType `json:"type"`
	// Integration URI.
	// Experimental.
	Uri *string `json:"uri"`
	// The ID of the VPC link for a private integration.
	//
	// Supported only for HTTP APIs.
	// Experimental.
	ConnectionId *string `json:"connectionId"`
	// The type of the network connection to the integration endpoint.
	// Experimental.
	ConnectionType HttpConnectionType `json:"connectionType"`
	// The HTTP method that must be used to invoke the underlying proxy.
	//
	// Required for `HttpIntegrationType.HTTP_PROXY`
	// Experimental.
	Method HttpMethod `json:"method"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `json:"secureServerName"`
}

// HTTP route in APIGateway is a combination of the HTTP method and the path component.
//
// This class models that combination.
// Experimental.
type HttpRouteKey interface {
	Key() *string
	Path() *string
}

// The jsii proxy struct for HttpRouteKey
type jsiiProxy_HttpRouteKey struct {
	_ byte // padding
}

func (j *jsiiProxy_HttpRouteKey) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpRouteKey) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}


// Create a route key with the combination of the path and the method.
// Experimental.
func HttpRouteKey_With(path *string, method HttpMethod) HttpRouteKey {
	_init_.Initialize()

	var returns HttpRouteKey

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRouteKey",
		"with",
		[]interface{}{path, method},
		&returns,
	)

	return returns
}

func HttpRouteKey_DEFAULT() HttpRouteKey {
	_init_.Initialize()
	var returns HttpRouteKey
	_jsii_.StaticGet(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRouteKey",
		"DEFAULT",
		&returns,
	)
	return returns
}

// Properties to initialize a new Route.
// Experimental.
type HttpRouteProps struct {
	// The integration to be configured on this route.
	// Experimental.
	Integration IHttpRouteIntegration `json:"integration"`
	// the API the route is associated with.
	// Experimental.
	HttpApi IHttpApi `json:"httpApi"`
	// The key to this route.
	//
	// This is a combination of an HTTP method and an HTTP path.
	// Experimental.
	RouteKey HttpRouteKey `json:"routeKey"`
	// The list of OIDC scopes to include in the authorization.
	//
	// These scopes will be merged with the scopes from the attached authorizer
	// Experimental.
	AuthorizationScopes *[]*string `json:"authorizationScopes"`
	// Authorizer for a WebSocket API or an HTTP API.
	// Experimental.
	Authorizer IHttpRouteAuthorizer `json:"authorizer"`
}

// Represents a stage where an instance of the API is deployed.
// Experimental.
type HttpStage interface {
	awscdk.Resource
	IHttpStage
	IStage
	Api() IHttpApi
	BaseApi() IApi
	DomainUrl() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	StageName() *string
	Url() *string
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDataProcessed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for HttpStage
type jsiiProxy_HttpStage struct {
	internal.Type__awscdkResource
	jsiiProxy_IHttpStage
	jsiiProxy_IStage
}

func (j *jsiiProxy_HttpStage) Api() IHttpApi {
	var returns IHttpApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpStage) BaseApi() IApi {
	var returns IApi
	_jsii_.Get(
		j,
		"baseApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpStage) DomainUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpStage) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpStage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpStage) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpStage) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpStage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpStage) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Experimental.
func NewHttpStage(scope constructs.Construct, id *string, props *HttpStageProps) HttpStage {
	_init_.Initialize()

	j := jsiiProxy_HttpStage{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpStage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpStage_Override(h HttpStage, scope constructs.Construct, id *string, props *HttpStageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpStage",
		[]interface{}{scope, id, props},
		h,
	)
}

// Import an existing stage into this CDK app.
// Experimental.
func HttpStage_FromHttpStageAttributes(scope constructs.Construct, id *string, attrs *HttpStageAttributes) IHttpStage {
	_init_.Initialize()

	var returns IHttpStage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpStage",
		"fromHttpStageAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func HttpStage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpStage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func HttpStage_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpStage",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (h *jsiiProxy_HttpStage) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		h,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (h *jsiiProxy_HttpStage) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (h *jsiiProxy_HttpStage) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (h *jsiiProxy_HttpStage) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Return the given named metric for this HTTP Api Gateway Stage.
// Experimental.
func (h *jsiiProxy_HttpStage) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of client-side errors captured in a given period.
// Experimental.
func (h *jsiiProxy_HttpStage) MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricClientError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the total number API requests in a given period.
// Experimental.
func (h *jsiiProxy_HttpStage) MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the amount of data processed in bytes.
// Experimental.
func (h *jsiiProxy_HttpStage) MetricDataProcessed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricDataProcessed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
// Experimental.
func (h *jsiiProxy_HttpStage) MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricIntegrationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The time between when API Gateway receives a request from a client and when it returns a response to the client.
//
// The latency includes the integration latency and other API Gateway overhead.
// Experimental.
func (h *jsiiProxy_HttpStage) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of server-side errors captured in a given period.
// Experimental.
func (h *jsiiProxy_HttpStage) MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricServerError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (h *jsiiProxy_HttpStage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The attributes used to import existing HttpStage.
// Experimental.
type HttpStageAttributes struct {
	// The name of the stage.
	// Experimental.
	StageName *string `json:"stageName"`
	// The API to which this stage is associated.
	// Experimental.
	Api IHttpApi `json:"api"`
}

// The options to create a new Stage for an HTTP API.
// Experimental.
type HttpStageOptions struct {
	// Whether updates to an API automatically trigger a new deployment.
	// Experimental.
	AutoDeploy *bool `json:"autoDeploy"`
	// The options for custom domain and api mapping.
	// Experimental.
	DomainMapping *DomainMappingOptions `json:"domainMapping"`
	// The name of the stage.
	//
	// See `StageName` class for more details.
	// Experimental.
	StageName *string `json:"stageName"`
}

// Properties to initialize an instance of `HttpStage`.
// Experimental.
type HttpStageProps struct {
	// Whether updates to an API automatically trigger a new deployment.
	// Experimental.
	AutoDeploy *bool `json:"autoDeploy"`
	// The options for custom domain and api mapping.
	// Experimental.
	DomainMapping *DomainMappingOptions `json:"domainMapping"`
	// The name of the stage.
	//
	// See `StageName` class for more details.
	// Experimental.
	StageName *string `json:"stageName"`
	// The HTTP API to which this stage is associated.
	// Experimental.
	HttpApi IHttpApi `json:"httpApi"`
}

// Represents a API Gateway HTTP/WebSocket API.
// Experimental.
type IApi interface {
	awscdk.IResource
	// Return the given named metric for this Api Gateway.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The default endpoint for an API.
	// Experimental.
	ApiEndpoint() *string
	// The identifier of this API Gateway API.
	// Experimental.
	ApiId() *string
}

// The jsii proxy for IApi
type jsiiProxy_IApi struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IApi) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IApi) ApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApi) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

// Represents an ApiGatewayV2 ApiMapping resource.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html
//
// Experimental.
type IApiMapping interface {
	awscdk.IResource
	// ID of the api mapping.
	// Experimental.
	ApiMappingId() *string
}

// The jsii proxy for IApiMapping
type jsiiProxy_IApiMapping struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IApiMapping) ApiMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiMappingId",
		&returns,
	)
	return returns
}

// Represents an Authorizer.
// Experimental.
type IAuthorizer interface {
	awscdk.IResource
	// Id of the Authorizer.
	// Experimental.
	AuthorizerId() *string
}

// The jsii proxy for IAuthorizer
type jsiiProxy_IAuthorizer struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAuthorizer) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

// Represents an APIGatewayV2 DomainName.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html
//
// Experimental.
type IDomainName interface {
	awscdk.IResource
	// The custom domain name.
	// Experimental.
	Name() *string
	// The domain name associated with the regional endpoint for this custom domain name.
	// Experimental.
	RegionalDomainName() *string
	// The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint.
	// Experimental.
	RegionalHostedZoneId() *string
}

// The jsii proxy for IDomainName
type jsiiProxy_IDomainName struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDomainName) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainName) RegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainName) RegionalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalHostedZoneId",
		&returns,
	)
	return returns
}

// Represents an HTTP API.
// Experimental.
type IHttpApi interface {
	IApi
	// Add a new VpcLink.
	// Experimental.
	AddVpcLink(options *VpcLinkProps) VpcLink
	// Metric for the number of client-side errors captured in a given period.
	// Experimental.
	MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number API requests in a given period.
	// Experimental.
	MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the amount of data processed in bytes.
	// Experimental.
	MetricDataProcessed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
	// Experimental.
	MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time between when API Gateway receives a request from a client and when it returns a response to the client.
	//
	// The latency includes the integration latency and other API Gateway overhead.
	// Experimental.
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of server-side errors captured in a given period.
	// Experimental.
	MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The identifier of this API Gateway HTTP API.
	// Deprecated: - use apiId instead
	HttpApiId() *string
}

// The jsii proxy for IHttpApi
type jsiiProxy_IHttpApi struct {
	jsiiProxy_IApi
}

func (i *jsiiProxy_IHttpApi) AddVpcLink(options *VpcLinkProps) VpcLink {
	var returns VpcLink

	_jsii_.Invoke(
		i,
		"addVpcLink",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpApi) MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricClientError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpApi) MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpApi) MetricDataProcessed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDataProcessed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpApi) MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIntegrationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpApi) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpApi) MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricServerError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IHttpApi) HttpApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpApiId",
		&returns,
	)
	return returns
}

// An authorizer for HTTP APIs.
// Experimental.
type IHttpAuthorizer interface {
	IAuthorizer
}

// The jsii proxy for IHttpAuthorizer
type jsiiProxy_IHttpAuthorizer struct {
	jsiiProxy_IAuthorizer
}

// Represents an Integration for an HTTP API.
// Experimental.
type IHttpIntegration interface {
	IIntegration
	// The HTTP API associated with this integration.
	// Experimental.
	HttpApi() IHttpApi
}

// The jsii proxy for IHttpIntegration
type jsiiProxy_IHttpIntegration struct {
	jsiiProxy_IIntegration
}

func (j *jsiiProxy_IHttpIntegration) HttpApi() IHttpApi {
	var returns IHttpApi
	_jsii_.Get(
		j,
		"httpApi",
		&returns,
	)
	return returns
}

// Represents a Route for an HTTP API.
// Experimental.
type IHttpRoute interface {
	IRoute
	// The HTTP API associated with this route.
	// Experimental.
	HttpApi() IHttpApi
	// Returns the path component of this HTTP route, `undefined` if the path is the catch-all route.
	// Experimental.
	Path() *string
}

// The jsii proxy for IHttpRoute
type jsiiProxy_IHttpRoute struct {
	jsiiProxy_IRoute
}

func (j *jsiiProxy_IHttpRoute) HttpApi() IHttpApi {
	var returns IHttpApi
	_jsii_.Get(
		j,
		"httpApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHttpRoute) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

// An authorizer that can attach to an Http Route.
// Experimental.
type IHttpRouteAuthorizer interface {
	// Bind this authorizer to a specified Http route.
	// Experimental.
	Bind(options *HttpRouteAuthorizerBindOptions) *HttpRouteAuthorizerConfig
}

// The jsii proxy for IHttpRouteAuthorizer
type jsiiProxy_IHttpRouteAuthorizer struct {
	_ byte // padding
}

func (i *jsiiProxy_IHttpRouteAuthorizer) Bind(options *HttpRouteAuthorizerBindOptions) *HttpRouteAuthorizerConfig {
	var returns *HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// The interface that various route integration classes will inherit.
// Experimental.
type IHttpRouteIntegration interface {
	// Bind this integration to the route.
	// Experimental.
	Bind(options *HttpRouteIntegrationBindOptions) *HttpRouteIntegrationConfig
}

// The jsii proxy for IHttpRouteIntegration
type jsiiProxy_IHttpRouteIntegration struct {
	_ byte // padding
}

func (i *jsiiProxy_IHttpRouteIntegration) Bind(options *HttpRouteIntegrationBindOptions) *HttpRouteIntegrationConfig {
	var returns *HttpRouteIntegrationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Represents the HttpStage.
// Experimental.
type IHttpStage interface {
	IStage
	// Metric for the number of client-side errors captured in a given period.
	// Experimental.
	MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number API requests in a given period.
	// Experimental.
	MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the amount of data processed in bytes.
	// Experimental.
	MetricDataProcessed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
	// Experimental.
	MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time between when API Gateway receives a request from a client and when it returns a response to the client.
	//
	// The latency includes the integration latency and other API Gateway overhead.
	// Experimental.
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of server-side errors captured in a given period.
	// Experimental.
	MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The API this stage is associated to.
	// Experimental.
	Api() IHttpApi
}

// The jsii proxy for IHttpStage
type jsiiProxy_IHttpStage struct {
	jsiiProxy_IStage
}

func (i *jsiiProxy_IHttpStage) MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricClientError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpStage) MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpStage) MetricDataProcessed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDataProcessed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpStage) MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIntegrationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpStage) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpStage) MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricServerError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IHttpStage) Api() IHttpApi {
	var returns IHttpApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

// Represents an integration to an API Route.
// Experimental.
type IIntegration interface {
	awscdk.IResource
	// Id of the integration.
	// Experimental.
	IntegrationId() *string
}

// The jsii proxy for IIntegration
type jsiiProxy_IIntegration struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IIntegration) IntegrationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationId",
		&returns,
	)
	return returns
}

// Represents a route.
// Experimental.
type IRoute interface {
	awscdk.IResource
	// Id of the Route.
	// Experimental.
	RouteId() *string
}

// The jsii proxy for IRoute
type jsiiProxy_IRoute struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRoute) RouteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeId",
		&returns,
	)
	return returns
}

// Represents a Stage.
// Experimental.
type IStage interface {
	awscdk.IResource
	// Return the given named metric for this HTTP Api Gateway Stage.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The name of the stage;
	//
	// its primary identifier.
	// Experimental.
	StageName() *string
	// The URL to this stage.
	// Experimental.
	Url() *string
}

// The jsii proxy for IStage
type jsiiProxy_IStage struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IStage) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IStage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStage) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

// Represents an API Gateway VpcLink.
// Experimental.
type IVpcLink interface {
	awscdk.IResource
	// The VPC to which this VPC Link is associated with.
	// Experimental.
	Vpc() awsec2.IVpc
	// Physical ID of the VpcLink resource.
	// Experimental.
	VpcLinkId() *string
}

// The jsii proxy for IVpcLink
type jsiiProxy_IVpcLink struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IVpcLink) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcLink) VpcLinkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcLinkId",
		&returns,
	)
	return returns
}

// Represents a WebSocket API.
// Experimental.
type IWebSocketApi interface {
	IApi
}

// The jsii proxy for IWebSocketApi
type jsiiProxy_IWebSocketApi struct {
	jsiiProxy_IApi
}

// Represents an Integration for an WebSocket API.
// Experimental.
type IWebSocketIntegration interface {
	IIntegration
	// The WebSocket API associated with this integration.
	// Experimental.
	WebSocketApi() IWebSocketApi
}

// The jsii proxy for IWebSocketIntegration
type jsiiProxy_IWebSocketIntegration struct {
	jsiiProxy_IIntegration
}

func (j *jsiiProxy_IWebSocketIntegration) WebSocketApi() IWebSocketApi {
	var returns IWebSocketApi
	_jsii_.Get(
		j,
		"webSocketApi",
		&returns,
	)
	return returns
}

// Represents a Route for an WebSocket API.
// Experimental.
type IWebSocketRoute interface {
	IRoute
	// The key to this route.
	// Experimental.
	RouteKey() *string
	// The WebSocket API associated with this route.
	// Experimental.
	WebSocketApi() IWebSocketApi
}

// The jsii proxy for IWebSocketRoute
type jsiiProxy_IWebSocketRoute struct {
	jsiiProxy_IRoute
}

func (j *jsiiProxy_IWebSocketRoute) RouteKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebSocketRoute) WebSocketApi() IWebSocketApi {
	var returns IWebSocketApi
	_jsii_.Get(
		j,
		"webSocketApi",
		&returns,
	)
	return returns
}

// The interface that various route integration classes will inherit.
// Experimental.
type IWebSocketRouteIntegration interface {
	// Bind this integration to the route.
	// Experimental.
	Bind(options *WebSocketRouteIntegrationBindOptions) *WebSocketRouteIntegrationConfig
}

// The jsii proxy for IWebSocketRouteIntegration
type jsiiProxy_IWebSocketRouteIntegration struct {
	_ byte // padding
}

func (i *jsiiProxy_IWebSocketRouteIntegration) Bind(options *WebSocketRouteIntegrationBindOptions) *WebSocketRouteIntegrationConfig {
	var returns *WebSocketRouteIntegrationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Represents the WebSocketStage.
// Experimental.
type IWebSocketStage interface {
	IStage
	// The API this stage is associated to.
	// Experimental.
	Api() IWebSocketApi
	// The callback URL to this stage.
	//
	// You can use the callback URL to send messages to the client from the backend system.
	// https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-basic-concept.html
	// https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html
	// Experimental.
	CallbackUrl() *string
}

// The jsii proxy for IWebSocketStage
type jsiiProxy_IWebSocketStage struct {
	jsiiProxy_IStage
}

func (j *jsiiProxy_IWebSocketStage) Api() IWebSocketApi {
	var returns IWebSocketApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebSocketStage) CallbackUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callbackUrl",
		&returns,
	)
	return returns
}

// Payload format version for lambda proxy integration.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
//
// Experimental.
type PayloadFormatVersion interface {
	Version() *string
}

// The jsii proxy struct for PayloadFormatVersion
type jsiiProxy_PayloadFormatVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_PayloadFormatVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// A custom payload version.
//
// Typically used if there is a version number that the CDK doesn't support yet
// Experimental.
func PayloadFormatVersion_Custom(version *string) PayloadFormatVersion {
	_init_.Initialize()

	var returns PayloadFormatVersion

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.PayloadFormatVersion",
		"custom",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func PayloadFormatVersion_VERSION_1_0() PayloadFormatVersion {
	_init_.Initialize()
	var returns PayloadFormatVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-apigatewayv2-alpha.PayloadFormatVersion",
		"VERSION_1_0",
		&returns,
	)
	return returns
}

func PayloadFormatVersion_VERSION_2_0() PayloadFormatVersion {
	_init_.Initialize()
	var returns PayloadFormatVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-apigatewayv2-alpha.PayloadFormatVersion",
		"VERSION_2_0",
		&returns,
	)
	return returns
}

// The attributes used to import existing Stage.
// Experimental.
type StageAttributes struct {
	// The name of the stage.
	// Experimental.
	StageName *string `json:"stageName"`
}

// Options required to create a new stage.
//
// Options that are common between HTTP and Websocket APIs.
// Experimental.
type StageOptions struct {
	// Whether updates to an API automatically trigger a new deployment.
	// Experimental.
	AutoDeploy *bool `json:"autoDeploy"`
	// The options for custom domain and api mapping.
	// Experimental.
	DomainMapping *DomainMappingOptions `json:"domainMapping"`
}

// Define a new VPC Link Specifies an API Gateway VPC link for a HTTP API to access resources in an Amazon Virtual Private Cloud (VPC).
// Experimental.
type VpcLink interface {
	awscdk.Resource
	IVpcLink
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	Vpc() awsec2.IVpc
	VpcLinkId() *string
	AddSecurityGroups(groups ...awsec2.ISecurityGroup)
	AddSubnets(subnets ...awsec2.ISubnet)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for VpcLink
type jsiiProxy_VpcLink struct {
	internal.Type__awscdkResource
	jsiiProxy_IVpcLink
}

func (j *jsiiProxy_VpcLink) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcLink) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcLink) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcLink) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcLink) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcLink) VpcLinkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcLinkId",
		&returns,
	)
	return returns
}


// Experimental.
func NewVpcLink(scope constructs.Construct, id *string, props *VpcLinkProps) VpcLink {
	_init_.Initialize()

	j := jsiiProxy_VpcLink{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.VpcLink",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewVpcLink_Override(v VpcLink, scope constructs.Construct, id *string, props *VpcLinkProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.VpcLink",
		[]interface{}{scope, id, props},
		v,
	)
}

// Import a VPC Link by specifying its attributes.
// Experimental.
func VpcLink_FromVpcLinkAttributes(scope constructs.Construct, id *string, attrs *VpcLinkAttributes) IVpcLink {
	_init_.Initialize()

	var returns IVpcLink

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.VpcLink",
		"fromVpcLinkAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func VpcLink_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.VpcLink",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func VpcLink_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.VpcLink",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds the provided security groups to the vpc link.
// Experimental.
func (v *jsiiProxy_VpcLink) AddSecurityGroups(groups ...awsec2.ISecurityGroup) {
	args := []interface{}{}
	for _, a := range groups {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		v,
		"addSecurityGroups",
		args,
	)
}

// Adds the provided subnets to the vpc link.
// Experimental.
func (v *jsiiProxy_VpcLink) AddSubnets(subnets ...awsec2.ISubnet) {
	args := []interface{}{}
	for _, a := range subnets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		v,
		"addSubnets",
		args,
	)
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (v *jsiiProxy_VpcLink) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (v *jsiiProxy_VpcLink) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (v *jsiiProxy_VpcLink) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (v *jsiiProxy_VpcLink) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (v *jsiiProxy_VpcLink) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes when importing a new VpcLink.
// Experimental.
type VpcLinkAttributes struct {
	// The VPC to which this VPC link is associated with.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The VPC Link id.
	// Experimental.
	VpcLinkId *string `json:"vpcLinkId"`
}

// Properties for a VpcLink.
// Experimental.
type VpcLinkProps struct {
	// The VPC in which the private resources reside.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// A list of security groups for the VPC link.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// A list of subnets for the VPC link.
	// Experimental.
	Subnets *awsec2.SubnetSelection `json:"subnets"`
	// The name used to label and identify the VPC link.
	// Experimental.
	VpcLinkName *string `json:"vpcLinkName"`
}

// Create a new API Gateway WebSocket API endpoint.
// Experimental.
type WebSocketApi interface {
	awscdk.Resource
	IApi
	IWebSocketApi
	ApiEndpoint() *string
	ApiId() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	WebSocketApiName() *string
	AddRoute(routeKey *string, options *WebSocketRouteOptions) WebSocketRoute
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for WebSocketApi
type jsiiProxy_WebSocketApi struct {
	internal.Type__awscdkResource
	jsiiProxy_IApi
	jsiiProxy_IWebSocketApi
}

func (j *jsiiProxy_WebSocketApi) ApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketApi) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketApi) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketApi) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketApi) WebSocketApiName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webSocketApiName",
		&returns,
	)
	return returns
}


// Experimental.
func NewWebSocketApi(scope constructs.Construct, id *string, props *WebSocketApiProps) WebSocketApi {
	_init_.Initialize()

	j := jsiiProxy_WebSocketApi{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketApi_Override(w WebSocketApi, scope constructs.Construct, id *string, props *WebSocketApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketApi",
		[]interface{}{scope, id, props},
		w,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func WebSocketApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func WebSocketApi_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketApi",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a new route.
// Experimental.
func (w *jsiiProxy_WebSocketApi) AddRoute(routeKey *string, options *WebSocketRouteOptions) WebSocketRoute {
	var returns WebSocketRoute

	_jsii_.Invoke(
		w,
		"addRoute",
		[]interface{}{routeKey, options},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (w *jsiiProxy_WebSocketApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		w,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (w *jsiiProxy_WebSocketApi) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (w *jsiiProxy_WebSocketApi) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (w *jsiiProxy_WebSocketApi) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Return the given named metric for this Api Gateway.
// Experimental.
func (w *jsiiProxy_WebSocketApi) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		w,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (w *jsiiProxy_WebSocketApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Props for WebSocket API.
// Experimental.
type WebSocketApiProps struct {
	// Name for the WebSocket API resource.
	// Experimental.
	ApiName *string `json:"apiName"`
	// Options to configure a '$connect' route.
	// Experimental.
	ConnectRouteOptions *WebSocketRouteOptions `json:"connectRouteOptions"`
	// Options to configure a '$default' route.
	// Experimental.
	DefaultRouteOptions *WebSocketRouteOptions `json:"defaultRouteOptions"`
	// The description of the API.
	// Experimental.
	Description *string `json:"description"`
	// Options to configure a '$disconnect' route.
	// Experimental.
	DisconnectRouteOptions *WebSocketRouteOptions `json:"disconnectRouteOptions"`
	// The route selection expression for the API.
	// Experimental.
	RouteSelectionExpression *string `json:"routeSelectionExpression"`
}

// The integration for an API route.
// Experimental.
type WebSocketIntegration interface {
	awscdk.Resource
	IWebSocketIntegration
	Env() *awscdk.ResourceEnvironment
	IntegrationId() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	WebSocketApi() IWebSocketApi
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for WebSocketIntegration
type jsiiProxy_WebSocketIntegration struct {
	internal.Type__awscdkResource
	jsiiProxy_IWebSocketIntegration
}

func (j *jsiiProxy_WebSocketIntegration) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketIntegration) IntegrationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketIntegration) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketIntegration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketIntegration) WebSocketApi() IWebSocketApi {
	var returns IWebSocketApi
	_jsii_.Get(
		j,
		"webSocketApi",
		&returns,
	)
	return returns
}


// Experimental.
func NewWebSocketIntegration(scope constructs.Construct, id *string, props *WebSocketIntegrationProps) WebSocketIntegration {
	_init_.Initialize()

	j := jsiiProxy_WebSocketIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketIntegration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketIntegration_Override(w WebSocketIntegration, scope constructs.Construct, id *string, props *WebSocketIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketIntegration",
		[]interface{}{scope, id, props},
		w,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func WebSocketIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func WebSocketIntegration_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketIntegration",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (w *jsiiProxy_WebSocketIntegration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		w,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (w *jsiiProxy_WebSocketIntegration) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (w *jsiiProxy_WebSocketIntegration) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (w *jsiiProxy_WebSocketIntegration) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (w *jsiiProxy_WebSocketIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The integration properties.
// Experimental.
type WebSocketIntegrationProps struct {
	// Integration type.
	// Experimental.
	IntegrationType WebSocketIntegrationType `json:"integrationType"`
	// Integration URI.
	// Experimental.
	IntegrationUri *string `json:"integrationUri"`
	// The WebSocket API to which this integration should be bound.
	// Experimental.
	WebSocketApi IWebSocketApi `json:"webSocketApi"`
}

// WebSocket Integration Types.
// Experimental.
type WebSocketIntegrationType string

const (
	WebSocketIntegrationType_AWS_PROXY WebSocketIntegrationType = "AWS_PROXY"
)

// Route class that creates the Route for API Gateway WebSocket API.
// Experimental.
type WebSocketRoute interface {
	awscdk.Resource
	IWebSocketRoute
	Env() *awscdk.ResourceEnvironment
	IntegrationResponseId() *string
	Node() constructs.Node
	PhysicalName() *string
	RouteId() *string
	RouteKey() *string
	Stack() awscdk.Stack
	WebSocketApi() IWebSocketApi
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for WebSocketRoute
type jsiiProxy_WebSocketRoute struct {
	internal.Type__awscdkResource
	jsiiProxy_IWebSocketRoute
}

func (j *jsiiProxy_WebSocketRoute) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketRoute) IntegrationResponseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationResponseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketRoute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketRoute) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketRoute) RouteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketRoute) RouteKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketRoute) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketRoute) WebSocketApi() IWebSocketApi {
	var returns IWebSocketApi
	_jsii_.Get(
		j,
		"webSocketApi",
		&returns,
	)
	return returns
}


// Experimental.
func NewWebSocketRoute(scope constructs.Construct, id *string, props *WebSocketRouteProps) WebSocketRoute {
	_init_.Initialize()

	j := jsiiProxy_WebSocketRoute{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRoute",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketRoute_Override(w WebSocketRoute, scope constructs.Construct, id *string, props *WebSocketRouteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRoute",
		[]interface{}{scope, id, props},
		w,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func WebSocketRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRoute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func WebSocketRoute_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRoute",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (w *jsiiProxy_WebSocketRoute) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		w,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (w *jsiiProxy_WebSocketRoute) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (w *jsiiProxy_WebSocketRoute) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (w *jsiiProxy_WebSocketRoute) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (w *jsiiProxy_WebSocketRoute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options to the WebSocketRouteIntegration during its bind operation.
// Experimental.
type WebSocketRouteIntegrationBindOptions struct {
	// The route to which this is being bound.
	// Experimental.
	Route IWebSocketRoute `json:"route"`
	// The current scope in which the bind is occurring.
	//
	// If the `WebSocketRouteIntegration` being bound creates additional constructs,
	// this will be used as their parent scope.
	// Experimental.
	Scope constructs.Construct `json:"scope"`
}

// Config returned back as a result of the bind.
// Experimental.
type WebSocketRouteIntegrationConfig struct {
	// Integration type.
	// Experimental.
	Type WebSocketIntegrationType `json:"type"`
	// Integration URI.
	// Experimental.
	Uri *string `json:"uri"`
}

// Options used to add route to the API.
// Experimental.
type WebSocketRouteOptions struct {
	// The integration to be configured on this route.
	// Experimental.
	Integration IWebSocketRouteIntegration `json:"integration"`
}

// Properties to initialize a new Route.
// Experimental.
type WebSocketRouteProps struct {
	// The integration to be configured on this route.
	// Experimental.
	Integration IWebSocketRouteIntegration `json:"integration"`
	// The key to this route.
	// Experimental.
	RouteKey *string `json:"routeKey"`
	// the API the route is associated with.
	// Experimental.
	WebSocketApi IWebSocketApi `json:"webSocketApi"`
}

// Represents a stage where an instance of the API is deployed.
// Experimental.
type WebSocketStage interface {
	awscdk.Resource
	IStage
	IWebSocketStage
	Api() IWebSocketApi
	BaseApi() IApi
	CallbackUrl() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	StageName() *string
	Url() *string
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for WebSocketStage
type jsiiProxy_WebSocketStage struct {
	internal.Type__awscdkResource
	jsiiProxy_IStage
	jsiiProxy_IWebSocketStage
}

func (j *jsiiProxy_WebSocketStage) Api() IWebSocketApi {
	var returns IWebSocketApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) BaseApi() IApi {
	var returns IApi
	_jsii_.Get(
		j,
		"baseApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) CallbackUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callbackUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Experimental.
func NewWebSocketStage(scope constructs.Construct, id *string, props *WebSocketStageProps) WebSocketStage {
	_init_.Initialize()

	j := jsiiProxy_WebSocketStage{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketStage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketStage_Override(w WebSocketStage, scope constructs.Construct, id *string, props *WebSocketStageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketStage",
		[]interface{}{scope, id, props},
		w,
	)
}

// Import an existing stage into this CDK app.
// Experimental.
func WebSocketStage_FromWebSocketStageAttributes(scope constructs.Construct, id *string, attrs *WebSocketStageAttributes) IWebSocketStage {
	_init_.Initialize()

	var returns IWebSocketStage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketStage",
		"fromWebSocketStageAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func WebSocketStage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketStage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func WebSocketStage_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketStage",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (w *jsiiProxy_WebSocketStage) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		w,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (w *jsiiProxy_WebSocketStage) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (w *jsiiProxy_WebSocketStage) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (w *jsiiProxy_WebSocketStage) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Return the given named metric for this HTTP Api Gateway Stage.
// Experimental.
func (w *jsiiProxy_WebSocketStage) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		w,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (w *jsiiProxy_WebSocketStage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The attributes used to import existing WebSocketStage.
// Experimental.
type WebSocketStageAttributes struct {
	// The name of the stage.
	// Experimental.
	StageName *string `json:"stageName"`
	// The API to which this stage is associated.
	// Experimental.
	Api IWebSocketApi `json:"api"`
}

// Properties to initialize an instance of `WebSocketStage`.
// Experimental.
type WebSocketStageProps struct {
	// Whether updates to an API automatically trigger a new deployment.
	// Experimental.
	AutoDeploy *bool `json:"autoDeploy"`
	// The options for custom domain and api mapping.
	// Experimental.
	DomainMapping *DomainMappingOptions `json:"domainMapping"`
	// The name of the stage.
	// Experimental.
	StageName *string `json:"stageName"`
	// The WebSocket API to which this stage is associated.
	// Experimental.
	WebSocketApi IWebSocketApi `json:"webSocketApi"`
}

