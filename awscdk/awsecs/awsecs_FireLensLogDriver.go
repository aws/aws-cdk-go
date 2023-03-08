package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// FireLens enables you to use task definition parameters to route logs to an AWS service   or AWS Partner Network (APN) destination for log storage and analytics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret secret
//
//   fireLensLogDriver := awscdk.Aws_ecs.NewFireLensLogDriver(&fireLensLogDriverProps{
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	options: map[string]*string{
//   		"optionsKey": jsii.String("options"),
//   	},
//   	secretOptions: map[string]*secret{
//   		"secretOptionsKey": secret,
//   	},
//   	tag: jsii.String("tag"),
//   })
//
// Experimental.
type FireLensLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
	// Experimental.
	Bind(_scope awscdk.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for FireLensLogDriver
type jsiiProxy_FireLensLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the FireLensLogDriver class.
// Experimental.
func NewFireLensLogDriver(props *FireLensLogDriverProps) FireLensLogDriver {
	_init_.Initialize()

	if err := validateNewFireLensLogDriverParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FireLensLogDriver{}

	_jsii_.Create(
		"monocdk.aws_ecs.FireLensLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the FireLensLogDriver class.
// Experimental.
func NewFireLensLogDriver_Override(f FireLensLogDriver, props *FireLensLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.FireLensLogDriver",
		[]interface{}{props},
		f,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
// Experimental.
func FireLensLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateFireLensLogDriver_AwsLogsParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.FireLensLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FireLensLogDriver) Bind(_scope awscdk.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	if err := f.validateBindParameters(_scope, _containerDefinition); err != nil {
		panic(err)
	}
	var returns *LogDriverConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

