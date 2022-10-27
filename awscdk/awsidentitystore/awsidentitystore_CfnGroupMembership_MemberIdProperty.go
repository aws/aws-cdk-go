package awsidentitystore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberIdProperty := &memberIdProperty{
//   	userId: jsii.String("userId"),
//   }
//
type CfnGroupMembership_MemberIdProperty struct {
	// `CfnGroupMembership.MemberIdProperty.UserId`.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

