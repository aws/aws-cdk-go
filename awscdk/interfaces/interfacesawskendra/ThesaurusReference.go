package interfacesawskendra


// A reference to a Thesaurus resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thesaurusReference := &ThesaurusReference{
//   	ThesaurusArn: jsii.String("thesaurusArn"),
//   }
//
type ThesaurusReference struct {
	// The Arn of the Thesaurus resource.
	ThesaurusArn *string `field:"required" json:"thesaurusArn" yaml:"thesaurusArn"`
}

