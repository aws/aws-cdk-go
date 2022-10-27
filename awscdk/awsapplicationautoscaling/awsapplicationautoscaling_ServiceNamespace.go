package awsapplicationautoscaling


// The service that supports Application AutoScaling.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var code code
//
//
//   handler := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_PYTHON_3_7(),
//   	handler: jsii.String("index.handler"),
//   	code: code,
//
//   	reservedConcurrentExecutions: jsii.Number(2),
//   })
//
//   fnVer := handler.currentVersion
//
//   target := appscaling.NewScalableTarget(this, jsii.String("ScalableTarget"), &scalableTargetProps{
//   	serviceNamespace: appscaling.serviceNamespace_LAMBDA,
//   	maxCapacity: jsii.Number(100),
//   	minCapacity: jsii.Number(10),
//   	resourceId: fmt.Sprintf("function:%v:%v", handler.functionName, fnVer.version),
//   	scalableDimension: jsii.String("lambda:function:ProvisionedConcurrency"),
//   })
//
//   target.scaleToTrackMetric(jsii.String("PceTracking"), &basicTargetTrackingScalingPolicyProps{
//   	targetValue: jsii.Number(0.9),
//   	predefinedMetric: appscaling.predefinedMetric_LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION,
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

