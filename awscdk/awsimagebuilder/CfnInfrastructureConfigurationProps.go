package awsimagebuilder


// Properties for defining a `CfnInfrastructureConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInfrastructureConfigurationProps := &CfnInfrastructureConfigurationProps{
//   	InstanceProfileName: jsii.String("instanceProfileName"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	InstanceMetadataOptions: &InstanceMetadataOptionsProperty{
//   		HttpPutResponseHopLimit: jsii.Number(123),
//   		HttpTokens: jsii.String("httpTokens"),
//   	},
//   	InstanceTypes: []*string{
//   		jsii.String("instanceTypes"),
//   	},
//   	KeyPair: jsii.String("keyPair"),
//   	Logging: &LoggingProperty{
//   		S3Logs: &S3LogsProperty{
//   			S3BucketName: jsii.String("s3BucketName"),
//   			S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   		},
//   	},
//   	Placement: &PlacementProperty{
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   		HostId: jsii.String("hostId"),
//   		HostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   		Tenancy: jsii.String("tenancy"),
//   	},
//   	ResourceTags: map[string]*string{
//   		"resourceTagsKey": jsii.String("resourceTags"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SnsTopicArn: jsii.String("snsTopicArn"),
//   	SubnetId: jsii.String("subnetId"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TerminateInstanceOnFailure: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html
//
type CfnInfrastructureConfigurationProps struct {
	// The instance profile to associate with the instance used to customize your Amazon EC2 AMI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-instanceprofilename
	//
	InstanceProfileName *string `field:"required" json:"instanceProfileName" yaml:"instanceProfileName"`
	// The name of the infrastructure configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the infrastructure configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The instance metadata options that you can set for the HTTP requests that pipeline builds use to launch EC2 build and test instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-instancemetadataoptions
	//
	InstanceMetadataOptions interface{} `field:"optional" json:"instanceMetadataOptions" yaml:"instanceMetadataOptions"`
	// The instance types of the infrastructure configuration.
	//
	// You can specify one or more instance types to use for this build. The service will pick one of these instance types based on availability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-instancetypes
	//
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The key pair of the infrastructure configuration.
	//
	// You can use this to log on to and debug the instance used to create your image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-keypair
	//
	KeyPair *string `field:"optional" json:"keyPair" yaml:"keyPair"`
	// The logging configuration of the infrastructure configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-logging
	//
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// The instance placement settings that define where the instances that are launched from your image will run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-placement
	//
	Placement interface{} `field:"optional" json:"placement" yaml:"placement"`
	// The metadata tags to assign to the Amazon EC2 instance that Image Builder launches during the build process.
	//
	// Tags are formatted as key value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-resourcetags
	//
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// The security group IDs to associate with the instance used to customize your Amazon EC2 AMI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The Amazon Resource Name (ARN) for the SNS topic to which we send image build event notifications.
	//
	// > EC2 Image Builder is unable to send notifications to SNS topics that are encrypted using keys from other accounts. The key that is used to encrypt the SNS topic must reside in the account that the Image Builder service runs under.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-snstopicarn
	//
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
	// The subnet ID in which to place the instance used to customize your Amazon EC2 AMI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// The metadata tags to assign to the infrastructure configuration resource that Image Builder creates as output.
	//
	// Tags are formatted as key value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The terminate instance on failure setting of the infrastructure configuration.
	//
	// Set to false if you want Image Builder to retain the instance used to configure your AMI if the build or test phase of your workflow fails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-terminateinstanceonfailure
	//
	TerminateInstanceOnFailure interface{} `field:"optional" json:"terminateInstanceOnFailure" yaml:"terminateInstanceOnFailure"`
}

