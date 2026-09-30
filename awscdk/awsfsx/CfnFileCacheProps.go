package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFileCache`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFileCacheProps := &CfnFileCacheProps{
//   	FileCacheType: jsii.String("fileCacheType"),
//   	FileCacheTypeVersion: jsii.String("fileCacheTypeVersion"),
//   	StorageCapacity: jsii.Number(123),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	CopyTagsToDataRepositoryAssociations: jsii.Boolean(false),
//   	DataRepositoryAssociations: []interface{}{
//   		&DataRepositoryAssociationProperty{
//   			DataRepositoryPath: jsii.String("dataRepositoryPath"),
//   			FileCachePath: jsii.String("fileCachePath"),
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LustreConfiguration: &LustreConfigurationProperty{
//   		DeploymentType: jsii.String("deploymentType"),
//   		MetadataConfiguration: &MetadataConfigurationProperty{
//   			StorageCapacity: jsii.Number(123),
//   		},
//   		PerUnitStorageThroughput: jsii.Number(123),
//
//   		// the properties below are optional
//   		LogConfiguration: &LogConfigurationProperty{
//   			Level: jsii.String("level"),
//
//   			// the properties below are optional
//   			Destination: jsii.String("destination"),
//   		},
//   		MountName: jsii.String("mountName"),
//   		WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html
//
type CfnFileCacheProps struct {
	// The type of cache, which must be LUSTRE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-filecachetype
	//
	FileCacheType *string `field:"required" json:"fileCacheType" yaml:"fileCacheType"`
	// The Lustre version of the cache, which must be 2.12.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-filecachetypeversion
	//
	FileCacheTypeVersion *string `field:"required" json:"fileCacheTypeVersion" yaml:"fileCacheTypeVersion"`
	// The storage capacity of the cache in gibibytes (GiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-storagecapacity
	//
	StorageCapacity *float64 `field:"required" json:"storageCapacity" yaml:"storageCapacity"`
	// A list of subnet IDs for the cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Whether tags should be copied to data repository associations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-copytagstodatarepositoryassociations
	//
	CopyTagsToDataRepositoryAssociations interface{} `field:"optional" json:"copyTagsToDataRepositoryAssociations" yaml:"copyTagsToDataRepositoryAssociations"`
	// Data repository associations for the cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-datarepositoryassociations
	//
	DataRepositoryAssociations interface{} `field:"optional" json:"dataRepositoryAssociations" yaml:"dataRepositoryAssociations"`
	// The KMS key ID for encrypting data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-lustreconfiguration
	//
	LustreConfiguration interface{} `field:"optional" json:"lustreConfiguration" yaml:"lustreConfiguration"`
	// A list of security group IDs for the cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Tags to associate with the file cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

