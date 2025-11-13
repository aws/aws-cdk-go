package interfacesawsappstream


// A reference to a AppBlock resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appBlockReference := &AppBlockReference{
//   	AppBlockArn: jsii.String("appBlockArn"),
//   }
//
type AppBlockReference struct {
	// The Arn of the AppBlock resource.
	AppBlockArn *string `field:"required" json:"appBlockArn" yaml:"appBlockArn"`
}

