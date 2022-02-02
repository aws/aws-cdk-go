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
//
// TODO: EXAMPLE
//
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
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
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
//
// TODO: EXAMPLE
//
// Experimental.
type BackupPlanProps struct {
	// The display name of the backup plan.
	// Experimental.
	BackupPlanName *string `json:"backupPlanName" yaml:"backupPlanName"`
	// Rules for the backup plan.
	//
	// Use `addRule()` to add rules after
	// instantiation.
	// Experimental.
	BackupPlanRules *[]BackupPlanRule `json:"backupPlanRules" yaml:"backupPlanRules"`
	// The backup vault where backups are stored.
	// Experimental.
	BackupVault IBackupVault `json:"backupVault" yaml:"backupVault"`
	// Enable Windows VSS backup.
	// See: https://docs.aws.amazon.com/aws-backup/latest/devguide/windows-backups.html
	//
	// Experimental.
	WindowsVss *bool `json:"windowsVss" yaml:"windowsVss"`
}

// A backup plan rule.
//
// TODO: EXAMPLE
//
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
//
// TODO: EXAMPLE
//
// Experimental.
type BackupPlanRuleProps struct {
	// The backup vault where backups are.
	// Experimental.
	BackupVault IBackupVault `json:"backupVault" yaml:"backupVault"`
	// The duration after a backup job is successfully started before it must be completed or it is canceled by AWS Backup.
	// Experimental.
	CompletionWindow awscdk.Duration `json:"completionWindow" yaml:"completionWindow"`
	// Specifies the duration after creation that a recovery point is deleted.
	//
	// Must be greater than `moveToColdStorageAfter`.
	// Experimental.
	DeleteAfter awscdk.Duration `json:"deleteAfter" yaml:"deleteAfter"`
	// Enables continuous backup and point-in-time restores (PITR).
	//
	// Property `deleteAfter` defines the retention period for the backup. It is mandatory if PITR is enabled.
	// If no value is specified, the retention period is set to 35 days which is the maximum retention period supported by PITR.
	//
	// Property `moveToColdStorageAfter` must not be specified because PITR does not support this option.
	// Experimental.
	EnableContinuousBackup *bool `json:"enableContinuousBackup" yaml:"enableContinuousBackup"`
	// Specifies the duration after creation that a recovery point is moved to cold storage.
	// Experimental.
	MoveToColdStorageAfter awscdk.Duration `json:"moveToColdStorageAfter" yaml:"moveToColdStorageAfter"`
	// A display name for the backup rule.
	// Experimental.
	RuleName *string `json:"ruleName" yaml:"ruleName"`
	// A CRON expression specifying when AWS Backup initiates a backup job.
	// Experimental.
	ScheduleExpression awsevents.Schedule `json:"scheduleExpression" yaml:"scheduleExpression"`
	// The duration after a backup is scheduled before a job is canceled if it doesn't start successfully.
	// Experimental.
	StartWindow awscdk.Duration `json:"startWindow" yaml:"startWindow"`
}

// A resource to backup.
//
// TODO: EXAMPLE
//
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
//
// TODO: EXAMPLE
//
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
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
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
//
// TODO: EXAMPLE
//
// Experimental.
type BackupSelectionOptions struct {
	// The resources to backup.
	//
	// Use the helper static methods defined on `BackupResource`.
	// Experimental.
	Resources *[]BackupResource `json:"resources" yaml:"resources"`
	// Whether to automatically give restores permissions to the role that AWS Backup uses.
	//
	// If `true`, the `AWSBackupServiceRolePolicyForRestores` managed
	// policy will be attached to the role.
	// Experimental.
	AllowRestores *bool `json:"allowRestores" yaml:"allowRestores"`
	// The name for this selection.
	// Experimental.
	BackupSelectionName *string `json:"backupSelectionName" yaml:"backupSelectionName"`
	// The role that AWS Backup uses to authenticate when backuping or restoring the resources.
	//
	// The `AWSBackupServiceRolePolicyForBackup` managed policy
	// will be attached to this role.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
}

// Properties for a BackupSelection.
//
// TODO: EXAMPLE
//
// Experimental.
type BackupSelectionProps struct {
	// The resources to backup.
	//
	// Use the helper static methods defined on `BackupResource`.
	// Experimental.
	Resources *[]BackupResource `json:"resources" yaml:"resources"`
	// Whether to automatically give restores permissions to the role that AWS Backup uses.
	//
	// If `true`, the `AWSBackupServiceRolePolicyForRestores` managed
	// policy will be attached to the role.
	// Experimental.
	AllowRestores *bool `json:"allowRestores" yaml:"allowRestores"`
	// The name for this selection.
	// Experimental.
	BackupSelectionName *string `json:"backupSelectionName" yaml:"backupSelectionName"`
	// The role that AWS Backup uses to authenticate when backuping or restoring the resources.
	//
	// The `AWSBackupServiceRolePolicyForBackup` managed policy
	// will be attached to this role.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// The backup plan for this selection.
	// Experimental.
	BackupPlan IBackupPlan `json:"backupPlan" yaml:"backupPlan"`
}

// A backup vault.
//
// TODO: EXAMPLE
//
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
	AddToAccessPolicy(statement awsiam.PolicyStatement)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	BlockRecoveryPointDeletion()
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

// Adds a statement to the vault access policy.
// Experimental.
func (b *jsiiProxy_BackupVault) AddToAccessPolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		b,
		"addToAccessPolicy",
		[]interface{}{statement},
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
// Experimental.
func (b *jsiiProxy_BackupVault) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Adds a statement to the vault access policy that prevents anyone from deleting a recovery point.
// Experimental.
func (b *jsiiProxy_BackupVault) BlockRecoveryPointDeletion() {
	_jsii_.InvokeVoid(
		b,
		"blockRecoveryPointDeletion",
		nil, // no parameters
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
//
// TODO: EXAMPLE
//
// Experimental.
type BackupVaultProps struct {
	// A resource-based policy that is used to manage access permissions on the backup vault.
	// Experimental.
	AccessPolicy awsiam.PolicyDocument `json:"accessPolicy" yaml:"accessPolicy"`
	// The name of a logical container where backups are stored.
	//
	// Backup vaults
	// are identified by names that are unique to the account used to create
	// them and the AWS Region where they are created.
	// Experimental.
	BackupVaultName *string `json:"backupVaultName" yaml:"backupVaultName"`
	// Whether to add statements to the vault access policy that prevents anyone from deleting a recovery point.
	// Experimental.
	BlockRecoveryPointDeletion *bool `json:"blockRecoveryPointDeletion" yaml:"blockRecoveryPointDeletion"`
	// The server-side encryption key to use to protect your backups.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// The vault events to send.
	// See: https://docs.aws.amazon.com/aws-backup/latest/devguide/sns-notifications.html
	//
	// Experimental.
	NotificationEvents *[]BackupVaultEvents `json:"notificationEvents" yaml:"notificationEvents"`
	// A SNS topic to send vault events to.
	// See: https://docs.aws.amazon.com/aws-backup/latest/devguide/sns-notifications.html
	//
	// Experimental.
	NotificationTopic awssns.ITopic `json:"notificationTopic" yaml:"notificationTopic"`
	// The removal policy to apply to the vault.
	//
	// Note that removing a vault
	// that contains recovery points will fail.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
}

// A CloudFormation `AWS::Backup::BackupPlan`.
//
// Contains an optional backup plan display name and an array of `BackupRule` objects, each of which specifies a backup rule. Each rule in a backup plan is a separate scheduled task and can back up a different selection of AWS resources.
//
// For a sample AWS CloudFormation template, see the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/assigning-resources.html#assigning-resources-cfn) .
//
// TODO: EXAMPLE
//
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// Specifies an object containing resource type and backup options.
//
// This is only supported for Windows VSS backups.
//
// TODO: EXAMPLE
//
type CfnBackupPlan_AdvancedBackupSettingResourceTypeProperty struct {
	// The backup option for the resource.
	//
	// Each option is a key-value pair.
	BackupOptions interface{} `json:"backupOptions" yaml:"backupOptions"`
	// The name of a resource type.
	//
	// The only supported resource type is EC2.
	ResourceType *string `json:"resourceType" yaml:"resourceType"`
}

// Specifies an object containing properties used to create a backup plan.
//
// TODO: EXAMPLE
//
type CfnBackupPlan_BackupPlanResourceTypeProperty struct {
	// The display name of a backup plan.
	BackupPlanName *string `json:"backupPlanName" yaml:"backupPlanName"`
	// An array of `BackupRule` objects, each of which specifies a scheduled task that is used to back up a selection of resources.
	BackupPlanRule interface{} `json:"backupPlanRule" yaml:"backupPlanRule"`
	// A list of backup options for each resource type.
	AdvancedBackupSettings interface{} `json:"advancedBackupSettings" yaml:"advancedBackupSettings"`
}

// Specifies an object containing properties used to schedule a task to back up a selection of resources.
//
// TODO: EXAMPLE
//
type CfnBackupPlan_BackupRuleResourceTypeProperty struct {
	// A display name for a backup rule.
	RuleName *string `json:"ruleName" yaml:"ruleName"`
	// The name of a logical container where backups are stored.
	//
	// Backup vaults are identified by names that are unique to the account used to create them and the AWS Region where they are created. They consist of letters, numbers, and hyphens.
	TargetBackupVault *string `json:"targetBackupVault" yaml:"targetBackupVault"`
	// A value in minutes after a backup job is successfully started before it must be completed or it is canceled by AWS Backup .
	CompletionWindowMinutes *float64 `json:"completionWindowMinutes" yaml:"completionWindowMinutes"`
	// An array of CopyAction objects, which contains the details of the copy operation.
	CopyActions interface{} `json:"copyActions" yaml:"copyActions"`
	// Enables continuous backup and point-in-time restores (PITR).
	EnableContinuousBackup interface{} `json:"enableContinuousBackup" yaml:"enableContinuousBackup"`
	// The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.
	//
	// AWS Backup transitions and expires backups automatically according to the lifecycle that you define.
	Lifecycle interface{} `json:"lifecycle" yaml:"lifecycle"`
	// To help organize your resources, you can assign your own metadata to the resources that you create.
	//
	// Each tag is a key-value pair.
	RecoveryPointTags interface{} `json:"recoveryPointTags" yaml:"recoveryPointTags"`
	// A CRON expression specifying when AWS Backup initiates a backup job.
	ScheduleExpression *string `json:"scheduleExpression" yaml:"scheduleExpression"`
	// An optional value that specifies a period of time in minutes after a backup is scheduled before a job is canceled if it doesn't start successfully.
	StartWindowMinutes *float64 `json:"startWindowMinutes" yaml:"startWindowMinutes"`
}

// Copies backups created by a backup rule to another vault.
//
// TODO: EXAMPLE
//
type CfnBackupPlan_CopyActionResourceTypeProperty struct {
	// An Amazon Resource Name (ARN) that uniquely identifies the destination backup vault for the copied backup.
	//
	// For example, `arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.`
	DestinationBackupVaultArn *string `json:"destinationBackupVaultArn" yaml:"destinationBackupVaultArn"`
	// Defines when a protected resource is transitioned to cold storage and when it expires.
	//
	// AWS Backup transitions and expires backups automatically according to the lifecycle that you define. If you do not specify a lifecycle, AWS Backup applies the lifecycle policy of the source backup to the destination backup.
	//
	// Backups transitioned to cold storage must be stored in cold storage for a minimum of 90 days.
	Lifecycle interface{} `json:"lifecycle" yaml:"lifecycle"`
}

// Specifies an object containing an array of `Transition` objects that determine how long in days before a recovery point transitions to cold storage or is deleted.
//
// TODO: EXAMPLE
//
type CfnBackupPlan_LifecycleResourceTypeProperty struct {
	// Specifies the number of days after creation that a recovery point is deleted.
	//
	// Must be greater than `MoveToColdStorageAfterDays` .
	DeleteAfterDays *float64 `json:"deleteAfterDays" yaml:"deleteAfterDays"`
	// Specifies the number of days after creation that a recovery point is moved to cold storage.
	MoveToColdStorageAfterDays *float64 `json:"moveToColdStorageAfterDays" yaml:"moveToColdStorageAfterDays"`
}

// Properties for defining a `CfnBackupPlan`.
//
// TODO: EXAMPLE
//
type CfnBackupPlanProps struct {
	// Uniquely identifies the backup plan to be associated with the selection of resources.
	BackupPlan interface{} `json:"backupPlan" yaml:"backupPlan"`
	// To help organize your resources, you can assign your own metadata to the resources that you create.
	//
	// Each tag is a key-value pair. The specified tags are assigned to all backups created with this plan.
	BackupPlanTags interface{} `json:"backupPlanTags" yaml:"backupPlanTags"`
}

// A CloudFormation `AWS::Backup::BackupSelection`.
//
// Specifies a set of resources to assign to a backup plan.
//
// For a sample AWS CloudFormation template, see the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/assigning-resources.html#assigning-resources-cfn) .
//
// TODO: EXAMPLE
//
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// Specifies an object containing properties used to assign a set of resources to a backup plan.
//
// TODO: EXAMPLE
//
type CfnBackupSelection_BackupSelectionResourceTypeProperty struct {
	// The ARN of the IAM role that AWS Backup uses to authenticate when backing up the target resource;
	//
	// for example, `arn:aws:iam::123456789012:role/S3Access` .
	IamRoleArn *string `json:"iamRoleArn" yaml:"iamRoleArn"`
	// The display name of a resource selection document.
	SelectionName *string `json:"selectionName" yaml:"selectionName"`
	// A list of conditions that you define to assign resources to your backup plans using tags.
	//
	// For example, `"StringEquals": {"Department": "accounting"` . Condition operators are case sensitive.
	//
	// `Conditions` differs from `ListOfTags` as follows:
	//
	// - When you specify more than one condition, you only assign the resources that match ALL conditions (using AND logic).
	// - `Conditions` supports `StringEquals` , `StringLike` , `StringNotEquals` , and `StringNotLike` . `ListOfTags` only supports `StringEquals` .
	Conditions interface{} `json:"conditions" yaml:"conditions"`
	// An array of conditions used to specify a set of resources to assign to a backup plan;
	//
	// for example, `"STRINGEQUALS": {"Department":"accounting"` .
	ListOfTags interface{} `json:"listOfTags" yaml:"listOfTags"`
	// A list of Amazon Resource Names (ARNs) to exclude from a backup plan.
	//
	// The maximum number of ARNs is 500 without wildcards, or 30 ARNs with wildcards.
	//
	// If you need to exclude many resources from a backup plan, consider a different resource selection strategy, such as assigning only one or a few resource types or refining your resource selection using tags.
	NotResources *[]*string `json:"notResources" yaml:"notResources"`
	// An array of strings that contain Amazon Resource Names (ARNs) of resources to assign to a backup plan.
	Resources *[]*string `json:"resources" yaml:"resources"`
}

// Specifies an object that contains an array of triplets made up of a condition type (such as `STRINGEQUALS` ), a key, and a value.
//
// Conditions are used to filter resources in a selection that is assigned to a backup plan.
//
// TODO: EXAMPLE
//
type CfnBackupSelection_ConditionResourceTypeProperty struct {
	// The key in a key-value pair.
	//
	// For example, in `"Department": "accounting"` , `"Department"` is the key.
	ConditionKey *string `json:"conditionKey" yaml:"conditionKey"`
	// An operation, such as `STRINGEQUALS` , that is applied to a key-value pair used to filter resources in a selection.
	ConditionType *string `json:"conditionType" yaml:"conditionType"`
	// The value in a key-value pair.
	//
	// For example, in `"Department": "accounting"` , `"accounting"` is the value.
	ConditionValue *string `json:"conditionValue" yaml:"conditionValue"`
}

// Properties for defining a `CfnBackupSelection`.
//
// TODO: EXAMPLE
//
type CfnBackupSelectionProps struct {
	// Uniquely identifies a backup plan.
	BackupPlanId *string `json:"backupPlanId" yaml:"backupPlanId"`
	// Specifies the body of a request to assign a set of resources to a backup plan.
	//
	// It includes an array of resources, an optional array of patterns to exclude resources, an optional role to provide access to the AWS service the resource belongs to, and an optional array of tags used to identify a set of resources.
	BackupSelection interface{} `json:"backupSelection" yaml:"backupSelection"`
}

// A CloudFormation `AWS::Backup::BackupVault`.
//
// Creates a logical container where backups are stored. A `CreateBackupVault` request includes a name, optionally one or more resource tags, an encryption key, and a request ID.
//
// Do not include sensitive data, such as passport numbers, in the name of a backup vault.
//
// For a sample AWS CloudFormation template, see the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/assigning-resources.html#assigning-resources-cfn) .
//
// TODO: EXAMPLE
//
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
	LockConfiguration() interface{}
	SetLockConfiguration(val interface{})
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

func (j *jsiiProxy_CfnBackupVault) LockConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lockConfiguration",
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

func (j *jsiiProxy_CfnBackupVault) SetLockConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"lockConfiguration",
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The `LockConfigurationType` property type specifies configuration for [AWS Backup Vault Lock](https://docs.aws.amazon.com/aws-backup/latest/devguide/vault-lock.html) .
//
// TODO: EXAMPLE
//
type CfnBackupVault_LockConfigurationTypeProperty struct {
	// The AWS Backup Vault Lock configuration that specifies the minimum retention period that the vault retains its recovery points.
	//
	// This setting can be useful if, for example, your organization's policies require you to retain certain data for at least seven years (2555 days).
	//
	// If this parameter is not specified, Vault Lock will not enforce a minimum retention period.
	//
	// If this parameter is specified, any backup or copy job to the vault must have a lifecycle policy with a retention period equal to or longer than the minimum retention period. If the job's retention period is shorter than that minimum retention period, then the vault fails that backup or copy job, and you should either modify your lifecycle settings or use a different vault. Recovery points already saved in the vault prior to Vault Lock are not affected.
	MinRetentionDays *float64 `json:"minRetentionDays" yaml:"minRetentionDays"`
	// The AWS Backup Vault Lock configuration that specifies the number of days before the lock date.
	//
	// For example, setting `ChangeableForDays` to 30 on Jan. 1, 2022 at 8pm UTC will set the lock date to Jan. 31, 2022 at 8pm UTC.
	//
	// AWS Backup enforces a 72-hour cooling-off period before Vault Lock takes effect and becomes immutable. Therefore, you must set `ChangeableForDays` to 3 or greater.
	//
	// Before the lock date, you can delete Vault Lock from the vault using `DeleteBackupVaultLockConfiguration` or change the Vault Lock configuration using `PutBackupVaultLockConfiguration` . On and after the lock date, the Vault Lock becomes immutable and cannot be changed or deleted.
	//
	// If this parameter is not specified, you can delete Vault Lock from the vault using `DeleteBackupVaultLockConfiguration` or change the Vault Lock configuration using `PutBackupVaultLockConfiguration` at any time.
	ChangeableForDays *float64 `json:"changeableForDays" yaml:"changeableForDays"`
	// The AWS Backup Vault Lock configuration that specifies the maximum retention period that the vault retains its recovery points.
	//
	// This setting can be useful if, for example, your organization's policies require you to destroy certain data after retaining it for four years (1460 days).
	//
	// If this parameter is not included, Vault Lock does not enforce a maximum retention period on the recovery points in the vault. If this parameter is included without a value, Vault Lock will not enforce a maximum retention period.
	//
	// If this parameter is specified, any backup or copy job to the vault must have a lifecycle policy with a retention period equal to or shorter than the maximum retention period. If the job's retention period is longer than that maximum retention period, then the vault fails the backup or copy job, and you should either modify your lifecycle settings or use a different vault. Recovery points already saved in the vault prior to Vault Lock are not affected.
	MaxRetentionDays *float64 `json:"maxRetentionDays" yaml:"maxRetentionDays"`
}

// Specifies an object containing SNS event notification properties for the target backup vault.
//
// TODO: EXAMPLE
//
type CfnBackupVault_NotificationObjectTypeProperty struct {
	// An array of events that indicate the status of jobs to back up resources to the backup vault.
	//
	// For valid events, see [BackupVaultEvents](https://docs.aws.amazon.com/aws-backup/latest/devguide/API_PutBackupVaultNotifications.html#API_PutBackupVaultNotifications_RequestSyntax) in the *AWS Backup API Guide* .
	BackupVaultEvents *[]*string `json:"backupVaultEvents" yaml:"backupVaultEvents"`
	// An ARN that uniquely identifies an Amazon Simple Notification Service (Amazon SNS) topic;
	//
	// for example, `arn:aws:sns:us-west-2:111122223333:MyTopic` .
	SnsTopicArn *string `json:"snsTopicArn" yaml:"snsTopicArn"`
}

// Properties for defining a `CfnBackupVault`.
//
// TODO: EXAMPLE
//
type CfnBackupVaultProps struct {
	// The name of a logical container where backups are stored.
	//
	// Backup vaults are identified by names that are unique to the account used to create them and the AWS Region where they are created. They consist of lowercase letters, numbers, and hyphens.
	BackupVaultName *string `json:"backupVaultName" yaml:"backupVaultName"`
	// A resource-based policy that is used to manage access permissions on the target backup vault.
	AccessPolicy interface{} `json:"accessPolicy" yaml:"accessPolicy"`
	// Metadata that you can assign to help organize the resources that you create.
	//
	// Each tag is a key-value pair.
	BackupVaultTags interface{} `json:"backupVaultTags" yaml:"backupVaultTags"`
	// The server-side encryption key that is used to protect your backups;
	//
	// for example, `arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab` .
	EncryptionKeyArn *string `json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// Configuration for [AWS Backup Vault Lock](https://docs.aws.amazon.com/aws-backup/latest/devguide/vault-lock.html) .
	LockConfiguration interface{} `json:"lockConfiguration" yaml:"lockConfiguration"`
	// The SNS event notifications for the specified backup vault.
	Notifications interface{} `json:"notifications" yaml:"notifications"`
}

// A CloudFormation `AWS::Backup::Framework`.
//
// Creates a framework with one or more controls. A framework is a collection of controls that you can use to evaluate your backup practices. By using pre-built customizable controls to define your policies, you can evaluate whether your backup practices comply with your policies and which resources are not yet in compliance.
//
// For a sample AWS CloudFormation template, see the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/bam-cfn-integration.html#bam-cfn-frameworks-template) .
//
// TODO: EXAMPLE
//
type CfnFramework interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrCreationTime() awscdk.IResolvable
	AttrDeploymentStatus() *string
	AttrFrameworkArn() *string
	AttrFrameworkStatus() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	FrameworkControls() interface{}
	SetFrameworkControls(val interface{})
	FrameworkDescription() *string
	SetFrameworkDescription(val *string)
	FrameworkName() *string
	SetFrameworkName(val *string)
	FrameworkTags() interface{}
	SetFrameworkTags(val interface{})
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

// The jsii proxy struct for CfnFramework
type jsiiProxy_CfnFramework struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFramework) AttrCreationTime() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) AttrDeploymentStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDeploymentStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) AttrFrameworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFrameworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) AttrFrameworkStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFrameworkStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) FrameworkControls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frameworkControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) FrameworkDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameworkDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) FrameworkName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameworkName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) FrameworkTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frameworkTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFramework) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Backup::Framework`.
func NewCfnFramework(scope awscdk.Construct, id *string, props *CfnFrameworkProps) CfnFramework {
	_init_.Initialize()

	j := jsiiProxy_CfnFramework{}

	_jsii_.Create(
		"monocdk.aws_backup.CfnFramework",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Backup::Framework`.
func NewCfnFramework_Override(c CfnFramework, scope awscdk.Construct, id *string, props *CfnFrameworkProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_backup.CfnFramework",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFramework) SetFrameworkControls(val interface{}) {
	_jsii_.Set(
		j,
		"frameworkControls",
		val,
	)
}

func (j *jsiiProxy_CfnFramework) SetFrameworkDescription(val *string) {
	_jsii_.Set(
		j,
		"frameworkDescription",
		val,
	)
}

func (j *jsiiProxy_CfnFramework) SetFrameworkName(val *string) {
	_jsii_.Set(
		j,
		"frameworkName",
		val,
	)
}

func (j *jsiiProxy_CfnFramework) SetFrameworkTags(val interface{}) {
	_jsii_.Set(
		j,
		"frameworkTags",
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
func CfnFramework_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.CfnFramework",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnFramework_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.CfnFramework",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnFramework_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.CfnFramework",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFramework_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_backup.CfnFramework",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnFramework) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnFramework) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnFramework) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnFramework) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnFramework) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnFramework) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnFramework) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnFramework) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnFramework) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnFramework) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnFramework) OnPrepare() {
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
func (c *jsiiProxy_CfnFramework) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnFramework) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnFramework) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnFramework) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFramework) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnFramework) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnFramework) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnFramework) ToString() *string {
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
func (c *jsiiProxy_CfnFramework) Validate() *[]*string {
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
func (c *jsiiProxy_CfnFramework) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A list of parameters for a control.
//
// A control can have zero, one, or more than one parameter. An example of a control with two parameters is: "backup plan frequency is at least `daily` and the retention period is at least `1 year` ". The first parameter is `daily` . The second parameter is `1 year` .
//
// TODO: EXAMPLE
//
type CfnFramework_ControlInputParameterProperty struct {
	// The name of a parameter, for example, `BackupPlanFrequency` .
	ParameterName *string `json:"parameterName" yaml:"parameterName"`
	// The value of parameter, for example, `hourly` .
	ParameterValue *string `json:"parameterValue" yaml:"parameterValue"`
}

// Contains detailed information about all of the controls of a framework.
//
// Each framework must contain at least one control.
//
// TODO: EXAMPLE
//
type CfnFramework_FrameworkControlProperty struct {
	// The name of a control.
	//
	// This name is between 1 and 256 characters.
	ControlName *string `json:"controlName" yaml:"controlName"`
	// A list of `ParameterName` and `ParameterValue` pairs.
	ControlInputParameters interface{} `json:"controlInputParameters" yaml:"controlInputParameters"`
	// The scope of a control.
	//
	// The control scope defines what the control will evaluate. Three examples of control scopes are: a specific backup plan, all backup plans with a specific tag, or all backup plans. For more information, see `ControlScope` .
	ControlScope interface{} `json:"controlScope" yaml:"controlScope"`
}

// Properties for defining a `CfnFramework`.
//
// TODO: EXAMPLE
//
type CfnFrameworkProps struct {
	// Contains detailed information about all of the controls of a framework.
	//
	// Each framework must contain at least one control.
	FrameworkControls interface{} `json:"frameworkControls" yaml:"frameworkControls"`
	// An optional description of the framework with a maximum 1,024 characters.
	FrameworkDescription *string `json:"frameworkDescription" yaml:"frameworkDescription"`
	// The unique name of a framework.
	//
	// This name is between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).
	FrameworkName *string `json:"frameworkName" yaml:"frameworkName"`
	// A list of tags with which to tag your framework.
	FrameworkTags interface{} `json:"frameworkTags" yaml:"frameworkTags"`
}

// A CloudFormation `AWS::Backup::ReportPlan`.
//
// Creates a report plan. A report plan is a document that contains information about the contents of the report and where AWS Backup will deliver it.
//
// If you call `CreateReportPlan` with a plan that already exists, you receive an `AlreadyExistsException` exception.
//
// For a sample AWS CloudFormation template, see the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/assigning-resources.html#assigning-resources-cfn) .
//
// TODO: EXAMPLE
//
type CfnReportPlan interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrReportPlanArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	ReportDeliveryChannel() interface{}
	SetReportDeliveryChannel(val interface{})
	ReportPlanDescription() *string
	SetReportPlanDescription(val *string)
	ReportPlanName() *string
	SetReportPlanName(val *string)
	ReportPlanTags() interface{}
	SetReportPlanTags(val interface{})
	ReportSetting() interface{}
	SetReportSetting(val interface{})
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

// The jsii proxy struct for CfnReportPlan
type jsiiProxy_CfnReportPlan struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnReportPlan) AttrReportPlanArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrReportPlanArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportPlan) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportPlan) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportPlan) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportPlan) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportPlan) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportPlan) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportPlan) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportPlan) ReportDeliveryChannel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportDeliveryChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportPlan) ReportPlanDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportPlanDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportPlan) ReportPlanName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportPlanName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportPlan) ReportPlanTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportPlanTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportPlan) ReportSetting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportPlan) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportPlan) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Backup::ReportPlan`.
func NewCfnReportPlan(scope awscdk.Construct, id *string, props *CfnReportPlanProps) CfnReportPlan {
	_init_.Initialize()

	j := jsiiProxy_CfnReportPlan{}

	_jsii_.Create(
		"monocdk.aws_backup.CfnReportPlan",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Backup::ReportPlan`.
func NewCfnReportPlan_Override(c CfnReportPlan, scope awscdk.Construct, id *string, props *CfnReportPlanProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_backup.CfnReportPlan",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnReportPlan) SetReportDeliveryChannel(val interface{}) {
	_jsii_.Set(
		j,
		"reportDeliveryChannel",
		val,
	)
}

func (j *jsiiProxy_CfnReportPlan) SetReportPlanDescription(val *string) {
	_jsii_.Set(
		j,
		"reportPlanDescription",
		val,
	)
}

func (j *jsiiProxy_CfnReportPlan) SetReportPlanName(val *string) {
	_jsii_.Set(
		j,
		"reportPlanName",
		val,
	)
}

func (j *jsiiProxy_CfnReportPlan) SetReportPlanTags(val interface{}) {
	_jsii_.Set(
		j,
		"reportPlanTags",
		val,
	)
}

func (j *jsiiProxy_CfnReportPlan) SetReportSetting(val interface{}) {
	_jsii_.Set(
		j,
		"reportSetting",
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
func CfnReportPlan_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.CfnReportPlan",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnReportPlan_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.CfnReportPlan",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnReportPlan_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.CfnReportPlan",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReportPlan_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_backup.CfnReportPlan",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnReportPlan) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnReportPlan) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnReportPlan) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnReportPlan) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnReportPlan) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnReportPlan) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnReportPlan) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnReportPlan) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnReportPlan) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnReportPlan) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnReportPlan) OnPrepare() {
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
func (c *jsiiProxy_CfnReportPlan) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnReportPlan) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnReportPlan) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnReportPlan) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnReportPlan) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnReportPlan) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnReportPlan) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnReportPlan) ToString() *string {
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
func (c *jsiiProxy_CfnReportPlan) Validate() *[]*string {
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
func (c *jsiiProxy_CfnReportPlan) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnReportPlan`.
//
// TODO: EXAMPLE
//
type CfnReportPlanProps struct {
	// Contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.
	ReportDeliveryChannel interface{} `json:"reportDeliveryChannel" yaml:"reportDeliveryChannel"`
	// Identifies the report template for the report. Reports are built using a report template. The report templates are:.
	//
	// `RESOURCE_COMPLIANCE_REPORT | CONTROL_COMPLIANCE_REPORT | BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT`
	//
	// If the report template is `RESOURCE_COMPLIANCE_REPORT` or `CONTROL_COMPLIANCE_REPORT` , this API resource also describes the report coverage by AWS Regions and frameworks.
	ReportSetting interface{} `json:"reportSetting" yaml:"reportSetting"`
	// An optional description of the report plan with a maximum 1,024 characters.
	ReportPlanDescription *string `json:"reportPlanDescription" yaml:"reportPlanDescription"`
	// The unique name of the report plan.
	//
	// This name is between 1 and 256 characters starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).
	ReportPlanName *string `json:"reportPlanName" yaml:"reportPlanName"`
	// A list of tags to tag your report plan.
	ReportPlanTags interface{} `json:"reportPlanTags" yaml:"reportPlanTags"`
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
//
// TODO: EXAMPLE
//
// Experimental.
type TagCondition struct {
	// The key in a key-value pair.
	//
	// For example, in `"ec2:ResourceTag/Department": "accounting"`,
	// `ec2:ResourceTag/Department` is the key.
	// Experimental.
	Key *string `json:"key" yaml:"key"`
	// The value in a key-value pair.
	//
	// For example, in `"ec2:ResourceTag/Department": "accounting"`,
	// `accounting` is the value.
	// Experimental.
	Value *string `json:"value" yaml:"value"`
	// An operation that is applied to a key-value pair used to filter resources in a selection.
	// Experimental.
	Operation TagOperation `json:"operation" yaml:"operation"`
}

// An operation that is applied to a key-value pair.
// Experimental.
type TagOperation string

const (
	TagOperation_STRING_EQUALS TagOperation = "STRING_EQUALS"
	TagOperation_DUMMY TagOperation = "DUMMY"
)

