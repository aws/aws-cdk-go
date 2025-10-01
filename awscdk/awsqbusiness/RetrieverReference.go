package awsqbusiness


// A reference to a Retriever resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retrieverReference := &RetrieverReference{
//   	ApplicationId: jsii.String("applicationId"),
//   	RetrieverArn: jsii.String("retrieverArn"),
//   	RetrieverId: jsii.String("retrieverId"),
//   }
//
type RetrieverReference struct {
	// The ApplicationId of the Retriever resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The ARN of the Retriever resource.
	RetrieverArn *string `field:"required" json:"retrieverArn" yaml:"retrieverArn"`
	// The RetrieverId of the Retriever resource.
	RetrieverId *string `field:"required" json:"retrieverId" yaml:"retrieverId"`
}

