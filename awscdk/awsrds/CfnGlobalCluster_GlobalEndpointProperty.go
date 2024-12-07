package awsrds


// The writer endpoint for the new global database cluster.
//
// This endpoint always points to the writer DB instance in the current primary cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalEndpointProperty := &GlobalEndpointProperty{
//   	Address: jsii.String("address"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-globalcluster-globalendpoint.html
//
type CfnGlobalCluster_GlobalEndpointProperty struct {
	// The writer endpoint for the new global database cluster.
	//
	// This endpoint always points to the writer DB instance in the current primary cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-globalcluster-globalendpoint.html#cfn-rds-globalcluster-globalendpoint-address
	//
	Address *string `field:"optional" json:"address" yaml:"address"`
}

