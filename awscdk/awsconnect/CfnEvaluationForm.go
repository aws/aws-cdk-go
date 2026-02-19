package awsconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an evaluation form for the specified Amazon Connect instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var evaluationFormSectionProperty_ EvaluationFormSectionProperty
//
//   cfnEvaluationForm := awscdk.Aws_connect.NewCfnEvaluationForm(this, jsii.String("MyCfnEvaluationForm"), &CfnEvaluationFormProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	Items: []interface{}{
//   		&EvaluationFormBaseItemProperty{
//   			Section: &EvaluationFormSectionProperty{
//   				RefId: jsii.String("refId"),
//   				Title: jsii.String("title"),
//
//   				// the properties below are optional
//   				Instructions: jsii.String("instructions"),
//   				Items: []interface{}{
//   					&EvaluationFormItemProperty{
//   						Question: &EvaluationFormQuestionProperty{
//   							QuestionType: jsii.String("questionType"),
//   							RefId: jsii.String("refId"),
//   							Title: jsii.String("title"),
//
//   							// the properties below are optional
//   							Enablement: &EvaluationFormItemEnablementConfigurationProperty{
//   								Action: jsii.String("action"),
//   								Condition: &EvaluationFormItemEnablementConditionProperty{
//   									Operands: []interface{}{
//   										&EvaluationFormItemEnablementConditionOperandProperty{
//   											Expression: &EvaluationFormItemEnablementExpressionProperty{
//   												Comparator: jsii.String("comparator"),
//   												Source: &EvaluationFormItemEnablementSourceProperty{
//   													Type: jsii.String("type"),
//
//   													// the properties below are optional
//   													RefId: jsii.String("refId"),
//   												},
//   												Values: []interface{}{
//   													&EvaluationFormItemEnablementSourceValueProperty{
//   														RefId: jsii.String("refId"),
//   														Type: jsii.String("type"),
//   													},
//   												},
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									Operator: jsii.String("operator"),
//   								},
//
//   								// the properties below are optional
//   								DefaultAction: jsii.String("defaultAction"),
//   							},
//   							Instructions: jsii.String("instructions"),
//   							NotApplicableEnabled: jsii.Boolean(false),
//   							QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   								MultiSelect: &EvaluationFormMultiSelectQuestionPropertiesProperty{
//   									Options: []interface{}{
//   										&EvaluationFormMultiSelectQuestionOptionProperty{
//   											RefId: jsii.String("refId"),
//   											Text: jsii.String("text"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									Automation: &EvaluationFormMultiSelectQuestionAutomationProperty{
//   										Options: []interface{}{
//   											&EvaluationFormMultiSelectQuestionAutomationOptionProperty{
//   												RuleCategory: &MultiSelectQuestionRuleCategoryAutomationProperty{
//   													Category: jsii.String("category"),
//   													Condition: jsii.String("condition"),
//   													OptionRefIds: []*string{
//   														jsii.String("optionRefIds"),
//   													},
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   											SourceType: jsii.String("sourceType"),
//   										},
//   										DefaultOptionRefIds: []*string{
//   											jsii.String("defaultOptionRefIds"),
//   										},
//   									},
//   									DisplayAs: jsii.String("displayAs"),
//   								},
//   								Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   									MaxValue: jsii.Number(123),
//   									MinValue: jsii.Number(123),
//
//   									// the properties below are optional
//   									Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   										AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   											SourceType: jsii.String("sourceType"),
//   										},
//   										PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   											Label: jsii.String("label"),
//   										},
//   									},
//   									Options: []interface{}{
//   										&EvaluationFormNumericQuestionOptionProperty{
//   											MaxValue: jsii.Number(123),
//   											MinValue: jsii.Number(123),
//
//   											// the properties below are optional
//   											AutomaticFail: jsii.Boolean(false),
//   											AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   												TargetSection: jsii.String("targetSection"),
//   											},
//   											Score: jsii.Number(123),
//   										},
//   									},
//   								},
//   								SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   									Options: []interface{}{
//   										&EvaluationFormSingleSelectQuestionOptionProperty{
//   											RefId: jsii.String("refId"),
//   											Text: jsii.String("text"),
//
//   											// the properties below are optional
//   											AutomaticFail: jsii.Boolean(false),
//   											AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   												TargetSection: jsii.String("targetSection"),
//   											},
//   											Score: jsii.Number(123),
//   										},
//   									},
//
//   									// the properties below are optional
//   									Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   										Options: []interface{}{
//   											&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   												RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   													Category: jsii.String("category"),
//   													Condition: jsii.String("condition"),
//   													OptionRefId: jsii.String("optionRefId"),
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   											SourceType: jsii.String("sourceType"),
//   										},
//   										DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   									},
//   									DisplayAs: jsii.String("displayAs"),
//   								},
//   								Text: &EvaluationFormTextQuestionPropertiesProperty{
//   									Automation: &EvaluationFormTextQuestionAutomationProperty{
//   										AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   											SourceType: jsii.String("sourceType"),
//   										},
//   									},
//   								},
//   							},
//   							Weight: jsii.Number(123),
//   						},
//   						Section: evaluationFormSectionProperty_,
//   					},
//   				},
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Status: jsii.String("status"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	AutoEvaluationConfiguration: &AutoEvaluationConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	Description: jsii.String("description"),
//   	LanguageConfiguration: &EvaluationFormLanguageConfigurationProperty{
//   		FormLanguage: jsii.String("formLanguage"),
//   	},
//   	ReviewConfiguration: &EvaluationReviewConfigurationProperty{
//   		ReviewNotificationRecipients: []interface{}{
//   			&EvaluationReviewNotificationRecipientProperty{
//   				Type: jsii.String("type"),
//   				Value: &EvaluationReviewNotificationRecipientValueProperty{
//   					UserId: jsii.String("userId"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		EligibilityDays: jsii.Number(123),
//   	},
//   	ScoringStrategy: &ScoringStrategyProperty{
//   		Mode: jsii.String("mode"),
//   		Status: jsii.String("status"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetConfiguration: &EvaluationFormTargetConfigurationProperty{
//   		ContactInteractionType: jsii.String("contactInteractionType"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html
//
type CfnEvaluationForm interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsconnect.IEvaluationFormRef
	awscdk.ITaggable
	// The Amazon Resource Name (ARN) of the evaluation form.
	AttrEvaluationFormArn() *string
	// The automatic evaluation configuration of an evaluation form.
	AutoEvaluationConfiguration() interface{}
	SetAutoEvaluationConfiguration(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the evaluation form.
	Description() *string
	SetDescription(val *string)
	Env() *interfaces.ResourceEnvironment
	// A reference to a EvaluationForm resource.
	EvaluationFormRef() *interfacesawsconnect.EvaluationFormReference
	// The identifier of the Amazon Connect instance.
	InstanceArn() *string
	SetInstanceArn(val *string)
	// Items that are part of the evaluation form.
	Items() interface{}
	SetItems(val interface{})
	// Configuration for language settings of this evaluation form.
	LanguageConfiguration() interface{}
	SetLanguageConfiguration(val interface{})
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
	ReviewConfiguration() interface{}
	SetReviewConfiguration(val interface{})
	// A scoring strategy of the evaluation form.
	ScoringStrategy() interface{}
	SetScoringStrategy(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The status of the evaluation form.
	Status() *string
	SetStatus(val *string)
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The tags used to organize, track, or control access for this resource.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// Configuration that specifies the target for this evaluation form.
	TargetConfiguration() interface{}
	SetTargetConfiguration(val interface{})
	// A title of the evaluation form.
	Title() *string
	SetTitle(val *string)
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnEvaluationForm
type jsiiProxy_CfnEvaluationForm struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsconnectIEvaluationFormRef
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnEvaluationForm) AttrEvaluationFormArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEvaluationFormArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) AutoEvaluationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoEvaluationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) EvaluationFormRef() *interfacesawsconnect.EvaluationFormReference {
	var returns *interfacesawsconnect.EvaluationFormReference
	_jsii_.Get(
		j,
		"evaluationFormRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) InstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) Items() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) LanguageConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"languageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) ReviewConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reviewConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) ScoringStrategy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scoringStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) TargetConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluationForm) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Connect::EvaluationForm`.
func NewCfnEvaluationForm(scope constructs.Construct, id *string, props *CfnEvaluationFormProps) CfnEvaluationForm {
	_init_.Initialize()

	if err := validateNewCfnEvaluationFormParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEvaluationForm{}

	_jsii_.Create(
		"aws-cdk-lib.aws_connect.CfnEvaluationForm",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Connect::EvaluationForm`.
func NewCfnEvaluationForm_Override(c CfnEvaluationForm, scope constructs.Construct, id *string, props *CfnEvaluationFormProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_connect.CfnEvaluationForm",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEvaluationForm)SetAutoEvaluationConfiguration(val interface{}) {
	if err := j.validateSetAutoEvaluationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoEvaluationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnEvaluationForm)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnEvaluationForm)SetInstanceArn(val *string) {
	if err := j.validateSetInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceArn",
		val,
	)
}

func (j *jsiiProxy_CfnEvaluationForm)SetItems(val interface{}) {
	if err := j.validateSetItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_CfnEvaluationForm)SetLanguageConfiguration(val interface{}) {
	if err := j.validateSetLanguageConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnEvaluationForm)SetReviewConfiguration(val interface{}) {
	if err := j.validateSetReviewConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reviewConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnEvaluationForm)SetScoringStrategy(val interface{}) {
	if err := j.validateSetScoringStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scoringStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnEvaluationForm)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_CfnEvaluationForm)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnEvaluationForm)SetTargetConfiguration(val interface{}) {
	if err := j.validateSetTargetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnEvaluationForm)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func CfnEvaluationForm_ArnForEvaluationForm(resource interfacesawsconnect.IEvaluationFormRef) *string {
	_init_.Initialize()

	if err := validateCfnEvaluationForm_ArnForEvaluationFormParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_connect.CfnEvaluationForm",
		"arnForEvaluationForm",
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
func CfnEvaluationForm_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEvaluationForm_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_connect.CfnEvaluationForm",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnEvaluationForm.
func CfnEvaluationForm_IsCfnEvaluationForm(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEvaluationForm_IsCfnEvaluationFormParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_connect.CfnEvaluationForm",
		"isCfnEvaluationForm",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnEvaluationForm_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEvaluationForm_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_connect.CfnEvaluationForm",
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
func CfnEvaluationForm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEvaluationForm_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_connect.CfnEvaluationForm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEvaluationForm_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_connect.CfnEvaluationForm",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEvaluationForm) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEvaluationForm) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEvaluationForm) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEvaluationForm) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEvaluationForm) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEvaluationForm) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEvaluationForm) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEvaluationForm) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEvaluationForm) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnEvaluationForm) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnEvaluationForm) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEvaluationForm) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEvaluationForm) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEvaluationForm) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEvaluationForm) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEvaluationForm) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnEvaluationForm) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnEvaluationForm) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEvaluationForm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEvaluationForm) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnEvaluationForm) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

