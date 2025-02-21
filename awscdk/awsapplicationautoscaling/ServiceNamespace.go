package awsapplicationautoscaling


// The service that supports Application AutoScaling.
//
// Example:
//   shardsScalableTarget := appscaling.NewScalableTarget(this, jsii.String("ElastiCacheRedisShardsScalableTarget"), &ScalableTargetProps{
//   	ServiceNamespace: appscaling.ServiceNamespace_ELASTICACHE,
//   	ScalableDimension: jsii.String("elasticache:replication-group:NodeGroups"),
//   	MinCapacity: jsii.Number(2),
//   	MaxCapacity: jsii.Number(10),
//   	ResourceId: jsii.String("replication-group/main-cluster"),
//   })
//
//   shardsScalableTarget.ScaleToTrackMetric(jsii.String("ElastiCacheRedisShardsCPUUtilization"), &BasicTargetTrackingScalingPolicyProps{
//   	TargetValue: jsii.Number(20),
//   	PredefinedMetric: appscaling.PredefinedMetric_ELASTICACHE_PRIMARY_ENGINE_CPU_UTILIZATION,
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

