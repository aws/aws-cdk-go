package awsrds


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointProperty := &endpointProperty{
//   	address: jsii.String("address"),
//   	port: jsii.String("port"),
//   }
//
type CfnDBCluster_EndpointProperty struct {
	// `CfnDBCluster.EndpointProperty.Address`.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// `CfnDBCluster.EndpointProperty.Port`.
	Port *string `field:"optional" json:"port" yaml:"port"`
}

