package interfacesawsapprunner


// A reference to a VpcConnector resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConnectorReference := &VpcConnectorReference{
//   	VpcConnectorArn: jsii.String("vpcConnectorArn"),
//   }
//
type VpcConnectorReference struct {
	// The VpcConnectorArn of the VpcConnector resource.
	VpcConnectorArn *string `field:"required" json:"vpcConnectorArn" yaml:"vpcConnectorArn"`
}

