package awsec2


// A reference to a FlowLog resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowLogReference := &FlowLogReference{
//   	FlowLogId: jsii.String("flowLogId"),
//   }
//
type FlowLogReference struct {
	// The Id of the FlowLog resource.
	FlowLogId *string `field:"required" json:"flowLogId" yaml:"flowLogId"`
}

