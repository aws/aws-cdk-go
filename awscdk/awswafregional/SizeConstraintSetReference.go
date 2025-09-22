package awswafregional


// A reference to a SizeConstraintSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sizeConstraintSetReference := &SizeConstraintSetReference{
//   	SizeConstraintSetId: jsii.String("sizeConstraintSetId"),
//   }
//
type SizeConstraintSetReference struct {
	// The Id of the SizeConstraintSet resource.
	SizeConstraintSetId *string `field:"required" json:"sizeConstraintSetId" yaml:"sizeConstraintSetId"`
}

