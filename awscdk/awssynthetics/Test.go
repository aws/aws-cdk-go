package awssynthetics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specify a test that the canary should run.
//
// Example:
//   canary := synthetics.NewCanary(this, jsii.String("MyCanary"), &CanaryProps{
//   	Schedule: synthetics.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(5))),
//   	Test: synthetics.Test_Custom(&CustomTestOptions{
//   		Code: synthetics.Code_FromAsset(path.join(__dirname, jsii.String("canary"))),
//   		Handler: jsii.String("index.handler"),
//   	}),
//   	Runtime: synthetics.Runtime_SYNTHETICS_NODEJS_PUPPETEER_7_0(),
//   	ResourcesToReplicateTags: []lAMBDA_FUNCTION{
//   		synthetics.ResourceToReplicateTags_*lAMBDA_FUNCTION,
//   	},
//   })
//
type Test interface {
	// The code that the canary should run.
	Code() Code
	// The handler of the canary.
	Handler() *string
}

// The jsii proxy struct for Test
type jsiiProxy_Test struct {
	_ byte // padding
}

func (j *jsiiProxy_Test) Code() Code {
	var returns Code
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Test) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}


// Specify a custom test with your own code.
//
// Returns: `Test` associated with the specified Code object.
func Test_Custom(options *CustomTestOptions) Test {
	_init_.Initialize()

	if err := validateTest_CustomParameters(options); err != nil {
		panic(err)
	}
	var returns Test

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.Test",
		"custom",
		[]interface{}{options},
		&returns,
	)

	return returns
}

