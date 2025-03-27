package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Information about the record that you want to create.
//
// The `AWS::Route53::RecordSet` type can be used as a standalone resource or as an embedded property in the `AWS::Route53::RecordSetGroup` type. Note that some `AWS::Route53::RecordSet` properties are valid only when used within `AWS::Route53::RecordSetGroup` .
//
// For more information, see [ChangeResourceRecordSets](https://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html) in the *Amazon Route 53 API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecordSet := awscdk.Aws_route53.NewCfnRecordSet(this, jsii.String("MyCfnRecordSet"), &CfnRecordSetProps{
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	AliasTarget: &AliasTargetProperty{
//   		DnsName: jsii.String("dnsName"),
//   		HostedZoneId: jsii.String("hostedZoneId"),
//
//   		// the properties below are optional
//   		EvaluateTargetHealth: jsii.Boolean(false),
//   	},
//   	CidrRoutingConfig: &CidrRoutingConfigProperty{
//   		CollectionId: jsii.String("collectionId"),
//   		LocationName: jsii.String("locationName"),
//   	},
//   	Comment: jsii.String("comment"),
//   	Failover: jsii.String("failover"),
//   	GeoLocation: &GeoLocationProperty{
//   		ContinentCode: jsii.String("continentCode"),
//   		CountryCode: jsii.String("countryCode"),
//   		SubdivisionCode: jsii.String("subdivisionCode"),
//   	},
//   	GeoProximityLocation: &GeoProximityLocationProperty{
//   		AwsRegion: jsii.String("awsRegion"),
//   		Bias: jsii.Number(123),
//   		Coordinates: &CoordinatesProperty{
//   			Latitude: jsii.String("latitude"),
//   			Longitude: jsii.String("longitude"),
//   		},
//   		LocalZoneGroup: jsii.String("localZoneGroup"),
//   	},
//   	HealthCheckId: jsii.String("healthCheckId"),
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	HostedZoneName: jsii.String("hostedZoneName"),
//   	MultiValueAnswer: jsii.Boolean(false),
//   	Region: jsii.String("region"),
//   	ResourceRecords: []*string{
//   		jsii.String("resourceRecords"),
//   	},
//   	SetIdentifier: jsii.String("setIdentifier"),
//   	Ttl: jsii.String("ttl"),
//   	Weight: jsii.Number(123),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordset.html
//
type CfnRecordSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// *Alias resource record sets only:* Information about the AWS resource, such as a CloudFront distribution or an Amazon S3 bucket, that you want to route traffic to.
	AliasTarget() interface{}
	SetAliasTarget(val interface{})
	// Specifies a coordinate of the east–west position of a geographic point on the surface of the Earth.
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The object that is specified in resource record set object when you are linking a resource record set to a CIDR location.
	CidrRoutingConfig() interface{}
	SetCidrRoutingConfig(val interface{})
	// *Optional:* Any comments you want to include about a change batch request.
	Comment() *string
	SetComment(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// *Failover resource record sets only:* To configure failover, you add the `Failover` element to two resource record sets.
	Failover() *string
	SetFailover(val *string)
	// *Geolocation resource record sets only:* A complex type that lets you control how Amazon Route 53 responds to DNS queries based on the geographic origin of the query.
	GeoLocation() interface{}
	SetGeoLocation(val interface{})
	// *GeoproximityLocation resource record sets only:* A complex type that lets you control how Route 53 responds to DNS queries based on the geographic origin of the query and your resources.
	GeoProximityLocation() interface{}
	SetGeoProximityLocation(val interface{})
	// If you want Amazon Route 53 to return this resource record set in response to a DNS query only when the status of a health check is healthy, include the `HealthCheckId` element and specify the ID of the applicable health check.
	HealthCheckId() *string
	SetHealthCheckId(val *string)
	// The ID of the hosted zone that you want to create records in.
	HostedZoneId() *string
	SetHostedZoneId(val *string)
	// The name of the hosted zone that you want to create records in.
	HostedZoneName() *string
	SetHostedZoneName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// *Multivalue answer resource record sets only* : To route traffic approximately randomly to multiple resources, such as web servers, create one multivalue answer record for each resource and specify `true` for `MultiValueAnswer` .
	MultiValueAnswer() interface{}
	SetMultiValueAnswer(val interface{})
	// For `ChangeResourceRecordSets` requests, the name of the record that you want to create, update, or delete.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// *Latency-based resource record sets only:* The Amazon EC2 Region where you created the resource that this resource record set refers to.
	Region() *string
	SetRegion(val *string)
	// One or more values that correspond with the value that you specified for the `Type` property.
	ResourceRecords() *[]*string
	SetResourceRecords(val *[]*string)
	// *Resource record sets that have a routing policy other than simple:* An identifier that differentiates among multiple resource record sets that have the same combination of name and type, such as multiple weighted resource record sets named acme.example.com that have a type of A. In a group of resource record sets that have the same name and type, the value of `SetIdentifier` must be unique for each resource record set.
	SetIdentifier() *string
	SetSetIdentifier(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The resource record cache time to live (TTL), in seconds.
	//
	// Note the following:.
	Ttl() *string
	SetTtl(val *string)
	// The DNS record type.
	Type() *string
	SetType(val *string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// *Weighted resource record sets only:* Among resource record sets that have the same combination of DNS name and type, a value that determines the proportion of DNS queries that Amazon Route 53 responds to using the current resource record set.
	Weight() *float64
	SetWeight(val *float64)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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

func (j *jsiiProxy_CfnRecordSet) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
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

func (j *jsiiProxy_CfnRecordSet) CidrRoutingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cidrRoutingConfig",
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

func (j *jsiiProxy_CfnRecordSet) GeoProximityLocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoProximityLocation",
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

func (j *jsiiProxy_CfnRecordSet) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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


func NewCfnRecordSet(scope constructs.Construct, id *string, props *CfnRecordSetProps) CfnRecordSet {
	_init_.Initialize()

	if err := validateNewCfnRecordSetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRecordSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnRecordSet_Override(c CfnRecordSet, scope constructs.Construct, id *string, props *CfnRecordSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetAliasTarget(val interface{}) {
	if err := j.validateSetAliasTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasTarget",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetCidrRoutingConfig(val interface{}) {
	if err := j.validateSetCidrRoutingConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrRoutingConfig",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetComment(val *string) {
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetFailover(val *string) {
	_jsii_.Set(
		j,
		"failover",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetGeoLocation(val interface{}) {
	if err := j.validateSetGeoLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geoLocation",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetGeoProximityLocation(val interface{}) {
	if err := j.validateSetGeoProximityLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geoProximityLocation",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetHealthCheckId(val *string) {
	_jsii_.Set(
		j,
		"healthCheckId",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetHostedZoneId(val *string) {
	_jsii_.Set(
		j,
		"hostedZoneId",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetHostedZoneName(val *string) {
	_jsii_.Set(
		j,
		"hostedZoneName",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetMultiValueAnswer(val interface{}) {
	if err := j.validateSetMultiValueAnswerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiValueAnswer",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetResourceRecords(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceRecords",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetSetIdentifier(val *string) {
	_jsii_.Set(
		j,
		"setIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetTtl(val *string) {
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetWeight(val *float64) {
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

	if err := validateCfnRecordSet_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnRecordSet_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRecordSet_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		"isCfnResource",
		[]interface{}{x},
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
func CfnRecordSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRecordSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_CfnRecordSet) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnRecordSet) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRecordSet) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRecordSet) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnRecordSet) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnRecordSet) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnRecordSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnRecordSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnRecordSet) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRecordSet) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRecordSet) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnRecordSet) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRecordSet) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRecordSet) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRecordSet) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRecordSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRecordSet) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

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
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

