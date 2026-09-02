package interfacesawssagemaker


// A reference to a Workforce resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workforceReference := &WorkforceReference{
//   	WorkforceArn: jsii.String("workforceArn"),
//   }
//
type WorkforceReference struct {
	// The WorkforceArn of the Workforce resource.
	WorkforceArn *string `field:"required" json:"workforceArn" yaml:"workforceArn"`
}

