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
// TODO: EXAMPLE
//
// Experimental.
type AssetCode interface {
	Code
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

// Called when the Job is initialized to allow this object to bind.
// Experimental.
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
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/glue/latest/dg/add-classifier.html#classifier-built-in
//
// Experimental.
type ClassificationString interface {
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
// TODO: EXAMPLE
//
// Experimental.
type CloudWatchEncryption struct {
	// Encryption mode.
	// Experimental.
	Mode CloudWatchEncryptionMode `json:"mode"`
	// The KMS key to be used to encrypt the data.
	// Experimental.
	KmsKey awskms.IKey `json:"kmsKey"`
}

// Encryption mode for CloudWatch Logs.
// See: https://docs.aws.amazon.com/glue/latest/webapi/API_CloudWatchEncryption.html#Glue-Type-CloudWatchEncryption-CloudWatchEncryptionMode
//
// Experimental.
type CloudWatchEncryptionMode string

const (
	CloudWatchEncryptionMode_KMS CloudWatchEncryptionMode = "KMS"
)

// Represents a Glue Job's Code assets (an asset can be a scripts, a jar, a python file or any other file).
//
// TODO: EXAMPLE
//
// Experimental.
type Code interface {
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

// Called when the Job is initialized to allow this object to bind.
// Experimental.
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
// TODO: EXAMPLE
//
// Experimental.
type CodeConfig struct {
	// The location of the code in S3.
	// Experimental.
	S3Location *awss3.Location `json:"s3Location"`
}

// A column of a table.
//
// TODO: EXAMPLE
//
// Experimental.
type Column struct {
	// Coment describing the column.
	// Experimental.
	Comment *string `json:"comment"`
	// Name of the column.
	// Experimental.
	Name *string `json:"name"`
	// Type of the column.
	// Experimental.
	Type *Type `json:"type"`
}

// An AWS Glue connection to a data source.
//
// TODO: EXAMPLE
//
// Experimental.
type Connection interface {
	awscdk.Resource
	IConnection
	ConnectionArn() *string
	ConnectionName() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddProperty(key *string, value *string)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Add additional connection parameters.
// Experimental.
func (c *jsiiProxy_Connection) AddProperty(key *string, value *string) {
	_jsii_.InvokeVoid(
		c,
		"addProperty",
		[]interface{}{key, value},
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
func (c *jsiiProxy_Connection) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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
// TODO: EXAMPLE
//
// Experimental.
type ConnectionOptions struct {
	// The name of the connection.
	// Experimental.
	ConnectionName *string `json:"connectionName"`
	// The description of the connection.
	// Experimental.
	Description *string `json:"description"`
	// A list of criteria that can be used in selecting this connection.
	//
	// This is useful for filtering the results of https://awscli.amazonaws.com/v2/documentation/api/latest/reference/glue/get-connections.html
	// Experimental.
	MatchCriteria *[]*string `json:"matchCriteria"`
	// Key-Value pairs that define parameters for the connection.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-connect.html
	//
	// Experimental.
	Properties *map[string]*string `json:"properties"`
	// The list of security groups needed to successfully make this connection e.g. to successfully connect to VPC.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The VPC subnet to connect to resources within a VPC.
	//
	// See more at https://docs.aws.amazon.com/glue/latest/dg/start-connecting.html.
	// Experimental.
	Subnet awsec2.ISubnet `json:"subnet"`
}

// Construction properties for {@link Connection}.
//
// TODO: EXAMPLE
//
// Experimental.
type ConnectionProps struct {
	// The name of the connection.
	// Experimental.
	ConnectionName *string `json:"connectionName"`
	// The description of the connection.
	// Experimental.
	Description *string `json:"description"`
	// A list of criteria that can be used in selecting this connection.
	//
	// This is useful for filtering the results of https://awscli.amazonaws.com/v2/documentation/api/latest/reference/glue/get-connections.html
	// Experimental.
	MatchCriteria *[]*string `json:"matchCriteria"`
	// Key-Value pairs that define parameters for the connection.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-connect.html
	//
	// Experimental.
	Properties *map[string]*string `json:"properties"`
	// The list of security groups needed to successfully make this connection e.g. to successfully connect to VPC.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The VPC subnet to connect to resources within a VPC.
	//
	// See more at https://docs.aws.amazon.com/glue/latest/dg/start-connecting.html.
	// Experimental.
	Subnet awsec2.ISubnet `json:"subnet"`
	// The type of the connection.
	// Experimental.
	Type ConnectionType `json:"type"`
}

// The type of the glue connection.
//
// If you need to use a connection type that doesn't exist as a static member, you
// can instantiate a `ConnectionType` object, e.g: `new ConnectionType('NEW_TYPE')`.
//
// TODO: EXAMPLE
//
// Experimental.
type ConnectionType interface {
	Name() *string
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

// The connection type name as expected by Connection resource.
// Experimental.
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
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
//
// Experimental.
type ContinuousLoggingProps struct {
	// Apply the provided conversion pattern.
	//
	// This is a Log4j Conversion Pattern to customize driver and executor logs.
	// Experimental.
	ConversionPattern *string `json:"conversionPattern"`
	// Enable continouous logging.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// Specify a custom CloudWatch log group name.
	// Experimental.
	LogGroup awslogs.ILogGroup `json:"logGroup"`
	// Specify a custom CloudWatch log stream prefix.
	// Experimental.
	LogStreamPrefix *string `json:"logStreamPrefix"`
	// Filter out non-useful Apache Spark driver/executor and Apache Hadoop YARN heartbeat log messages.
	// Experimental.
	Quiet *bool `json:"quiet"`
}

// Defines the input/output formats and ser/de for a single DataFormat.
//
// TODO: EXAMPLE
//
// Experimental.
type DataFormat interface {
	ClassificationString() ClassificationString
	InputFormat() InputFormat
	OutputFormat() OutputFormat
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
// TODO: EXAMPLE
//
// Experimental.
type DataFormatProps struct {
	// Classification string given to tables with this data format.
	// Experimental.
	ClassificationString ClassificationString `json:"classificationString"`
	// `InputFormat` for this data format.
	// Experimental.
	InputFormat InputFormat `json:"inputFormat"`
	// `OutputFormat` for this data format.
	// Experimental.
	OutputFormat OutputFormat `json:"outputFormat"`
	// Serialization library for this data format.
	// Experimental.
	SerializationLibrary SerializationLibrary `json:"serializationLibrary"`
}

// A Glue database.
//
// TODO: EXAMPLE
//
// Experimental.
type Database interface {
	awscdk.Resource
	IDatabase
	CatalogArn() *string
	CatalogId() *string
	DatabaseArn() *string
	DatabaseName() *string
	Env() *awscdk.ResourceEnvironment
	LocationUri() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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
func (d *jsiiProxy_Database) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// TODO: EXAMPLE
//
// Experimental.
type DatabaseProps struct {
	// The name of the database.
	// Experimental.
	DatabaseName *string `json:"databaseName"`
	// The location of the database (for example, an HDFS path).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-databaseinput.html
	//
	// Experimental.
	LocationUri *string `json:"locationUri"`
}

// AWS Glue version determines the versions of Apache Spark and Python that are available to the job.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/glue/latest/dg/add-job.html.
//
// If you need to use a GlueVersion that doesn't exist as a static member, you
// can instantiate a `GlueVersion` object, e.g: `GlueVersion.of('1.5')`.
//
// Experimental.
type GlueVersion interface {
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
// TODO: EXAMPLE
//
// Experimental.
type InputFormat interface {
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
// TODO: EXAMPLE
//
// Experimental.
type Job interface {
	awscdk.Resource
	IJob
	Env() *awscdk.ResourceEnvironment
	GrantPrincipal() awsiam.IPrincipal
	JobArn() *string
	JobName() *string
	Node() constructs.Node
	PhysicalName() *string
	Role() awsiam.IRole
	SparkUILoggingLocation() *SparkUILoggingLocation
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Metric(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimeout(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnFailure(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnStateChange(id *string, jobState JobState, options *awsevents.OnEventOptions) awsevents.Rule
	OnSuccess(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnTimeout(id *string, options *awsevents.OnEventOptions) awsevents.Rule
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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
func (j *jsiiProxy_Job) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		j,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
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

// Create a CloudWatch metric.
// See: https://docs.aws.amazon.com/glue/latest/dg/monitoring-awsglue-with-cloudwatch-metrics.html
//
// Experimental.
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

// Return a CloudWatch Metric indicating job failure.
//
// This metric is based on the Rule returned by no-args onFailure() call.
// Experimental.
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

// Return a CloudWatch Metric indicating job success.
//
// This metric is based on the Rule returned by no-args onSuccess() call.
// Experimental.
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

// Return a CloudWatch Metric indicating job timeout.
//
// This metric is based on the Rule returned by no-args onTimeout() call.
// Experimental.
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

// Create a CloudWatch Event Rule for this Glue Job when it's in a given state.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types
//
// Experimental.
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

// Return a CloudWatch Event Rule matching FAILED state.
// Experimental.
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

// Create a CloudWatch Event Rule for the transition into the input jobState.
// Experimental.
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

// Create a CloudWatch Event Rule matching JobState.SUCCEEDED.
// Experimental.
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

// Return a CloudWatch Event Rule matching TIMEOUT state.
// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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
// TODO: EXAMPLE
//
// Experimental.
type JobAttributes struct {
	// The name of the job.
	// Experimental.
	JobName *string `json:"jobName"`
	// The IAM role assumed by Glue to run this job.
	// Experimental.
	Role awsiam.IRole `json:"role"`
}

// Job bookmarks encryption configuration.
//
// TODO: EXAMPLE
//
// Experimental.
type JobBookmarksEncryption struct {
	// Encryption mode.
	// Experimental.
	Mode JobBookmarksEncryptionMode `json:"mode"`
	// The KMS key to be used to encrypt the data.
	// Experimental.
	KmsKey awskms.IKey `json:"kmsKey"`
}

// Encryption mode for Job Bookmarks.
// See: https://docs.aws.amazon.com/glue/latest/webapi/API_JobBookmarksEncryption.html#Glue-Type-JobBookmarksEncryption-JobBookmarksEncryptionMode
//
// Experimental.
type JobBookmarksEncryptionMode string

const (
	JobBookmarksEncryptionMode_CLIENT_SIDE_KMS JobBookmarksEncryptionMode = "CLIENT_SIDE_KMS"
)

// The executable properties related to the Glue job's GlueVersion, JobType and code.
//
// TODO: EXAMPLE
//
// Experimental.
type JobExecutable interface {
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

// Called during Job initialization to get JobExecutableConfig.
// Experimental.
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
// TODO: EXAMPLE
//
// Experimental.
type JobExecutableConfig struct {
	// The Scala class that serves as the entry point for the job.
	//
	// This applies only if your the job langauage is Scala.
	// See: `--class` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ClassName *string `json:"className"`
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	// See: `--extra-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraFiles *[]Code `json:"extraFiles"`
	// Additional Java .jar files that AWS Glue adds to the Java classpath before executing your script.
	// See: `--extra-jars` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJars *[]Code `json:"extraJars"`
	// Setting this value to true prioritizes the customer's extra JAR files in the classpath.
	// See: `--user-jars-first` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJarsFirst *bool `json:"extraJarsFirst"`
	// Additional Python files that AWS Glue adds to the Python path before executing your script.
	// See: `--extra-py-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraPythonFiles *[]Code `json:"extraPythonFiles"`
	// Glue version.
	// See: https://docs.aws.amazon.com/glue/latest/dg/release-notes.html
	//
	// Experimental.
	GlueVersion GlueVersion `json:"glueVersion"`
	// The language of the job (Scala or Python).
	// See: `--job-language` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	Language JobLanguage `json:"language"`
	// The Python version to use.
	// Experimental.
	PythonVersion PythonVersion `json:"pythonVersion"`
	// The script that is executed by a job.
	// Experimental.
	Script Code `json:"script"`
	// Specify the type of the job whether it's an Apache Spark ETL or streaming one or if it's a Python shell job.
	// Experimental.
	Type JobType `json:"type"`
}

// Runtime language of the Glue job.
// Experimental.
type JobLanguage string

const (
	JobLanguage_PYTHON JobLanguage = "PYTHON"
	JobLanguage_SCALA JobLanguage = "SCALA"
)

// Construction properties for {@link Job}.
//
// TODO: EXAMPLE
//
// Experimental.
type JobProps struct {
	// The job's executable properties.
	// Experimental.
	Executable JobExecutable `json:"executable"`
	// The {@link Connection}s used for this job.
	//
	// Connections are used to connect to other AWS Service or resources within a VPC.
	// Experimental.
	Connections *[]IConnection `json:"connections"`
	// Enables continuous logging with the specified props.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ContinuousLogging *ContinuousLoggingProps `json:"continuousLogging"`
	// The default arguments for this job, specified as name-value pairs.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html for a list of reserved parameters
	//
	// Experimental.
	DefaultArguments *map[string]*string `json:"defaultArguments"`
	// The description of the job.
	// Experimental.
	Description *string `json:"description"`
	// Enables the collection of metrics for job profiling.
	// See: `--enable-metrics` at https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	EnableProfilingMetrics *bool `json:"enableProfilingMetrics"`
	// The name of the job.
	// Experimental.
	JobName *string `json:"jobName"`
	// The number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.
	//
	// Cannot be used for Glue version 2.0 and later - workerType and workerCount should be used instead.
	// Experimental.
	MaxCapacity *float64 `json:"maxCapacity"`
	// The maximum number of concurrent runs allowed for the job.
	//
	// An error is returned when this threshold is reached. The maximum value you can specify is controlled by a service limit.
	// Experimental.
	MaxConcurrentRuns *float64 `json:"maxConcurrentRuns"`
	// The maximum number of times to retry this job after a job run fails.
	// Experimental.
	MaxRetries *float64 `json:"maxRetries"`
	// The number of minutes to wait after a job run starts, before sending a job run delay notification.
	// Experimental.
	NotifyDelayAfter awscdk.Duration `json:"notifyDelayAfter"`
	// The IAM role assumed by Glue to run this job.
	//
	// If providing a custom role, it needs to trust the Glue service principal (glue.amazonaws.com) and be granted sufficient permissions.
	// See: https://docs.aws.amazon.com/glue/latest/dg/getting-started-access.html
	//
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The {@link SecurityConfiguration} to use for this job.
	// Experimental.
	SecurityConfiguration ISecurityConfiguration `json:"securityConfiguration"`
	// Enables the Spark UI debugging and monitoring with the specified props.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	SparkUI *SparkUIProps `json:"sparkUI"`
	// The tags to add to the resources on which the job runs.
	// Experimental.
	Tags *map[string]*string `json:"tags"`
	// The maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The number of workers of a defined {@link WorkerType} that are allocated when a job runs.
	// Experimental.
	WorkerCount *float64 `json:"workerCount"`
	// The type of predefined worker that is allocated when a job runs.
	// Experimental.
	WorkerType WorkerType `json:"workerType"`
}

// Job states emitted by Glue to CloudWatch Events.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types for more information.
//
// Experimental.
type JobState string

const (
	JobState_SUCCEEDED JobState = "SUCCEEDED"
	JobState_FAILED JobState = "FAILED"
	JobState_TIMEOUT JobState = "TIMEOUT"
	JobState_STARTING JobState = "STARTING"
	JobState_RUNNING JobState = "RUNNING"
	JobState_STOPPING JobState = "STOPPING"
	JobState_STOPPED JobState = "STOPPED"
)

// The job type.
//
// If you need to use a JobType that doesn't exist as a static member, you
// can instantiate a `JobType` object, e.g: `JobType.of('other name')`.
//
// TODO: EXAMPLE
//
// Experimental.
type JobType interface {
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
	MetricType_GAUGE MetricType = "GAUGE"
	MetricType_COUNT MetricType = "COUNT"
)

// Absolute class name of the Hadoop `OutputFormat` to use when writing table files.
//
// TODO: EXAMPLE
//
// Experimental.
type OutputFormat interface {
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

// Props for creating a Python shell job executable.
//
// TODO: EXAMPLE
//
// Experimental.
type PythonShellExecutableProps struct {
	// Glue version.
	// See: https://docs.aws.amazon.com/glue/latest/dg/release-notes.html
	//
	// Experimental.
	GlueVersion GlueVersion `json:"glueVersion"`
	// The Python version to use.
	// Experimental.
	PythonVersion PythonVersion `json:"pythonVersion"`
	// The script that executes a job.
	// Experimental.
	Script Code `json:"script"`
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	//
	// Only individual files are supported, directories are not supported.
	// See: `--extra-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraFiles *[]Code `json:"extraFiles"`
	// Additional Python files that AWS Glue adds to the Python path before executing your script.
	//
	// Only individual files are supported, directories are not supported.
	// See: `--extra-py-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraPythonFiles *[]Code `json:"extraPythonFiles"`
}

// Props for creating a Python Spark (ETL or Streaming) job executable.
//
// TODO: EXAMPLE
//
// Experimental.
type PythonSparkJobExecutableProps struct {
	// Glue version.
	// See: https://docs.aws.amazon.com/glue/latest/dg/release-notes.html
	//
	// Experimental.
	GlueVersion GlueVersion `json:"glueVersion"`
	// The Python version to use.
	// Experimental.
	PythonVersion PythonVersion `json:"pythonVersion"`
	// The script that executes a job.
	// Experimental.
	Script Code `json:"script"`
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	//
	// Only individual files are supported, directories are not supported.
	// See: `--extra-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraFiles *[]Code `json:"extraFiles"`
	// Additional Java .jar files that AWS Glue adds to the Java classpath before executing your script. Only individual files are supported, directories are not supported.
	// See: `--extra-jars` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJars *[]Code `json:"extraJars"`
	// Setting this value to true prioritizes the customer's extra JAR files in the classpath.
	// See: `--user-jars-first` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJarsFirst *bool `json:"extraJarsFirst"`
	// Additional Python files that AWS Glue adds to the Python path before executing your script.
	//
	// Only individual files are supported, directories are not supported.
	// See: `--extra-py-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraPythonFiles *[]Code `json:"extraPythonFiles"`
}

// Python version.
//
// TODO: EXAMPLE
//
// Experimental.
type PythonVersion string

const (
	PythonVersion_THREE PythonVersion = "THREE"
	PythonVersion_TWO PythonVersion = "TWO"
)

// Glue job Code from an S3 bucket.
//
// TODO: EXAMPLE
//
// Experimental.
type S3Code interface {
	Code
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

// Called when the Job is initialized to allow this object to bind.
// Experimental.
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
// TODO: EXAMPLE
//
// Experimental.
type S3Encryption struct {
	// Encryption mode.
	// Experimental.
	Mode S3EncryptionMode `json:"mode"`
	// The KMS key to be used to encrypt the data.
	// Experimental.
	KmsKey awskms.IKey `json:"kmsKey"`
}

// Encryption mode for S3.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/glue/latest/webapi/API_S3Encryption.html#Glue-Type-S3Encryption-S3EncryptionMode
//
// Experimental.
type S3EncryptionMode string

const (
	S3EncryptionMode_S3_MANAGED S3EncryptionMode = "S3_MANAGED"
	S3EncryptionMode_KMS S3EncryptionMode = "KMS"
)

// Props for creating a Scala Spark (ETL or Streaming) job executable.
//
// TODO: EXAMPLE
//
// Experimental.
type ScalaJobExecutableProps struct {
	// The fully qualified Scala class name that serves as the entry point for the job.
	// See: `--class` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ClassName *string `json:"className"`
	// Glue version.
	// See: https://docs.aws.amazon.com/glue/latest/dg/release-notes.html
	//
	// Experimental.
	GlueVersion GlueVersion `json:"glueVersion"`
	// The script that executes a job.
	// Experimental.
	Script Code `json:"script"`
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	//
	// Only individual files are supported, directories are not supported.
	// See: `--extra-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraFiles *[]Code `json:"extraFiles"`
	// Additional Java .jar files that AWS Glue adds to the Java classpath before executing your script. Only individual files are supported, directories are not supported.
	// See: `--extra-jars` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJars *[]Code `json:"extraJars"`
	// Setting this value to true prioritizes the customer's extra JAR files in the classpath.
	// See: `--user-jars-first` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJarsFirst *bool `json:"extraJarsFirst"`
}

// TODO: EXAMPLE
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
// TODO: Bounds
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
// TODO: EXAMPLE
//
// Experimental.
type SecurityConfiguration interface {
	awscdk.Resource
	ISecurityConfiguration
	CloudWatchEncryptionKey() awskms.IKey
	Env() *awscdk.ResourceEnvironment
	JobBookmarksEncryptionKey() awskms.IKey
	Node() constructs.Node
	PhysicalName() *string
	S3EncryptionKey() awskms.IKey
	SecurityConfigurationName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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
func (s *jsiiProxy_SecurityConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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
// TODO: EXAMPLE
//
// Experimental.
type SecurityConfigurationProps struct {
	// The name of the security configuration.
	// Experimental.
	SecurityConfigurationName *string `json:"securityConfigurationName"`
	// The encryption configuration for Amazon CloudWatch Logs.
	// Experimental.
	CloudWatchEncryption *CloudWatchEncryption `json:"cloudWatchEncryption"`
	// The encryption configuration for Glue Job Bookmarks.
	// Experimental.
	JobBookmarksEncryption *JobBookmarksEncryption `json:"jobBookmarksEncryption"`
	// The encryption configuration for Amazon Simple Storage Service (Amazon S3) data.
	// Experimental.
	S3Encryption *S3Encryption `json:"s3Encryption"`
}

// Serialization library to use when serializing/deserializing (SerDe) table records.
//
// TODO: EXAMPLE
//
// See: https://cwiki.apache.org/confluence/display/Hive/SerDe
//
// Experimental.
type SerializationLibrary interface {
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
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
//
// Experimental.
type SparkUILoggingLocation struct {
	// The bucket where the Glue job stores the logs.
	// Experimental.
	Bucket awss3.IBucket `json:"bucket"`
	// The path inside the bucket (objects prefix) where the Glue job stores the logs.
	// Experimental.
	Prefix *string `json:"prefix"`
}

// Properties for enabling Spark UI monitoring feature for Spark-based Glue jobs.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
//
// Experimental.
type SparkUIProps struct {
	// The bucket where the Glue job stores the logs.
	// Experimental.
	Bucket awss3.IBucket `json:"bucket"`
	// Enable Spark UI.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// The path inside the bucket (objects prefix) where the Glue job stores the logs.
	// Experimental.
	Prefix *string `json:"prefix"`
}

// A Glue table.
//
// TODO: EXAMPLE
//
// Experimental.
type Table interface {
	awscdk.Resource
	ITable
	Bucket() awss3.IBucket
	Columns() *[]*Column
	Compressed() *bool
	Database() IDatabase
	DataFormat() DataFormat
	Encryption() TableEncryption
	EncryptionKey() awskms.IKey
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PartitionKeys() *[]*Column
	PhysicalName() *string
	S3Prefix() *string
	Stack() awscdk.Stack
	TableArn() *string
	TableName() *string
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	GrantReadWrite(grantee awsiam.IGrantable) awsiam.Grant
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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
func (t *jsiiProxy_Table) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
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

// Grant read permissions to the table and the underlying data stored in S3 to an IAM principal.
// Experimental.
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

// Grant read and write permissions to the table and the underlying data stored in S3 to an IAM principal.
// Experimental.
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

// Grant write permissions to the table and the underlying data stored in S3 to an IAM principal.
// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// TODO: EXAMPLE
//
// Experimental.
type TableAttributes struct {
	// Experimental.
	TableArn *string `json:"tableArn"`
	// Experimental.
	TableName *string `json:"tableName"`
}

// Encryption options for a Table.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/athena/latest/ug/encryption.html
//
// Experimental.
type TableEncryption string

const (
	TableEncryption_UNENCRYPTED TableEncryption = "UNENCRYPTED"
	TableEncryption_S3_MANAGED TableEncryption = "S3_MANAGED"
	TableEncryption_KMS TableEncryption = "KMS"
	TableEncryption_KMS_MANAGED TableEncryption = "KMS_MANAGED"
	TableEncryption_CLIENT_SIDE_KMS TableEncryption = "CLIENT_SIDE_KMS"
)

// TODO: EXAMPLE
//
// Experimental.
type TableProps struct {
	// Columns of the table.
	// Experimental.
	Columns *[]*Column `json:"columns"`
	// Database in which to store the table.
	// Experimental.
	Database IDatabase `json:"database"`
	// Storage type of the table's data.
	// Experimental.
	DataFormat DataFormat `json:"dataFormat"`
	// Name of the table.
	// Experimental.
	TableName *string `json:"tableName"`
	// S3 bucket in which to store data.
	// Experimental.
	Bucket awss3.IBucket `json:"bucket"`
	// Indicates whether the table's data is compressed or not.
	// Experimental.
	Compressed *bool `json:"compressed"`
	// Description of the table.
	// Experimental.
	Description *string `json:"description"`
	// The kind of encryption to secure the data with.
	//
	// You can only provide this option if you are not explicitly passing in a bucket.
	//
	// If you choose `SSE-KMS`, you *can* provide an un-managed KMS key with `encryptionKey`.
	// If you choose `CSE-KMS`, you *must* provide an un-managed KMS key with `encryptionKey`.
	// Experimental.
	Encryption TableEncryption `json:"encryption"`
	// External KMS key to use for bucket encryption.
	//
	// The `encryption` property must be `SSE-KMS` or `CSE-KMS`.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// Partition columns of the table.
	// Experimental.
	PartitionKeys *[]*Column `json:"partitionKeys"`
	// S3 prefix under which table objects are stored.
	// Experimental.
	S3Prefix *string `json:"s3Prefix"`
	// Indicates whether the table data is stored in subdirectories.
	// Experimental.
	StoredAsSubDirectories *bool `json:"storedAsSubDirectories"`
}

// Represents a type of a column in a table schema.
//
// TODO: EXAMPLE
//
// Experimental.
type Type struct {
	// Glue InputString for this type.
	// Experimental.
	InputString *string `json:"inputString"`
	// Indicates whether this type is a primitive data type.
	// Experimental.
	IsPrimitive *bool `json:"isPrimitive"`
}

// The type of predefined worker that is allocated when a job runs.
//
// If you need to use a WorkerType that doesn't exist as a static member, you
// can instantiate a `WorkerType` object, e.g: `WorkerType.of('other type')`.
//
// TODO: EXAMPLE
//
// Experimental.
type WorkerType interface {
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

