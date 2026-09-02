package interfacesawskendra


// A reference to a QuerySuggestionsBlockList resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   querySuggestionsBlockListReference := &QuerySuggestionsBlockListReference{
//   	QuerySuggestionsBlockListArn: jsii.String("querySuggestionsBlockListArn"),
//   }
//
type QuerySuggestionsBlockListReference struct {
	// The Arn of the QuerySuggestionsBlockList resource.
	QuerySuggestionsBlockListArn *string `field:"required" json:"querySuggestionsBlockListArn" yaml:"querySuggestionsBlockListArn"`
}

