package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDataRepositoryAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataRepositoryAssociationProps := &cfnDataRepositoryAssociationProps{
//   	dataRepositoryPath: jsii.String("dataRepositoryPath"),
//   	fileSystemId: jsii.String("fileSystemId"),
//   	fileSystemPath: jsii.String("fileSystemPath"),
//
//   	// the properties below are optional
//   	batchImportMetaDataOnCreate: jsii.Boolean(false),
//   	importedFileChunkSize: jsii.Number(123),
//   	s3: &s3Property{
//   		autoExportPolicy: &autoExportPolicyProperty{
//   			events: []*string{
//   				jsii.String("events"),
//   			},
//   		},
//   		autoImportPolicy: &autoImportPolicyProperty{
//   			events: []*string{
//   				jsii.String("events"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDataRepositoryAssociationProps struct {
	// The path to the data repository that will be linked to the cache or file system.
	//
	// - For Amazon File Cache, the path can be an NFS data repository that will be linked to the cache. The path can be in one of two formats:
	//
	// - If you are not using the `DataRepositorySubdirectories` parameter, the path is to an NFS Export directory (or one of its subdirectories) in the format `nsf://nfs-domain-name/exportpath` . You can therefore link a single NFS Export to a single data repository association.
	// - If you are using the `DataRepositorySubdirectories` parameter, the path is the domain name of the NFS file system in the format `nfs://filer-domain-name` , which indicates the root of the subdirectories specified with the `DataRepositorySubdirectories` parameter.
	// - For Amazon File Cache, the path can be an S3 bucket or prefix in the format `s3://myBucket/myPrefix/` .
	// - For Amazon FSx for Lustre, the path can be an S3 bucket or prefix in the format `s3://myBucket/myPrefix/` .
	DataRepositoryPath *string `field:"required" json:"dataRepositoryPath" yaml:"dataRepositoryPath"`
	// `AWS::FSx::DataRepositoryAssociation.FileSystemId`.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// A path on the Amazon FSx for Lustre file system that points to a high-level directory (such as `/ns1/` ) or subdirectory (such as `/ns1/subdir/` ) that will be mapped 1-1 with `DataRepositoryPath` .
	//
	// The leading forward slash in the name is required. Two data repository associations cannot have overlapping file system paths. For example, if a data repository is associated with file system path `/ns1/` , then you cannot link another data repository with file system path `/ns1/ns2` .
	//
	// This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.
	//
	// > If you specify only a forward slash ( `/` ) as the file system path, you can link only one data repository to the file system. You can only specify "/" as the file system path for the first data repository associated with a file system.
	FileSystemPath *string `field:"required" json:"fileSystemPath" yaml:"fileSystemPath"`
	// A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created.
	//
	// The task runs if this flag is set to `true` .
	//
	// > `BatchImportMetaDataOnCreate` is not supported for data repositories linked to an Amazon File Cache resource.
	BatchImportMetaDataOnCreate interface{} `field:"optional" json:"batchImportMetaDataOnCreate" yaml:"batchImportMetaDataOnCreate"`
	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.
	//
	// The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system or cache.
	//
	// The default chunk size is 1,024 MiB (1 GiB) and can go as high as 512,000 MiB (500 GiB). Amazon S3 objects have a maximum size of 5 TB.
	ImportedFileChunkSize *float64 `field:"optional" json:"importedFileChunkSize" yaml:"importedFileChunkSize"`
	// The configuration for an Amazon S3 data repository linked to an Amazon FSx for Lustre file system with a data repository association.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// `AWS::FSx::DataRepositoryAssociation.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

