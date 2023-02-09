// The CDK Construct Library for AWS::Cloud9
package awscdkcloud9alpha


// The image ID used for creating an Amazon EC2 environment.
//
// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var vpc vpc
//
//
//   user := iam.NewUser(this, jsii.String("user"))
//   user.addManagedPolicy(iam.managedPolicy.fromAwsManagedPolicyName(jsii.String("AWSCloud9Administrator")))
//   cloud9.NewEc2Environment(this, jsii.String("C9Env"), &ec2EnvironmentProps{
//   	vpc: vpc,
//   	imageId: cloud9.imageId_AMAZON_LINUX_2,
//
//   	owner: cloud9.owner.user(user),
//   })
//
// Experimental.
type ImageId string

const (
	// Create using Amazon Linux 2.
	// Experimental.
	ImageId_AMAZON_LINUX_2 ImageId = "AMAZON_LINUX_2"
	// Create using Ubunut 18.04.
	// Experimental.
	ImageId_UBUNTU_18_04 ImageId = "UBUNTU_18_04"
)

