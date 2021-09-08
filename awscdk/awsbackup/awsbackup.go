package awsbackup

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsbackup/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsefs"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/constructs-go/constructs/v3"
)

// A backup plan.
// Experimental.
type BackupPlan interface {
	awscdk.Resource
	IBackupPlan
	BackupPlanArn() *string
	BackupPlanId() *string
	BackupVault() IBackupVault
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	VersionId() *string
	AddRule(rule BackupPlanRule)
	AddSelection(id *string, options *BackupSelectionOptions) BackupSelection
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for BackupPlan
type jsiiProxy_BackupPlan struct {
	internal.Type__awscdkResource
	jsiiProxy_IBackupPlan
}

func (j *jsiiProxy_BackupPlan) BackupPlanArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupPlanArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlan) BackupPlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupPlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlan) BackupVault() IBackupVault {
	var returns IBackupVault
	_jsii_.Get(
		j,
		"backupVault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlan) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlan) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlan) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlan) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlan) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}


// Experimental.
func NewBackupPlan(scope constructs.Construct, id *string, props *BackupPlanProps) BackupPlan {
	_init_.Initialize()

	j := jsiiProxy_BackupPlan{}

	_jsii_.Create(
		"monocdk.aws_backup.BackupPlan",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBackupPlan_Override(b BackupPlan, scope constructs.Construct, id *string, props *BackupPlanProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_backup.BackupPlan",
		[]interface{}{scope, id, props},
		b,
	)
}

// Daily with 35 day retention.
// Experimental.
func BackupPlan_Daily35DayRetention(scope constructs.Construct, id *string, backupVault IBackupVault) BackupPlan {
	_init_.Initialize()

	var returns BackupPlan

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlan",
		"daily35DayRetention",
		[]interface{}{scope, id, backupVault},
		&returns,
	)

	return returns
}

// Daily and monthly with 1 year retention.
// Experimental.
func BackupPlan_DailyMonthly1YearRetention(scope constructs.Construct, id *string, backupVault IBackupVault) BackupPlan {
	_init_.Initialize()

	var returns BackupPlan

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlan",
		"dailyMonthly1YearRetention",
		[]interface{}{scope, id, backupVault},
		&returns,
	)

	return returns
}

// Daily, weekly and monthly with 5 year retention.
// Experimental.
func BackupPlan_DailyWeeklyMonthly5YearRetention(scope constructs.Construct, id *string, backupVault IBackupVault) BackupPlan {
	_init_.Initialize()

	var returns BackupPlan

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlan",
		"dailyWeeklyMonthly5YearRetention",
		[]interface{}{scope, id, backupVault},
		&returns,
	)

	return returns
}

// Daily, weekly and monthly with 7 year retention.
// Experimental.
func BackupPlan_DailyWeeklyMonthly7YearRetention(scope constructs.Construct, id *string, backupVault IBackupVault) BackupPlan {
	_init_.Initialize()

	var returns BackupPlan

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlan",
		"dailyWeeklyMonthly7YearRetention",
		[]interface{}{scope, id, backupVault},
		&returns,
	)

	return returns
}

// Import an existing backup plan.
// Experimental.
func BackupPlan_FromBackupPlanId(scope constructs.Construct, id *string, backupPlanId *string) IBackupPlan {
	_init_.Initialize()

	var returns IBackupPlan

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlan",
		"fromBackupPlanId",
		[]interface{}{scope, id, backupPlanId},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func BackupPlan_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlan",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func BackupPlan_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlan",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a rule to a plan.
// Experimental.
func (b *jsiiProxy_BackupPlan) AddRule(rule BackupPlanRule) {
	_jsii_.InvokeVoid(
		b,
		"addRule",
		[]interface{}{rule},
	)
}

// Adds a selection to this plan.
// Experimental.
func (b *jsiiProxy_BackupPlan) AddSelection(id *string, options *BackupSelectionOptions) BackupSelection {
	var returns BackupSelection

	_jsii_.Invoke(
		b,
		"addSelection",
		[]interface{}{id, options},
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (b *jsiiProxy_BackupPlan) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (b *jsiiProxy_BackupPlan) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_BackupPlan) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_BackupPlan) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
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
func (b *jsiiProxy_BackupPlan) OnPrepare() {
	_jsii_.InvokeVoid(
		b,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (b *jsiiProxy_BackupPlan) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
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
func (b *jsiiProxy_BackupPlan) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (b *jsiiProxy_BackupPlan) Prepare() {
	_jsii_.InvokeVoid(
		b,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (b *jsiiProxy_BackupPlan) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (b *jsiiProxy_BackupPlan) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
// Experimental.
func (b *jsiiProxy_BackupPlan) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a BackupPlan.
// Experimental.
type BackupPlanProps struct {
	// The display name of the backup plan.
	// Experimental.
	BackupPlanName *string `json:"backupPlanName"`
	// Rules for the backup plan.
	//
	// Use `addRule()` to add rules after
	// instantiation.
	// Experimental.
	BackupPlanRules *[]BackupPlanRule `json:"backupPlanRules"`
	// The backup vault where backups are stored.
	// Experimental.
	BackupVault IBackupVault `json:"backupVault"`
}

// A backup plan rule.
// Experimental.
type BackupPlanRule interface {
	Props() *BackupPlanRuleProps
}

// The jsii proxy struct for BackupPlanRule
type jsiiProxy_BackupPlanRule struct {
	_ byte // padding
}

func (j *jsiiProxy_BackupPlanRule) Props() *BackupPlanRuleProps {
	var returns *BackupPlanRuleProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewBackupPlanRule(props *BackupPlanRuleProps) BackupPlanRule {
	_init_.Initialize()

	j := jsiiProxy_BackupPlanRule{}

	_jsii_.Create(
		"monocdk.aws_backup.BackupPlanRule",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewBackupPlanRule_Override(b BackupPlanRule, props *BackupPlanRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_backup.BackupPlanRule",
		[]interface{}{props},
		b,
	)
}

// Daily with 35 days retention.
// Experimental.
func BackupPlanRule_Daily(backupVault IBackupVault) BackupPlanRule {
	_init_.Initialize()

	var returns BackupPlanRule

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlanRule",
		"daily",
		[]interface{}{backupVault},
		&returns,
	)

	return returns
}

// Monthly 1 year retention, move to cold storage after 1 month.
// Experimental.
func BackupPlanRule_Monthly1Year(backupVault IBackupVault) BackupPlanRule {
	_init_.Initialize()

	var returns BackupPlanRule

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlanRule",
		"monthly1Year",
		[]interface{}{backupVault},
		&returns,
	)

	return returns
}

// Monthly 5 year retention, move to cold storage after 3 months.
// Experimental.
func BackupPlanRule_Monthly5Year(backupVault IBackupVault) BackupPlanRule {
	_init_.Initialize()

	var returns BackupPlanRule

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlanRule",
		"monthly5Year",
		[]interface{}{backupVault},
		&returns,
	)

	return returns
}

// Monthly 7 year retention, move to cold storage after 3 months.
// Experimental.
func BackupPlanRule_Monthly7Year(backupVault IBackupVault) BackupPlanRule {
	_init_.Initialize()

	var returns BackupPlanRule

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlanRule",
		"monthly7Year",
		[]interface{}{backupVault},
		&returns,
	)

	return returns
}

// Weekly with 3 months retention.
// Experimental.
func BackupPlanRule_Weekly(backupVault IBackupVault) BackupPlanRule {
	_init_.Initialize()

	var returns BackupPlanRule

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlanRule",
		"weekly",
		[]interface{}{backupVault},
		&returns,
	)

	return returns
}

// Properties for a BackupPlanRule.
// Experimental.
type BackupPlanRuleProps struct {
	// The backup vault where backups are.
	// Experimental.
	BackupVault IBackupVault `json:"backupVault"`
	// The duration after a backup job is successfully started before it must be completed or it is canceled by AWS Backup.
	// Experimental.
	CompletionWindow awscdk.Duration `json:"completionWindow"`
	// Specifies the duration after creation that a recovery point is deleted.
	//
	// Must be greater than `moveToColdStorageAfter`.
	// Experimental.
	DeleteAfter awscdk.Duration `json:"deleteAfter"`
	// Specifies the duration after creation that a recovery point is moved to cold storage.
	// Experimental.
	MoveToColdStorageAfter awscdk.Duration `json:"moveToColdStorageAfter"`
	// A display name for the backup rule.
	// Experimental.
	RuleName *string `json:"ruleName"`
	// A CRON expression specifying when AWS Backup initiates a backup job.
	// Experimental.
	ScheduleExpression awsevents.Schedule `json:"scheduleExpression"`
	// The duration after a backup is scheduled before a job is canceled if it doesn't start successfully.
	// Experimental.
	StartWindow awscdk.Duration `json:"startWindow"`
}

// A resource to backup.
// Experimental.
type BackupResource interface {
	Construct() awscdk.Construct
	Resource() *string
	TagCondition() *TagCondition
}

// The jsii proxy struct for BackupResource
type jsiiProxy_BackupResource struct {
	_ byte // padding
}

func (j *jsiiProxy_BackupResource) Construct() awscdk.Construct {
	var returns awscdk.Construct
	_jsii_.Get(
		j,
		"construct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupResource) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupResource) TagCondition() *TagCondition {
	var returns *TagCondition
	_jsii_.Get(
		j,
		"tagCondition",
		&returns,
	)
	return returns
}


// Experimental.
func NewBackupResource(resource *string, tagCondition *TagCondition, construct constructs.Construct) BackupResource {
	_init_.Initialize()

	j := jsiiProxy_BackupResource{}

	_jsii_.Create(
		"monocdk.aws_backup.BackupResource",
		[]interface{}{resource, tagCondition, construct},
		&j,
	)

	return &j
}

// Experimental.
func NewBackupResource_Override(b BackupResource, resource *string, tagCondition *TagCondition, construct constructs.Construct) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_backup.BackupResource",
		[]interface{}{resource, tagCondition, construct},
		b,
	)
}

// A list of ARNs or match patterns such as `arn:aws:ec2:us-east-1:123456789012:volume/*`.
// Experimental.
func BackupResource_FromArn(arn *string) BackupResource {
	_init_.Initialize()

	var returns BackupResource

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupResource",
		"fromArn",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

// Adds all supported resources in a construct.
// Experimental.
func BackupResource_FromConstruct(construct constructs.Construct) BackupResource {
	_init_.Initialize()

	var returns BackupResource

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupResource",
		"fromConstruct",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// A DynamoDB table.
// Experimental.
func BackupResource_FromDynamoDbTable(table awsdynamodb.ITable) BackupResource {
	_init_.Initialize()

	var returns BackupResource

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupResource",
		"fromDynamoDbTable",
		[]interface{}{table},
		&returns,
	)

	return returns
}

// An EC2 instance.
// Experimental.
func BackupResource_FromEc2Instance(instance awsec2.IInstance) BackupResource {
	_init_.Initialize()

	var returns BackupResource

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupResource",
		"fromEc2Instance",
		[]interface{}{instance},
		&returns,
	)

	return returns
}

// An EFS file system.
// Experimental.
func BackupResource_FromEfsFileSystem(fileSystem awsefs.IFileSystem) BackupResource {
	_init_.Initialize()

	var returns BackupResource

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupResource",
		"fromEfsFileSystem",
		[]interface{}{fileSystem},
		&returns,
	)

	return returns
}

// A RDS database instance.
// Experimental.
func BackupResource_FromRdsDatabaseInstance(instance awsrds.IDatabaseInstance) BackupResource {
	_init_.Initialize()

	var returns BackupResource

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupResource",
		"fromRdsDatabaseInstance",
		[]interface{}{instance},
		&returns,
	)

	return returns
}

// A tag condition.
// Experimental.
func BackupResource_FromTag(key *string, value *string, operation TagOperation) BackupResource {
	_init_.Initialize()

	var returns BackupResource

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupResource",
		"fromTag",
		[]interface{}{key, value, operation},
		&returns,
	)

	return returns
}

// A backup selection.
// Experimental.
type BackupSelection interface {
	awscdk.Resource
	awsiam.IGrantable
	BackupPlanId() *string
	Env() *awscdk.ResourceEnvironment
	GrantPrincipal() awsiam.IPrincipal
	Node() awscdk.ConstructNode
	PhysicalName() *string
	SelectionId() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for BackupSelection
type jsiiProxy_BackupSelection struct {
	internal.Type__awscdkResource
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_BackupSelection) BackupPlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupPlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelection) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelection) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelection) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelection) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelection) SelectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelection) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewBackupSelection(scope constructs.Construct, id *string, props *BackupSelectionProps) BackupSelection {
	_init_.Initialize()

	j := jsiiProxy_BackupSelection{}

	_jsii_.Create(
		"monocdk.aws_backup.BackupSelection",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBackupSelection_Override(b BackupSelection, scope constructs.Construct, id *string, props *BackupSelectionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_backup.BackupSelection",
		[]interface{}{scope, id, props},
		b,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func BackupSelection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupSelection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func BackupSelection_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupSelection",
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (b *jsiiProxy_BackupSelection) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (b *jsiiProxy_BackupSelection) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_BackupSelection) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_BackupSelection) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
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
func (b *jsiiProxy_BackupSelection) OnPrepare() {
	_jsii_.InvokeVoid(
		b,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (b *jsiiProxy_BackupSelection) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
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
func (b *jsiiProxy_BackupSelection) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (b *jsiiProxy_BackupSelection) Prepare() {
	_jsii_.InvokeVoid(
		b,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (b *jsiiProxy_BackupSelection) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (b *jsiiProxy_BackupSelection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_BackupSelection) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for a BackupSelection.
// Experimental.
type BackupSelectionOptions struct {
	// The resources to backup.
	//
	// Use the helper static methods defined on `BackupResource`.
	// Experimental.
	Resources *[]BackupResource `json:"resources"`
	// Whether to automatically give restores permissions to the role that AWS Backup uses.
	//
	// If `true`, the `AWSBackupServiceRolePolicyForRestores` managed
	// policy will be attached to the role.
	// Experimental.
	AllowRestores *bool `json:"allowRestores"`
	// The name for this selection.
	// Experimental.
	BackupSelectionName *string `json:"backupSelectionName"`
	// The role that AWS Backup uses to authenticate when backuping or restoring the resources.
	//
	// The `AWSBackupServiceRolePolicyForBackup` managed policy
	// will be attached to this role.
	// Experimental.
	Role awsiam.IRole `json:"role"`
}

// Properties for a BackupSelection.
// Experimental.
type BackupSelectionProps struct {
	// The resources to backup.
	//
	// Use the helper static methods defined on `BackupResource`.
	// Experimental.
	Resources *[]BackupResource `json:"resources"`
	// Whether to automatically give restores permissions to the role that AWS Backup uses.
	//
	// If `true`, the `AWSBackupServiceRolePolicyForRestores` managed
	// policy will be attached to the role.
	// Experimental.
	AllowRestores *bool `json:"allowRestores"`
	// The name for this selection.
	// Experimental.
	BackupSelectionName *string `json:"backupSelectionName"`
	// The role that AWS Backup uses to authenticate when backuping or restoring the resources.
	//
	// The `AWSBackupServiceRolePolicyForBackup` managed policy
	// will be attached to this role.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The backup plan for this selection.
	// Experimental.
	BackupPlan IBackupPlan `json:"backupPlan"`
}

// A backup vault.
// Experimental.
type BackupVault interface {
	awscdk.Resource
	IBackupVault
	BackupVaultArn() *string
	BackupVaultName() *string
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for BackupVault
type jsiiProxy_BackupVault struct {
	internal.Type__awscdkResource
	jsiiProxy_IBackupVault
}

func (j *jsiiProxy_BackupVault) BackupVaultArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupVault) BackupVaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupVault) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupVault) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupVault) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupVault) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewBackupVault(scope constructs.Construct, id *string, props *BackupVaultProps) BackupVault {
	_init_.Initialize()

	j := jsiiProxy_BackupVault{}

	_jsii_.Create(
		"monocdk.aws_backup.BackupVault",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBackupVault_Override(b BackupVault, scope constructs.Construct, id *string, props *BackupVaultProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_backup.BackupVault",
		[]interface{}{scope, id, props},
		b,
	)
}

// Import an existing backup vault by arn.
// Experimental.
func BackupVault_FromBackupVaultArn(scope constructs.Construct, id *string, backupVaultArn *string) IBackupVault {
	_init_.Initialize()

	var returns IBackupVault

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupVault",
		"fromBackupVaultArn",
		[]interface{}{scope, id, backupVaultArn},
		&returns,
	)

	return returns
}

// Import an existing backup vault by name.
// Experimental.
func BackupVault_FromBackupVaultName(scope constructs.Construct, id *string, backupVaultName *string) IBackupVault {
	_init_.Initialize()

	var returns IBackupVault

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupVault",
		"fromBackupVaultName",
		[]interface{}{scope, id, backupVaultName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func BackupVault_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupVault",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func BackupVault_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupVault",
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (b *jsiiProxy_BackupVault) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (b *jsiiProxy_BackupVault) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_BackupVault) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_BackupVault) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the actions defined in actions to the given grantee on this Backup Vault resource.
// Experimental.
func (b *jsiiProxy_BackupVault) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grant",
		args,
		&returns,
	)

	return returns
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
func (b *jsiiProxy_BackupVault) OnPrepare() {
	_jsii_.InvokeVoid(
		b,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (b *jsiiProxy_BackupVault) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
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
func (b *jsiiProxy_BackupVault) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (b *jsiiProxy_BackupVault) Prepare() {
	_jsii_.InvokeVoid(
		b,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (b *jsiiProxy_BackupVault) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (b *jsiiProxy_BackupVault) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_BackupVault) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Backup vault events.
// Experimental.
type BackupVaultEvents string

const (
	BackupVaultEvents_BACKUP_JOB_STARTED BackupVaultEvents = "BACKUP_JOB_STARTED"
	BackupVaultEvents_BACKUP_JOB_COMPLETED BackupVaultEvents = "BACKUP_JOB_COMPLETED"
	BackupVaultEvents_BACKUP_JOB_SUCCESSFUL BackupVaultEvents = "BACKUP_JOB_SUCCESSFUL"
	BackupVaultEvents_BACKUP_JOB_FAILED BackupVaultEvents = "BACKUP_JOB_FAILED"
	BackupVaultEvents_BACKUP_JOB_EXPIRED BackupVaultEvents = "BACKUP_JOB_EXPIRED"
	BackupVaultEvents_RESTORE_JOB_STARTED BackupVaultEvents = "RESTORE_JOB_STARTED"
	BackupVaultEvents_RESTORE_JOB_COMPLETED BackupVaultEvents = "RESTORE_JOB_COMPLETED"
	BackupVaultEvents_RESTORE_JOB_SUCCESSFUL BackupVaultEvents = "RESTORE_JOB_SUCCESSFUL"
	BackupVaultEvents_RESTORE_JOB_FAILED BackupVaultEvents = "RESTORE_JOB_FAILED"
	BackupVaultEvents_COPY_JOB_STARTED BackupVaultEvents = "COPY_JOB_STARTED"
	BackupVaultEvents_COPY_JOB_SUCCESSFUL BackupVaultEvents = "COPY_JOB_SUCCESSFUL"
	BackupVaultEvents_COPY_JOB_FAILED BackupVaultEvents = "COPY_JOB_FAILED"
	BackupVaultEvents_RECOVERY_POINT_MODIFIED BackupVaultEvents = "RECOVERY_POINT_MODIFIED"
	BackupVaultEvents_BACKUP_PLAN_CREATED BackupVaultEvents = "BACKUP_PLAN_CREATED"
	BackupVaultEvents_BACKUP_PLAN_MODIFIED BackupVaultEvents = "BACKUP_PLAN_MODIFIED"
)

// Properties for a BackupVault.
// Experimental.
type BackupVaultProps struct {
	// A resource-based policy that is used to manage access permissions on the backup vault.
	// Experimental.
	AccessPolicy awsiam.PolicyDocument `json:"accessPolicy"`
	// The name of a logical container where backups are stored.
	//
	// Backup vaults
	// are identified by names that are unique to the account used to create
	// them and the AWS Region where they are created.
	// Experimental.
	BackupVaultName *string `json:"backupVaultName"`
	// Whether to add statements to the vault access policy that prevents anyone from deleting a recovery point.
	// Experimental.
	BlockRecoveryPointDeletion *bool `json:"blockRecoveryPointDeletion"`
	// The server-side encryption key to use to protect your backups.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// The vault events to send.
	// See: https://docs.aws.amazon.com/aws-backup/latest/devguide/sns-notifications.html
	//
	// Experimental.
	NotificationEvents *[]BackupVaultEvents `json:"notificationEvents"`
	// A SNS topic to send vault events to.
	// See: https://docs.aws.amazon.com/aws-backup/latest/devguide/sns-notifications.html
	//
	// Experimental.
	NotificationTopic awssns.ITopic `json:"notificationTopic"`
	// The removal policy to apply to the vault.
	//
	// Note that removing a vault
	// that contains recovery points will fail.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
}

// A CloudFormation `AWS::Backup::BackupPlan`.
type CfnBackupPlan interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrBackupPlanArn() *string
	AttrBackupPlanId() *string
	AttrVersionId() *string
	BackupPlan() interface{}
	SetBackupPlan(val interface{})
	BackupPlanTags() interface{}
	SetBackupPlanTags(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
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

// The jsii proxy struct for CfnBackupPlan
type jsiiProxy_CfnBackupPlan struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBackupPlan) AttrBackupPlanArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrBackupPlanArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupPlan) AttrBackupPlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrBackupPlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupPlan) AttrVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupPlan) BackupPlan() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupPlan) BackupPlanTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupPlanTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupPlan) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupPlan) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupPlan) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupPlan) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupPlan) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupPlan) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupPlan) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupPlan) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupPlan) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Backup::BackupPlan`.
func NewCfnBackupPlan(scope awscdk.Construct, id *string, props *CfnBackupPlanProps) CfnBackupPlan {
	_init_.Initialize()

	j := jsiiProxy_CfnBackupPlan{}

	_jsii_.Create(
		"monocdk.aws_backup.CfnBackupPlan",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Backup::BackupPlan`.
func NewCfnBackupPlan_Override(c CfnBackupPlan, scope awscdk.Construct, id *string, props *CfnBackupPlanProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_backup.CfnBackupPlan",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBackupPlan) SetBackupPlan(val interface{}) {
	_jsii_.Set(
		j,
		"backupPlan",
		val,
	)
}

func (j *jsiiProxy_CfnBackupPlan) SetBackupPlanTags(val interface{}) {
	_jsii_.Set(
		j,
		"backupPlanTags",
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
func CfnBackupPlan_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.CfnBackupPlan",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnBackupPlan_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.CfnBackupPlan",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnBackupPlan_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.CfnBackupPlan",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBackupPlan_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_backup.CfnBackupPlan",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnBackupPlan) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnBackupPlan) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnBackupPlan) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnBackupPlan) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnBackupPlan) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnBackupPlan) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnBackupPlan) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnBackupPlan) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnBackupPlan) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnBackupPlan) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnBackupPlan) OnPrepare() {
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
func (c *jsiiProxy_CfnBackupPlan) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnBackupPlan) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnBackupPlan) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnBackupPlan) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBackupPlan) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnBackupPlan) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnBackupPlan) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnBackupPlan) ToString() *string {
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
func (c *jsiiProxy_CfnBackupPlan) Validate() *[]*string {
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
func (c *jsiiProxy_CfnBackupPlan) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnBackupPlan_AdvancedBackupSettingResourceTypeProperty struct {
	// `CfnBackupPlan.AdvancedBackupSettingResourceTypeProperty.BackupOptions`.
	BackupOptions interface{} `json:"backupOptions"`
	// `CfnBackupPlan.AdvancedBackupSettingResourceTypeProperty.ResourceType`.
	ResourceType *string `json:"resourceType"`
}

type CfnBackupPlan_BackupPlanResourceTypeProperty struct {
	// `CfnBackupPlan.BackupPlanResourceTypeProperty.BackupPlanName`.
	BackupPlanName *string `json:"backupPlanName"`
	// `CfnBackupPlan.BackupPlanResourceTypeProperty.BackupPlanRule`.
	BackupPlanRule interface{} `json:"backupPlanRule"`
	// `CfnBackupPlan.BackupPlanResourceTypeProperty.AdvancedBackupSettings`.
	AdvancedBackupSettings interface{} `json:"advancedBackupSettings"`
}

type CfnBackupPlan_BackupRuleResourceTypeProperty struct {
	// `CfnBackupPlan.BackupRuleResourceTypeProperty.RuleName`.
	RuleName *string `json:"ruleName"`
	// `CfnBackupPlan.BackupRuleResourceTypeProperty.TargetBackupVault`.
	TargetBackupVault *string `json:"targetBackupVault"`
	// `CfnBackupPlan.BackupRuleResourceTypeProperty.CompletionWindowMinutes`.
	CompletionWindowMinutes *float64 `json:"completionWindowMinutes"`
	// `CfnBackupPlan.BackupRuleResourceTypeProperty.CopyActions`.
	CopyActions interface{} `json:"copyActions"`
	// `CfnBackupPlan.BackupRuleResourceTypeProperty.EnableContinuousBackup`.
	EnableContinuousBackup interface{} `json:"enableContinuousBackup"`
	// `CfnBackupPlan.BackupRuleResourceTypeProperty.Lifecycle`.
	Lifecycle interface{} `json:"lifecycle"`
	// `CfnBackupPlan.BackupRuleResourceTypeProperty.RecoveryPointTags`.
	RecoveryPointTags interface{} `json:"recoveryPointTags"`
	// `CfnBackupPlan.BackupRuleResourceTypeProperty.ScheduleExpression`.
	ScheduleExpression *string `json:"scheduleExpression"`
	// `CfnBackupPlan.BackupRuleResourceTypeProperty.StartWindowMinutes`.
	StartWindowMinutes *float64 `json:"startWindowMinutes"`
}

type CfnBackupPlan_CopyActionResourceTypeProperty struct {
	// `CfnBackupPlan.CopyActionResourceTypeProperty.DestinationBackupVaultArn`.
	DestinationBackupVaultArn *string `json:"destinationBackupVaultArn"`
	// `CfnBackupPlan.CopyActionResourceTypeProperty.Lifecycle`.
	Lifecycle interface{} `json:"lifecycle"`
}

type CfnBackupPlan_LifecycleResourceTypeProperty struct {
	// `CfnBackupPlan.LifecycleResourceTypeProperty.DeleteAfterDays`.
	DeleteAfterDays *float64 `json:"deleteAfterDays"`
	// `CfnBackupPlan.LifecycleResourceTypeProperty.MoveToColdStorageAfterDays`.
	MoveToColdStorageAfterDays *float64 `json:"moveToColdStorageAfterDays"`
}

// Properties for defining a `AWS::Backup::BackupPlan`.
type CfnBackupPlanProps struct {
	// `AWS::Backup::BackupPlan.BackupPlan`.
	BackupPlan interface{} `json:"backupPlan"`
	// `AWS::Backup::BackupPlan.BackupPlanTags`.
	BackupPlanTags interface{} `json:"backupPlanTags"`
}

// A CloudFormation `AWS::Backup::BackupSelection`.
type CfnBackupSelection interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrBackupPlanId() *string
	AttrId() *string
	AttrSelectionId() *string
	BackupPlanId() *string
	SetBackupPlanId(val *string)
	BackupSelection() interface{}
	SetBackupSelection(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
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

// The jsii proxy struct for CfnBackupSelection
type jsiiProxy_CfnBackupSelection struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBackupSelection) AttrBackupPlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrBackupPlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupSelection) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupSelection) AttrSelectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSelectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupSelection) BackupPlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupPlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupSelection) BackupSelection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupSelection) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupSelection) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupSelection) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupSelection) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupSelection) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupSelection) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupSelection) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupSelection) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupSelection) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Backup::BackupSelection`.
func NewCfnBackupSelection(scope awscdk.Construct, id *string, props *CfnBackupSelectionProps) CfnBackupSelection {
	_init_.Initialize()

	j := jsiiProxy_CfnBackupSelection{}

	_jsii_.Create(
		"monocdk.aws_backup.CfnBackupSelection",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Backup::BackupSelection`.
func NewCfnBackupSelection_Override(c CfnBackupSelection, scope awscdk.Construct, id *string, props *CfnBackupSelectionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_backup.CfnBackupSelection",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBackupSelection) SetBackupPlanId(val *string) {
	_jsii_.Set(
		j,
		"backupPlanId",
		val,
	)
}

func (j *jsiiProxy_CfnBackupSelection) SetBackupSelection(val interface{}) {
	_jsii_.Set(
		j,
		"backupSelection",
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
func CfnBackupSelection_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.CfnBackupSelection",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnBackupSelection_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.CfnBackupSelection",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnBackupSelection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.CfnBackupSelection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBackupSelection_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_backup.CfnBackupSelection",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnBackupSelection) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnBackupSelection) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnBackupSelection) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnBackupSelection) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnBackupSelection) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnBackupSelection) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnBackupSelection) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnBackupSelection) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnBackupSelection) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnBackupSelection) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnBackupSelection) OnPrepare() {
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
func (c *jsiiProxy_CfnBackupSelection) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnBackupSelection) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnBackupSelection) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnBackupSelection) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBackupSelection) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnBackupSelection) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnBackupSelection) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnBackupSelection) ToString() *string {
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
func (c *jsiiProxy_CfnBackupSelection) Validate() *[]*string {
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
func (c *jsiiProxy_CfnBackupSelection) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnBackupSelection_BackupSelectionResourceTypeProperty struct {
	// `CfnBackupSelection.BackupSelectionResourceTypeProperty.IamRoleArn`.
	IamRoleArn *string `json:"iamRoleArn"`
	// `CfnBackupSelection.BackupSelectionResourceTypeProperty.SelectionName`.
	SelectionName *string `json:"selectionName"`
	// `CfnBackupSelection.BackupSelectionResourceTypeProperty.ListOfTags`.
	ListOfTags interface{} `json:"listOfTags"`
	// `CfnBackupSelection.BackupSelectionResourceTypeProperty.Resources`.
	Resources *[]*string `json:"resources"`
}

type CfnBackupSelection_ConditionResourceTypeProperty struct {
	// `CfnBackupSelection.ConditionResourceTypeProperty.ConditionKey`.
	ConditionKey *string `json:"conditionKey"`
	// `CfnBackupSelection.ConditionResourceTypeProperty.ConditionType`.
	ConditionType *string `json:"conditionType"`
	// `CfnBackupSelection.ConditionResourceTypeProperty.ConditionValue`.
	ConditionValue *string `json:"conditionValue"`
}

// Properties for defining a `AWS::Backup::BackupSelection`.
type CfnBackupSelectionProps struct {
	// `AWS::Backup::BackupSelection.BackupPlanId`.
	BackupPlanId *string `json:"backupPlanId"`
	// `AWS::Backup::BackupSelection.BackupSelection`.
	BackupSelection interface{} `json:"backupSelection"`
}

// A CloudFormation `AWS::Backup::BackupVault`.
type CfnBackupVault interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AccessPolicy() interface{}
	SetAccessPolicy(val interface{})
	AttrBackupVaultArn() *string
	AttrBackupVaultName() *string
	BackupVaultName() *string
	SetBackupVaultName(val *string)
	BackupVaultTags() interface{}
	SetBackupVaultTags(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EncryptionKeyArn() *string
	SetEncryptionKeyArn(val *string)
	LogicalId() *string
	Node() awscdk.ConstructNode
	Notifications() interface{}
	SetNotifications(val interface{})
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

// The jsii proxy struct for CfnBackupVault
type jsiiProxy_CfnBackupVault struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBackupVault) AccessPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupVault) AttrBackupVaultArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrBackupVaultArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupVault) AttrBackupVaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrBackupVaultName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupVault) BackupVaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupVault) BackupVaultTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupVaultTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupVault) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupVault) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupVault) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupVault) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupVault) EncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupVault) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupVault) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupVault) Notifications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupVault) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupVault) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupVault) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Backup::BackupVault`.
func NewCfnBackupVault(scope awscdk.Construct, id *string, props *CfnBackupVaultProps) CfnBackupVault {
	_init_.Initialize()

	j := jsiiProxy_CfnBackupVault{}

	_jsii_.Create(
		"monocdk.aws_backup.CfnBackupVault",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Backup::BackupVault`.
func NewCfnBackupVault_Override(c CfnBackupVault, scope awscdk.Construct, id *string, props *CfnBackupVaultProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_backup.CfnBackupVault",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBackupVault) SetAccessPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"accessPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnBackupVault) SetBackupVaultName(val *string) {
	_jsii_.Set(
		j,
		"backupVaultName",
		val,
	)
}

func (j *jsiiProxy_CfnBackupVault) SetBackupVaultTags(val interface{}) {
	_jsii_.Set(
		j,
		"backupVaultTags",
		val,
	)
}

func (j *jsiiProxy_CfnBackupVault) SetEncryptionKeyArn(val *string) {
	_jsii_.Set(
		j,
		"encryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_CfnBackupVault) SetNotifications(val interface{}) {
	_jsii_.Set(
		j,
		"notifications",
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
func CfnBackupVault_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.CfnBackupVault",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnBackupVault_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.CfnBackupVault",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnBackupVault_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.CfnBackupVault",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBackupVault_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_backup.CfnBackupVault",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnBackupVault) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnBackupVault) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnBackupVault) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnBackupVault) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnBackupVault) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnBackupVault) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnBackupVault) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnBackupVault) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnBackupVault) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnBackupVault) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnBackupVault) OnPrepare() {
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
func (c *jsiiProxy_CfnBackupVault) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnBackupVault) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnBackupVault) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnBackupVault) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBackupVault) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnBackupVault) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnBackupVault) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnBackupVault) ToString() *string {
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
func (c *jsiiProxy_CfnBackupVault) Validate() *[]*string {
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
func (c *jsiiProxy_CfnBackupVault) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnBackupVault_NotificationObjectTypeProperty struct {
	// `CfnBackupVault.NotificationObjectTypeProperty.BackupVaultEvents`.
	BackupVaultEvents *[]*string `json:"backupVaultEvents"`
	// `CfnBackupVault.NotificationObjectTypeProperty.SNSTopicArn`.
	SnsTopicArn *string `json:"snsTopicArn"`
}

// Properties for defining a `AWS::Backup::BackupVault`.
type CfnBackupVaultProps struct {
	// `AWS::Backup::BackupVault.BackupVaultName`.
	BackupVaultName *string `json:"backupVaultName"`
	// `AWS::Backup::BackupVault.AccessPolicy`.
	AccessPolicy interface{} `json:"accessPolicy"`
	// `AWS::Backup::BackupVault.BackupVaultTags`.
	BackupVaultTags interface{} `json:"backupVaultTags"`
	// `AWS::Backup::BackupVault.EncryptionKeyArn`.
	EncryptionKeyArn *string `json:"encryptionKeyArn"`
	// `AWS::Backup::BackupVault.Notifications`.
	Notifications interface{} `json:"notifications"`
}

// A backup plan.
// Experimental.
type IBackupPlan interface {
	awscdk.IResource
	// The identifier of the backup plan.
	// Experimental.
	BackupPlanId() *string
}

// The jsii proxy for IBackupPlan
type jsiiProxy_IBackupPlan struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IBackupPlan) BackupPlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupPlanId",
		&returns,
	)
	return returns
}

// A backup vault.
// Experimental.
type IBackupVault interface {
	awscdk.IResource
	// Grant the actions defined in actions to the given grantee on this backup vault.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// The ARN of the backup vault.
	// Experimental.
	BackupVaultArn() *string
	// The name of a logical container where backups are stored.
	// Experimental.
	BackupVaultName() *string
}

// The jsii proxy for IBackupVault
type jsiiProxy_IBackupVault struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IBackupVault) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IBackupVault) BackupVaultArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBackupVault) BackupVaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultName",
		&returns,
	)
	return returns
}

// A tag condition.
// Experimental.
type TagCondition struct {
	// The key in a key-value pair.
	//
	// For example, in `"ec2:ResourceTag/Department": "accounting"`,
	// `ec2:ResourceTag/Department` is the key.
	// Experimental.
	Key *string `json:"key"`
	// The value in a key-value pair.
	//
	// For example, in `"ec2:ResourceTag/Department": "accounting"`,
	// `accounting` is the value.
	// Experimental.
	Value *string `json:"value"`
	// An operation that is applied to a key-value pair used to filter resources in a selection.
	// Experimental.
	Operation TagOperation `json:"operation"`
}

// An operation that is applied to a key-value pair.
// Experimental.
type TagOperation string

const (
	TagOperation_STRING_EQUALS TagOperation = "STRING_EQUALS"
	TagOperation_DUMMY TagOperation = "DUMMY"
)

