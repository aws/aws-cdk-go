package awsstepfunctionstasks


// AutoScaling Adjustment Type.
type EmrCreateCluster_ScalingAdjustmentType string

const (
	// CHANGE_IN_CAPACITY.
	EmrCreateCluster_ScalingAdjustmentType_CHANGE_IN_CAPACITY EmrCreateCluster_ScalingAdjustmentType = "CHANGE_IN_CAPACITY"
	// PERCENT_CHANGE_IN_CAPACITY.
	EmrCreateCluster_ScalingAdjustmentType_PERCENT_CHANGE_IN_CAPACITY EmrCreateCluster_ScalingAdjustmentType = "PERCENT_CHANGE_IN_CAPACITY"
	// EXACT_CAPACITY.
	EmrCreateCluster_ScalingAdjustmentType_EXACT_CAPACITY EmrCreateCluster_ScalingAdjustmentType = "EXACT_CAPACITY"
)

