package awsimagebuilderalpha


// The for loop iterates on a range of integers specified within a boundary outlined by the start and end of the variables.
//
// The iterating values are in the set [start, end] and includes boundary values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   componentDocumentForLoop := &ComponentDocumentForLoop{
//   	End: jsii.Number(123),
//   	Start: jsii.Number(123),
//   	UpdateBy: jsii.Number(123),
//   }
//
// Experimental.
type ComponentDocumentForLoop struct {
	// Ending value of iteration.
	//
	// Does not accept chaining expressions.
	// Experimental.
	End *float64 `field:"required" json:"end" yaml:"end"`
	// Starting value of iteration.
	//
	// Does not accept chaining expressions.
	// Experimental.
	Start *float64 `field:"required" json:"start" yaml:"start"`
	// Difference by which an iterating value is updated through addition.
	//
	// It must be a negative or positive non-zero
	// value. Does not accept chaining expressions.
	// Experimental.
	UpdateBy *float64 `field:"required" json:"updateBy" yaml:"updateBy"`
}

