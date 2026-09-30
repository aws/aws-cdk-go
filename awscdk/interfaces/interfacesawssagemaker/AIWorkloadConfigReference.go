package interfacesawssagemaker


// A reference to a AIWorkloadConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIWorkloadConfigReference := &AIWorkloadConfigReference{
//   	AiWorkloadConfigArn: jsii.String("aiWorkloadConfigArn"),
//   }
//
type AIWorkloadConfigReference struct {
	// The AIWorkloadConfigArn of the AIWorkloadConfig resource.
	AiWorkloadConfigArn *string `field:"required" json:"aiWorkloadConfigArn" yaml:"aiWorkloadConfigArn"`
}

