// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha


// The options to add a field to an Intermediate Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appsync_alpha "github.com/aws/aws-cdk-go/awscdkappsyncalpha"
//
//   var field field
//
//   addFieldOptions := &addFieldOptions{
//   	field: field,
//   	fieldName: jsii.String("fieldName"),
//   }
//
// Experimental.
type AddFieldOptions struct {
	// The resolvable field to add.
	//
	// This option must be configured for Object, Interface,
	// Input and Union Types.
	// Experimental.
	Field IField `field:"optional" json:"field" yaml:"field"`
	// The name of the field.
	//
	// This option must be configured for Object, Interface,
	// Input and Enum Types.
	// Experimental.
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
}

