package awssynthetics


// Resources that tags applied to a canary should be replicated to.
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
type ResourceToReplicateTags string

const (
	// Replicate canary tags to the Lambda function.
	//
	// When specified, CloudWatch Synthetics will keep the tags of the canary
	// and the Lambda function synchronized. Any future changes made to the
	// canary's tags will also be applied to the function.
	ResourceToReplicateTags_LAMBDA_FUNCTION ResourceToReplicateTags = "LAMBDA_FUNCTION"
)

