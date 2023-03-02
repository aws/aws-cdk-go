package awsfsx


// The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association.
//
// The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3Property := &s3Property{
//   	autoExportPolicy: &autoExportPolicyProperty{
//   		events: []*string{
//   			jsii.String("events"),
//   		},
//   	},
//   	autoImportPolicy: &autoImportPolicyProperty{
//   		events: []*string{
//   			jsii.String("events"),
//   		},
//   	},
//   }
//
type CfnDataRepositoryAssociation_S3Property struct {
	// Describes a data repository association's automatic export policy.
	//
	// The `AutoExportPolicy` defines the types of updated objects on the file system that will be automatically exported to the data repository. As you create, modify, or delete files, Amazon FSx for Lustre automatically exports the defined changes asynchronously once your application finishes modifying the file.
	//
	// This `AutoExportPolicy` is supported only for Amazon FSx for Lustre file systems with the `Persistent_2` deployment type.
	AutoExportPolicy interface{} `field:"optional" json:"autoExportPolicy" yaml:"autoExportPolicy"`
	// Describes the data repository association's automatic import policy.
	//
	// The AutoImportPolicy defines how Amazon FSx keeps your file metadata and directory listings up to date by importing changes to your Amazon FSx for Lustre file system as you modify objects in a linked S3 bucket.
	//
	// The `AutoImportPolicy` is supported only for Amazon FSx for Lustre file systems with the `Persistent_2` deployment type.
	AutoImportPolicy interface{} `field:"optional" json:"autoImportPolicy" yaml:"autoImportPolicy"`
}

