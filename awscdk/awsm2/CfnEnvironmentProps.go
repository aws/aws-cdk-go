package awsm2


// Properties for defining a `CfnEnvironment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentProps := &CfnEnvironmentProps{
//   	EngineType: jsii.String("engineType"),
//   	InstanceType: jsii.String("instanceType"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	HighAvailabilityConfig: &HighAvailabilityConfigProperty{
//   		DesiredCapacity: jsii.Number(123),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
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
type CfnEnvironmentProps struct {
	// The target platform for the runtime environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-enginetype
	//
	EngineType *string `field:"required" json:"engineType" yaml:"engineType"`
	// The instance type of the runtime environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The name of the runtime environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the runtime environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version of the runtime engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-engineversion
	//
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Defines the details of a high availability configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-highavailabilityconfig
	//
	HighAvailabilityConfig interface{} `field:"optional" json:"highAvailabilityConfig" yaml:"highAvailabilityConfig"`
	// The identifier of a customer managed key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html#cfn-m2-environment-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Configures the maintenance window you want for the runtime environment.
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

