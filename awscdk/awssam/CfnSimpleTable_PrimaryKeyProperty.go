package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   primaryKeyProperty := &PrimaryKeyProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   }
//
type CfnSimpleTable_PrimaryKeyProperty struct {
	// `CfnSimpleTable.PrimaryKeyProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnSimpleTable.PrimaryKeyProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

