package awsfsx


// Describes a data repository association's automatic export policy.
//
// The `AutoExportPolicy` defines the types of updated objects on the file system that will be automatically exported to the data repository. As you create, modify, or delete files, Amazon FSx for Lustre automatically exports the defined changes asynchronously once your application finishes modifying the file.
//
// The `AutoExportPolicy` is only supported on Amazon FSx for Lustre file systems with a data repository association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoExportPolicyProperty := &AutoExportPolicyProperty{
//   	Events: []*string{
//   		jsii.String("events"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-datarepositoryassociation-autoexportpolicy.html
//
type CfnDataRepositoryAssociation_AutoExportPolicyProperty struct {
	// The `AutoExportPolicy` can have the following event values:.
	//
	// - `NEW` - New files and directories are automatically exported to the data repository as they are added to the file system.
	// - `CHANGED` - Changes to files and directories on the file system are automatically exported to the data repository.
	// - `DELETED` - Files and directories are automatically deleted on the data repository when they are deleted on the file system.
	//
	// You can define any combination of event types for your `AutoExportPolicy` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-datarepositoryassociation-autoexportpolicy.html#cfn-fsx-datarepositoryassociation-autoexportpolicy-events
	//
	Events *[]*string `field:"required" json:"events" yaml:"events"`
}

