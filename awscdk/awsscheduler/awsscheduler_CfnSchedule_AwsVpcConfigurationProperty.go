package awsscheduler


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
type CfnSchedule_AwsVpcConfigurationProperty struct {
	// `CfnSchedule.AwsVpcConfigurationProperty.Subnets`.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// `CfnSchedule.AwsVpcConfigurationProperty.AssignPublicIp`.
	AssignPublicIp *string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// `CfnSchedule.AwsVpcConfigurationProperty.SecurityGroups`.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

