package interfacesawslookoutmetrics


// A reference to a Alert resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alertReference := &AlertReference{
//   	AlertArn: jsii.String("alertArn"),
//   }
//
type AlertReference struct {
	// The Arn of the Alert resource.
	AlertArn *string `field:"required" json:"alertArn" yaml:"alertArn"`
}

