package awsstepfunctionstasks


// The Cluster ScaleDownBehavior specifies the way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_RunJobFlow.html#EMR-RunJobFlow-request-ScaleDownBehavior
//
type EmrCreateCluster_EmrClusterScaleDownBehavior string

const (
	// Indicates that Amazon EMR terminates nodes at the instance-hour boundary, regardless of when the request to terminate the instance was submitted.
	//
	// This option is only available with Amazon EMR 5.1.0 and later and is the default for clusters created using that version
	EmrCreateCluster_EmrClusterScaleDownBehavior_TERMINATE_AT_INSTANCE_HOUR EmrCreateCluster_EmrClusterScaleDownBehavior = "TERMINATE_AT_INSTANCE_HOUR"
	// Indicates that Amazon EMR adds nodes to a deny list and drains tasks from nodes before terminating the Amazon EC2 instances, regardless of the instance-hour boundary.
	EmrCreateCluster_EmrClusterScaleDownBehavior_TERMINATE_AT_TASK_COMPLETION EmrCreateCluster_EmrClusterScaleDownBehavior = "TERMINATE_AT_TASK_COMPLETION"
)

