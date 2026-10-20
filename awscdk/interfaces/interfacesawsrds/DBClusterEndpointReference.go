package interfacesawsrds


// A reference to a DBClusterEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBClusterEndpointReference := &DBClusterEndpointReference{
//   	DbClusterEndpointArn: jsii.String("dbClusterEndpointArn"),
//   }
//
type DBClusterEndpointReference struct {
	// The DBClusterEndpointArn of the DBClusterEndpoint resource.
	DbClusterEndpointArn *string `field:"required" json:"dbClusterEndpointArn" yaml:"dbClusterEndpointArn"`
}

