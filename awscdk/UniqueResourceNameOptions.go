package awscdk


// Options for creating a unique resource name.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   uniqueResourceNameOptions := &UniqueResourceNameOptions{
//   	AllowedSpecialCharacters: jsii.String("allowedSpecialCharacters"),
//   	MaxLength: jsii.Number(123),
//   	Separator: jsii.String("separator"),
//   }
//
type UniqueResourceNameOptions struct {
	// Non-alphanumeric characters allowed in the unique resource name.
	// Default: - none.
	//
	AllowedSpecialCharacters *string `field:"optional" json:"allowedSpecialCharacters" yaml:"allowedSpecialCharacters"`
	// The maximum length of the unique resource name.
	// Default: - 256.
	//
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// The separator used between the path components.
	// Default: - none.
	//
	Separator *string `field:"optional" json:"separator" yaml:"separator"`
}

