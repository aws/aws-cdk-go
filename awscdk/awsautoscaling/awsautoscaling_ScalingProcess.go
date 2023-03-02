package awsautoscaling


// Experimental.
type ScalingProcess string

const (
	// Experimental.
	ScalingProcess_LAUNCH ScalingProcess = "LAUNCH"
	// Experimental.
	ScalingProcess_TERMINATE ScalingProcess = "TERMINATE"
	// Experimental.
	ScalingProcess_HEALTH_CHECK ScalingProcess = "HEALTH_CHECK"
	// Experimental.
	ScalingProcess_REPLACE_UNHEALTHY ScalingProcess = "REPLACE_UNHEALTHY"
	// Experimental.
	ScalingProcess_AZ_REBALANCE ScalingProcess = "AZ_REBALANCE"
	// Experimental.
	ScalingProcess_ALARM_NOTIFICATION ScalingProcess = "ALARM_NOTIFICATION"
	// Experimental.
	ScalingProcess_SCHEDULED_ACTIONS ScalingProcess = "SCHEDULED_ACTIONS"
	// Experimental.
	ScalingProcess_ADD_TO_LOAD_BALANCER ScalingProcess = "ADD_TO_LOAD_BALANCER"
)

