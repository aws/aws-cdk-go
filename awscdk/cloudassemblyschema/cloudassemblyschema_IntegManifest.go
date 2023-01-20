package cloudassemblyschema


// Definitions for the integration testing manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integManifest := &integManifest{
//   	testCases: map[string]testCase{
//   		"testCasesKey": &testCase{
//   			"stacks": []*string{
//   				jsii.String("stacks"),
//   			},
//
//   			// the properties below are optional
//   			"allowDestroy": []*string{
//   				jsii.String("allowDestroy"),
//   			},
//   			"assertionStack": jsii.String("assertionStack"),
//   			"assertionStackName": jsii.String("assertionStackName"),
//   			"cdkCommandOptions": &CdkCommands{
//   				"deploy": &DeployCommand{
//   					"args": &DeployOptions{
//   						"all": jsii.Boolean(false),
//   						"app": jsii.String("app"),
//   						"assetMetadata": jsii.Boolean(false),
//   						"caBundlePath": jsii.String("caBundlePath"),
//   						"changeSetName": jsii.String("changeSetName"),
//   						"ci": jsii.Boolean(false),
//   						"color": jsii.Boolean(false),
//   						"context": map[string]*string{
//   							"contextKey": jsii.String("context"),
//   						},
//   						"debug": jsii.Boolean(false),
//   						"ec2Creds": jsii.Boolean(false),
//   						"exclusively": jsii.Boolean(false),
//   						"execute": jsii.Boolean(false),
//   						"force": jsii.Boolean(false),
//   						"ignoreErrors": jsii.Boolean(false),
//   						"json": jsii.Boolean(false),
//   						"lookups": jsii.Boolean(false),
//   						"notices": jsii.Boolean(false),
//   						"notificationArns": []*string{
//   							jsii.String("notificationArns"),
//   						},
//   						"output": jsii.String("output"),
//   						"outputsFile": jsii.String("outputsFile"),
//   						"parameters": map[string]*string{
//   							"parametersKey": jsii.String("parameters"),
//   						},
//   						"pathMetadata": jsii.Boolean(false),
//   						"profile": jsii.String("profile"),
//   						"proxy": jsii.String("proxy"),
//   						"requireApproval": awscdk.cloud_assembly_schema.RequireApproval_NEVER,
//   						"reuseAssets": []*string{
//   							jsii.String("reuseAssets"),
//   						},
//   						"roleArn": jsii.String("roleArn"),
//   						"rollback": jsii.Boolean(false),
//   						"stacks": []*string{
//   							jsii.String("stacks"),
//   						},
//   						"staging": jsii.Boolean(false),
//   						"strict": jsii.Boolean(false),
//   						"toolkitStackName": jsii.String("toolkitStackName"),
//   						"trace": jsii.Boolean(false),
//   						"usePreviousParameters": jsii.Boolean(false),
//   						"verbose": jsii.Boolean(false),
//   						"versionReporting": jsii.Boolean(false),
//   					},
//   					"enabled": jsii.Boolean(false),
//   					"expectedMessage": jsii.String("expectedMessage"),
//   					"expectError": jsii.Boolean(false),
//   				},
//   				"destroy": &DestroyCommand{
//   					"args": &DestroyOptions{
//   						"all": jsii.Boolean(false),
//   						"app": jsii.String("app"),
//   						"assetMetadata": jsii.Boolean(false),
//   						"caBundlePath": jsii.String("caBundlePath"),
//   						"color": jsii.Boolean(false),
//   						"context": map[string]*string{
//   							"contextKey": jsii.String("context"),
//   						},
//   						"debug": jsii.Boolean(false),
//   						"ec2Creds": jsii.Boolean(false),
//   						"exclusively": jsii.Boolean(false),
//   						"force": jsii.Boolean(false),
//   						"ignoreErrors": jsii.Boolean(false),
//   						"json": jsii.Boolean(false),
//   						"lookups": jsii.Boolean(false),
//   						"notices": jsii.Boolean(false),
//   						"output": jsii.String("output"),
//   						"pathMetadata": jsii.Boolean(false),
//   						"profile": jsii.String("profile"),
//   						"proxy": jsii.String("proxy"),
//   						"roleArn": jsii.String("roleArn"),
//   						"stacks": []*string{
//   							jsii.String("stacks"),
//   						},
//   						"staging": jsii.Boolean(false),
//   						"strict": jsii.Boolean(false),
//   						"trace": jsii.Boolean(false),
//   						"verbose": jsii.Boolean(false),
//   						"versionReporting": jsii.Boolean(false),
//   					},
//   					"enabled": jsii.Boolean(false),
//   					"expectedMessage": jsii.String("expectedMessage"),
//   					"expectError": jsii.Boolean(false),
//   				},
//   			},
//   			"diffAssets": jsii.Boolean(false),
//   			"hooks": &Hooks{
//   				"postDeploy": []*string{
//   					jsii.String("postDeploy"),
//   				},
//   				"postDestroy": []*string{
//   					jsii.String("postDestroy"),
//   				},
//   				"preDeploy": []*string{
//   					jsii.String("preDeploy"),
//   				},
//   				"preDestroy": []*string{
//   					jsii.String("preDestroy"),
//   				},
//   			},
//   			"regions": []*string{
//   				jsii.String("regions"),
//   			},
//   			"stackUpdateWorkflow": jsii.Boolean(false),
//   		},
//   	},
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	enableLookups: jsii.Boolean(false),
//   	synthContext: map[string]*string{
//   		"synthContextKey": jsii.String("synthContext"),
//   	},
//   }
//
type IntegManifest struct {
	// test cases.
	TestCases *map[string]*TestCase `field:"required" json:"testCases" yaml:"testCases"`
	// Version of the manifest.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Enable lookups for this test.
	//
	// If lookups are enabled
	// then `stackUpdateWorkflow` must be set to false.
	// Lookups should only be enabled when you are explicitely testing
	// lookups.
	EnableLookups *bool `field:"optional" json:"enableLookups" yaml:"enableLookups"`
	// Additional context to use when performing a synth.
	//
	// Any context provided here will override
	// any default context.
	SynthContext *map[string]*string `field:"optional" json:"synthContext" yaml:"synthContext"`
}

