package awslookoutmetrics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslookoutmetrics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::LookoutMetrics::AnomalyDetector`.
//
// The `AWS::LookoutMetrics::AnomalyDetector` type creates an anomaly detector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnomalyDetector := awscdk.Aws_lookoutmetrics.NewCfnAnomalyDetector(this, jsii.String("MyCfnAnomalyDetector"), &cfnAnomalyDetectorProps{
//   	anomalyDetectorConfig: &anomalyDetectorConfigProperty{
//   		anomalyDetectorFrequency: jsii.String("anomalyDetectorFrequency"),
//   	},
//   	metricSetList: []interface{}{
//   		&metricSetProperty{
//   			metricList: []interface{}{
//   				&metricProperty{
//   					aggregationFunction: jsii.String("aggregationFunction"),
//   					metricName: jsii.String("metricName"),
//
//   					// the properties below are optional
//   					namespace: jsii.String("namespace"),
//   				},
//   			},
//   			metricSetName: jsii.String("metricSetName"),
//   			metricSource: &metricSourceProperty{
//   				appFlowConfig: &appFlowConfigProperty{
//   					flowName: jsii.String("flowName"),
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				cloudwatchConfig: &cloudwatchConfigProperty{
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				rdsSourceConfig: &rDSSourceConfigProperty{
//   					databaseHost: jsii.String("databaseHost"),
//   					databaseName: jsii.String("databaseName"),
//   					databasePort: jsii.Number(123),
//   					dbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   					roleArn: jsii.String("roleArn"),
//   					secretManagerArn: jsii.String("secretManagerArn"),
//   					tableName: jsii.String("tableName"),
//   					vpcConfiguration: &vpcConfigurationProperty{
//   						securityGroupIdList: []*string{
//   							jsii.String("securityGroupIdList"),
//   						},
//   						subnetIdList: []*string{
//   							jsii.String("subnetIdList"),
//   						},
//   					},
//   				},
//   				redshiftSourceConfig: &redshiftSourceConfigProperty{
//   					clusterIdentifier: jsii.String("clusterIdentifier"),
//   					databaseHost: jsii.String("databaseHost"),
//   					databaseName: jsii.String("databaseName"),
//   					databasePort: jsii.Number(123),
//   					roleArn: jsii.String("roleArn"),
//   					secretManagerArn: jsii.String("secretManagerArn"),
//   					tableName: jsii.String("tableName"),
//   					vpcConfiguration: &vpcConfigurationProperty{
//   						securityGroupIdList: []*string{
//   							jsii.String("securityGroupIdList"),
//   						},
//   						subnetIdList: []*string{
//   							jsii.String("subnetIdList"),
//   						},
//   					},
//   				},
//   				s3SourceConfig: &s3SourceConfigProperty{
//   					fileFormatDescriptor: &fileFormatDescriptorProperty{
//   						csvFormatDescriptor: &csvFormatDescriptorProperty{
//   							charset: jsii.String("charset"),
//   							containsHeader: jsii.Boolean(false),
//   							delimiter: jsii.String("delimiter"),
//   							fileCompression: jsii.String("fileCompression"),
//   							headerList: []*string{
//   								jsii.String("headerList"),
//   							},
//   							quoteSymbol: jsii.String("quoteSymbol"),
//   						},
//   						jsonFormatDescriptor: &jsonFormatDescriptorProperty{
//   							charset: jsii.String("charset"),
//   							fileCompression: jsii.String("fileCompression"),
//   						},
//   					},
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					historicalDataPathList: []*string{
//   						jsii.String("historicalDataPathList"),
//   					},
//   					templatedPathList: []*string{
//   						jsii.String("templatedPathList"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			dimensionList: []*string{
//   				jsii.String("dimensionList"),
//   			},
//   			metricSetDescription: jsii.String("metricSetDescription"),
//   			metricSetFrequency: jsii.String("metricSetFrequency"),
//   			offset: jsii.Number(123),
//   			timestampColumn: &timestampColumnProperty{
//   				columnFormat: jsii.String("columnFormat"),
//   				columnName: jsii.String("columnName"),
//   			},
//   			timezone: jsii.String("timezone"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	anomalyDetectorDescription: jsii.String("anomalyDetectorDescription"),
//   	anomalyDetectorName: jsii.String("anomalyDetectorName"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   })
//
type CfnAnomalyDetector interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Contains information about the configuration of the anomaly detector.
	AnomalyDetectorConfig() interface{}
	SetAnomalyDetectorConfig(val interface{})
	// A description of the detector.
	AnomalyDetectorDescription() *string
	SetAnomalyDetectorDescription(val *string)
	// The name of the detector.
	AnomalyDetectorName() *string
	SetAnomalyDetectorName(val *string)
	// The Amazon Resource Name (ARN) of the detector.
	//
	// For example, `arn:aws:lookoutmetrics:us-east-2:123456789012:AnomalyDetector:my-detector`.
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The ARN of the KMS key to use to encrypt your data.
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
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
	// The detector's dataset.
	MetricSetList() interface{}
	SetMetricSetList(val interface{})
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnAnomalyDetector
type jsiiProxy_CfnAnomalyDetector struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAnomalyDetector) AnomalyDetectorConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anomalyDetectorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) AnomalyDetectorDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anomalyDetectorDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) AnomalyDetectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anomalyDetectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) MetricSetList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricSetList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::LookoutMetrics::AnomalyDetector`.
func NewCfnAnomalyDetector(scope constructs.Construct, id *string, props *CfnAnomalyDetectorProps) CfnAnomalyDetector {
	_init_.Initialize()

	j := jsiiProxy_CfnAnomalyDetector{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAnomalyDetector",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LookoutMetrics::AnomalyDetector`.
func NewCfnAnomalyDetector_Override(c CfnAnomalyDetector, scope constructs.Construct, id *string, props *CfnAnomalyDetectorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAnomalyDetector",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector) SetAnomalyDetectorConfig(val interface{}) {
	_jsii_.Set(
		j,
		"anomalyDetectorConfig",
		val,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector) SetAnomalyDetectorDescription(val *string) {
	_jsii_.Set(
		j,
		"anomalyDetectorDescription",
		val,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector) SetAnomalyDetectorName(val *string) {
	_jsii_.Set(
		j,
		"anomalyDetectorName",
		val,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector) SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector) SetMetricSetList(val interface{}) {
	_jsii_.Set(
		j,
		"metricSetList",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAnomalyDetector_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAnomalyDetector",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAnomalyDetector_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAnomalyDetector",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnAnomalyDetector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAnomalyDetector",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAnomalyDetector_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAnomalyDetector",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAnomalyDetector) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAnomalyDetector) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAnomalyDetector) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAnomalyDetector) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAnomalyDetector) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAnomalyDetector) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

