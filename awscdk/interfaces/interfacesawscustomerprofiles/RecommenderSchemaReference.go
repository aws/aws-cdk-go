package interfacesawscustomerprofiles


// A reference to a RecommenderSchema resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recommenderSchemaReference := &RecommenderSchemaReference{
//   	RecommenderSchemaArn: jsii.String("recommenderSchemaArn"),
//   }
//
type RecommenderSchemaReference struct {
	// The Arn of the RecommenderSchema resource.
	RecommenderSchemaArn *string `field:"required" json:"recommenderSchemaArn" yaml:"recommenderSchemaArn"`
}

