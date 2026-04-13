package awssecurityagent


// AWS resource configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   aWSResourcesProperty := &AWSResourcesProperty{
//   	IamRoles: []*string{
//   		jsii.String("iamRoles"),
//   	},
//   	LambdaFunctionArns: []*string{
//   		jsii.String("lambdaFunctionArns"),
//   	},
//   	LogGroups: []*string{
//   		jsii.String("logGroups"),
//   	},
//   	S3Buckets: []*string{
//   		jsii.String("s3Buckets"),
//   	},
//   	SecretArns: []*string{
//   		jsii.String("secretArns"),
//   	},
//   	Vpcs: []interface{}{
//   		&VpcConfigProperty{
//   			SecurityGroupArns: []*string{
//   				jsii.String("securityGroupArns"),
//   			},
//   			SubnetArns: []*string{
//   				jsii.String("subnetArns"),
//   			},
//   			VpcArn: jsii.String("vpcArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-awsresources.html
//
type CfnAgentSpacePropsMixin_AWSResourcesProperty struct {
	// IAM role ARNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-awsresources.html#cfn-securityagent-agentspace-awsresources-iamroles
	//
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// Lambda function ARNs used to retrieve tester credentials for pentests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-awsresources.html#cfn-securityagent-agentspace-awsresources-lambdafunctionarns
	//
	LambdaFunctionArns *[]*string `field:"optional" json:"lambdaFunctionArns" yaml:"lambdaFunctionArns"`
	// CloudWatch log group ARNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-awsresources.html#cfn-securityagent-agentspace-awsresources-loggroups
	//
	LogGroups *[]*string `field:"optional" json:"logGroups" yaml:"logGroups"`
	// S3 bucket ARNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-awsresources.html#cfn-securityagent-agentspace-awsresources-s3buckets
	//
	S3Buckets *[]*string `field:"optional" json:"s3Buckets" yaml:"s3Buckets"`
	// SecretsManager secret ARNs used to store tester credentials for pentests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-awsresources.html#cfn-securityagent-agentspace-awsresources-secretarns
	//
	SecretArns *[]*string `field:"optional" json:"secretArns" yaml:"secretArns"`
	// VPC configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-awsresources.html#cfn-securityagent-agentspace-awsresources-vpcs
	//
	Vpcs interface{} `field:"optional" json:"vpcs" yaml:"vpcs"`
}

