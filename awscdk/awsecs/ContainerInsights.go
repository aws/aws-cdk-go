package awsecs


// The CloudWatch Container Insights setting.
//
// Example:
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	ContainerInsightsV2: ecs.ContainerInsights_ENHANCED,
//   })
//
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cloudwatch-container-insights.html
//
type ContainerInsights string

const (
	// Enable CloudWatch Container Insights for the cluster.
	ContainerInsights_ENABLED ContainerInsights = "ENABLED"
	// Disable CloudWatch Container Insights for the cluster.
	ContainerInsights_DISABLED ContainerInsights = "DISABLED"
	// Enable CloudWatch Container Insights with enhanced observability for the cluster.
	ContainerInsights_ENHANCED ContainerInsights = "ENHANCED"
)

