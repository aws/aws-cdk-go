package awsqbusiness


// A reference to a Index resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   indexReference := &IndexReference{
//   	ApplicationId: jsii.String("applicationId"),
//   	IndexArn: jsii.String("indexArn"),
//   	IndexId: jsii.String("indexId"),
//   }
//
type IndexReference struct {
	// The ApplicationId of the Index resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The ARN of the Index resource.
	IndexArn *string `field:"required" json:"indexArn" yaml:"indexArn"`
	// The IndexId of the Index resource.
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
}

