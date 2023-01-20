package awsrds


// The result of binding a `ProxyTarget` to a `DatabaseProxy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var databaseCluster databaseCluster
//   var databaseInstance databaseInstance
//
//   proxyTargetConfig := &proxyTargetConfig{
//   	engineFamily: jsii.String("engineFamily"),
//
//   	// the properties below are optional
//   	dbClusters: []iDatabaseCluster{
//   		databaseCluster,
//   	},
//   	dbInstances: []iDatabaseInstance{
//   		databaseInstance,
//   	},
//   }
//
type ProxyTargetConfig struct {
	// The engine family of the database instance or cluster this proxy connects with.
	EngineFamily *string `field:"required" json:"engineFamily" yaml:"engineFamily"`
	// The database clusters to which this proxy connects.
	//
	// Either this or `dbInstances` will be set and the other `undefined`.
	DbClusters *[]IDatabaseCluster `field:"optional" json:"dbClusters" yaml:"dbClusters"`
	// The database instances to which this proxy connects.
	//
	// Either this or `dbClusters` will be set and the other `undefined`.
	DbInstances *[]IDatabaseInstance `field:"optional" json:"dbInstances" yaml:"dbInstances"`
}

