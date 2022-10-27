// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Example:
//   type stackUnderTestProps struct {
//   	stackProps
//   	architecture architecture
//   }
//
//   type stackUnderTest struct {
//   	stack
//   }
//
//   func newStackUnderTest(scope construct, id *string, props stackUnderTestProps) *stackUnderTest {
//   	this := &stackUnderTest{}
//   	newStack_Override(this, scope, id, props)
//
//   	lambda.NewFunction(this, jsii.String("Handler"), &functionProps{
//   		runtime: lambda.runtime_NODEJS_14_X(),
//   		handler: jsii.String("index.handler"),
//   		code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   		architecture: props.architecture,
//   	})
//   	return this
//   }
//
//   // Beginning of the test suite
//   app := awscdk.NewApp()
//
//   awscdkintegtestsalpha.NewIntegTest(app, jsii.String("DifferentArchitectures"), &integTestProps{
//   	testCases: []*stack{
//   		NewStackUnderTest(app, jsii.String("Stack1"), &stackUnderTestProps{
//   			architecture: lambda.*architecture_ARM_64(),
//   		}),
//   		NewStackUnderTest(app, jsii.String("Stack2"), &stackUnderTestProps{
//   			architecture: lambda.*architecture_X86_64(),
//   		}),
//   	},
//   })
//
type StackProps struct {
	// Include runtime versioning information in this Stack.
	AnalyticsReporting *bool `field:"optional" json:"analyticsReporting" yaml:"analyticsReporting"`
	// A description of the stack.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The AWS environment (account/region) where this stack will be deployed.
	//
	// Set the `region`/`account` fields of `env` to either a concrete value to
	// select the indicated environment (recommended for production stacks), or to
	// the values of environment variables
	// `CDK_DEFAULT_REGION`/`CDK_DEFAULT_ACCOUNT` to let the target environment
	// depend on the AWS credentials/configuration that the CDK CLI is executed
	// under (recommended for development stacks).
	//
	// If the `Stack` is instantiated inside a `Stage`, any undefined
	// `region`/`account` fields from `env` will default to the same field on the
	// encompassing `Stage`, if configured there.
	//
	// If either `region` or `account` are not set nor inherited from `Stage`, the
	// Stack will be considered "*environment-agnostic*"". Environment-agnostic
	// stacks can be deployed to any environment but may not be able to take
	// advantage of all features of the CDK. For example, they will not be able to
	// use environmental context lookups such as `ec2.Vpc.fromLookup` and will not
	// automatically translate Service Principals to the right format based on the
	// environment's AWS partition, and other such enhancements.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   // Use a concrete account and region to deploy this stack to:
	//   // `.account` and `.region` will simply return these values.
	//   // Use a concrete account and region to deploy this stack to:
	//   // `.account` and `.region` will simply return these values.
	//   awscdk.Newstack(app, jsii.String("Stack1"), &stackProps{
	//   	env: &environment{
	//   		account: jsii.String("123456789012"),
	//   		region: jsii.String("us-east-1"),
	//   	},
	//   })
	//
	//   // Use the CLI's current credentials to determine the target environment:
	//   // `.account` and `.region` will reflect the account+region the CLI
	//   // is configured to use (based on the user CLI credentials)
	//   // Use the CLI's current credentials to determine the target environment:
	//   // `.account` and `.region` will reflect the account+region the CLI
	//   // is configured to use (based on the user CLI credentials)
	//   awscdk.Newstack(app, jsii.String("Stack2"), &stackProps{
	//   	env: &environment{
	//   		account: process.env.cDK_DEFAULT_ACCOUNT,
	//   		region: process.env.cDK_DEFAULT_REGION,
	//   	},
	//   })
	//
	//   // Define multiple stacks stage associated with an environment
	//   myStage := awscdk.NewStage(app, jsii.String("MyStage"), &stageProps{
	//   	env: &environment{
	//   		account: jsii.String("123456789012"),
	//   		region: jsii.String("us-east-1"),
	//   	},
	//   })
	//
	//   // both of these stacks will use the stage's account/region:
	//   // `.account` and `.region` will resolve to the concrete values as above
	//   // both of these stacks will use the stage's account/region:
	//   // `.account` and `.region` will resolve to the concrete values as above
	//   NewMyStack(myStage, jsii.String("Stack1"))
	//   NewYourStack(myStage, jsii.String("Stack2"))
	//
	//   // Define an environment-agnostic stack:
	//   // `.account` and `.region` will resolve to `{ "Ref": "AWS::AccountId" }` and `{ "Ref": "AWS::Region" }` respectively.
	//   // which will only resolve to actual values by CloudFormation during deployment.
	//   // Define an environment-agnostic stack:
	//   // `.account` and `.region` will resolve to `{ "Ref": "AWS::AccountId" }` and `{ "Ref": "AWS::Region" }` respectively.
	//   // which will only resolve to actual values by CloudFormation during deployment.
	//   NewMyStack(app, jsii.String("Stack1"))
	//
	Env *Environment `field:"optional" json:"env" yaml:"env"`
	// Name to deploy the stack with.
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
	// Synthesis method to use while deploying this stack.
	Synthesizer IStackSynthesizer `field:"optional" json:"synthesizer" yaml:"synthesizer"`
	// Stack tags that will be applied to all the taggable resources and the stack itself.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Whether to enable termination protection for this stack.
	TerminationProtection *bool `field:"optional" json:"terminationProtection" yaml:"terminationProtection"`
}

