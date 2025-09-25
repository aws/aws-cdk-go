package awsmacie


// A reference to a AllowList resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allowListReference := &AllowListReference{
//   	AllowListArn: jsii.String("allowListArn"),
//   	AllowListId: jsii.String("allowListId"),
//   }
//
type AllowListReference struct {
	// The ARN of the AllowList resource.
	AllowListArn *string `field:"required" json:"allowListArn" yaml:"allowListArn"`
	// The Id of the AllowList resource.
	AllowListId *string `field:"required" json:"allowListId" yaml:"allowListId"`
}

