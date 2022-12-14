package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties common to single-user and multi-user rotation options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var interfaceVpcEndpoint interfaceVpcEndpoint
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//
//   commonRotationUserOptions := &commonRotationUserOptions{
//   	automaticallyAfter: cdk.duration.minutes(jsii.Number(30)),
//   	endpoint: interfaceVpcEndpoint,
//   	excludeCharacters: jsii.String("excludeCharacters"),
//   	securityGroup: securityGroup,
//   	vpcSubnets: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//   	},
//   }
//
type CommonRotationUserOptions struct {
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	AutomaticallyAfter awscdk.Duration `field:"optional" json:"automaticallyAfter" yaml:"automaticallyAfter"`
	// The VPC interface endpoint to use for the Secrets Manager API.
	//
	// If you enable private DNS hostnames for your VPC private endpoint (the default), you don't
	// need to specify an endpoint. The standard Secrets Manager DNS hostname the Secrets Manager
	// CLI and SDKs use by default (https://secretsmanager.<region>.amazonaws.com) automatically
	// resolves to your VPC endpoint.
	Endpoint awsec2.IInterfaceVpcEndpoint `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Specifies characters to not include in generated passwords.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// The security group for the Lambda rotation function.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Where to place the rotation Lambda function.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

