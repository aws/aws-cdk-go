package awsworkspacesweb


// A reference to a NetworkSettings resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkSettingsReference := &NetworkSettingsReference{
//   	NetworkSettingsArn: jsii.String("networkSettingsArn"),
//   }
//
type NetworkSettingsReference struct {
	// The NetworkSettingsArn of the NetworkSettings resource.
	NetworkSettingsArn *string `field:"required" json:"networkSettingsArn" yaml:"networkSettingsArn"`
}

