package awsiotthingsgraph


// A reference to a FlowTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowTemplateReference := &FlowTemplateReference{
//   	FlowTemplateId: jsii.String("flowTemplateId"),
//   }
//
type FlowTemplateReference struct {
	// The Id of the FlowTemplate resource.
	FlowTemplateId *string `field:"required" json:"flowTemplateId" yaml:"flowTemplateId"`
}

