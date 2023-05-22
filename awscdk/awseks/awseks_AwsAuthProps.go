package awseks


// Configuration props for the AwsAuth construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   awsAuthProps := &AwsAuthProps{
//   	Cluster: cluster,
//   }
//
// Experimental.
type AwsAuthProps struct {
	// The EKS cluster to apply this configuration to.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	Cluster Cluster `field:"required" json:"cluster" yaml:"cluster"`
}

