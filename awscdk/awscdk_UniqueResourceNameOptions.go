// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Options for creating a unique resource name.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   uniqueResourceNameOptions := &uniqueResourceNameOptions{
//   	allowedSpecialCharacters: jsii.String("allowedSpecialCharacters"),
//   	maxLength: jsii.Number(123),
//   	separator: jsii.String("separator"),
//   }
//
type UniqueResourceNameOptions struct {
	// Non-alphanumeric characters allowed in the unique resource name.
	AllowedSpecialCharacters *string `field:"optional" json:"allowedSpecialCharacters" yaml:"allowedSpecialCharacters"`
	// The maximum length of the unique resource name.
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// The separator used between the path components.
	Separator *string `field:"optional" json:"separator" yaml:"separator"`
}

