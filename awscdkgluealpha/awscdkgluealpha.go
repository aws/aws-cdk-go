// The CDK Construct Library for AWS::Glue
package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdkgluealpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Job Code from a local file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage dockerImage
//   var grantable iGrantable
//   var localBundling iLocalBundling
//
//   assetCode := glue_alpha.NewAssetCode(jsii.String("path"), &assetOptions{
//   	assetHash: jsii.String("assetHash"),
//   	assetHashType: cdk.assetHashType_SOURCE,
//   	bundling: &bundlingOptions{
//   		image: dockerImage,
//
//   		// the properties below are optional
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		entrypoint: []*string{
//   			jsii.String("entrypoint"),
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		local: localBundling,
//   		network: jsii.String("network"),
//   		outputType: cdk.bundlingOutput_ARCHIVED,
//   		securityOpt: jsii.String("securityOpt"),
//   		user: jsii.String("user"),
//   		volumes: []dockerVolume{
//   			&dockerVolume{
//   				containerPath: jsii.String("containerPath"),
//   				hostPath: jsii.String("hostPath"),
//
//   				// the properties below are optional
//   				consistency: cdk.dockerVolumeConsistency_CONSISTENT,
//   			},
//   		},
//   		workingDirectory: jsii.String("workingDirectory"),
//   	},
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	followSymlinks: cdk.symlinkFollowMode_NEVER,
//   	ignoreMode: cdk.ignoreMode_GLOB,
//   	readers: []*iGrantable{
//   		grantable,
//   	},
//   })
//
// Experimental.
type AssetCode interface {
	Code
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig
}

// The jsii proxy struct for AssetCode
type jsiiProxy_AssetCode struct {
	jsiiProxy_Code
}

// Experimental.
func NewAssetCode(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	j := jsiiProxy_AssetCode{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.AssetCode",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetCode_Override(a AssetCode, path *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.AssetCode",
		[]interface{}{path, options},
		a,
	)
}

// Job code from a local disk path.
// Experimental.
func AssetCode_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.AssetCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Job code as an S3 object.
// Experimental.
func AssetCode_FromBucket(bucket awss3.IBucket, key *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.AssetCode",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetCode) Bind(scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, grantable},
		&returns,
	)

	return returns
}

// Classification string given to tables with this data format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   classificationString := glue_alpha.classificationString_AVRO()
//
// See: https://docs.aws.amazon.com/glue/latest/dg/add-classifier.html#classifier-built-in
//
// Experimental.
type ClassificationString interface {
	// Experimental.
	Value() *string
}

// The jsii proxy struct for ClassificationString
type jsiiProxy_ClassificationString struct {
	_ byte // padding
}

func (j *jsiiProxy_ClassificationString) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewClassificationString(value *string) ClassificationString {
	_init_.Initialize()

	j := jsiiProxy_ClassificationString{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Experimental.
func NewClassificationString_Override(c ClassificationString, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		[]interface{}{value},
		c,
	)
}

func ClassificationString_AVRO() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		"AVRO",
		&returns,
	)
	return returns
}

func ClassificationString_CSV() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		"CSV",
		&returns,
	)
	return returns
}

func ClassificationString_JSON() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		"JSON",
		&returns,
	)
	return returns
}

func ClassificationString_ORC() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		"ORC",
		&returns,
	)
	return returns
}

func ClassificationString_PARQUET() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		"PARQUET",
		&returns,
	)
	return returns
}

func ClassificationString_XML() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		"XML",
		&returns,
	)
	return returns
}

// CloudWatch Logs encryption configuration.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &securityConfigurationProps{
//   	securityConfigurationName: jsii.String("name"),
//   	cloudWatchEncryption: &cloudWatchEncryption{
//   		mode: glue.cloudWatchEncryptionMode_KMS,
//   	},
//   	jobBookmarksEncryption: &jobBookmarksEncryption{
//   		mode: glue.jobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
//   	},
//   	s3Encryption: &s3Encryption{
//   		mode: glue.s3EncryptionMode_KMS,
//   	},
//   })
//
// Experimental.
type CloudWatchEncryption struct {
	// Encryption mode.
	// Experimental.
	Mode CloudWatchEncryptionMode `field:"required" json:"mode" yaml:"mode"`
	// The KMS key to be used to encrypt the data.
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

// Encryption mode for CloudWatch Logs.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &securityConfigurationProps{
//   	securityConfigurationName: jsii.String("name"),
//   	cloudWatchEncryption: &cloudWatchEncryption{
//   		mode: glue.cloudWatchEncryptionMode_KMS,
//   	},
//   	jobBookmarksEncryption: &jobBookmarksEncryption{
//   		mode: glue.jobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
//   	},
//   	s3Encryption: &s3Encryption{
//   		mode: glue.s3EncryptionMode_KMS,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/glue/latest/webapi/API_CloudWatchEncryption.html#Glue-Type-CloudWatchEncryption-CloudWatchEncryptionMode
//
// Experimental.
type CloudWatchEncryptionMode string

const (
	// Server-side encryption (SSE) with an AWS KMS key managed by the account owner.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html
	//
	// Experimental.
	CloudWatchEncryptionMode_KMS CloudWatchEncryptionMode = "KMS"
)

// Represents a Glue Job's Code assets (an asset can be a scripts, a jar, a python file or any other file).
//
// Example:
//   var bucket bucket
//
//   glue.NewJob(this, jsii.String("PythonShellJob"), &jobProps{
//   	executable: glue.jobExecutable.pythonShell(&pythonShellExecutableProps{
//   		glueVersion: glue.glueVersion_V1_0(),
//   		pythonVersion: glue.pythonVersion_THREE,
//   		script: glue.code.fromBucket(bucket, jsii.String("script.py")),
//   	}),
//   	description: jsii.String("an example Python Shell job"),
//   })
//
// Experimental.
type Code interface {
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig
}

// The jsii proxy struct for Code
type jsiiProxy_Code struct {
	_ byte // padding
}

// Experimental.
func NewCode_Override(c Code) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.Code",
		nil, // no parameters
		c,
	)
}

// Job code from a local disk path.
// Experimental.
func Code_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Code",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Job code as an S3 object.
// Experimental.
func Code_FromBucket(bucket awss3.IBucket, key *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Code",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Code) Bind(scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, grantable},
		&returns,
	)

	return returns
}

// Result of binding `Code` into a `Job`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   codeConfig := &codeConfig{
//   	s3Location: &location{
//   		bucketName: jsii.String("bucketName"),
//   		objectKey: jsii.String("objectKey"),
//
//   		// the properties below are optional
//   		objectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
// Experimental.
type CodeConfig struct {
	// The location of the code in S3.
	// Experimental.
	S3Location *awss3.Location `field:"required" json:"s3Location" yaml:"s3Location"`
}

// A column of a table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   column := &column{
//   	name: jsii.String("name"),
//   	type: &type{
//   		inputString: jsii.String("inputString"),
//   		isPrimitive: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	comment: jsii.String("comment"),
//   }
//
// Experimental.
type Column struct {
	// Name of the column.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of the column.
	// Experimental.
	Type *Type `field:"required" json:"type" yaml:"type"`
	// Coment describing the column.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

// An AWS Glue connection to a data source.
//
// Example:
//   var securityGroup securityGroup
//   var subnet subnet
//
//   glue.NewConnection(this, jsii.String("MyConnection"), &connectionProps{
//   	type: glue.connectionType_NETWORK(),
//   	// The security groups granting AWS Glue inbound access to the data source within the VPC
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	// The VPC subnet which contains the data source
//   	subnet: subnet,
//   })
//
// Experimental.
type Connection interface {
	awscdk.Resource
	IConnection
	// The ARN of the connection.
	// Experimental.
	ConnectionArn() *string
	// The name of the connection.
	// Experimental.
	ConnectionName() *string
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
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	// Add additional connection parameters.
	// Experimental.
	AddProperty(key *string, value *string)
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Connection
type jsiiProxy_Connection struct {
	internal.Type__awscdkResource
	jsiiProxy_IConnection
}

func (j *jsiiProxy_Connection) ConnectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) ConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewConnection(scope constructs.Construct, id *string, props *ConnectionProps) Connection {
	_init_.Initialize()

	j := jsiiProxy_Connection{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.Connection",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewConnection_Override(c Connection, scope constructs.Construct, id *string, props *ConnectionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.Connection",
		[]interface{}{scope, id, props},
		c,
	)
}

// Creates a Connection construct that represents an external connection.
// Experimental.
func Connection_FromConnectionArn(scope constructs.Construct, id *string, connectionArn *string) IConnection {
	_init_.Initialize()

	var returns IConnection

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Connection",
		"fromConnectionArn",
		[]interface{}{scope, id, connectionArn},
		&returns,
	)

	return returns
}

// Creates a Connection construct that represents an external connection.
// Experimental.
func Connection_FromConnectionName(scope constructs.Construct, id *string, connectionName *string) IConnection {
	_init_.Initialize()

	var returns IConnection

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Connection",
		"fromConnectionName",
		[]interface{}{scope, id, connectionName},
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
// Experimental.
func Connection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Connection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func Connection_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Connection",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Connection_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Connection",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connection) AddProperty(key *string, value *string) {
	_jsii_.InvokeVoid(
		c,
		"addProperty",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_Connection) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_Connection) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connection) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connection) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Base Connection Options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//   var subnet subnet
//
//   connectionOptions := &connectionOptions{
//   	connectionName: jsii.String("connectionName"),
//   	description: jsii.String("description"),
//   	matchCriteria: []*string{
//   		jsii.String("matchCriteria"),
//   	},
//   	properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	subnet: subnet,
//   }
//
// Experimental.
type ConnectionOptions struct {
	// The name of the connection.
	// Experimental.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// The description of the connection.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of criteria that can be used in selecting this connection.
	//
	// This is useful for filtering the results of https://awscli.amazonaws.com/v2/documentation/api/latest/reference/glue/get-connections.html
	// Experimental.
	MatchCriteria *[]*string `field:"optional" json:"matchCriteria" yaml:"matchCriteria"`
	// Key-Value pairs that define parameters for the connection.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-connect.html
	//
	// Experimental.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// The list of security groups needed to successfully make this connection e.g. to successfully connect to VPC.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The VPC subnet to connect to resources within a VPC.
	//
	// See more at https://docs.aws.amazon.com/glue/latest/dg/start-connecting.html.
	// Experimental.
	Subnet awsec2.ISubnet `field:"optional" json:"subnet" yaml:"subnet"`
}

// Construction properties for {@link Connection}.
//
// Example:
//   var securityGroup securityGroup
//   var subnet subnet
//
//   glue.NewConnection(this, jsii.String("MyConnection"), &connectionProps{
//   	type: glue.connectionType_NETWORK(),
//   	// The security groups granting AWS Glue inbound access to the data source within the VPC
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	// The VPC subnet which contains the data source
//   	subnet: subnet,
//   })
//
// Experimental.
type ConnectionProps struct {
	// The name of the connection.
	// Experimental.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// The description of the connection.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of criteria that can be used in selecting this connection.
	//
	// This is useful for filtering the results of https://awscli.amazonaws.com/v2/documentation/api/latest/reference/glue/get-connections.html
	// Experimental.
	MatchCriteria *[]*string `field:"optional" json:"matchCriteria" yaml:"matchCriteria"`
	// Key-Value pairs that define parameters for the connection.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-connect.html
	//
	// Experimental.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// The list of security groups needed to successfully make this connection e.g. to successfully connect to VPC.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The VPC subnet to connect to resources within a VPC.
	//
	// See more at https://docs.aws.amazon.com/glue/latest/dg/start-connecting.html.
	// Experimental.
	Subnet awsec2.ISubnet `field:"optional" json:"subnet" yaml:"subnet"`
	// The type of the connection.
	// Experimental.
	Type ConnectionType `field:"required" json:"type" yaml:"type"`
}

// The type of the glue connection.
//
// If you need to use a connection type that doesn't exist as a static member, you
// can instantiate a `ConnectionType` object, e.g: `new ConnectionType('NEW_TYPE')`.
//
// Example:
//   var securityGroup securityGroup
//   var subnet subnet
//
//   glue.NewConnection(this, jsii.String("MyConnection"), &connectionProps{
//   	type: glue.connectionType_NETWORK(),
//   	// The security groups granting AWS Glue inbound access to the data source within the VPC
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	// The VPC subnet which contains the data source
//   	subnet: subnet,
//   })
//
// Experimental.
type ConnectionType interface {
	// The name of this ConnectionType, as expected by Connection resource.
	// Experimental.
	Name() *string
	// The connection type name as expected by Connection resource.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectionType
type jsiiProxy_ConnectionType struct {
	_ byte // padding
}

func (j *jsiiProxy_ConnectionType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewConnectionType(name *string) ConnectionType {
	_init_.Initialize()

	j := jsiiProxy_ConnectionType{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		[]interface{}{name},
		&j,
	)

	return &j
}

// Experimental.
func NewConnectionType_Override(c ConnectionType, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		[]interface{}{name},
		c,
	)
}

func ConnectionType_JDBC() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"JDBC",
		&returns,
	)
	return returns
}

func ConnectionType_KAFKA() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"KAFKA",
		&returns,
	)
	return returns
}

func ConnectionType_MONGODB() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"MONGODB",
		&returns,
	)
	return returns
}

func ConnectionType_NETWORK() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"NETWORK",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ConnectionType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for enabling Continuous Logging for Glue Jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var logGroup logGroup
//
//   continuousLoggingProps := &continuousLoggingProps{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	conversionPattern: jsii.String("conversionPattern"),
//   	logGroup: logGroup,
//   	logStreamPrefix: jsii.String("logStreamPrefix"),
//   	quiet: jsii.Boolean(false),
//   }
//
// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
//
// Experimental.
type ContinuousLoggingProps struct {
	// Enable continouous logging.
	// Experimental.
	Enabled *bool `field:"required" json:"enabled" yaml:"enabled"`
	// Apply the provided conversion pattern.
	//
	// This is a Log4j Conversion Pattern to customize driver and executor logs.
	// Experimental.
	ConversionPattern *string `field:"optional" json:"conversionPattern" yaml:"conversionPattern"`
	// Specify a custom CloudWatch log group name.
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Specify a custom CloudWatch log stream prefix.
	// Experimental.
	LogStreamPrefix *string `field:"optional" json:"logStreamPrefix" yaml:"logStreamPrefix"`
	// Filter out non-useful Apache Spark driver/executor and Apache Hadoop YARN heartbeat log messages.
	// Experimental.
	Quiet *bool `field:"optional" json:"quiet" yaml:"quiet"`
}

// Defines the input/output formats and ser/de for a single DataFormat.
//
// Example:
//   var myDatabase database
//
//   glue.NewTable(this, jsii.String("MyTable"), &tableProps{
//   	database: myDatabase,
//   	tableName: jsii.String("my_table"),
//   	columns: []column{
//   		&column{
//   			name: jsii.String("col1"),
//   			type: glue.schema_STRING(),
//   		},
//   	},
//   	partitionKeys: []*column{
//   		&column{
//   			name: jsii.String("year"),
//   			type: glue.*schema_SMALL_INT(),
//   		},
//   		&column{
//   			name: jsii.String("month"),
//   			type: glue.*schema_SMALL_INT(),
//   		},
//   	},
//   	dataFormat: glue.dataFormat_JSON(),
//   	enablePartitionFiltering: jsii.Boolean(true),
//   })
//
// Experimental.
type DataFormat interface {
	// Classification string given to tables with this data format.
	// Experimental.
	ClassificationString() ClassificationString
	// `InputFormat` for this data format.
	// Experimental.
	InputFormat() InputFormat
	// `OutputFormat` for this data format.
	// Experimental.
	OutputFormat() OutputFormat
	// Serialization library for this data format.
	// Experimental.
	SerializationLibrary() SerializationLibrary
}

// The jsii proxy struct for DataFormat
type jsiiProxy_DataFormat struct {
	_ byte // padding
}

func (j *jsiiProxy_DataFormat) ClassificationString() ClassificationString {
	var returns ClassificationString
	_jsii_.Get(
		j,
		"classificationString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFormat) InputFormat() InputFormat {
	var returns InputFormat
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFormat) OutputFormat() OutputFormat {
	var returns OutputFormat
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFormat) SerializationLibrary() SerializationLibrary {
	var returns SerializationLibrary
	_jsii_.Get(
		j,
		"serializationLibrary",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataFormat(props *DataFormatProps) DataFormat {
	_init_.Initialize()

	j := jsiiProxy_DataFormat{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.DataFormat",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewDataFormat_Override(d DataFormat, props *DataFormatProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.DataFormat",
		[]interface{}{props},
		d,
	)
}

func DataFormat_APACHE_LOGS() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.DataFormat",
		"APACHE_LOGS",
		&returns,
	)
	return returns
}

func DataFormat_AVRO() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.DataFormat",
		"AVRO",
		&returns,
	)
	return returns
}

func DataFormat_CLOUDTRAIL_LOGS() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.DataFormat",
		"CLOUDTRAIL_LOGS",
		&returns,
	)
	return returns
}

func DataFormat_CSV() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.DataFormat",
		"CSV",
		&returns,
	)
	return returns
}

func DataFormat_JSON() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.DataFormat",
		"JSON",
		&returns,
	)
	return returns
}

func DataFormat_LOGSTASH() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.DataFormat",
		"LOGSTASH",
		&returns,
	)
	return returns
}

func DataFormat_ORC() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.DataFormat",
		"ORC",
		&returns,
	)
	return returns
}

func DataFormat_PARQUET() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.DataFormat",
		"PARQUET",
		&returns,
	)
	return returns
}

func DataFormat_TSV() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.DataFormat",
		"TSV",
		&returns,
	)
	return returns
}

// Properties of a DataFormat instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   var classificationString classificationString
//   var inputFormat inputFormat
//   var outputFormat outputFormat
//   var serializationLibrary serializationLibrary
//
//   dataFormatProps := &dataFormatProps{
//   	inputFormat: inputFormat,
//   	outputFormat: outputFormat,
//   	serializationLibrary: serializationLibrary,
//
//   	// the properties below are optional
//   	classificationString: classificationString,
//   }
//
// Experimental.
type DataFormatProps struct {
	// `InputFormat` for this data format.
	// Experimental.
	InputFormat InputFormat `field:"required" json:"inputFormat" yaml:"inputFormat"`
	// `OutputFormat` for this data format.
	// Experimental.
	OutputFormat OutputFormat `field:"required" json:"outputFormat" yaml:"outputFormat"`
	// Serialization library for this data format.
	// Experimental.
	SerializationLibrary SerializationLibrary `field:"required" json:"serializationLibrary" yaml:"serializationLibrary"`
	// Classification string given to tables with this data format.
	// Experimental.
	ClassificationString ClassificationString `field:"optional" json:"classificationString" yaml:"classificationString"`
}

// A Glue database.
//
// Example:
//   glue.NewDatabase(this, jsii.String("MyDatabase"), &databaseProps{
//   	databaseName: jsii.String("my_database"),
//   })
//
// Experimental.
type Database interface {
	awscdk.Resource
	IDatabase
	// ARN of the Glue catalog in which this database is stored.
	// Experimental.
	CatalogArn() *string
	// The catalog id of the database (usually, the AWS account id).
	// Experimental.
	CatalogId() *string
	// ARN of this database.
	// Experimental.
	DatabaseArn() *string
	// Name of this database.
	// Experimental.
	DatabaseName() *string
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
	// Location URI of this database.
	// Experimental.
	LocationUri() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Database
type jsiiProxy_Database struct {
	internal.Type__awscdkResource
	jsiiProxy_IDatabase
}

func (j *jsiiProxy_Database) CatalogArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) DatabaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) LocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDatabase(scope constructs.Construct, id *string, props *DatabaseProps) Database {
	_init_.Initialize()

	j := jsiiProxy_Database{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.Database",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDatabase_Override(d Database, scope constructs.Construct, id *string, props *DatabaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.Database",
		[]interface{}{scope, id, props},
		d,
	)
}

// Experimental.
func Database_FromDatabaseArn(scope constructs.Construct, id *string, databaseArn *string) IDatabase {
	_init_.Initialize()

	var returns IDatabase

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Database",
		"fromDatabaseArn",
		[]interface{}{scope, id, databaseArn},
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
// Experimental.
func Database_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Database",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func Database_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Database",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Database_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Database",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_Database) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   glue.NewDatabase(this, jsii.String("MyDatabase"), &databaseProps{
//   	databaseName: jsii.String("my_database"),
//   })
//
// Experimental.
type DatabaseProps struct {
	// The name of the database.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The location of the database (for example, an HDFS path).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-databaseinput.html
	//
	// Experimental.
	LocationUri *string `field:"optional" json:"locationUri" yaml:"locationUri"`
}

// AWS Glue version determines the versions of Apache Spark and Python that are available to the job.
//
// Example:
//   var bucket bucket
//
//   glue.NewJob(this, jsii.String("PythonShellJob"), &jobProps{
//   	executable: glue.jobExecutable.pythonShell(&pythonShellExecutableProps{
//   		glueVersion: glue.glueVersion_V1_0(),
//   		pythonVersion: glue.pythonVersion_THREE,
//   		script: glue.code.fromBucket(bucket, jsii.String("script.py")),
//   	}),
//   	description: jsii.String("an example Python Shell job"),
//   })
//
// See: https://docs.aws.amazon.com/glue/latest/dg/add-job.html.
//
// If you need to use a GlueVersion that doesn't exist as a static member, you
// can instantiate a `GlueVersion` object, e.g: `GlueVersion.of('1.5')`.
//
// Experimental.
type GlueVersion interface {
	// The name of this GlueVersion, as expected by Job resource.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for GlueVersion
type jsiiProxy_GlueVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_GlueVersion) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Custom Glue version.
// Experimental.
func GlueVersion_Of(version *string) GlueVersion {
	_init_.Initialize()

	var returns GlueVersion

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.GlueVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func GlueVersion_V0_9() GlueVersion {
	_init_.Initialize()
	var returns GlueVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.GlueVersion",
		"V0_9",
		&returns,
	)
	return returns
}

func GlueVersion_V1_0() GlueVersion {
	_init_.Initialize()
	var returns GlueVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.GlueVersion",
		"V1_0",
		&returns,
	)
	return returns
}

func GlueVersion_V2_0() GlueVersion {
	_init_.Initialize()
	var returns GlueVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.GlueVersion",
		"V2_0",
		&returns,
	)
	return returns
}

func GlueVersion_V3_0() GlueVersion {
	_init_.Initialize()
	var returns GlueVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.GlueVersion",
		"V3_0",
		&returns,
	)
	return returns
}

// Interface representing a created or an imported {@link Connection}.
// Experimental.
type IConnection interface {
	awscdk.IResource
	// The ARN of the connection.
	// Experimental.
	ConnectionArn() *string
	// The name of the connection.
	// Experimental.
	ConnectionName() *string
}

// The jsii proxy for IConnection
type jsiiProxy_IConnection struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IConnection) ConnectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnection) ConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionName",
		&returns,
	)
	return returns
}

// Experimental.
type IDatabase interface {
	awscdk.IResource
	// The ARN of the catalog.
	// Experimental.
	CatalogArn() *string
	// The catalog id of the database (usually, the AWS account id).
	// Experimental.
	CatalogId() *string
	// The ARN of the database.
	// Experimental.
	DatabaseArn() *string
	// The name of the database.
	// Experimental.
	DatabaseName() *string
}

// The jsii proxy for IDatabase
type jsiiProxy_IDatabase struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDatabase) CatalogArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabase) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabase) DatabaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabase) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

// Interface representing a created or an imported {@link Job}.
// Experimental.
type IJob interface {
	awsiam.IGrantable
	awscdk.IResource
	// Create a CloudWatch metric.
	// See: https://docs.aws.amazon.com/glue/latest/dg/monitoring-awsglue-with-cloudwatch-metrics.html
	//
	// Experimental.
	Metric(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Create a CloudWatch Metric indicating job failure.
	// Experimental.
	MetricFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Create a CloudWatch Metric indicating job success.
	// Experimental.
	MetricSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Create a CloudWatch Metric indicating job timeout.
	// Experimental.
	MetricTimeout(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Defines a CloudWatch event rule triggered when something happens with this job.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types
	//
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule triggered when this job moves to the FAILED state.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types
	//
	// Experimental.
	OnFailure(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule triggered when this job moves to the input jobState.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types
	//
	// Experimental.
	OnStateChange(id *string, jobState JobState, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule triggered when this job moves to the SUCCEEDED state.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types
	//
	// Experimental.
	OnSuccess(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule triggered when this job moves to the TIMEOUT state.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types
	//
	// Experimental.
	OnTimeout(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The ARN of the job.
	// Experimental.
	JobArn() *string
	// The name of the job.
	// Experimental.
	JobName() *string
}

// The jsii proxy for IJob
type jsiiProxy_IJob struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IJob) Metric(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, type_, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) MetricFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFailure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) MetricSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSuccess",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) MetricTimeout(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTimeout",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) OnFailure(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onFailure",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) OnStateChange(id *string, jobState JobState, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onStateChange",
		[]interface{}{id, jobState, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) OnSuccess(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onSuccess",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) OnTimeout(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onTimeout",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IJob) JobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJob) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJob) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJob) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJob) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Interface representing a created or an imported {@link SecurityConfiguration}.
// Experimental.
type ISecurityConfiguration interface {
	awscdk.IResource
	// The name of the security configuration.
	// Experimental.
	SecurityConfigurationName() *string
}

// The jsii proxy for ISecurityConfiguration
type jsiiProxy_ISecurityConfiguration struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ISecurityConfiguration) SecurityConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationName",
		&returns,
	)
	return returns
}

// Experimental.
type ITable interface {
	awscdk.IResource
	// Experimental.
	TableArn() *string
	// Experimental.
	TableName() *string
}

// The jsii proxy for ITable
type jsiiProxy_ITable struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ITable) TableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

// Absolute class name of the Hadoop `InputFormat` to use when reading table files.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   inputFormat := glue_alpha.inputFormat_AVRO()
//
// Experimental.
type InputFormat interface {
	// Experimental.
	ClassName() *string
}

// The jsii proxy struct for InputFormat
type jsiiProxy_InputFormat struct {
	_ byte // padding
}

func (j *jsiiProxy_InputFormat) ClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"className",
		&returns,
	)
	return returns
}


// Experimental.
func NewInputFormat(className *string) InputFormat {
	_init_.Initialize()

	j := jsiiProxy_InputFormat{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.InputFormat",
		[]interface{}{className},
		&j,
	)

	return &j
}

// Experimental.
func NewInputFormat_Override(i InputFormat, className *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.InputFormat",
		[]interface{}{className},
		i,
	)
}

func InputFormat_AVRO() InputFormat {
	_init_.Initialize()
	var returns InputFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.InputFormat",
		"AVRO",
		&returns,
	)
	return returns
}

func InputFormat_CLOUDTRAIL() InputFormat {
	_init_.Initialize()
	var returns InputFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.InputFormat",
		"CLOUDTRAIL",
		&returns,
	)
	return returns
}

func InputFormat_ORC() InputFormat {
	_init_.Initialize()
	var returns InputFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.InputFormat",
		"ORC",
		&returns,
	)
	return returns
}

func InputFormat_PARQUET() InputFormat {
	_init_.Initialize()
	var returns InputFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.InputFormat",
		"PARQUET",
		&returns,
	)
	return returns
}

func InputFormat_TEXT() InputFormat {
	_init_.Initialize()
	var returns InputFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.InputFormat",
		"TEXT",
		&returns,
	)
	return returns
}

// A Glue Job.
//
// Example:
//   var bucket bucket
//
//   glue.NewJob(this, jsii.String("PythonShellJob"), &jobProps{
//   	executable: glue.jobExecutable.pythonShell(&pythonShellExecutableProps{
//   		glueVersion: glue.glueVersion_V1_0(),
//   		pythonVersion: glue.pythonVersion_THREE,
//   		script: glue.code.fromBucket(bucket, jsii.String("script.py")),
//   	}),
//   	description: jsii.String("an example Python Shell job"),
//   })
//
// Experimental.
type Job interface {
	awscdk.Resource
	IJob
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
	// The principal this Glue Job is running as.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// The ARN of the job.
	// Experimental.
	JobArn() *string
	// The name of the job.
	// Experimental.
	JobName() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The IAM role Glue assumes to run this job.
	// Experimental.
	Role() awsiam.IRole
	// The Spark UI logs location if Spark UI monitoring and debugging is enabled.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	SparkUILoggingLocation() *SparkUILoggingLocation
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
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
	// Create a CloudWatch metric.
	// See: https://docs.aws.amazon.com/glue/latest/dg/monitoring-awsglue-with-cloudwatch-metrics.html
	//
	// Experimental.
	Metric(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a CloudWatch Metric indicating job failure.
	//
	// This metric is based on the Rule returned by no-args onFailure() call.
	// Experimental.
	MetricFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a CloudWatch Metric indicating job success.
	//
	// This metric is based on the Rule returned by no-args onSuccess() call.
	// Experimental.
	MetricSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a CloudWatch Metric indicating job timeout.
	//
	// This metric is based on the Rule returned by no-args onTimeout() call.
	// Experimental.
	MetricTimeout(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Create a CloudWatch Event Rule for this Glue Job when it's in a given state.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types
	//
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Return a CloudWatch Event Rule matching FAILED state.
	// Experimental.
	OnFailure(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Create a CloudWatch Event Rule for the transition into the input jobState.
	// Experimental.
	OnStateChange(id *string, jobState JobState, options *awsevents.OnEventOptions) awsevents.Rule
	// Create a CloudWatch Event Rule matching JobState.SUCCEEDED.
	// Experimental.
	OnSuccess(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Return a CloudWatch Event Rule matching TIMEOUT state.
	// Experimental.
	OnTimeout(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Job
type jsiiProxy_Job struct {
	internal.Type__awscdkResource
	jsiiProxy_IJob
}

func (j *jsiiProxy_Job) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) JobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) SparkUILoggingLocation() *SparkUILoggingLocation {
	var returns *SparkUILoggingLocation
	_jsii_.Get(
		j,
		"sparkUILoggingLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewJob(scope constructs.Construct, id *string, props *JobProps) Job {
	_init_.Initialize()

	j := jsiiProxy_Job{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.Job",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewJob_Override(j Job, scope constructs.Construct, id *string, props *JobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.Job",
		[]interface{}{scope, id, props},
		j,
	)
}

// Creates a Glue Job.
// Experimental.
func Job_FromJobAttributes(scope constructs.Construct, id *string, attrs *JobAttributes) IJob {
	_init_.Initialize()

	var returns IJob

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Job",
		"fromJobAttributes",
		[]interface{}{scope, id, attrs},
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
// Experimental.
func Job_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Job",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func Job_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Job",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Job_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Job",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		j,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_Job) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) Metric(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		j,
		"metric",
		[]interface{}{metricName, type_, props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) MetricFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		j,
		"metricFailure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) MetricSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		j,
		"metricSuccess",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) MetricTimeout(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		j,
		"metricTimeout",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		j,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) OnFailure(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		j,
		"onFailure",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) OnStateChange(id *string, jobState JobState, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		j,
		"onStateChange",
		[]interface{}{id, jobState, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) OnSuccess(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		j,
		"onSuccess",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) OnTimeout(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		j,
		"onTimeout",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes for importing {@link Job}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   jobAttributes := &jobAttributes{
//   	jobName: jsii.String("jobName"),
//
//   	// the properties below are optional
//   	role: role,
//   }
//
// Experimental.
type JobAttributes struct {
	// The name of the job.
	// Experimental.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// The IAM role assumed by Glue to run this job.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

// Job bookmarks encryption configuration.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &securityConfigurationProps{
//   	securityConfigurationName: jsii.String("name"),
//   	cloudWatchEncryption: &cloudWatchEncryption{
//   		mode: glue.cloudWatchEncryptionMode_KMS,
//   	},
//   	jobBookmarksEncryption: &jobBookmarksEncryption{
//   		mode: glue.jobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
//   	},
//   	s3Encryption: &s3Encryption{
//   		mode: glue.s3EncryptionMode_KMS,
//   	},
//   })
//
// Experimental.
type JobBookmarksEncryption struct {
	// Encryption mode.
	// Experimental.
	Mode JobBookmarksEncryptionMode `field:"required" json:"mode" yaml:"mode"`
	// The KMS key to be used to encrypt the data.
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

// Encryption mode for Job Bookmarks.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &securityConfigurationProps{
//   	securityConfigurationName: jsii.String("name"),
//   	cloudWatchEncryption: &cloudWatchEncryption{
//   		mode: glue.cloudWatchEncryptionMode_KMS,
//   	},
//   	jobBookmarksEncryption: &jobBookmarksEncryption{
//   		mode: glue.jobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
//   	},
//   	s3Encryption: &s3Encryption{
//   		mode: glue.s3EncryptionMode_KMS,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/glue/latest/webapi/API_JobBookmarksEncryption.html#Glue-Type-JobBookmarksEncryption-JobBookmarksEncryptionMode
//
// Experimental.
type JobBookmarksEncryptionMode string

const (
	// Client-side encryption (CSE) with an AWS KMS key managed by the account owner.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html
	//
	// Experimental.
	JobBookmarksEncryptionMode_CLIENT_SIDE_KMS JobBookmarksEncryptionMode = "CLIENT_SIDE_KMS"
)

// The executable properties related to the Glue job's GlueVersion, JobType and code.
//
// Example:
//   var bucket bucket
//
//   glue.NewJob(this, jsii.String("PythonShellJob"), &jobProps{
//   	executable: glue.jobExecutable.pythonShell(&pythonShellExecutableProps{
//   		glueVersion: glue.glueVersion_V1_0(),
//   		pythonVersion: glue.pythonVersion_THREE,
//   		script: glue.code.fromBucket(bucket, jsii.String("script.py")),
//   	}),
//   	description: jsii.String("an example Python Shell job"),
//   })
//
// Experimental.
type JobExecutable interface {
	// Called during Job initialization to get JobExecutableConfig.
	// Experimental.
	Bind() *JobExecutableConfig
}

// The jsii proxy struct for JobExecutable
type jsiiProxy_JobExecutable struct {
	_ byte // padding
}

// Create a custom JobExecutable.
// Experimental.
func JobExecutable_Of(config *JobExecutableConfig) JobExecutable {
	_init_.Initialize()

	var returns JobExecutable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobExecutable",
		"of",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Create Python executable props for Apache Spark ETL job.
// Experimental.
func JobExecutable_PythonEtl(props *PythonSparkJobExecutableProps) JobExecutable {
	_init_.Initialize()

	var returns JobExecutable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobExecutable",
		"pythonEtl",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create Python executable props for python shell jobs.
// Experimental.
func JobExecutable_PythonShell(props *PythonShellExecutableProps) JobExecutable {
	_init_.Initialize()

	var returns JobExecutable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobExecutable",
		"pythonShell",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create Python executable props for Apache Spark Streaming job.
// Experimental.
func JobExecutable_PythonStreaming(props *PythonSparkJobExecutableProps) JobExecutable {
	_init_.Initialize()

	var returns JobExecutable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobExecutable",
		"pythonStreaming",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create Scala executable props for Apache Spark ETL job.
// Experimental.
func JobExecutable_ScalaEtl(props *ScalaJobExecutableProps) JobExecutable {
	_init_.Initialize()

	var returns JobExecutable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobExecutable",
		"scalaEtl",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create Scala executable props for Apache Spark Streaming job.
// Experimental.
func JobExecutable_ScalaStreaming(props *ScalaJobExecutableProps) JobExecutable {
	_init_.Initialize()

	var returns JobExecutable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobExecutable",
		"scalaStreaming",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobExecutable) Bind() *JobExecutableConfig {
	var returns *JobExecutableConfig

	_jsii_.Invoke(
		j,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Result of binding a `JobExecutable` into a `Job`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   var code code
//   var glueVersion glueVersion
//   var jobType jobType
//
//   jobExecutableConfig := &jobExecutableConfig{
//   	glueVersion: glueVersion,
//   	language: glue_alpha.jobLanguage_SCALA,
//   	script: code,
//   	type: jobType,
//
//   	// the properties below are optional
//   	className: jsii.String("className"),
//   	extraFiles: []*code{
//   		code,
//   	},
//   	extraJars: []*code{
//   		code,
//   	},
//   	extraJarsFirst: jsii.Boolean(false),
//   	extraPythonFiles: []*code{
//   		code,
//   	},
//   	pythonVersion: glue_alpha.pythonVersion_TWO,
//   }
//
// Experimental.
type JobExecutableConfig struct {
	// Glue version.
	// See: https://docs.aws.amazon.com/glue/latest/dg/release-notes.html
	//
	// Experimental.
	GlueVersion GlueVersion `field:"required" json:"glueVersion" yaml:"glueVersion"`
	// The language of the job (Scala or Python).
	// See: `--job-language` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	Language JobLanguage `field:"required" json:"language" yaml:"language"`
	// The script that is executed by a job.
	// Experimental.
	Script Code `field:"required" json:"script" yaml:"script"`
	// Specify the type of the job whether it's an Apache Spark ETL or streaming one or if it's a Python shell job.
	// Experimental.
	Type JobType `field:"required" json:"type" yaml:"type"`
	// The Scala class that serves as the entry point for the job.
	//
	// This applies only if your the job langauage is Scala.
	// See: `--class` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	// See: `--extra-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraFiles *[]Code `field:"optional" json:"extraFiles" yaml:"extraFiles"`
	// Additional Java .jar files that AWS Glue adds to the Java classpath before executing your script.
	// See: `--extra-jars` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJars *[]Code `field:"optional" json:"extraJars" yaml:"extraJars"`
	// Setting this value to true prioritizes the customer's extra JAR files in the classpath.
	// See: `--user-jars-first` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJarsFirst *bool `field:"optional" json:"extraJarsFirst" yaml:"extraJarsFirst"`
	// Additional Python files that AWS Glue adds to the Python path before executing your script.
	// See: `--extra-py-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraPythonFiles *[]Code `field:"optional" json:"extraPythonFiles" yaml:"extraPythonFiles"`
	// The Python version to use.
	// Experimental.
	PythonVersion PythonVersion `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
}

// Runtime language of the Glue job.
// Experimental.
type JobLanguage string

const (
	// Scala.
	// Experimental.
	JobLanguage_SCALA JobLanguage = "SCALA"
	// Python.
	// Experimental.
	JobLanguage_PYTHON JobLanguage = "PYTHON"
)

// Construction properties for {@link Job}.
//
// Example:
//   var bucket bucket
//
//   glue.NewJob(this, jsii.String("PythonShellJob"), &jobProps{
//   	executable: glue.jobExecutable.pythonShell(&pythonShellExecutableProps{
//   		glueVersion: glue.glueVersion_V1_0(),
//   		pythonVersion: glue.pythonVersion_THREE,
//   		script: glue.code.fromBucket(bucket, jsii.String("script.py")),
//   	}),
//   	description: jsii.String("an example Python Shell job"),
//   })
//
// Experimental.
type JobProps struct {
	// The job's executable properties.
	// Experimental.
	Executable JobExecutable `field:"required" json:"executable" yaml:"executable"`
	// The {@link Connection}s used for this job.
	//
	// Connections are used to connect to other AWS Service or resources within a VPC.
	// Experimental.
	Connections *[]IConnection `field:"optional" json:"connections" yaml:"connections"`
	// Enables continuous logging with the specified props.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ContinuousLogging *ContinuousLoggingProps `field:"optional" json:"continuousLogging" yaml:"continuousLogging"`
	// The default arguments for this job, specified as name-value pairs.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html for a list of reserved parameters
	//
	// Experimental.
	DefaultArguments *map[string]*string `field:"optional" json:"defaultArguments" yaml:"defaultArguments"`
	// The description of the job.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enables the collection of metrics for job profiling.
	// See: `--enable-metrics` at https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	EnableProfilingMetrics *bool `field:"optional" json:"enableProfilingMetrics" yaml:"enableProfilingMetrics"`
	// The name of the job.
	// Experimental.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// The number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.
	//
	// Cannot be used for Glue version 2.0 and later - workerType and workerCount should be used instead.
	// Experimental.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The maximum number of concurrent runs allowed for the job.
	//
	// An error is returned when this threshold is reached. The maximum value you can specify is controlled by a service limit.
	// Experimental.
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
	// The maximum number of times to retry this job after a job run fails.
	// Experimental.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// The number of minutes to wait after a job run starts, before sending a job run delay notification.
	// Experimental.
	NotifyDelayAfter awscdk.Duration `field:"optional" json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
	// The IAM role assumed by Glue to run this job.
	//
	// If providing a custom role, it needs to trust the Glue service principal (glue.amazonaws.com) and be granted sufficient permissions.
	// See: https://docs.aws.amazon.com/glue/latest/dg/getting-started-access.html
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The {@link SecurityConfiguration} to use for this job.
	// Experimental.
	SecurityConfiguration ISecurityConfiguration `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// Enables the Spark UI debugging and monitoring with the specified props.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	SparkUI *SparkUIProps `field:"optional" json:"sparkUI" yaml:"sparkUI"`
	// The tags to add to the resources on which the job runs.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The number of workers of a defined {@link WorkerType} that are allocated when a job runs.
	// Experimental.
	WorkerCount *float64 `field:"optional" json:"workerCount" yaml:"workerCount"`
	// The type of predefined worker that is allocated when a job runs.
	// Experimental.
	WorkerType WorkerType `field:"optional" json:"workerType" yaml:"workerType"`
}

// Job states emitted by Glue to CloudWatch Events.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types for more information.
//
// Experimental.
type JobState string

const (
	// State indicating job run succeeded.
	// Experimental.
	JobState_SUCCEEDED JobState = "SUCCEEDED"
	// State indicating job run failed.
	// Experimental.
	JobState_FAILED JobState = "FAILED"
	// State indicating job run timed out.
	// Experimental.
	JobState_TIMEOUT JobState = "TIMEOUT"
	// State indicating job is starting.
	// Experimental.
	JobState_STARTING JobState = "STARTING"
	// State indicating job is running.
	// Experimental.
	JobState_RUNNING JobState = "RUNNING"
	// State indicating job is stopping.
	// Experimental.
	JobState_STOPPING JobState = "STOPPING"
	// State indicating job stopped.
	// Experimental.
	JobState_STOPPED JobState = "STOPPED"
)

// The job type.
//
// If you need to use a JobType that doesn't exist as a static member, you
// can instantiate a `JobType` object, e.g: `JobType.of('other name')`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   jobType := glue_alpha.jobType_ETL()
//
// Experimental.
type JobType interface {
	// The name of this JobType, as expected by Job resource.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for JobType
type jsiiProxy_JobType struct {
	_ byte // padding
}

func (j *jsiiProxy_JobType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Custom type name.
// Experimental.
func JobType_Of(name *string) JobType {
	_init_.Initialize()

	var returns JobType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobType",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func JobType_ETL() JobType {
	_init_.Initialize()
	var returns JobType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.JobType",
		"ETL",
		&returns,
	)
	return returns
}

func JobType_PYTHON_SHELL() JobType {
	_init_.Initialize()
	var returns JobType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.JobType",
		"PYTHON_SHELL",
		&returns,
	)
	return returns
}

func JobType_STREAMING() JobType {
	_init_.Initialize()
	var returns JobType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.JobType",
		"STREAMING",
		&returns,
	)
	return returns
}

// The Glue CloudWatch metric type.
// See: https://docs.aws.amazon.com/glue/latest/dg/monitoring-awsglue-with-cloudwatch-metrics.html
//
// Experimental.
type MetricType string

const (
	// A value at a point in time.
	// Experimental.
	MetricType_GAUGE MetricType = "GAUGE"
	// An aggregate number.
	// Experimental.
	MetricType_COUNT MetricType = "COUNT"
)

// Absolute class name of the Hadoop `OutputFormat` to use when writing table files.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   outputFormat := glue_alpha.NewOutputFormat(jsii.String("className"))
//
// Experimental.
type OutputFormat interface {
	// Experimental.
	ClassName() *string
}

// The jsii proxy struct for OutputFormat
type jsiiProxy_OutputFormat struct {
	_ byte // padding
}

func (j *jsiiProxy_OutputFormat) ClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"className",
		&returns,
	)
	return returns
}


// Experimental.
func NewOutputFormat(className *string) OutputFormat {
	_init_.Initialize()

	j := jsiiProxy_OutputFormat{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.OutputFormat",
		[]interface{}{className},
		&j,
	)

	return &j
}

// Experimental.
func NewOutputFormat_Override(o OutputFormat, className *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.OutputFormat",
		[]interface{}{className},
		o,
	)
}

func OutputFormat_AVRO() InputFormat {
	_init_.Initialize()
	var returns InputFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.OutputFormat",
		"AVRO",
		&returns,
	)
	return returns
}

func OutputFormat_HIVE_IGNORE_KEY_TEXT() OutputFormat {
	_init_.Initialize()
	var returns OutputFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.OutputFormat",
		"HIVE_IGNORE_KEY_TEXT",
		&returns,
	)
	return returns
}

func OutputFormat_ORC() InputFormat {
	_init_.Initialize()
	var returns InputFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.OutputFormat",
		"ORC",
		&returns,
	)
	return returns
}

func OutputFormat_PARQUET() OutputFormat {
	_init_.Initialize()
	var returns OutputFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.OutputFormat",
		"PARQUET",
		&returns,
	)
	return returns
}

// Properties of a Partition Index.
//
// Example:
//   var myTable table
//
//   myTable.addPartitionIndex(&partitionIndex{
//   	indexName: jsii.String("my-index"),
//   	keyNames: []*string{
//   		jsii.String("year"),
//   	},
//   })
//
// Experimental.
type PartitionIndex struct {
	// The partition key names that comprise the partition index.
	//
	// The names must correspond to a name in the
	// table's partition keys.
	// Experimental.
	KeyNames *[]*string `field:"required" json:"keyNames" yaml:"keyNames"`
	// The name of the partition index.
	// Experimental.
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
}

// Props for creating a Python shell job executable.
//
// Example:
//   var bucket bucket
//
//   glue.NewJob(this, jsii.String("PythonShellJob"), &jobProps{
//   	executable: glue.jobExecutable.pythonShell(&pythonShellExecutableProps{
//   		glueVersion: glue.glueVersion_V1_0(),
//   		pythonVersion: glue.pythonVersion_THREE,
//   		script: glue.code.fromBucket(bucket, jsii.String("script.py")),
//   	}),
//   	description: jsii.String("an example Python Shell job"),
//   })
//
// Experimental.
type PythonShellExecutableProps struct {
	// Glue version.
	// See: https://docs.aws.amazon.com/glue/latest/dg/release-notes.html
	//
	// Experimental.
	GlueVersion GlueVersion `field:"required" json:"glueVersion" yaml:"glueVersion"`
	// The Python version to use.
	// Experimental.
	PythonVersion PythonVersion `field:"required" json:"pythonVersion" yaml:"pythonVersion"`
	// The script that executes a job.
	// Experimental.
	Script Code `field:"required" json:"script" yaml:"script"`
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	//
	// Only individual files are supported, directories are not supported.
	// See: `--extra-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraFiles *[]Code `field:"optional" json:"extraFiles" yaml:"extraFiles"`
	// Additional Python files that AWS Glue adds to the Python path before executing your script.
	//
	// Only individual files are supported, directories are not supported.
	// See: `--extra-py-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraPythonFiles *[]Code `field:"optional" json:"extraPythonFiles" yaml:"extraPythonFiles"`
}

// Props for creating a Python Spark (ETL or Streaming) job executable.
//
// Example:
//   glue.NewJob(this, jsii.String("PythonSparkStreamingJob"), &jobProps{
//   	executable: glue.jobExecutable.pythonStreaming(&pythonSparkJobExecutableProps{
//   		glueVersion: glue.glueVersion_V2_0(),
//   		pythonVersion: glue.pythonVersion_THREE,
//   		script: glue.code.fromAsset(path.join(__dirname, jsii.String("job-script/hello_world.py"))),
//   	}),
//   	description: jsii.String("an example Python Streaming job"),
//   })
//
// Experimental.
type PythonSparkJobExecutableProps struct {
	// Glue version.
	// See: https://docs.aws.amazon.com/glue/latest/dg/release-notes.html
	//
	// Experimental.
	GlueVersion GlueVersion `field:"required" json:"glueVersion" yaml:"glueVersion"`
	// The Python version to use.
	// Experimental.
	PythonVersion PythonVersion `field:"required" json:"pythonVersion" yaml:"pythonVersion"`
	// The script that executes a job.
	// Experimental.
	Script Code `field:"required" json:"script" yaml:"script"`
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	//
	// Only individual files are supported, directories are not supported.
	// See: `--extra-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraFiles *[]Code `field:"optional" json:"extraFiles" yaml:"extraFiles"`
	// Additional Java .jar files that AWS Glue adds to the Java classpath before executing your script. Only individual files are supported, directories are not supported.
	// See: `--extra-jars` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJars *[]Code `field:"optional" json:"extraJars" yaml:"extraJars"`
	// Setting this value to true prioritizes the customer's extra JAR files in the classpath.
	// See: `--user-jars-first` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJarsFirst *bool `field:"optional" json:"extraJarsFirst" yaml:"extraJarsFirst"`
	// Additional Python files that AWS Glue adds to the Python path before executing your script.
	//
	// Only individual files are supported, directories are not supported.
	// See: `--extra-py-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraPythonFiles *[]Code `field:"optional" json:"extraPythonFiles" yaml:"extraPythonFiles"`
}

// Python version.
//
// Example:
//   glue.NewJob(this, jsii.String("PythonSparkStreamingJob"), &jobProps{
//   	executable: glue.jobExecutable.pythonStreaming(&pythonSparkJobExecutableProps{
//   		glueVersion: glue.glueVersion_V2_0(),
//   		pythonVersion: glue.pythonVersion_THREE,
//   		script: glue.code.fromAsset(path.join(__dirname, jsii.String("job-script/hello_world.py"))),
//   	}),
//   	description: jsii.String("an example Python Streaming job"),
//   })
//
// Experimental.
type PythonVersion string

const (
	// Python 2 (the exact version depends on GlueVersion and JobCommand used).
	// Experimental.
	PythonVersion_TWO PythonVersion = "TWO"
	// Python 3 (the exact version depends on GlueVersion and JobCommand used).
	// Experimental.
	PythonVersion_THREE PythonVersion = "THREE"
)

// Glue job Code from an S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   s3Code := glue_alpha.NewS3Code(bucket, jsii.String("key"))
//
// Experimental.
type S3Code interface {
	Code
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(_scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig
}

// The jsii proxy struct for S3Code
type jsiiProxy_S3Code struct {
	jsiiProxy_Code
}

// Experimental.
func NewS3Code(bucket awss3.IBucket, key *string) S3Code {
	_init_.Initialize()

	j := jsiiProxy_S3Code{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.S3Code",
		[]interface{}{bucket, key},
		&j,
	)

	return &j
}

// Experimental.
func NewS3Code_Override(s S3Code, bucket awss3.IBucket, key *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.S3Code",
		[]interface{}{bucket, key},
		s,
	)
}

// Job code from a local disk path.
// Experimental.
func S3Code_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.S3Code",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Job code as an S3 object.
// Experimental.
func S3Code_FromBucket(bucket awss3.IBucket, key *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.S3Code",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Code) Bind(_scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, grantable},
		&returns,
	)

	return returns
}

// S3 encryption configuration.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &securityConfigurationProps{
//   	securityConfigurationName: jsii.String("name"),
//   	cloudWatchEncryption: &cloudWatchEncryption{
//   		mode: glue.cloudWatchEncryptionMode_KMS,
//   	},
//   	jobBookmarksEncryption: &jobBookmarksEncryption{
//   		mode: glue.jobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
//   	},
//   	s3Encryption: &s3Encryption{
//   		mode: glue.s3EncryptionMode_KMS,
//   	},
//   })
//
// Experimental.
type S3Encryption struct {
	// Encryption mode.
	// Experimental.
	Mode S3EncryptionMode `field:"required" json:"mode" yaml:"mode"`
	// The KMS key to be used to encrypt the data.
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

// Encryption mode for S3.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &securityConfigurationProps{
//   	securityConfigurationName: jsii.String("name"),
//   	cloudWatchEncryption: &cloudWatchEncryption{
//   		mode: glue.cloudWatchEncryptionMode_KMS,
//   	},
//   	jobBookmarksEncryption: &jobBookmarksEncryption{
//   		mode: glue.jobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
//   	},
//   	s3Encryption: &s3Encryption{
//   		mode: glue.s3EncryptionMode_KMS,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/glue/latest/webapi/API_S3Encryption.html#Glue-Type-S3Encryption-S3EncryptionMode
//
// Experimental.
type S3EncryptionMode string

const (
	// Server side encryption (SSE) with an Amazon S3-managed key.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html
	//
	// Experimental.
	S3EncryptionMode_S3_MANAGED S3EncryptionMode = "S3_MANAGED"
	// Server-side encryption (SSE) with an AWS KMS key managed by the account owner.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html
	//
	// Experimental.
	S3EncryptionMode_KMS S3EncryptionMode = "KMS"
)

// Props for creating a Scala Spark (ETL or Streaming) job executable.
//
// Example:
//   var bucket bucket
//
//   glue.NewJob(this, jsii.String("ScalaSparkEtlJob"), &jobProps{
//   	executable: glue.jobExecutable.scalaEtl(&scalaJobExecutableProps{
//   		glueVersion: glue.glueVersion_V2_0(),
//   		script: glue.code.fromBucket(bucket, jsii.String("src/com/example/HelloWorld.scala")),
//   		className: jsii.String("com.example.HelloWorld"),
//   		extraJars: []*code{
//   			glue.*code.fromBucket(bucket, jsii.String("jars/HelloWorld.jar")),
//   		},
//   	}),
//   	description: jsii.String("an example Scala ETL job"),
//   })
//
// Experimental.
type ScalaJobExecutableProps struct {
	// The fully qualified Scala class name that serves as the entry point for the job.
	// See: `--class` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ClassName *string `field:"required" json:"className" yaml:"className"`
	// Glue version.
	// See: https://docs.aws.amazon.com/glue/latest/dg/release-notes.html
	//
	// Experimental.
	GlueVersion GlueVersion `field:"required" json:"glueVersion" yaml:"glueVersion"`
	// The script that executes a job.
	// Experimental.
	Script Code `field:"required" json:"script" yaml:"script"`
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	//
	// Only individual files are supported, directories are not supported.
	// See: `--extra-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraFiles *[]Code `field:"optional" json:"extraFiles" yaml:"extraFiles"`
	// Additional Java .jar files that AWS Glue adds to the Java classpath before executing your script. Only individual files are supported, directories are not supported.
	// See: `--extra-jars` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJars *[]Code `field:"optional" json:"extraJars" yaml:"extraJars"`
	// Setting this value to true prioritizes the customer's extra JAR files in the classpath.
	// See: `--user-jars-first` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJarsFirst *bool `field:"optional" json:"extraJarsFirst" yaml:"extraJarsFirst"`
}

// Example:
//   var myDatabase database
//
//   glue.NewTable(this, jsii.String("MyTable"), &tableProps{
//   	database: myDatabase,
//   	tableName: jsii.String("my_table"),
//   	columns: []column{
//   		&column{
//   			name: jsii.String("col1"),
//   			type: glue.schema_STRING(),
//   		},
//   	},
//   	partitionKeys: []*column{
//   		&column{
//   			name: jsii.String("year"),
//   			type: glue.*schema_SMALL_INT(),
//   		},
//   		&column{
//   			name: jsii.String("month"),
//   			type: glue.*schema_SMALL_INT(),
//   		},
//   	},
//   	dataFormat: glue.dataFormat_JSON(),
//   	enablePartitionFiltering: jsii.Boolean(true),
//   })
//
// See: https://docs.aws.amazon.com/athena/latest/ug/data-types.html
//
// Experimental.
type Schema interface {
}

// The jsii proxy struct for Schema
type jsiiProxy_Schema struct {
	_ byte // padding
}

// Experimental.
func NewSchema() Schema {
	_init_.Initialize()

	j := jsiiProxy_Schema{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.Schema",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSchema_Override(s Schema) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.Schema",
		nil, // no parameters
		s,
	)
}

// Creates an array of some other type.
// Experimental.
func Schema_Array(itemType *Type) *Type {
	_init_.Initialize()

	var returns *Type

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Schema",
		"array",
		[]interface{}{itemType},
		&returns,
	)

	return returns
}

// Fixed length character data, with a specified length between 1 and 255.
// Experimental.
func Schema_Char(length *float64) *Type {
	_init_.Initialize()

	var returns *Type

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Schema",
		"char",
		[]interface{}{length},
		&returns,
	)

	return returns
}

// Creates a decimal type.
//
// TODO: Bounds.
// Experimental.
func Schema_Decimal(precision *float64, scale *float64) *Type {
	_init_.Initialize()

	var returns *Type

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Schema",
		"decimal",
		[]interface{}{precision, scale},
		&returns,
	)

	return returns
}

// Creates a map of some primitive key type to some value type.
// Experimental.
func Schema_Map(keyType *Type, valueType *Type) *Type {
	_init_.Initialize()

	var returns *Type

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Schema",
		"map",
		[]interface{}{keyType, valueType},
		&returns,
	)

	return returns
}

// Creates a nested structure containing individually named and typed columns.
// Experimental.
func Schema_Struct(columns *[]*Column) *Type {
	_init_.Initialize()

	var returns *Type

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Schema",
		"struct",
		[]interface{}{columns},
		&returns,
	)

	return returns
}

// Variable length character data, with a specified length between 1 and 65535.
// Experimental.
func Schema_Varchar(length *float64) *Type {
	_init_.Initialize()

	var returns *Type

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Schema",
		"varchar",
		[]interface{}{length},
		&returns,
	)

	return returns
}

func Schema_BIG_INT() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.Schema",
		"BIG_INT",
		&returns,
	)
	return returns
}

func Schema_BINARY() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.Schema",
		"BINARY",
		&returns,
	)
	return returns
}

func Schema_BOOLEAN() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.Schema",
		"BOOLEAN",
		&returns,
	)
	return returns
}

func Schema_DATE() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.Schema",
		"DATE",
		&returns,
	)
	return returns
}

func Schema_DOUBLE() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.Schema",
		"DOUBLE",
		&returns,
	)
	return returns
}

func Schema_FLOAT() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.Schema",
		"FLOAT",
		&returns,
	)
	return returns
}

func Schema_INTEGER() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.Schema",
		"INTEGER",
		&returns,
	)
	return returns
}

func Schema_SMALL_INT() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.Schema",
		"SMALL_INT",
		&returns,
	)
	return returns
}

func Schema_STRING() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.Schema",
		"STRING",
		&returns,
	)
	return returns
}

func Schema_TIMESTAMP() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.Schema",
		"TIMESTAMP",
		&returns,
	)
	return returns
}

func Schema_TINY_INT() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.Schema",
		"TINY_INT",
		&returns,
	)
	return returns
}

// A security configuration is a set of security properties that can be used by AWS Glue to encrypt data at rest.
//
// The following scenarios show some of the ways that you can use a security configuration.
// - Attach a security configuration to an AWS Glue crawler to write encrypted Amazon CloudWatch Logs.
// - Attach a security configuration to an extract, transform, and load (ETL) job to write encrypted Amazon Simple Storage Service (Amazon S3) targets and encrypted CloudWatch Logs.
// - Attach a security configuration to an ETL job to write its jobs bookmarks as encrypted Amazon S3 data.
// - Attach a security configuration to a development endpoint to write encrypted Amazon S3 targets.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &securityConfigurationProps{
//   	securityConfigurationName: jsii.String("name"),
//   	cloudWatchEncryption: &cloudWatchEncryption{
//   		mode: glue.cloudWatchEncryptionMode_KMS,
//   	},
//   	jobBookmarksEncryption: &jobBookmarksEncryption{
//   		mode: glue.jobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
//   	},
//   	s3Encryption: &s3Encryption{
//   		mode: glue.s3EncryptionMode_KMS,
//   	},
//   })
//
// Experimental.
type SecurityConfiguration interface {
	awscdk.Resource
	ISecurityConfiguration
	// The KMS key used in CloudWatch encryption if it requires a kms key.
	// Experimental.
	CloudWatchEncryptionKey() awskms.IKey
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
	// The KMS key used in job bookmarks encryption if it requires a kms key.
	// Experimental.
	JobBookmarksEncryptionKey() awskms.IKey
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The KMS key used in S3 encryption if it requires a kms key.
	// Experimental.
	S3EncryptionKey() awskms.IKey
	// The name of the security configuration.
	// Experimental.
	SecurityConfigurationName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityConfiguration
type jsiiProxy_SecurityConfiguration struct {
	internal.Type__awscdkResource
	jsiiProxy_ISecurityConfiguration
}

func (j *jsiiProxy_SecurityConfiguration) CloudWatchEncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"cloudWatchEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityConfiguration) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityConfiguration) JobBookmarksEncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"jobBookmarksEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityConfiguration) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityConfiguration) S3EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"s3EncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityConfiguration) SecurityConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewSecurityConfiguration(scope constructs.Construct, id *string, props *SecurityConfigurationProps) SecurityConfiguration {
	_init_.Initialize()

	j := jsiiProxy_SecurityConfiguration{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.SecurityConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSecurityConfiguration_Override(s SecurityConfiguration, scope constructs.Construct, id *string, props *SecurityConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.SecurityConfiguration",
		[]interface{}{scope, id, props},
		s,
	)
}

// Creates a Connection construct that represents an external security configuration.
// Experimental.
func SecurityConfiguration_FromSecurityConfigurationName(scope constructs.Construct, id *string, securityConfigurationName *string) ISecurityConfiguration {
	_init_.Initialize()

	var returns ISecurityConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.SecurityConfiguration",
		"fromSecurityConfigurationName",
		[]interface{}{scope, id, securityConfigurationName},
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
// Experimental.
func SecurityConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.SecurityConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func SecurityConfiguration_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.SecurityConfiguration",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func SecurityConfiguration_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.SecurityConfiguration",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_SecurityConfiguration) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityConfiguration) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityConfiguration) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Constructions properties of {@link SecurityConfiguration}.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &securityConfigurationProps{
//   	securityConfigurationName: jsii.String("name"),
//   	cloudWatchEncryption: &cloudWatchEncryption{
//   		mode: glue.cloudWatchEncryptionMode_KMS,
//   	},
//   	jobBookmarksEncryption: &jobBookmarksEncryption{
//   		mode: glue.jobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
//   	},
//   	s3Encryption: &s3Encryption{
//   		mode: glue.s3EncryptionMode_KMS,
//   	},
//   })
//
// Experimental.
type SecurityConfigurationProps struct {
	// The name of the security configuration.
	// Experimental.
	SecurityConfigurationName *string `field:"required" json:"securityConfigurationName" yaml:"securityConfigurationName"`
	// The encryption configuration for Amazon CloudWatch Logs.
	// Experimental.
	CloudWatchEncryption *CloudWatchEncryption `field:"optional" json:"cloudWatchEncryption" yaml:"cloudWatchEncryption"`
	// The encryption configuration for Glue Job Bookmarks.
	// Experimental.
	JobBookmarksEncryption *JobBookmarksEncryption `field:"optional" json:"jobBookmarksEncryption" yaml:"jobBookmarksEncryption"`
	// The encryption configuration for Amazon Simple Storage Service (Amazon S3) data.
	// Experimental.
	S3Encryption *S3Encryption `field:"optional" json:"s3Encryption" yaml:"s3Encryption"`
}

// Serialization library to use when serializing/deserializing (SerDe) table records.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   serializationLibrary := glue_alpha.serializationLibrary_AVRO()
//
// See: https://cwiki.apache.org/confluence/display/Hive/SerDe
//
// Experimental.
type SerializationLibrary interface {
	// Experimental.
	ClassName() *string
}

// The jsii proxy struct for SerializationLibrary
type jsiiProxy_SerializationLibrary struct {
	_ byte // padding
}

func (j *jsiiProxy_SerializationLibrary) ClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"className",
		&returns,
	)
	return returns
}


// Experimental.
func NewSerializationLibrary(className *string) SerializationLibrary {
	_init_.Initialize()

	j := jsiiProxy_SerializationLibrary{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		[]interface{}{className},
		&j,
	)

	return &j
}

// Experimental.
func NewSerializationLibrary_Override(s SerializationLibrary, className *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		[]interface{}{className},
		s,
	)
}

func SerializationLibrary_AVRO() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"AVRO",
		&returns,
	)
	return returns
}

func SerializationLibrary_CLOUDTRAIL() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"CLOUDTRAIL",
		&returns,
	)
	return returns
}

func SerializationLibrary_GROK() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"GROK",
		&returns,
	)
	return returns
}

func SerializationLibrary_HIVE_JSON() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"HIVE_JSON",
		&returns,
	)
	return returns
}

func SerializationLibrary_LAZY_SIMPLE() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"LAZY_SIMPLE",
		&returns,
	)
	return returns
}

func SerializationLibrary_OPEN_CSV() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"OPEN_CSV",
		&returns,
	)
	return returns
}

func SerializationLibrary_OPENX_JSON() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"OPENX_JSON",
		&returns,
	)
	return returns
}

func SerializationLibrary_ORC() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"ORC",
		&returns,
	)
	return returns
}

func SerializationLibrary_PARQUET() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"PARQUET",
		&returns,
	)
	return returns
}

func SerializationLibrary_REGEXP() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"REGEXP",
		&returns,
	)
	return returns
}

// The Spark UI logging location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   sparkUILoggingLocation := &sparkUILoggingLocation{
//   	bucket: bucket,
//
//   	// the properties below are optional
//   	prefix: jsii.String("prefix"),
//   }
//
// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
//
// Experimental.
type SparkUILoggingLocation struct {
	// The bucket where the Glue job stores the logs.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The path inside the bucket (objects prefix) where the Glue job stores the logs.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

// Properties for enabling Spark UI monitoring feature for Spark-based Glue jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   sparkUIProps := &sparkUIProps{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	bucket: bucket,
//   	prefix: jsii.String("prefix"),
//   }
//
// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
//
// Experimental.
type SparkUIProps struct {
	// Enable Spark UI.
	// Experimental.
	Enabled *bool `field:"required" json:"enabled" yaml:"enabled"`
	// The bucket where the Glue job stores the logs.
	// Experimental.
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// The path inside the bucket (objects prefix) where the Glue job stores the logs.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

// A Glue table.
//
// Example:
//   var myDatabase database
//
//   glue.NewTable(this, jsii.String("MyTable"), &tableProps{
//   	database: myDatabase,
//   	tableName: jsii.String("my_table"),
//   	columns: []column{
//   		&column{
//   			name: jsii.String("col1"),
//   			type: glue.schema_STRING(),
//   		},
//   	},
//   	partitionKeys: []*column{
//   		&column{
//   			name: jsii.String("year"),
//   			type: glue.*schema_SMALL_INT(),
//   		},
//   		&column{
//   			name: jsii.String("month"),
//   			type: glue.*schema_SMALL_INT(),
//   		},
//   	},
//   	dataFormat: glue.dataFormat_JSON(),
//   	enablePartitionFiltering: jsii.Boolean(true),
//   })
//
// Experimental.
type Table interface {
	awscdk.Resource
	ITable
	// S3 bucket in which the table's data resides.
	// Experimental.
	Bucket() awss3.IBucket
	// This table's columns.
	// Experimental.
	Columns() *[]*Column
	// Indicates whether the table's data is compressed or not.
	// Experimental.
	Compressed() *bool
	// Database this table belongs to.
	// Experimental.
	Database() IDatabase
	// Format of this table's data files.
	// Experimental.
	DataFormat() DataFormat
	// The type of encryption enabled for the table.
	// Experimental.
	Encryption() TableEncryption
	// The KMS key used to secure the data if `encryption` is set to `CSE-KMS` or `SSE-KMS`.
	//
	// Otherwise, `undefined`.
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
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// This table's partition indexes.
	// Experimental.
	PartitionIndexes() *[]*PartitionIndex
	// This table's partition keys if the table is partitioned.
	// Experimental.
	PartitionKeys() *[]*Column
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// S3 Key Prefix under which this table's files are stored in S3.
	// Experimental.
	S3Prefix() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// ARN of this table.
	// Experimental.
	TableArn() *string
	// Name of this table.
	// Experimental.
	TableName() *string
	// Add a partition index to the table.
	//
	// You can have a maximum of 3 partition
	// indexes to a table. Partition index keys must be a subset of the table's
	// partition keys.
	// See: https://docs.aws.amazon.com/glue/latest/dg/partition-indexes.html
	//
	// Experimental.
	AddPartitionIndex(index *PartitionIndex)
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
	// Grant the given identity custom permissions.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions *[]*string) awsiam.Grant
	// Grant read permissions to the table and the underlying data stored in S3 to an IAM principal.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grant read and write permissions to the table and the underlying data stored in S3 to an IAM principal.
	// Experimental.
	GrantReadWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity custom permissions to ALL underlying resources of the table.
	//
	// Permissions will be granted to the catalog, the database, and the table.
	// Experimental.
	GrantToUnderlyingResources(grantee awsiam.IGrantable, actions *[]*string) awsiam.Grant
	// Grant write permissions to the table and the underlying data stored in S3 to an IAM principal.
	// Experimental.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Table
type jsiiProxy_Table struct {
	internal.Type__awscdkResource
	jsiiProxy_ITable
}

func (j *jsiiProxy_Table) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) Columns() *[]*Column {
	var returns *[]*Column
	_jsii_.Get(
		j,
		"columns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) Compressed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"compressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) Database() IDatabase {
	var returns IDatabase
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) DataFormat() DataFormat {
	var returns DataFormat
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) Encryption() TableEncryption {
	var returns TableEncryption
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) PartitionIndexes() *[]*PartitionIndex {
	var returns *[]*PartitionIndex
	_jsii_.Get(
		j,
		"partitionIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) PartitionKeys() *[]*Column {
	var returns *[]*Column
	_jsii_.Get(
		j,
		"partitionKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) S3Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) TableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}


// Experimental.
func NewTable(scope constructs.Construct, id *string, props *TableProps) Table {
	_init_.Initialize()

	j := jsiiProxy_Table{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.Table",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTable_Override(t Table, scope constructs.Construct, id *string, props *TableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.Table",
		[]interface{}{scope, id, props},
		t,
	)
}

// Experimental.
func Table_FromTableArn(scope constructs.Construct, id *string, tableArn *string) ITable {
	_init_.Initialize()

	var returns ITable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Table",
		"fromTableArn",
		[]interface{}{scope, id, tableArn},
		&returns,
	)

	return returns
}

// Creates a Table construct that represents an external table.
// Experimental.
func Table_FromTableAttributes(scope constructs.Construct, id *string, attrs *TableAttributes) ITable {
	_init_.Initialize()

	var returns ITable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Table",
		"fromTableAttributes",
		[]interface{}{scope, id, attrs},
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
// Experimental.
func Table_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Table",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func Table_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Table",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Table_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Table",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Table) AddPartitionIndex(index *PartitionIndex) {
	_jsii_.InvokeVoid(
		t,
		"addPartitionIndex",
		[]interface{}{index},
	)
}

func (t *jsiiProxy_Table) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_Table) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Table) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Table) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Table) Grant(grantee awsiam.IGrantable, actions *[]*string) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grant",
		[]interface{}{grantee, actions},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Table) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Table) GrantReadWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grantReadWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Table) GrantToUnderlyingResources(grantee awsiam.IGrantable, actions *[]*string) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grantToUnderlyingResources",
		[]interface{}{grantee, actions},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Table) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Table) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   tableAttributes := &tableAttributes{
//   	tableArn: jsii.String("tableArn"),
//   	tableName: jsii.String("tableName"),
//   }
//
// Experimental.
type TableAttributes struct {
	// Experimental.
	TableArn *string `field:"required" json:"tableArn" yaml:"tableArn"`
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

// Encryption options for a Table.
//
// Example:
//   var myDatabase database
//
//   glue.NewTable(this, jsii.String("MyTable"), &tableProps{
//   	encryption: glue.tableEncryption_S3_MANAGED,
//   	// ...
//   	database: myDatabase,
//   	tableName: jsii.String("my_table"),
//   	columns: []column{
//   		&column{
//   			name: jsii.String("col1"),
//   			type: glue.schema_STRING(),
//   		},
//   	},
//   	dataFormat: glue.dataFormat_JSON(),
//   })
//
// See: https://docs.aws.amazon.com/athena/latest/ug/encryption.html
//
// Experimental.
type TableEncryption string

const (
	// Experimental.
	TableEncryption_UNENCRYPTED TableEncryption = "UNENCRYPTED"
	// Server side encryption (SSE) with an Amazon S3-managed key.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html
	//
	// Experimental.
	TableEncryption_S3_MANAGED TableEncryption = "S3_MANAGED"
	// Server-side encryption (SSE) with an AWS KMS key managed by the account owner.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html
	//
	// Experimental.
	TableEncryption_KMS TableEncryption = "KMS"
	// Server-side encryption (SSE) with an AWS KMS key managed by the KMS service.
	// Experimental.
	TableEncryption_KMS_MANAGED TableEncryption = "KMS_MANAGED"
	// Client-side encryption (CSE) with an AWS KMS key managed by the account owner.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html
	//
	// Experimental.
	TableEncryption_CLIENT_SIDE_KMS TableEncryption = "CLIENT_SIDE_KMS"
)

// Example:
//   var myDatabase database
//
//   glue.NewTable(this, jsii.String("MyTable"), &tableProps{
//   	database: myDatabase,
//   	tableName: jsii.String("my_table"),
//   	columns: []column{
//   		&column{
//   			name: jsii.String("col1"),
//   			type: glue.schema_STRING(),
//   		},
//   	},
//   	partitionKeys: []*column{
//   		&column{
//   			name: jsii.String("year"),
//   			type: glue.*schema_SMALL_INT(),
//   		},
//   		&column{
//   			name: jsii.String("month"),
//   			type: glue.*schema_SMALL_INT(),
//   		},
//   	},
//   	dataFormat: glue.dataFormat_JSON(),
//   	enablePartitionFiltering: jsii.Boolean(true),
//   })
//
// Experimental.
type TableProps struct {
	// Columns of the table.
	// Experimental.
	Columns *[]*Column `field:"required" json:"columns" yaml:"columns"`
	// Database in which to store the table.
	// Experimental.
	Database IDatabase `field:"required" json:"database" yaml:"database"`
	// Storage type of the table's data.
	// Experimental.
	DataFormat DataFormat `field:"required" json:"dataFormat" yaml:"dataFormat"`
	// Name of the table.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// S3 bucket in which to store data.
	// Experimental.
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// Indicates whether the table's data is compressed or not.
	// Experimental.
	Compressed *bool `field:"optional" json:"compressed" yaml:"compressed"`
	// Description of the table.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enables partition filtering.
	// See: https://docs.aws.amazon.com/athena/latest/ug/glue-best-practices.html#glue-best-practices-partition-index
	//
	// Experimental.
	EnablePartitionFiltering *bool `field:"optional" json:"enablePartitionFiltering" yaml:"enablePartitionFiltering"`
	// The kind of encryption to secure the data with.
	//
	// You can only provide this option if you are not explicitly passing in a bucket.
	//
	// If you choose `SSE-KMS`, you *can* provide an un-managed KMS key with `encryptionKey`.
	// If you choose `CSE-KMS`, you *must* provide an un-managed KMS key with `encryptionKey`.
	// Experimental.
	Encryption TableEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for bucket encryption.
	//
	// The `encryption` property must be `SSE-KMS` or `CSE-KMS`.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Partition indexes on the table.
	//
	// A maximum of 3 indexes
	// are allowed on a table. Keys in the index must be part
	// of the table's partition keys.
	// Experimental.
	PartitionIndexes *[]*PartitionIndex `field:"optional" json:"partitionIndexes" yaml:"partitionIndexes"`
	// Partition columns of the table.
	// Experimental.
	PartitionKeys *[]*Column `field:"optional" json:"partitionKeys" yaml:"partitionKeys"`
	// S3 prefix under which table objects are stored.
	// Experimental.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
	// Indicates whether the table data is stored in subdirectories.
	// Experimental.
	StoredAsSubDirectories *bool `field:"optional" json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
}

// Represents a type of a column in a table schema.
//
// Example:
//   var myDatabase database
//
//   glue.NewTable(this, jsii.String("MyTable"), &tableProps{
//   	database: myDatabase,
//   	tableName: jsii.String("my_table"),
//   	columns: []column{
//   		&column{
//   			name: jsii.String("col1"),
//   			type: glue.schema_STRING(),
//   		},
//   	},
//   	partitionKeys: []*column{
//   		&column{
//   			name: jsii.String("year"),
//   			type: glue.*schema_SMALL_INT(),
//   		},
//   		&column{
//   			name: jsii.String("month"),
//   			type: glue.*schema_SMALL_INT(),
//   		},
//   	},
//   	dataFormat: glue.dataFormat_JSON(),
//   	enablePartitionFiltering: jsii.Boolean(true),
//   })
//
// Experimental.
type Type struct {
	// Glue InputString for this type.
	// Experimental.
	InputString *string `field:"required" json:"inputString" yaml:"inputString"`
	// Indicates whether this type is a primitive data type.
	// Experimental.
	IsPrimitive *bool `field:"required" json:"isPrimitive" yaml:"isPrimitive"`
}

// The type of predefined worker that is allocated when a job runs.
//
// If you need to use a WorkerType that doesn't exist as a static member, you
// can instantiate a `WorkerType` object, e.g: `WorkerType.of('other type')`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   workerType := glue_alpha.workerType_G_1X()
//
// Experimental.
type WorkerType interface {
	// The name of this WorkerType, as expected by Job resource.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for WorkerType
type jsiiProxy_WorkerType struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkerType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Custom worker type.
// Experimental.
func WorkerType_Of(workerType *string) WorkerType {
	_init_.Initialize()

	var returns WorkerType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.WorkerType",
		"of",
		[]interface{}{workerType},
		&returns,
	)

	return returns
}

func WorkerType_G_1X() WorkerType {
	_init_.Initialize()
	var returns WorkerType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.WorkerType",
		"G_1X",
		&returns,
	)
	return returns
}

func WorkerType_G_2X() WorkerType {
	_init_.Initialize()
	var returns WorkerType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.WorkerType",
		"G_2X",
		&returns,
	)
	return returns
}

func WorkerType_STANDARD() WorkerType {
	_init_.Initialize()
	var returns WorkerType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.WorkerType",
		"STANDARD",
		&returns,
	)
	return returns
}

