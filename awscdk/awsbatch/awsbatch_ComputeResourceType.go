package awsbatch


// Property to specify if the compute environment uses On-Demand, SpotFleet, Fargate, or Fargate Spot compute resources.
//
// Example:
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//
//   spotEnvironment := batch.NewComputeEnvironment(this, jsii.String("MySpotEnvironment"), &computeEnvironmentProps{
//   	computeResources: &computeResources{
//   		type: batch.computeResourceType_SPOT,
//   		bidPercentage: jsii.Number(75),
//   		 // Bids for resources at 75% of the on-demand price
//   		vpc: vpc,
//   	},
//   })
//
// Experimental.
type ComputeResourceType string

const (
	// Resources will be EC2 On-Demand resources.
	// Experimental.
	ComputeResourceType_ON_DEMAND ComputeResourceType = "ON_DEMAND"
	// Resources will be EC2 SpotFleet resources.
	// Experimental.
	ComputeResourceType_SPOT ComputeResourceType = "SPOT"
	// Resources will be Fargate resources.
	// Experimental.
	ComputeResourceType_FARGATE ComputeResourceType = "FARGATE"
	// Resources will be Fargate Spot resources.
	//
	// Fargate Spot uses spare capacity in the AWS cloud to run your fault-tolerant,
	// time-flexible jobs at up to a 70% discount. If AWS needs the resources back,
	// jobs running on Fargate Spot will be interrupted with two minutes of notification.
	// Experimental.
	ComputeResourceType_FARGATE_SPOT ComputeResourceType = "FARGATE_SPOT"
)

