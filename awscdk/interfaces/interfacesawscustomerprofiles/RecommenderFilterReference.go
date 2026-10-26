package interfacesawscustomerprofiles


// A reference to a RecommenderFilter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recommenderFilterReference := &RecommenderFilterReference{
//   	RecommenderFilterArn: jsii.String("recommenderFilterArn"),
//   }
//
type RecommenderFilterReference struct {
	// The RecommenderFilterArn of the RecommenderFilter resource.
	RecommenderFilterArn *string `field:"required" json:"recommenderFilterArn" yaml:"recommenderFilterArn"`
}

