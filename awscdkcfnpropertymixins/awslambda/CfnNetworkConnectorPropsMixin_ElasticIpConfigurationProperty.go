package awslambda


// Specifies whether Lambda assigns an Elastic IP address to each of the network connector's elastic network interfaces (ENIs).
//
// If you don't specify an Elastic IP configuration, no Elastic IP address is assigned. When you update this configuration, the change applies to the ENIs of the new connector version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   elasticIpConfigurationProperty := &ElasticIpConfigurationProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-networkconnector-elasticipconfiguration.html
//
type CfnNetworkConnectorPropsMixin_ElasticIpConfigurationProperty struct {
	// The Elastic IP address type for the connector's ENIs.
	//
	// With None, Lambda doesn't assign an Elastic IP address. With ServiceManaged, Lambda allocates a service-managed Elastic IP address in your account for each ENI, associates it with that ENI, and releases it when it deletes the ENI. The connector's operator role must allow the ec2:AllocateAddress and ec2:AssociateAddress actions on Elastic IP addresses and on the connector's ENIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-networkconnector-elasticipconfiguration.html#cfn-lambda-networkconnector-elasticipconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

