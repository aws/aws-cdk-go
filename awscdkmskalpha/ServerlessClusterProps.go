package awscdkmskalpha


// Properties for a MSK Serverless Cluster.
//
// Example:
//   var vpc vpc
//
//
//   serverlessCluster := msk.NewServerlessCluster(this, jsii.String("ServerlessCluster"), &ServerlessClusterProps{
//   	ClusterName: jsii.String("MyServerlessCluster"),
//   	VpcConfigs: []vpcConfig{
//   		&vpcConfig{
//   			Vpc: *Vpc,
//   		},
//   	},
//   })
//
// Experimental.
type ServerlessClusterProps struct {
	// The configuration of the Amazon VPCs for the cluster.
	//
	// You can specify up to 5 VPC configurations.
	// Experimental.
	VpcConfigs *[]*VpcConfig `field:"required" json:"vpcConfigs" yaml:"vpcConfigs"`
	// The physical name of the cluster.
	// Default: - auto generate.
	//
	// Experimental.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
}

