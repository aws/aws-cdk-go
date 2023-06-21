package awscodebuild


// The type returned from `IArtifacts#bind`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactsConfig := &ArtifactsConfig{
//   	ArtifactsProperty: &ArtifactsProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		ArtifactIdentifier: jsii.String("artifactIdentifier"),
//   		EncryptionDisabled: jsii.Boolean(false),
//   		Location: jsii.String("location"),
//   		Name: jsii.String("name"),
//   		NamespaceType: jsii.String("namespaceType"),
//   		OverrideArtifactName: jsii.Boolean(false),
//   		Packaging: jsii.String("packaging"),
//   		Path: jsii.String("path"),
//   	},
//   }
//
type ArtifactsConfig struct {
	// The low-level CloudFormation artifacts property.
	ArtifactsProperty *CfnProject_ArtifactsProperty `field:"required" json:"artifactsProperty" yaml:"artifactsProperty"`
}

