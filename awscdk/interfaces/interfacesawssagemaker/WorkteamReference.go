package interfacesawssagemaker


// A reference to a Workteam resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workteamReference := &WorkteamReference{
//   	WorkteamId: jsii.String("workteamId"),
//   	WorkteamName: jsii.String("workteamName"),
//   }
//
type WorkteamReference struct {
	// The Id of the Workteam resource.
	WorkteamId *string `field:"required" json:"workteamId" yaml:"workteamId"`
	// The WorkteamName of the Workteam resource.
	WorkteamName *string `field:"required" json:"workteamName" yaml:"workteamName"`
}

