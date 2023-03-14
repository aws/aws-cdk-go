package awscloudfront


// An origin in an origin group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originGroupMemberProperty := &originGroupMemberProperty{
//   	originId: jsii.String("originId"),
//   }
//
type CfnDistribution_OriginGroupMemberProperty struct {
	// The ID for an origin in an origin group.
	OriginId *string `field:"required" json:"originId" yaml:"originId"`
}

