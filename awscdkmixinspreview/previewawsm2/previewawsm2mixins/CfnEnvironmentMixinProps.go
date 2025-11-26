package previewawsm2mixins


// Properties for CfnEnvironmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEnvironmentMixinProps := &CfnEnvironmentMixinProps{
//   	Description: jsii.String("description"),
//   	EngineType: jsii.String("engineType"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	HighAvailabilityConfig: &HighAvailabilityConfigProperty{
//   		DesiredCapacity: jsii.Number(123),
//   	},
//   	InstanceType: jsii.String("instanceType"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Name: jsii.String("name"),
//   	NetworkType: jsii.String("networkType"),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	PubliclyAccessible: jsii.Boolean(false),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	StorageConfigurations: []interface{}{
//   		&StorageConfigurationProperty{
//   			Efs: &EfsStorageConfigurationProperty{
//   				FileSystemId: jsii.String("fileSystemId"),
//   				MountPoint: jsii.String("mountPoint"),
//   			},
//   			Fsx: &FsxStorageConfigurationProperty{
//   				FileSystemId: jsii.String("fileSystemId"),
//   				MountPoint: jsii.String("mountPoint"),
//   			},
//   		},
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html
//
type CfnEnvironmentMixinProps struct {
	// The description of the runtime environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The target platform for the runtime environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-enginetype
	//
	EngineType *string `field:"optional" json:"engineType" yaml:"engineType"`
	// The version of the runtime engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-engineversion
	//
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// > AWS Mainframe Modernization Service (Managed Runtime Environment experience) will no longer be open to new customers starting on November 7, 2025.
	//
	// If you would like to use the service, please sign up prior to November 7, 2025. For capabilities similar to AWS Mainframe Modernization Service (Managed Runtime Environment experience) explore AWS Mainframe Modernization Service (Self-Managed Experience). Existing customers can continue to use the service as normal. For more information, see [AWS Mainframe Modernization availability change](https://docs.aws.amazon.com/m2/latest/userguide/mainframe-modernization-availability-change.html) .
	//
	// Defines the details of a high availability configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-highavailabilityconfig
	//
	HighAvailabilityConfig interface{} `field:"optional" json:"highAvailabilityConfig" yaml:"highAvailabilityConfig"`
	// The instance type of the runtime environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The identifier of a customer managed key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of the runtime environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The network type supported by the runtime environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-networktype
	//
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Configures the maintenance window that you want for the runtime environment.
	//
	// The maintenance window must have the format `ddd:hh24:mi-ddd:hh24:mi` and must be less than 24 hours. The following two examples are valid maintenance windows: `sun:23:45-mon:00:15` or `sat:01:00-sat:03:00` .
	//
	// If you do not provide a value, a random system-generated value will be assigned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-preferredmaintenancewindow
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Specifies whether the runtime environment is publicly accessible.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-publiclyaccessible
	//
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The list of security groups for the VPC associated with this runtime environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// > AWS Mainframe Modernization Service (Managed Runtime Environment experience) will no longer be open to new customers starting on November 7, 2025.
	//
	// If you would like to use the service, please sign up prior to November 7, 2025. For capabilities similar to AWS Mainframe Modernization Service (Managed Runtime Environment experience) explore AWS Mainframe Modernization Service (Self-Managed Experience). Existing customers can continue to use the service as normal. For more information, see [AWS Mainframe Modernization availability change](https://docs.aws.amazon.com/m2/latest/userguide/mainframe-modernization-availability-change.html) .
	//
	// Defines the storage configuration for a runtime environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-storageconfigurations
	//
	StorageConfigurations interface{} `field:"optional" json:"storageConfigurations" yaml:"storageConfigurations"`
	// The list of subnets associated with the VPC for this runtime environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

