package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Scala Streaming Jobs class.
//
// A Streaming job is similar to an ETL job, except that it performs ETL on data streams
// using the Apache Spark Structured Streaming framework.
// These jobs will default to use Python 3.9.
//
// Similar to ETL jobs, streaming job supports Scala and Python languages. Similar to ETL,
// it supports G1 and G2 worker type and 2.0, 3.0 and 4.0 version. We’ll default to G2 worker
// and 4.0 version for streaming jobs which developers can override.
// We will enable —enable-metrics, —enable-spark-ui, —enable-continuous-cloudwatch-log.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var code code
//   var connection connection
//   var logGroup logGroup
//   var role role
//   var securityConfiguration securityConfiguration
//
//   scalaSparkStreamingJob := glue_alpha.NewScalaSparkStreamingJob(this, jsii.String("MyScalaSparkStreamingJob"), &ScalaSparkStreamingJobProps{
//   	ClassName: jsii.String("className"),
//   	Role: role,
//   	Script: code,
//
//   	// the properties below are optional
//   	Connections: []iConnection{
//   		connection,
//   	},
//   	ContinuousLogging: &ContinuousLoggingProps{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		ConversionPattern: jsii.String("conversionPattern"),
//   		LogGroup: logGroup,
//   		LogStreamPrefix: jsii.String("logStreamPrefix"),
//   		Quiet: jsii.Boolean(false),
//   	},
//   	DefaultArguments: map[string]*string{
//   		"defaultArgumentsKey": jsii.String("defaultArguments"),
//   	},
//   	Description: jsii.String("description"),
//   	EnableProfilingMetrics: jsii.Boolean(false),
//   	ExtraFiles: []*code{
//   		code,
//   	},
//   	ExtraJars: []*code{
//   		code,
//   	},
//   	ExtraJarsFirst: jsii.Boolean(false),
//   	GlueVersion: glue_alpha.GlueVersion_V0_9,
//   	JobName: jsii.String("jobName"),
//   	JobRunQueuingEnabled: jsii.Boolean(false),
//   	MaxConcurrentRuns: jsii.Number(123),
//   	MaxRetries: jsii.Number(123),
//   	NumberOfWorkers: jsii.Number(123),
//   	SecurityConfiguration: securityConfiguration,
//   	SparkUI: &SparkUIProps{
//   		Bucket: bucket,
//   		Prefix: jsii.String("prefix"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	WorkerType: glue_alpha.WorkerType_STANDARD,
//   })
//
// Experimental.
type ScalaSparkStreamingJob interface {
	SparkJob
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
	// The principal to grant permissions to.
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
	//   cross-environment scenarios.
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
	// Returns the job arn.
	// Experimental.
	BuildJobArn(scope constructs.Construct, jobName *string) *string
	// Check no usage of reserved arguments.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	CheckNoReservedArgs(defaultArguments *map[string]*string) *map[string]*string
	// Experimental.
	CodeS3ObjectUrl(code Code) *string
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
	// Experimental.
	NonExecutableCommonArguments(props *SparkJobProps) *map[string]*string
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
	// Setup Continuous Logging Properties.
	//
	// Returns: String containing the args for the continuous logging command.
	// Experimental.
	SetupContinuousLogging(role awsiam.IRole, props *ContinuousLoggingProps) interface{}
	// Set the arguments for extra {@link Code}-related properties.
	// Experimental.
	SetupExtraCodeArguments(args *map[string]*string, props *SparkExtraCodeProps)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ScalaSparkStreamingJob
type jsiiProxy_ScalaSparkStreamingJob struct {
	jsiiProxy_SparkJob
}

func (j *jsiiProxy_ScalaSparkStreamingJob) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalaSparkStreamingJob) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalaSparkStreamingJob) JobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalaSparkStreamingJob) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalaSparkStreamingJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalaSparkStreamingJob) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalaSparkStreamingJob) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalaSparkStreamingJob) SparkUILoggingLocation() *SparkUILoggingLocation {
	var returns *SparkUILoggingLocation
	_jsii_.Get(
		j,
		"sparkUILoggingLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalaSparkStreamingJob) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// ScalaSparkStreamingJob constructor.
// Experimental.
func NewScalaSparkStreamingJob(scope constructs.Construct, id *string, props *ScalaSparkStreamingJobProps) ScalaSparkStreamingJob {
	_init_.Initialize()

	if err := validateNewScalaSparkStreamingJobParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ScalaSparkStreamingJob{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.ScalaSparkStreamingJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// ScalaSparkStreamingJob constructor.
// Experimental.
func NewScalaSparkStreamingJob_Override(s ScalaSparkStreamingJob, scope constructs.Construct, id *string, props *ScalaSparkStreamingJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.ScalaSparkStreamingJob",
		[]interface{}{scope, id, props},
		s,
	)
}

// Identifies an existing Glue Job from a subset of attributes that can be referenced from within another Stack or Construct.
// Experimental.
func ScalaSparkStreamingJob_FromJobAttributes(scope constructs.Construct, id *string, attrs *JobAttributes) IJob {
	_init_.Initialize()

	if err := validateScalaSparkStreamingJob_FromJobAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IJob

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.ScalaSparkStreamingJob",
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
func ScalaSparkStreamingJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateScalaSparkStreamingJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.ScalaSparkStreamingJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func ScalaSparkStreamingJob_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateScalaSparkStreamingJob_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.ScalaSparkStreamingJob",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ScalaSparkStreamingJob_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateScalaSparkStreamingJob_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.ScalaSparkStreamingJob",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func ScalaSparkStreamingJob_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ScalaSparkStreamingJob",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_ScalaSparkStreamingJob) BuildJobArn(scope constructs.Construct, jobName *string) *string {
	if err := s.validateBuildJobArnParameters(scope, jobName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"buildJobArn",
		[]interface{}{scope, jobName},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) CheckNoReservedArgs(defaultArguments *map[string]*string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"checkNoReservedArgs",
		[]interface{}{defaultArguments},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) CodeS3ObjectUrl(code Code) *string {
	if err := s.validateCodeS3ObjectUrlParameters(code); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"codeS3ObjectUrl",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (s *jsiiProxy_ScalaSparkStreamingJob) GetResourceNameAttribute(nameAttr *string) *string {
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

func (s *jsiiProxy_ScalaSparkStreamingJob) Metric(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricParameters(metricName, type_, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, type_, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) MetricFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricFailureParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricFailure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) MetricSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricSuccessParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricSuccess",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) MetricTimeout(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricTimeoutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTimeout",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) NonExecutableCommonArguments(props *SparkJobProps) *map[string]*string {
	if err := s.validateNonExecutableCommonArgumentsParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"nonExecutableCommonArguments",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := s.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) OnFailure(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := s.validateOnFailureParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onFailure",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) OnStateChange(id *string, jobState JobState, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := s.validateOnStateChangeParameters(id, jobState, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onStateChange",
		[]interface{}{id, jobState, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) OnSuccess(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := s.validateOnSuccessParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onSuccess",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) OnTimeout(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := s.validateOnTimeoutParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onTimeout",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) SetupContinuousLogging(role awsiam.IRole, props *ContinuousLoggingProps) interface{} {
	if err := s.validateSetupContinuousLoggingParameters(role, props); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"setupContinuousLogging",
		[]interface{}{role, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalaSparkStreamingJob) SetupExtraCodeArguments(args *map[string]*string, props *SparkExtraCodeProps) {
	if err := s.validateSetupExtraCodeArgumentsParameters(args, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"setupExtraCodeArguments",
		[]interface{}{args, props},
	)
}

func (s *jsiiProxy_ScalaSparkStreamingJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

