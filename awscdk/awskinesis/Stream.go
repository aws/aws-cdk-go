package awskinesis

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Kinesis stream.
//
// Can be encrypted with a KMS key.
//
// Example:
//   lambdaRole := iam.NewRole(this, jsii.String("Role"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
//   	Description: jsii.String("Example role..."),
//   })
//
//   stream := kinesis.NewStream(this, jsii.String("MyEncryptedStream"), &StreamProps{
//   	Encryption: kinesis.StreamEncryption_KMS,
//   })
//
//   // give lambda permissions to read stream
//   stream.grantRead(lambdaRole)
//
type Stream interface {
	awscdk.Resource
	IStream
	// Indicates if a stream resource policy should automatically be created upon the first call to `addToResourcePolicy`.
	//
	// Set by subclasses.
	AutoCreatePolicy() *bool
	// Optional KMS encryption key associated with this stream.
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
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The ARN of the stream.
	StreamArn() *string
	// The name of the stream.
	StreamName() *string
	// Adds a statement to the IAM resource policy associated with this stream.
	//
	// If this stream was created in this stack (`new Stream`), a resource policy
	// will be automatically created upon the first call to `addToResourcePolicy`. If
	// the stream is imported (`Stream.import`), then this is a no-op.
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
	// Grant the indicated permissions on this stream to the given IAM principal (Role/Group/User).
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant read permissions for this stream and its contents to an IAM principal (Role/Group/User).
	//
	// If an encryption key is used, permission to ues the key to decrypt the
	// contents of the stream will also be granted.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grants read/write permissions for this stream and its contents to an IAM principal (Role/Group/User).
	//
	// If an encryption key is used, permission to use the key for
	// encrypt/decrypt will also be granted.
	GrantReadWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Grant write permissions for this stream and its contents to an IAM principal (Role/Group/User).
	//
	// If an encryption key is used, permission to ues the key to encrypt the
	// contents of the stream will also be granted.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Return stream metric based from its metric name.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of records retrieved from the shard, measured over the specified time period.
	//
	// Minimum, Maximum, and
	// Average statistics represent the records in a single GetRecords operation for the stream in the specified time
	// period.
	//
	// average
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricGetRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of bytes retrieved from the Kinesis stream, measured over the specified time period.
	//
	// Minimum, Maximum,
	// and Average statistics represent the bytes in a single GetRecords operation for the stream in the specified time
	// period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricGetRecordsBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The age of the last record in all GetRecords calls made against a Kinesis stream, measured over the specified time period.
	//
	// Age is the difference between the current time and when the last record of the GetRecords call was written
	// to the stream. The Minimum and Maximum statistics can be used to track the progress of Kinesis consumer
	// applications. A value of zero indicates that the records being read are completely caught up with the stream.
	//
	// The metric defaults to maximum over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricGetRecordsIteratorAgeMilliseconds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of successful GetRecords operations per stream, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricGetRecordsLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of successful GetRecords operations per stream, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricGetRecordsSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of bytes successfully put to the Kinesis stream over the specified time period.
	//
	// This metric includes
	// bytes from PutRecord and PutRecords operations. Minimum, Maximum, and Average statistics represent the bytes in a
	// single put operation for the stream in the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of records successfully put to the Kinesis stream over the specified time period.
	//
	// This metric includes
	// record counts from PutRecord and PutRecords operations. Minimum, Maximum, and Average statistics represent the
	// records in a single put operation for the stream in the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of bytes put to the Kinesis stream using the PutRecord operation over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time taken per PutRecord operation, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of bytes put to the Kinesis stream using the PutRecords operation over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordsBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of records rejected due to internal failures in a PutRecords operation per Kinesis data stream, measured over the specified time period.
	//
	// Occasional internal failures are to be expected and should be retried.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordsFailedRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time taken per PutRecords operation, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordsLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of PutRecords operations where at least one record succeeded, per Kinesis stream, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordsSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of successful records in a PutRecords operation per Kinesis data stream, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordsSuccessfulRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of records rejected due to throttling in a PutRecords operation per Kinesis data stream, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordsThrottledRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of records sent in a PutRecords operation per Kinesis data stream, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordsTotalRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of successful PutRecord operations per Kinesis stream, measured over the specified time period.
	//
	// Average
	// reflects the percentage of successful writes to a stream.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of GetRecords calls throttled for the stream over the specified time period.
	//
	// The most commonly used
	// statistic for this metric is Average.
	//
	// When the Minimum statistic has a value of 1, all records were throttled for the stream during the specified time
	// period.
	//
	// When the Maximum statistic has a value of 0 (zero), no records were throttled for the stream during the specified
	// time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricReadProvisionedThroughputExceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of records rejected due to throttling for the stream over the specified time period.
	//
	// This metric
	// includes throttling from PutRecord and PutRecords operations.
	//
	// When the Minimum statistic has a non-zero value, records were being throttled for the stream during the specified
	// time period.
	//
	// When the Maximum statistic has a value of 0 (zero), no records were being throttled for the stream during the
	// specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricWriteProvisionedThroughputExceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Stream
type jsiiProxy_Stream struct {
	internal.Type__awscdkResource
	jsiiProxy_IStream
}

func (j *jsiiProxy_Stream) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) StreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) StreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamName",
		&returns,
	)
	return returns
}


func NewStream(scope constructs.Construct, id *string, props *StreamProps) Stream {
	_init_.Initialize()

	if err := validateNewStreamParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Stream{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesis.Stream",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStream_Override(s Stream, scope constructs.Construct, id *string, props *StreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesis.Stream",
		[]interface{}{scope, id, props},
		s,
	)
}

// Import an existing Kinesis Stream provided an ARN.
func Stream_FromStreamArn(scope constructs.Construct, id *string, streamArn *string) IStream {
	_init_.Initialize()

	if err := validateStream_FromStreamArnParameters(scope, id, streamArn); err != nil {
		panic(err)
	}
	var returns IStream

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesis.Stream",
		"fromStreamArn",
		[]interface{}{scope, id, streamArn},
		&returns,
	)

	return returns
}

// Creates a Stream construct that represents an external stream.
func Stream_FromStreamAttributes(scope constructs.Construct, id *string, attrs *StreamAttributes) IStream {
	_init_.Initialize()

	if err := validateStream_FromStreamAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IStream

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesis.Stream",
		"fromStreamAttributes",
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
func Stream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStream_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesis.Stream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Stream_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateStream_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesis.Stream",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Stream_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateStream_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesis.Stream",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func Stream_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesis.Stream",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_Stream) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := s.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		s,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_Stream) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := s.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) GetResourceNameAttribute(nameAttr *string) *string {
	if err := s.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := s.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) GrantReadWrite(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateGrantReadWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantReadWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateGrantWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricGetRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricGetRecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricGetRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricGetRecordsBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricGetRecordsBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricGetRecordsBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricGetRecordsIteratorAgeMilliseconds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricGetRecordsIteratorAgeMillisecondsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricGetRecordsIteratorAgeMilliseconds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricGetRecordsLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricGetRecordsLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricGetRecordsLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricGetRecordsSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricGetRecordsSuccessParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricGetRecordsSuccess",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricIncomingBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricIncomingBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricIncomingRecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricIncomingRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricPutRecordBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricPutRecordBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricPutRecordBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricPutRecordLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricPutRecordLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricPutRecordLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricPutRecordsBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricPutRecordsBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricPutRecordsBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricPutRecordsFailedRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricPutRecordsFailedRecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricPutRecordsFailedRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricPutRecordsLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricPutRecordsLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricPutRecordsLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricPutRecordsSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricPutRecordsSuccessParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricPutRecordsSuccess",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricPutRecordsSuccessfulRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricPutRecordsSuccessfulRecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricPutRecordsSuccessfulRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricPutRecordsThrottledRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricPutRecordsThrottledRecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricPutRecordsThrottledRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricPutRecordsTotalRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricPutRecordsTotalRecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricPutRecordsTotalRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricPutRecordSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricPutRecordSuccessParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricPutRecordSuccess",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricReadProvisionedThroughputExceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricReadProvisionedThroughputExceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricReadProvisionedThroughputExceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) MetricWriteProvisionedThroughputExceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricWriteProvisionedThroughputExceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricWriteProvisionedThroughputExceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

