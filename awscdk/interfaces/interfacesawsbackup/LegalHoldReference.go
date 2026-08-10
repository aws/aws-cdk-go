package interfacesawsbackup


// A reference to a LegalHold resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   legalHoldReference := &LegalHoldReference{
//   	LegalHoldArn: jsii.String("legalHoldArn"),
//   }
//
type LegalHoldReference struct {
	// The Arn of the LegalHold resource.
	LegalHoldArn *string `field:"required" json:"legalHoldArn" yaml:"legalHoldArn"`
}

