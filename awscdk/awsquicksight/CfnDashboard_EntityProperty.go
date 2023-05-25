package awsquicksight


// An object, structure, or sub-structure of an analysis, template, or dashboard.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityProperty := &EntityProperty{
//   	Path: jsii.String("path"),
//   }
//
type CfnDashboard_EntityProperty struct {
	// The hierarchical path of the entity within the analysis, template, or dashboard definition tree.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

