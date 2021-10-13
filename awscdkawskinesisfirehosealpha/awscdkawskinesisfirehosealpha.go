// The CDK Construct Library for AWS::KinesisFirehose
package awscdkawskinesisfirehosealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkawskinesisfirehosealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdkawskinesisfirehosealpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Options when binding a DataProcessor to a delivery stream destination.
// Experimental.
type DataProcessorBindOptions struct {
	// The IAM role assumed by Kinesis Data Firehose to write to the destination that this DataProcessor will bind to.
	// Experimental.
	Role awsiam.IRole `json:"role"`
}

// The full configuration of a data processor.
// Experimental.
type DataProcessorConfig struct {
	// The key-value pair that identifies the underlying processor resource.
	// Experimental.
	ProcessorIdentifier *DataProcessorIdentifier `json:"processorIdentifier"`
	// The type of the underlying processor resource.
	//
	// Must be an accepted value in `CfnDeliveryStream.ProcessorProperty.Type`.
	//
	// TODO: EXAMPLE
	//
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processor.html#cfn-kinesisfirehose-deliverystream-processor-type
	//
	// Experimental.
	ProcessorType *string `json:"processorType"`
}

// The key-value pair that identifies the underlying processor resource.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processorparameter.html
//
// Experimental.
type DataProcessorIdentifier struct {
	// The parameter name that corresponds to the processor resource's identifier.
	//
	// Must be an accepted value in `CfnDeliveryStream.ProcessoryParameterProperty.ParameterName`.
	// Experimental.
	ParameterName *string `json:"parameterName"`
	// The identifier of the underlying processor resource.
	// Experimental.
	ParameterValue *string `json:"parameterValue"`
}

// Configure the data processor.
// Experimental.
type DataProcessorProps struct {
	// The length of time Kinesis Data Firehose will buffer incoming data before calling the processor.
	//
	// s
	// Experimental.
	BufferInterval awscdk.Duration `json:"bufferInterval"`
	// The amount of incoming data Kinesis Data Firehose will buffer before calling the processor.
	// Experimental.
	BufferSize awscdk.Size `json:"bufferSize"`
	// The number of times Kinesis Data Firehose will retry the processor invocation after a failure due to network timeout or invocation limits.
	// Experimental.
	Retries *float64 `json:"retries"`
}

// Create a Kinesis Data Firehose delivery stream.
// Experimental.
type DeliveryStream interface {
	awscdk.Resource
	IDeliveryStream
	Connections() awsec2.Connections
	DeliveryStreamArn() *string
	DeliveryStreamName() *string
	Env() *awscdk.ResourceEnvironment
	GrantPrincipal() awsiam.IPrincipal
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	GrantPutRecords(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricBackupToS3Bytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricBackupToS3DataFreshness(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricBackupToS3Records(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for DeliveryStream
type jsiiProxy_DeliveryStream struct {
	internal.Type__awscdkResource
	jsiiProxy_IDeliveryStream
}

func (j *jsiiProxy_DeliveryStream) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) DeliveryStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) DeliveryStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDeliveryStream(scope constructs.Construct, id *string, props *DeliveryStreamProps) DeliveryStream {
	_init_.Initialize()

	j := jsiiProxy_DeliveryStream{}

	_jsii_.Create(
		"@aws-cdk/aws-kinesisfirehose-alpha.DeliveryStream",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDeliveryStream_Override(d DeliveryStream, scope constructs.Construct, id *string, props *DeliveryStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-kinesisfirehose-alpha.DeliveryStream",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import an existing delivery stream from its ARN.
// Experimental.
func DeliveryStream_FromDeliveryStreamArn(scope constructs.Construct, id *string, deliveryStreamArn *string) IDeliveryStream {
	_init_.Initialize()

	var returns IDeliveryStream

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisfirehose-alpha.DeliveryStream",
		"fromDeliveryStreamArn",
		[]interface{}{scope, id, deliveryStreamArn},
		&returns,
	)

	return returns
}

// Import an existing delivery stream from its attributes.
// Experimental.
func DeliveryStream_FromDeliveryStreamAttributes(scope constructs.Construct, id *string, attrs *DeliveryStreamAttributes) IDeliveryStream {
	_init_.Initialize()

	var returns IDeliveryStream

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisfirehose-alpha.DeliveryStream",
		"fromDeliveryStreamAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing delivery stream from its name.
// Experimental.
func DeliveryStream_FromDeliveryStreamName(scope constructs.Construct, id *string, deliveryStreamName *string) IDeliveryStream {
	_init_.Initialize()

	var returns IDeliveryStream

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisfirehose-alpha.DeliveryStream",
		"fromDeliveryStreamName",
		[]interface{}{scope, id, deliveryStreamName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DeliveryStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisfirehose-alpha.DeliveryStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func DeliveryStream_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisfirehose-alpha.DeliveryStream",
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
func (d *jsiiProxy_DeliveryStream) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (d *jsiiProxy_DeliveryStream) GeneratePhysicalName() *string {
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
func (d *jsiiProxy_DeliveryStream) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (d *jsiiProxy_DeliveryStream) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the `grantee` identity permissions to perform `actions`.
// Experimental.
func (d *jsiiProxy_DeliveryStream) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grant",
		args,
		&returns,
	)

	return returns
}

// Grant the `grantee` identity permissions to perform `firehose:PutRecord` and `firehose:PutRecordBatch` actions on this delivery stream.
// Experimental.
func (d *jsiiProxy_DeliveryStream) GrantPutRecords(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantPutRecords",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this delivery stream.
// Experimental.
func (d *jsiiProxy_DeliveryStream) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of bytes delivered to Amazon S3 for backup over the specified time period.
//
// By default, this metric will be calculated as an average over a period of 5 minutes.
// Experimental.
func (d *jsiiProxy_DeliveryStream) MetricBackupToS3Bytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricBackupToS3Bytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the age (from getting into Kinesis Data Firehose to now) of the oldest record in Kinesis Data Firehose.
//
// Any record older than this age has been delivered to the Amazon S3 bucket for backup.
//
// By default, this metric will be calculated as an average over a period of 5 minutes.
// Experimental.
func (d *jsiiProxy_DeliveryStream) MetricBackupToS3DataFreshness(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricBackupToS3DataFreshness",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of records delivered to Amazon S3 for backup over the specified time period.
//
// By default, this metric will be calculated as an average over a period of 5 minutes.
// Experimental.
func (d *jsiiProxy_DeliveryStream) MetricBackupToS3Records(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricBackupToS3Records",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of bytes ingested successfully into the delivery stream over the specified time period after throttling.
//
// By default, this metric will be calculated as an average over a period of 5 minutes.
// Experimental.
func (d *jsiiProxy_DeliveryStream) MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricIncomingBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of records ingested successfully into the delivery stream over the specified time period after throttling.
//
// By default, this metric will be calculated as an average over a period of 5 minutes.
// Experimental.
func (d *jsiiProxy_DeliveryStream) MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricIncomingRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DeliveryStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A full specification of a delivery stream that can be used to import it fluently into the CDK application.
// Experimental.
type DeliveryStreamAttributes struct {
	// The ARN of the delivery stream.
	//
	// At least one of deliveryStreamArn and deliveryStreamName must be provided.
	// Experimental.
	DeliveryStreamArn *string `json:"deliveryStreamArn"`
	// The name of the delivery stream.
	//
	// At least one of deliveryStreamName and deliveryStreamArn  must be provided.
	// Experimental.
	DeliveryStreamName *string `json:"deliveryStreamName"`
	// The IAM role associated with this delivery stream.
	//
	// Assumed by Kinesis Data Firehose to read from sources and encrypt data server-side.
	// Experimental.
	Role awsiam.IRole `json:"role"`
}

// Properties for a new delivery stream.
// Experimental.
type DeliveryStreamProps struct {
	// The destinations that this delivery stream will deliver data to.
	//
	// Only a singleton array is supported at this time.
	// Experimental.
	Destinations *[]IDestination `json:"destinations"`
	// A name for the delivery stream.
	// Experimental.
	DeliveryStreamName *string `json:"deliveryStreamName"`
	// Indicates the type of customer master key (CMK) to use for server-side encryption, if any.
	// Experimental.
	Encryption StreamEncryption `json:"encryption"`
	// Customer managed key to server-side encrypt data in the stream.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// The IAM role associated with this delivery stream.
	//
	// Assumed by Kinesis Data Firehose to read from sources and encrypt data server-side.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The Kinesis data stream to use as a source for this delivery stream.
	// Experimental.
	SourceStream awskinesis.IStream `json:"sourceStream"`
}

// Options when binding a destination to a delivery stream.
// Experimental.
type DestinationBindOptions struct {
}

// A Kinesis Data Firehose delivery stream destination configuration.
// Experimental.
type DestinationConfig struct {
	// Any resources that were created by the destination when binding it to the stack that must be deployed before the delivery stream is deployed.
	// Experimental.
	Dependables *[]constructs.IDependable `json:"dependables"`
	// S3 destination configuration properties.
	// Experimental.
	ExtendedS3DestinationConfiguration *awskinesisfirehose.CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty `json:"extendedS3DestinationConfiguration"`
}

// A data processor that Kinesis Data Firehose will call to transform records before delivering data.
// Experimental.
type IDataProcessor interface {
	// Binds this processor to a destination of a delivery stream.
	//
	// Implementers should use this method to grant processor invocation permissions to the provided stream and return the
	// necessary configuration to register as a processor.
	// Experimental.
	Bind(scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig
	// The constructor props of the DataProcessor.
	// Experimental.
	Props() *DataProcessorProps
}

// The jsii proxy for IDataProcessor
type jsiiProxy_IDataProcessor struct {
	_ byte // padding
}

func (i *jsiiProxy_IDataProcessor) Bind(scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig {
	var returns *DataProcessorConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDataProcessor) Props() *DataProcessorProps {
	var returns *DataProcessorProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

// Represents a Kinesis Data Firehose delivery stream.
// Experimental.
type IDeliveryStream interface {
	awsec2.IConnectable
	awsiam.IGrantable
	awscdk.IResource
	// Grant the `grantee` identity permissions to perform `actions`.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the `grantee` identity permissions to perform `firehose:PutRecord` and `firehose:PutRecordBatch` actions on this delivery stream.
	// Experimental.
	GrantPutRecords(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this delivery stream.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of bytes delivered to Amazon S3 for backup over the specified time period.
	//
	// By default, this metric will be calculated as an average over a period of 5 minutes.
	// Experimental.
	MetricBackupToS3Bytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the age (from getting into Kinesis Data Firehose to now) of the oldest record in Kinesis Data Firehose.
	//
	// Any record older than this age has been delivered to the Amazon S3 bucket for backup.
	//
	// By default, this metric will be calculated as an average over a period of 5 minutes.
	// Experimental.
	MetricBackupToS3DataFreshness(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of records delivered to Amazon S3 for backup over the specified time period.
	//
	// By default, this metric will be calculated as an average over a period of 5 minutes.
	// Experimental.
	MetricBackupToS3Records(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of bytes ingested successfully into the delivery stream over the specified time period after throttling.
	//
	// By default, this metric will be calculated as an average over a period of 5 minutes.
	// Experimental.
	MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of records ingested successfully into the delivery stream over the specified time period after throttling.
	//
	// By default, this metric will be calculated as an average over a period of 5 minutes.
	// Experimental.
	MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ARN of the delivery stream.
	// Experimental.
	DeliveryStreamArn() *string
	// The name of the delivery stream.
	// Experimental.
	DeliveryStreamName() *string
}

// The jsii proxy for IDeliveryStream
type jsiiProxy_IDeliveryStream struct {
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IDeliveryStream) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IDeliveryStream) GrantPutRecords(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPutRecords",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) MetricBackupToS3Bytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricBackupToS3Bytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) MetricBackupToS3DataFreshness(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricBackupToS3DataFreshness",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) MetricBackupToS3Records(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricBackupToS3Records",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIncomingBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIncomingRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDeliveryStream) DeliveryStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) DeliveryStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// A Kinesis Data Firehose delivery stream destination.
// Experimental.
type IDestination interface {
	// Binds this destination to the Kinesis Data Firehose delivery stream.
	//
	// Implementers should use this method to bind resources to the stack and initialize values using the provided stream.
	// Experimental.
	Bind(scope constructs.Construct, options *DestinationBindOptions) *DestinationConfig
}

// The jsii proxy for IDestination
type jsiiProxy_IDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_IDestination) Bind(scope constructs.Construct, options *DestinationBindOptions) *DestinationConfig {
	var returns *DestinationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

// Use an AWS Lambda function to transform records.
// Experimental.
type LambdaFunctionProcessor interface {
	IDataProcessor
	Props() *DataProcessorProps
	Bind(_scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig
}

// The jsii proxy struct for LambdaFunctionProcessor
type jsiiProxy_LambdaFunctionProcessor struct {
	jsiiProxy_IDataProcessor
}

func (j *jsiiProxy_LambdaFunctionProcessor) Props() *DataProcessorProps {
	var returns *DataProcessorProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaFunctionProcessor(lambdaFunction awslambda.IFunction, props *DataProcessorProps) LambdaFunctionProcessor {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionProcessor{}

	_jsii_.Create(
		"@aws-cdk/aws-kinesisfirehose-alpha.LambdaFunctionProcessor",
		[]interface{}{lambdaFunction, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaFunctionProcessor_Override(l LambdaFunctionProcessor, lambdaFunction awslambda.IFunction, props *DataProcessorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-kinesisfirehose-alpha.LambdaFunctionProcessor",
		[]interface{}{lambdaFunction, props},
		l,
	)
}

// Binds this processor to a destination of a delivery stream.
//
// Implementers should use this method to grant processor invocation permissions to the provided stream and return the
// necessary configuration to register as a processor.
// Experimental.
func (l *jsiiProxy_LambdaFunctionProcessor) Bind(_scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig {
	var returns *DataProcessorConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}

// Options for server-side encryption of a delivery stream.
// Experimental.
type StreamEncryption string

const (
	StreamEncryption_UNENCRYPTED StreamEncryption = "UNENCRYPTED"
	StreamEncryption_CUSTOMER_MANAGED StreamEncryption = "CUSTOMER_MANAGED"
	StreamEncryption_AWS_OWNED StreamEncryption = "AWS_OWNED"
)

