package awsimagebuilderalpha


// The looping construct of a component defines a repeated sequence of instructions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   componentDocumentLoop := &ComponentDocumentLoop{
//   	For: &ComponentDocumentForLoop{
//   		End: jsii.Number(123),
//   		Start: jsii.Number(123),
//   		UpdateBy: jsii.Number(123),
//   	},
//   	ForEach: []*string{
//   		jsii.String("forEach"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// Experimental.
type ComponentDocumentLoop struct {
	// The for loop iterates on a range of integers specified within a boundary outlined by the start and end of the variables.
	// Default: - none if `forEach` is provided. otherwise, `for` is required.
	//
	// Experimental.
	For *ComponentDocumentForLoop `field:"optional" json:"for" yaml:"for"`
	// The forEach loop iterates on an explicit list of values, which can be strings and chained expressions.
	// Default: - none if `for` is provided. otherwise, `forEach` is required.
	//
	// Experimental.
	ForEach *[]*string `field:"optional" json:"forEach" yaml:"forEach"`
	// The name of the loop, which can be used to reference.
	// Default: loop.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

