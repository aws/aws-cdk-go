package interfacesawscustomerprofiles


// A reference to a Recommender resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recommenderReference := &RecommenderReference{
//   	DomainName: jsii.String("domainName"),
//   	RecommenderArn: jsii.String("recommenderArn"),
//   	RecommenderName: jsii.String("recommenderName"),
//   }
//
type RecommenderReference struct {
	// The DomainName of the Recommender resource.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The ARN of the Recommender resource.
	RecommenderArn *string `field:"required" json:"recommenderArn" yaml:"recommenderArn"`
	// The RecommenderName of the Recommender resource.
	RecommenderName *string `field:"required" json:"recommenderName" yaml:"recommenderName"`
}

