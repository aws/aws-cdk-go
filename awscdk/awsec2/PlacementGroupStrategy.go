package awsec2


// Which strategy to use when launching instances.
//
// Example:
//   var instanceType instanceType
//
//
//   pg := ec2.NewPlacementGroup(this, jsii.String("test-pg"), &PlacementGroupProps{
//   	Strategy: ec2.PlacementGroupStrategy_SPREAD,
//   })
//
//   ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2023(),
//   	PlacementGroup: pg,
//   })
//
type PlacementGroupStrategy string

const (
	// Packs instances close together inside an Availability Zone.
	//
	// This strategy enables workloads to achieve the low-latency network
	// performance necessary for tightly-coupled node-to-node communication that
	// is typical of high-performance computing (HPC) applications.
	PlacementGroupStrategy_CLUSTER PlacementGroupStrategy = "CLUSTER"
	// Spreads your instances across logical partitions such that groups of instances in one partition do not share the underlying hardware with groups of instances in different partitions.
	//
	// This strategy is typically used by large distributed and replicated workloads,
	// such as Hadoop, Cassandra, and Kafka.
	PlacementGroupStrategy_PARTITION PlacementGroupStrategy = "PARTITION"
	// Strictly places a small group of instances across distinct underlying hardware to reduce correlated failures.
	PlacementGroupStrategy_SPREAD PlacementGroupStrategy = "SPREAD"
)

