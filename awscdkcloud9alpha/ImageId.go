package awscdkcloud9alpha


// The image ID used for creating an Amazon EC2 environment.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var vpc vpc
//
//
//   user := iam.NewUser(this, jsii.String("user"))
//   user.AddManagedPolicy(iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("AWSCloud9Administrator")))
//   cloud9.NewEc2Environment(this, jsii.String("C9Env"), &Ec2EnvironmentProps{
//   	Vpc: Vpc,
//   	ImageId: cloud9.ImageId_AMAZON_LINUX_2,
//
//   	Owner: cloud9.Owner_User(user),
//   })
//
// Experimental.
type ImageId string

const (
	// Create using Amazon Linux 2.
	// Experimental.
	ImageId_AMAZON_LINUX_2 ImageId = "AMAZON_LINUX_2"
	// Create using Amazon Linux 2023.
	// Experimental.
	ImageId_AMAZON_LINUX_2023 ImageId = "AMAZON_LINUX_2023"
	// Create using Ubuntu 18.04.
	// Deprecated: Since Ubuntu 18.04 has ended standard support as of May 31, 2023, we recommend you choose Ubuntu 22.04.
	ImageId_UBUNTU_18_04 ImageId = "UBUNTU_18_04"
	// Create using Ubuntu 22.04.
	// Experimental.
	ImageId_UBUNTU_22_04 ImageId = "UBUNTU_22_04"
)

