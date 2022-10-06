package awsrds


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   readEndpointProperty := &readEndpointProperty{
//   	address: jsii.String("address"),
//   }
//
type CfnDBCluster_ReadEndpointProperty struct {
	// `CfnDBCluster.ReadEndpointProperty.Address`.
	Address *string `field:"optional" json:"address" yaml:"address"`
}

