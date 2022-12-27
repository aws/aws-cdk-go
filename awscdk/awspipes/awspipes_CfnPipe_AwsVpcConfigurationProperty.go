package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsVpcConfigurationProperty := &awsVpcConfigurationProperty{
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	assignPublicIp: jsii.String("assignPublicIp"),
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   }
//
type CfnPipe_AwsVpcConfigurationProperty struct {
	// `CfnPipe.AwsVpcConfigurationProperty.Subnets`.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// `CfnPipe.AwsVpcConfigurationProperty.AssignPublicIp`.
	AssignPublicIp *string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// `CfnPipe.AwsVpcConfigurationProperty.SecurityGroups`.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

