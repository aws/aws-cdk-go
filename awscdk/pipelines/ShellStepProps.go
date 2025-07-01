package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for a `ShellStep`.
//
// Example:
//   var pipeline codePipeline
//
//   preprod := NewMyApplicationStage(this, jsii.String("PreProd"))
//   prod := NewMyApplicationStage(this, jsii.String("Prod"))
//   topic := sns.NewTopic(this, jsii.String("ChangeApprovalTopic"))
//
//   pipeline.AddStage(preprod, &AddStageOpts{
//   	Post: []step{
//   		pipelines.NewShellStep(jsii.String("Validate Endpoint"), &ShellStepProps{
//   			Commands: []*string{
//   				jsii.String("curl -Ssf https://my.webservice.com/"),
//   			},
//   		}),
//   	},
//   })
//   pipeline.AddStage(prod, &AddStageOpts{
//   	Pre: []*step{
//   		pipelines.NewManualApprovalStep(jsii.String("PromoteToProd"), &ManualApprovalStepProps{
//   			//All options below are optional
//   			Comment: jsii.String("Please validate changes"),
//   			ReviewUrl: jsii.String("https://my.webservice.com/"),
//   			NotificationTopic: topic,
//   		}),
//   	},
//   })
//
type ShellStepProps struct {
	// Commands to run.
	Commands *[]*string `field:"required" json:"commands" yaml:"commands"`
	// Additional FileSets to put in other directories.
	//
	// Specifies a mapping from directory name to FileSets. During the
	// script execution, the FileSets will be available in the directories
	// indicated.
	//
	// The directory names may be relative. For example, you can put
	// the main input and an additional input side-by-side with the
	// following configuration:
	//
	// ```ts
	// const script = new pipelines.ShellStep('MainScript', {
	//   commands: ['npm ci','npm run build','npx cdk synth'],
	//   input: pipelines.CodePipelineSource.gitHub('org/source1', 'main'),
	//   additionalInputs: {
	//     '../siblingdir': pipelines.CodePipelineSource.gitHub('org/source2', 'main'),
	//   }
	// });
	// ```.
	// Default: - No additional inputs.
	//
	AdditionalInputs *map[string]IFileSetProducer `field:"optional" json:"additionalInputs" yaml:"additionalInputs"`
	// Environment variables to set.
	// Default: - No environment variables.
	//
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// Set environment variables based on Stack Outputs.
	//
	// `ShellStep`s following stack or stage deployments may
	// access the `CfnOutput`s of those stacks to get access to
	// --for example--automatically generated resource names or
	// endpoint URLs.
	// Default: - No environment variables created from stack outputs.
	//
	EnvFromCfnOutputs *map[string]awscdk.CfnOutput `field:"optional" json:"envFromCfnOutputs" yaml:"envFromCfnOutputs"`
	// FileSet to run these scripts on.
	//
	// The files in the FileSet will be placed in the working directory when
	// the script is executed. Use `additionalInputs` to download file sets
	// to other directories as well.
	// Default: - No input specified.
	//
	Input IFileSetProducer `field:"optional" json:"input" yaml:"input"`
	// Installation commands to run before the regular commands.
	//
	// For deployment engines that support it, install commands will be classified
	// differently in the job history from the regular `commands`.
	// Default: - No installation commands.
	//
	InstallCommands *[]*string `field:"optional" json:"installCommands" yaml:"installCommands"`
	// The directory that will contain the primary output fileset.
	//
	// After running the script, the contents of the given directory
	// will be treated as the primary output of this Step.
	// Default: - No primary output.
	//
	PrimaryOutputDirectory *string `field:"optional" json:"primaryOutputDirectory" yaml:"primaryOutputDirectory"`
}

