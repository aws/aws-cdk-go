package awsopsworks


// Describes an Elastic IP address.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elasticIpProperty := &elasticIpProperty{
//   	ip: jsii.String("ip"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   }
//
type CfnStack_ElasticIpProperty struct {
	// The IP address.
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// The name, which can be a maximum of 32 characters.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

