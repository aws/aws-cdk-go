package awsstepfunctionstasks


// Spot Timeout Actions.
type EmrCreateCluster_SpotTimeoutAction string

const (
	// SWITCH_TO_ON_DEMAND.
	EmrCreateCluster_SpotTimeoutAction_SWITCH_TO_ON_DEMAND EmrCreateCluster_SpotTimeoutAction = "SWITCH_TO_ON_DEMAND"
	// TERMINATE_CLUSTER.
	EmrCreateCluster_SpotTimeoutAction_TERMINATE_CLUSTER EmrCreateCluster_SpotTimeoutAction = "TERMINATE_CLUSTER"
)

