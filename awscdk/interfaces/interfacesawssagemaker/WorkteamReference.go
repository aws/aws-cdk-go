package interfacesawssagemaker


// A reference to a Workteam resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workteamReference := &WorkteamReference{
//   	WorkteamArn: jsii.String("workteamArn"),
//   }
//
type WorkteamReference struct {
	// The WorkteamArn of the Workteam resource.
	WorkteamArn *string `field:"required" json:"workteamArn" yaml:"workteamArn"`
}

