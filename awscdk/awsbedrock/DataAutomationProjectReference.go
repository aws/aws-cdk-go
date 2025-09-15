package awsbedrock


// A reference to a DataAutomationProject resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataAutomationProjectReference := &DataAutomationProjectReference{
//   	ProjectArn: jsii.String("projectArn"),
//   }
//
type DataAutomationProjectReference struct {
	// The ProjectArn of the DataAutomationProject resource.
	ProjectArn *string `field:"required" json:"projectArn" yaml:"projectArn"`
}

