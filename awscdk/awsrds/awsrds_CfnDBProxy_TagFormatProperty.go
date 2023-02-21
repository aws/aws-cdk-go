package awsrds


// Metadata assigned to a DB proxy consisting of a key-value pair.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagFormatProperty := &tagFormatProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnDBProxy_TagFormatProperty struct {
	// A key is the required name of the tag.
	//
	// The string value can be 1-128 Unicode characters in length and can't be prefixed with `aws:` . The string can contain only the set of Unicode letters, digits, white-space, '_', '.', '/', '=', '+', '-' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-]*)$").
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A value is the optional value of the tag.
	//
	// The string value can be 1-256 Unicode characters in length and can't be prefixed with `aws:` . The string can contain only the set of Unicode letters, digits, white-space, '_', '.', '/', '=', '+', '-' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-]*)$").
	Value *string `field:"optional" json:"value" yaml:"value"`
}

