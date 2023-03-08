package awseks


// Enable or disable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs.
//
// By default, cluster control plane logs aren't exported to CloudWatch Logs. For more information, see [Amazon EKS Cluster control plane logs](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) in the **Amazon EKS User Guide** .
//
// > When updating a resource, you must include this `Logging` property if the previous CloudFormation template of the resource had it. > CloudWatch Logs ingestion, archive storage, and data scanning rates apply to exported control plane logs. For more information, see [CloudWatch Pricing](https://docs.aws.amazon.com/cloudwatch/pricing/) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingProperty := &loggingProperty{
//   	clusterLogging: &clusterLoggingProperty{
//   		enabledTypes: []interface{}{
//   			&loggingTypeConfigProperty{
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
type CfnCluster_LoggingProperty struct {
	// The cluster control plane logging configuration for your cluster.
	ClusterLogging interface{} `field:"optional" json:"clusterLogging" yaml:"clusterLogging"`
}

