package awssynthetics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Runtime options for a canary.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//
//   canary := synthetics.NewCanary(this, jsii.String("MyCanary"), &CanaryProps{
//   	Schedule: synthetics.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(5))),
//   	Test: synthetics.Test_Custom(&CustomTestOptions{
//   		Code: synthetics.Code_FromAsset(path.join(__dirname, jsii.String("canary"))),
//   		Handler: jsii.String("index.handler"),
//   	}),
//   	Runtime: synthetics.Runtime_SYNTHETICS_NODEJS_PUPPETEER_6_2(),
//   	Memory: cdk.Size_Mebibytes(jsii.Number(1024)),
//   })
//
type Runtime interface {
	// The Lambda runtime family.
	Family() RuntimeFamily
	// The name of the runtime version.
	Name() *string
}

// The jsii proxy struct for Runtime
type jsiiProxy_Runtime struct {
	_ byte // padding
}

func (j *jsiiProxy_Runtime) Family() RuntimeFamily {
	var returns RuntimeFamily
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runtime) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func NewRuntime(name *string, family RuntimeFamily) Runtime {
	_init_.Initialize()

	if err := validateNewRuntimeParameters(name, family); err != nil {
		panic(err)
	}
	j := jsiiProxy_Runtime{}

	_jsii_.Create(
		"aws-cdk-lib.aws_synthetics.Runtime",
		[]interface{}{name, family},
		&j,
	)

	return &j
}

func NewRuntime_Override(r Runtime, name *string, family RuntimeFamily) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_synthetics.Runtime",
		[]interface{}{name, family},
		r,
	)
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_3_5() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_3_5",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_3_6() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_3_6",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_3_7() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_3_7",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_3_8() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_3_8",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_3_9() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_3_9",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_4_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_4_0",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_5_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_5_0",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_5_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_5_1",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_5_2() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_5_2",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_6_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_6_0",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_6_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_6_1",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_6_2() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_6_2",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_7_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_7_0",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_8_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_8_0",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_PYTHON_SELENIUM_1_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_PYTHON_SELENIUM_1_0",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_PYTHON_SELENIUM_1_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_PYTHON_SELENIUM_1_1",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_PYTHON_SELENIUM_1_2() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_PYTHON_SELENIUM_1_2",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_PYTHON_SELENIUM_1_3() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_PYTHON_SELENIUM_1_3",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_PYTHON_SELENIUM_2_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_PYTHON_SELENIUM_2_0",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_PYTHON_SELENIUM_2_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_PYTHON_SELENIUM_2_1",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_PYTHON_SELENIUM_3_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_PYTHON_SELENIUM_3_0",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_PYTHON_SELENIUM_4_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.Runtime",
		"SYNTHETICS_PYTHON_SELENIUM_4_0",
		&returns,
	)
	return returns
}

