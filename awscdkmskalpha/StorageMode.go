package awscdkmskalpha


// The storage mode for the cluster brokers.
//
// Example:
//   var vpc vpc
//   var bucket iBucket
//
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &ClusterProps{
//   	ClusterName: jsii.String("myCluster"),
//   	KafkaVersion: msk.KafkaVersion_V3_6_0(),
//   	Vpc: Vpc,
//   	StorageMode: msk.StorageMode_TIERED,
//   })
//
// Experimental.
type StorageMode string

const (
	// Local storage mode utilizes network attached EBS storage.
	// Experimental.
	StorageMode_LOCAL StorageMode = "LOCAL"
	// Tiered storage mode utilizes EBS storage and Tiered storage.
	// Experimental.
	StorageMode_TIERED StorageMode = "TIERED"
)

