package interfacesawsmedialive


// A reference to a Offering resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   offeringReference := &OfferingReference{
//   	OfferingArn: jsii.String("offeringArn"),
//   }
//
type OfferingReference struct {
	// The Arn of the Offering resource.
	OfferingArn *string `field:"required" json:"offeringArn" yaml:"offeringArn"`
}

