package interfacesawsfms


// A reference to a ApplicationsList resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationsListReference := &ApplicationsListReference{
//   	ApplicationsListArn: jsii.String("applicationsListArn"),
//   }
//
type ApplicationsListReference struct {
	// The Arn of the ApplicationsList resource.
	ApplicationsListArn *string `field:"required" json:"applicationsListArn" yaml:"applicationsListArn"`
}

