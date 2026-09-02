package interfacesawscases


// A reference to a Case resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   caseReference := &CaseReference{
//   	CaseArn: jsii.String("caseArn"),
//   }
//
type CaseReference struct {
	// The Arn of the Case resource.
	CaseArn *string `field:"required" json:"caseArn" yaml:"caseArn"`
}

