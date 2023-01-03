package awsm2


// Properties for defining a `CfnEnvironment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentProps := &cfnEnvironmentProps{
//   	engineType: jsii.String("engineType"),
//   	instanceType: jsii.String("instanceType"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	engineVersion: jsii.String("engineVersion"),
//   	highAvailabilityConfig: &highAvailabilityConfigProperty{
//   		desiredCapacity: jsii.Number(123),
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	storageConfigurations: []interface{}{
//   		&storageConfigurationProperty{
//   			efs: &efsStorageConfigurationProperty{
//   				fileSystemId: jsii.String("fileSystemId"),
//   				mountPoint: jsii.String("mountPoint"),
//   			},
//   			fsx: &fsxStorageConfigurationProperty{
//   				fileSystemId: jsii.String("fileSystemId"),
//   				mountPoint: jsii.String("mountPoint"),
//   			},
//   		},
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnEnvironmentProps struct {
	// `AWS::M2::Environment.EngineType`.
	EngineType *string `field:"required" json:"engineType" yaml:"engineType"`
	// `AWS::M2::Environment.InstanceType`.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// `AWS::M2::Environment.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::M2::Environment.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::M2::Environment.EngineVersion`.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// `AWS::M2::Environment.HighAvailabilityConfig`.
	HighAvailabilityConfig interface{} `field:"optional" json:"highAvailabilityConfig" yaml:"highAvailabilityConfig"`
	// `AWS::M2::Environment.KmsKeyId`.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// `AWS::M2::Environment.PreferredMaintenanceWindow`.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// `AWS::M2::Environment.PubliclyAccessible`.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// `AWS::M2::Environment.SecurityGroupIds`.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// `AWS::M2::Environment.StorageConfigurations`.
	StorageConfigurations interface{} `field:"optional" json:"storageConfigurations" yaml:"storageConfigurations"`
	// `AWS::M2::Environment.SubnetIds`.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// `AWS::M2::Environment.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

