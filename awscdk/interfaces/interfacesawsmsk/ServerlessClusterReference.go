package interfacesawsmsk


// A reference to a ServerlessCluster resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverlessClusterReference := &ServerlessClusterReference{
//   	ServerlessClusterArn: jsii.String("serverlessClusterArn"),
//   }
//
type ServerlessClusterReference struct {
	// The Arn of the ServerlessCluster resource.
	ServerlessClusterArn *string `field:"required" json:"serverlessClusterArn" yaml:"serverlessClusterArn"`
}

