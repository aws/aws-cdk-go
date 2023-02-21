package awslightsail


// `Networking` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the public ports and the monthly amount of data transfer allocated for the instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkingProperty := &networkingProperty{
//   	ports: []interface{}{
//   		&portProperty{
//   			accessDirection: jsii.String("accessDirection"),
//   			accessFrom: jsii.String("accessFrom"),
//   			accessType: jsii.String("accessType"),
//   			cidrListAliases: []*string{
//   				jsii.String("cidrListAliases"),
//   			},
//   			cidrs: []*string{
//   				jsii.String("cidrs"),
//   			},
//   			commonName: jsii.String("commonName"),
//   			fromPort: jsii.Number(123),
//   			ipv6Cidrs: []*string{
//   				jsii.String("ipv6Cidrs"),
//   			},
//   			protocol: jsii.String("protocol"),
//   			toPort: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	monthlyTransfer: jsii.Number(123),
//   }
//
type CfnInstance_NetworkingProperty struct {
	// An array of ports to open on the instance.
	Ports interface{} `field:"required" json:"ports" yaml:"ports"`
	// The monthly amount of data transfer, in GB, allocated for the instance.
	MonthlyTransfer *float64 `field:"optional" json:"monthlyTransfer" yaml:"monthlyTransfer"`
}

