package awsdocdb


// The storage type of the DocDB cluster.
//
// Example:
//   var vpc vpc
//
//
//   cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	MasterUser: &Login{
//   		Username: jsii.String("myuser"),
//   	},
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_MEMORY5, ec2.InstanceSize_LARGE),
//   	Vpc: Vpc,
//   	StorageType: docdb.StorageType_IOPT1,
//   })
//
type StorageType string

const (
	// Standard storage.
	StorageType_STANDARD StorageType = "STANDARD"
	// I/O-optimized storage.
	StorageType_IOPT1 StorageType = "IOPT1"
)

