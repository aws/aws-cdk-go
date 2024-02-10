package awss3


// An interface that represents the location of a specific object in an S3 Bucket.
//
// Example:
//   lambdaStack := cdk.NewStack(app, jsii.String("LambdaStack"))
//   lambdaCode := lambda.Code_FromCfnParameters()
//   lambda.NewFunction(lambdaStack, jsii.String("Lambda"), &FunctionProps{
//   	Code: lambdaCode,
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   })
//   // other resources that your Lambda needs, added to the lambdaStack...
//
//   pipelineStack := cdk.NewStack(app, jsii.String("PipelineStack"))
//   pipeline := codepipeline.NewPipeline(pipelineStack, jsii.String("Pipeline"), &PipelineProps{
//   	CrossAccountKeys: jsii.Boolean(true),
//   })
//
//   // add the source code repository containing this code to your Pipeline,
//   // and the source code of the Lambda Function, if they're separate
//   cdkSourceOutput := codepipeline.NewArtifact()
//   cdkSourceAction := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
//   	Repository: codecommit.NewRepository(pipelineStack, jsii.String("CdkCodeRepo"), &RepositoryProps{
//   		RepositoryName: jsii.String("CdkCodeRepo"),
//   	}),
//   	ActionName: jsii.String("CdkCode_Source"),
//   	Output: cdkSourceOutput,
//   })
//   lambdaSourceOutput := codepipeline.NewArtifact()
//   lambdaSourceAction := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
//   	Repository: codecommit.NewRepository(pipelineStack, jsii.String("LambdaCodeRepo"), &RepositoryProps{
//   		RepositoryName: jsii.String("LambdaCodeRepo"),
//   	}),
//   	ActionName: jsii.String("LambdaCode_Source"),
//   	Output: lambdaSourceOutput,
//   })
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Source"),
//   	Actions: []iAction{
//   		cdkSourceAction,
//   		lambdaSourceAction,
//   	},
//   })
//
//   // synthesize the Lambda CDK template, using CodeBuild
//   // the below values are just examples, assuming your CDK code is in TypeScript/JavaScript -
//   // adjust the build environment and/or commands accordingly
//   cdkBuildProject := codebuild.NewProject(pipelineStack, jsii.String("CdkBuildProject"), &ProjectProps{
//   	Environment: &BuildEnvironment{
//   		BuildImage: codebuild.LinuxBuildImage_STANDARD_7_0(),
//   	},
//   	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string]*string{
//   			"install": map[string]*string{
//   				"commands": jsii.String("npm install"),
//   			},
//   			"build": map[string][]*string{
//   				"commands": []*string{
//   					jsii.String("npm run build"),
//   					jsii.String("npm run cdk synth LambdaStack -- -o ."),
//   				},
//   			},
//   		},
//   		"artifacts": map[string]*string{
//   			"files": jsii.String("LambdaStack.template.yaml"),
//   		},
//   	}),
//   })
//   cdkBuildOutput := codepipeline.NewArtifact()
//   cdkBuildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   	ActionName: jsii.String("CDK_Build"),
//   	Project: cdkBuildProject,
//   	Input: cdkSourceOutput,
//   	Outputs: []artifact{
//   		cdkBuildOutput,
//   	},
//   })
//
//   // build your Lambda code, using CodeBuild
//   // again, this example assumes your Lambda is written in TypeScript/JavaScript -
//   // make sure to adjust the build environment and/or commands if they don't match your specific situation
//   lambdaBuildProject := codebuild.NewProject(pipelineStack, jsii.String("LambdaBuildProject"), &ProjectProps{
//   	Environment: &BuildEnvironment{
//   		BuildImage: codebuild.LinuxBuildImage_STANDARD_7_0(),
//   	},
//   	BuildSpec: codebuild.BuildSpec_*FromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string]*string{
//   			"install": map[string]*string{
//   				"commands": jsii.String("npm install"),
//   			},
//   			"build": map[string]*string{
//   				"commands": jsii.String("npm run build"),
//   			},
//   		},
//   		"artifacts": map[string][]*string{
//   			"files": []*string{
//   				jsii.String("index.js"),
//   				jsii.String("node_modules/**/*"),
//   			},
//   		},
//   	}),
//   })
//   lambdaBuildOutput := codepipeline.NewArtifact()
//   lambdaBuildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   	ActionName: jsii.String("Lambda_Build"),
//   	Project: lambdaBuildProject,
//   	Input: lambdaSourceOutput,
//   	Outputs: []*artifact{
//   		lambdaBuildOutput,
//   	},
//   })
//
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Build"),
//   	Actions: []*iAction{
//   		cdkBuildAction,
//   		lambdaBuildAction,
//   	},
//   })
//
//   // finally, deploy your Lambda Stack
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Deploy"),
//   	Actions: []*iAction{
//   		codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&CloudFormationCreateUpdateStackActionProps{
//   			ActionName: jsii.String("Lambda_CFN_Deploy"),
//   			TemplatePath: cdkBuildOutput.AtPath(jsii.String("LambdaStack.template.yaml")),
//   			StackName: jsii.String("LambdaStackDeployedName"),
//   			AdminPermissions: jsii.Boolean(true),
//   			ParameterOverrides: lambdaCode.Assign(lambdaBuildOutput.s3Location),
//   			ExtraInputs: []*artifact{
//   				lambdaBuildOutput,
//   			},
//   		}),
//   	},
//   })
//
type Location struct {
	// The name of the S3 Bucket the object is in.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The path inside the Bucket where the object is located at.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
	// The S3 object version.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

