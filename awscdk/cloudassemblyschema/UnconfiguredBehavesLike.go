package cloudassemblyschema


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var v1 interface{}
//   var v2 interface{}
//
//   unconfiguredBehavesLike := &UnconfiguredBehavesLike{
//   	V1: v1,
//   	V2: v2,
//   }
//
type UnconfiguredBehavesLike struct {
	// Historical accident, don't use.
	//
	// This value may be present, but it should never be used. The actual value is
	// in the `v2` field, regardless of the version of the CDK library.
	// Default: - ignore.
	//
	V1 interface{} `field:"optional" json:"v1" yaml:"v1"`
	// The value of the flag that produces the same behavior as when the flag is not configured at all.
	//
	// Even though it is called 'v2', this is the official name of this field. In
	// any future versions of CDK (v3, v4, ...), this field will still be called 'v2'.
	//
	// The structure of this field is a historical accident. See the comment on
	// `unconfiguredBehavesLike` for more information.
	// Default: false.
	//
	V2 interface{} `field:"optional" json:"v2" yaml:"v2"`
}

