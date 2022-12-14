package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &networkConfigurationProperty{
//   	awsvpcConfiguration: &awsVpcConfigurationProperty{
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//
//   		// the properties below are optional
//   		assignPublicIp: jsii.String("assignPublicIp"),
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   	},
//   }
//
type CfnPipe_NetworkConfigurationProperty struct {
	// `CfnPipe.NetworkConfigurationProperty.AwsvpcConfiguration`.
	AwsvpcConfiguration interface{} `field:"optional" json:"awsvpcConfiguration" yaml:"awsvpcConfiguration"`
}

