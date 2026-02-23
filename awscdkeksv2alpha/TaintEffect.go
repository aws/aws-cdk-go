package awscdkeksv2alpha


// Effect types of kubernetes node taint.
//
// Note: These values are specifically for AWS EKS NodeGroups and use the AWS API format.
// When using AWS CLI or API, taint effects must be NO_SCHEDULE, PREFER_NO_SCHEDULE, or NO_EXECUTE.
// When using Kubernetes directly or kubectl, taint effects must be NoSchedule, PreferNoSchedule, or NoExecute.
//
// For Kubernetes manifests (like Karpenter NodePools), use string literals with PascalCase format:
// - 'NoSchedule' instead of TaintEffect.NO_SCHEDULE
// - 'PreferNoSchedule' instead of TaintEffect.PREFER_NO_SCHEDULE
// - 'NoExecute' instead of TaintEffect.NO_EXECUTE
// See: https://docs.aws.amazon.com/eks/latest/userguide/node-taints-managed-node-groups.html
//
// Deprecated.
type TaintEffect string

const (
	// NoSchedule.
	// Deprecated.
	TaintEffect_NO_SCHEDULE TaintEffect = "NO_SCHEDULE"
	// PreferNoSchedule.
	// Deprecated.
	TaintEffect_PREFER_NO_SCHEDULE TaintEffect = "PREFER_NO_SCHEDULE"
	// NoExecute.
	// Deprecated.
	TaintEffect_NO_EXECUTE TaintEffect = "NO_EXECUTE"
)

