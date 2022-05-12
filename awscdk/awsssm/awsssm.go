package awsssm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::SSM::Association`.
//
// The `AWS::SSM::Association` resource creates a State Manager association for your managed instances. A State Manager association defines the state that you want to maintain on your instances. For example, an association can specify that anti-virus software must be installed and running on your instances, or that certain ports must be closed. For static targets, the association specifies a schedule for when the configuration is reapplied. For dynamic targets, such as an AWS Resource Groups or an AWS Auto Scaling Group, State Manager applies the configuration when new instances are added to the group. The association also specifies actions to take when applying the configuration. For example, an association for anti-virus software might run once a day. If the software is not installed, then State Manager installs it. If the software is installed, but the service is not running, then the association might instruct State Manager to start the service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnAssociation := awscdk.Aws_ssm.NewCfnAssociation(this, jsii.String("MyCfnAssociation"), &cfnAssociationProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	applyOnlyAtCronInterval: jsii.Boolean(false),
//   	associationName: jsii.String("associationName"),
//   	automationTargetParameterName: jsii.String("automationTargetParameterName"),
//   	calendarNames: []*string{
//   		jsii.String("calendarNames"),
//   	},
//   	complianceSeverity: jsii.String("complianceSeverity"),
//   	documentVersion: jsii.String("documentVersion"),
//   	instanceId: jsii.String("instanceId"),
//   	maxConcurrency: jsii.String("maxConcurrency"),
//   	maxErrors: jsii.String("maxErrors"),
//   	outputLocation: &instanceAssociationOutputLocationProperty{
//   		s3Location: &s3OutputLocationProperty{
//   			outputS3BucketName: jsii.String("outputS3BucketName"),
//   			outputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   			outputS3Region: jsii.String("outputS3Region"),
//   		},
//   	},
//   	parameters: map[string]interface{}{
//   		"parametersKey": parameters,
//   	},
//   	scheduleExpression: jsii.String("scheduleExpression"),
//   	syncCompliance: jsii.String("syncCompliance"),
//   	targets: []interface{}{
//   		&targetProperty{
//   			key: jsii.String("key"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	waitForSuccessTimeoutSeconds: jsii.Number(123),
//   })
//
type CfnAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// By default, when you create a new association, the system runs it immediately after it is created and then according to the schedule you specified.
	//
	// Specify this option if you don't want an association to run immediately after you create it. This parameter is not supported for rate expressions.
	ApplyOnlyAtCronInterval() interface{}
	SetApplyOnlyAtCronInterval(val interface{})
	// Specify a descriptive name for the association.
	AssociationName() *string
	SetAssociationName(val *string)
	// The association ID.
	AttrAssociationId() *string
	// Choose the parameter that will define how your automation will branch out.
	//
	// This target is required for associations that use an Automation runbook and target resources by using rate controls. Automation is a capability of AWS Systems Manager .
	AutomationTargetParameterName() *string
	SetAutomationTargetParameterName(val *string)
	// The names or Amazon Resource Names (ARNs) of the Change Calendar type documents your associations are gated under.
	//
	// The associations only run when that Change Calendar is open. For more information, see [AWS Systems Manager Change Calendar](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-change-calendar) .
	CalendarNames() *[]*string
	SetCalendarNames(val *[]*string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The severity level that is assigned to the association.
	ComplianceSeverity() *string
	SetComplianceSeverity(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The version of the SSM document to associate with the target.
	//
	// > Note the following important information.
	// >
	// > - State Manager doesn't support running associations that use a new version of a document if that document is shared from another account. State Manager always runs the `default` version of a document if shared from another account, even though the Systems Manager console shows that a new version was processed. If you want to run an association using a new version of a document shared form another account, you must set the document version to `default` .
	// > - `DocumentVersion` is not valid for documents owned by AWS , such as `AWS-RunPatchBaseline` or `AWS-UpdateSSMAgent` . If you specify `DocumentVersion` for an AWS document, the system returns the following error: "Error occurred during operation 'CreateAssociation'." (RequestToken: <token>, HandlerErrorCode: GeneralServiceException).
	DocumentVersion() *string
	SetDocumentVersion(val *string)
	// The ID of the instance that the SSM document is associated with.
	//
	// You must specify the `InstanceId` or `Targets` property.
	//
	// > `InstanceId` has been deprecated. To specify an instance ID for an association, use the `Targets` parameter. If you use the parameter `InstanceId` , you cannot use the parameters `AssociationName` , `DocumentVersion` , `MaxErrors` , `MaxConcurrency` , `OutputLocation` , or `ScheduleExpression` . To use these parameters, you must use the `Targets` parameter.
	InstanceId() *string
	SetInstanceId(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The maximum number of targets allowed to run the association at the same time.
	//
	// You can specify a number, for example 10, or a percentage of the target set, for example 10%. The default value is 100%, which means all targets run the association at the same time.
	//
	// If a new managed node starts and attempts to run an association while Systems Manager is running `MaxConcurrency` associations, the association is allowed to run. During the next association interval, the new managed node will process its association within the limit specified for `MaxConcurrency` .
	MaxConcurrency() *string
	SetMaxConcurrency(val *string)
	// The number of errors that are allowed before the system stops sending requests to run the association on additional targets.
	//
	// You can specify either an absolute number of errors, for example 10, or a percentage of the target set, for example 10%. If you specify 3, for example, the system stops sending requests when the fourth error is received. If you specify 0, then the system stops sending requests after the first error is returned. If you run an association on 50 managed nodes and set `MaxError` to 10%, then the system stops sending the request when the sixth error is received.
	//
	// Executions that are already running an association when `MaxErrors` is reached are allowed to complete, but some of these executions may fail as well. If you need to ensure that there won't be more than max-errors failed executions, set `MaxConcurrency` to 1 so that executions proceed one at a time.
	MaxErrors() *string
	SetMaxErrors(val *string)
	// The name of the SSM document that contains the configuration information for the instance.
	//
	// You can specify `Command` or `Automation` documents. The documents can be AWS -predefined documents, documents you created, or a document that is shared with you from another account. For SSM documents that are shared with you from other AWS accounts , you must specify the complete SSM document ARN, in the following format:
	//
	// `arn:partition:ssm:region:account-id:document/document-name`
	//
	// For example: `arn:aws:ssm:us-east-2:12345678912:document/My-Shared-Document`
	//
	// For AWS -predefined documents and SSM documents you created in your account, you only need to specify the document name. For example, AWS -ApplyPatchBaseline or My-Document.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// An Amazon Simple Storage Service (Amazon S3) bucket where you want to store the output details of the request.
	OutputLocation() interface{}
	SetOutputLocation(val interface{})
	// The parameters for the runtime configuration of the document.
	Parameters() interface{}
	SetParameters(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// A cron expression that specifies a schedule when the association runs.
	//
	// The schedule runs in Coordinated Universal Time (UTC).
	ScheduleExpression() *string
	SetScheduleExpression(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The mode for generating association compliance.
	//
	// You can specify `AUTO` or `MANUAL` . In `AUTO` mode, the system uses the status of the association execution to determine the compliance status. If the association execution runs successfully, then the association is `COMPLIANT` . If the association execution doesn't run successfully, the association is `NON-COMPLIANT` .
	//
	// In `MANUAL` mode, you must specify the `AssociationId` as a parameter for the PutComplianceItems API action. In this case, compliance data is not managed by State Manager. It is managed by your direct call to the PutComplianceItems API action.
	//
	// By default, all associations use `AUTO` mode.
	SyncCompliance() *string
	SetSyncCompliance(val *string)
	// The targets for the association.
	//
	// You must specify the `InstanceId` or `Targets` property. You can target all instances in an AWS account by specifying the `InstanceIds` key with a value of `*` . To view a JSON and a YAML example that targets all instances, see "Create an association for all managed instances in an AWS account " on the Examples page.
	Targets() interface{}
	SetTargets(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The number of seconds the service should wait for the association status to show "Success" before proceeding with the stack execution.
	//
	// If the association status doesn't show "Success" after the specified number of seconds, then stack creation fails.
	WaitForSuccessTimeoutSeconds() *float64
	SetWaitForSuccessTimeoutSeconds(val *float64)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnAssociation
type jsiiProxy_CfnAssociation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAssociation) ApplyOnlyAtCronInterval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyOnlyAtCronInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) AssociationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) AttrAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) AutomationTargetParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationTargetParameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) CalendarNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"calendarNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) ComplianceSeverity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceSeverity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) DocumentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) MaxConcurrency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) MaxErrors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) OutputLocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) ScheduleExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) SyncCompliance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncCompliance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) Targets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociation) WaitForSuccessTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForSuccessTimeoutSeconds",
		&returns,
	)
	return returns
}


// Create a new `AWS::SSM::Association`.
func NewCfnAssociation(scope awscdk.Construct, id *string, props *CfnAssociationProps) CfnAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnAssociation{}

	_jsii_.Create(
		"monocdk.aws_ssm.CfnAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SSM::Association`.
func NewCfnAssociation_Override(c CfnAssociation, scope awscdk.Construct, id *string, props *CfnAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ssm.CfnAssociation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAssociation) SetApplyOnlyAtCronInterval(val interface{}) {
	_jsii_.Set(
		j,
		"applyOnlyAtCronInterval",
		val,
	)
}

func (j *jsiiProxy_CfnAssociation) SetAssociationName(val *string) {
	_jsii_.Set(
		j,
		"associationName",
		val,
	)
}

func (j *jsiiProxy_CfnAssociation) SetAutomationTargetParameterName(val *string) {
	_jsii_.Set(
		j,
		"automationTargetParameterName",
		val,
	)
}

func (j *jsiiProxy_CfnAssociation) SetCalendarNames(val *[]*string) {
	_jsii_.Set(
		j,
		"calendarNames",
		val,
	)
}

func (j *jsiiProxy_CfnAssociation) SetComplianceSeverity(val *string) {
	_jsii_.Set(
		j,
		"complianceSeverity",
		val,
	)
}

func (j *jsiiProxy_CfnAssociation) SetDocumentVersion(val *string) {
	_jsii_.Set(
		j,
		"documentVersion",
		val,
	)
}

func (j *jsiiProxy_CfnAssociation) SetInstanceId(val *string) {
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_CfnAssociation) SetMaxConcurrency(val *string) {
	_jsii_.Set(
		j,
		"maxConcurrency",
		val,
	)
}

func (j *jsiiProxy_CfnAssociation) SetMaxErrors(val *string) {
	_jsii_.Set(
		j,
		"maxErrors",
		val,
	)
}

func (j *jsiiProxy_CfnAssociation) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnAssociation) SetOutputLocation(val interface{}) {
	_jsii_.Set(
		j,
		"outputLocation",
		val,
	)
}

func (j *jsiiProxy_CfnAssociation) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnAssociation) SetScheduleExpression(val *string) {
	_jsii_.Set(
		j,
		"scheduleExpression",
		val,
	)
}

func (j *jsiiProxy_CfnAssociation) SetSyncCompliance(val *string) {
	_jsii_.Set(
		j,
		"syncCompliance",
		val,
	)
}

func (j *jsiiProxy_CfnAssociation) SetTargets(val interface{}) {
	_jsii_.Set(
		j,
		"targets",
		val,
	)
}

func (j *jsiiProxy_CfnAssociation) SetWaitForSuccessTimeoutSeconds(val *float64) {
	_jsii_.Set(
		j,
		"waitForSuccessTimeoutSeconds",
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
func CfnAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnAssociation",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAssociation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_ssm.CfnAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAssociation) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAssociation) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAssociation) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAssociation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAssociation) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssociation) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAssociation) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAssociation) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnAssociation) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAssociation) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAssociation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssociation) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssociation) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssociation) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssociation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `InstanceAssociationOutputLocation` is a property of the [AWS::SSM::Association](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html) resource that specifies an Amazon S3 bucket where you want to store the results of this association request.
//
// For the minimal permissions required to enable Amazon S3 output for an association, see [Creating associations](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-state-assoc.html) in the *Systems Manager User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceAssociationOutputLocationProperty := &instanceAssociationOutputLocationProperty{
//   	s3Location: &s3OutputLocationProperty{
//   		outputS3BucketName: jsii.String("outputS3BucketName"),
//   		outputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   		outputS3Region: jsii.String("outputS3Region"),
//   	},
//   }
//
type CfnAssociation_InstanceAssociationOutputLocationProperty struct {
	// `S3OutputLocation` is a property of the [InstanceAssociationOutputLocation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html) property that specifies an Amazon S3 bucket where you want to store the results of this request.
	S3Location interface{} `field:"optional" json:"s3Location" yaml:"s3Location"`
}

// `S3OutputLocation` is a property of the [AWS::SSM::Association](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html) resource that specifies an Amazon S3 bucket where you want to store the results of this association request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3OutputLocationProperty := &s3OutputLocationProperty{
//   	outputS3BucketName: jsii.String("outputS3BucketName"),
//   	outputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   	outputS3Region: jsii.String("outputS3Region"),
//   }
//
type CfnAssociation_S3OutputLocationProperty struct {
	// The name of the S3 bucket.
	OutputS3BucketName *string `field:"optional" json:"outputS3BucketName" yaml:"outputS3BucketName"`
	// The S3 bucket subfolder.
	OutputS3KeyPrefix *string `field:"optional" json:"outputS3KeyPrefix" yaml:"outputS3KeyPrefix"`
	// The AWS Region of the S3 bucket.
	OutputS3Region *string `field:"optional" json:"outputS3Region" yaml:"outputS3Region"`
}

// `Target` is a property of the [AWS::SSM::Association](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html) resource that specifies the targets for an SSM document in Systems Manager . You can target all instances in an AWS account by specifying the `InstanceIds` key with a value of `*` . To view a JSON and a YAML example that targets all instances, see "Create an association for all managed instances in an AWS account " on the Examples page.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetProperty := &targetProperty{
//   	key: jsii.String("key"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnAssociation_TargetProperty struct {
	// User-defined criteria for sending commands that target managed nodes that meet the criteria.
	Key *string `field:"required" json:"key" yaml:"key"`
	// User-defined criteria that maps to `Key` .
	//
	// For example, if you specified `tag:ServerRole` , you could specify `value:WebServer` to run a command on instances that include EC2 tags of `ServerRole,WebServer` .
	//
	// Depending on the type of target, the maximum number of values for a key might be lower than the global maximum of 50.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

// Properties for defining a `CfnAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnAssociationProps := &cfnAssociationProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	applyOnlyAtCronInterval: jsii.Boolean(false),
//   	associationName: jsii.String("associationName"),
//   	automationTargetParameterName: jsii.String("automationTargetParameterName"),
//   	calendarNames: []*string{
//   		jsii.String("calendarNames"),
//   	},
//   	complianceSeverity: jsii.String("complianceSeverity"),
//   	documentVersion: jsii.String("documentVersion"),
//   	instanceId: jsii.String("instanceId"),
//   	maxConcurrency: jsii.String("maxConcurrency"),
//   	maxErrors: jsii.String("maxErrors"),
//   	outputLocation: &instanceAssociationOutputLocationProperty{
//   		s3Location: &s3OutputLocationProperty{
//   			outputS3BucketName: jsii.String("outputS3BucketName"),
//   			outputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   			outputS3Region: jsii.String("outputS3Region"),
//   		},
//   	},
//   	parameters: map[string]interface{}{
//   		"parametersKey": parameters,
//   	},
//   	scheduleExpression: jsii.String("scheduleExpression"),
//   	syncCompliance: jsii.String("syncCompliance"),
//   	targets: []interface{}{
//   		&targetProperty{
//   			key: jsii.String("key"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	waitForSuccessTimeoutSeconds: jsii.Number(123),
//   }
//
type CfnAssociationProps struct {
	// The name of the SSM document that contains the configuration information for the instance.
	//
	// You can specify `Command` or `Automation` documents. The documents can be AWS -predefined documents, documents you created, or a document that is shared with you from another account. For SSM documents that are shared with you from other AWS accounts , you must specify the complete SSM document ARN, in the following format:
	//
	// `arn:partition:ssm:region:account-id:document/document-name`
	//
	// For example: `arn:aws:ssm:us-east-2:12345678912:document/My-Shared-Document`
	//
	// For AWS -predefined documents and SSM documents you created in your account, you only need to specify the document name. For example, AWS -ApplyPatchBaseline or My-Document.
	Name *string `field:"required" json:"name" yaml:"name"`
	// By default, when you create a new association, the system runs it immediately after it is created and then according to the schedule you specified.
	//
	// Specify this option if you don't want an association to run immediately after you create it. This parameter is not supported for rate expressions.
	ApplyOnlyAtCronInterval interface{} `field:"optional" json:"applyOnlyAtCronInterval" yaml:"applyOnlyAtCronInterval"`
	// Specify a descriptive name for the association.
	AssociationName *string `field:"optional" json:"associationName" yaml:"associationName"`
	// Choose the parameter that will define how your automation will branch out.
	//
	// This target is required for associations that use an Automation runbook and target resources by using rate controls. Automation is a capability of AWS Systems Manager .
	AutomationTargetParameterName *string `field:"optional" json:"automationTargetParameterName" yaml:"automationTargetParameterName"`
	// The names or Amazon Resource Names (ARNs) of the Change Calendar type documents your associations are gated under.
	//
	// The associations only run when that Change Calendar is open. For more information, see [AWS Systems Manager Change Calendar](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-change-calendar) .
	CalendarNames *[]*string `field:"optional" json:"calendarNames" yaml:"calendarNames"`
	// The severity level that is assigned to the association.
	ComplianceSeverity *string `field:"optional" json:"complianceSeverity" yaml:"complianceSeverity"`
	// The version of the SSM document to associate with the target.
	//
	// > Note the following important information.
	// >
	// > - State Manager doesn't support running associations that use a new version of a document if that document is shared from another account. State Manager always runs the `default` version of a document if shared from another account, even though the Systems Manager console shows that a new version was processed. If you want to run an association using a new version of a document shared form another account, you must set the document version to `default` .
	// > - `DocumentVersion` is not valid for documents owned by AWS , such as `AWS-RunPatchBaseline` or `AWS-UpdateSSMAgent` . If you specify `DocumentVersion` for an AWS document, the system returns the following error: "Error occurred during operation 'CreateAssociation'." (RequestToken: <token>, HandlerErrorCode: GeneralServiceException).
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// The ID of the instance that the SSM document is associated with.
	//
	// You must specify the `InstanceId` or `Targets` property.
	//
	// > `InstanceId` has been deprecated. To specify an instance ID for an association, use the `Targets` parameter. If you use the parameter `InstanceId` , you cannot use the parameters `AssociationName` , `DocumentVersion` , `MaxErrors` , `MaxConcurrency` , `OutputLocation` , or `ScheduleExpression` . To use these parameters, you must use the `Targets` parameter.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The maximum number of targets allowed to run the association at the same time.
	//
	// You can specify a number, for example 10, or a percentage of the target set, for example 10%. The default value is 100%, which means all targets run the association at the same time.
	//
	// If a new managed node starts and attempts to run an association while Systems Manager is running `MaxConcurrency` associations, the association is allowed to run. During the next association interval, the new managed node will process its association within the limit specified for `MaxConcurrency` .
	MaxConcurrency *string `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// The number of errors that are allowed before the system stops sending requests to run the association on additional targets.
	//
	// You can specify either an absolute number of errors, for example 10, or a percentage of the target set, for example 10%. If you specify 3, for example, the system stops sending requests when the fourth error is received. If you specify 0, then the system stops sending requests after the first error is returned. If you run an association on 50 managed nodes and set `MaxError` to 10%, then the system stops sending the request when the sixth error is received.
	//
	// Executions that are already running an association when `MaxErrors` is reached are allowed to complete, but some of these executions may fail as well. If you need to ensure that there won't be more than max-errors failed executions, set `MaxConcurrency` to 1 so that executions proceed one at a time.
	MaxErrors *string `field:"optional" json:"maxErrors" yaml:"maxErrors"`
	// An Amazon Simple Storage Service (Amazon S3) bucket where you want to store the output details of the request.
	OutputLocation interface{} `field:"optional" json:"outputLocation" yaml:"outputLocation"`
	// The parameters for the runtime configuration of the document.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A cron expression that specifies a schedule when the association runs.
	//
	// The schedule runs in Coordinated Universal Time (UTC).
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The mode for generating association compliance.
	//
	// You can specify `AUTO` or `MANUAL` . In `AUTO` mode, the system uses the status of the association execution to determine the compliance status. If the association execution runs successfully, then the association is `COMPLIANT` . If the association execution doesn't run successfully, the association is `NON-COMPLIANT` .
	//
	// In `MANUAL` mode, you must specify the `AssociationId` as a parameter for the PutComplianceItems API action. In this case, compliance data is not managed by State Manager. It is managed by your direct call to the PutComplianceItems API action.
	//
	// By default, all associations use `AUTO` mode.
	SyncCompliance *string `field:"optional" json:"syncCompliance" yaml:"syncCompliance"`
	// The targets for the association.
	//
	// You must specify the `InstanceId` or `Targets` property. You can target all instances in an AWS account by specifying the `InstanceIds` key with a value of `*` . To view a JSON and a YAML example that targets all instances, see "Create an association for all managed instances in an AWS account " on the Examples page.
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// The number of seconds the service should wait for the association status to show "Success" before proceeding with the stack execution.
	//
	// If the association status doesn't show "Success" after the specified number of seconds, then stack creation fails.
	WaitForSuccessTimeoutSeconds *float64 `field:"optional" json:"waitForSuccessTimeoutSeconds" yaml:"waitForSuccessTimeoutSeconds"`
}

// A CloudFormation `AWS::SSM::Document`.
//
// The `AWS::SSM::Document` resource creates a Systems Manager (SSM) document in AWS Systems Manager . This document defines the actions that Systems Manager performs on your AWS resources.
//
// > This resource does not support CloudFormation drift detection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var content interface{}
//
//   cfnDocument := awscdk.Aws_ssm.NewCfnDocument(this, jsii.String("MyCfnDocument"), &cfnDocumentProps{
//   	content: content,
//
//   	// the properties below are optional
//   	attachments: []interface{}{
//   		&attachmentsSourceProperty{
//   			key: jsii.String("key"),
//   			name: jsii.String("name"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	documentFormat: jsii.String("documentFormat"),
//   	documentType: jsii.String("documentType"),
//   	name: jsii.String("name"),
//   	requires: []interface{}{
//   		&documentRequiresProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetType: jsii.String("targetType"),
//   	updateMethod: jsii.String("updateMethod"),
//   	versionName: jsii.String("versionName"),
//   })
//
type CfnDocument interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A list of key-value pairs that describe attachments to a version of a document.
	Attachments() interface{}
	SetAttachments(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The content for the new SSM document in JSON or YAML.
	//
	// > This parameter also supports `String` data types.
	Content() interface{}
	SetContent(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Specify the document format for the request.
	//
	// JSON is the default format.
	DocumentFormat() *string
	SetDocumentFormat(val *string)
	// The type of document to create.
	//
	// *Allowed Values* : `ApplicationConfigurationSchema` | `Automation` | `Automation.ChangeTemplate` | `Command` | `DeploymentStrategy` | `Package` | `Policy` | `Session`
	DocumentType() *string
	SetDocumentType(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// A name for the SSM document.
	//
	// > You can't use the following strings as document name prefixes. These are reserved by AWS for use as document name prefixes:
	// >
	// > - `aws-`
	// > - `amazon`
	// > - `amzn`.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// A list of SSM documents required by a document.
	//
	// This parameter is used exclusively by AWS AppConfig . When a user creates an AWS AppConfig configuration in an SSM document, the user must also specify a required document for validation purposes. In this case, an `ApplicationConfiguration` document requires an `ApplicationConfigurationSchema` document for validation purposes. For more information, see [What is AWS AppConfig ?](https://docs.aws.amazon.com/appconfig/latest/userguide/what-is-appconfig.html) in the *AWS AppConfig User Guide* .
	Requires() interface{}
	SetRequires(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// AWS CloudFormation resource tags to apply to the document.
	//
	// Use tags to help you identify and categorize resources.
	Tags() awscdk.TagManager
	// Specify a target type to define the kinds of resources the document can run on.
	//
	// For example, to run a document on EC2 instances, specify the following value: `/AWS::EC2::Instance` . If you specify a value of '/' the document can run on all types of resources. If you don't specify a value, the document can't run on any resources. For a list of valid resource types, see [AWS resource and property types reference](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html) in the *AWS CloudFormation User Guide* .
	TargetType() *string
	SetTargetType(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// `AWS::SSM::Document.UpdateMethod`.
	UpdateMethod() *string
	SetUpdateMethod(val *string)
	// An optional field specifying the version of the artifact you are creating with the document.
	//
	// For example, "Release 12, Update 6". This value is unique across all versions of a document, and can't be changed.
	VersionName() *string
	SetVersionName(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDocument
type jsiiProxy_CfnDocument struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDocument) Attachments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) Content() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) DocumentFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) DocumentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) Requires() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requires",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) TargetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) UpdateMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocument) VersionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionName",
		&returns,
	)
	return returns
}


// Create a new `AWS::SSM::Document`.
func NewCfnDocument(scope awscdk.Construct, id *string, props *CfnDocumentProps) CfnDocument {
	_init_.Initialize()

	j := jsiiProxy_CfnDocument{}

	_jsii_.Create(
		"monocdk.aws_ssm.CfnDocument",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SSM::Document`.
func NewCfnDocument_Override(c CfnDocument, scope awscdk.Construct, id *string, props *CfnDocumentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ssm.CfnDocument",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDocument) SetAttachments(val interface{}) {
	_jsii_.Set(
		j,
		"attachments",
		val,
	)
}

func (j *jsiiProxy_CfnDocument) SetContent(val interface{}) {
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_CfnDocument) SetDocumentFormat(val *string) {
	_jsii_.Set(
		j,
		"documentFormat",
		val,
	)
}

func (j *jsiiProxy_CfnDocument) SetDocumentType(val *string) {
	_jsii_.Set(
		j,
		"documentType",
		val,
	)
}

func (j *jsiiProxy_CfnDocument) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDocument) SetRequires(val interface{}) {
	_jsii_.Set(
		j,
		"requires",
		val,
	)
}

func (j *jsiiProxy_CfnDocument) SetTargetType(val *string) {
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

func (j *jsiiProxy_CfnDocument) SetUpdateMethod(val *string) {
	_jsii_.Set(
		j,
		"updateMethod",
		val,
	)
}

func (j *jsiiProxy_CfnDocument) SetVersionName(val *string) {
	_jsii_.Set(
		j,
		"versionName",
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
func CfnDocument_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnDocument",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDocument_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnDocument",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDocument_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnDocument",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDocument_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_ssm.CfnDocument",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDocument) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDocument) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDocument) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDocument) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDocument) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDocument) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDocument) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDocument) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocument) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocument) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDocument) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDocument) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDocument) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocument) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDocument) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDocument) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocument) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocument) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDocument) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocument) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocument) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Identifying information about a document attachment, including the file name and a key-value pair that identifies the location of an attachment to a document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attachmentsSourceProperty := &attachmentsSourceProperty{
//   	key: jsii.String("key"),
//   	name: jsii.String("name"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnDocument_AttachmentsSourceProperty struct {
	// The key of a key-value pair that identifies the location of an attachment to a document.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The name of the document attachment file.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of a key-value pair that identifies the location of an attachment to a document.
	//
	// The format for *Value* depends on the type of key you specify.
	//
	// - For the key *SourceUrl* , the value is an S3 bucket location. For example:
	//
	// `"Values": [ "s3://doc-example-bucket/my-folder" ]`
	// - For the key *S3FileUrl* , the value is a file in an S3 bucket. For example:
	//
	// `"Values": [ "s3://doc-example-bucket/my-folder/my-file.py" ]`
	// - For the key *AttachmentReference* , the value is constructed from the name of another SSM document in your account, a version number of that document, and a file attached to that document version that you want to reuse. For example:
	//
	// `"Values": [ "MyOtherDocument/3/my-other-file.py" ]`
	//
	// However, if the SSM document is shared with you from another account, the full SSM document ARN must be specified instead of the document name only. For example:
	//
	// `"Values": [ "arn:aws:ssm:us-east-2:111122223333:document/OtherAccountDocument/3/their-file.py" ]`
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

// An SSM document required by the current document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentRequiresProperty := &documentRequiresProperty{
//   	name: jsii.String("name"),
//   	version: jsii.String("version"),
//   }
//
type CfnDocument_DocumentRequiresProperty struct {
	// The name of the required SSM document.
	//
	// The name can be an Amazon Resource Name (ARN).
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The document version required by the current document.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// Properties for defining a `CfnDocument`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var content interface{}
//
//   cfnDocumentProps := &cfnDocumentProps{
//   	content: content,
//
//   	// the properties below are optional
//   	attachments: []interface{}{
//   		&attachmentsSourceProperty{
//   			key: jsii.String("key"),
//   			name: jsii.String("name"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	documentFormat: jsii.String("documentFormat"),
//   	documentType: jsii.String("documentType"),
//   	name: jsii.String("name"),
//   	requires: []interface{}{
//   		&documentRequiresProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetType: jsii.String("targetType"),
//   	updateMethod: jsii.String("updateMethod"),
//   	versionName: jsii.String("versionName"),
//   }
//
type CfnDocumentProps struct {
	// The content for the new SSM document in JSON or YAML.
	//
	// > This parameter also supports `String` data types.
	Content interface{} `field:"required" json:"content" yaml:"content"`
	// A list of key-value pairs that describe attachments to a version of a document.
	Attachments interface{} `field:"optional" json:"attachments" yaml:"attachments"`
	// Specify the document format for the request.
	//
	// JSON is the default format.
	DocumentFormat *string `field:"optional" json:"documentFormat" yaml:"documentFormat"`
	// The type of document to create.
	//
	// *Allowed Values* : `ApplicationConfigurationSchema` | `Automation` | `Automation.ChangeTemplate` | `Command` | `DeploymentStrategy` | `Package` | `Policy` | `Session`
	DocumentType *string `field:"optional" json:"documentType" yaml:"documentType"`
	// A name for the SSM document.
	//
	// > You can't use the following strings as document name prefixes. These are reserved by AWS for use as document name prefixes:
	// >
	// > - `aws-`
	// > - `amazon`
	// > - `amzn`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of SSM documents required by a document.
	//
	// This parameter is used exclusively by AWS AppConfig . When a user creates an AWS AppConfig configuration in an SSM document, the user must also specify a required document for validation purposes. In this case, an `ApplicationConfiguration` document requires an `ApplicationConfigurationSchema` document for validation purposes. For more information, see [What is AWS AppConfig ?](https://docs.aws.amazon.com/appconfig/latest/userguide/what-is-appconfig.html) in the *AWS AppConfig User Guide* .
	Requires interface{} `field:"optional" json:"requires" yaml:"requires"`
	// AWS CloudFormation resource tags to apply to the document.
	//
	// Use tags to help you identify and categorize resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specify a target type to define the kinds of resources the document can run on.
	//
	// For example, to run a document on EC2 instances, specify the following value: `/AWS::EC2::Instance` . If you specify a value of '/' the document can run on all types of resources. If you don't specify a value, the document can't run on any resources. For a list of valid resource types, see [AWS resource and property types reference](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html) in the *AWS CloudFormation User Guide* .
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
	// `AWS::SSM::Document.UpdateMethod`.
	UpdateMethod *string `field:"optional" json:"updateMethod" yaml:"updateMethod"`
	// An optional field specifying the version of the artifact you are creating with the document.
	//
	// For example, "Release 12, Update 6". This value is unique across all versions of a document, and can't be changed.
	VersionName *string `field:"optional" json:"versionName" yaml:"versionName"`
}

// A CloudFormation `AWS::SSM::MaintenanceWindow`.
//
// The `AWS::SSM::MaintenanceWindow` resource represents general information about a maintenance window for AWS Systems Manager . Maintenance Windows let you define a schedule for when to perform potentially disruptive actions on your instances, such as patching an operating system (OS), updating drivers, or installing software. Each maintenance window has a schedule, a duration, a set of registered targets, and a set of registered tasks.
//
// For more information, see [Systems Manager Maintenance Windows](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-maintenance.html) in the *AWS Systems Manager User Guide* and [CreateMaintenanceWindow](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_CreateMaintenanceWindow.html) in the *AWS Systems Manager API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMaintenanceWindow := awscdk.Aws_ssm.NewCfnMaintenanceWindow(this, jsii.String("MyCfnMaintenanceWindow"), &cfnMaintenanceWindowProps{
//   	allowUnassociatedTargets: jsii.Boolean(false),
//   	cutoff: jsii.Number(123),
//   	duration: jsii.Number(123),
//   	name: jsii.String("name"),
//   	schedule: jsii.String("schedule"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	endDate: jsii.String("endDate"),
//   	scheduleOffset: jsii.Number(123),
//   	scheduleTimezone: jsii.String("scheduleTimezone"),
//   	startDate: jsii.String("startDate"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnMaintenanceWindow interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Enables a maintenance window task to run on managed instances, even if you have not registered those instances as targets.
	//
	// If enabled, then you must specify the unregistered instances (by instance ID) when you register a task with the maintenance window.
	AllowUnassociatedTargets() interface{}
	SetAllowUnassociatedTargets(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The number of hours before the end of the maintenance window that AWS Systems Manager stops scheduling new tasks for execution.
	Cutoff() *float64
	SetCutoff(val *float64)
	// A description of the maintenance window.
	Description() *string
	SetDescription(val *string)
	// The duration of the maintenance window in hours.
	Duration() *float64
	SetDuration(val *float64)
	// The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become inactive.
	EndDate() *string
	SetEndDate(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the maintenance window.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The schedule of the maintenance window in the form of a cron or rate expression.
	Schedule() *string
	SetSchedule(val *string)
	// The number of days to wait to run a maintenance window after the scheduled cron expression date and time.
	ScheduleOffset() *float64
	SetScheduleOffset(val *float64)
	// The time zone that the scheduled maintenance window executions are based on, in Internet Assigned Numbers Authority (IANA) format.
	ScheduleTimezone() *string
	SetScheduleTimezone(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become active.
	//
	// StartDate allows you to delay activation of the Maintenance Window until the specified future date.
	StartDate() *string
	SetStartDate(val *string)
	// Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs).
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a maintenance window to identify the type of tasks it will run, the types of targets, and the environment it will run in.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnMaintenanceWindow
type jsiiProxy_CfnMaintenanceWindow struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMaintenanceWindow) AllowUnassociatedTargets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnassociatedTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) Cutoff() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cutoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) Duration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) EndDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) ScheduleOffset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduleOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) ScheduleTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) StartDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindow) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SSM::MaintenanceWindow`.
func NewCfnMaintenanceWindow(scope awscdk.Construct, id *string, props *CfnMaintenanceWindowProps) CfnMaintenanceWindow {
	_init_.Initialize()

	j := jsiiProxy_CfnMaintenanceWindow{}

	_jsii_.Create(
		"monocdk.aws_ssm.CfnMaintenanceWindow",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SSM::MaintenanceWindow`.
func NewCfnMaintenanceWindow_Override(c CfnMaintenanceWindow, scope awscdk.Construct, id *string, props *CfnMaintenanceWindowProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ssm.CfnMaintenanceWindow",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindow) SetAllowUnassociatedTargets(val interface{}) {
	_jsii_.Set(
		j,
		"allowUnassociatedTargets",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindow) SetCutoff(val *float64) {
	_jsii_.Set(
		j,
		"cutoff",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindow) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindow) SetDuration(val *float64) {
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindow) SetEndDate(val *string) {
	_jsii_.Set(
		j,
		"endDate",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindow) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindow) SetSchedule(val *string) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindow) SetScheduleOffset(val *float64) {
	_jsii_.Set(
		j,
		"scheduleOffset",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindow) SetScheduleTimezone(val *string) {
	_jsii_.Set(
		j,
		"scheduleTimezone",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindow) SetStartDate(val *string) {
	_jsii_.Set(
		j,
		"startDate",
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
func CfnMaintenanceWindow_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnMaintenanceWindow",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnMaintenanceWindow_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnMaintenanceWindow",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnMaintenanceWindow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnMaintenanceWindow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMaintenanceWindow_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_ssm.CfnMaintenanceWindow",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindow) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindow) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindow) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindow) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindow) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindow) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindow) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindow) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindow) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindow) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindow) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnMaintenanceWindow) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindow) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindow) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindow) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnMaintenanceWindow) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindow) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindow) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindow) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindow) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnMaintenanceWindow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMaintenanceWindowProps := &cfnMaintenanceWindowProps{
//   	allowUnassociatedTargets: jsii.Boolean(false),
//   	cutoff: jsii.Number(123),
//   	duration: jsii.Number(123),
//   	name: jsii.String("name"),
//   	schedule: jsii.String("schedule"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	endDate: jsii.String("endDate"),
//   	scheduleOffset: jsii.Number(123),
//   	scheduleTimezone: jsii.String("scheduleTimezone"),
//   	startDate: jsii.String("startDate"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnMaintenanceWindowProps struct {
	// Enables a maintenance window task to run on managed instances, even if you have not registered those instances as targets.
	//
	// If enabled, then you must specify the unregistered instances (by instance ID) when you register a task with the maintenance window.
	AllowUnassociatedTargets interface{} `field:"required" json:"allowUnassociatedTargets" yaml:"allowUnassociatedTargets"`
	// The number of hours before the end of the maintenance window that AWS Systems Manager stops scheduling new tasks for execution.
	Cutoff *float64 `field:"required" json:"cutoff" yaml:"cutoff"`
	// The duration of the maintenance window in hours.
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// The name of the maintenance window.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schedule of the maintenance window in the form of a cron or rate expression.
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// A description of the maintenance window.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become inactive.
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// The number of days to wait to run a maintenance window after the scheduled cron expression date and time.
	ScheduleOffset *float64 `field:"optional" json:"scheduleOffset" yaml:"scheduleOffset"`
	// The time zone that the scheduled maintenance window executions are based on, in Internet Assigned Numbers Authority (IANA) format.
	ScheduleTimezone *string `field:"optional" json:"scheduleTimezone" yaml:"scheduleTimezone"`
	// The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become active.
	//
	// StartDate allows you to delay activation of the Maintenance Window until the specified future date.
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs).
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a maintenance window to identify the type of tasks it will run, the types of targets, and the environment it will run in.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SSM::MaintenanceWindowTarget`.
//
// The `AWS::SSM::MaintenanceWindowTarget` resource registers a target with a maintenance window for AWS Systems Manager . For more information, see [RegisterTargetWithMaintenanceWindow](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_RegisterTargetWithMaintenanceWindow.html) in the *AWS Systems Manager API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMaintenanceWindowTarget := awscdk.Aws_ssm.NewCfnMaintenanceWindowTarget(this, jsii.String("MyCfnMaintenanceWindowTarget"), &cfnMaintenanceWindowTargetProps{
//   	resourceType: jsii.String("resourceType"),
//   	targets: []interface{}{
//   		&targetsProperty{
//   			key: jsii.String("key"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	windowId: jsii.String("windowId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	ownerInformation: jsii.String("ownerInformation"),
//   })
//
type CfnMaintenanceWindowTarget interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// A description for the target.
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name for the maintenance window target.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// A user-provided value that will be included in any Amazon CloudWatch Events events that are raised while running tasks for these targets in this maintenance window.
	OwnerInformation() *string
	SetOwnerInformation(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The type of target that is being registered with the maintenance window.
	ResourceType() *string
	SetResourceType(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The targets to register with the maintenance window.
	//
	// In other words, the instances to run commands on when the maintenance window runs.
	//
	// You must specify targets by using the `WindowTargetIds` parameter.
	Targets() interface{}
	SetTargets(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The ID of the maintenance window to register the target with.
	WindowId() *string
	SetWindowId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnMaintenanceWindowTarget
type jsiiProxy_CfnMaintenanceWindowTarget struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) OwnerInformation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) Targets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) WindowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowId",
		&returns,
	)
	return returns
}


// Create a new `AWS::SSM::MaintenanceWindowTarget`.
func NewCfnMaintenanceWindowTarget(scope awscdk.Construct, id *string, props *CfnMaintenanceWindowTargetProps) CfnMaintenanceWindowTarget {
	_init_.Initialize()

	j := jsiiProxy_CfnMaintenanceWindowTarget{}

	_jsii_.Create(
		"monocdk.aws_ssm.CfnMaintenanceWindowTarget",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SSM::MaintenanceWindowTarget`.
func NewCfnMaintenanceWindowTarget_Override(c CfnMaintenanceWindowTarget, scope awscdk.Construct, id *string, props *CfnMaintenanceWindowTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ssm.CfnMaintenanceWindowTarget",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) SetOwnerInformation(val *string) {
	_jsii_.Set(
		j,
		"ownerInformation",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) SetResourceType(val *string) {
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) SetTargets(val interface{}) {
	_jsii_.Set(
		j,
		"targets",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTarget) SetWindowId(val *string) {
	_jsii_.Set(
		j,
		"windowId",
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
func CfnMaintenanceWindowTarget_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnMaintenanceWindowTarget",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnMaintenanceWindowTarget_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnMaintenanceWindowTarget",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnMaintenanceWindowTarget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnMaintenanceWindowTarget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMaintenanceWindowTarget_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_ssm.CfnMaintenanceWindowTarget",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTarget) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `Targets` property type specifies adding a target to a maintenance window target in AWS Systems Manager .
//
// `Targets` is a property of the [AWS::SSM::MaintenanceWindowTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtarget.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetsProperty := &targetsProperty{
//   	key: jsii.String("key"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnMaintenanceWindowTarget_TargetsProperty struct {
	// User-defined criteria for sending commands that target managed nodes that meet the criteria.
	Key *string `field:"required" json:"key" yaml:"key"`
	// User-defined criteria that maps to `Key` .
	//
	// For example, if you specified `tag:ServerRole` , you could specify `value:WebServer` to run a command on instances that include EC2 tags of `ServerRole,WebServer` .
	//
	// Depending on the type of target, the maximum number of values for a key might be lower than the global maximum of 50.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

// Properties for defining a `CfnMaintenanceWindowTarget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMaintenanceWindowTargetProps := &cfnMaintenanceWindowTargetProps{
//   	resourceType: jsii.String("resourceType"),
//   	targets: []interface{}{
//   		&targetsProperty{
//   			key: jsii.String("key"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	windowId: jsii.String("windowId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	ownerInformation: jsii.String("ownerInformation"),
//   }
//
type CfnMaintenanceWindowTargetProps struct {
	// The type of target that is being registered with the maintenance window.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// The targets to register with the maintenance window.
	//
	// In other words, the instances to run commands on when the maintenance window runs.
	//
	// You must specify targets by using the `WindowTargetIds` parameter.
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// The ID of the maintenance window to register the target with.
	WindowId *string `field:"required" json:"windowId" yaml:"windowId"`
	// A description for the target.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name for the maintenance window target.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A user-provided value that will be included in any Amazon CloudWatch Events events that are raised while running tasks for these targets in this maintenance window.
	OwnerInformation *string `field:"optional" json:"ownerInformation" yaml:"ownerInformation"`
}

// A CloudFormation `AWS::SSM::MaintenanceWindowTask`.
//
// The `AWS::SSM::MaintenanceWindowTask` resource defines information about a task for an AWS Systems Manager maintenance window. For more information, see [RegisterTaskWithMaintenanceWindow](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_RegisterTaskWithMaintenanceWindow.html) in the *AWS Systems Manager API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var taskParameters interface{}
//
//   cfnMaintenanceWindowTask := awscdk.Aws_ssm.NewCfnMaintenanceWindowTask(this, jsii.String("MyCfnMaintenanceWindowTask"), &cfnMaintenanceWindowTaskProps{
//   	priority: jsii.Number(123),
//   	taskArn: jsii.String("taskArn"),
//   	taskType: jsii.String("taskType"),
//   	windowId: jsii.String("windowId"),
//
//   	// the properties below are optional
//   	cutoffBehavior: jsii.String("cutoffBehavior"),
//   	description: jsii.String("description"),
//   	loggingInfo: &loggingInfoProperty{
//   		region: jsii.String("region"),
//   		s3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		s3Prefix: jsii.String("s3Prefix"),
//   	},
//   	maxConcurrency: jsii.String("maxConcurrency"),
//   	maxErrors: jsii.String("maxErrors"),
//   	name: jsii.String("name"),
//   	serviceRoleArn: jsii.String("serviceRoleArn"),
//   	targets: []interface{}{
//   		&targetProperty{
//   			key: jsii.String("key"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	taskInvocationParameters: &taskInvocationParametersProperty{
//   		maintenanceWindowAutomationParameters: &maintenanceWindowAutomationParametersProperty{
//   			documentVersion: jsii.String("documentVersion"),
//   			parameters: parameters,
//   		},
//   		maintenanceWindowLambdaParameters: &maintenanceWindowLambdaParametersProperty{
//   			clientContext: jsii.String("clientContext"),
//   			payload: jsii.String("payload"),
//   			qualifier: jsii.String("qualifier"),
//   		},
//   		maintenanceWindowRunCommandParameters: &maintenanceWindowRunCommandParametersProperty{
//   			cloudWatchOutputConfig: &cloudWatchOutputConfigProperty{
//   				cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   				cloudWatchOutputEnabled: jsii.Boolean(false),
//   			},
//   			comment: jsii.String("comment"),
//   			documentHash: jsii.String("documentHash"),
//   			documentHashType: jsii.String("documentHashType"),
//   			documentVersion: jsii.String("documentVersion"),
//   			notificationConfig: &notificationConfigProperty{
//   				notificationArn: jsii.String("notificationArn"),
//
//   				// the properties below are optional
//   				notificationEvents: []*string{
//   					jsii.String("notificationEvents"),
//   				},
//   				notificationType: jsii.String("notificationType"),
//   			},
//   			outputS3BucketName: jsii.String("outputS3BucketName"),
//   			outputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   			parameters: parameters,
//   			serviceRoleArn: jsii.String("serviceRoleArn"),
//   			timeoutSeconds: jsii.Number(123),
//   		},
//   		maintenanceWindowStepFunctionsParameters: &maintenanceWindowStepFunctionsParametersProperty{
//   			input: jsii.String("input"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	taskParameters: taskParameters,
//   })
//
type CfnMaintenanceWindowTask interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The specification for whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached.
	CutoffBehavior() *string
	SetCutoffBehavior(val *string)
	// A description of the task.
	Description() *string
	SetDescription(val *string)
	// Information about an Amazon S3 bucket to write task-level logs to.
	//
	// > `LoggingInfo` has been deprecated. To specify an Amazon S3 bucket to contain logs, instead use the `OutputS3BucketName` and `OutputS3KeyPrefix` options in the `TaskInvocationParameters` structure. For information about how Systems Manager handles these options for the supported maintenance window task types, see [AWS Systems Manager MaintenanceWindowTask TaskInvocationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html) .
	LoggingInfo() interface{}
	SetLoggingInfo(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The maximum number of targets this task can be run for, in parallel.
	//
	// > Although this element is listed as "Required: No", a value can be omitted only when you are registering or updating a [targetless task](https://docs.aws.amazon.com/systems-manager/latest/userguide/maintenance-windows-targetless-tasks.html) You must provide a value in all other cases.
	// >
	// > For maintenance window tasks without a target specified, you can't supply a value for this option. Instead, the system inserts a placeholder value of `1` . This value doesn't affect the running of your task.
	MaxConcurrency() *string
	SetMaxConcurrency(val *string)
	// The maximum number of errors allowed before this task stops being scheduled.
	//
	// > Although this element is listed as "Required: No", a value can be omitted only when you are registering or updating a [targetless task](https://docs.aws.amazon.com/systems-manager/latest/userguide/maintenance-windows-targetless-tasks.html) You must provide a value in all other cases.
	// >
	// > For maintenance window tasks without a target specified, you can't supply a value for this option. Instead, the system inserts a placeholder value of `1` . This value doesn't affect the running of your task.
	MaxErrors() *string
	SetMaxErrors(val *string)
	// The task name.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The priority of the task in the maintenance window.
	//
	// The lower the number, the higher the priority. Tasks that have the same priority are scheduled in parallel.
	Priority() *float64
	SetPriority(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) service role to use to publish Amazon Simple Notification Service (Amazon SNS) notifications for maintenance window Run Command tasks.
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The targets, either instances or window target IDs.
	//
	// - Specify instances using `Key=InstanceIds,Values= *instanceid1* , *instanceid2*` .
	// - Specify window target IDs using `Key=WindowTargetIds,Values= *window-target-id-1* , *window-target-id-2*` .
	Targets() interface{}
	SetTargets(val interface{})
	// The resource that the task uses during execution.
	//
	// For `RUN_COMMAND` and `AUTOMATION` task types, `TaskArn` is the SSM document name or Amazon Resource Name (ARN).
	//
	// For `LAMBDA` tasks, `TaskArn` is the function name or ARN.
	//
	// For `STEP_FUNCTIONS` tasks, `TaskArn` is the state machine ARN.
	TaskArn() *string
	SetTaskArn(val *string)
	// The parameters to pass to the task when it runs.
	//
	// Populate only the fields that match the task type. All other fields should be empty.
	//
	// > When you update a maintenance window task that has options specified in `TaskInvocationParameters` , you must provide again all the `TaskInvocationParameters` values that you want to retain. The values you do not specify again are removed. For example, suppose that when you registered a Run Command task, you specified `TaskInvocationParameters` values for `Comment` , `NotificationConfig` , and `OutputS3BucketName` . If you update the maintenance window task and specify only a different `OutputS3BucketName` value, the values for `Comment` and `NotificationConfig` are removed.
	TaskInvocationParameters() interface{}
	SetTaskInvocationParameters(val interface{})
	// The parameters to pass to the task when it runs.
	//
	// > `TaskParameters` has been deprecated. To specify parameters to pass to a task when it runs, instead use the `Parameters` option in the `TaskInvocationParameters` structure. For information about how Systems Manager handles these options for the supported maintenance window task types, see [MaintenanceWindowTaskInvocationParameters](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_MaintenanceWindowTaskInvocationParameters.html) .
	TaskParameters() interface{}
	SetTaskParameters(val interface{})
	// The type of task.
	//
	// Valid values: `RUN_COMMAND` , `AUTOMATION` , `LAMBDA` , `STEP_FUNCTIONS` .
	TaskType() *string
	SetTaskType(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The ID of the maintenance window where the task is registered.
	WindowId() *string
	SetWindowId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnMaintenanceWindowTask
type jsiiProxy_CfnMaintenanceWindowTask struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) CutoffBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cutoffBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) LoggingInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) MaxConcurrency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) MaxErrors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) Targets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) TaskArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) TaskInvocationParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taskInvocationParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) TaskParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taskParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) TaskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) WindowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowId",
		&returns,
	)
	return returns
}


// Create a new `AWS::SSM::MaintenanceWindowTask`.
func NewCfnMaintenanceWindowTask(scope awscdk.Construct, id *string, props *CfnMaintenanceWindowTaskProps) CfnMaintenanceWindowTask {
	_init_.Initialize()

	j := jsiiProxy_CfnMaintenanceWindowTask{}

	_jsii_.Create(
		"monocdk.aws_ssm.CfnMaintenanceWindowTask",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SSM::MaintenanceWindowTask`.
func NewCfnMaintenanceWindowTask_Override(c CfnMaintenanceWindowTask, scope awscdk.Construct, id *string, props *CfnMaintenanceWindowTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ssm.CfnMaintenanceWindowTask",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetCutoffBehavior(val *string) {
	_jsii_.Set(
		j,
		"cutoffBehavior",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetLoggingInfo(val interface{}) {
	_jsii_.Set(
		j,
		"loggingInfo",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetMaxConcurrency(val *string) {
	_jsii_.Set(
		j,
		"maxConcurrency",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetMaxErrors(val *string) {
	_jsii_.Set(
		j,
		"maxErrors",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetServiceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetTargets(val interface{}) {
	_jsii_.Set(
		j,
		"targets",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetTaskArn(val *string) {
	_jsii_.Set(
		j,
		"taskArn",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetTaskInvocationParameters(val interface{}) {
	_jsii_.Set(
		j,
		"taskInvocationParameters",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetTaskParameters(val interface{}) {
	_jsii_.Set(
		j,
		"taskParameters",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetTaskType(val *string) {
	_jsii_.Set(
		j,
		"taskType",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetWindowId(val *string) {
	_jsii_.Set(
		j,
		"windowId",
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
func CfnMaintenanceWindowTask_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnMaintenanceWindowTask",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnMaintenanceWindowTask_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnMaintenanceWindowTask",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnMaintenanceWindowTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnMaintenanceWindowTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMaintenanceWindowTask_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_ssm.CfnMaintenanceWindowTask",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Configuration options for sending command output to Amazon CloudWatch Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchOutputConfigProperty := &cloudWatchOutputConfigProperty{
//   	cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   	cloudWatchOutputEnabled: jsii.Boolean(false),
//   }
//
type CfnMaintenanceWindowTask_CloudWatchOutputConfigProperty struct {
	// The name of the CloudWatch Logs log group where you want to send command output.
	//
	// If you don't specify a group name, AWS Systems Manager automatically creates a log group for you. The log group uses the following naming format:
	//
	// `aws/ssm/ *SystemsManagerDocumentName*`.
	CloudWatchLogGroupName *string `field:"optional" json:"cloudWatchLogGroupName" yaml:"cloudWatchLogGroupName"`
	// Enables Systems Manager to send command output to CloudWatch Logs.
	CloudWatchOutputEnabled interface{} `field:"optional" json:"cloudWatchOutputEnabled" yaml:"cloudWatchOutputEnabled"`
}

// The `LoggingInfo` property type specifies information about the Amazon S3 bucket to write instance-level logs to.
//
// `LoggingInfo` is a property of the [AWS::SSM::MaintenanceWindowTask](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html) resource.
//
// > `LoggingInfo` has been deprecated. To specify an Amazon S3 bucket to contain logs, instead use the `OutputS3BucketName` and `OutputS3KeyPrefix` options in the `TaskInvocationParameters` structure. For information about how Systems Manager handles these options for the supported maintenance window task types, see [AWS Systems Manager MaintenanceWindowTask TaskInvocationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingInfoProperty := &loggingInfoProperty{
//   	region: jsii.String("region"),
//   	s3Bucket: jsii.String("s3Bucket"),
//
//   	// the properties below are optional
//   	s3Prefix: jsii.String("s3Prefix"),
//   }
//
type CfnMaintenanceWindowTask_LoggingInfoProperty struct {
	// The AWS Region where the S3 bucket is located.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The name of an S3 bucket where execution logs are stored .
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The Amazon S3 bucket subfolder.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
}

// The `MaintenanceWindowAutomationParameters` property type specifies the parameters for an `AUTOMATION` task type for a maintenance window task in AWS Systems Manager .
//
// `MaintenanceWindowAutomationParameters` is a property of the [TaskInvocationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html) property type.
//
// For information about available parameters in Automation runbooks, you can view the content of the runbook itself in the Systems Manager console. For information, see [View runbook content](https://docs.aws.amazon.com/systems-manager/latest/userguide/automation-documents-reference-details.html#view-automation-json) in the *AWS Systems Manager User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   maintenanceWindowAutomationParametersProperty := &maintenanceWindowAutomationParametersProperty{
//   	documentVersion: jsii.String("documentVersion"),
//   	parameters: parameters,
//   }
//
type CfnMaintenanceWindowTask_MaintenanceWindowAutomationParametersProperty struct {
	// The version of an Automation runbook to use during task execution.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// The parameters for the AUTOMATION task.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

// The `MaintenanceWindowLambdaParameters` property type specifies the parameters for a `LAMBDA` task type for a maintenance window task in AWS Systems Manager .
//
// `MaintenanceWindowLambdaParameters` is a property of the [TaskInvocationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceWindowLambdaParametersProperty := &maintenanceWindowLambdaParametersProperty{
//   	clientContext: jsii.String("clientContext"),
//   	payload: jsii.String("payload"),
//   	qualifier: jsii.String("qualifier"),
//   }
//
type CfnMaintenanceWindowTask_MaintenanceWindowLambdaParametersProperty struct {
	// Client-specific information to pass to the AWS Lambda function that you're invoking.
	//
	// You can then use the `context` variable to process the client information in your AWS Lambda function.
	ClientContext *string `field:"optional" json:"clientContext" yaml:"clientContext"`
	// JSON to provide to your AWS Lambda function as input.
	//
	// > Although `Type` is listed as "String" for this property, the payload content must be formatted as a Base64-encoded binary data object.
	//
	// *Length Constraint:* 4096.
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// An AWS Lambda function version or alias name.
	//
	// If you specify a function version, the action uses the qualified function Amazon Resource Name (ARN) to invoke a specific Lambda function. If you specify an alias name, the action uses the alias ARN to invoke the Lambda function version that the alias points to.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

// The `MaintenanceWindowRunCommandParameters` property type specifies the parameters for a `RUN_COMMAND` task type for a maintenance window task in AWS Systems Manager .
//
// This means that these parameters are the same as those for the `SendCommand` API call. For more information about `SendCommand` parameters, see [SendCommand](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_SendCommand.html) in the *AWS Systems Manager API Reference* .
//
// For information about available parameters in SSM Command documents, you can view the content of the document itself in the Systems Manager console. For information, see [Viewing SSM command document content](https://docs.aws.amazon.com/systems-manager/latest/userguide/viewing-ssm-document-content.html) in the *AWS Systems Manager User Guide* .
//
// `MaintenanceWindowRunCommandParameters` is a property of the [TaskInvocationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   maintenanceWindowRunCommandParametersProperty := &maintenanceWindowRunCommandParametersProperty{
//   	cloudWatchOutputConfig: &cloudWatchOutputConfigProperty{
//   		cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   		cloudWatchOutputEnabled: jsii.Boolean(false),
//   	},
//   	comment: jsii.String("comment"),
//   	documentHash: jsii.String("documentHash"),
//   	documentHashType: jsii.String("documentHashType"),
//   	documentVersion: jsii.String("documentVersion"),
//   	notificationConfig: &notificationConfigProperty{
//   		notificationArn: jsii.String("notificationArn"),
//
//   		// the properties below are optional
//   		notificationEvents: []*string{
//   			jsii.String("notificationEvents"),
//   		},
//   		notificationType: jsii.String("notificationType"),
//   	},
//   	outputS3BucketName: jsii.String("outputS3BucketName"),
//   	outputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   	parameters: parameters,
//   	serviceRoleArn: jsii.String("serviceRoleArn"),
//   	timeoutSeconds: jsii.Number(123),
//   }
//
type CfnMaintenanceWindowTask_MaintenanceWindowRunCommandParametersProperty struct {
	// Configuration options for sending command output to Amazon CloudWatch Logs.
	CloudWatchOutputConfig interface{} `field:"optional" json:"cloudWatchOutputConfig" yaml:"cloudWatchOutputConfig"`
	// Information about the command or commands to run.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The SHA-256 or SHA-1 hash created by the system when the document was created.
	//
	// SHA-1 hashes have been deprecated.
	DocumentHash *string `field:"optional" json:"documentHash" yaml:"documentHash"`
	// The SHA-256 or SHA-1 hash type.
	//
	// SHA-1 hashes are deprecated.
	DocumentHashType *string `field:"optional" json:"documentHashType" yaml:"documentHashType"`
	// The AWS Systems Manager document (SSM document) version to use in the request.
	//
	// You can specify `$DEFAULT` , `$LATEST` , or a specific version number. If you run commands by using the AWS CLI, then you must escape the first two options by using a backslash. If you specify a version number, then you don't need to use the backslash. For example:
	//
	// `--document-version "\$DEFAULT"`
	//
	// `--document-version "\$LATEST"`
	//
	// `--document-version "3"`.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// Configurations for sending notifications about command status changes on a per-managed node basis.
	NotificationConfig interface{} `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// The name of the Amazon Simple Storage Service (Amazon S3) bucket.
	OutputS3BucketName *string `field:"optional" json:"outputS3BucketName" yaml:"outputS3BucketName"`
	// The S3 bucket subfolder.
	OutputS3KeyPrefix *string `field:"optional" json:"outputS3KeyPrefix" yaml:"outputS3KeyPrefix"`
	// The parameters for the `RUN_COMMAND` task execution.
	//
	// The supported parameters are the same as those for the `SendCommand` API call. For more information, see [SendCommand](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_SendCommand.html) in the *AWS Systems Manager API Reference* .
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) service role to use to publish Amazon Simple Notification Service (Amazon SNS) notifications for maintenance window Run Command tasks.
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// If this time is reached and the command hasn't already started running, it doesn't run.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

// The `MaintenanceWindowStepFunctionsParameters` property type specifies the parameters for the execution of a `STEP_FUNCTIONS` task in a Systems Manager maintenance window.
//
// `MaintenanceWindowStepFunctionsParameters` is a property of the [TaskInvocationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceWindowStepFunctionsParametersProperty := &maintenanceWindowStepFunctionsParametersProperty{
//   	input: jsii.String("input"),
//   	name: jsii.String("name"),
//   }
//
type CfnMaintenanceWindowTask_MaintenanceWindowStepFunctionsParametersProperty struct {
	// The inputs for the `STEP_FUNCTIONS` task.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// The name of the `STEP_FUNCTIONS` task.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// The `NotificationConfig` property type specifies configurations for sending notifications for a maintenance window task in AWS Systems Manager .
//
// `NotificationConfig` is a property of the [MaintenanceWindowRunCommandParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationConfigProperty := &notificationConfigProperty{
//   	notificationArn: jsii.String("notificationArn"),
//
//   	// the properties below are optional
//   	notificationEvents: []*string{
//   		jsii.String("notificationEvents"),
//   	},
//   	notificationType: jsii.String("notificationType"),
//   }
//
type CfnMaintenanceWindowTask_NotificationConfigProperty struct {
	// An Amazon Resource Name (ARN) for an Amazon Simple Notification Service (Amazon SNS) topic.
	//
	// Run Command pushes notifications about command status changes to this topic.
	NotificationArn *string `field:"required" json:"notificationArn" yaml:"notificationArn"`
	// The different events that you can receive notifications for.
	//
	// These events include the following: `All` (events), `InProgress` , `Success` , `TimedOut` , `Cancelled` , `Failed` . To learn more about these events, see [Configuring Amazon SNS Notifications for AWS Systems Manager](https://docs.aws.amazon.com/systems-manager/latest/userguide/monitoring-sns-notifications.html) in the *AWS Systems Manager User Guide* .
	NotificationEvents *[]*string `field:"optional" json:"notificationEvents" yaml:"notificationEvents"`
	// The notification type.
	//
	// - `Command` : Receive notification when the status of a command changes.
	// - `Invocation` : For commands sent to multiple instances, receive notification on a per-instance basis when the status of a command changes.
	NotificationType *string `field:"optional" json:"notificationType" yaml:"notificationType"`
}

// The `Target` property type specifies targets (either instances or window target IDs).
//
// You specify instances by using `Key=InstanceIds,Values=< *instanceid1* >,< *instanceid2* >` . You specify window target IDs using `Key=WindowTargetIds,Values=< *window-target-id-1* >,< *window-target-id-2* >` for a maintenance window task in AWS Systems Manager .
//
// `Target` is a property of the [AWS::SSM::MaintenanceWindowTask](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html) property type.
//
// > To use `resource-groups:Name` as the key for a maintenance window target, specify the resource group as a `AWS::SSM::MaintenanceWindowTarget` type, and use the `Ref` function to specify the target for `AWS::SSM::MaintenanceWindowTask` . For an example, see *Create a Run Command task that targets instances using a resource group name* in [AWS::SSM::MaintenanceWindowTask Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#aws-resource-ssm-maintenancewindowtask--examples) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetProperty := &targetProperty{
//   	key: jsii.String("key"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnMaintenanceWindowTask_TargetProperty struct {
	// User-defined criteria for sending commands that target instances that meet the criteria.
	//
	// `Key` can be `InstanceIds` or `WindowTargetIds` . For more information about how to target instances within a maintenance window task, see [About 'register-task-with-maintenance-window' Options and Values](https://docs.aws.amazon.com/systems-manager/latest/userguide/register-tasks-options.html) in the *AWS Systems Manager User Guide* .
	Key *string `field:"required" json:"key" yaml:"key"`
	// User-defined criteria that maps to `Key` .
	//
	// For example, if you specify `InstanceIds` , you can specify `i-1234567890abcdef0,i-9876543210abcdef0` to run a command on two EC2 instances. For more information about how to target instances within a maintenance window task, see [About 'register-task-with-maintenance-window' Options and Values](https://docs.aws.amazon.com/systems-manager/latest/userguide/register-tasks-options.html) in the *AWS Systems Manager User Guide* .
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

// The `TaskInvocationParameters` property type specifies the task execution parameters for a maintenance window task in AWS Systems Manager .
//
// `TaskInvocationParameters` is a property of the [AWS::SSM::MaintenanceWindowTask](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   taskInvocationParametersProperty := &taskInvocationParametersProperty{
//   	maintenanceWindowAutomationParameters: &maintenanceWindowAutomationParametersProperty{
//   		documentVersion: jsii.String("documentVersion"),
//   		parameters: parameters,
//   	},
//   	maintenanceWindowLambdaParameters: &maintenanceWindowLambdaParametersProperty{
//   		clientContext: jsii.String("clientContext"),
//   		payload: jsii.String("payload"),
//   		qualifier: jsii.String("qualifier"),
//   	},
//   	maintenanceWindowRunCommandParameters: &maintenanceWindowRunCommandParametersProperty{
//   		cloudWatchOutputConfig: &cloudWatchOutputConfigProperty{
//   			cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   			cloudWatchOutputEnabled: jsii.Boolean(false),
//   		},
//   		comment: jsii.String("comment"),
//   		documentHash: jsii.String("documentHash"),
//   		documentHashType: jsii.String("documentHashType"),
//   		documentVersion: jsii.String("documentVersion"),
//   		notificationConfig: &notificationConfigProperty{
//   			notificationArn: jsii.String("notificationArn"),
//
//   			// the properties below are optional
//   			notificationEvents: []*string{
//   				jsii.String("notificationEvents"),
//   			},
//   			notificationType: jsii.String("notificationType"),
//   		},
//   		outputS3BucketName: jsii.String("outputS3BucketName"),
//   		outputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   		parameters: parameters,
//   		serviceRoleArn: jsii.String("serviceRoleArn"),
//   		timeoutSeconds: jsii.Number(123),
//   	},
//   	maintenanceWindowStepFunctionsParameters: &maintenanceWindowStepFunctionsParametersProperty{
//   		input: jsii.String("input"),
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnMaintenanceWindowTask_TaskInvocationParametersProperty struct {
	// The parameters for an `AUTOMATION` task type.
	MaintenanceWindowAutomationParameters interface{} `field:"optional" json:"maintenanceWindowAutomationParameters" yaml:"maintenanceWindowAutomationParameters"`
	// The parameters for a `LAMBDA` task type.
	MaintenanceWindowLambdaParameters interface{} `field:"optional" json:"maintenanceWindowLambdaParameters" yaml:"maintenanceWindowLambdaParameters"`
	// The parameters for a `RUN_COMMAND` task type.
	MaintenanceWindowRunCommandParameters interface{} `field:"optional" json:"maintenanceWindowRunCommandParameters" yaml:"maintenanceWindowRunCommandParameters"`
	// The parameters for a `STEP_FUNCTIONS` task type.
	MaintenanceWindowStepFunctionsParameters interface{} `field:"optional" json:"maintenanceWindowStepFunctionsParameters" yaml:"maintenanceWindowStepFunctionsParameters"`
}

// Properties for defining a `CfnMaintenanceWindowTask`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var taskParameters interface{}
//
//   cfnMaintenanceWindowTaskProps := &cfnMaintenanceWindowTaskProps{
//   	priority: jsii.Number(123),
//   	taskArn: jsii.String("taskArn"),
//   	taskType: jsii.String("taskType"),
//   	windowId: jsii.String("windowId"),
//
//   	// the properties below are optional
//   	cutoffBehavior: jsii.String("cutoffBehavior"),
//   	description: jsii.String("description"),
//   	loggingInfo: &loggingInfoProperty{
//   		region: jsii.String("region"),
//   		s3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		s3Prefix: jsii.String("s3Prefix"),
//   	},
//   	maxConcurrency: jsii.String("maxConcurrency"),
//   	maxErrors: jsii.String("maxErrors"),
//   	name: jsii.String("name"),
//   	serviceRoleArn: jsii.String("serviceRoleArn"),
//   	targets: []interface{}{
//   		&targetProperty{
//   			key: jsii.String("key"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	taskInvocationParameters: &taskInvocationParametersProperty{
//   		maintenanceWindowAutomationParameters: &maintenanceWindowAutomationParametersProperty{
//   			documentVersion: jsii.String("documentVersion"),
//   			parameters: parameters,
//   		},
//   		maintenanceWindowLambdaParameters: &maintenanceWindowLambdaParametersProperty{
//   			clientContext: jsii.String("clientContext"),
//   			payload: jsii.String("payload"),
//   			qualifier: jsii.String("qualifier"),
//   		},
//   		maintenanceWindowRunCommandParameters: &maintenanceWindowRunCommandParametersProperty{
//   			cloudWatchOutputConfig: &cloudWatchOutputConfigProperty{
//   				cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   				cloudWatchOutputEnabled: jsii.Boolean(false),
//   			},
//   			comment: jsii.String("comment"),
//   			documentHash: jsii.String("documentHash"),
//   			documentHashType: jsii.String("documentHashType"),
//   			documentVersion: jsii.String("documentVersion"),
//   			notificationConfig: &notificationConfigProperty{
//   				notificationArn: jsii.String("notificationArn"),
//
//   				// the properties below are optional
//   				notificationEvents: []*string{
//   					jsii.String("notificationEvents"),
//   				},
//   				notificationType: jsii.String("notificationType"),
//   			},
//   			outputS3BucketName: jsii.String("outputS3BucketName"),
//   			outputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   			parameters: parameters,
//   			serviceRoleArn: jsii.String("serviceRoleArn"),
//   			timeoutSeconds: jsii.Number(123),
//   		},
//   		maintenanceWindowStepFunctionsParameters: &maintenanceWindowStepFunctionsParametersProperty{
//   			input: jsii.String("input"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	taskParameters: taskParameters,
//   }
//
type CfnMaintenanceWindowTaskProps struct {
	// The priority of the task in the maintenance window.
	//
	// The lower the number, the higher the priority. Tasks that have the same priority are scheduled in parallel.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The resource that the task uses during execution.
	//
	// For `RUN_COMMAND` and `AUTOMATION` task types, `TaskArn` is the SSM document name or Amazon Resource Name (ARN).
	//
	// For `LAMBDA` tasks, `TaskArn` is the function name or ARN.
	//
	// For `STEP_FUNCTIONS` tasks, `TaskArn` is the state machine ARN.
	TaskArn *string `field:"required" json:"taskArn" yaml:"taskArn"`
	// The type of task.
	//
	// Valid values: `RUN_COMMAND` , `AUTOMATION` , `LAMBDA` , `STEP_FUNCTIONS` .
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// The ID of the maintenance window where the task is registered.
	WindowId *string `field:"required" json:"windowId" yaml:"windowId"`
	// The specification for whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached.
	CutoffBehavior *string `field:"optional" json:"cutoffBehavior" yaml:"cutoffBehavior"`
	// A description of the task.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Information about an Amazon S3 bucket to write task-level logs to.
	//
	// > `LoggingInfo` has been deprecated. To specify an Amazon S3 bucket to contain logs, instead use the `OutputS3BucketName` and `OutputS3KeyPrefix` options in the `TaskInvocationParameters` structure. For information about how Systems Manager handles these options for the supported maintenance window task types, see [AWS Systems Manager MaintenanceWindowTask TaskInvocationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html) .
	LoggingInfo interface{} `field:"optional" json:"loggingInfo" yaml:"loggingInfo"`
	// The maximum number of targets this task can be run for, in parallel.
	//
	// > Although this element is listed as "Required: No", a value can be omitted only when you are registering or updating a [targetless task](https://docs.aws.amazon.com/systems-manager/latest/userguide/maintenance-windows-targetless-tasks.html) You must provide a value in all other cases.
	// >
	// > For maintenance window tasks without a target specified, you can't supply a value for this option. Instead, the system inserts a placeholder value of `1` . This value doesn't affect the running of your task.
	MaxConcurrency *string `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// The maximum number of errors allowed before this task stops being scheduled.
	//
	// > Although this element is listed as "Required: No", a value can be omitted only when you are registering or updating a [targetless task](https://docs.aws.amazon.com/systems-manager/latest/userguide/maintenance-windows-targetless-tasks.html) You must provide a value in all other cases.
	// >
	// > For maintenance window tasks without a target specified, you can't supply a value for this option. Instead, the system inserts a placeholder value of `1` . This value doesn't affect the running of your task.
	MaxErrors *string `field:"optional" json:"maxErrors" yaml:"maxErrors"`
	// The task name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) service role to use to publish Amazon Simple Notification Service (Amazon SNS) notifications for maintenance window Run Command tasks.
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// The targets, either instances or window target IDs.
	//
	// - Specify instances using `Key=InstanceIds,Values= *instanceid1* , *instanceid2*` .
	// - Specify window target IDs using `Key=WindowTargetIds,Values= *window-target-id-1* , *window-target-id-2*` .
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// The parameters to pass to the task when it runs.
	//
	// Populate only the fields that match the task type. All other fields should be empty.
	//
	// > When you update a maintenance window task that has options specified in `TaskInvocationParameters` , you must provide again all the `TaskInvocationParameters` values that you want to retain. The values you do not specify again are removed. For example, suppose that when you registered a Run Command task, you specified `TaskInvocationParameters` values for `Comment` , `NotificationConfig` , and `OutputS3BucketName` . If you update the maintenance window task and specify only a different `OutputS3BucketName` value, the values for `Comment` and `NotificationConfig` are removed.
	TaskInvocationParameters interface{} `field:"optional" json:"taskInvocationParameters" yaml:"taskInvocationParameters"`
	// The parameters to pass to the task when it runs.
	//
	// > `TaskParameters` has been deprecated. To specify parameters to pass to a task when it runs, instead use the `Parameters` option in the `TaskInvocationParameters` structure. For information about how Systems Manager handles these options for the supported maintenance window task types, see [MaintenanceWindowTaskInvocationParameters](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_MaintenanceWindowTaskInvocationParameters.html) .
	TaskParameters interface{} `field:"optional" json:"taskParameters" yaml:"taskParameters"`
}

// A CloudFormation `AWS::SSM::Parameter`.
//
// The `AWS::SSM::Parameter` resource creates an SSM parameter in AWS Systems Manager Parameter Store.
//
// > To create an SSM parameter, you must have the AWS Identity and Access Management ( IAM ) permissions `ssm:PutParameter` and `ssm:AddTagsToResource` . On stack creation, AWS CloudFormation adds the following three tags to the parameter: `aws:cloudformation:stack-name` , `aws:cloudformation:logical-id` , and `aws:cloudformation:stack-id` , in addition to any custom tags you specify.
// >
// > To add, update, or remove tags during stack update, you must have IAM permissions for both `ssm:AddTagsToResource` and `ssm:RemoveTagsFromResource` . For more information, see [Managing Access Using Policies](https://docs.aws.amazon.com/systems-manager/latest/userguide/security-iam.html#security_iam_access-manage) in the *AWS Systems Manager User Guide* .
//
// For information about valid values for parameters, see [Requirements and Constraints for Parameter Names](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-paramstore-su-create.html#sysman-parameter-name-constraints) in the *AWS Systems Manager User Guide* and [PutParameter](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PutParameter.html) in the *AWS Systems Manager API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnParameter := awscdk.Aws_ssm.NewCfnParameter(this, jsii.String("MyCfnParameter"), &cfnParameterProps{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//
//   	// the properties below are optional
//   	allowedPattern: jsii.String("allowedPattern"),
//   	dataType: jsii.String("dataType"),
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	policies: jsii.String("policies"),
//   	tags: tags,
//   	tier: jsii.String("tier"),
//   })
//
type CfnParameter interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A regular expression used to validate the parameter value.
	//
	// For example, for String types with values restricted to numbers, you can specify the following: `AllowedPattern=^\d+$`.
	AllowedPattern() *string
	SetAllowedPattern(val *string)
	// Returns the type of the parameter.
	//
	// Valid values are `String` or `StringList` .
	AttrType() *string
	// Returns the value of the parameter.
	AttrValue() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The data type of the parameter, such as `text` or `aws:ec2:image` .
	//
	// The default is `text` .
	DataType() *string
	SetDataType(val *string)
	// Information about the parameter.
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the parameter.
	//
	// > The maximum length constraint listed below includes capacity for additional system attributes that aren't part of the name. The maximum length for a parameter name, including the full length of the parameter ARN, is 1011 characters. For example, the length of the following parameter name is 65 characters, not 20 characters: `arn:aws:ssm:us-east-2:111222333444:parameter/ExampleParameterName`
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Information about the policies assigned to a parameter.
	//
	// [Assigning parameter policies](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-policies.html) in the *AWS Systems Manager User Guide* .
	Policies() *string
	SetPolicies(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs).
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a Systems Manager parameter to identify the type of resource to which it applies, the environment, or the type of configuration data referenced by the parameter.
	Tags() awscdk.TagManager
	// The parameter tier.
	Tier() *string
	SetTier(val *string)
	// The type of parameter.
	//
	// > AWS CloudFormation doesn't support creating a `SecureString` parameter type.
	//
	// *Allowed Values* : String | StringList.
	Type() *string
	SetType(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The parameter value.
	//
	// > If type is `StringList` , the system returns a comma-separated string with no spaces between commas in the `Value` field.
	Value() *string
	SetValue(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnParameter
type jsiiProxy_CfnParameter struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnParameter) AllowedPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) AttrType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) AttrValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) DataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Policies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Create a new `AWS::SSM::Parameter`.
func NewCfnParameter(scope awscdk.Construct, id *string, props *CfnParameterProps) CfnParameter {
	_init_.Initialize()

	j := jsiiProxy_CfnParameter{}

	_jsii_.Create(
		"monocdk.aws_ssm.CfnParameter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SSM::Parameter`.
func NewCfnParameter_Override(c CfnParameter, scope awscdk.Construct, id *string, props *CfnParameterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ssm.CfnParameter",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnParameter) SetAllowedPattern(val *string) {
	_jsii_.Set(
		j,
		"allowedPattern",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetDataType(val *string) {
	_jsii_.Set(
		j,
		"dataType",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetPolicies(val *string) {
	_jsii_.Set(
		j,
		"policies",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetTier(val *string) {
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
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
func CfnParameter_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnParameter",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnParameter_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnParameter",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnParameter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnParameter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnParameter_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_ssm.CfnParameter",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnParameter) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnParameter) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnParameter) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnParameter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnParameter) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnParameter) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnParameter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnParameter) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnParameter) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnParameter) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnParameter) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnParameter) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnParameter) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnParameter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnParameter) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnParameter) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnParameter) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnParameter) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnParameter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnParameter) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnParameter) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnParameter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnParameterProps := &cfnParameterProps{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//
//   	// the properties below are optional
//   	allowedPattern: jsii.String("allowedPattern"),
//   	dataType: jsii.String("dataType"),
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	policies: jsii.String("policies"),
//   	tags: tags,
//   	tier: jsii.String("tier"),
//   }
//
type CfnParameterProps struct {
	// The type of parameter.
	//
	// > AWS CloudFormation doesn't support creating a `SecureString` parameter type.
	//
	// *Allowed Values* : String | StringList.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The parameter value.
	//
	// > If type is `StringList` , the system returns a comma-separated string with no spaces between commas in the `Value` field.
	Value *string `field:"required" json:"value" yaml:"value"`
	// A regular expression used to validate the parameter value.
	//
	// For example, for String types with values restricted to numbers, you can specify the following: `AllowedPattern=^\d+$`.
	AllowedPattern *string `field:"optional" json:"allowedPattern" yaml:"allowedPattern"`
	// The data type of the parameter, such as `text` or `aws:ec2:image` .
	//
	// The default is `text` .
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// Information about the parameter.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter.
	//
	// > The maximum length constraint listed below includes capacity for additional system attributes that aren't part of the name. The maximum length for a parameter name, including the full length of the parameter ARN, is 1011 characters. For example, the length of the following parameter name is 65 characters, not 20 characters: `arn:aws:ssm:us-east-2:111222333444:parameter/ExampleParameterName`
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Information about the policies assigned to a parameter.
	//
	// [Assigning parameter policies](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-policies.html) in the *AWS Systems Manager User Guide* .
	Policies *string `field:"optional" json:"policies" yaml:"policies"`
	// Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs).
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a Systems Manager parameter to identify the type of resource to which it applies, the environment, or the type of configuration data referenced by the parameter.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The parameter tier.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

// A CloudFormation `AWS::SSM::PatchBaseline`.
//
// The `AWS::SSM::PatchBaseline` resource defines the basic information for an AWS Systems Manager patch baseline. A patch baseline defines which patches are approved for installation on your instances.
//
// For more information, see [CreatePatchBaseline](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_CreatePatchBaseline.html) in the *AWS Systems Manager API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPatchBaseline := awscdk.Aws_ssm.NewCfnPatchBaseline(this, jsii.String("MyCfnPatchBaseline"), &cfnPatchBaselineProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	approvalRules: &ruleGroupProperty{
//   		patchRules: []interface{}{
//   			&ruleProperty{
//   				approveAfterDays: jsii.Number(123),
//   				approveUntilDate: jsii.String("approveUntilDate"),
//   				complianceLevel: jsii.String("complianceLevel"),
//   				enableNonSecurity: jsii.Boolean(false),
//   				patchFilterGroup: &patchFilterGroupProperty{
//   					patchFilters: []interface{}{
//   						&patchFilterProperty{
//   							key: jsii.String("key"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	approvedPatches: []*string{
//   		jsii.String("approvedPatches"),
//   	},
//   	approvedPatchesComplianceLevel: jsii.String("approvedPatchesComplianceLevel"),
//   	approvedPatchesEnableNonSecurity: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	globalFilters: &patchFilterGroupProperty{
//   		patchFilters: []interface{}{
//   			&patchFilterProperty{
//   				key: jsii.String("key"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	operatingSystem: jsii.String("operatingSystem"),
//   	patchGroups: []*string{
//   		jsii.String("patchGroups"),
//   	},
//   	rejectedPatches: []*string{
//   		jsii.String("rejectedPatches"),
//   	},
//   	rejectedPatchesAction: jsii.String("rejectedPatchesAction"),
//   	sources: []interface{}{
//   		&patchSourceProperty{
//   			configuration: jsii.String("configuration"),
//   			name: jsii.String("name"),
//   			products: []*string{
//   				jsii.String("products"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnPatchBaseline interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A set of rules used to include patches in the baseline.
	ApprovalRules() interface{}
	SetApprovalRules(val interface{})
	// A list of explicitly approved patches for the baseline.
	//
	// For information about accepted formats for lists of approved patches and rejected patches, see [About package name formats for approved and rejected patch lists](https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html) in the *AWS Systems Manager User Guide* .
	ApprovedPatches() *[]*string
	SetApprovedPatches(val *[]*string)
	// Defines the compliance level for approved patches.
	//
	// When an approved patch is reported as missing, this value describes the severity of the compliance violation. The default value is `UNSPECIFIED` .
	ApprovedPatchesComplianceLevel() *string
	SetApprovedPatchesComplianceLevel(val *string)
	// Indicates whether the list of approved patches includes non-security updates that should be applied to the managed nodes.
	//
	// The default value is `false` . Applies to Linux managed nodes only.
	ApprovedPatchesEnableNonSecurity() interface{}
	SetApprovedPatchesEnableNonSecurity(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// A description of the patch baseline.
	Description() *string
	SetDescription(val *string)
	// A set of global filters used to include patches in the baseline.
	GlobalFilters() interface{}
	SetGlobalFilters(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the patch baseline.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Defines the operating system the patch baseline applies to.
	//
	// The default value is `WINDOWS` .
	OperatingSystem() *string
	SetOperatingSystem(val *string)
	// The name of the patch group to be registered with the patch baseline.
	PatchGroups() *[]*string
	SetPatchGroups(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// A list of explicitly rejected patches for the baseline.
	//
	// For information about accepted formats for lists of approved patches and rejected patches, see [About package name formats for approved and rejected patch lists](https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html) in the *AWS Systems Manager User Guide* .
	RejectedPatches() *[]*string
	SetRejectedPatches(val *[]*string)
	// The action for Patch Manager to take on patches included in the `RejectedPackages` list.
	//
	// - *`ALLOW_AS_DEPENDENCY`* : A package in the `Rejected` patches list is installed only if it is a dependency of another package. It is considered compliant with the patch baseline, and its status is reported as `InstalledOther` . This is the default action if no option is specified.
	// - *`BLOCK`* : Packages in the `RejectedPatches` list, and packages that include them as dependencies, aren't installed under any circumstances. If a package was installed before it was added to the Rejected patches list, it is considered non-compliant with the patch baseline, and its status is reported as `InstalledRejected` .
	RejectedPatchesAction() *string
	SetRejectedPatchesAction(val *string)
	// Information about the patches to use to update the managed nodes, including target operating systems and source repositories.
	//
	// Applies to Linux managed nodes only.
	Sources() interface{}
	SetSources(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Optional metadata that you assign to a resource.
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a patch baseline to identify the severity level of patches it specifies and the operating system family it applies to.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnPatchBaseline
type jsiiProxy_CfnPatchBaseline struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPatchBaseline) ApprovalRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) ApprovedPatches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"approvedPatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) ApprovedPatchesComplianceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvedPatchesComplianceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) ApprovedPatchesEnableNonSecurity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvedPatchesEnableNonSecurity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) GlobalFilters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) OperatingSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) PatchGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"patchGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) RejectedPatches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rejectedPatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) RejectedPatchesAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rejectedPatchesAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) Sources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaseline) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SSM::PatchBaseline`.
func NewCfnPatchBaseline(scope awscdk.Construct, id *string, props *CfnPatchBaselineProps) CfnPatchBaseline {
	_init_.Initialize()

	j := jsiiProxy_CfnPatchBaseline{}

	_jsii_.Create(
		"monocdk.aws_ssm.CfnPatchBaseline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SSM::PatchBaseline`.
func NewCfnPatchBaseline_Override(c CfnPatchBaseline, scope awscdk.Construct, id *string, props *CfnPatchBaselineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ssm.CfnPatchBaseline",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPatchBaseline) SetApprovalRules(val interface{}) {
	_jsii_.Set(
		j,
		"approvalRules",
		val,
	)
}

func (j *jsiiProxy_CfnPatchBaseline) SetApprovedPatches(val *[]*string) {
	_jsii_.Set(
		j,
		"approvedPatches",
		val,
	)
}

func (j *jsiiProxy_CfnPatchBaseline) SetApprovedPatchesComplianceLevel(val *string) {
	_jsii_.Set(
		j,
		"approvedPatchesComplianceLevel",
		val,
	)
}

func (j *jsiiProxy_CfnPatchBaseline) SetApprovedPatchesEnableNonSecurity(val interface{}) {
	_jsii_.Set(
		j,
		"approvedPatchesEnableNonSecurity",
		val,
	)
}

func (j *jsiiProxy_CfnPatchBaseline) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnPatchBaseline) SetGlobalFilters(val interface{}) {
	_jsii_.Set(
		j,
		"globalFilters",
		val,
	)
}

func (j *jsiiProxy_CfnPatchBaseline) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnPatchBaseline) SetOperatingSystem(val *string) {
	_jsii_.Set(
		j,
		"operatingSystem",
		val,
	)
}

func (j *jsiiProxy_CfnPatchBaseline) SetPatchGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"patchGroups",
		val,
	)
}

func (j *jsiiProxy_CfnPatchBaseline) SetRejectedPatches(val *[]*string) {
	_jsii_.Set(
		j,
		"rejectedPatches",
		val,
	)
}

func (j *jsiiProxy_CfnPatchBaseline) SetRejectedPatchesAction(val *string) {
	_jsii_.Set(
		j,
		"rejectedPatchesAction",
		val,
	)
}

func (j *jsiiProxy_CfnPatchBaseline) SetSources(val interface{}) {
	_jsii_.Set(
		j,
		"sources",
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
func CfnPatchBaseline_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnPatchBaseline",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPatchBaseline_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnPatchBaseline",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPatchBaseline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnPatchBaseline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPatchBaseline_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_ssm.CfnPatchBaseline",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPatchBaseline) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPatchBaseline) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPatchBaseline) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPatchBaseline) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPatchBaseline) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPatchBaseline) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPatchBaseline) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPatchBaseline) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPatchBaseline) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPatchBaseline) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPatchBaseline) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPatchBaseline) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPatchBaseline) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPatchBaseline) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPatchBaseline) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPatchBaseline) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPatchBaseline) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPatchBaseline) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPatchBaseline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPatchBaseline) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPatchBaseline) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `PatchFilterGroup` property type specifies a set of patch filters for an AWS Systems Manager patch baseline, typically used for approval rules for a Systems Manager patch baseline.
//
// `PatchFilterGroup` is the property type for the `GlobalFilters` property of the [AWS::SSM::PatchBaseline](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html) resource and the `PatchFilterGroup` property of the [Rule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   patchFilterGroupProperty := &patchFilterGroupProperty{
//   	patchFilters: []interface{}{
//   		&patchFilterProperty{
//   			key: jsii.String("key"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
type CfnPatchBaseline_PatchFilterGroupProperty struct {
	// The set of patch filters that make up the group.
	PatchFilters interface{} `field:"optional" json:"patchFilters" yaml:"patchFilters"`
}

// The `PatchFilter` property type defines a patch filter for an AWS Systems Manager patch baseline.
//
// The `PatchFilters` property of the [PatchFilterGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-patchfiltergroup.html) property type contains a list of `PatchFilter` property types.
//
// You can view lists of valid values for the patch properties by running the `DescribePatchProperties` command. For more information, see [DescribePatchProperties](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DescribePatchProperties.html) in the *AWS Systems Manager API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   patchFilterProperty := &patchFilterProperty{
//   	key: jsii.String("key"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnPatchBaseline_PatchFilterProperty struct {
	// The key for the filter.
	//
	// For information about valid keys, see [PatchFilter](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchFilter.html) in the *AWS Systems Manager API Reference* .
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the filter key.
	//
	// For information about valid values for each key based on operating system type, see [PatchFilter](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchFilter.html) in the *AWS Systems Manager API Reference* .
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

// `PatchSource` is the property type for the `Sources` resource of the [AWS::SSM::PatchBaseline](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html) resource.
//
// The AWS CloudFormation `AWS::SSM::PatchSource` resource is used to provide information about the patches to use to update target instances, including target operating systems and source repository. Applies to Linux instances only.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   patchSourceProperty := &patchSourceProperty{
//   	configuration: jsii.String("configuration"),
//   	name: jsii.String("name"),
//   	products: []*string{
//   		jsii.String("products"),
//   	},
//   }
//
type CfnPatchBaseline_PatchSourceProperty struct {
	// The value of the yum repo configuration. For example:.
	//
	// `[main]`
	//
	// `name=MyCustomRepository`
	//
	// `baseurl=https://my-custom-repository`
	//
	// `enabled=1`
	//
	// > For information about other options available for your yum repository configuration, see [dnf.conf(5)](https://docs.aws.amazon.com/https://man7.org/linux/man-pages/man5/dnf.conf.5.html) .
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// The name specified to identify the patch source.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The specific operating system versions a patch repository applies to, such as "Ubuntu16.04", "AmazonLinux2016.09", "RedhatEnterpriseLinux7.2" or "Suse12.7". For lists of supported product values, see [PatchFilter](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchFilter.html) in the *AWS Systems Manager API Reference* .
	Products *[]*string `field:"optional" json:"products" yaml:"products"`
}

// The `RuleGroup` property type specifies a set of rules that define the approval rules for an AWS Systems Manager patch baseline.
//
// `RuleGroup` is the property type for the `ApprovalRules` property of the [AWS::SSM::PatchBaseline](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleGroupProperty := &ruleGroupProperty{
//   	patchRules: []interface{}{
//   		&ruleProperty{
//   			approveAfterDays: jsii.Number(123),
//   			approveUntilDate: jsii.String("approveUntilDate"),
//   			complianceLevel: jsii.String("complianceLevel"),
//   			enableNonSecurity: jsii.Boolean(false),
//   			patchFilterGroup: &patchFilterGroupProperty{
//   				patchFilters: []interface{}{
//   					&patchFilterProperty{
//   						key: jsii.String("key"),
//   						values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnPatchBaseline_RuleGroupProperty struct {
	// The rules that make up the rule group.
	PatchRules interface{} `field:"optional" json:"patchRules" yaml:"patchRules"`
}

// The `Rule` property type specifies an approval rule for a Systems Manager patch baseline.
//
// The `PatchRules` property of the [RuleGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rulegroup.html) property type contains a list of `Rule` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleProperty := &ruleProperty{
//   	approveAfterDays: jsii.Number(123),
//   	approveUntilDate: jsii.String("approveUntilDate"),
//   	complianceLevel: jsii.String("complianceLevel"),
//   	enableNonSecurity: jsii.Boolean(false),
//   	patchFilterGroup: &patchFilterGroupProperty{
//   		patchFilters: []interface{}{
//   			&patchFilterProperty{
//   				key: jsii.String("key"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnPatchBaseline_RuleProperty struct {
	// The number of days after the release date of each patch matched by the rule that the patch is marked as approved in the patch baseline.
	//
	// For example, a value of `7` means that patches are approved seven days after they are released.
	//
	// You must specify a value for `ApproveAfterDays` .
	//
	// Exception: Not supported on Debian Server or Ubuntu Server.
	ApproveAfterDays *float64 `field:"optional" json:"approveAfterDays" yaml:"approveAfterDays"`
	// The cutoff date for auto approval of released patches.
	//
	// Any patches released on or before this date are installed automatically. Not supported on Debian Server or Ubuntu Server.
	//
	// Enter dates in the format `YYYY-MM-DD` . For example, `2021-12-31` .
	ApproveUntilDate *string `field:"optional" json:"approveUntilDate" yaml:"approveUntilDate"`
	// A compliance severity level for all approved patches in a patch baseline.
	//
	// Valid compliance severity levels include the following: `UNSPECIFIED` , `CRITICAL` , `HIGH` , `MEDIUM` , `LOW` , and `INFORMATIONAL` .
	ComplianceLevel *string `field:"optional" json:"complianceLevel" yaml:"complianceLevel"`
	// For managed nodes identified by the approval rule filters, enables a patch baseline to apply non-security updates available in the specified repository.
	//
	// The default value is `false` . Applies to Linux managed nodes only.
	EnableNonSecurity interface{} `field:"optional" json:"enableNonSecurity" yaml:"enableNonSecurity"`
	// The patch filter group that defines the criteria for the rule.
	PatchFilterGroup interface{} `field:"optional" json:"patchFilterGroup" yaml:"patchFilterGroup"`
}

// Properties for defining a `CfnPatchBaseline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPatchBaselineProps := &cfnPatchBaselineProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	approvalRules: &ruleGroupProperty{
//   		patchRules: []interface{}{
//   			&ruleProperty{
//   				approveAfterDays: jsii.Number(123),
//   				approveUntilDate: jsii.String("approveUntilDate"),
//   				complianceLevel: jsii.String("complianceLevel"),
//   				enableNonSecurity: jsii.Boolean(false),
//   				patchFilterGroup: &patchFilterGroupProperty{
//   					patchFilters: []interface{}{
//   						&patchFilterProperty{
//   							key: jsii.String("key"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	approvedPatches: []*string{
//   		jsii.String("approvedPatches"),
//   	},
//   	approvedPatchesComplianceLevel: jsii.String("approvedPatchesComplianceLevel"),
//   	approvedPatchesEnableNonSecurity: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	globalFilters: &patchFilterGroupProperty{
//   		patchFilters: []interface{}{
//   			&patchFilterProperty{
//   				key: jsii.String("key"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	operatingSystem: jsii.String("operatingSystem"),
//   	patchGroups: []*string{
//   		jsii.String("patchGroups"),
//   	},
//   	rejectedPatches: []*string{
//   		jsii.String("rejectedPatches"),
//   	},
//   	rejectedPatchesAction: jsii.String("rejectedPatchesAction"),
//   	sources: []interface{}{
//   		&patchSourceProperty{
//   			configuration: jsii.String("configuration"),
//   			name: jsii.String("name"),
//   			products: []*string{
//   				jsii.String("products"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPatchBaselineProps struct {
	// The name of the patch baseline.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A set of rules used to include patches in the baseline.
	ApprovalRules interface{} `field:"optional" json:"approvalRules" yaml:"approvalRules"`
	// A list of explicitly approved patches for the baseline.
	//
	// For information about accepted formats for lists of approved patches and rejected patches, see [About package name formats for approved and rejected patch lists](https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html) in the *AWS Systems Manager User Guide* .
	ApprovedPatches *[]*string `field:"optional" json:"approvedPatches" yaml:"approvedPatches"`
	// Defines the compliance level for approved patches.
	//
	// When an approved patch is reported as missing, this value describes the severity of the compliance violation. The default value is `UNSPECIFIED` .
	ApprovedPatchesComplianceLevel *string `field:"optional" json:"approvedPatchesComplianceLevel" yaml:"approvedPatchesComplianceLevel"`
	// Indicates whether the list of approved patches includes non-security updates that should be applied to the managed nodes.
	//
	// The default value is `false` . Applies to Linux managed nodes only.
	ApprovedPatchesEnableNonSecurity interface{} `field:"optional" json:"approvedPatchesEnableNonSecurity" yaml:"approvedPatchesEnableNonSecurity"`
	// A description of the patch baseline.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A set of global filters used to include patches in the baseline.
	GlobalFilters interface{} `field:"optional" json:"globalFilters" yaml:"globalFilters"`
	// Defines the operating system the patch baseline applies to.
	//
	// The default value is `WINDOWS` .
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// The name of the patch group to be registered with the patch baseline.
	PatchGroups *[]*string `field:"optional" json:"patchGroups" yaml:"patchGroups"`
	// A list of explicitly rejected patches for the baseline.
	//
	// For information about accepted formats for lists of approved patches and rejected patches, see [About package name formats for approved and rejected patch lists](https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html) in the *AWS Systems Manager User Guide* .
	RejectedPatches *[]*string `field:"optional" json:"rejectedPatches" yaml:"rejectedPatches"`
	// The action for Patch Manager to take on patches included in the `RejectedPackages` list.
	//
	// - *`ALLOW_AS_DEPENDENCY`* : A package in the `Rejected` patches list is installed only if it is a dependency of another package. It is considered compliant with the patch baseline, and its status is reported as `InstalledOther` . This is the default action if no option is specified.
	// - *`BLOCK`* : Packages in the `RejectedPatches` list, and packages that include them as dependencies, aren't installed under any circumstances. If a package was installed before it was added to the Rejected patches list, it is considered non-compliant with the patch baseline, and its status is reported as `InstalledRejected` .
	RejectedPatchesAction *string `field:"optional" json:"rejectedPatchesAction" yaml:"rejectedPatchesAction"`
	// Information about the patches to use to update the managed nodes, including target operating systems and source repositories.
	//
	// Applies to Linux managed nodes only.
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// Optional metadata that you assign to a resource.
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a patch baseline to identify the severity level of patches it specifies and the operating system family it applies to.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SSM::ResourceDataSync`.
//
// The `AWS::SSM::ResourceDataSync` resource creates, updates, or deletes a resource data sync for AWS Systems Manager . A resource data sync helps you view data from multiple sources in a single location. Systems Manager offers two types of resource data sync: `SyncToDestination` and `SyncFromSource` .
//
// You can configure Systems Manager Inventory to use the `SyncToDestination` type to synchronize Inventory data from multiple AWS Regions to a single Amazon S3 bucket.
//
// You can configure Systems Manager Explorer to use the `SyncFromSource` type to synchronize operational work items (OpsItems) and operational data (OpsData) from multiple AWS Regions . This type can synchronize OpsItems and OpsData from multiple AWS accounts and Regions or from an `EntireOrganization` by using AWS Organizations .
//
// A resource data sync is an asynchronous operation that returns immediately. After a successful initial sync is completed, the system continuously syncs data.
//
// By default, data is not encrypted in Amazon S3 . We strongly recommend that you enable encryption in Amazon S3 to ensure secure data storage. We also recommend that you secure access to the Amazon S3 bucket by creating a restrictive bucket policy.
//
// For more information, see [Configuring Inventory Collection](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-inventory-configuring.html#sysman-inventory-datasync) and [Setting Up Systems Manager Explorer to Display Data from Multiple Accounts and Regions](https://docs.aws.amazon.com/systems-manager/latest/userguide/Explorer-resource-data-sync.html) in the *AWS Systems Manager User Guide* .
//
// Important: The following *Syntax* section shows all fields that are supported for a resource data sync. The *Examples* section below shows the recommended way to specify configurations for each sync type. Please see the *Examples* section when you create your resource data sync.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceDataSync := awscdk.Aws_ssm.NewCfnResourceDataSync(this, jsii.String("MyCfnResourceDataSync"), &cfnResourceDataSyncProps{
//   	syncName: jsii.String("syncName"),
//
//   	// the properties below are optional
//   	bucketName: jsii.String("bucketName"),
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	bucketRegion: jsii.String("bucketRegion"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   	s3Destination: &s3DestinationProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketRegion: jsii.String("bucketRegion"),
//   		syncFormat: jsii.String("syncFormat"),
//
//   		// the properties below are optional
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		kmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	syncFormat: jsii.String("syncFormat"),
//   	syncSource: &syncSourceProperty{
//   		sourceRegions: []*string{
//   			jsii.String("sourceRegions"),
//   		},
//   		sourceType: jsii.String("sourceType"),
//
//   		// the properties below are optional
//   		awsOrganizationsSource: &awsOrganizationsSourceProperty{
//   			organizationSourceType: jsii.String("organizationSourceType"),
//
//   			// the properties below are optional
//   			organizationalUnits: []*string{
//   				jsii.String("organizationalUnits"),
//   			},
//   		},
//   		includeFutureRegions: jsii.Boolean(false),
//   	},
//   	syncType: jsii.String("syncType"),
//   })
//
type CfnResourceDataSync interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of the resource data sync.
	AttrSyncName() *string
	// The name of the S3 bucket where the aggregated data is stored.
	BucketName() *string
	SetBucketName(val *string)
	// An Amazon S3 prefix for the bucket.
	BucketPrefix() *string
	SetBucketPrefix(val *string)
	// The AWS Region with the S3 bucket targeted by the resource data sync.
	BucketRegion() *string
	SetBucketRegion(val *string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The ARN of an encryption key for a destination in Amazon S3 .
	//
	// You can use a KMS key to encrypt inventory data in Amazon S3 . You must specify a key that exist in the same region as the destination Amazon S3 bucket.
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
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// Configuration information for the target S3 bucket.
	S3Destination() interface{}
	SetS3Destination(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A supported sync format.
	//
	// The following format is currently supported: JsonSerDe.
	SyncFormat() *string
	SetSyncFormat(val *string)
	// A name for the resource data sync.
	SyncName() *string
	SetSyncName(val *string)
	// Information about the source where the data was synchronized.
	SyncSource() interface{}
	SetSyncSource(val interface{})
	// The type of resource data sync.
	//
	// If `SyncType` is `SyncToDestination` , then the resource data sync synchronizes data to an S3 bucket. If the `SyncType` is `SyncFromSource` then the resource data sync synchronizes data from AWS Organizations or from multiple AWS Regions .
	SyncType() *string
	SetSyncType(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnResourceDataSync
type jsiiProxy_CfnResourceDataSync struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResourceDataSync) AttrSyncName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSyncName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) BucketPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) BucketRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) S3Destination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) SyncFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) SyncName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) SyncSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) SyncType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSync) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SSM::ResourceDataSync`.
func NewCfnResourceDataSync(scope awscdk.Construct, id *string, props *CfnResourceDataSyncProps) CfnResourceDataSync {
	_init_.Initialize()

	j := jsiiProxy_CfnResourceDataSync{}

	_jsii_.Create(
		"monocdk.aws_ssm.CfnResourceDataSync",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SSM::ResourceDataSync`.
func NewCfnResourceDataSync_Override(c CfnResourceDataSync, scope awscdk.Construct, id *string, props *CfnResourceDataSyncProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ssm.CfnResourceDataSync",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResourceDataSync) SetBucketName(val *string) {
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_CfnResourceDataSync) SetBucketPrefix(val *string) {
	_jsii_.Set(
		j,
		"bucketPrefix",
		val,
	)
}

func (j *jsiiProxy_CfnResourceDataSync) SetBucketRegion(val *string) {
	_jsii_.Set(
		j,
		"bucketRegion",
		val,
	)
}

func (j *jsiiProxy_CfnResourceDataSync) SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_CfnResourceDataSync) SetS3Destination(val interface{}) {
	_jsii_.Set(
		j,
		"s3Destination",
		val,
	)
}

func (j *jsiiProxy_CfnResourceDataSync) SetSyncFormat(val *string) {
	_jsii_.Set(
		j,
		"syncFormat",
		val,
	)
}

func (j *jsiiProxy_CfnResourceDataSync) SetSyncName(val *string) {
	_jsii_.Set(
		j,
		"syncName",
		val,
	)
}

func (j *jsiiProxy_CfnResourceDataSync) SetSyncSource(val interface{}) {
	_jsii_.Set(
		j,
		"syncSource",
		val,
	)
}

func (j *jsiiProxy_CfnResourceDataSync) SetSyncType(val *string) {
	_jsii_.Set(
		j,
		"syncType",
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
func CfnResourceDataSync_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnResourceDataSync",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnResourceDataSync_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnResourceDataSync",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnResourceDataSync_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.CfnResourceDataSync",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourceDataSync_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_ssm.CfnResourceDataSync",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourceDataSync) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResourceDataSync) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResourceDataSync) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResourceDataSync) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResourceDataSync) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResourceDataSync) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResourceDataSync) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResourceDataSync) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDataSync) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDataSync) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResourceDataSync) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResourceDataSync) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnResourceDataSync) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDataSync) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResourceDataSync) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResourceDataSync) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDataSync) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDataSync) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnResourceDataSync) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDataSync) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDataSync) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Information about the `AwsOrganizationsSource` resource data sync source.
//
// A sync source of this type can synchronize data from AWS Organizations or, if an AWS organization isn't present, from multiple AWS Regions .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsOrganizationsSourceProperty := &awsOrganizationsSourceProperty{
//   	organizationSourceType: jsii.String("organizationSourceType"),
//
//   	// the properties below are optional
//   	organizationalUnits: []*string{
//   		jsii.String("organizationalUnits"),
//   	},
//   }
//
type CfnResourceDataSync_AwsOrganizationsSourceProperty struct {
	// If an AWS organization is present, this is either `OrganizationalUnits` or `EntireOrganization` .
	//
	// For `OrganizationalUnits` , the data is aggregated from a set of organization units. For `EntireOrganization` , the data is aggregated from the entire AWS organization.
	OrganizationSourceType *string `field:"required" json:"organizationSourceType" yaml:"organizationSourceType"`
	// The AWS Organizations organization units included in the sync.
	OrganizationalUnits *[]*string `field:"optional" json:"organizationalUnits" yaml:"organizationalUnits"`
}

// Information about the target S3 bucket for the resource data sync.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DestinationProperty := &s3DestinationProperty{
//   	bucketName: jsii.String("bucketName"),
//   	bucketRegion: jsii.String("bucketRegion"),
//   	syncFormat: jsii.String("syncFormat"),
//
//   	// the properties below are optional
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
type CfnResourceDataSync_S3DestinationProperty struct {
	// The name of the S3 bucket where the aggregated data is stored.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The AWS Region with the S3 bucket targeted by the resource data sync.
	BucketRegion *string `field:"required" json:"bucketRegion" yaml:"bucketRegion"`
	// A supported sync format.
	//
	// The following format is currently supported: JsonSerDe.
	SyncFormat *string `field:"required" json:"syncFormat" yaml:"syncFormat"`
	// An Amazon S3 prefix for the bucket.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The ARN of an encryption key for a destination in Amazon S3.
	//
	// Must belong to the same Region as the destination S3 bucket.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

// Information about the source of the data included in the resource data sync.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   syncSourceProperty := &syncSourceProperty{
//   	sourceRegions: []*string{
//   		jsii.String("sourceRegions"),
//   	},
//   	sourceType: jsii.String("sourceType"),
//
//   	// the properties below are optional
//   	awsOrganizationsSource: &awsOrganizationsSourceProperty{
//   		organizationSourceType: jsii.String("organizationSourceType"),
//
//   		// the properties below are optional
//   		organizationalUnits: []*string{
//   			jsii.String("organizationalUnits"),
//   		},
//   	},
//   	includeFutureRegions: jsii.Boolean(false),
//   }
//
type CfnResourceDataSync_SyncSourceProperty struct {
	// The `SyncSource` AWS Regions included in the resource data sync.
	SourceRegions *[]*string `field:"required" json:"sourceRegions" yaml:"sourceRegions"`
	// The type of data source for the resource data sync.
	//
	// `SourceType` is either `AwsOrganizations` (if an organization is present in AWS Organizations ) or `SingleAccountMultiRegions` .
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// Information about the AwsOrganizationsSource resource data sync source.
	//
	// A sync source of this type can synchronize data from AWS Organizations .
	AwsOrganizationsSource interface{} `field:"optional" json:"awsOrganizationsSource" yaml:"awsOrganizationsSource"`
	// Whether to automatically synchronize and aggregate data from new AWS Regions when those Regions come online.
	IncludeFutureRegions interface{} `field:"optional" json:"includeFutureRegions" yaml:"includeFutureRegions"`
}

// Properties for defining a `CfnResourceDataSync`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceDataSyncProps := &cfnResourceDataSyncProps{
//   	syncName: jsii.String("syncName"),
//
//   	// the properties below are optional
//   	bucketName: jsii.String("bucketName"),
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	bucketRegion: jsii.String("bucketRegion"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   	s3Destination: &s3DestinationProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketRegion: jsii.String("bucketRegion"),
//   		syncFormat: jsii.String("syncFormat"),
//
//   		// the properties below are optional
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		kmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	syncFormat: jsii.String("syncFormat"),
//   	syncSource: &syncSourceProperty{
//   		sourceRegions: []*string{
//   			jsii.String("sourceRegions"),
//   		},
//   		sourceType: jsii.String("sourceType"),
//
//   		// the properties below are optional
//   		awsOrganizationsSource: &awsOrganizationsSourceProperty{
//   			organizationSourceType: jsii.String("organizationSourceType"),
//
//   			// the properties below are optional
//   			organizationalUnits: []*string{
//   				jsii.String("organizationalUnits"),
//   			},
//   		},
//   		includeFutureRegions: jsii.Boolean(false),
//   	},
//   	syncType: jsii.String("syncType"),
//   }
//
type CfnResourceDataSyncProps struct {
	// A name for the resource data sync.
	SyncName *string `field:"required" json:"syncName" yaml:"syncName"`
	// The name of the S3 bucket where the aggregated data is stored.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// An Amazon S3 prefix for the bucket.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The AWS Region with the S3 bucket targeted by the resource data sync.
	BucketRegion *string `field:"optional" json:"bucketRegion" yaml:"bucketRegion"`
	// The ARN of an encryption key for a destination in Amazon S3 .
	//
	// You can use a KMS key to encrypt inventory data in Amazon S3 . You must specify a key that exist in the same region as the destination Amazon S3 bucket.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Configuration information for the target S3 bucket.
	S3Destination interface{} `field:"optional" json:"s3Destination" yaml:"s3Destination"`
	// A supported sync format.
	//
	// The following format is currently supported: JsonSerDe.
	SyncFormat *string `field:"optional" json:"syncFormat" yaml:"syncFormat"`
	// Information about the source where the data was synchronized.
	SyncSource interface{} `field:"optional" json:"syncSource" yaml:"syncSource"`
	// The type of resource data sync.
	//
	// If `SyncType` is `SyncToDestination` , then the resource data sync synchronizes data to an S3 bucket. If the `SyncType` is `SyncFromSource` then the resource data sync synchronizes data from AWS Organizations or from multiple AWS Regions .
	SyncType *string `field:"optional" json:"syncType" yaml:"syncType"`
}

// Common attributes for string parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   commonStringParameterAttributes := &commonStringParameterAttributes{
//   	parameterName: jsii.String("parameterName"),
//
//   	// the properties below are optional
//   	simpleName: jsii.Boolean(false),
//   }
//
// Experimental.
type CommonStringParameterAttributes struct {
	// The name of the parameter store value.
	//
	// This value can be a token or a concrete string. If it is a concrete string
	// and includes "/" it must also be prefixed with a "/" (fully-qualified).
	// Experimental.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Indicates of the parameter name is a simple name (i.e. does not include "/" separators).
	//
	// This is only required only if `parameterName` is a token, which means we
	// are unable to detect if the name is simple or "path-like" for the purpose
	// of rendering SSM parameter ARNs.
	//
	// If `parameterName` is not specified, `simpleName` must be `true` (or
	// undefined) since the name generated by AWS CloudFormation is always a
	// simple name.
	// Experimental.
	SimpleName *bool `field:"optional" json:"simpleName" yaml:"simpleName"`
}

// An SSM Parameter reference.
// Experimental.
type IParameter interface {
	awscdk.IResource
	// Grants read (DescribeParameter, GetParameter, GetParameterHistory) permissions on the SSM Parameter.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grants write (PutParameter) permissions on the SSM Parameter.
	// Experimental.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the SSM Parameter resource.
	// Experimental.
	ParameterArn() *string
	// The name of the SSM Parameter resource.
	// Experimental.
	ParameterName() *string
	// The type of the SSM Parameter resource.
	// Experimental.
	ParameterType() *string
}

// The jsii proxy for IParameter
type jsiiProxy_IParameter struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IParameter) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IParameter) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IParameter) ParameterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameter) ParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameter) ParameterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterType",
		&returns,
	)
	return returns
}

// A StringList SSM Parameter.
// Experimental.
type IStringListParameter interface {
	IParameter
	// The parameter value.
	//
	// Value must not nest another parameter. Do not use {{}} in the value. Values in the array
	// cannot contain commas (``,``).
	// Experimental.
	StringListValue() *[]*string
}

// The jsii proxy for IStringListParameter
type jsiiProxy_IStringListParameter struct {
	jsiiProxy_IParameter
}

func (j *jsiiProxy_IStringListParameter) StringListValue() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stringListValue",
		&returns,
	)
	return returns
}

// A String SSM Parameter.
// Experimental.
type IStringParameter interface {
	IParameter
	// The parameter value.
	//
	// Value must not nest another parameter. Do not use {{}} in the value.
	// Experimental.
	StringValue() *string
}

// The jsii proxy for IStringParameter
type jsiiProxy_IStringParameter struct {
	jsiiProxy_IParameter
}

func (j *jsiiProxy_IStringParameter) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

// SSM parameter data type.
// Experimental.
type ParameterDataType string

const (
	// Text.
	// Experimental.
	ParameterDataType_TEXT ParameterDataType = "TEXT"
	// Aws Ec2 Image.
	// Experimental.
	ParameterDataType_AWS_EC2_IMAGE ParameterDataType = "AWS_EC2_IMAGE"
)

// Properties needed to create a new SSM Parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterOptions := &parameterOptions{
//   	allowedPattern: jsii.String("allowedPattern"),
//   	description: jsii.String("description"),
//   	parameterName: jsii.String("parameterName"),
//   	simpleName: jsii.Boolean(false),
//   	tier: awscdk.Aws_ssm.parameterTier_ADVANCED,
//   }
//
// Experimental.
type ParameterOptions struct {
	// A regular expression used to validate the parameter value.
	//
	// For example, for String types with values restricted to
	// numbers, you can specify the following: ``^\d+$``.
	// Experimental.
	AllowedPattern *string `field:"optional" json:"allowedPattern" yaml:"allowedPattern"`
	// Information about the parameter that you want to add to the system.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter.
	// Experimental.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// Indicates of the parameter name is a simple name (i.e. does not include "/" separators).
	//
	// This is only required only if `parameterName` is a token, which means we
	// are unable to detect if the name is simple or "path-like" for the purpose
	// of rendering SSM parameter ARNs.
	//
	// If `parameterName` is not specified, `simpleName` must be `true` (or
	// undefined) since the name generated by AWS CloudFormation is always a
	// simple name.
	// Experimental.
	SimpleName *bool `field:"optional" json:"simpleName" yaml:"simpleName"`
	// The tier of the string parameter.
	// Experimental.
	Tier ParameterTier `field:"optional" json:"tier" yaml:"tier"`
}

// SSM parameter tier.
//
// Example:
//   ssm.NewStringParameter(this, jsii.String("Parameter"), &stringParameterProps{
//   	allowedPattern: jsii.String(".*"),
//   	description: jsii.String("The value Foo"),
//   	parameterName: jsii.String("FooParameter"),
//   	stringValue: jsii.String("Foo"),
//   	tier: ssm.parameterTier_ADVANCED,
//   })
//
// Experimental.
type ParameterTier string

const (
	// String.
	// Experimental.
	ParameterTier_ADVANCED ParameterTier = "ADVANCED"
	// String.
	// Experimental.
	ParameterTier_INTELLIGENT_TIERING ParameterTier = "INTELLIGENT_TIERING"
	// String.
	// Experimental.
	ParameterTier_STANDARD ParameterTier = "STANDARD"
)

// SSM parameter type.
// Experimental.
type ParameterType string

const (
	// String.
	// Experimental.
	ParameterType_STRING ParameterType = "STRING"
	// Secure String.
	//
	// Parameter Store uses an AWS Key Management Service (KMS) customer master key (CMK) to encrypt the parameter value.
	// Parameters of type SecureString cannot be created directly from a CDK application.
	// Experimental.
	ParameterType_SECURE_STRING ParameterType = "SECURE_STRING"
	// String List.
	// Experimental.
	ParameterType_STRING_LIST ParameterType = "STRING_LIST"
	// An Amazon EC2 image ID, such as ami-0ff8a91507f77f867.
	// Experimental.
	ParameterType_AWS_EC2_IMAGE_ID ParameterType = "AWS_EC2_IMAGE_ID"
)

// Attributes for secure string parameters.
//
// Example:
//   // Retrieve the latest value of the non-secret parameter
//   // with name "/My/String/Parameter".
//   stringValue := ssm.stringParameter.fromStringParameterAttributes(this, jsii.String("MyValue"), &stringParameterAttributes{
//   	parameterName: jsii.String("/My/Public/Parameter"),
//   }).stringValue
//   stringValueVersionFromToken := ssm.stringParameter.fromStringParameterAttributes(this, jsii.String("MyValueVersionFromToken"), &stringParameterAttributes{
//   	parameterName: jsii.String("/My/Public/Parameter"),
//   	// parameter version from token
//   	version: parameterVersion,
//   }).stringValue
//
//   // Retrieve a specific version of the secret (SecureString) parameter.
//   // 'version' is always required.
//   secretValue := ssm.stringParameter.fromSecureStringParameterAttributes(this, jsii.String("MySecureValue"), &secureStringParameterAttributes{
//   	parameterName: jsii.String("/My/Secret/Parameter"),
//   	version: jsii.Number(5),
//   })
//   secretValueVersionFromToken := ssm.stringParameter.fromSecureStringParameterAttributes(this, jsii.String("MySecureValueVersionFromToken"), &secureStringParameterAttributes{
//   	parameterName: jsii.String("/My/Secret/Parameter"),
//   	// parameter version from token
//   	version: parameterVersion,
//   })
//
// Experimental.
type SecureStringParameterAttributes struct {
	// The name of the parameter store value.
	//
	// This value can be a token or a concrete string. If it is a concrete string
	// and includes "/" it must also be prefixed with a "/" (fully-qualified).
	// Experimental.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Indicates of the parameter name is a simple name (i.e. does not include "/" separators).
	//
	// This is only required only if `parameterName` is a token, which means we
	// are unable to detect if the name is simple or "path-like" for the purpose
	// of rendering SSM parameter ARNs.
	//
	// If `parameterName` is not specified, `simpleName` must be `true` (or
	// undefined) since the name generated by AWS CloudFormation is always a
	// simple name.
	// Experimental.
	SimpleName *bool `field:"optional" json:"simpleName" yaml:"simpleName"`
	// The encryption key that is used to encrypt this parameter.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The version number of the value you wish to retrieve.
	// Experimental.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

// Creates a new StringList SSM Parameter.
//
// Example:
//   // Create a new SSM Parameter holding a String
//   param := ssm.NewStringParameter(stack, jsii.String("StringParameter"), &stringParameterProps{
//   	// description: 'Some user-friendly description',
//   	// name: 'ParameterName',
//   	stringValue: jsii.String("Initial parameter value"),
//   })
//
//   // Grant read access to some Role
//   param.grantRead(role)
//
//   // Create a new SSM Parameter holding a StringList
//   listParameter := ssm.NewStringListParameter(stack, jsii.String("StringListParameter"), &stringListParameterProps{
//   	// description: 'Some user-friendly description',
//   	// name: 'ParameterName',
//   	stringListValue: []*string{
//   		jsii.String("Initial parameter value A"),
//   		jsii.String("Initial parameter value B"),
//   	},
//   })
//
// Experimental.
type StringListParameter interface {
	awscdk.Resource
	IParameter
	IStringListParameter
	// The encryption key that is used to encrypt this parameter.
	//
	// * @default - default master key.
	// Experimental.
	EncryptionKey() awskms.IKey
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The ARN of the SSM Parameter resource.
	// Experimental.
	ParameterArn() *string
	// The name of the SSM Parameter resource.
	// Experimental.
	ParameterName() *string
	// The type of the SSM Parameter resource.
	// Experimental.
	ParameterType() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The parameter value.
	//
	// Value must not nest another parameter. Do not use {{}} in the value. Values in the array
	// cannot contain commas (``,``).
	// Experimental.
	StringListValue() *[]*string
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
	// Grants read (DescribeParameter, GetParameter, GetParameterHistory) permissions on the SSM Parameter.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grants write (PutParameter) permissions on the SSM Parameter.
	// Experimental.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for StringListParameter
type jsiiProxy_StringListParameter struct {
	internal.Type__awscdkResource
	jsiiProxy_IParameter
	jsiiProxy_IStringListParameter
}

func (j *jsiiProxy_StringListParameter) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringListParameter) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringListParameter) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringListParameter) ParameterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringListParameter) ParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringListParameter) ParameterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringListParameter) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringListParameter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringListParameter) StringListValue() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stringListValue",
		&returns,
	)
	return returns
}


// Experimental.
func NewStringListParameter(scope constructs.Construct, id *string, props *StringListParameterProps) StringListParameter {
	_init_.Initialize()

	j := jsiiProxy_StringListParameter{}

	_jsii_.Create(
		"monocdk.aws_ssm.StringListParameter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStringListParameter_Override(s StringListParameter, scope constructs.Construct, id *string, props *StringListParameterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ssm.StringListParameter",
		[]interface{}{scope, id, props},
		s,
	)
}

// Imports an external parameter of type string list.
//
// Returns a token and should not be parsed.
// Experimental.
func StringListParameter_FromStringListParameterName(scope constructs.Construct, id *string, stringListParameterName *string) IStringListParameter {
	_init_.Initialize()

	var returns IStringListParameter

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.StringListParameter",
		"fromStringListParameterName",
		[]interface{}{scope, id, stringListParameterName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func StringListParameter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.StringListParameter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func StringListParameter_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.StringListParameter",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringListParameter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_StringListParameter) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringListParameter) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringListParameter) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringListParameter) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringListParameter) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringListParameter) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringListParameter) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_StringListParameter) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringListParameter) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringListParameter) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_StringListParameter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringListParameter) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties needed to create a StringList SSM Parameter.
//
// Example:
//   // Create a new SSM Parameter holding a String
//   param := ssm.NewStringParameter(stack, jsii.String("StringParameter"), &stringParameterProps{
//   	// description: 'Some user-friendly description',
//   	// name: 'ParameterName',
//   	stringValue: jsii.String("Initial parameter value"),
//   })
//
//   // Grant read access to some Role
//   param.grantRead(role)
//
//   // Create a new SSM Parameter holding a StringList
//   listParameter := ssm.NewStringListParameter(stack, jsii.String("StringListParameter"), &stringListParameterProps{
//   	// description: 'Some user-friendly description',
//   	// name: 'ParameterName',
//   	stringListValue: []*string{
//   		jsii.String("Initial parameter value A"),
//   		jsii.String("Initial parameter value B"),
//   	},
//   })
//
// Experimental.
type StringListParameterProps struct {
	// A regular expression used to validate the parameter value.
	//
	// For example, for String types with values restricted to
	// numbers, you can specify the following: ``^\d+$``.
	// Experimental.
	AllowedPattern *string `field:"optional" json:"allowedPattern" yaml:"allowedPattern"`
	// Information about the parameter that you want to add to the system.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter.
	// Experimental.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// Indicates of the parameter name is a simple name (i.e. does not include "/" separators).
	//
	// This is only required only if `parameterName` is a token, which means we
	// are unable to detect if the name is simple or "path-like" for the purpose
	// of rendering SSM parameter ARNs.
	//
	// If `parameterName` is not specified, `simpleName` must be `true` (or
	// undefined) since the name generated by AWS CloudFormation is always a
	// simple name.
	// Experimental.
	SimpleName *bool `field:"optional" json:"simpleName" yaml:"simpleName"`
	// The tier of the string parameter.
	// Experimental.
	Tier ParameterTier `field:"optional" json:"tier" yaml:"tier"`
	// The values of the parameter.
	//
	// It may not reference another parameter and ``{{}}`` cannot be used in the value.
	// Experimental.
	StringListValue *[]*string `field:"required" json:"stringListValue" yaml:"stringListValue"`
}

// Creates a new String SSM Parameter.
//
// Example:
//   vpc := ec2.vpc.fromVpcAttributes(this, jsii.String("VPC"), &vpcAttributes{
//   	vpcId: jsii.String("vpc-1234"),
//   	availabilityZones: []*string{
//   		jsii.String("us-east-1a"),
//   		jsii.String("us-east-1b"),
//   	},
//
//   	// Either pass literals for all IDs
//   	publicSubnetIds: []*string{
//   		jsii.String("s-12345"),
//   		jsii.String("s-67890"),
//   	},
//
//   	// OR: import a list of known length
//   	privateSubnetIds: awscdk.Fn.importListValue(jsii.String("PrivateSubnetIds"), jsii.Number(2)),
//
//   	// OR: split an imported string to a list of known length
//   	isolatedSubnetIds: awscdk.Fn.split(jsii.String(","), ssm.stringParameter.valueForStringParameter(this, jsii.String("MyParameter")), jsii.Number(2)),
//   })
//
// Experimental.
type StringParameter interface {
	awscdk.Resource
	IParameter
	IStringParameter
	// The encryption key that is used to encrypt this parameter.
	//
	// * @default - default master key.
	// Experimental.
	EncryptionKey() awskms.IKey
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The ARN of the SSM Parameter resource.
	// Experimental.
	ParameterArn() *string
	// The name of the SSM Parameter resource.
	// Experimental.
	ParameterName() *string
	// The type of the SSM Parameter resource.
	// Experimental.
	ParameterType() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The parameter value.
	//
	// Value must not nest another parameter. Do not use {{}} in the value.
	// Experimental.
	StringValue() *string
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
	// Grants read (DescribeParameter, GetParameter, GetParameterHistory) permissions on the SSM Parameter.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grants write (PutParameter) permissions on the SSM Parameter.
	// Experimental.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for StringParameter
type jsiiProxy_StringParameter struct {
	internal.Type__awscdkResource
	jsiiProxy_IParameter
	jsiiProxy_IStringParameter
}

func (j *jsiiProxy_StringParameter) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) ParameterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) ParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) ParameterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}


// Experimental.
func NewStringParameter(scope constructs.Construct, id *string, props *StringParameterProps) StringParameter {
	_init_.Initialize()

	j := jsiiProxy_StringParameter{}

	_jsii_.Create(
		"monocdk.aws_ssm.StringParameter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStringParameter_Override(s StringParameter, scope constructs.Construct, id *string, props *StringParameterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ssm.StringParameter",
		[]interface{}{scope, id, props},
		s,
	)
}

// Imports a secure string parameter from the SSM parameter store.
// Experimental.
func StringParameter_FromSecureStringParameterAttributes(scope constructs.Construct, id *string, attrs *SecureStringParameterAttributes) IStringParameter {
	_init_.Initialize()

	var returns IStringParameter

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.StringParameter",
		"fromSecureStringParameterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports an external string parameter with name and optional version.
// Experimental.
func StringParameter_FromStringParameterAttributes(scope constructs.Construct, id *string, attrs *StringParameterAttributes) IStringParameter {
	_init_.Initialize()

	var returns IStringParameter

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.StringParameter",
		"fromStringParameterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports an external string parameter by name.
// Experimental.
func StringParameter_FromStringParameterName(scope constructs.Construct, id *string, stringParameterName *string) IStringParameter {
	_init_.Initialize()

	var returns IStringParameter

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.StringParameter",
		"fromStringParameterName",
		[]interface{}{scope, id, stringParameterName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func StringParameter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.StringParameter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func StringParameter_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.StringParameter",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Returns a token that will resolve (during deployment).
// Deprecated: Use `SecretValue.ssmSecure()` instead, it will correctly type the imported value as a `SecretValue` and allow importing without version.
func StringParameter_ValueForSecureStringParameter(scope constructs.Construct, parameterName *string, version *float64) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.StringParameter",
		"valueForSecureStringParameter",
		[]interface{}{scope, parameterName, version},
		&returns,
	)

	return returns
}

// Returns a token that will resolve (during deployment) to the string value of an SSM string parameter.
// Experimental.
func StringParameter_ValueForStringParameter(scope constructs.Construct, parameterName *string, version *float64) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.StringParameter",
		"valueForStringParameter",
		[]interface{}{scope, parameterName, version},
		&returns,
	)

	return returns
}

// Returns a token that will resolve (during deployment) to the string value of an SSM string parameter.
// Experimental.
func StringParameter_ValueForTypedStringParameter(scope constructs.Construct, parameterName *string, type_ ParameterType, version *float64) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.StringParameter",
		"valueForTypedStringParameter",
		[]interface{}{scope, parameterName, type_, version},
		&returns,
	)

	return returns
}

// Reads the value of an SSM parameter during synthesis through an environmental context provider.
//
// Requires that the stack this scope is defined in will have explicit
// account/region information. Otherwise, it will fail during synthesis.
// Experimental.
func StringParameter_ValueFromLookup(scope awscdk.Construct, parameterName *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.aws_ssm.StringParameter",
		"valueFromLookup",
		[]interface{}{scope, parameterName},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringParameter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_StringParameter) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringParameter) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringParameter) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringParameter) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringParameter) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringParameter) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringParameter) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_StringParameter) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringParameter) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringParameter) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_StringParameter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringParameter) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes for parameters of various types of string.
//
// Example:
//   // Retrieve the latest value of the non-secret parameter
//   // with name "/My/String/Parameter".
//   stringValue := ssm.stringParameter.fromStringParameterAttributes(this, jsii.String("MyValue"), &stringParameterAttributes{
//   	parameterName: jsii.String("/My/Public/Parameter"),
//   }).stringValue
//   stringValueVersionFromToken := ssm.stringParameter.fromStringParameterAttributes(this, jsii.String("MyValueVersionFromToken"), &stringParameterAttributes{
//   	parameterName: jsii.String("/My/Public/Parameter"),
//   	// parameter version from token
//   	version: parameterVersion,
//   }).stringValue
//
//   // Retrieve a specific version of the secret (SecureString) parameter.
//   // 'version' is always required.
//   secretValue := ssm.stringParameter.fromSecureStringParameterAttributes(this, jsii.String("MySecureValue"), &secureStringParameterAttributes{
//   	parameterName: jsii.String("/My/Secret/Parameter"),
//   	version: jsii.Number(5),
//   })
//   secretValueVersionFromToken := ssm.stringParameter.fromSecureStringParameterAttributes(this, jsii.String("MySecureValueVersionFromToken"), &secureStringParameterAttributes{
//   	parameterName: jsii.String("/My/Secret/Parameter"),
//   	// parameter version from token
//   	version: parameterVersion,
//   })
//
// See: ParameterType.
//
// Experimental.
type StringParameterAttributes struct {
	// The name of the parameter store value.
	//
	// This value can be a token or a concrete string. If it is a concrete string
	// and includes "/" it must also be prefixed with a "/" (fully-qualified).
	// Experimental.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Indicates of the parameter name is a simple name (i.e. does not include "/" separators).
	//
	// This is only required only if `parameterName` is a token, which means we
	// are unable to detect if the name is simple or "path-like" for the purpose
	// of rendering SSM parameter ARNs.
	//
	// If `parameterName` is not specified, `simpleName` must be `true` (or
	// undefined) since the name generated by AWS CloudFormation is always a
	// simple name.
	// Experimental.
	SimpleName *bool `field:"optional" json:"simpleName" yaml:"simpleName"`
	// The type of the string parameter.
	// Experimental.
	Type ParameterType `field:"optional" json:"type" yaml:"type"`
	// The version number of the value you wish to retrieve.
	// Experimental.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

// Properties needed to create a String SSM parameter.
//
// Example:
//   ssm.NewStringParameter(this, jsii.String("Parameter"), &stringParameterProps{
//   	allowedPattern: jsii.String(".*"),
//   	description: jsii.String("The value Foo"),
//   	parameterName: jsii.String("FooParameter"),
//   	stringValue: jsii.String("Foo"),
//   	tier: ssm.parameterTier_ADVANCED,
//   })
//
// Experimental.
type StringParameterProps struct {
	// A regular expression used to validate the parameter value.
	//
	// For example, for String types with values restricted to
	// numbers, you can specify the following: ``^\d+$``.
	// Experimental.
	AllowedPattern *string `field:"optional" json:"allowedPattern" yaml:"allowedPattern"`
	// Information about the parameter that you want to add to the system.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter.
	// Experimental.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// Indicates of the parameter name is a simple name (i.e. does not include "/" separators).
	//
	// This is only required only if `parameterName` is a token, which means we
	// are unable to detect if the name is simple or "path-like" for the purpose
	// of rendering SSM parameter ARNs.
	//
	// If `parameterName` is not specified, `simpleName` must be `true` (or
	// undefined) since the name generated by AWS CloudFormation is always a
	// simple name.
	// Experimental.
	SimpleName *bool `field:"optional" json:"simpleName" yaml:"simpleName"`
	// The tier of the string parameter.
	// Experimental.
	Tier ParameterTier `field:"optional" json:"tier" yaml:"tier"`
	// The value of the parameter.
	//
	// It may not reference another parameter and ``{{}}`` cannot be used in the value.
	// Experimental.
	StringValue *string `field:"required" json:"stringValue" yaml:"stringValue"`
	// The data type of the parameter, such as `text` or `aws:ec2:image`.
	// Experimental.
	DataType ParameterDataType `field:"optional" json:"dataType" yaml:"dataType"`
	// The type of the string parameter.
	// Experimental.
	Type ParameterType `field:"optional" json:"type" yaml:"type"`
}

