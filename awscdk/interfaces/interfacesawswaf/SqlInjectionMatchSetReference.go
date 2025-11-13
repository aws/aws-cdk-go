package interfacesawswaf


// A reference to a SqlInjectionMatchSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqlInjectionMatchSetReference := &SqlInjectionMatchSetReference{
//   	SqlInjectionMatchSetId: jsii.String("sqlInjectionMatchSetId"),
//   }
//
type SqlInjectionMatchSetReference struct {
	// The Id of the SqlInjectionMatchSet resource.
	SqlInjectionMatchSetId *string `field:"required" json:"sqlInjectionMatchSetId" yaml:"sqlInjectionMatchSetId"`
}

