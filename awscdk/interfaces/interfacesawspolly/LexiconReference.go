package interfacesawspolly


// A reference to a Lexicon resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lexiconReference := &LexiconReference{
//   	LexiconArn: jsii.String("lexiconArn"),
//   }
//
type LexiconReference struct {
	// The Arn of the Lexicon resource.
	LexiconArn *string `field:"required" json:"lexiconArn" yaml:"lexiconArn"`
}

