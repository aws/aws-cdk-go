package awscloudwatch


// Specify how missing data points are treated during alarm evaluation.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	Timeout: awscdk.Duration_Minutes(jsii.Number(5)),
//   })
//
//   if fn.Timeout {
//   	cloudwatch.NewAlarm(this, jsii.String("MyAlarm"), &AlarmProps{
//   		Metric: fn.metricDuration().With(&MetricOptions{
//   			Statistic: jsii.String("Maximum"),
//   		}),
//   		EvaluationPeriods: jsii.Number(1),
//   		DatapointsToAlarm: jsii.Number(1),
//   		Threshold: fn.*Timeout.ToMilliseconds(),
//   		TreatMissingData: cloudwatch.TreatMissingData_IGNORE,
//   		AlarmName: jsii.String("My Lambda Timeout"),
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

