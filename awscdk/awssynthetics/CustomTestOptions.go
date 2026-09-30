package awssynthetics


// Properties for specifying a test.
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
type CustomTestOptions struct {
	// The code of the canary script.
	Code Code `field:"required" json:"code" yaml:"code"`
	// The handler for the code.
	//
	// Must end with `.handler`.
	Handler *string `field:"required" json:"handler" yaml:"handler"`
}

