package interfacesawskendra


// A reference to a FeaturedResultsSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   featuredResultsSetReference := &FeaturedResultsSetReference{
//   	FeaturedResultsSetArn: jsii.String("featuredResultsSetArn"),
//   }
//
type FeaturedResultsSetReference struct {
	// The Arn of the FeaturedResultsSet resource.
	FeaturedResultsSetArn *string `field:"required" json:"featuredResultsSetArn" yaml:"featuredResultsSetArn"`
}

