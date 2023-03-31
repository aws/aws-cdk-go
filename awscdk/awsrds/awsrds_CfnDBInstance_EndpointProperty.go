package awsrds


// This data type represents the information you need to connect to an Amazon RDS DB instance.
//
// This data type is used as a response element in the following actions:
//
// - `CreateDBInstance`
// - `DescribeDBInstances`
// - `DeleteDBInstance`
//
// For the data structure that represents Amazon Aurora DB cluster endpoints, see `DBClusterEndpoint` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointProperty := &endpointProperty{
//   	address: jsii.String("address"),
//   	hostedZoneId: jsii.String("hostedZoneId"),
//   	port: jsii.String("port"),
//   }
//
type CfnDBInstance_EndpointProperty struct {
	// Specifies the DNS address of the DB instance.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// Specifies the port that the database engine is listening on.
	Port *string `field:"optional" json:"port" yaml:"port"`
}

