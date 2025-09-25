package awsmacie


// A reference to a FindingsFilter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   findingsFilterReference := &FindingsFilterReference{
//   	FindingsFilterArn: jsii.String("findingsFilterArn"),
//   	FindingsFilterId: jsii.String("findingsFilterId"),
//   }
//
type FindingsFilterReference struct {
	// The ARN of the FindingsFilter resource.
	FindingsFilterArn *string `field:"required" json:"findingsFilterArn" yaml:"findingsFilterArn"`
	// The Id of the FindingsFilter resource.
	FindingsFilterId *string `field:"required" json:"findingsFilterId" yaml:"findingsFilterId"`
}

