package interfacesawspersonalize


// A reference to a Recommender resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recommenderReference := &RecommenderReference{
//   	RecommenderArn: jsii.String("recommenderArn"),
//   }
//
type RecommenderReference struct {
	// The RecommenderArn of the Recommender resource.
	RecommenderArn *string `field:"required" json:"recommenderArn" yaml:"recommenderArn"`
}

