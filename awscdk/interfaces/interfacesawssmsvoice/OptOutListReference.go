package interfacesawssmsvoice


// A reference to a OptOutList resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   optOutListReference := &OptOutListReference{
//   	OptOutListArn: jsii.String("optOutListArn"),
//   	OptOutListName: jsii.String("optOutListName"),
//   }
//
type OptOutListReference struct {
	// The ARN of the OptOutList resource.
	OptOutListArn *string `field:"required" json:"optOutListArn" yaml:"optOutListArn"`
	// The OptOutListName of the OptOutList resource.
	OptOutListName *string `field:"required" json:"optOutListName" yaml:"optOutListName"`
}

