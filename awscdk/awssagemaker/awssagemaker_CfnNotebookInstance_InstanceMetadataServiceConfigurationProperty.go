package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceMetadataServiceConfigurationProperty := &instanceMetadataServiceConfigurationProperty{
//   	minimumInstanceMetadataServiceVersion: jsii.String("minimumInstanceMetadataServiceVersion"),
//   }
//
type CfnNotebookInstance_InstanceMetadataServiceConfigurationProperty struct {
	// `CfnNotebookInstance.InstanceMetadataServiceConfigurationProperty.MinimumInstanceMetadataServiceVersion`.
	MinimumInstanceMetadataServiceVersion *string `field:"required" json:"minimumInstanceMetadataServiceVersion" yaml:"minimumInstanceMetadataServiceVersion"`
}

