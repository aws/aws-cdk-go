package awslightsail


// `Networking` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the public ports and the monthly amount of data transfer allocated for the instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkingProperty := &NetworkingProperty{
//   	Ports: []interface{}{
//   		&PortProperty{
//   			AccessDirection: jsii.String("accessDirection"),
//   			AccessFrom: jsii.String("accessFrom"),
//   			AccessType: jsii.String("accessType"),
//   			CidrListAliases: []*string{
//   				jsii.String("cidrListAliases"),
//   			},
//   			Cidrs: []*string{
//   				jsii.String("cidrs"),
//   			},
//   			CommonName: jsii.String("commonName"),
//   			FromPort: jsii.Number(123),
//   			Ipv6Cidrs: []*string{
//   				jsii.String("ipv6Cidrs"),
//   			},
//   			Protocol: jsii.String("protocol"),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	MonthlyTransfer: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-networking.html
//
type CfnInstance_NetworkingProperty struct {
	// An array of ports to open on the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-networking.html#cfn-lightsail-instance-networking-ports
	//
	Ports interface{} `field:"required" json:"ports" yaml:"ports"`
	// The monthly amount of data transfer, in GB, allocated for the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-networking.html#cfn-lightsail-instance-networking-monthlytransfer
	//
	MonthlyTransfer *float64 `field:"optional" json:"monthlyTransfer" yaml:"monthlyTransfer"`
}

