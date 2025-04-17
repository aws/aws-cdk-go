package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Options to add the multi user rotation.
//
// Example:
//   var instance databaseInstance
//   var mySecurityGroup securityGroup
//
//
//   instance.addRotationSingleUser(&RotationSingleUserOptions{
//   	AutomaticallyAfter: awscdk.Duration_Days(jsii.Number(7)),
//   	 // defaults to 30 days
//   	ExcludeCharacters: jsii.String("!@#$%^&*"),
//   	 // defaults to the set " %+~`#/// here*()|[]{}:;<>?!'/@\"\\"
//   	SecurityGroup: mySecurityGroup,
//   	 // defaults to an auto-created security group
//   	RotateImmediatelyOnUpdate: jsii.Boolean(false),
//   })
//
type RotationSingleUserOptions struct {
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	// Default: - 30 days.
	//
	AutomaticallyAfter awscdk.Duration `field:"optional" json:"automaticallyAfter" yaml:"automaticallyAfter"`
	// The VPC interface endpoint to use for the Secrets Manager API.
	//
	// If you enable private DNS hostnames for your VPC private endpoint (the default), you don't
	// need to specify an endpoint. The standard Secrets Manager DNS hostname the Secrets Manager
	// CLI and SDKs use by default (https://secretsmanager.<region>.amazonaws.com) automatically
	// resolves to your VPC endpoint.
	// Default: https://secretsmanager.<region>.amazonaws.com
	//
	Endpoint awsec2.IInterfaceVpcEndpoint `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Specifies characters to not include in generated passwords.
	// Default: " %+~`#$&*()|[]{}:;<>?!'/@\"\\"
	//
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.
	// Default: true.
	//
	RotateImmediatelyOnUpdate *bool `field:"optional" json:"rotateImmediatelyOnUpdate" yaml:"rotateImmediatelyOnUpdate"`
	// The security group for the Lambda rotation function.
	// Default: - a new security group is created.
	//
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Where to place the rotation Lambda function.
	// Default: - same placement as instance or cluster.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

