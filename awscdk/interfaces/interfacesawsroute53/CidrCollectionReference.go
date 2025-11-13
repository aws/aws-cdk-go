package interfacesawsroute53


// A reference to a CidrCollection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cidrCollectionReference := &CidrCollectionReference{
//   	CidrCollectionArn: jsii.String("cidrCollectionArn"),
//   	CidrCollectionId: jsii.String("cidrCollectionId"),
//   }
//
type CidrCollectionReference struct {
	// The ARN of the CidrCollection resource.
	CidrCollectionArn *string `field:"required" json:"cidrCollectionArn" yaml:"cidrCollectionArn"`
	// The Id of the CidrCollection resource.
	CidrCollectionId *string `field:"required" json:"cidrCollectionId" yaml:"cidrCollectionId"`
}

