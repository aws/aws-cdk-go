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
	// `AWS::FSx::DataRepositoryAssociation.DataRepositoryPath`.
	DataRepositoryPath *string `field:"required" json:"dataRepositoryPath" yaml:"dataRepositoryPath"`
	// `AWS::FSx::DataRepositoryAssociation.FileSystemId`.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// `AWS::FSx::DataRepositoryAssociation.FileSystemPath`.
	FileSystemPath *string `field:"required" json:"fileSystemPath" yaml:"fileSystemPath"`
	// `AWS::FSx::DataRepositoryAssociation.BatchImportMetaDataOnCreate`.
	BatchImportMetaDataOnCreate interface{} `field:"optional" json:"batchImportMetaDataOnCreate" yaml:"batchImportMetaDataOnCreate"`
	// `AWS::FSx::DataRepositoryAssociation.ImportedFileChunkSize`.
	ImportedFileChunkSize *float64 `field:"optional" json:"importedFileChunkSize" yaml:"importedFileChunkSize"`
	// `AWS::FSx::DataRepositoryAssociation.S3`.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// `AWS::FSx::DataRepositoryAssociation.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

