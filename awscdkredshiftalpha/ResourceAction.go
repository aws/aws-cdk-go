package awscdkredshiftalpha


// The Amazon Redshift operation.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdkredshiftalpha"
//
//   var vpc iVpc
//
//
//   // Pause the cluster
//   // Pause the cluster
//   awscdkredshiftalpha.NewCluster(this, jsii.String("PausedCluster"), &ClusterProps{
//   	MasterUser: &Login{
//   		MasterUsername: jsii.String("admin"),
//   	},
//   	Vpc: Vpc,
//   	ResourceAction: awscdkredshiftalpha.ResourceAction_PAUSE,
//   })
//
//   // Resume the cluster
//   // Resume the cluster
//   awscdkredshiftalpha.NewCluster(this, jsii.String("ResumedCluster"), &ClusterProps{
//   	MasterUser: &Login{
//   		MasterUsername: jsii.String("admin"),
//   	},
//   	Vpc: Vpc,
//   	ResourceAction: awscdkredshiftalpha.ResourceAction_RESUME,
//   })
//
//   // Failover the cluster
//   // Failover the cluster
//   awscdkredshiftalpha.NewCluster(this, jsii.String("FailOverCluster"), &ClusterProps{
//   	MasterUser: &Login{
//   		MasterUsername: jsii.String("admin"),
//   	},
//   	// VPC must have 3 AZs for the cluster which executes failover action
//   	Vpc: Vpc,
//   	// Must be a multi-AZ cluster to failover
//   	MultiAz: jsii.Boolean(true),
//   	ResourceAction: awscdkredshiftalpha.ResourceAction_FAILOVER_PRIMARY_COMPUTE,
//   })
//
// Experimental.
type ResourceAction string

const (
	// Pause the cluster.
	// Experimental.
	ResourceAction_PAUSE_CLUSTER ResourceAction = "PAUSE_CLUSTER"
	// Resume the cluster.
	// Experimental.
	ResourceAction_RESUME_CLUSTER ResourceAction = "RESUME_CLUSTER"
	// Failing over to the other availability zone.
	// See: https://docs.aws.amazon.com/redshift/latest/mgmt/test-cluster-multi-az.html
	//
	// Experimental.
	ResourceAction_FAILOVER_PRIMARY_COMPUTE ResourceAction = "FAILOVER_PRIMARY_COMPUTE"
)

