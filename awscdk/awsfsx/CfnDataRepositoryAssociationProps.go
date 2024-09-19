package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataRepositoryAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataRepositoryAssociationProps := &CfnDataRepositoryAssociationProps{
//   	DataRepositoryPath: jsii.String("dataRepositoryPath"),
//   	FileSystemId: jsii.String("fileSystemId"),
//   	FileSystemPath: jsii.String("fileSystemPath"),
//
//   	// the properties below are optional
//   	BatchImportMetaDataOnCreate: jsii.Boolean(false),
//   	ImportedFileChunkSize: jsii.Number(123),
//   	S3: &S3Property{
//   		AutoExportPolicy: &AutoExportPolicyProperty{
//   			Events: []*string{
//   				jsii.String("events"),
//   			},
//   		},
//   		AutoImportPolicy: &AutoImportPolicyProperty{
//   			Events: []*string{
//   				jsii.String("events"),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-datarepositoryassociation.html
//
type CfnDataRepositoryAssociationProps struct {
	// The path to the Amazon S3 data repository that will be linked to the file system.
	//
	// The path can be an S3 bucket or prefix in the format `s3://bucket-name/prefix/` . This path specifies where in the S3 data repository files will be imported from or exported to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-datarepositoryassociation.html#cfn-fsx-datarepositoryassociation-datarepositorypath
	//
	DataRepositoryPath *string `field:"required" json:"dataRepositoryPath" yaml:"dataRepositoryPath"`
	// The ID of the file system on which the data repository association is configured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-datarepositoryassociation.html#cfn-fsx-datarepositoryassociation-filesystemid
	//
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// A path on the Amazon FSx for Lustre file system that points to a high-level directory (such as `/ns1/` ) or subdirectory (such as `/ns1/subdir/` ) that will be mapped 1-1 with `DataRepositoryPath` .
	//
	// The leading forward slash in the name is required. Two data repository associations cannot have overlapping file system paths. For example, if a data repository is associated with file system path `/ns1/` , then you cannot link another data repository with file system path `/ns1/ns2` .
	//
	// This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.
	//
	// > If you specify only a forward slash ( `/` ) as the file system path, you can link only one data repository to the file system. You can only specify "/" as the file system path for the first data repository associated with a file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-datarepositoryassociation.html#cfn-fsx-datarepositoryassociation-filesystempath
	//
	FileSystemPath *string `field:"required" json:"fileSystemPath" yaml:"fileSystemPath"`
	// A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created.
	//
	// The task runs if this flag is set to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-datarepositoryassociation.html#cfn-fsx-datarepositoryassociation-batchimportmetadataoncreate
	//
	BatchImportMetaDataOnCreate interface{} `field:"optional" json:"batchImportMetaDataOnCreate" yaml:"batchImportMetaDataOnCreate"`
	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.
	//
	// The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system or cache.
	//
	// The default chunk size is 1,024 MiB (1 GiB) and can go as high as 512,000 MiB (500 GiB). Amazon S3 objects have a maximum size of 5 TB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-datarepositoryassociation.html#cfn-fsx-datarepositoryassociation-importedfilechunksize
	//
	ImportedFileChunkSize *float64 `field:"optional" json:"importedFileChunkSize" yaml:"importedFileChunkSize"`
	// The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association.
	//
	// The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-datarepositoryassociation.html#cfn-fsx-datarepositoryassociation-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// A list of `Tag` values, with a maximum of 50 elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-datarepositoryassociation.html#cfn-fsx-datarepositoryassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

