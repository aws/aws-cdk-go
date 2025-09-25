package awsodb


// A reference to a OdbPeeringConnection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   odbPeeringConnectionReference := &OdbPeeringConnectionReference{
//   	OdbPeeringConnectionArn: jsii.String("odbPeeringConnectionArn"),
//   }
//
type OdbPeeringConnectionReference struct {
	// The OdbPeeringConnectionArn of the OdbPeeringConnection resource.
	OdbPeeringConnectionArn *string `field:"required" json:"odbPeeringConnectionArn" yaml:"odbPeeringConnectionArn"`
}

