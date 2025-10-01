package awscloudformation


// A reference to a WaitConditionHandle resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   waitConditionHandleReference := &WaitConditionHandleReference{
//   	WaitConditionHandleId: jsii.String("waitConditionHandleId"),
//   }
//
type WaitConditionHandleReference struct {
	// The Id of the WaitConditionHandle resource.
	WaitConditionHandleId *string `field:"required" json:"waitConditionHandleId" yaml:"waitConditionHandleId"`
}

