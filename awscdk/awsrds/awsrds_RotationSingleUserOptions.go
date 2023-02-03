package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Options to add the multi user rotation.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var instance databaseInstance
//   var mySecurityGroup securityGroup
//
//   instance.addRotationSingleUser(&rotationSingleUserOptions{
//   	automaticallyAfter: cdk.duration.days(jsii.Number(7)),
//   	 // defaults to 30 days
//   	excludeCharacters: jsii.String("!@#$%^&*"),
//   	 // defaults to the set " %+~`#/// here*()|[]{}:;<>?!'/@\"\\"
//   	securityGroup: mySecurityGroup,
//   })
//
type RotationSingleUserOptions struct {
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

