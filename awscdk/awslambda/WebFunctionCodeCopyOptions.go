package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// Configuration for the Lambda function that copies staged code into its destination bucket.
//
// The function is a stack-level singleton: every copy whose options resolve to the same
// configuration shares one function. Supplying different options produces a separate
// function, since a single Lambda cannot hold two conflicting configurations.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc IVpc
//
//
//   lambda.WebFunctionCode_FromAssetToManagedBucket(path.join(__dirname, jsii.String("my-app")), &WebFunctionManagedBucketAssetOptions{
//   	CopyOptions: &WebFunctionCodeCopyOptions{
//   		Vpc: *Vpc,
//   		VpcSubnets: &SubnetSelection{
//   			SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   		},
//   	},
//   })
//
type WebFunctionCodeCopyOptions struct {
	// The log group the Lambda function sends logs to.
	// Default: - a log group is created in the stack, shared by every copy that uses the same function.
	//
	LogGroup interfacesawslogs.ILogGroupRef `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Execution role for the Lambda function that performs the copy.
	// Default: - a role is created with permissions to read the source object and write the
	// destination object.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security groups for the Lambda function.
	//
	// Only used if `vpc` is supplied.
	// Default: - a dedicated security group is created.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The VPC to run the Lambda function in.
	//
	// Only needed if your organization requires all functions to run inside a VPC. The function
	// calls S3, so the VPC needs either a NAT gateway or an S3 VPC endpoint.
	// Default: - the function is not placed in a VPC.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Which subnets to place the Lambda function in.
	//
	// Only used if `vpc` is supplied.
	// Default: - the VPC's private subnets.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

