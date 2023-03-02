package awsapigatewayv2


// Properties for defining a `CfnVpcLink`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcLinkProps := &cfnVpcLinkProps{
//   	name: jsii.String("name"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnVpcLinkProps struct {
	// The name of the VPC link.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of subnet IDs to include in the VPC link.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// A list of security group IDs for the VPC link.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The collection of tags.
	//
	// Each tag element is associated with a given resource.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

