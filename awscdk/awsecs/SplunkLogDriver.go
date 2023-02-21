package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A log driver that sends log information to splunk Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret secret
//
//   splunkLogDriver := awscdk.Aws_ecs.NewSplunkLogDriver(&SplunkLogDriverProps{
//   	SecretToken: secret,
//   	Url: jsii.String("url"),
//
//   	// the properties below are optional
//   	CaName: jsii.String("caName"),
//   	CaPath: jsii.String("caPath"),
//   	Env: []*string{
//   		jsii.String("env"),
//   	},
//   	EnvRegex: jsii.String("envRegex"),
//   	Format: awscdk.*Aws_ecs.SplunkLogFormat_INLINE,
//   	Gzip: jsii.Boolean(false),
//   	GzipLevel: jsii.Number(123),
//   	Index: jsii.String("index"),
//   	InsecureSkipVerify: jsii.String("insecureSkipVerify"),
//   	Labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	Source: jsii.String("source"),
//   	SourceType: jsii.String("sourceType"),
//   	Tag: jsii.String("tag"),
//   	VerifyConnection: jsii.Boolean(false),
//   })
//
type SplunkLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
	Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for SplunkLogDriver
type jsiiProxy_SplunkLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the SplunkLogDriver class.
func NewSplunkLogDriver(props *SplunkLogDriverProps) SplunkLogDriver {
	_init_.Initialize()

	if err := validateNewSplunkLogDriverParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SplunkLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.SplunkLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the SplunkLogDriver class.
func NewSplunkLogDriver_Override(s SplunkLogDriver, props *SplunkLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.SplunkLogDriver",
		[]interface{}{props},
		s,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func SplunkLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateSplunkLogDriver_AwsLogsParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.SplunkLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SplunkLogDriver) Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	if err := s.validateBindParameters(_scope, _containerDefinition); err != nil {
		panic(err)
	}
	var returns *LogDriverConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

