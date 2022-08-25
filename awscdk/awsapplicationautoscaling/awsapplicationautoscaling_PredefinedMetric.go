package awsapplicationautoscaling


// One of the predefined autoscaling metrics.
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
type PredefinedMetric string

const (
	// Average percentage of instances in an AppStream fleet that are being used.
	PredefinedMetric_APPSTREAM_AVERAGE_CAPACITY_UTILIZATION PredefinedMetric = "APPSTREAM_AVERAGE_CAPACITY_UTILIZATION"
	// Percentage of provisioned read capacity units utilized by a Keyspaces table.
	PredefinedMetric_CASSANDRA_READ_CAPACITY_UTILIZATION PredefinedMetric = "CASSANDRA_READ_CAPACITY_UTILIZATION"
	// Percentage of provisioned write capacity units utilized by a Keyspaces table.
	PredefinedMetric_CASSANDRA_WRITE_CAPACITY_UTILIZATION PredefinedMetric = "CASSANDRA_WRITE_CAPACITY_UTILIZATION"
	// Percentage of provisioned inference units utilized by a Comprehend endpoint.
	PredefinedMetric_COMPREHEND_INFERENCE_UTILIZATION PredefinedMetric = "COMPREHEND_INFERENCE_UTILIZATION"
	// Average CPU Utilization of read replica instances in a Neptune DB cluster.
	PredefinedMetric_NEPTURE_READER_AVERAGE_CPU_UTILIZATION PredefinedMetric = "NEPTURE_READER_AVERAGE_CPU_UTILIZATION"
	// Percentage of provisioned read capacity units consumed by a DynamoDB table.
	PredefinedMetric_DYNAMODB_READ_CAPACITY_UTILIZATION PredefinedMetric = "DYNAMODB_READ_CAPACITY_UTILIZATION"
	// Percentage of provisioned write capacity units consumed by a DynamoDB table.
	//
	// Suffix `dummy` is necessary due to jsii bug (https://github.com/aws/jsii/issues/2782).
	// Duplicate values will be dropped, so this suffix is added as a workaround.
	// The value will be replaced when this enum is used.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	PredefinedMetric_DYNAMODB_WRITE_CAPACITY_UTILIZATION PredefinedMetric = "DYNAMODB_WRITE_CAPACITY_UTILIZATION"
	// DYANMODB_WRITE_CAPACITY_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Deprecated: use `PredefinedMetric.DYNAMODB_WRITE_CAPACITY_UTILIZATION`
	PredefinedMetric_DYANMODB_WRITE_CAPACITY_UTILIZATION PredefinedMetric = "DYANMODB_WRITE_CAPACITY_UTILIZATION"
	// ALB_REQUEST_COUNT_PER_TARGET.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	PredefinedMetric_ALB_REQUEST_COUNT_PER_TARGET PredefinedMetric = "ALB_REQUEST_COUNT_PER_TARGET"
	// RDS_READER_AVERAGE_CPU_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	PredefinedMetric_RDS_READER_AVERAGE_CPU_UTILIZATION PredefinedMetric = "RDS_READER_AVERAGE_CPU_UTILIZATION"
	// RDS_READER_AVERAGE_DATABASE_CONNECTIONS.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	PredefinedMetric_RDS_READER_AVERAGE_DATABASE_CONNECTIONS PredefinedMetric = "RDS_READER_AVERAGE_DATABASE_CONNECTIONS"
	// EC2_SPOT_FLEET_REQUEST_AVERAGE_CPU_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	PredefinedMetric_EC2_SPOT_FLEET_REQUEST_AVERAGE_CPU_UTILIZATION PredefinedMetric = "EC2_SPOT_FLEET_REQUEST_AVERAGE_CPU_UTILIZATION"
	// EC2_SPOT_FLEET_REQUEST_AVERAGE_NETWORK_IN.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	PredefinedMetric_EC2_SPOT_FLEET_REQUEST_AVERAGE_NETWORK_IN PredefinedMetric = "EC2_SPOT_FLEET_REQUEST_AVERAGE_NETWORK_IN"
	// EC2_SPOT_FLEET_REQUEST_AVERAGE_NETWORK_OUT.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	PredefinedMetric_EC2_SPOT_FLEET_REQUEST_AVERAGE_NETWORK_OUT PredefinedMetric = "EC2_SPOT_FLEET_REQUEST_AVERAGE_NETWORK_OUT"
	// SAGEMAKER_VARIANT_INVOCATIONS_PER_INSTANCE.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	PredefinedMetric_SAGEMAKER_VARIANT_INVOCATIONS_PER_INSTANCE PredefinedMetric = "SAGEMAKER_VARIANT_INVOCATIONS_PER_INSTANCE"
	// ECS_SERVICE_AVERAGE_CPU_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	PredefinedMetric_ECS_SERVICE_AVERAGE_CPU_UTILIZATION PredefinedMetric = "ECS_SERVICE_AVERAGE_CPU_UTILIZATION"
	// ECS_SERVICE_AVERAGE_MEMORY_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	PredefinedMetric_ECS_SERVICE_AVERAGE_MEMORY_UTILIZATION PredefinedMetric = "ECS_SERVICE_AVERAGE_MEMORY_UTILIZATION"
	// LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/monitoring-metrics.html#monitoring-metrics-concurrency
	//
	PredefinedMetric_LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION PredefinedMetric = "LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION"
	// KAFKA_BROKER_STORAGE_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	PredefinedMetric_KAFKA_BROKER_STORAGE_UTILIZATION PredefinedMetric = "KAFKA_BROKER_STORAGE_UTILIZATION"
	// ELASTIC_CACHE_PRIMARY_ENGINE_CPU_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	PredefinedMetric_ELASTICACHE_PRIMARY_ENGINE_CPU_UTILIZATION PredefinedMetric = "ELASTICACHE_PRIMARY_ENGINE_CPU_UTILIZATION"
	// ELASTIC_CACHE_REPLICA_ENGINE_CPU_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	PredefinedMetric_ELASTICACHE_REPLICA_ENGINE_CPU_UTILIZATION PredefinedMetric = "ELASTICACHE_REPLICA_ENGINE_CPU_UTILIZATION"
	// ELASTIC_CACHE_REPLICA_ENGINE_CPU_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	PredefinedMetric_ELASTICACHE_DATABASE_MEMORY_USAGE_COUNTED_FOR_EVICT_PERCENTAGE PredefinedMetric = "ELASTICACHE_DATABASE_MEMORY_USAGE_COUNTED_FOR_EVICT_PERCENTAGE"
)

