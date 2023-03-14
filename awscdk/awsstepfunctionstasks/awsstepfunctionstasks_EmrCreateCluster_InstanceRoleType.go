package awsstepfunctionstasks


// Instance Role Types.
// Experimental.
type EmrCreateCluster_InstanceRoleType string

const (
	// Master Node.
	// Experimental.
	EmrCreateCluster_InstanceRoleType_MASTER EmrCreateCluster_InstanceRoleType = "MASTER"
	// Core Node.
	// Experimental.
	EmrCreateCluster_InstanceRoleType_CORE EmrCreateCluster_InstanceRoleType = "CORE"
	// Task Node.
	// Experimental.
	EmrCreateCluster_InstanceRoleType_TASK EmrCreateCluster_InstanceRoleType = "TASK"
)

