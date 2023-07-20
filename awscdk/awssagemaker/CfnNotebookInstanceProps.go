package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnNotebookInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNotebookInstanceProps := &CfnNotebookInstanceProps{
//   	InstanceType: jsii.String("instanceType"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	AcceleratorTypes: []*string{
//   		jsii.String("acceleratorTypes"),
//   	},
//   	AdditionalCodeRepositories: []*string{
//   		jsii.String("additionalCodeRepositories"),
//   	},
//   	DefaultCodeRepository: jsii.String("defaultCodeRepository"),
//   	DirectInternetAccess: jsii.String("directInternetAccess"),
//   	InstanceMetadataServiceConfiguration: &InstanceMetadataServiceConfigurationProperty{
//   		MinimumInstanceMetadataServiceVersion: jsii.String("minimumInstanceMetadataServiceVersion"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LifecycleConfigName: jsii.String("lifecycleConfigName"),
//   	NotebookInstanceName: jsii.String("notebookInstanceName"),
//   	PlatformIdentifier: jsii.String("platformIdentifier"),
//   	RootAccess: jsii.String("rootAccess"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetId: jsii.String("subnetId"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VolumeSizeInGb: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html
//
type CfnNotebookInstanceProps struct {
	// The type of ML compute instance to launch for the notebook instance.
	//
	// > Expect some interruption of service if this parameter is changed as CloudFormation stops a notebook instance and starts it up again to update it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// When you send any requests to AWS resources from the notebook instance, SageMaker assumes this role to perform tasks on your behalf.
	//
	// You must grant this role necessary permissions so SageMaker can perform these tasks. The policy must allow the SageMaker service principal (sagemaker.amazonaws.com) permissions to assume this role. For more information, see [SageMaker Roles](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html) .
	//
	// > To be able to pass this role to SageMaker, the caller of this API must have the `iam:PassRole` permission.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A list of Amazon Elastic Inference (EI) instance types to associate with the notebook instance.
	//
	// Currently, only one instance type can be associated with a notebook instance. For more information, see [Using Elastic Inference in Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) .
	//
	// *Valid Values:* `ml.eia1.medium | ml.eia1.large | ml.eia1.xlarge | ml.eia2.medium | ml.eia2.large | ml.eia2.xlarge` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-acceleratortypes
	//
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// An array of up to three Git repositories associated with the notebook instance.
	//
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance. For more information, see [Associating Git Repositories with SageMaker Notebook Instances](https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-additionalcoderepositories
	//
	AdditionalCodeRepositories *[]*string `field:"optional" json:"additionalCodeRepositories" yaml:"additionalCodeRepositories"`
	// The Git repository associated with the notebook instance as its default code repository.
	//
	// This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. When you open a notebook instance, it opens in the directory that contains this repository. For more information, see [Associating Git Repositories with SageMaker Notebook Instances](https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-defaultcoderepository
	//
	DefaultCodeRepository *string `field:"optional" json:"defaultCodeRepository" yaml:"defaultCodeRepository"`
	// Sets whether SageMaker provides internet access to the notebook instance.
	//
	// If you set this to `Disabled` this notebook instance is able to access resources only in your VPC, and is not be able to connect to SageMaker training and endpoint services unless you configure a NAT Gateway in your VPC.
	//
	// For more information, see [Notebook Instances Are Internet-Enabled by Default](https://docs.aws.amazon.com/sagemaker/latest/dg/appendix-additional-considerations.html#appendix-notebook-and-internet-access) . You can set the value of this parameter to `Disabled` only if you set a value for the `SubnetId` parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-directinternetaccess
	//
	DirectInternetAccess *string `field:"optional" json:"directInternetAccess" yaml:"directInternetAccess"`
	// Information on the IMDS configuration of the notebook instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-instancemetadataserviceconfiguration
	//
	InstanceMetadataServiceConfiguration interface{} `field:"optional" json:"instanceMetadataServiceConfiguration" yaml:"instanceMetadataServiceConfiguration"`
	// The Amazon Resource Name (ARN) of a AWS Key Management Service key that SageMaker uses to encrypt data on the storage volume attached to your notebook instance.
	//
	// The KMS key you provide must be enabled. For information, see [Enabling and Disabling Keys](https://docs.aws.amazon.com/kms/latest/developerguide/enabling-keys.html) in the *AWS Key Management Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of a lifecycle configuration to associate with the notebook instance.
	//
	// For information about lifecycle configurations, see [Customize a Notebook Instance](https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html) in the *Amazon SageMaker Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-lifecycleconfigname
	//
	LifecycleConfigName *string `field:"optional" json:"lifecycleConfigName" yaml:"lifecycleConfigName"`
	// The name of the new notebook instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-notebookinstancename
	//
	NotebookInstanceName *string `field:"optional" json:"notebookInstanceName" yaml:"notebookInstanceName"`
	// The platform identifier of the notebook instance runtime environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-platformidentifier
	//
	PlatformIdentifier *string `field:"optional" json:"platformIdentifier" yaml:"platformIdentifier"`
	// Whether root access is enabled or disabled for users of the notebook instance. The default value is `Enabled` .
	//
	// > Lifecycle configurations need root access to be able to set up a notebook instance. Because of this, lifecycle configurations associated with a notebook instance always run with root access even if you disable root access for users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-rootaccess
	//
	RootAccess *string `field:"optional" json:"rootAccess" yaml:"rootAccess"`
	// The VPC security group IDs, in the form sg-xxxxxxxx.
	//
	// The security groups must be for the same VPC as specified in the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnet in a VPC to which you would like to have a connectivity from your ML compute instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) .
	//
	// You can add tags later by using the `CreateTags` API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The size, in GB, of the ML storage volume to attach to the notebook instance.
	//
	// The default value is 5 GB.
	//
	// > Expect some interruption of service if this parameter is changed as CloudFormation stops a notebook instance and starts it up again to update it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-volumesizeingb
	//
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

