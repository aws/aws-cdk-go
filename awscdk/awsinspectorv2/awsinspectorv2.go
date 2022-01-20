package awsinspectorv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsinspectorv2/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::InspectorV2::Filter`.
//
// TODO: EXAMPLE
//
type CfnFilter interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	FilterAction() *string
	SetFilterAction(val *string)
	FilterCriteria() interface{}
	SetFilterCriteria(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnFilter
type jsiiProxy_CfnFilter struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFilter) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) FilterAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) FilterCriteria() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::InspectorV2::Filter`.
func NewCfnFilter(scope awscdk.Construct, id *string, props *CfnFilterProps) CfnFilter {
	_init_.Initialize()

	j := jsiiProxy_CfnFilter{}

	_jsii_.Create(
		"monocdk.aws_inspectorv2.CfnFilter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::InspectorV2::Filter`.
func NewCfnFilter_Override(c CfnFilter, scope awscdk.Construct, id *string, props *CfnFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_inspectorv2.CfnFilter",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFilter) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFilter) SetFilterAction(val *string) {
	_jsii_.Set(
		j,
		"filterAction",
		val,
	)
}

func (j *jsiiProxy_CfnFilter) SetFilterCriteria(val interface{}) {
	_jsii_.Set(
		j,
		"filterCriteria",
		val,
	)
}

func (j *jsiiProxy_CfnFilter) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnFilter_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_inspectorv2.CfnFilter",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnFilter_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_inspectorv2.CfnFilter",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnFilter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_inspectorv2.CfnFilter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFilter_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_inspectorv2.CfnFilter",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnFilter) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnFilter) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnFilter) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnFilter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnFilter) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnFilter) AddPropertyOverride(propertyPath *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnFilter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnFilter) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnFilter) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnFilter) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnFilter) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnFilter) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnFilter) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnFilter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnFilter) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFilter) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnFilter) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnFilter) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnFilter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnFilter) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnFilter) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnFilter_DateFilterProperty struct {
	// `CfnFilter.DateFilterProperty.EndInclusive`.
	EndInclusive *float64 `json:"endInclusive" yaml:"endInclusive"`
	// `CfnFilter.DateFilterProperty.StartInclusive`.
	StartInclusive *float64 `json:"startInclusive" yaml:"startInclusive"`
}

// TODO: EXAMPLE
//
type CfnFilter_FilterCriteriaProperty struct {
	// `CfnFilter.FilterCriteriaProperty.AwsAccountId`.
	AwsAccountId interface{} `json:"awsAccountId" yaml:"awsAccountId"`
	// `CfnFilter.FilterCriteriaProperty.ComponentId`.
	ComponentId interface{} `json:"componentId" yaml:"componentId"`
	// `CfnFilter.FilterCriteriaProperty.ComponentType`.
	ComponentType interface{} `json:"componentType" yaml:"componentType"`
	// `CfnFilter.FilterCriteriaProperty.Ec2InstanceImageId`.
	Ec2InstanceImageId interface{} `json:"ec2InstanceImageId" yaml:"ec2InstanceImageId"`
	// `CfnFilter.FilterCriteriaProperty.Ec2InstanceSubnetId`.
	Ec2InstanceSubnetId interface{} `json:"ec2InstanceSubnetId" yaml:"ec2InstanceSubnetId"`
	// `CfnFilter.FilterCriteriaProperty.Ec2InstanceVpcId`.
	Ec2InstanceVpcId interface{} `json:"ec2InstanceVpcId" yaml:"ec2InstanceVpcId"`
	// `CfnFilter.FilterCriteriaProperty.EcrImageArchitecture`.
	EcrImageArchitecture interface{} `json:"ecrImageArchitecture" yaml:"ecrImageArchitecture"`
	// `CfnFilter.FilterCriteriaProperty.EcrImageHash`.
	EcrImageHash interface{} `json:"ecrImageHash" yaml:"ecrImageHash"`
	// `CfnFilter.FilterCriteriaProperty.EcrImagePushedAt`.
	EcrImagePushedAt interface{} `json:"ecrImagePushedAt" yaml:"ecrImagePushedAt"`
	// `CfnFilter.FilterCriteriaProperty.EcrImageRegistry`.
	EcrImageRegistry interface{} `json:"ecrImageRegistry" yaml:"ecrImageRegistry"`
	// `CfnFilter.FilterCriteriaProperty.EcrImageRepositoryName`.
	EcrImageRepositoryName interface{} `json:"ecrImageRepositoryName" yaml:"ecrImageRepositoryName"`
	// `CfnFilter.FilterCriteriaProperty.EcrImageTags`.
	EcrImageTags interface{} `json:"ecrImageTags" yaml:"ecrImageTags"`
	// `CfnFilter.FilterCriteriaProperty.FindingArn`.
	FindingArn interface{} `json:"findingArn" yaml:"findingArn"`
	// `CfnFilter.FilterCriteriaProperty.FindingStatus`.
	FindingStatus interface{} `json:"findingStatus" yaml:"findingStatus"`
	// `CfnFilter.FilterCriteriaProperty.FindingType`.
	FindingType interface{} `json:"findingType" yaml:"findingType"`
	// `CfnFilter.FilterCriteriaProperty.FirstObservedAt`.
	FirstObservedAt interface{} `json:"firstObservedAt" yaml:"firstObservedAt"`
	// `CfnFilter.FilterCriteriaProperty.InspectorScore`.
	InspectorScore interface{} `json:"inspectorScore" yaml:"inspectorScore"`
	// `CfnFilter.FilterCriteriaProperty.LastObservedAt`.
	LastObservedAt interface{} `json:"lastObservedAt" yaml:"lastObservedAt"`
	// `CfnFilter.FilterCriteriaProperty.NetworkProtocol`.
	NetworkProtocol interface{} `json:"networkProtocol" yaml:"networkProtocol"`
	// `CfnFilter.FilterCriteriaProperty.PortRange`.
	PortRange interface{} `json:"portRange" yaml:"portRange"`
	// `CfnFilter.FilterCriteriaProperty.RelatedVulnerabilities`.
	RelatedVulnerabilities interface{} `json:"relatedVulnerabilities" yaml:"relatedVulnerabilities"`
	// `CfnFilter.FilterCriteriaProperty.ResourceId`.
	ResourceId interface{} `json:"resourceId" yaml:"resourceId"`
	// `CfnFilter.FilterCriteriaProperty.ResourceTags`.
	ResourceTags interface{} `json:"resourceTags" yaml:"resourceTags"`
	// `CfnFilter.FilterCriteriaProperty.ResourceType`.
	ResourceType interface{} `json:"resourceType" yaml:"resourceType"`
	// `CfnFilter.FilterCriteriaProperty.Severity`.
	Severity interface{} `json:"severity" yaml:"severity"`
	// `CfnFilter.FilterCriteriaProperty.Title`.
	Title interface{} `json:"title" yaml:"title"`
	// `CfnFilter.FilterCriteriaProperty.UpdatedAt`.
	UpdatedAt interface{} `json:"updatedAt" yaml:"updatedAt"`
	// `CfnFilter.FilterCriteriaProperty.VendorSeverity`.
	VendorSeverity interface{} `json:"vendorSeverity" yaml:"vendorSeverity"`
	// `CfnFilter.FilterCriteriaProperty.VulnerabilityId`.
	VulnerabilityId interface{} `json:"vulnerabilityId" yaml:"vulnerabilityId"`
	// `CfnFilter.FilterCriteriaProperty.VulnerabilitySource`.
	VulnerabilitySource interface{} `json:"vulnerabilitySource" yaml:"vulnerabilitySource"`
	// `CfnFilter.FilterCriteriaProperty.VulnerablePackages`.
	VulnerablePackages interface{} `json:"vulnerablePackages" yaml:"vulnerablePackages"`
}

// TODO: EXAMPLE
//
type CfnFilter_MapFilterProperty struct {
	// `CfnFilter.MapFilterProperty.Comparison`.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// `CfnFilter.MapFilterProperty.Key`.
	Key *string `json:"key" yaml:"key"`
	// `CfnFilter.MapFilterProperty.Value`.
	Value *string `json:"value" yaml:"value"`
}

// TODO: EXAMPLE
//
type CfnFilter_NumberFilterProperty struct {
	// `CfnFilter.NumberFilterProperty.LowerInclusive`.
	LowerInclusive *float64 `json:"lowerInclusive" yaml:"lowerInclusive"`
	// `CfnFilter.NumberFilterProperty.UpperInclusive`.
	UpperInclusive *float64 `json:"upperInclusive" yaml:"upperInclusive"`
}

// TODO: EXAMPLE
//
type CfnFilter_PackageFilterProperty struct {
	// `CfnFilter.PackageFilterProperty.Architecture`.
	Architecture interface{} `json:"architecture" yaml:"architecture"`
	// `CfnFilter.PackageFilterProperty.Epoch`.
	Epoch interface{} `json:"epoch" yaml:"epoch"`
	// `CfnFilter.PackageFilterProperty.Name`.
	Name interface{} `json:"name" yaml:"name"`
	// `CfnFilter.PackageFilterProperty.Release`.
	Release interface{} `json:"release" yaml:"release"`
	// `CfnFilter.PackageFilterProperty.SourceLayerHash`.
	SourceLayerHash interface{} `json:"sourceLayerHash" yaml:"sourceLayerHash"`
	// `CfnFilter.PackageFilterProperty.Version`.
	Version interface{} `json:"version" yaml:"version"`
}

// TODO: EXAMPLE
//
type CfnFilter_PortRangeFilterProperty struct {
	// `CfnFilter.PortRangeFilterProperty.BeginInclusive`.
	BeginInclusive *float64 `json:"beginInclusive" yaml:"beginInclusive"`
	// `CfnFilter.PortRangeFilterProperty.EndInclusive`.
	EndInclusive *float64 `json:"endInclusive" yaml:"endInclusive"`
}

// TODO: EXAMPLE
//
type CfnFilter_StringFilterProperty struct {
	// `CfnFilter.StringFilterProperty.Comparison`.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// `CfnFilter.StringFilterProperty.Value`.
	Value *string `json:"value" yaml:"value"`
}

// Properties for defining a `CfnFilter`.
//
// TODO: EXAMPLE
//
type CfnFilterProps struct {
	// `AWS::InspectorV2::Filter.FilterAction`.
	FilterAction *string `json:"filterAction" yaml:"filterAction"`
	// `AWS::InspectorV2::Filter.FilterCriteria`.
	FilterCriteria interface{} `json:"filterCriteria" yaml:"filterCriteria"`
	// `AWS::InspectorV2::Filter.Name`.
	Name *string `json:"name" yaml:"name"`
	// `AWS::InspectorV2::Filter.Description`.
	Description *string `json:"description" yaml:"description"`
}

