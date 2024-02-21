package awsapplicationautoscaling


// The service that supports Application AutoScaling.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var code code
//
//
//   handler := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_PYTHON_3_12(),
//   	Handler: jsii.String("index.handler"),
//   	Code: Code,
//
//   	ReservedConcurrentExecutions: jsii.Number(2),
//   })
//
//   fnVer := handler.currentVersion
//
//   target := appscaling.NewScalableTarget(this, jsii.String("ScalableTarget"), &ScalableTargetProps{
//   	ServiceNamespace: appscaling.ServiceNamespace_LAMBDA,
//   	MaxCapacity: jsii.Number(100),
//   	MinCapacity: jsii.Number(10),
//   	ResourceId: fmt.Sprintf("function:%v:%v", handler.FunctionName, fnVer.Version),
//   	ScalableDimension: jsii.String("lambda:function:ProvisionedConcurrency"),
//   })
//
//   target.ScaleToTrackMetric(jsii.String("PceTracking"), &BasicTargetTrackingScalingPolicyProps{
//   	TargetValue: jsii.Number(0.9),
//   	PredefinedMetric: appscaling.PredefinedMetric_LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION,
//   })
//
type ServiceNamespace string

const (
	// Elastic Container Service.
	ServiceNamespace_ECS ServiceNamespace = "ECS"
	// Elastic Map Reduce.
	ServiceNamespace_ELASTIC_MAP_REDUCE ServiceNamespace = "ELASTIC_MAP_REDUCE"
	// Elastic Compute Cloud.
	ServiceNamespace_EC2 ServiceNamespace = "EC2"
	// App Stream.
	ServiceNamespace_APPSTREAM ServiceNamespace = "APPSTREAM"
	// Dynamo DB.
	ServiceNamespace_DYNAMODB ServiceNamespace = "DYNAMODB"
	// Relational Database Service.
	ServiceNamespace_RDS ServiceNamespace = "RDS"
	// SageMaker.
	ServiceNamespace_SAGEMAKER ServiceNamespace = "SAGEMAKER"
	// Custom Resource.
	ServiceNamespace_CUSTOM_RESOURCE ServiceNamespace = "CUSTOM_RESOURCE"
	// Lambda.
	ServiceNamespace_LAMBDA ServiceNamespace = "LAMBDA"
	// Comprehend.
	ServiceNamespace_COMPREHEND ServiceNamespace = "COMPREHEND"
	// Kafka.
	ServiceNamespace_KAFKA ServiceNamespace = "KAFKA"
	// ElastiCache.
	ServiceNamespace_ELASTICACHE ServiceNamespace = "ELASTICACHE"
	// Neptune.
	ServiceNamespace_NEPTUNE ServiceNamespace = "NEPTUNE"
)

