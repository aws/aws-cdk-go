package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFileCachePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnFileCacheMixinProps := &CfnFileCacheMixinProps{
//   	CopyTagsToDataRepositoryAssociations: jsii.Boolean(false),
//   	DataRepositoryAssociations: []interface{}{
//   		&DataRepositoryAssociationProperty{
//   			DataRepositoryPath: jsii.String("dataRepositoryPath"),
//   			FileCachePath: jsii.String("fileCachePath"),
//   		},
//   	},
//   	FileCacheType: jsii.String("fileCacheType"),
//   	FileCacheTypeVersion: jsii.String("fileCacheTypeVersion"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LustreConfiguration: &LustreConfigurationProperty{
//   		DeploymentType: jsii.String("deploymentType"),
//   		LogConfiguration: &LogConfigurationProperty{
//   			Destination: jsii.String("destination"),
//   			Level: jsii.String("level"),
//   		},
//   		MetadataConfiguration: &MetadataConfigurationProperty{
//   			StorageCapacity: jsii.Number(123),
//   		},
//   		MountName: jsii.String("mountName"),
//   		PerUnitStorageThroughput: jsii.Number(123),
//   		WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	StorageCapacity: jsii.Number(123),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
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
type CfnFileCacheMixinProps struct {
	// Whether tags should be copied to data repository associations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-copytagstodatarepositoryassociations
	//
	CopyTagsToDataRepositoryAssociations interface{} `field:"optional" json:"copyTagsToDataRepositoryAssociations" yaml:"copyTagsToDataRepositoryAssociations"`
	// Data repository associations for the cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-datarepositoryassociations
	//
	DataRepositoryAssociations interface{} `field:"optional" json:"dataRepositoryAssociations" yaml:"dataRepositoryAssociations"`
	// The type of cache, which must be LUSTRE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-filecachetype
	//
	FileCacheType *string `field:"optional" json:"fileCacheType" yaml:"fileCacheType"`
	// The Lustre version of the cache, which must be 2.12.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-filecachetypeversion
	//
	FileCacheTypeVersion *string `field:"optional" json:"fileCacheTypeVersion" yaml:"fileCacheTypeVersion"`
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
	// The storage capacity of the cache in gibibytes (GiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-storagecapacity
	//
	StorageCapacity *float64 `field:"optional" json:"storageCapacity" yaml:"storageCapacity"`
	// A list of subnet IDs for the cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Tags to associate with the file cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html#cfn-fsx-filecache-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

