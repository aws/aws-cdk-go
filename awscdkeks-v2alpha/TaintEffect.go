package awscdkeks-v2alpha


// Effect types of kubernetes node taint.
// Experimental.
type TaintEffect string

const (
	// NoSchedule.
	// Experimental.
	TaintEffect_NO_SCHEDULE TaintEffect = "NO_SCHEDULE"
	// PreferNoSchedule.
	// Experimental.
	TaintEffect_PREFER_NO_SCHEDULE TaintEffect = "PREFER_NO_SCHEDULE"
	// NoExecute.
	// Experimental.
	TaintEffect_NO_EXECUTE TaintEffect = "NO_EXECUTE"
)

