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
//   endpointProperty := &EndpointProperty{
//   	Address: jsii.String("address"),
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	Port: jsii.String("port"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-endpoint.html
//
type CfnDBInstance_EndpointProperty struct {
	// Specifies the DNS address of the DB instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-endpoint.html#cfn-rds-dbinstance-endpoint-address
	//
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-endpoint.html#cfn-rds-dbinstance-endpoint-hostedzoneid
	//
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// Specifies the port that the database engine is listening on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-endpoint.html#cfn-rds-dbinstance-endpoint-port
	//
	Port *string `field:"optional" json:"port" yaml:"port"`
}

