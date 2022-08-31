package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// A log driver that sends log information to splunk Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret secret
//   var secretValue secretValue
//
//   splunkLogDriver := awscdk.Aws_ecs.NewSplunkLogDriver(&splunkLogDriverProps{
//   	url: jsii.String("url"),
//
//   	// the properties below are optional
//   	caName: jsii.String("caName"),
//   	caPath: jsii.String("caPath"),
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	format: awscdk.*Aws_ecs.splunkLogFormat_INLINE,
//   	gzip: jsii.Boolean(false),
//   	gzipLevel: jsii.Number(123),
//   	index: jsii.String("index"),
//   	insecureSkipVerify: jsii.String("insecureSkipVerify"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	secretToken: secret,
//   	source: jsii.String("source"),
//   	sourceType: jsii.String("sourceType"),
//   	tag: jsii.String("tag"),
//   	token: secretValue,
//   	verifyConnection: jsii.Boolean(false),
//   })
//
// Experimental.
type SplunkLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
	// Experimental.
	Bind(_scope awscdk.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for SplunkLogDriver
type jsiiProxy_SplunkLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the SplunkLogDriver class.
// Experimental.
func NewSplunkLogDriver(props *SplunkLogDriverProps) SplunkLogDriver {
	_init_.Initialize()

	j := jsiiProxy_SplunkLogDriver{}

	_jsii_.Create(
		"monocdk.aws_ecs.SplunkLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the SplunkLogDriver class.
// Experimental.
func NewSplunkLogDriver_Override(s SplunkLogDriver, props *SplunkLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.SplunkLogDriver",
		[]interface{}{props},
		s,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
// Experimental.
func SplunkLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.SplunkLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SplunkLogDriver) Bind(_scope awscdk.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	var returns *LogDriverConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

