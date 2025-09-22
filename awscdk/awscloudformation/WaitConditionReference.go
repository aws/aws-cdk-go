package awscloudformation


// A reference to a WaitCondition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   waitConditionReference := &WaitConditionReference{
//   	WaitConditionId: jsii.String("waitConditionId"),
//   }
//
type WaitConditionReference struct {
	// The Id of the WaitCondition resource.
	WaitConditionId *string `field:"required" json:"waitConditionId" yaml:"waitConditionId"`
}

