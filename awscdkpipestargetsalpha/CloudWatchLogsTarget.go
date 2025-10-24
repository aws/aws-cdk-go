package awscdkpipestargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/internal"
)

// An EventBridge Pipes target that sends messages to a CloudWatch Logs log group.
//
// Example:
//   var sourceQueue Queue
//   var targetLogGroup LogGroup
//
//
//   logGroupTarget := targets.NewCloudWatchLogsTarget(targetLogGroup)
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: logGroupTarget,
//   })
//
// Experimental.
type CloudWatchLogsTarget interface {
	awscdkpipesalpha.ITarget
	// The ARN of the target resource.
	// Experimental.
	TargetArn() *string
	// Bind this target to a pipe.
	// Experimental.
	Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.TargetConfig
	// Grant the pipe role to push to the target.
	// Experimental.
	GrantPush(grantee awsiam.IRole)
}

// The jsii proxy struct for CloudWatchLogsTarget
type jsiiProxy_CloudWatchLogsTarget struct {
	internal.Type__awscdkpipesalphaITarget
}

func (j *jsiiProxy_CloudWatchLogsTarget) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewCloudWatchLogsTarget(logGroup awslogs.ILogGroup, parameters *CloudWatchLogsTargetParameters) CloudWatchLogsTarget {
	_init_.Initialize()

	if err := validateNewCloudWatchLogsTargetParameters(logGroup, parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudWatchLogsTarget{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.CloudWatchLogsTarget",
		[]interface{}{logGroup, parameters},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudWatchLogsTarget_Override(c CloudWatchLogsTarget, logGroup awslogs.ILogGroup, parameters *CloudWatchLogsTargetParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.CloudWatchLogsTarget",
		[]interface{}{logGroup, parameters},
		c,
	)
}

func (c *jsiiProxy_CloudWatchLogsTarget) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.TargetConfig {
	if err := c.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *awscdkpipesalpha.TargetConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudWatchLogsTarget) GrantPush(grantee awsiam.IRole) {
	if err := c.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"grantPush",
		[]interface{}{grantee},
	)
}

