package awsiam


// A reference to a Group resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupReference := &GroupReference{
//   	GroupArn: jsii.String("groupArn"),
//   	GroupName: jsii.String("groupName"),
//   }
//
type GroupReference struct {
	// The ARN of the Group resource.
	GroupArn *string `field:"required" json:"groupArn" yaml:"groupArn"`
	// The GroupName of the Group resource.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
}

