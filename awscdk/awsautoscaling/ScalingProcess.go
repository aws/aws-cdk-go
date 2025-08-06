package awsautoscaling


type ScalingProcess string

const (
	ScalingProcess_LAUNCH ScalingProcess = "LAUNCH"
	ScalingProcess_TERMINATE ScalingProcess = "TERMINATE"
	ScalingProcess_HEALTH_CHECK ScalingProcess = "HEALTH_CHECK"
	ScalingProcess_REPLACE_UNHEALTHY ScalingProcess = "REPLACE_UNHEALTHY"
	ScalingProcess_AZ_REBALANCE ScalingProcess = "AZ_REBALANCE"
	ScalingProcess_ALARM_NOTIFICATION ScalingProcess = "ALARM_NOTIFICATION"
	ScalingProcess_SCHEDULED_ACTIONS ScalingProcess = "SCHEDULED_ACTIONS"
	ScalingProcess_ADD_TO_LOAD_BALANCER ScalingProcess = "ADD_TO_LOAD_BALANCER"
	ScalingProcess_INSTANCE_REFRESH ScalingProcess = "INSTANCE_REFRESH"
)

