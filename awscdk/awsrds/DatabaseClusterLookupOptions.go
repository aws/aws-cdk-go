package awsrds


// Properties for looking up an existing DatabaseCluster.
//
// Example:
//   var myUserRole role
//
//
//   clusterFromLookup := rds.DatabaseCluster_FromLookup(this, jsii.String("ClusterFromLookup"), &DatabaseClusterLookupOptions{
//   	ClusterIdentifier: jsii.String("my-cluster-id"),
//   })
//
//   // Grant a connection
//   clusterFromLookup.GrantConnect(myUserRole, jsii.String("my-user-id"))
//
type DatabaseClusterLookupOptions struct {
	// The cluster identifier of the DatabaseCluster.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
}

