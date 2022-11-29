package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnNotebookInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNotebookInstanceProps := &cfnNotebookInstanceProps{
//   	instanceType: jsii.String("instanceType"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	acceleratorTypes: []*string{
//   		jsii.String("acceleratorTypes"),
//   	},
//   	additionalCodeRepositories: []*string{
//   		jsii.String("additionalCodeRepositories"),
//   	},
//   	defaultCodeRepository: jsii.String("defaultCodeRepository"),
//   	directInternetAccess: jsii.String("directInternetAccess"),
//   	instanceMetadataServiceConfiguration: &instanceMetadataServiceConfigurationProperty{
//   		minimumInstanceMetadataServiceVersion: jsii.String("minimumInstanceMetadataServiceVersion"),
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	lifecycleConfigName: jsii.String("lifecycleConfigName"),
//   	notebookInstanceName: jsii.String("notebookInstanceName"),
//   	platformIdentifier: jsii.String("platformIdentifier"),
//   	rootAccess: jsii.String("rootAccess"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetId: jsii.String("subnetId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	volumeSizeInGb: jsii.Number(123),
//   }
//
type CfnNotebookInstanceProps struct {
	// The type of ML compute instance to launch for the notebook instance.
	//
	// > Expect some interruption of service if this parameter is changed as CloudFormation stops a notebook instance and starts it up again to update it.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// When you send any requests to AWS resources from the notebook instance, SageMaker assumes this role to perform tasks on your behalf.
	//
	// You must grant this role necessary permissions so SageMaker can perform these tasks. The policy must allow the SageMaker service principal (sagemaker.amazonaws.com) permissions to assume this role. For more information, see [SageMaker Roles](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html) .
	//
	// > To be able to pass this role to SageMaker, the caller of this API must have the `iam:PassRole` permission.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A list of Amazon Elastic Inference (EI) instance types to associate with the notebook instance.
	//
	// Currently, only one instance type can be associated with a notebook instance. For more information, see [Using Elastic Inference in Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) .
	//
	// *Valid Values:* `ml.eia1.medium | ml.eia1.large | ml.eia1.xlarge | ml.eia2.medium | ml.eia2.large | ml.eia2.xlarge` .
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// An array of up to three Git repositories associated with the notebook instance.
	//
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance. For more information, see [Associating Git Repositories with SageMaker Notebook Instances](https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html) .
	AdditionalCodeRepositories *[]*string `field:"optional" json:"additionalCodeRepositories" yaml:"additionalCodeRepositories"`
	// The Git repository associated with the notebook instance as its default code repository.
	//
	// This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. When you open a notebook instance, it opens in the directory that contains this repository. For more information, see [Associating Git Repositories with SageMaker Notebook Instances](https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html) .
	DefaultCodeRepository *string `field:"optional" json:"defaultCodeRepository" yaml:"defaultCodeRepository"`
	// Sets whether SageMaker provides internet access to the notebook instance.
	//
	// If you set this to `Disabled` this notebook instance is able to access resources only in your VPC, and is not be able to connect to SageMaker training and endpoint services unless you configure a NAT Gateway in your VPC.
	//
	// For more information, see [Notebook Instances Are Internet-Enabled by Default](https://docs.aws.amazon.com/sagemaker/latest/dg/appendix-additional-considerations.html#appendix-notebook-and-internet-access) . You can set the value of this parameter to `Disabled` only if you set a value for the `SubnetId` parameter.
	DirectInternetAccess *string `field:"optional" json:"directInternetAccess" yaml:"directInternetAccess"`
	// `AWS::SageMaker::NotebookInstance.InstanceMetadataServiceConfiguration`.
	InstanceMetadataServiceConfiguration interface{} `field:"optional" json:"instanceMetadataServiceConfiguration" yaml:"instanceMetadataServiceConfiguration"`
	// The Amazon Resource Name (ARN) of a AWS Key Management Service key that SageMaker uses to encrypt data on the storage volume attached to your notebook instance.
	//
	// The KMS key you provide must be enabled. For information, see [Enabling and Disabling Keys](https://docs.aws.amazon.com/kms/latest/developerguide/enabling-keys.html) in the *AWS Key Management Service Developer Guide* .
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of a lifecycle configuration to associate with the notebook instance.
	//
	// For information about lifecycle configurations, see [Customize a Notebook Instance](https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html) in the *Amazon SageMaker Developer Guide* .
	LifecycleConfigName *string `field:"optional" json:"lifecycleConfigName" yaml:"lifecycleConfigName"`
	// The name of the new notebook instance.
	NotebookInstanceName *string `field:"optional" json:"notebookInstanceName" yaml:"notebookInstanceName"`
	// The platform identifier of the notebook instance runtime environment.
	PlatformIdentifier *string `field:"optional" json:"platformIdentifier" yaml:"platformIdentifier"`
	// Whether root access is enabled or disabled for users of the notebook instance. The default value is `Enabled` .
	//
	// > Lifecycle configurations need root access to be able to set up a notebook instance. Because of this, lifecycle configurations associated with a notebook instance always run with root access even if you disable root access for users.
	RootAccess *string `field:"optional" json:"rootAccess" yaml:"rootAccess"`
	// The VPC security group IDs, in the form sg-xxxxxxxx.
	//
	// The security groups must be for the same VPC as specified in the subnet.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnet in a VPC to which you would like to have a connectivity from your ML compute instance.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) .
	//
	// You can add tags later by using the `CreateTags` API.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The size, in GB, of the ML storage volume to attach to the notebook instance.
	//
	// The default value is 5 GB.
	//
	// > Expect some interruption of service if this parameter is changed as CloudFormation stops a notebook instance and starts it up again to update it.
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

