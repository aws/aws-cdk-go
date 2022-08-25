package awscodebuild


// The type returned from {@link IArtifacts#bind}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactsConfig := &artifactsConfig{
//   	artifactsProperty: &artifactsProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		artifactIdentifier: jsii.String("artifactIdentifier"),
//   		encryptionDisabled: jsii.Boolean(false),
//   		location: jsii.String("location"),
//   		name: jsii.String("name"),
//   		namespaceType: jsii.String("namespaceType"),
//   		overrideArtifactName: jsii.Boolean(false),
//   		packaging: jsii.String("packaging"),
//   		path: jsii.String("path"),
//   	},
//   }
//
type ArtifactsConfig struct {
	// The low-level CloudFormation artifacts property.
	ArtifactsProperty *CfnProject_ArtifactsProperty `field:"required" json:"artifactsProperty" yaml:"artifactsProperty"`
}

