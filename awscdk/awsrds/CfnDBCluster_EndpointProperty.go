package awsrds


// The `Endpoint` return value specifies the connection endpoint for the primary instance of the DB cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointProperty := &EndpointProperty{
//   	Address: jsii.String("address"),
//   	Port: jsii.String("port"),
//   }
//
type CfnDBCluster_EndpointProperty struct {
	// Specifies the connection endpoint for the primary instance of the DB cluster.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Specifies the port that the database engine is listening on.
	Port *string `field:"optional" json:"port" yaml:"port"`
}

