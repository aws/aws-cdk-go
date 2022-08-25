package pipelines


// Properties for a `StackDeployment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stackDeploymentProps := &stackDeploymentProps{
//   	absoluteTemplatePath: jsii.String("absoluteTemplatePath"),
//   	constructPath: jsii.String("constructPath"),
//   	stackArtifactId: jsii.String("stackArtifactId"),
//   	stackName: jsii.String("stackName"),
//
//   	// the properties below are optional
//   	account: jsii.String("account"),
//   	assets: []stackAsset{
//   		&stackAsset{
//   			assetId: jsii.String("assetId"),
//   			assetManifestPath: jsii.String("assetManifestPath"),
//   			assetSelector: jsii.String("assetSelector"),
//   			assetType: awscdk.Pipelines.assetType_FILE,
//   			isTemplate: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			assetPublishingRoleArn: jsii.String("assetPublishingRoleArn"),
//   		},
//   	},
//   	assumeRoleArn: jsii.String("assumeRoleArn"),
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	region: jsii.String("region"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	templateS3Uri: jsii.String("templateS3Uri"),
//   }
//
type StackDeploymentProps struct {
	// Template path on disk to cloud assembly (cdk.out).
	AbsoluteTemplatePath *string `field:"required" json:"absoluteTemplatePath" yaml:"absoluteTemplatePath"`
	// Construct path for this stack.
	ConstructPath *string `field:"required" json:"constructPath" yaml:"constructPath"`
	// Artifact ID for this stack.
	StackArtifactId *string `field:"required" json:"stackArtifactId" yaml:"stackArtifactId"`
	// Name for this stack.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// Account where the stack should be deployed.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Assets referenced by this stack.
	Assets *[]*StackAsset `field:"optional" json:"assets" yaml:"assets"`
	// Role to assume before deploying this stack.
	AssumeRoleArn *string `field:"optional" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// Execution role to pass to CloudFormation.
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Region where the stack should be deployed.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Tags to apply to the stack.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The S3 URL which points to the template asset location in the publishing bucket.
	TemplateS3Uri *string `field:"optional" json:"templateS3Uri" yaml:"templateS3Uri"`
}

