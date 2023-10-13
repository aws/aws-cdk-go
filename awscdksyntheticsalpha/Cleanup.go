package awscdksyntheticsalpha


// Different ways to clean up underlying Canary resources when the Canary is deleted.
//
// Example:
//   canary := synthetics.NewCanary(this, jsii.String("Canary"), &CanaryProps{
//   	Test: synthetics.Test_Custom(&CustomTestOptions{
//   		Handler: jsii.String("index.handler"),
//   		Code: synthetics.Code_FromInline(jsii.String("/* Synthetics handler code")),
//   	}),
//   	Cleanup: synthetics.Cleanup_LAMBDA,
//   	Runtime: synthetics.Runtime_SYNTHETICS_NODEJS_PUPPETEER_4_0(),
//   })
//
// Deprecated.
type Cleanup string

const (
	// Clean up nothing.
	//
	// The user is responsible for cleaning up
	// all resources left behind by the Canary.
	// Deprecated.
	Cleanup_NOTHING Cleanup = "NOTHING"
	// Clean up the underlying Lambda function only.
	//
	// The user is
	// responsible for cleaning up all other resources left behind
	// by the Canary.
	// Deprecated.
	Cleanup_LAMBDA Cleanup = "LAMBDA"
)

