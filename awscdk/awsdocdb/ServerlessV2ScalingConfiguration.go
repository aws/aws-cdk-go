package awsdocdb


// ServerlessV2 scaling configuration for DocumentDB clusters.
//
// Example:
//   var vpc Vpc
//
//   cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	MasterUser: &Login{
//   		Username: jsii.String("myuser"),
//   	},
//   	Vpc: Vpc,
//   	ServerlessV2ScalingConfiguration: &ServerlessV2ScalingConfiguration{
//   		MinCapacity: jsii.Number(0.5),
//   		MaxCapacity: jsii.Number(2),
//   	},
//   	EngineVersion: jsii.String("5.0.0"),
//   })
//
type ServerlessV2ScalingConfiguration struct {
	// The maximum number of DocumentDB capacity units (DCUs) for a DocumentDB instance in a DocumentDB Serverless cluster.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum number of DocumentDB capacity units (DCUs) for a DocumentDB instance in a DocumentDB Serverless cluster.
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
}

