package awsrds


// The type of storage.
//
// Example:
//   var vpc vpc
//
//
//   iopsInstance := rds.NewDatabaseInstance(this, jsii.String("IopsInstance"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
//   		Version: rds.MysqlEngineVersion_VER_8_0_30(),
//   	}),
//   	Vpc: Vpc,
//   	StorageType: rds.StorageType_IO1,
//   	Iops: jsii.Number(5000),
//   })
//
//   gp3Instance := rds.NewDatabaseInstance(this, jsii.String("Gp3Instance"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_*Mysql(&MySqlInstanceEngineProps{
//   		Version: rds.MysqlEngineVersion_VER_8_0_30(),
//   	}),
//   	Vpc: Vpc,
//   	AllocatedStorage: jsii.Number(500),
//   	StorageType: rds.StorageType_GP3,
//   	StorageThroughput: jsii.Number(500),
//   })
//
// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html
//
type StorageType string

const (
	// Standard.
	//
	// Amazon RDS supports magnetic storage for backward compatibility. It is recommended to use
	// General Purpose SSD or Provisioned IOPS SSD for any new storage needs.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#CHAP_Storage.Magnetic
	//
	StorageType_STANDARD StorageType = "STANDARD"
	// General purpose SSD (gp2).
	//
	// Baseline performance determined by volume size.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#Concepts.Storage.GeneralSSD
	//
	StorageType_GP2 StorageType = "GP2"
	// General purpose SSD (gp3).
	//
	// Performance scales independently from storage.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#Concepts.Storage.GeneralSSD
	//
	StorageType_GP3 StorageType = "GP3"
	// Provisioned IOPS SSD (io1).
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS
	//
	StorageType_IO1 StorageType = "IO1"
	// Provisioned IOPS SSD (io2).
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS
	//
	StorageType_IO2 StorageType = "IO2"
)

