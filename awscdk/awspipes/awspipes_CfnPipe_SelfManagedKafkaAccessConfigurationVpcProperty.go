package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selfManagedKafkaAccessConfigurationVpcProperty := &selfManagedKafkaAccessConfigurationVpcProperty{
//   	securityGroup: []*string{
//   		jsii.String("securityGroup"),
//   	},
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
type CfnPipe_SelfManagedKafkaAccessConfigurationVpcProperty struct {
	// `CfnPipe.SelfManagedKafkaAccessConfigurationVpcProperty.SecurityGroup`.
	SecurityGroup *[]*string `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// `CfnPipe.SelfManagedKafkaAccessConfigurationVpcProperty.Subnets`.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

