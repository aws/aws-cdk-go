package awseks


// A resource weight entry for the scheduler scoring strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceWeightProperty := &ResourceWeightProperty{
//   	Name: jsii.String("name"),
//   	Weight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourceweight.html
//
type CfnCluster_ResourceWeightProperty struct {
	// The name of the resource (for example, cpu or memory).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourceweight.html#cfn-eks-cluster-resourceweight-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The weight assigned to the resource for scoring.
	//
	// Must be between 1 and 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourceweight.html#cfn-eks-cluster-resourceweight-weight
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

