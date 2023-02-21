package awsecs


type ContainerDependencyCondition string

const (
	// This condition emulates the behavior of links and volumes today.
	//
	// It validates that a dependent container is started before permitting other containers to start.
	ContainerDependencyCondition_START ContainerDependencyCondition = "START"
	// This condition validates that a dependent container runs to completion (exits) before permitting other containers to start.
	//
	// This can be useful for nonessential containers that run a script and then exit.
	ContainerDependencyCondition_COMPLETE ContainerDependencyCondition = "COMPLETE"
	// This condition is the same as COMPLETE, but it also requires that the container exits with a zero status.
	ContainerDependencyCondition_SUCCESS ContainerDependencyCondition = "SUCCESS"
	// This condition validates that the dependent container passes its Docker health check before permitting other containers to start.
	//
	// This requires that the dependent container has health checks configured. This condition is confirmed only at task startup.
	ContainerDependencyCondition_HEALTHY ContainerDependencyCondition = "HEALTHY"
)

