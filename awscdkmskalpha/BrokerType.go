package awscdkmskalpha


// The broker type for the cluster.
//
// Example:
//   var vpc Vpc
//
//
//   expressCluster := msk.NewCluster(this, jsii.String("ExpressCluster"), &ClusterProps{
//   	ClusterName: jsii.String("MyExpressCluster"),
//   	KafkaVersion: msk.KafkaVersion_V3_8_X(),
//   	Vpc: Vpc,
//   	BrokerType: msk.BrokerType_EXPRESS,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_M7G, ec2.InstanceSize_XLARGE),
//   })
//
// Experimental.
type BrokerType string

const (
	// Standard brokers provide high-availability guarantees.
	// Experimental.
	BrokerType_STANDARD BrokerType = "STANDARD"
	// Express brokers are a low-cost option for development, testing, and workloads that don't require the high availability guarantees of standard MSK cluster.
	// Experimental.
	BrokerType_EXPRESS BrokerType = "EXPRESS"
)

