package awsiam


// Result of a call to grantOnKey().
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var grant Grant
//
//   grantOnKeyResult := &GrantOnKeyResult{
//   	Grant: grant,
//   }
//
type GrantOnKeyResult struct {
	// The Grant object, if a grant was created.
	// Default: No grant.
	//
	Grant Grant `field:"optional" json:"grant" yaml:"grant"`
}

