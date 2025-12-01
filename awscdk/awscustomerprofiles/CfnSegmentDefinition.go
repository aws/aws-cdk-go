package awscustomerprofiles

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscustomerprofiles"
	"github.com/aws/constructs-go/constructs/v10"
)

// A segment definition resource of Amazon Connect Customer Profiles.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSegmentDefinition := awscdk.Aws_customerprofiles.NewCfnSegmentDefinition(this, jsii.String("MyCfnSegmentDefinition"), &CfnSegmentDefinitionProps{
//   	DisplayName: jsii.String("displayName"),
//   	DomainName: jsii.String("domainName"),
//   	SegmentDefinitionName: jsii.String("segmentDefinitionName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	SegmentGroups: &SegmentGroupProperty{
//   		Groups: []interface{}{
//   			&GroupProperty{
//   				Dimensions: []interface{}{
//   					&DimensionProperty{
//   						CalculatedAttributes: map[string]interface{}{
//   							"calculatedAttributesKey": &CalculatedAttributeDimensionProperty{
//   								"dimensionType": jsii.String("dimensionType"),
//   								"values": []*string{
//   									jsii.String("values"),
//   								},
//
//   								// the properties below are optional
//   								"conditionOverrides": &ConditionOverridesProperty{
//   									"range": &RangeOverrideProperty{
//   										"start": jsii.Number(123),
//   										"unit": jsii.String("unit"),
//
//   										// the properties below are optional
//   										"end": jsii.Number(123),
//   									},
//   								},
//   							},
//   						},
//   						ProfileAttributes: &ProfileAttributesProperty{
//   							AccountNumber: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							AdditionalInformation: &ExtraLengthValueProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							Address: &AddressDimensionProperty{
//   								City: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Country: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								County: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								PostalCode: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Province: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								State: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   							Attributes: map[string]interface{}{
//   								"attributesKey": &AttributeDimensionProperty{
//   									"dimensionType": jsii.String("dimensionType"),
//   									"values": []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   							BillingAddress: &AddressDimensionProperty{
//   								City: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Country: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								County: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								PostalCode: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Province: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								State: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   							BirthDate: &DateDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							BusinessEmailAddress: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							BusinessName: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							BusinessPhoneNumber: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							EmailAddress: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							FirstName: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							GenderString: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							HomePhoneNumber: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							LastName: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							MailingAddress: &AddressDimensionProperty{
//   								City: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Country: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								County: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								PostalCode: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Province: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								State: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   							MiddleName: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							MobilePhoneNumber: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							PartyTypeString: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							PersonalEmailAddress: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							PhoneNumber: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							ProfileType: &ProfileTypeDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							ShippingAddress: &AddressDimensionProperty{
//   								City: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Country: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								County: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								PostalCode: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Province: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								State: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   				SourceSegments: []interface{}{
//   					&SourceSegmentProperty{
//   						SegmentDefinitionName: jsii.String("segmentDefinitionName"),
//   					},
//   				},
//   				SourceType: jsii.String("sourceType"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Include: jsii.String("include"),
//   	},
//   	SegmentSqlQuery: jsii.String("segmentSqlQuery"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-segmentdefinition.html
//
type CfnSegmentDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawscustomerprofiles.ISegmentDefinitionRef
	awscdk.ITaggableV2
	// When the segment definition was created.
	AttrCreatedAt() *string
	// The arn of the segment definition.
	AttrSegmentDefinitionArn() *string
	// The SQL query that defines the segment criteria.
	AttrSegmentType() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the segment definition.
	Description() *string
	SetDescription(val *string)
	// Display name of the segment definition.
	DisplayName() *string
	SetDisplayName(val *string)
	// The name of the domain.
	DomainName() *string
	SetDomainName(val *string)
	Env() *interfaces.ResourceEnvironment
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
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Name of the segment definition.
	SegmentDefinitionName() *string
	SetSegmentDefinitionName(val *string)
	// A reference to a SegmentDefinition resource.
	SegmentDefinitionRef() *interfacesawscustomerprofiles.SegmentDefinitionReference
	// Contains all groups of the segment definition.
	SegmentGroups() interface{}
	SetSegmentGroups(val interface{})
	// The SQL query that defines the segment criteria.
	SegmentSqlQuery() *string
	SetSegmentSqlQuery(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags belonging to the segment definition.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
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

// The jsii proxy struct for CfnSegmentDefinition
type jsiiProxy_CfnSegmentDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawscustomerprofilesISegmentDefinitionRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnSegmentDefinition) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) AttrSegmentDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSegmentDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) AttrSegmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSegmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) SegmentDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) SegmentDefinitionRef() *interfacesawscustomerprofiles.SegmentDefinitionReference {
	var returns *interfacesawscustomerprofiles.SegmentDefinitionReference
	_jsii_.Get(
		j,
		"segmentDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) SegmentGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"segmentGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) SegmentSqlQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentSqlQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinition) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::CustomerProfiles::SegmentDefinition`.
func NewCfnSegmentDefinition(scope constructs.Construct, id *string, props *CfnSegmentDefinitionProps) CfnSegmentDefinition {
	_init_.Initialize()

	if err := validateNewCfnSegmentDefinitionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSegmentDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_customerprofiles.CfnSegmentDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CustomerProfiles::SegmentDefinition`.
func NewCfnSegmentDefinition_Override(c CfnSegmentDefinition, scope constructs.Construct, id *string, props *CfnSegmentDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_customerprofiles.CfnSegmentDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSegmentDefinition)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnSegmentDefinition)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_CfnSegmentDefinition)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnSegmentDefinition)SetSegmentDefinitionName(val *string) {
	if err := j.validateSetSegmentDefinitionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentDefinitionName",
		val,
	)
}

func (j *jsiiProxy_CfnSegmentDefinition)SetSegmentGroups(val interface{}) {
	if err := j.validateSetSegmentGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentGroups",
		val,
	)
}

func (j *jsiiProxy_CfnSegmentDefinition)SetSegmentSqlQuery(val *string) {
	_jsii_.Set(
		j,
		"segmentSqlQuery",
		val,
	)
}

func (j *jsiiProxy_CfnSegmentDefinition)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func CfnSegmentDefinition_ArnForSegmentDefinition(resource interfacesawscustomerprofiles.ISegmentDefinitionRef) *string {
	_init_.Initialize()

	if err := validateCfnSegmentDefinition_ArnForSegmentDefinitionParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnSegmentDefinition",
		"arnForSegmentDefinition",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSegmentDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSegmentDefinition_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnSegmentDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnSegmentDefinition_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSegmentDefinition_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnSegmentDefinition",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnSegmentDefinition.
func CfnSegmentDefinition_IsCfnSegmentDefinition(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSegmentDefinition_IsCfnSegmentDefinitionParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnSegmentDefinition",
		"isCfnSegmentDefinition",
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
func CfnSegmentDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSegmentDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnSegmentDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSegmentDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_customerprofiles.CfnSegmentDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSegmentDefinition) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSegmentDefinition) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSegmentDefinition) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSegmentDefinition) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSegmentDefinition) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSegmentDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSegmentDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSegmentDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSegmentDefinition) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnSegmentDefinition) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnSegmentDefinition) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSegmentDefinition) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSegmentDefinition) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSegmentDefinition) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSegmentDefinition) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSegmentDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnSegmentDefinition) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnSegmentDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSegmentDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSegmentDefinition) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

