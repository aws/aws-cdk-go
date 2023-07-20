package awshealthlake


// The time that a Data Store was created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   createdAtProperty := &CreatedAtProperty{
//   	Nanos: jsii.Number(123),
//   	Seconds: jsii.String("seconds"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-createdat.html
//
type CfnFHIRDatastore_CreatedAtProperty struct {
	// Nanoseconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-createdat.html#cfn-healthlake-fhirdatastore-createdat-nanos
	//
	Nanos *float64 `field:"required" json:"nanos" yaml:"nanos"`
	// Seconds since epoch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-createdat.html#cfn-healthlake-fhirdatastore-createdat-seconds
	//
	Seconds *string `field:"required" json:"seconds" yaml:"seconds"`
}

