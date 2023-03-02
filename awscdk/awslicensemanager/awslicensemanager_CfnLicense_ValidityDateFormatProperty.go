package awslicensemanager


// Date and time range during which the license is valid, in ISO8601-UTC format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   validityDateFormatProperty := &validityDateFormatProperty{
//   	begin: jsii.String("begin"),
//   	end: jsii.String("end"),
//   }
//
type CfnLicense_ValidityDateFormatProperty struct {
	// Start of the time range.
	Begin *string `field:"required" json:"begin" yaml:"begin"`
	// End of the time range.
	End *string `field:"required" json:"end" yaml:"end"`
}

