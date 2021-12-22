package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A DNS A record.
//
// TODO: EXAMPLE
//
type ARecord interface {
	RecordSet
	DomainName() *string
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

// The jsii proxy struct for ARecord
type jsiiProxy_ARecord struct {
	jsiiProxy_RecordSet
}

func (j *jsiiProxy_ARecord) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ARecord) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ARecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ARecord) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ARecord) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewARecord(scope constructs.Construct, id *string, props *ARecordProps) ARecord {
	_init_.Initialize()

	j := jsiiProxy_ARecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.ARecord",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewARecord_Override(a ARecord, scope constructs.Construct, id *string, props *ARecordProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.ARecord",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ARecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.ARecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ARecord_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.ARecord",
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
func (a *jsiiProxy_ARecord) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_ARecord) GeneratePhysicalName() *string {
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
func (a *jsiiProxy_ARecord) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (a *jsiiProxy_ARecord) GetResourceNameAttribute(nameAttr *string) *string {
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
func (a *jsiiProxy_ARecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a ARecord.
//
// TODO: EXAMPLE
//
type ARecordProps struct {
	// A comment to add on the record.
	Comment *string `json:"comment"`
	// The domain name for this record.
	RecordName *string `json:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `json:"ttl"`
	// The hosted zone in which to define the new record.
	Zone IHostedZone `json:"zone"`
	// The target.
	Target RecordTarget `json:"target"`
}

// A DNS AAAA record.
//
// TODO: EXAMPLE
//
type AaaaRecord interface {
	RecordSet
	DomainName() *string
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

// The jsii proxy struct for AaaaRecord
type jsiiProxy_AaaaRecord struct {
	jsiiProxy_RecordSet
}

func (j *jsiiProxy_AaaaRecord) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AaaaRecord) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AaaaRecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AaaaRecord) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AaaaRecord) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewAaaaRecord(scope constructs.Construct, id *string, props *AaaaRecordProps) AaaaRecord {
	_init_.Initialize()

	j := jsiiProxy_AaaaRecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.AaaaRecord",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAaaaRecord_Override(a AaaaRecord, scope constructs.Construct, id *string, props *AaaaRecordProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.AaaaRecord",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AaaaRecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.AaaaRecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func AaaaRecord_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.AaaaRecord",
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
func (a *jsiiProxy_AaaaRecord) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_AaaaRecord) GeneratePhysicalName() *string {
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
func (a *jsiiProxy_AaaaRecord) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (a *jsiiProxy_AaaaRecord) GetResourceNameAttribute(nameAttr *string) *string {
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
func (a *jsiiProxy_AaaaRecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a AaaaRecord.
//
// TODO: EXAMPLE
//
type AaaaRecordProps struct {
	// A comment to add on the record.
	Comment *string `json:"comment"`
	// The domain name for this record.
	RecordName *string `json:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `json:"ttl"`
	// The hosted zone in which to define the new record.
	Zone IHostedZone `json:"zone"`
	// The target.
	Target RecordTarget `json:"target"`
}

// Represents the properties of an alias target destination.
//
// TODO: EXAMPLE
//
type AliasRecordTargetConfig struct {
	// DNS name of the target.
	DnsName *string `json:"dnsName"`
	// Hosted zone ID of the target.
	HostedZoneId *string `json:"hostedZoneId"`
}

// A DNS Amazon CAA record.
//
// A CAA record to restrict certificate authorities allowed
// to issue certificates for a domain to Amazon only.
//
// TODO: EXAMPLE
//
type CaaAmazonRecord interface {
	CaaRecord
	DomainName() *string
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

// The jsii proxy struct for CaaAmazonRecord
type jsiiProxy_CaaAmazonRecord struct {
	jsiiProxy_CaaRecord
}

func (j *jsiiProxy_CaaAmazonRecord) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CaaAmazonRecord) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CaaAmazonRecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CaaAmazonRecord) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CaaAmazonRecord) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCaaAmazonRecord(scope constructs.Construct, id *string, props *CaaAmazonRecordProps) CaaAmazonRecord {
	_init_.Initialize()

	j := jsiiProxy_CaaAmazonRecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CaaAmazonRecord",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCaaAmazonRecord_Override(c CaaAmazonRecord, scope constructs.Construct, id *string, props *CaaAmazonRecordProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CaaAmazonRecord",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CaaAmazonRecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CaaAmazonRecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func CaaAmazonRecord_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CaaAmazonRecord",
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
func (c *jsiiProxy_CaaAmazonRecord) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CaaAmazonRecord) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
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
func (c *jsiiProxy_CaaAmazonRecord) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
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
func (c *jsiiProxy_CaaAmazonRecord) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CaaAmazonRecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a CaaAmazonRecord.
//
// TODO: EXAMPLE
//
type CaaAmazonRecordProps struct {
	// A comment to add on the record.
	Comment *string `json:"comment"`
	// The domain name for this record.
	RecordName *string `json:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `json:"ttl"`
	// The hosted zone in which to define the new record.
	Zone IHostedZone `json:"zone"`
}

// A DNS CAA record.
//
// TODO: EXAMPLE
//
type CaaRecord interface {
	RecordSet
	DomainName() *string
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

// The jsii proxy struct for CaaRecord
type jsiiProxy_CaaRecord struct {
	jsiiProxy_RecordSet
}

func (j *jsiiProxy_CaaRecord) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CaaRecord) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CaaRecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CaaRecord) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CaaRecord) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCaaRecord(scope constructs.Construct, id *string, props *CaaRecordProps) CaaRecord {
	_init_.Initialize()

	j := jsiiProxy_CaaRecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CaaRecord",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCaaRecord_Override(c CaaRecord, scope constructs.Construct, id *string, props *CaaRecordProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CaaRecord",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CaaRecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CaaRecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func CaaRecord_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CaaRecord",
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
func (c *jsiiProxy_CaaRecord) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CaaRecord) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
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
func (c *jsiiProxy_CaaRecord) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
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
func (c *jsiiProxy_CaaRecord) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CaaRecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a CaaRecord.
//
// TODO: EXAMPLE
//
type CaaRecordProps struct {
	// A comment to add on the record.
	Comment *string `json:"comment"`
	// The domain name for this record.
	RecordName *string `json:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `json:"ttl"`
	// The hosted zone in which to define the new record.
	Zone IHostedZone `json:"zone"`
	// The values.
	Values *[]*CaaRecordValue `json:"values"`
}

// Properties for a CAA record value.
//
// TODO: EXAMPLE
//
type CaaRecordValue struct {
	// The flag.
	Flag *float64 `json:"flag"`
	// The tag.
	Tag CaaTag `json:"tag"`
	// The value associated with the tag.
	Value *string `json:"value"`
}

// The CAA tag.
type CaaTag string

const (
	CaaTag_IODEF CaaTag = "IODEF"
	CaaTag_ISSUE CaaTag = "ISSUE"
	CaaTag_ISSUEWILD CaaTag = "ISSUEWILD"
)

// A CloudFormation `AWS::Route53::DNSSEC`.
//
// TODO: EXAMPLE
//
type CfnDNSSEC interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	HostedZoneId() *string
	SetHostedZoneId(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDNSSEC
type jsiiProxy_CfnDNSSEC struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDNSSEC) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDNSSEC) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDNSSEC) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDNSSEC) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDNSSEC) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDNSSEC) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDNSSEC) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDNSSEC) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDNSSEC) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDNSSEC) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53::DNSSEC`.
func NewCfnDNSSEC(scope constructs.Construct, id *string, props *CfnDNSSECProps) CfnDNSSEC {
	_init_.Initialize()

	j := jsiiProxy_CfnDNSSEC{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnDNSSEC",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53::DNSSEC`.
func NewCfnDNSSEC_Override(c CfnDNSSEC, scope constructs.Construct, id *string, props *CfnDNSSECProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnDNSSEC",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDNSSEC) SetHostedZoneId(val *string) {
	_jsii_.Set(
		j,
		"hostedZoneId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDNSSEC_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnDNSSEC",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDNSSEC_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnDNSSEC",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDNSSEC_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnDNSSEC",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDNSSEC_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53.CfnDNSSEC",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDNSSEC) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnDNSSEC) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDNSSEC) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnDNSSEC) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDNSSEC) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDNSSEC) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDNSSEC) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnDNSSEC) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDNSSEC) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDNSSEC) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDNSSEC) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDNSSEC) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnDNSSEC) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnDNSSEC) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDNSSEC) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::Route53::DNSSEC`.
//
// TODO: EXAMPLE
//
type CfnDNSSECProps struct {
	// `AWS::Route53::DNSSEC.HostedZoneId`.
	HostedZoneId *string `json:"hostedZoneId"`
}

// A CloudFormation `AWS::Route53::HealthCheck`.
//
// TODO: EXAMPLE
//
type CfnHealthCheck interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrHealthCheckId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	HealthCheckConfig() interface{}
	SetHealthCheckConfig(val interface{})
	HealthCheckTags() interface{}
	SetHealthCheckTags(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnHealthCheck
type jsiiProxy_CfnHealthCheck struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnHealthCheck) AttrHealthCheckId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrHealthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHealthCheck) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHealthCheck) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHealthCheck) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHealthCheck) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHealthCheck) HealthCheckConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHealthCheck) HealthCheckTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHealthCheck) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHealthCheck) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHealthCheck) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHealthCheck) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHealthCheck) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53::HealthCheck`.
func NewCfnHealthCheck(scope constructs.Construct, id *string, props *CfnHealthCheckProps) CfnHealthCheck {
	_init_.Initialize()

	j := jsiiProxy_CfnHealthCheck{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnHealthCheck",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53::HealthCheck`.
func NewCfnHealthCheck_Override(c CfnHealthCheck, scope constructs.Construct, id *string, props *CfnHealthCheckProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnHealthCheck",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnHealthCheck) SetHealthCheckConfig(val interface{}) {
	_jsii_.Set(
		j,
		"healthCheckConfig",
		val,
	)
}

func (j *jsiiProxy_CfnHealthCheck) SetHealthCheckTags(val interface{}) {
	_jsii_.Set(
		j,
		"healthCheckTags",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnHealthCheck_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnHealthCheck",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnHealthCheck_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnHealthCheck",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnHealthCheck_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnHealthCheck",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHealthCheck_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53.CfnHealthCheck",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnHealthCheck) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnHealthCheck) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnHealthCheck) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnHealthCheck) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnHealthCheck) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnHealthCheck) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnHealthCheck) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnHealthCheck) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnHealthCheck) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnHealthCheck) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnHealthCheck) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnHealthCheck) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnHealthCheck) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnHealthCheck) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHealthCheck) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnHealthCheck_AlarmIdentifierProperty struct {
	// `CfnHealthCheck.AlarmIdentifierProperty.Name`.
	Name *string `json:"name"`
	// `CfnHealthCheck.AlarmIdentifierProperty.Region`.
	Region *string `json:"region"`
}

// TODO: EXAMPLE
//
type CfnHealthCheck_HealthCheckConfigProperty struct {
	// `CfnHealthCheck.HealthCheckConfigProperty.AlarmIdentifier`.
	AlarmIdentifier interface{} `json:"alarmIdentifier"`
	// `CfnHealthCheck.HealthCheckConfigProperty.ChildHealthChecks`.
	ChildHealthChecks *[]*string `json:"childHealthChecks"`
	// `CfnHealthCheck.HealthCheckConfigProperty.EnableSNI`.
	EnableSni interface{} `json:"enableSni"`
	// `CfnHealthCheck.HealthCheckConfigProperty.FailureThreshold`.
	FailureThreshold *float64 `json:"failureThreshold"`
	// `CfnHealthCheck.HealthCheckConfigProperty.FullyQualifiedDomainName`.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName"`
	// `CfnHealthCheck.HealthCheckConfigProperty.HealthThreshold`.
	HealthThreshold *float64 `json:"healthThreshold"`
	// `CfnHealthCheck.HealthCheckConfigProperty.InsufficientDataHealthStatus`.
	InsufficientDataHealthStatus *string `json:"insufficientDataHealthStatus"`
	// `CfnHealthCheck.HealthCheckConfigProperty.Inverted`.
	Inverted interface{} `json:"inverted"`
	// `CfnHealthCheck.HealthCheckConfigProperty.IPAddress`.
	IpAddress *string `json:"ipAddress"`
	// `CfnHealthCheck.HealthCheckConfigProperty.MeasureLatency`.
	MeasureLatency interface{} `json:"measureLatency"`
	// `CfnHealthCheck.HealthCheckConfigProperty.Port`.
	Port *float64 `json:"port"`
	// `CfnHealthCheck.HealthCheckConfigProperty.Regions`.
	Regions *[]*string `json:"regions"`
	// `CfnHealthCheck.HealthCheckConfigProperty.RequestInterval`.
	RequestInterval *float64 `json:"requestInterval"`
	// `CfnHealthCheck.HealthCheckConfigProperty.ResourcePath`.
	ResourcePath *string `json:"resourcePath"`
	// `CfnHealthCheck.HealthCheckConfigProperty.SearchString`.
	SearchString *string `json:"searchString"`
	// `CfnHealthCheck.HealthCheckConfigProperty.Type`.
	Type *string `json:"type"`
}

// TODO: EXAMPLE
//
type CfnHealthCheck_HealthCheckTagProperty struct {
	// `CfnHealthCheck.HealthCheckTagProperty.Key`.
	Key *string `json:"key"`
	// `CfnHealthCheck.HealthCheckTagProperty.Value`.
	Value *string `json:"value"`
}

// Properties for defining a `AWS::Route53::HealthCheck`.
//
// TODO: EXAMPLE
//
type CfnHealthCheckProps struct {
	// `AWS::Route53::HealthCheck.HealthCheckConfig`.
	HealthCheckConfig interface{} `json:"healthCheckConfig"`
	// `AWS::Route53::HealthCheck.HealthCheckTags`.
	HealthCheckTags interface{} `json:"healthCheckTags"`
}

// A CloudFormation `AWS::Route53::HostedZone`.
//
// TODO: EXAMPLE
//
type CfnHostedZone interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
	AttrNameServers() *[]*string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	HostedZoneConfig() interface{}
	SetHostedZoneConfig(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	QueryLoggingConfig() interface{}
	SetQueryLoggingConfig(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	Vpcs() interface{}
	SetVpcs(val interface{})
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnHostedZone
type jsiiProxy_CfnHostedZone struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnHostedZone) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedZone) AttrNameServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrNameServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedZone) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedZone) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedZone) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedZone) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedZone) HostedZoneConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostedZoneConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedZone) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedZone) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedZone) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedZone) QueryLoggingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryLoggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedZone) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedZone) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedZone) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedZone) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedZone) Vpcs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcs",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53::HostedZone`.
func NewCfnHostedZone(scope constructs.Construct, id *string, props *CfnHostedZoneProps) CfnHostedZone {
	_init_.Initialize()

	j := jsiiProxy_CfnHostedZone{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnHostedZone",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53::HostedZone`.
func NewCfnHostedZone_Override(c CfnHostedZone, scope constructs.Construct, id *string, props *CfnHostedZoneProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnHostedZone",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnHostedZone) SetHostedZoneConfig(val interface{}) {
	_jsii_.Set(
		j,
		"hostedZoneConfig",
		val,
	)
}

func (j *jsiiProxy_CfnHostedZone) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnHostedZone) SetQueryLoggingConfig(val interface{}) {
	_jsii_.Set(
		j,
		"queryLoggingConfig",
		val,
	)
}

func (j *jsiiProxy_CfnHostedZone) SetVpcs(val interface{}) {
	_jsii_.Set(
		j,
		"vpcs",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnHostedZone_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnHostedZone",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnHostedZone_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnHostedZone",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnHostedZone_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnHostedZone",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHostedZone_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53.CfnHostedZone",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnHostedZone) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnHostedZone) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnHostedZone) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnHostedZone) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnHostedZone) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnHostedZone) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnHostedZone) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnHostedZone) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnHostedZone) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnHostedZone) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnHostedZone) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnHostedZone) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnHostedZone) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnHostedZone) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHostedZone) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnHostedZone_HostedZoneConfigProperty struct {
	// `CfnHostedZone.HostedZoneConfigProperty.Comment`.
	Comment *string `json:"comment"`
}

// TODO: EXAMPLE
//
type CfnHostedZone_HostedZoneTagProperty struct {
	// `CfnHostedZone.HostedZoneTagProperty.Key`.
	Key *string `json:"key"`
	// `CfnHostedZone.HostedZoneTagProperty.Value`.
	Value *string `json:"value"`
}

// TODO: EXAMPLE
//
type CfnHostedZone_QueryLoggingConfigProperty struct {
	// `CfnHostedZone.QueryLoggingConfigProperty.CloudWatchLogsLogGroupArn`.
	CloudWatchLogsLogGroupArn *string `json:"cloudWatchLogsLogGroupArn"`
}

// TODO: EXAMPLE
//
type CfnHostedZone_VPCProperty struct {
	// `CfnHostedZone.VPCProperty.VPCId`.
	VpcId *string `json:"vpcId"`
	// `CfnHostedZone.VPCProperty.VPCRegion`.
	VpcRegion *string `json:"vpcRegion"`
}

// Properties for defining a `AWS::Route53::HostedZone`.
//
// TODO: EXAMPLE
//
type CfnHostedZoneProps struct {
	// `AWS::Route53::HostedZone.HostedZoneConfig`.
	HostedZoneConfig interface{} `json:"hostedZoneConfig"`
	// `AWS::Route53::HostedZone.HostedZoneTags`.
	HostedZoneTags *[]*CfnHostedZone_HostedZoneTagProperty `json:"hostedZoneTags"`
	// `AWS::Route53::HostedZone.Name`.
	Name *string `json:"name"`
	// `AWS::Route53::HostedZone.QueryLoggingConfig`.
	QueryLoggingConfig interface{} `json:"queryLoggingConfig"`
	// `AWS::Route53::HostedZone.VPCs`.
	Vpcs interface{} `json:"vpcs"`
}

// A CloudFormation `AWS::Route53::KeySigningKey`.
//
// TODO: EXAMPLE
//
type CfnKeySigningKey interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	HostedZoneId() *string
	SetHostedZoneId(val *string)
	KeyManagementServiceArn() *string
	SetKeyManagementServiceArn(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Status() *string
	SetStatus(val *string)
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnKeySigningKey
type jsiiProxy_CfnKeySigningKey struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnKeySigningKey) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeySigningKey) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeySigningKey) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeySigningKey) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeySigningKey) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeySigningKey) KeyManagementServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyManagementServiceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeySigningKey) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeySigningKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeySigningKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeySigningKey) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeySigningKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeySigningKey) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeySigningKey) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53::KeySigningKey`.
func NewCfnKeySigningKey(scope constructs.Construct, id *string, props *CfnKeySigningKeyProps) CfnKeySigningKey {
	_init_.Initialize()

	j := jsiiProxy_CfnKeySigningKey{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnKeySigningKey",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53::KeySigningKey`.
func NewCfnKeySigningKey_Override(c CfnKeySigningKey, scope constructs.Construct, id *string, props *CfnKeySigningKeyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnKeySigningKey",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnKeySigningKey) SetHostedZoneId(val *string) {
	_jsii_.Set(
		j,
		"hostedZoneId",
		val,
	)
}

func (j *jsiiProxy_CfnKeySigningKey) SetKeyManagementServiceArn(val *string) {
	_jsii_.Set(
		j,
		"keyManagementServiceArn",
		val,
	)
}

func (j *jsiiProxy_CfnKeySigningKey) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnKeySigningKey) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnKeySigningKey_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnKeySigningKey",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnKeySigningKey_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnKeySigningKey",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnKeySigningKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnKeySigningKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnKeySigningKey_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53.CfnKeySigningKey",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnKeySigningKey) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnKeySigningKey) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnKeySigningKey) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnKeySigningKey) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnKeySigningKey) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnKeySigningKey) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnKeySigningKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnKeySigningKey) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnKeySigningKey) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnKeySigningKey) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnKeySigningKey) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnKeySigningKey) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnKeySigningKey) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnKeySigningKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnKeySigningKey) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::Route53::KeySigningKey`.
//
// TODO: EXAMPLE
//
type CfnKeySigningKeyProps struct {
	// `AWS::Route53::KeySigningKey.HostedZoneId`.
	HostedZoneId *string `json:"hostedZoneId"`
	// `AWS::Route53::KeySigningKey.KeyManagementServiceArn`.
	KeyManagementServiceArn *string `json:"keyManagementServiceArn"`
	// `AWS::Route53::KeySigningKey.Name`.
	Name *string `json:"name"`
	// `AWS::Route53::KeySigningKey.Status`.
	Status *string `json:"status"`
}

// A CloudFormation `AWS::Route53::RecordSet`.
//
// TODO: EXAMPLE
//
type CfnRecordSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AliasTarget() interface{}
	SetAliasTarget(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Comment() *string
	SetComment(val *string)
	CreationStack() *[]*string
	Failover() *string
	SetFailover(val *string)
	GeoLocation() interface{}
	SetGeoLocation(val interface{})
	HealthCheckId() *string
	SetHealthCheckId(val *string)
	HostedZoneId() *string
	SetHostedZoneId(val *string)
	HostedZoneName() *string
	SetHostedZoneName(val *string)
	LogicalId() *string
	MultiValueAnswer() interface{}
	SetMultiValueAnswer(val interface{})
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Region() *string
	SetRegion(val *string)
	ResourceRecords() *[]*string
	SetResourceRecords(val *[]*string)
	SetIdentifier() *string
	SetSetIdentifier(val *string)
	Stack() awscdk.Stack
	Ttl() *string
	SetTtl(val *string)
	Type() *string
	SetType(val *string)
	UpdatedProperites() *map[string]interface{}
	Weight() *float64
	SetWeight(val *float64)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnRecordSet
type jsiiProxy_CfnRecordSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRecordSet) AliasTarget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aliasTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Failover() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) GeoLocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) HealthCheckId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) HostedZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) MultiValueAnswer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiValueAnswer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) ResourceRecords() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) SetIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53::RecordSet`.
func NewCfnRecordSet(scope constructs.Construct, id *string, props *CfnRecordSetProps) CfnRecordSet {
	_init_.Initialize()

	j := jsiiProxy_CfnRecordSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53::RecordSet`.
func NewCfnRecordSet_Override(c CfnRecordSet, scope constructs.Construct, id *string, props *CfnRecordSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRecordSet) SetAliasTarget(val interface{}) {
	_jsii_.Set(
		j,
		"aliasTarget",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet) SetComment(val *string) {
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet) SetFailover(val *string) {
	_jsii_.Set(
		j,
		"failover",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet) SetGeoLocation(val interface{}) {
	_jsii_.Set(
		j,
		"geoLocation",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet) SetHealthCheckId(val *string) {
	_jsii_.Set(
		j,
		"healthCheckId",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet) SetHostedZoneId(val *string) {
	_jsii_.Set(
		j,
		"hostedZoneId",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet) SetHostedZoneName(val *string) {
	_jsii_.Set(
		j,
		"hostedZoneName",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet) SetMultiValueAnswer(val interface{}) {
	_jsii_.Set(
		j,
		"multiValueAnswer",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet) SetResourceRecords(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceRecords",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet) SetSetIdentifier(val *string) {
	_jsii_.Set(
		j,
		"setIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet) SetTtl(val *string) {
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet) SetWeight(val *float64) {
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRecordSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRecordSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnRecordSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRecordSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnRecordSet) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnRecordSet) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnRecordSet) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnRecordSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnRecordSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnRecordSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnRecordSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnRecordSet) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnRecordSet) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnRecordSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnRecordSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRecordSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnRecordSet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnRecordSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRecordSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnRecordSet_AliasTargetProperty struct {
	// `CfnRecordSet.AliasTargetProperty.DNSName`.
	DnsName *string `json:"dnsName"`
	// `CfnRecordSet.AliasTargetProperty.EvaluateTargetHealth`.
	EvaluateTargetHealth interface{} `json:"evaluateTargetHealth"`
	// `CfnRecordSet.AliasTargetProperty.HostedZoneId`.
	HostedZoneId *string `json:"hostedZoneId"`
}

// TODO: EXAMPLE
//
type CfnRecordSet_GeoLocationProperty struct {
	// `CfnRecordSet.GeoLocationProperty.ContinentCode`.
	ContinentCode *string `json:"continentCode"`
	// `CfnRecordSet.GeoLocationProperty.CountryCode`.
	CountryCode *string `json:"countryCode"`
	// `CfnRecordSet.GeoLocationProperty.SubdivisionCode`.
	SubdivisionCode *string `json:"subdivisionCode"`
}

// A CloudFormation `AWS::Route53::RecordSetGroup`.
//
// TODO: EXAMPLE
//
type CfnRecordSetGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Comment() *string
	SetComment(val *string)
	CreationStack() *[]*string
	HostedZoneId() *string
	SetHostedZoneId(val *string)
	HostedZoneName() *string
	SetHostedZoneName(val *string)
	LogicalId() *string
	Node() constructs.Node
	RecordSets() interface{}
	SetRecordSets(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnRecordSetGroup
type jsiiProxy_CfnRecordSetGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRecordSetGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSetGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSetGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSetGroup) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSetGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSetGroup) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSetGroup) HostedZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSetGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSetGroup) RecordSets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSetGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSetGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSetGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53::RecordSetGroup`.
func NewCfnRecordSetGroup(scope constructs.Construct, id *string, props *CfnRecordSetGroupProps) CfnRecordSetGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnRecordSetGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnRecordSetGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53::RecordSetGroup`.
func NewCfnRecordSetGroup_Override(c CfnRecordSetGroup, scope constructs.Construct, id *string, props *CfnRecordSetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnRecordSetGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRecordSetGroup) SetComment(val *string) {
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSetGroup) SetHostedZoneId(val *string) {
	_jsii_.Set(
		j,
		"hostedZoneId",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSetGroup) SetHostedZoneName(val *string) {
	_jsii_.Set(
		j,
		"hostedZoneName",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSetGroup) SetRecordSets(val interface{}) {
	_jsii_.Set(
		j,
		"recordSets",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRecordSetGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnRecordSetGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRecordSetGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnRecordSetGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnRecordSetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnRecordSetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRecordSetGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53.CfnRecordSetGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnRecordSetGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnRecordSetGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnRecordSetGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnRecordSetGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnRecordSetGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnRecordSetGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnRecordSetGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnRecordSetGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnRecordSetGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnRecordSetGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnRecordSetGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRecordSetGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnRecordSetGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnRecordSetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRecordSetGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnRecordSetGroup_AliasTargetProperty struct {
	// `CfnRecordSetGroup.AliasTargetProperty.DNSName`.
	DnsName *string `json:"dnsName"`
	// `CfnRecordSetGroup.AliasTargetProperty.EvaluateTargetHealth`.
	EvaluateTargetHealth interface{} `json:"evaluateTargetHealth"`
	// `CfnRecordSetGroup.AliasTargetProperty.HostedZoneId`.
	HostedZoneId *string `json:"hostedZoneId"`
}

// TODO: EXAMPLE
//
type CfnRecordSetGroup_GeoLocationProperty struct {
	// `CfnRecordSetGroup.GeoLocationProperty.ContinentCode`.
	ContinentCode *string `json:"continentCode"`
	// `CfnRecordSetGroup.GeoLocationProperty.CountryCode`.
	CountryCode *string `json:"countryCode"`
	// `CfnRecordSetGroup.GeoLocationProperty.SubdivisionCode`.
	SubdivisionCode *string `json:"subdivisionCode"`
}

// TODO: EXAMPLE
//
type CfnRecordSetGroup_RecordSetProperty struct {
	// `CfnRecordSetGroup.RecordSetProperty.AliasTarget`.
	AliasTarget interface{} `json:"aliasTarget"`
	// `CfnRecordSetGroup.RecordSetProperty.Comment`.
	Comment *string `json:"comment"`
	// `CfnRecordSetGroup.RecordSetProperty.Failover`.
	Failover *string `json:"failover"`
	// `CfnRecordSetGroup.RecordSetProperty.GeoLocation`.
	GeoLocation interface{} `json:"geoLocation"`
	// `CfnRecordSetGroup.RecordSetProperty.HealthCheckId`.
	HealthCheckId *string `json:"healthCheckId"`
	// `CfnRecordSetGroup.RecordSetProperty.HostedZoneId`.
	HostedZoneId *string `json:"hostedZoneId"`
	// `CfnRecordSetGroup.RecordSetProperty.HostedZoneName`.
	HostedZoneName *string `json:"hostedZoneName"`
	// `CfnRecordSetGroup.RecordSetProperty.MultiValueAnswer`.
	MultiValueAnswer interface{} `json:"multiValueAnswer"`
	// `CfnRecordSetGroup.RecordSetProperty.Name`.
	Name *string `json:"name"`
	// `CfnRecordSetGroup.RecordSetProperty.Region`.
	Region *string `json:"region"`
	// `CfnRecordSetGroup.RecordSetProperty.ResourceRecords`.
	ResourceRecords *[]*string `json:"resourceRecords"`
	// `CfnRecordSetGroup.RecordSetProperty.SetIdentifier`.
	SetIdentifier *string `json:"setIdentifier"`
	// `CfnRecordSetGroup.RecordSetProperty.TTL`.
	Ttl *string `json:"ttl"`
	// `CfnRecordSetGroup.RecordSetProperty.Type`.
	Type *string `json:"type"`
	// `CfnRecordSetGroup.RecordSetProperty.Weight`.
	Weight *float64 `json:"weight"`
}

// Properties for defining a `AWS::Route53::RecordSetGroup`.
//
// TODO: EXAMPLE
//
type CfnRecordSetGroupProps struct {
	// `AWS::Route53::RecordSetGroup.Comment`.
	Comment *string `json:"comment"`
	// `AWS::Route53::RecordSetGroup.HostedZoneId`.
	HostedZoneId *string `json:"hostedZoneId"`
	// `AWS::Route53::RecordSetGroup.HostedZoneName`.
	HostedZoneName *string `json:"hostedZoneName"`
	// `AWS::Route53::RecordSetGroup.RecordSets`.
	RecordSets interface{} `json:"recordSets"`
}

// Properties for defining a `AWS::Route53::RecordSet`.
//
// TODO: EXAMPLE
//
type CfnRecordSetProps struct {
	// `AWS::Route53::RecordSet.AliasTarget`.
	AliasTarget interface{} `json:"aliasTarget"`
	// `AWS::Route53::RecordSet.Comment`.
	Comment *string `json:"comment"`
	// `AWS::Route53::RecordSet.Failover`.
	Failover *string `json:"failover"`
	// `AWS::Route53::RecordSet.GeoLocation`.
	GeoLocation interface{} `json:"geoLocation"`
	// `AWS::Route53::RecordSet.HealthCheckId`.
	HealthCheckId *string `json:"healthCheckId"`
	// `AWS::Route53::RecordSet.HostedZoneId`.
	HostedZoneId *string `json:"hostedZoneId"`
	// `AWS::Route53::RecordSet.HostedZoneName`.
	HostedZoneName *string `json:"hostedZoneName"`
	// `AWS::Route53::RecordSet.MultiValueAnswer`.
	MultiValueAnswer interface{} `json:"multiValueAnswer"`
	// `AWS::Route53::RecordSet.Name`.
	Name *string `json:"name"`
	// `AWS::Route53::RecordSet.Region`.
	Region *string `json:"region"`
	// `AWS::Route53::RecordSet.ResourceRecords`.
	ResourceRecords *[]*string `json:"resourceRecords"`
	// `AWS::Route53::RecordSet.SetIdentifier`.
	SetIdentifier *string `json:"setIdentifier"`
	// `AWS::Route53::RecordSet.TTL`.
	Ttl *string `json:"ttl"`
	// `AWS::Route53::RecordSet.Type`.
	Type *string `json:"type"`
	// `AWS::Route53::RecordSet.Weight`.
	Weight *float64 `json:"weight"`
}

// A DNS CNAME record.
//
// TODO: EXAMPLE
//
type CnameRecord interface {
	RecordSet
	DomainName() *string
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

// The jsii proxy struct for CnameRecord
type jsiiProxy_CnameRecord struct {
	jsiiProxy_RecordSet
}

func (j *jsiiProxy_CnameRecord) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CnameRecord) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CnameRecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CnameRecord) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CnameRecord) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCnameRecord(scope constructs.Construct, id *string, props *CnameRecordProps) CnameRecord {
	_init_.Initialize()

	j := jsiiProxy_CnameRecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CnameRecord",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCnameRecord_Override(c CnameRecord, scope constructs.Construct, id *string, props *CnameRecordProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CnameRecord",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CnameRecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CnameRecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func CnameRecord_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CnameRecord",
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
func (c *jsiiProxy_CnameRecord) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CnameRecord) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
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
func (c *jsiiProxy_CnameRecord) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
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
func (c *jsiiProxy_CnameRecord) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CnameRecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a CnameRecord.
//
// TODO: EXAMPLE
//
type CnameRecordProps struct {
	// A comment to add on the record.
	Comment *string `json:"comment"`
	// The domain name for this record.
	RecordName *string `json:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `json:"ttl"`
	// The hosted zone in which to define the new record.
	Zone IHostedZone `json:"zone"`
	// The domain name.
	DomainName *string `json:"domainName"`
}

// Common properties to create a Route 53 hosted zone.
//
// TODO: EXAMPLE
//
type CommonHostedZoneProps struct {
	// Any comments that you want to include about the hosted zone.
	Comment *string `json:"comment"`
	// The Amazon Resource Name (ARN) for the log group that you want Amazon Route 53 to send query logs to.
	QueryLogsLogGroupArn *string `json:"queryLogsLogGroupArn"`
	// The name of the domain.
	//
	// For resource record types that include a domain
	// name, specify a fully qualified domain name.
	ZoneName *string `json:"zoneName"`
}

// A Cross Account Zone Delegation record.
//
// TODO: EXAMPLE
//
type CrossAccountZoneDelegationRecord interface {
	constructs.Construct
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for CrossAccountZoneDelegationRecord
type jsiiProxy_CrossAccountZoneDelegationRecord struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CrossAccountZoneDelegationRecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewCrossAccountZoneDelegationRecord(scope constructs.Construct, id *string, props *CrossAccountZoneDelegationRecordProps) CrossAccountZoneDelegationRecord {
	_init_.Initialize()

	j := jsiiProxy_CrossAccountZoneDelegationRecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CrossAccountZoneDelegationRecord",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCrossAccountZoneDelegationRecord_Override(c CrossAccountZoneDelegationRecord, scope constructs.Construct, id *string, props *CrossAccountZoneDelegationRecordProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CrossAccountZoneDelegationRecord",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CrossAccountZoneDelegationRecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CrossAccountZoneDelegationRecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CrossAccountZoneDelegationRecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a CrossAccountZoneDelegationRecord.
//
// TODO: EXAMPLE
//
type CrossAccountZoneDelegationRecordProps struct {
	// The zone to be delegated.
	DelegatedZone IHostedZone `json:"delegatedZone"`
	// The delegation role in the parent account.
	DelegationRole awsiam.IRole `json:"delegationRole"`
	// The hosted zone id in the parent account.
	ParentHostedZoneId *string `json:"parentHostedZoneId"`
	// The hosted zone name in the parent account.
	ParentHostedZoneName *string `json:"parentHostedZoneName"`
	// The removal policy to apply to the record set.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `json:"ttl"`
}

// A DNS DS record.
//
// TODO: EXAMPLE
//
type DsRecord interface {
	RecordSet
	DomainName() *string
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

// The jsii proxy struct for DsRecord
type jsiiProxy_DsRecord struct {
	jsiiProxy_RecordSet
}

func (j *jsiiProxy_DsRecord) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DsRecord) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DsRecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DsRecord) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DsRecord) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewDsRecord(scope constructs.Construct, id *string, props *DsRecordProps) DsRecord {
	_init_.Initialize()

	j := jsiiProxy_DsRecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.DsRecord",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDsRecord_Override(d DsRecord, scope constructs.Construct, id *string, props *DsRecordProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.DsRecord",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DsRecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.DsRecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DsRecord_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.DsRecord",
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
func (d *jsiiProxy_DsRecord) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DsRecord) GeneratePhysicalName() *string {
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
func (d *jsiiProxy_DsRecord) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (d *jsiiProxy_DsRecord) GetResourceNameAttribute(nameAttr *string) *string {
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
func (d *jsiiProxy_DsRecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a DSRecord.
//
// TODO: EXAMPLE
//
type DsRecordProps struct {
	// A comment to add on the record.
	Comment *string `json:"comment"`
	// The domain name for this record.
	RecordName *string `json:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `json:"ttl"`
	// The hosted zone in which to define the new record.
	Zone IHostedZone `json:"zone"`
	// The DS values.
	Values *[]*string `json:"values"`
}

// Container for records, and records contain information about how to route traffic for a specific domain, such as example.com and its subdomains (acme.example.com, zenith.example.com).
//
// TODO: EXAMPLE
//
type HostedZone interface {
	awscdk.Resource
	IHostedZone
	Env() *awscdk.ResourceEnvironment
	HostedZoneArn() *string
	HostedZoneId() *string
	HostedZoneNameServers() *[]*string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	Vpcs() *[]*CfnHostedZone_VPCProperty
	ZoneName() *string
	AddVpc(vpc awsec2.IVpc)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for HostedZone
type jsiiProxy_HostedZone struct {
	internal.Type__awscdkResource
	jsiiProxy_IHostedZone
}

func (j *jsiiProxy_HostedZone) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedZone) HostedZoneArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedZone) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedZone) HostedZoneNameServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostedZoneNameServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedZone) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedZone) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedZone) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedZone) Vpcs() *[]*CfnHostedZone_VPCProperty {
	var returns *[]*CfnHostedZone_VPCProperty
	_jsii_.Get(
		j,
		"vpcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedZone) ZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneName",
		&returns,
	)
	return returns
}


func NewHostedZone(scope constructs.Construct, id *string, props *HostedZoneProps) HostedZone {
	_init_.Initialize()

	j := jsiiProxy_HostedZone{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.HostedZone",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewHostedZone_Override(h HostedZone, scope constructs.Construct, id *string, props *HostedZoneProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.HostedZone",
		[]interface{}{scope, id, props},
		h,
	)
}

// Imports a hosted zone from another stack.
//
// Use when both hosted zone ID and hosted zone name are known.
func HostedZone_FromHostedZoneAttributes(scope constructs.Construct, id *string, attrs *HostedZoneAttributes) IHostedZone {
	_init_.Initialize()

	var returns IHostedZone

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.HostedZone",
		"fromHostedZoneAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import a Route 53 hosted zone defined either outside the CDK, or in a different CDK stack.
//
// Use when hosted zone ID is known. Hosted zone name becomes unavailable through this query.
func HostedZone_FromHostedZoneId(scope constructs.Construct, id *string, hostedZoneId *string) IHostedZone {
	_init_.Initialize()

	var returns IHostedZone

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.HostedZone",
		"fromHostedZoneId",
		[]interface{}{scope, id, hostedZoneId},
		&returns,
	)

	return returns
}

// Lookup a hosted zone in the current account/region based on query parameters.
//
// Requires environment, you must specify env for the stack.
//
// Use to easily query hosted zones.
// See: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
//
func HostedZone_FromLookup(scope constructs.Construct, id *string, query *HostedZoneProviderProps) IHostedZone {
	_init_.Initialize()

	var returns IHostedZone

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.HostedZone",
		"fromLookup",
		[]interface{}{scope, id, query},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func HostedZone_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.HostedZone",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func HostedZone_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.HostedZone",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add another VPC to this private hosted zone.
func (h *jsiiProxy_HostedZone) AddVpc(vpc awsec2.IVpc) {
	_jsii_.InvokeVoid(
		h,
		"addVpc",
		[]interface{}{vpc},
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
func (h *jsiiProxy_HostedZone) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		h,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (h *jsiiProxy_HostedZone) GeneratePhysicalName() *string {
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
func (h *jsiiProxy_HostedZone) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (h *jsiiProxy_HostedZone) GetResourceNameAttribute(nameAttr *string) *string {
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
func (h *jsiiProxy_HostedZone) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Reference to a hosted zone.
//
// TODO: EXAMPLE
//
type HostedZoneAttributes struct {
	// Identifier of the hosted zone.
	HostedZoneId *string `json:"hostedZoneId"`
	// Name of the hosted zone.
	ZoneName *string `json:"zoneName"`
}

// Properties of a new hosted zone.
//
// TODO: EXAMPLE
//
type HostedZoneProps struct {
	// Any comments that you want to include about the hosted zone.
	Comment *string `json:"comment"`
	// The Amazon Resource Name (ARN) for the log group that you want Amazon Route 53 to send query logs to.
	QueryLogsLogGroupArn *string `json:"queryLogsLogGroupArn"`
	// The name of the domain.
	//
	// For resource record types that include a domain
	// name, specify a fully qualified domain name.
	ZoneName *string `json:"zoneName"`
	// A VPC that you want to associate with this hosted zone.
	//
	// When you specify
	// this property, a private hosted zone will be created.
	//
	// You can associate additional VPCs to this private zone using `addVpc(vpc)`.
	Vpcs *[]awsec2.IVpc `json:"vpcs"`
}

// Zone properties for looking up the Hosted Zone.
//
// TODO: EXAMPLE
//
type HostedZoneProviderProps struct {
	// The zone domain e.g. example.com.
	DomainName *string `json:"domainName"`
	// Whether the zone that is being looked up is a private hosted zone.
	PrivateZone *bool `json:"privateZone"`
	// Specifies the ID of the VPC associated with a private hosted zone.
	//
	// If a VPC ID is provided and privateZone is false, no results will be returned
	// and an error will be raised
	VpcId *string `json:"vpcId"`
}

// Classes that are valid alias record targets, like CloudFront distributions and load balancers, should implement this interface.
type IAliasRecordTarget interface {
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(record IRecordSet, zone IHostedZone) *AliasRecordTargetConfig
}

// The jsii proxy for IAliasRecordTarget
type jsiiProxy_IAliasRecordTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_IAliasRecordTarget) Bind(record IRecordSet, zone IHostedZone) *AliasRecordTargetConfig {
	var returns *AliasRecordTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{record, zone},
		&returns,
	)

	return returns
}

// Imported or created hosted zone.
type IHostedZone interface {
	awscdk.IResource
	// ARN of this hosted zone, such as arn:${Partition}:route53:::hostedzone/${Id}.
	HostedZoneArn() *string
	// ID of this hosted zone, such as "Z23ABC4XYZL05B".
	HostedZoneId() *string
	// Returns the set of name servers for the specific hosted zone. For example: ns1.example.com.
	//
	// This attribute will be undefined for private hosted zones or hosted zones imported from another stack.
	HostedZoneNameServers() *[]*string
	// FQDN of this hosted zone.
	ZoneName() *string
}

// The jsii proxy for IHostedZone
type jsiiProxy_IHostedZone struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IHostedZone) HostedZoneArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHostedZone) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHostedZone) HostedZoneNameServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostedZoneNameServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHostedZone) ZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneName",
		&returns,
	)
	return returns
}

// Represents a Route 53 private hosted zone.
type IPrivateHostedZone interface {
	IHostedZone
}

// The jsii proxy for IPrivateHostedZone
type jsiiProxy_IPrivateHostedZone struct {
	jsiiProxy_IHostedZone
}

// Represents a Route 53 public hosted zone.
type IPublicHostedZone interface {
	IHostedZone
}

// The jsii proxy for IPublicHostedZone
type jsiiProxy_IPublicHostedZone struct {
	jsiiProxy_IHostedZone
}

// A record set.
type IRecordSet interface {
	awscdk.IResource
	// The domain name of the record.
	DomainName() *string
}

// The jsii proxy for IRecordSet
type jsiiProxy_IRecordSet struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRecordSet) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

// A DNS MX record.
//
// TODO: EXAMPLE
//
type MxRecord interface {
	RecordSet
	DomainName() *string
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

// The jsii proxy struct for MxRecord
type jsiiProxy_MxRecord struct {
	jsiiProxy_RecordSet
}

func (j *jsiiProxy_MxRecord) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MxRecord) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MxRecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MxRecord) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MxRecord) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewMxRecord(scope constructs.Construct, id *string, props *MxRecordProps) MxRecord {
	_init_.Initialize()

	j := jsiiProxy_MxRecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.MxRecord",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewMxRecord_Override(m MxRecord, scope constructs.Construct, id *string, props *MxRecordProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.MxRecord",
		[]interface{}{scope, id, props},
		m,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func MxRecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.MxRecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func MxRecord_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.MxRecord",
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
func (m *jsiiProxy_MxRecord) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		m,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (m *jsiiProxy_MxRecord) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		m,
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
func (m *jsiiProxy_MxRecord) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		m,
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
func (m *jsiiProxy_MxRecord) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (m *jsiiProxy_MxRecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a MxRecord.
//
// TODO: EXAMPLE
//
type MxRecordProps struct {
	// A comment to add on the record.
	Comment *string `json:"comment"`
	// The domain name for this record.
	RecordName *string `json:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `json:"ttl"`
	// The hosted zone in which to define the new record.
	Zone IHostedZone `json:"zone"`
	// The values.
	Values *[]*MxRecordValue `json:"values"`
}

// Properties for a MX record value.
//
// TODO: EXAMPLE
//
type MxRecordValue struct {
	// The mail server host name.
	HostName *string `json:"hostName"`
	// The priority.
	Priority *float64 `json:"priority"`
}

// A DNS NS record.
//
// TODO: EXAMPLE
//
type NsRecord interface {
	RecordSet
	DomainName() *string
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

// The jsii proxy struct for NsRecord
type jsiiProxy_NsRecord struct {
	jsiiProxy_RecordSet
}

func (j *jsiiProxy_NsRecord) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsRecord) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsRecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsRecord) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsRecord) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewNsRecord(scope constructs.Construct, id *string, props *NsRecordProps) NsRecord {
	_init_.Initialize()

	j := jsiiProxy_NsRecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.NsRecord",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNsRecord_Override(n NsRecord, scope constructs.Construct, id *string, props *NsRecordProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.NsRecord",
		[]interface{}{scope, id, props},
		n,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func NsRecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.NsRecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func NsRecord_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.NsRecord",
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
func (n *jsiiProxy_NsRecord) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (n *jsiiProxy_NsRecord) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		n,
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
func (n *jsiiProxy_NsRecord) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		n,
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
func (n *jsiiProxy_NsRecord) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (n *jsiiProxy_NsRecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a NSRecord.
//
// TODO: EXAMPLE
//
type NsRecordProps struct {
	// A comment to add on the record.
	Comment *string `json:"comment"`
	// The domain name for this record.
	RecordName *string `json:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `json:"ttl"`
	// The hosted zone in which to define the new record.
	Zone IHostedZone `json:"zone"`
	// The NS values.
	Values *[]*string `json:"values"`
}

// Create a Route53 private hosted zone for use in one or more VPCs.
//
// Note that `enableDnsHostnames` and `enableDnsSupport` must have been enabled
// for the VPC you're configuring for private hosted zones.
//
// TODO: EXAMPLE
//
type PrivateHostedZone interface {
	HostedZone
	IPrivateHostedZone
	Env() *awscdk.ResourceEnvironment
	HostedZoneArn() *string
	HostedZoneId() *string
	HostedZoneNameServers() *[]*string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	Vpcs() *[]*CfnHostedZone_VPCProperty
	ZoneName() *string
	AddVpc(vpc awsec2.IVpc)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for PrivateHostedZone
type jsiiProxy_PrivateHostedZone struct {
	jsiiProxy_HostedZone
	jsiiProxy_IPrivateHostedZone
}

func (j *jsiiProxy_PrivateHostedZone) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateHostedZone) HostedZoneArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateHostedZone) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateHostedZone) HostedZoneNameServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostedZoneNameServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateHostedZone) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateHostedZone) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateHostedZone) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateHostedZone) Vpcs() *[]*CfnHostedZone_VPCProperty {
	var returns *[]*CfnHostedZone_VPCProperty
	_jsii_.Get(
		j,
		"vpcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateHostedZone) ZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneName",
		&returns,
	)
	return returns
}


func NewPrivateHostedZone(scope constructs.Construct, id *string, props *PrivateHostedZoneProps) PrivateHostedZone {
	_init_.Initialize()

	j := jsiiProxy_PrivateHostedZone{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.PrivateHostedZone",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPrivateHostedZone_Override(p PrivateHostedZone, scope constructs.Construct, id *string, props *PrivateHostedZoneProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.PrivateHostedZone",
		[]interface{}{scope, id, props},
		p,
	)
}

// Imports a hosted zone from another stack.
//
// Use when both hosted zone ID and hosted zone name are known.
func PrivateHostedZone_FromHostedZoneAttributes(scope constructs.Construct, id *string, attrs *HostedZoneAttributes) IHostedZone {
	_init_.Initialize()

	var returns IHostedZone

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.PrivateHostedZone",
		"fromHostedZoneAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import a Route 53 hosted zone defined either outside the CDK, or in a different CDK stack.
//
// Use when hosted zone ID is known. Hosted zone name becomes unavailable through this query.
func PrivateHostedZone_FromHostedZoneId(scope constructs.Construct, id *string, hostedZoneId *string) IHostedZone {
	_init_.Initialize()

	var returns IHostedZone

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.PrivateHostedZone",
		"fromHostedZoneId",
		[]interface{}{scope, id, hostedZoneId},
		&returns,
	)

	return returns
}

// Lookup a hosted zone in the current account/region based on query parameters.
//
// Requires environment, you must specify env for the stack.
//
// Use to easily query hosted zones.
// See: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
//
func PrivateHostedZone_FromLookup(scope constructs.Construct, id *string, query *HostedZoneProviderProps) IHostedZone {
	_init_.Initialize()

	var returns IHostedZone

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.PrivateHostedZone",
		"fromLookup",
		[]interface{}{scope, id, query},
		&returns,
	)

	return returns
}

// Import a Route 53 private hosted zone defined either outside the CDK, or in a different CDK stack.
func PrivateHostedZone_FromPrivateHostedZoneId(scope constructs.Construct, id *string, privateHostedZoneId *string) IPrivateHostedZone {
	_init_.Initialize()

	var returns IPrivateHostedZone

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.PrivateHostedZone",
		"fromPrivateHostedZoneId",
		[]interface{}{scope, id, privateHostedZoneId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PrivateHostedZone_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.PrivateHostedZone",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func PrivateHostedZone_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.PrivateHostedZone",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add another VPC to this private hosted zone.
func (p *jsiiProxy_PrivateHostedZone) AddVpc(vpc awsec2.IVpc) {
	_jsii_.InvokeVoid(
		p,
		"addVpc",
		[]interface{}{vpc},
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
func (p *jsiiProxy_PrivateHostedZone) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_PrivateHostedZone) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PrivateHostedZone) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PrivateHostedZone) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PrivateHostedZone) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to create a Route 53 private hosted zone.
//
// TODO: EXAMPLE
//
type PrivateHostedZoneProps struct {
	// Any comments that you want to include about the hosted zone.
	Comment *string `json:"comment"`
	// The Amazon Resource Name (ARN) for the log group that you want Amazon Route 53 to send query logs to.
	QueryLogsLogGroupArn *string `json:"queryLogsLogGroupArn"`
	// The name of the domain.
	//
	// For resource record types that include a domain
	// name, specify a fully qualified domain name.
	ZoneName *string `json:"zoneName"`
	// A VPC that you want to associate with this hosted zone.
	//
	// Private hosted zones must be associated with at least one VPC. You can
	// associated additional VPCs using `addVpc(vpc)`.
	Vpc awsec2.IVpc `json:"vpc"`
}

// Create a Route53 public hosted zone.
//
// TODO: EXAMPLE
//
type PublicHostedZone interface {
	HostedZone
	IPublicHostedZone
	CrossAccountZoneDelegationRole() awsiam.Role
	Env() *awscdk.ResourceEnvironment
	HostedZoneArn() *string
	HostedZoneId() *string
	HostedZoneNameServers() *[]*string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	Vpcs() *[]*CfnHostedZone_VPCProperty
	ZoneName() *string
	AddDelegation(delegate IPublicHostedZone, opts *ZoneDelegationOptions)
	AddVpc(_vpc awsec2.IVpc)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for PublicHostedZone
type jsiiProxy_PublicHostedZone struct {
	jsiiProxy_HostedZone
	jsiiProxy_IPublicHostedZone
}

func (j *jsiiProxy_PublicHostedZone) CrossAccountZoneDelegationRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"crossAccountZoneDelegationRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicHostedZone) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicHostedZone) HostedZoneArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicHostedZone) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicHostedZone) HostedZoneNameServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostedZoneNameServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicHostedZone) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicHostedZone) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicHostedZone) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicHostedZone) Vpcs() *[]*CfnHostedZone_VPCProperty {
	var returns *[]*CfnHostedZone_VPCProperty
	_jsii_.Get(
		j,
		"vpcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicHostedZone) ZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneName",
		&returns,
	)
	return returns
}


func NewPublicHostedZone(scope constructs.Construct, id *string, props *PublicHostedZoneProps) PublicHostedZone {
	_init_.Initialize()

	j := jsiiProxy_PublicHostedZone{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.PublicHostedZone",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPublicHostedZone_Override(p PublicHostedZone, scope constructs.Construct, id *string, props *PublicHostedZoneProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.PublicHostedZone",
		[]interface{}{scope, id, props},
		p,
	)
}

// Imports a hosted zone from another stack.
//
// Use when both hosted zone ID and hosted zone name are known.
func PublicHostedZone_FromHostedZoneAttributes(scope constructs.Construct, id *string, attrs *HostedZoneAttributes) IHostedZone {
	_init_.Initialize()

	var returns IHostedZone

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.PublicHostedZone",
		"fromHostedZoneAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import a Route 53 hosted zone defined either outside the CDK, or in a different CDK stack.
//
// Use when hosted zone ID is known. Hosted zone name becomes unavailable through this query.
func PublicHostedZone_FromHostedZoneId(scope constructs.Construct, id *string, hostedZoneId *string) IHostedZone {
	_init_.Initialize()

	var returns IHostedZone

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.PublicHostedZone",
		"fromHostedZoneId",
		[]interface{}{scope, id, hostedZoneId},
		&returns,
	)

	return returns
}

// Lookup a hosted zone in the current account/region based on query parameters.
//
// Requires environment, you must specify env for the stack.
//
// Use to easily query hosted zones.
// See: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
//
func PublicHostedZone_FromLookup(scope constructs.Construct, id *string, query *HostedZoneProviderProps) IHostedZone {
	_init_.Initialize()

	var returns IHostedZone

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.PublicHostedZone",
		"fromLookup",
		[]interface{}{scope, id, query},
		&returns,
	)

	return returns
}

// Import a Route 53 public hosted zone defined either outside the CDK, or in a different CDK stack.
func PublicHostedZone_FromPublicHostedZoneId(scope constructs.Construct, id *string, publicHostedZoneId *string) IPublicHostedZone {
	_init_.Initialize()

	var returns IPublicHostedZone

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.PublicHostedZone",
		"fromPublicHostedZoneId",
		[]interface{}{scope, id, publicHostedZoneId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PublicHostedZone_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.PublicHostedZone",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func PublicHostedZone_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.PublicHostedZone",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a delegation from this zone to a designated zone.
func (p *jsiiProxy_PublicHostedZone) AddDelegation(delegate IPublicHostedZone, opts *ZoneDelegationOptions) {
	_jsii_.InvokeVoid(
		p,
		"addDelegation",
		[]interface{}{delegate, opts},
	)
}

// Add another VPC to this private hosted zone.
func (p *jsiiProxy_PublicHostedZone) AddVpc(_vpc awsec2.IVpc) {
	_jsii_.InvokeVoid(
		p,
		"addVpc",
		[]interface{}{_vpc},
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
func (p *jsiiProxy_PublicHostedZone) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_PublicHostedZone) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PublicHostedZone) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PublicHostedZone) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PublicHostedZone) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a PublicHostedZone.
//
// TODO: EXAMPLE
//
type PublicHostedZoneProps struct {
	// Any comments that you want to include about the hosted zone.
	Comment *string `json:"comment"`
	// The Amazon Resource Name (ARN) for the log group that you want Amazon Route 53 to send query logs to.
	QueryLogsLogGroupArn *string `json:"queryLogsLogGroupArn"`
	// The name of the domain.
	//
	// For resource record types that include a domain
	// name, specify a fully qualified domain name.
	ZoneName *string `json:"zoneName"`
	// Whether to create a CAA record to restrict certificate authorities allowed to issue certificates for this domain to Amazon only.
	CaaAmazon *bool `json:"caaAmazon"`
	// A principal which is trusted to assume a role for zone delegation.
	CrossAccountZoneDelegationPrincipal awsiam.IPrincipal `json:"crossAccountZoneDelegationPrincipal"`
	// The name of the role created for cross account delegation.
	CrossAccountZoneDelegationRoleName *string `json:"crossAccountZoneDelegationRoleName"`
}

// A record set.
//
// TODO: EXAMPLE
//
type RecordSet interface {
	awscdk.Resource
	IRecordSet
	DomainName() *string
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

// The jsii proxy struct for RecordSet
type jsiiProxy_RecordSet struct {
	internal.Type__awscdkResource
	jsiiProxy_IRecordSet
}

func (j *jsiiProxy_RecordSet) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecordSet) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecordSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecordSet) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecordSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewRecordSet(scope constructs.Construct, id *string, props *RecordSetProps) RecordSet {
	_init_.Initialize()

	j := jsiiProxy_RecordSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.RecordSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRecordSet_Override(r RecordSet, scope constructs.Construct, id *string, props *RecordSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.RecordSet",
		[]interface{}{scope, id, props},
		r,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func RecordSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.RecordSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func RecordSet_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.RecordSet",
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
func (r *jsiiProxy_RecordSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_RecordSet) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
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
func (r *jsiiProxy_RecordSet) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
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
func (r *jsiiProxy_RecordSet) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_RecordSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for a RecordSet.
//
// TODO: EXAMPLE
//
type RecordSetOptions struct {
	// A comment to add on the record.
	Comment *string `json:"comment"`
	// The domain name for this record.
	RecordName *string `json:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `json:"ttl"`
	// The hosted zone in which to define the new record.
	Zone IHostedZone `json:"zone"`
}

// Construction properties for a RecordSet.
//
// TODO: EXAMPLE
//
type RecordSetProps struct {
	// A comment to add on the record.
	Comment *string `json:"comment"`
	// The domain name for this record.
	RecordName *string `json:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `json:"ttl"`
	// The hosted zone in which to define the new record.
	Zone IHostedZone `json:"zone"`
	// The record type.
	RecordType RecordType `json:"recordType"`
	// The target for this record, either `RecordTarget.fromValues()` or `RecordTarget.fromAlias()`.
	Target RecordTarget `json:"target"`
}

// Type union for a record that accepts multiple types of target.
//
// TODO: EXAMPLE
//
type RecordTarget interface {
	AliasTarget() IAliasRecordTarget
	Values() *[]*string
}

// The jsii proxy struct for RecordTarget
type jsiiProxy_RecordTarget struct {
	_ byte // padding
}

func (j *jsiiProxy_RecordTarget) AliasTarget() IAliasRecordTarget {
	var returns IAliasRecordTarget
	_jsii_.Get(
		j,
		"aliasTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecordTarget) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}


func NewRecordTarget(values *[]*string, aliasTarget IAliasRecordTarget) RecordTarget {
	_init_.Initialize()

	j := jsiiProxy_RecordTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.RecordTarget",
		[]interface{}{values, aliasTarget},
		&j,
	)

	return &j
}

func NewRecordTarget_Override(r RecordTarget, values *[]*string, aliasTarget IAliasRecordTarget) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.RecordTarget",
		[]interface{}{values, aliasTarget},
		r,
	)
}

// Use an alias as target.
func RecordTarget_FromAlias(aliasTarget IAliasRecordTarget) RecordTarget {
	_init_.Initialize()

	var returns RecordTarget

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.RecordTarget",
		"fromAlias",
		[]interface{}{aliasTarget},
		&returns,
	)

	return returns
}

// Use ip addresses as target.
func RecordTarget_FromIpAddresses(ipAddresses ...*string) RecordTarget {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range ipAddresses {
		args = append(args, a)
	}

	var returns RecordTarget

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.RecordTarget",
		"fromIpAddresses",
		args,
		&returns,
	)

	return returns
}

// Use string values as target.
func RecordTarget_FromValues(values ...*string) RecordTarget {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range values {
		args = append(args, a)
	}

	var returns RecordTarget

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.RecordTarget",
		"fromValues",
		args,
		&returns,
	)

	return returns
}

// The record type.
type RecordType string

const (
	RecordType_A RecordType = "A"
	RecordType_AAAA RecordType = "AAAA"
	RecordType_CAA RecordType = "CAA"
	RecordType_CNAME RecordType = "CNAME"
	RecordType_DS RecordType = "DS"
	RecordType_MX RecordType = "MX"
	RecordType_NAPTR RecordType = "NAPTR"
	RecordType_NS RecordType = "NS"
	RecordType_PTR RecordType = "PTR"
	RecordType_SOA RecordType = "SOA"
	RecordType_SPF RecordType = "SPF"
	RecordType_SRV RecordType = "SRV"
	RecordType_TXT RecordType = "TXT"
)

// A DNS SRV record.
//
// TODO: EXAMPLE
//
type SrvRecord interface {
	RecordSet
	DomainName() *string
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

// The jsii proxy struct for SrvRecord
type jsiiProxy_SrvRecord struct {
	jsiiProxy_RecordSet
}

func (j *jsiiProxy_SrvRecord) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SrvRecord) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SrvRecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SrvRecord) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SrvRecord) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewSrvRecord(scope constructs.Construct, id *string, props *SrvRecordProps) SrvRecord {
	_init_.Initialize()

	j := jsiiProxy_SrvRecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.SrvRecord",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSrvRecord_Override(s SrvRecord, scope constructs.Construct, id *string, props *SrvRecordProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.SrvRecord",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SrvRecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.SrvRecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func SrvRecord_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.SrvRecord",
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
func (s *jsiiProxy_SrvRecord) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_SrvRecord) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
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
func (s *jsiiProxy_SrvRecord) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
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
func (s *jsiiProxy_SrvRecord) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_SrvRecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a SrvRecord.
//
// TODO: EXAMPLE
//
type SrvRecordProps struct {
	// A comment to add on the record.
	Comment *string `json:"comment"`
	// The domain name for this record.
	RecordName *string `json:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `json:"ttl"`
	// The hosted zone in which to define the new record.
	Zone IHostedZone `json:"zone"`
	// The values.
	Values *[]*SrvRecordValue `json:"values"`
}

// Properties for a SRV record value.
//
// TODO: EXAMPLE
//
type SrvRecordValue struct {
	// The server host name.
	HostName *string `json:"hostName"`
	// The port.
	Port *float64 `json:"port"`
	// The priority.
	Priority *float64 `json:"priority"`
	// The weight.
	Weight *float64 `json:"weight"`
}

// A DNS TXT record.
//
// TODO: EXAMPLE
//
type TxtRecord interface {
	RecordSet
	DomainName() *string
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

// The jsii proxy struct for TxtRecord
type jsiiProxy_TxtRecord struct {
	jsiiProxy_RecordSet
}

func (j *jsiiProxy_TxtRecord) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TxtRecord) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TxtRecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TxtRecord) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TxtRecord) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewTxtRecord(scope constructs.Construct, id *string, props *TxtRecordProps) TxtRecord {
	_init_.Initialize()

	j := jsiiProxy_TxtRecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.TxtRecord",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewTxtRecord_Override(t TxtRecord, scope constructs.Construct, id *string, props *TxtRecordProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.TxtRecord",
		[]interface{}{scope, id, props},
		t,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func TxtRecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.TxtRecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func TxtRecord_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.TxtRecord",
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
func (t *jsiiProxy_TxtRecord) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_TxtRecord) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
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
func (t *jsiiProxy_TxtRecord) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		t,
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
func (t *jsiiProxy_TxtRecord) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (t *jsiiProxy_TxtRecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a TxtRecord.
//
// TODO: EXAMPLE
//
type TxtRecordProps struct {
	// A comment to add on the record.
	Comment *string `json:"comment"`
	// The domain name for this record.
	RecordName *string `json:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `json:"ttl"`
	// The hosted zone in which to define the new record.
	Zone IHostedZone `json:"zone"`
	// The text values.
	Values *[]*string `json:"values"`
}

// A Private DNS configuration for a VPC endpoint service.
//
// TODO: EXAMPLE
//
type VpcEndpointServiceDomainName interface {
	constructs.Construct
	DomainName() *string
	SetDomainName(val *string)
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for VpcEndpointServiceDomainName
type jsiiProxy_VpcEndpointServiceDomainName struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_VpcEndpointServiceDomainName) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointServiceDomainName) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewVpcEndpointServiceDomainName(scope constructs.Construct, id *string, props *VpcEndpointServiceDomainNameProps) VpcEndpointServiceDomainName {
	_init_.Initialize()

	j := jsiiProxy_VpcEndpointServiceDomainName{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.VpcEndpointServiceDomainName",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewVpcEndpointServiceDomainName_Override(v VpcEndpointServiceDomainName, scope constructs.Construct, id *string, props *VpcEndpointServiceDomainNameProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.VpcEndpointServiceDomainName",
		[]interface{}{scope, id, props},
		v,
	)
}

func (j *jsiiProxy_VpcEndpointServiceDomainName) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func VpcEndpointServiceDomainName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.VpcEndpointServiceDomainName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (v *jsiiProxy_VpcEndpointServiceDomainName) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to configure a VPC Endpoint Service domain name.
//
// TODO: EXAMPLE
//
type VpcEndpointServiceDomainNameProps struct {
	// The domain name to use.
	//
	// This domain name must be owned by this account (registered through Route53),
	// or delegated to this account. Domain ownership will be verified by AWS before
	// private DNS can be used.
	// See: https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-services-dns-validation.html
	//
	DomainName *string `json:"domainName"`
	// The VPC Endpoint Service to configure Private DNS for.
	EndpointService awsec2.IVpcEndpointService `json:"endpointService"`
	// The public hosted zone to use for the domain.
	PublicHostedZone IPublicHostedZone `json:"publicHostedZone"`
}

// Options available when creating a delegation relationship from one PublicHostedZone to another.
//
// TODO: EXAMPLE
//
type ZoneDelegationOptions struct {
	// A comment to add on the DNS record created to incorporate the delegation.
	Comment *string `json:"comment"`
	// The TTL (Time To Live) of the DNS delegation record in DNS caches.
	Ttl awscdk.Duration `json:"ttl"`
}

// A record to delegate further lookups to a different set of name servers.
//
// TODO: EXAMPLE
//
type ZoneDelegationRecord interface {
	RecordSet
	DomainName() *string
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

// The jsii proxy struct for ZoneDelegationRecord
type jsiiProxy_ZoneDelegationRecord struct {
	jsiiProxy_RecordSet
}

func (j *jsiiProxy_ZoneDelegationRecord) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDelegationRecord) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDelegationRecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDelegationRecord) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDelegationRecord) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewZoneDelegationRecord(scope constructs.Construct, id *string, props *ZoneDelegationRecordProps) ZoneDelegationRecord {
	_init_.Initialize()

	j := jsiiProxy_ZoneDelegationRecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.ZoneDelegationRecord",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewZoneDelegationRecord_Override(z ZoneDelegationRecord, scope constructs.Construct, id *string, props *ZoneDelegationRecordProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.ZoneDelegationRecord",
		[]interface{}{scope, id, props},
		z,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ZoneDelegationRecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.ZoneDelegationRecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ZoneDelegationRecord_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.ZoneDelegationRecord",
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
func (z *jsiiProxy_ZoneDelegationRecord) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		z,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (z *jsiiProxy_ZoneDelegationRecord) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		z,
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
func (z *jsiiProxy_ZoneDelegationRecord) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		z,
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
func (z *jsiiProxy_ZoneDelegationRecord) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (z *jsiiProxy_ZoneDelegationRecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a ZoneDelegationRecord.
//
// TODO: EXAMPLE
//
type ZoneDelegationRecordProps struct {
	// A comment to add on the record.
	Comment *string `json:"comment"`
	// The domain name for this record.
	RecordName *string `json:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `json:"ttl"`
	// The hosted zone in which to define the new record.
	Zone IHostedZone `json:"zone"`
	// The name servers to report in the delegation records.
	NameServers *[]*string `json:"nameServers"`
}

