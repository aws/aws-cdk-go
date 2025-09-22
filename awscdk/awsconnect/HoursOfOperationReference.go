package awsconnect


// A reference to a HoursOfOperation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hoursOfOperationReference := &HoursOfOperationReference{
//   	HoursOfOperationArn: jsii.String("hoursOfOperationArn"),
//   }
//
type HoursOfOperationReference struct {
	// The HoursOfOperationArn of the HoursOfOperation resource.
	HoursOfOperationArn *string `field:"required" json:"hoursOfOperationArn" yaml:"hoursOfOperationArn"`
}

