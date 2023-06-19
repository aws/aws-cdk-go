package cloudassemblyschema


// Artifact properties for nested cloud assemblies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nestedCloudAssemblyProperties := &NestedCloudAssemblyProperties{
//   	DirectoryName: jsii.String("directoryName"),
//
//   	// the properties below are optional
//   	DisplayName: jsii.String("displayName"),
//   }
//
// Experimental.
type NestedCloudAssemblyProperties struct {
	// Relative path to the nested cloud assembly.
	// Experimental.
	DirectoryName *string `field:"required" json:"directoryName" yaml:"directoryName"`
	// Display name for the cloud assembly.
	// Experimental.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

