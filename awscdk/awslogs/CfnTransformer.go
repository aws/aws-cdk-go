package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates or updates a *log transformer* for a single log group.
//
// You use log transformers to transform log events into a different format, making them easier for you to process and analyze. You can also transform logs from different sources into standardized formats that contains relevant, source-specific information.
//
// After you have created a transformer, CloudWatch Logs performs the transformations at the time of log ingestion. You can then refer to the transformed versions of the logs during operations such as querying with CloudWatch Logs Insights or creating metric filters or subscription filers.
//
// You can also use a transformer to copy metadata from metadata keys into the log events themselves. This metadata can include log group name, log stream name, account ID and Region.
//
// A transformer for a log group is a series of processors, where each processor applies one type of transformation to the log events ingested into this log group. The processors work one after another, in the order that you list them, like a pipeline. For more information about the available processors to use in a transformer, see [Processors that you can use](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-Processors) .
//
// Having log events in standardized format enables visibility across your applications for your log analysis, reporting, and alarming needs. CloudWatch Logs provides transformation for common log types with out-of-the-box transformation templates for major AWS log sources such as VPC flow logs, Lambda, and Amazon RDS. You can use pre-built transformation templates or create custom transformation policies.
//
// You can create transformers only for the log groups in the Standard log class.
//
// You can also set up a transformer at the account level. For more information, see [PutAccountPolicy](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutAccountPolicy.html) . If there is both a log-group level transformer created with `PutTransformer` and an account-level transformer that could apply to the same log group, the log group uses only the log-group level transformer. It ignores the account-level transformer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransformer := awscdk.Aws_logs.NewCfnTransformer(this, jsii.String("MyCfnTransformer"), &CfnTransformerProps{
//   	LogGroupIdentifier: jsii.String("logGroupIdentifier"),
//   	TransformerConfig: []interface{}{
//   		&ProcessorProperty{
//   			AddKeys: &AddKeysProperty{
//   				Entries: []interface{}{
//   					&AddKeyEntryProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//
//   						// the properties below are optional
//   						OverwriteIfExists: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			CopyValue: &CopyValueProperty{
//   				Entries: []interface{}{
//   					&CopyValueEntryProperty{
//   						Source: jsii.String("source"),
//   						Target: jsii.String("target"),
//
//   						// the properties below are optional
//   						OverwriteIfExists: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			Csv: &CsvProperty{
//   				Columns: []*string{
//   					jsii.String("columns"),
//   				},
//   				Delimiter: jsii.String("delimiter"),
//   				QuoteCharacter: jsii.String("quoteCharacter"),
//   				Source: jsii.String("source"),
//   			},
//   			DateTimeConverter: &DateTimeConverterProperty{
//   				MatchPatterns: []*string{
//   					jsii.String("matchPatterns"),
//   				},
//   				Source: jsii.String("source"),
//   				Target: jsii.String("target"),
//
//   				// the properties below are optional
//   				Locale: jsii.String("locale"),
//   				SourceTimezone: jsii.String("sourceTimezone"),
//   				TargetFormat: jsii.String("targetFormat"),
//   				TargetTimezone: jsii.String("targetTimezone"),
//   			},
//   			DeleteKeys: &DeleteKeysProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   			Grok: &GrokProperty{
//   				Match: jsii.String("match"),
//
//   				// the properties below are optional
//   				Source: jsii.String("source"),
//   			},
//   			ListToMap: &ListToMapProperty{
//   				Key: jsii.String("key"),
//   				Source: jsii.String("source"),
//
//   				// the properties below are optional
//   				Flatten: jsii.Boolean(false),
//   				FlattenedElement: jsii.String("flattenedElement"),
//   				Target: jsii.String("target"),
//   				ValueKey: jsii.String("valueKey"),
//   			},
//   			LowerCaseString: &LowerCaseStringProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   			MoveKeys: &MoveKeysProperty{
//   				Entries: []interface{}{
//   					&MoveKeyEntryProperty{
//   						Source: jsii.String("source"),
//   						Target: jsii.String("target"),
//
//   						// the properties below are optional
//   						OverwriteIfExists: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			ParseCloudfront: &ParseCloudfrontProperty{
//   				Source: jsii.String("source"),
//   			},
//   			ParseJson: &ParseJSONProperty{
//   				Destination: jsii.String("destination"),
//   				Source: jsii.String("source"),
//   			},
//   			ParseKeyValue: &ParseKeyValueProperty{
//   				Destination: jsii.String("destination"),
//   				FieldDelimiter: jsii.String("fieldDelimiter"),
//   				KeyPrefix: jsii.String("keyPrefix"),
//   				KeyValueDelimiter: jsii.String("keyValueDelimiter"),
//   				NonMatchValue: jsii.String("nonMatchValue"),
//   				OverwriteIfExists: jsii.Boolean(false),
//   				Source: jsii.String("source"),
//   			},
//   			ParsePostgres: &ParsePostgresProperty{
//   				Source: jsii.String("source"),
//   			},
//   			ParseRoute53: &ParseRoute53Property{
//   				Source: jsii.String("source"),
//   			},
//   			ParseToOcsf: &ParseToOCSFProperty{
//   				EventSource: jsii.String("eventSource"),
//   				OcsfVersion: jsii.String("ocsfVersion"),
//
//   				// the properties below are optional
//   				Source: jsii.String("source"),
//   			},
//   			ParseVpc: &ParseVPCProperty{
//   				Source: jsii.String("source"),
//   			},
//   			ParseWaf: &ParseWAFProperty{
//   				Source: jsii.String("source"),
//   			},
//   			RenameKeys: &RenameKeysProperty{
//   				Entries: []interface{}{
//   					&RenameKeyEntryProperty{
//   						Key: jsii.String("key"),
//   						RenameTo: jsii.String("renameTo"),
//
//   						// the properties below are optional
//   						OverwriteIfExists: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			SplitString: &SplitStringProperty{
//   				Entries: []interface{}{
//   					&SplitStringEntryProperty{
//   						Delimiter: jsii.String("delimiter"),
//   						Source: jsii.String("source"),
//   					},
//   				},
//   			},
//   			SubstituteString: &SubstituteStringProperty{
//   				Entries: []interface{}{
//   					&SubstituteStringEntryProperty{
//   						From: jsii.String("from"),
//   						Source: jsii.String("source"),
//   						To: jsii.String("to"),
//   					},
//   				},
//   			},
//   			TrimString: &TrimStringProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   			TypeConverter: &TypeConverterProperty{
//   				Entries: []interface{}{
//   					&TypeConverterEntryProperty{
//   						Key: jsii.String("key"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			UpperCaseString: &UpperCaseStringProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-transformer.html
//
type CfnTransformer interface {
	awscdk.CfnResource
	ITransformerRef
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specify either the name or ARN of the log group to create the transformer for.
	LogGroupIdentifier() *string
	SetLogGroupIdentifier(val *string)
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// This structure is an array that contains the configuration of this log transformer.
	TransformerConfig() interface{}
	SetTransformerConfig(val interface{})
	// A reference to a Transformer resource.
	TransformerRef() *TransformerReference
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

// The jsii proxy struct for CfnTransformer
type jsiiProxy_CfnTransformer struct {
	internal.Type__awscdkCfnResource
	jsiiProxy_ITransformerRef
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTransformer) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformer) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformer) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformer) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformer) LogGroupIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformer) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformer) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformer) TransformerConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformer) TransformerRef() *TransformerReference {
	var returns *TransformerReference
	_jsii_.Get(
		j,
		"transformerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformer) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformer) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnTransformer(scope constructs.Construct, id *string, props *CfnTransformerProps) CfnTransformer {
	_init_.Initialize()

	if err := validateNewCfnTransformerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransformer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.CfnTransformer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnTransformer_Override(c CfnTransformer, scope constructs.Construct, id *string, props *CfnTransformerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.CfnTransformer",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTransformer)SetLogGroupIdentifier(val *string) {
	if err := j.validateSetLogGroupIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnTransformer)SetTransformerConfig(val interface{}) {
	if err := j.validateSetTransformerConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformerConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTransformer_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransformer_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.CfnTransformer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnTransformer_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransformer_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.CfnTransformer",
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
func CfnTransformer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransformer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.CfnTransformer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTransformer_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CfnTransformer",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransformer) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTransformer) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTransformer) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTransformer) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTransformer) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTransformer) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTransformer) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTransformer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTransformer) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnTransformer) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnTransformer) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTransformer) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTransformer) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTransformer) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTransformer) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTransformer) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnTransformer) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnTransformer) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTransformer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTransformer) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

