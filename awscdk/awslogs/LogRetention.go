package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
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
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   logRetention := awscdk.Aws_logs.NewLogRetention(this, jsii.String("MyLogRetention"), &LogRetentionProps{
//   	LogGroupName: jsii.String("logGroupName"),
//   	Retention: awscdk.*Aws_logs.RetentionDays_ONE_DAY,
//
//   	// the properties below are optional
//   	LogGroupRegion: jsii.String("logGroupRegion"),
//   	LogRetentionRetryOptions: &LogRetentionRetryOptions{
//   		Base: cdk.Duration_Minutes(jsii.Number(30)),
//   		MaxRetries: jsii.Number(123),
//   	},
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	Role: role,
//   })
//
type LogRetention interface {
	constructs.Construct
	// The ARN of the LogGroup.
	LogGroupArn() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for LogRetention
type jsiiProxy_LogRetention struct {
	internal.Type__constructsConstruct
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

func (j *jsiiProxy_LogRetention) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewLogRetention(scope constructs.Construct, id *string, props *LogRetentionProps) LogRetention {
	_init_.Initialize()

	if err := validateNewLogRetentionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogRetention{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.LogRetention",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLogRetention_Override(l LogRetention, scope constructs.Construct, id *string, props *LogRetentionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.LogRetention",
		[]interface{}{scope, id, props},
		l,
	)
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
func LogRetention_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogRetention_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.LogRetention",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
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

