package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// A database secret.
//
// Example:
//   // Build a data source for AppSync to access the database.
//   var api graphqlApi
//   // Create username and password secret for DB Cluster
//   secret := rds.NewDatabaseSecret(this, jsii.String("AuroraSecret"), &DatabaseSecretProps{
//   	Username: jsii.String("clusteradmin"),
//   })
//
//   // The VPC to place the cluster in
//   vpc := ec2.NewVpc(this, jsii.String("AuroraVpc"))
//
//   // Create the serverless cluster, provide all values needed to customise the database.
//   cluster := rds.NewServerlessCluster(this, jsii.String("AuroraCluster"), &ServerlessClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AURORA_MYSQL(),
//   	Vpc: Vpc,
//   	Credentials: map[string]*string{
//   		"username": jsii.String("clusteradmin"),
//   	},
//   	ClusterIdentifier: jsii.String("db-endpoint-test"),
//   	DefaultDatabaseName: jsii.String("demos"),
//   })
//   rdsDS := api.AddRdsDataSource(jsii.String("rds"), cluster, secret, jsii.String("demos"))
//
//   // Set up a resolver for an RDS query.
//   rdsDS.CreateResolver(jsii.String("QueryGetDemosRdsResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Query"),
//   	FieldName: jsii.String("getDemosRds"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromString(jsii.String(`
//   	  {
//   	    "version": "2018-05-29",
//   	    "statements": [
//   	      "SELECT * FROM demos"
//   	    ]
//   	  }
//   	  `)),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
//   	    $utils.toJson($utils.rds.toJsonObject($ctx.result)[0])
//   	  `)),
//   })
//
//   // Set up a resolver for an RDS mutation.
//   rdsDS.CreateResolver(jsii.String("MutationAddDemoRdsResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Mutation"),
//   	FieldName: jsii.String("addDemoRds"),
//   	RequestMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
//   	  {
//   	    "version": "2018-05-29",
//   	    "statements": [
//   	      "INSERT INTO demos VALUES (:id, :version)",
//   	      "SELECT * WHERE id = :id"
//   	    ],
//   	    "variableMap": {
//   	      ":id": $util.toJson($util.autoId()),
//   	      ":version": $util.toJson($ctx.args.version)
//   	    }
//   	  }
//   	  `)),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
//   	    $utils.toJson($utils.rds.toJsonObject($ctx.result)[1][0])
//   	  `)),
//   })
//
type DatabaseSecret interface {
	awssecretsmanager.Secret
	// Provides an identifier for this secret for use in IAM policies.
	//
	// If there is a full ARN, this is just the ARN;
	// if we have a partial ARN -- due to either importing by secret name or partial ARN --
	// then we need to add a suffix to capture the full ARN's format.
	ArnForPolicies() *string
	AutoCreatePolicy() *bool
	// The customer-managed encryption key that is used to encrypt this secret, if any.
	//
	// When not specified, the default
	// KMS key for the account and region is being used.
	EncryptionKey() awskms.IKey
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The string of the characters that are excluded in this secret when it is generated.
	ExcludeCharacters() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The ARN of the secret in AWS Secrets Manager.
	//
	// Will return the full ARN if available, otherwise a partial arn.
	// For secrets imported by the deprecated `fromSecretName`, it will return the `secretName`.
	SecretArn() *string
	// The full ARN of the secret in AWS Secrets Manager, which is the ARN including the Secrets Manager-supplied 6-character suffix.
	//
	// This is equal to `secretArn` in most cases, but is undefined when a full ARN is not available (e.g., secrets imported by name).
	SecretFullArn() *string
	// The name of the secret.
	//
	// For "owned" secrets, this will be the full resource name (secret name + suffix), unless the
	// '@aws-cdk/aws-secretsmanager:parseOwnedSecretName' feature flag is set.
	SecretName() *string
	// Retrieve the value of the stored secret as a `SecretValue`.
	SecretValue() awscdk.SecretValue
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds a replica region for the secret.
	AddReplicaRegion(region *string, encryptionKey awskms.IKey)
	// Adds a rotation schedule to the secret.
	AddRotationSchedule(id *string, options *awssecretsmanager.RotationScheduleOptions) awssecretsmanager.RotationSchedule
	// Adds a statement to the IAM resource policy associated with this secret.
	//
	// If this secret was created in this stack, a resource policy will be
	// automatically created upon the first call to `addToResourcePolicy`. If
	// the secret is imported, then this is a no-op.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Attach a target to this secret.
	//
	// Returns: An attached secret.
	Attach(target awssecretsmanager.ISecretAttachmentTarget) awssecretsmanager.ISecret
	// Denies the `DeleteSecret` action to all principals within the current account.
	DenyAccountRootDelete()
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grants reading the secret value to some role.
	GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant
	// Grants writing and updating the secret value to some role.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Interpret the secret as a JSON object and return a field's value from it as a `SecretValue`.
	SecretValueFromJson(jsonField *string) awscdk.SecretValue
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for DatabaseSecret
type jsiiProxy_DatabaseSecret struct {
	internal.Type__awssecretsmanagerSecret
}

func (j *jsiiProxy_DatabaseSecret) ArnForPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnForPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) ExcludeCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) SecretFullArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretFullArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) SecretValue() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"secretValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewDatabaseSecret(scope constructs.Construct, id *string, props *DatabaseSecretProps) DatabaseSecret {
	_init_.Initialize()

	if err := validateNewDatabaseSecretParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseSecret{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDatabaseSecret_Override(d DatabaseSecret, scope constructs.Construct, id *string, props *DatabaseSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import an existing secret into the Stack.
func DatabaseSecret_FromSecretAttributes(scope constructs.Construct, id *string, attrs *awssecretsmanager.SecretAttributes) awssecretsmanager.ISecret {
	_init_.Initialize()

	if err := validateDatabaseSecret_FromSecretAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns awssecretsmanager.ISecret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		"fromSecretAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports a secret by complete ARN.
//
// The complete ARN is the ARN with the Secrets Manager-supplied suffix.
func DatabaseSecret_FromSecretCompleteArn(scope constructs.Construct, id *string, secretCompleteArn *string) awssecretsmanager.ISecret {
	_init_.Initialize()

	if err := validateDatabaseSecret_FromSecretCompleteArnParameters(scope, id, secretCompleteArn); err != nil {
		panic(err)
	}
	var returns awssecretsmanager.ISecret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		"fromSecretCompleteArn",
		[]interface{}{scope, id, secretCompleteArn},
		&returns,
	)

	return returns
}

// Imports a secret by secret name.
//
// A secret with this name must exist in the same account & region.
// Replaces the deprecated `fromSecretName`.
// Please note this method returns ISecret that only contains partial ARN and could lead to AccessDeniedException
// when you pass the partial ARN to CLI or SDK to get the secret value. If your secret name ends with a hyphen and
// 6 characters, you should always use fromSecretCompleteArn() to avoid potential AccessDeniedException.
// See: https://docs.aws.amazon.com/secretsmanager/latest/userguide/troubleshoot.html#ARN_secretnamehyphen
//
func DatabaseSecret_FromSecretNameV2(scope constructs.Construct, id *string, secretName *string) awssecretsmanager.ISecret {
	_init_.Initialize()

	if err := validateDatabaseSecret_FromSecretNameV2Parameters(scope, id, secretName); err != nil {
		panic(err)
	}
	var returns awssecretsmanager.ISecret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		"fromSecretNameV2",
		[]interface{}{scope, id, secretName},
		&returns,
	)

	return returns
}

// Imports a secret by partial ARN.
//
// The partial ARN is the ARN without the Secrets Manager-supplied suffix.
func DatabaseSecret_FromSecretPartialArn(scope constructs.Construct, id *string, secretPartialArn *string) awssecretsmanager.ISecret {
	_init_.Initialize()

	if err := validateDatabaseSecret_FromSecretPartialArnParameters(scope, id, secretPartialArn); err != nil {
		panic(err)
	}
	var returns awssecretsmanager.ISecret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		"fromSecretPartialArn",
		[]interface{}{scope, id, secretPartialArn},
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
func DatabaseSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseSecret_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func DatabaseSecret_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDatabaseSecret_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DatabaseSecret_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDatabaseSecret_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Secret.
func DatabaseSecret_IsSecret(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseSecret_IsSecretParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		"isSecret",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabaseSecret_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabaseSecret) AddReplicaRegion(region *string, encryptionKey awskms.IKey) {
	if err := d.validateAddReplicaRegionParameters(region); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addReplicaRegion",
		[]interface{}{region, encryptionKey},
	)
}

func (d *jsiiProxy_DatabaseSecret) AddRotationSchedule(id *string, options *awssecretsmanager.RotationScheduleOptions) awssecretsmanager.RotationSchedule {
	if err := d.validateAddRotationScheduleParameters(id, options); err != nil {
		panic(err)
	}
	var returns awssecretsmanager.RotationSchedule

	_jsii_.Invoke(
		d,
		"addRotationSchedule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := d.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		d,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := d.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DatabaseSecret) Attach(target awssecretsmanager.ISecretAttachmentTarget) awssecretsmanager.ISecret {
	if err := d.validateAttachParameters(target); err != nil {
		panic(err)
	}
	var returns awssecretsmanager.ISecret

	_jsii_.Invoke(
		d,
		"attach",
		[]interface{}{target},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) DenyAccountRootDelete() {
	_jsii_.InvokeVoid(
		d,
		"denyAccountRootDelete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecret) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := d.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) GetResourceNameAttribute(nameAttr *string) *string {
	if err := d.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant {
	if err := d.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantRead",
		[]interface{}{grantee, versionStages},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) SecretValueFromJson(jsonField *string) awscdk.SecretValue {
	if err := d.validateSecretValueFromJsonParameters(jsonField); err != nil {
		panic(err)
	}
	var returns awscdk.SecretValue

	_jsii_.Invoke(
		d,
		"secretValueFromJson",
		[]interface{}{jsonField},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

