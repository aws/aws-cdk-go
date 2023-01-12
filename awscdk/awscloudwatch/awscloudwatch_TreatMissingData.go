package awscloudwatch


// Specify how missing data points are treated during alarm evaluation.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_18_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	timeout: cdk.duration.minutes(jsii.Number(5)),
//   })
//
//   if fn.timeout {
//   	cloudwatch.NewAlarm(this, jsii.String("MyAlarm"), &alarmProps{
//   		metric: fn.metricDuration().with(&metricOptions{
//   			statistic: jsii.String("Maximum"),
//   		}),
//   		evaluationPeriods: jsii.Number(1),
//   		datapointsToAlarm: jsii.Number(1),
//   		threshold: fn.timeout.toMilliseconds(),
//   		treatMissingData: cloudwatch.treatMissingData_IGNORE,
//   		alarmName: jsii.String("My Lambda Timeout"),
//   	})
//   }
//
type TreatMissingData string

const (
	// Missing data points are treated as breaching the threshold.
	TreatMissingData_BREACHING TreatMissingData = "BREACHING"
	// Missing data points are treated as being within the threshold.
	TreatMissingData_NOT_BREACHING TreatMissingData = "NOT_BREACHING"
	// The current alarm state is maintained.
	TreatMissingData_IGNORE TreatMissingData = "IGNORE"
	// The alarm does not consider missing data points when evaluating whether to change state.
	TreatMissingData_MISSING TreatMissingData = "MISSING"
)

