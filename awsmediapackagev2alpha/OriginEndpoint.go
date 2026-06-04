package awsmediapackagev2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediapackagev2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediapackagev2"
	"github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines an AWS Elemental MediaPackage V2 Origin Endpoint.
//
// Example:
//   var channel Channel
//   var spekeRole IRole
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("TsEndpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Ts(&TsSegmentProps{
//   		Encryption: awsmediapackagev2alpha.TsEncryption_Speke(&TsSpekeEncryptionProps{
//   			Method: awsmediapackagev2alpha.TsEncryptionMethod_SAMPLE_AES,
//   			ResourceId: jsii.String("my-content-id"),
//   			Url: jsii.String("https://example.com/speke"),
//   			Role: spekeRole,
//   		}),
//   	}),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   		}),
//   	},
//   })
//
// Experimental.
type OriginEndpoint interface {
	awscdk.Resource
	IOriginEndpoint
	// Indicates if an origin endpoint resource policy should automatically be created upon the first call to `addToResourcePolicy`.
	// Experimental.
	AutoCreatePolicy() *bool
	// Experimental.
	SetAutoCreatePolicy(val *bool)
	// CDN authorization configuration to be applied when the policy is created.
	//
	// Set by subclass constructors when `cdnAuth` is provided in props.
	// Experimental.
	CdnAuthConfig() *CdnAuthConfiguration
	// Experimental.
	SetCdnAuthConfig(val *CdnAuthConfiguration)
	// The name of the channel group associated with the origin endpoint configuration.
	// Experimental.
	ChannelGroupName() *string
	// The channel name associated with the origin endpoint.
	// Experimental.
	ChannelName() *string
	// The timestamp of the creation of the origin endpoint.
	// Experimental.
	CreatedAt() *string
	// Experimental.
	DashManifests() *[]*awsmediapackagev2.CfnOriginEndpoint_DashManifestConfigurationProperty
	// Experimental.
	SetDashManifests(val *[]*awsmediapackagev2.CfnOriginEndpoint_DashManifestConfigurationProperty)
	// Array containing DASH Manifests created by the OriginEndpoint.
	// Experimental.
	DashManifestUrls() *[]*string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed in a Stack (those created by
	// creating new class instances like `new Role()`, `new Bucket()`, etc.), this
	// is always the same as the environment of the stack they belong to.
	//
	// For referenced resources (those obtained from referencing methods like
	// `Role.fromRoleArn()`, `Bucket.fromBucketName()`, etc.), they might be
	// different than the stack they were imported into.
	// Experimental.
	Env() *interfaces.ResourceEnvironment
	// Experimental.
	HlsManifests() *[]*awsmediapackagev2.CfnOriginEndpoint_HlsManifestConfigurationProperty
	// Experimental.
	SetHlsManifests(val *[]*awsmediapackagev2.CfnOriginEndpoint_HlsManifestConfigurationProperty)
	// Array containing HLS Manifests created by the OriginEndpoint.
	// Experimental.
	HlsManifestUrls() *[]*string
	// Experimental.
	LlHlsManifests() *[]*awsmediapackagev2.CfnOriginEndpoint_LowLatencyHlsManifestConfigurationProperty
	// Experimental.
	SetLlHlsManifests(val *[]*awsmediapackagev2.CfnOriginEndpoint_LowLatencyHlsManifestConfigurationProperty)
	// Array containing Low Latency HLS Manifests created by the OriginEndpoint.
	// Experimental.
	LowLatencyHlsManifestUrls() *[]*string
	// The timestamp of the modification of the origin endpoint.
	// Experimental.
	ModifiedAt() *string
	// Experimental.
	MssManifests() *[]*awsmediapackagev2.CfnOriginEndpoint_MssManifestConfigurationProperty
	// Experimental.
	SetMssManifests(val *[]*awsmediapackagev2.CfnOriginEndpoint_MssManifestConfigurationProperty)
	// Array containing MSS Manifests created by the OriginEndpoint.
	// Experimental.
	MssManifestUrls() *[]*string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The Amazon Resource Name (ARN) of the origin endpoint.
	// Experimental.
	OriginEndpointArn() *string
	// The name of the origin endpoint associated with the origin endpoint configuration.
	// Experimental.
	OriginEndpointName() *string
	// A reference to this Origin Endpoint resource.
	//
	// Required by the auto-generated IOriginEndpointRef interface.
	// Experimental.
	OriginEndpointRef() *interfacesawsmediapackagev2.OriginEndpointReference
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The resource policy associated with this origin endpoint.
	//
	// If `autoCreatePolicy` is true, an `OriginEndpointPolicy` will be created upon the
	// first call to addToResourcePolicy(s).
	// Experimental.
	Policy() OriginEndpointPolicy
	// Experimental.
	SetPolicy(val OriginEndpointPolicy)
	// Experimental.
	Segment() *SegmentConfiguration
	// Experimental.
	SetSegment(val *SegmentConfiguration)
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Configure origin endpoint policy.
	//
	// You can only add 1 OriginEndpointPolicy to an OriginEndpoint. If you have already
	// defined one, this will append to the policy already created.
	//
	// To configure CDN authentication, set `cdnAuth` on the construct's props instead.
	// Experimental.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Override the cross-stack reference strength for this resource.
	//
	// When set, any cross-stack reference to this resource will use the specified
	// mechanism instead of the global default determined by the
	// `@aws-cdk/core:defaultCrossStackReferences` context key. This is useful for
	// selectively weakening specific references to avoid the "deadly embrace" problem
	// without changing the app-wide default.
	// Experimental.
	ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength)
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Create a CloudWatch metric.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Egress Bytes.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricEgressBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Egress Request Count.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricEgressRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Egress Response time.
	// Default: - average over 60 seconds.
	//
	// Experimental.
	MetricEgressResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Ingress Bytes.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricIngressBytes(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Ingress Request Count.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricIngressRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Ingress response time.
	// Default: - average over 60 seconds.
	//
	// Experimental.
	MetricIngressResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Validate and modify Segment configuration for endpoint.
	// Experimental.
	SegmentValidation(segmentContainerType ContainerType, segment *SegmentConfiguration) *awsmediapackagev2.CfnOriginEndpoint_SegmentProperty
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for OriginEndpoint
type jsiiProxy_OriginEndpoint struct {
	internal.Type__awscdkResource
	jsiiProxy_IOriginEndpoint
}

func (j *jsiiProxy_OriginEndpoint) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) CdnAuthConfig() *CdnAuthConfiguration {
	var returns *CdnAuthConfiguration
	_jsii_.Get(
		j,
		"cdnAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) ChannelGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) ChannelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) DashManifests() *[]*awsmediapackagev2.CfnOriginEndpoint_DashManifestConfigurationProperty {
	var returns *[]*awsmediapackagev2.CfnOriginEndpoint_DashManifestConfigurationProperty
	_jsii_.Get(
		j,
		"dashManifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) DashManifestUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dashManifestUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) HlsManifests() *[]*awsmediapackagev2.CfnOriginEndpoint_HlsManifestConfigurationProperty {
	var returns *[]*awsmediapackagev2.CfnOriginEndpoint_HlsManifestConfigurationProperty
	_jsii_.Get(
		j,
		"hlsManifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) HlsManifestUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hlsManifestUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) LlHlsManifests() *[]*awsmediapackagev2.CfnOriginEndpoint_LowLatencyHlsManifestConfigurationProperty {
	var returns *[]*awsmediapackagev2.CfnOriginEndpoint_LowLatencyHlsManifestConfigurationProperty
	_jsii_.Get(
		j,
		"llHlsManifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) LowLatencyHlsManifestUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lowLatencyHlsManifestUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) MssManifests() *[]*awsmediapackagev2.CfnOriginEndpoint_MssManifestConfigurationProperty {
	var returns *[]*awsmediapackagev2.CfnOriginEndpoint_MssManifestConfigurationProperty
	_jsii_.Get(
		j,
		"mssManifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) MssManifestUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mssManifestUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) OriginEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) OriginEndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originEndpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) OriginEndpointRef() *interfacesawsmediapackagev2.OriginEndpointReference {
	var returns *interfacesawsmediapackagev2.OriginEndpointReference
	_jsii_.Get(
		j,
		"originEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) Policy() OriginEndpointPolicy {
	var returns OriginEndpointPolicy
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) Segment() *SegmentConfiguration {
	var returns *SegmentConfiguration
	_jsii_.Get(
		j,
		"segment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewOriginEndpoint(scope constructs.Construct, id *string, props *OriginEndpointProps) OriginEndpoint {
	_init_.Initialize()

	if err := validateNewOriginEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OriginEndpoint{}

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.OriginEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewOriginEndpoint_Override(o OriginEndpoint, scope constructs.Construct, id *string, props *OriginEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.OriginEndpoint",
		[]interface{}{scope, id, props},
		o,
	)
}

func (j *jsiiProxy_OriginEndpoint)SetAutoCreatePolicy(val *bool) {
	if err := j.validateSetAutoCreatePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoCreatePolicy",
		val,
	)
}

func (j *jsiiProxy_OriginEndpoint)SetCdnAuthConfig(val *CdnAuthConfiguration) {
	if err := j.validateSetCdnAuthConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdnAuthConfig",
		val,
	)
}

func (j *jsiiProxy_OriginEndpoint)SetDashManifests(val *[]*awsmediapackagev2.CfnOriginEndpoint_DashManifestConfigurationProperty) {
	if err := j.validateSetDashManifestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dashManifests",
		val,
	)
}

func (j *jsiiProxy_OriginEndpoint)SetHlsManifests(val *[]*awsmediapackagev2.CfnOriginEndpoint_HlsManifestConfigurationProperty) {
	if err := j.validateSetHlsManifestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hlsManifests",
		val,
	)
}

func (j *jsiiProxy_OriginEndpoint)SetLlHlsManifests(val *[]*awsmediapackagev2.CfnOriginEndpoint_LowLatencyHlsManifestConfigurationProperty) {
	if err := j.validateSetLlHlsManifestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"llHlsManifests",
		val,
	)
}

func (j *jsiiProxy_OriginEndpoint)SetMssManifests(val *[]*awsmediapackagev2.CfnOriginEndpoint_MssManifestConfigurationProperty) {
	if err := j.validateSetMssManifestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mssManifests",
		val,
	)
}

func (j *jsiiProxy_OriginEndpoint)SetPolicy(val OriginEndpointPolicy) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_OriginEndpoint)SetSegment(val *SegmentConfiguration) {
	if err := j.validateSetSegmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segment",
		val,
	)
}

// Creates an OriginEndpoint construct that represents an external (imported) Origin Endpoint from its ARN.
//
// The ARN is expected to be in the format:
// `arn:<partition>:mediapackagev2:<region>:<account>:channelGroup/<groupName>/channel/<channelName>/originEndpoint/<endpointName>`.
// Experimental.
func OriginEndpoint_FromOriginEndpointArn(scope constructs.Construct, id *string, originEndpointArn *string) IOriginEndpoint {
	_init_.Initialize()

	if err := validateOriginEndpoint_FromOriginEndpointArnParameters(scope, id, originEndpointArn); err != nil {
		panic(err)
	}
	var returns IOriginEndpoint

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.OriginEndpoint",
		"fromOriginEndpointArn",
		[]interface{}{scope, id, originEndpointArn},
		&returns,
	)

	return returns
}

// Creates an OriginEndpoint construct that represents an external (imported) Origin Endpoint.
// Experimental.
func OriginEndpoint_FromOriginEndpointAttributes(scope constructs.Construct, id *string, attrs *OriginEndpointAttributes) IOriginEndpoint {
	_init_.Initialize()

	if err := validateOriginEndpoint_FromOriginEndpointAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IOriginEndpoint

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.OriginEndpoint",
		"fromOriginEndpointAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func OriginEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOriginEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.OriginEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func OriginEndpoint_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateOriginEndpoint_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.OriginEndpoint",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func OriginEndpoint_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateOriginEndpoint_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.OriginEndpoint",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func OriginEndpoint_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediapackagev2-alpha.OriginEndpoint",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OriginEndpoint) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := o.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		o,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginEndpoint) ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength) {
	if err := o.validateApplyCrossStackReferenceStrengthParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"applyCrossStackReferenceStrength",
		[]interface{}{strength},
	)
}

func (o *jsiiProxy_OriginEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := o.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (o *jsiiProxy_OriginEndpoint) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginEndpoint) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := o.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginEndpoint) GetResourceNameAttribute(nameAttr *string) *string {
	if err := o.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginEndpoint) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginEndpoint) MetricEgressBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricEgressBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricEgressBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginEndpoint) MetricEgressRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricEgressRequestCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricEgressRequestCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginEndpoint) MetricEgressResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricEgressResponseTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricEgressResponseTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginEndpoint) MetricIngressBytes(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricIngressBytesParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricIngressBytes",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginEndpoint) MetricIngressRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricIngressRequestCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricIngressRequestCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginEndpoint) MetricIngressResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricIngressResponseTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricIngressResponseTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginEndpoint) SegmentValidation(segmentContainerType ContainerType, segment *SegmentConfiguration) *awsmediapackagev2.CfnOriginEndpoint_SegmentProperty {
	if err := o.validateSegmentValidationParameters(segmentContainerType, segment); err != nil {
		panic(err)
	}
	var returns *awsmediapackagev2.CfnOriginEndpoint_SegmentProperty

	_jsii_.Invoke(
		o,
		"segmentValidation",
		[]interface{}{segmentContainerType, segment},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginEndpoint) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		o,
		"with",
		args,
		&returns,
	)

	return returns
}

