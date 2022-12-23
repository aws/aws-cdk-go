package awsfsx


// Describes the data repository association's automatic import policy.
//
// The AutoImportPolicy defines how Amazon FSx keeps your file metadata and directory listings up to date by importing changes to your Amazon FSx for Lustre file system as you modify objects in a linked S3 bucket.
//
// The `AutoImportPolicy` is supported only for Amazon FSx for Lustre file systems with the `Persistent_2` deployment type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoImportPolicyProperty := &autoImportPolicyProperty{
//   	events: []*string{
//   		jsii.String("events"),
//   	},
//   }
//
type CfnDataRepositoryAssociation_AutoImportPolicyProperty struct {
	// The `AutoImportPolicy` can have the following event values:.
	//
	// - `NEW` - Amazon FSx automatically imports metadata of files added to the linked S3 bucket that do not currently exist in the FSx file system.
	// - `CHANGED` - Amazon FSx automatically updates file metadata and invalidates existing file content on the file system as files change in the data repository.
	// - `DELETED` - Amazon FSx automatically deletes files on the file system as corresponding files are deleted in the data repository.
	//
	// You can define any combination of event types for your `AutoImportPolicy` .
	Events *[]*string `field:"required" json:"events" yaml:"events"`
}

