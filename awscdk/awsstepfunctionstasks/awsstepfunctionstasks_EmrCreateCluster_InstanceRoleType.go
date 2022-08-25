package awsstepfunctionstasks


// Instance Role Types.
type EmrCreateCluster_InstanceRoleType string

const (
	// Master Node.
	EmrCreateCluster_InstanceRoleType_MASTER EmrCreateCluster_InstanceRoleType = "MASTER"
	// Core Node.
	EmrCreateCluster_InstanceRoleType_CORE EmrCreateCluster_InstanceRoleType = "CORE"
	// Task Node.
	EmrCreateCluster_InstanceRoleType_TASK EmrCreateCluster_InstanceRoleType = "TASK"
)

