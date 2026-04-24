package cloudassemblyschema


// Metadata type of a PropertyMutation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   propertyMutationMetadataEntry := &PropertyMutationMetadataEntry{
//   	PropertyName: jsii.String("propertyName"),
//   	StackTrace: []*string{
//   		jsii.String("stackTrace"),
//   	},
//   }
//
type PropertyMutationMetadataEntry struct {
	// Name of the property.
	PropertyName *string `field:"required" json:"propertyName" yaml:"propertyName"`
	// Stack trace of the mutation.
	StackTrace *[]*string `field:"required" json:"stackTrace" yaml:"stackTrace"`
}

