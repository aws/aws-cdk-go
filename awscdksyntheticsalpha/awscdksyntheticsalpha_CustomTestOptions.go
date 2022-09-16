// The CDK Construct Library for AWS::Synthetics
package awscdksyntheticsalpha


// Properties for specifying a test.
//
// Example:
//   canary := synthetics.NewCanary(this, jsii.String("MyCanary"), &canaryProps{
//   	schedule: synthetics.schedule.rate(awscdk.Duration.minutes(jsii.Number(5))),
//   	test: synthetics.test.custom(&customTestOptions{
//   		code: synthetics.code.fromAsset(path.join(__dirname, jsii.String("canary"))),
//   		handler: jsii.String("index.handler"),
//   	}),
//   	runtime: synthetics.runtime_SYNTHETICS_NODEJS_PUPPETEER_3_1(),
//   	environmentVariables: map[string]*string{
//   		"stage": jsii.String("prod"),
//   	},
//   })
//
// Experimental.
type CustomTestOptions struct {
	// The code of the canary script.
	// Experimental.
	Code Code `field:"required" json:"code" yaml:"code"`
	// The handler for the code.
	//
	// Must end with `.handler`.
	// Experimental.
	Handler *string `field:"required" json:"handler" yaml:"handler"`
}

