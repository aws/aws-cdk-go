package awshealthlake


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   createdAtProperty := &createdAtProperty{
//   	nanos: jsii.Number(123),
//   	seconds: jsii.String("seconds"),
//   }
//
type CfnFHIRDatastore_CreatedAtProperty struct {
	// `CfnFHIRDatastore.CreatedAtProperty.Nanos`.
	Nanos *float64 `field:"required" json:"nanos" yaml:"nanos"`
	// `CfnFHIRDatastore.CreatedAtProperty.Seconds`.
	Seconds *string `field:"required" json:"seconds" yaml:"seconds"`
}

