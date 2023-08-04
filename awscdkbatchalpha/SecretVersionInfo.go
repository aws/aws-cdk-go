package awscdkbatchalpha


// Specify the secret's version id or version stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import batch_alpha "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//
//   secretVersionInfo := &SecretVersionInfo{
//   	VersionId: jsii.String("versionId"),
//   	VersionStage: jsii.String("versionStage"),
//   }
//
// Experimental.
type SecretVersionInfo struct {
	// version id of the secret.
	// Default: - use default version id.
	//
	// Experimental.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
	// version stage of the secret.
	// Default: - use default version stage.
	//
	// Experimental.
	VersionStage *string `field:"optional" json:"versionStage" yaml:"versionStage"`
}

