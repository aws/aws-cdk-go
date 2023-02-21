package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Creates a custom resource to control the retention policy of a CloudWatch Logs log group.
//
// The log group is created if it doesn't already exist. The policy
// is removed when `retentionDays` is `undefined` or equal to `Infinity`.
// Log group can be created in the region that is different from stack region by
// specifying `logGroupRegion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var role role
//
//   logRetention := awscdk.Aws_logs.NewLogRetention(this, jsii.String("MyLogRetention"), &logRetentionProps{
//   	logGroupName: jsii.String("logGroupName"),
//   	retention: awscdk.*Aws_logs.retentionDays_ONE_DAY,
//
//   	// the properties below are optional
//   	logGroupRegion: jsii.String("logGroupRegion"),
//   	logRetentionRetryOptions: &logRetentionRetryOptions{
//   		base: duration,
//   		maxRetries: jsii.Number(123),
//   	},
//   	role: role,
//   })
//
// Experimental.
type LogRetention interface {
	awscdk.Construct
	// The ARN of the LogGroup.
	// Experimental.
	LogGroupArn() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
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

// The jsii proxy struct for LogRetention
type jsiiProxy_LogRetention struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_LogRetention) LogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogRetention) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewLogRetention(scope constructs.Construct, id *string, props *LogRetentionProps) LogRetention {
	_init_.Initialize()

	if err := validateNewLogRetentionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogRetention{}

	_jsii_.Create(
		"monocdk.aws_logs.LogRetention",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLogRetention_Override(l LogRetention, scope constructs.Construct, id *string, props *LogRetentionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.LogRetention",
		[]interface{}{scope, id, props},
		l,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func LogRetention_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogRetention_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.LogRetention",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogRetention) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogRetention) OnSynthesize(session constructs.ISynthesisSession) {
	if err := l.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (l *jsiiProxy_LogRetention) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogRetention) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogRetention) Synthesize(session awscdk.ISynthesisSession) {
	if err := l.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

func (l *jsiiProxy_LogRetention) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogRetention) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

