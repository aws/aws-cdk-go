package interfacesawseks


// A reference to a Nodegroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodegroupReference := &NodegroupReference{
//   	NodegroupArn: jsii.String("nodegroupArn"),
//   	NodegroupId: jsii.String("nodegroupId"),
//   }
//
type NodegroupReference struct {
	// The ARN of the Nodegroup resource.
	NodegroupArn *string `field:"required" json:"nodegroupArn" yaml:"nodegroupArn"`
	// The Id of the Nodegroup resource.
	NodegroupId *string `field:"required" json:"nodegroupId" yaml:"nodegroupId"`
}

