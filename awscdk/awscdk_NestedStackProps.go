// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Initialization props for the `NestedStack` construct.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/constructs-go/constructs"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   /**
//    * This file showcases how to split up a RestApi's Resources and Methods across nested stacks.
//    *
//    * The root stack 'RootStack' first defines a RestApi.
//    * Two nested stacks BooksStack and PetsStack, create corresponding Resources '/books' and '/pets'.
//    * They are then deployed to a 'prod' Stage via a third nested stack - DeployStack.
//    *
//    * To verify this worked, go to the APIGateway
//    */
//
//   type rootStack struct {
//   	stack
//   }
//
//   func newRootStack(scope construct) *rootStack {
//   	this := &rootStack{}
//   	newStack_Override(this, scope, jsii.String("integ-restapi-import-RootStack"))
//
//   	restApi := awscdk.NewRestApi(this, jsii.String("RestApi"), &restApiProps{
//   		cloudWatchRole: jsii.Boolean(true),
//   		deploy: jsii.Boolean(false),
//   	})
//   	restApi.root.addMethod(jsii.String("ANY"))
//
//   	petsStack := NewPetsStack(this, &resourceNestedStackProps{
//   		restApiId: restApi.restApiId,
//   		rootResourceId: restApi.restApiRootResourceId,
//   	})
//   	booksStack := NewBooksStack(this, &resourceNestedStackProps{
//   		restApiId: restApi.restApiId,
//   		rootResourceId: restApi.restApiRootResourceId,
//   	})
//   	NewDeployStack(this, &deployStackProps{
//   		restApiId: restApi.restApiId,
//   		methods: petsStack.methods.concat(booksStack.methods),
//   	})
//
//   	awscdk.NewCfnOutput(this, jsii.String("PetsURL"), &cfnOutputProps{
//   		value: fmt.Sprintf("https://%v.execute-api.%v.amazonaws.com/prod/pets", restApi.restApiId, this.region),
//   	})
//
//   	awscdk.NewCfnOutput(this, jsii.String("BooksURL"), &cfnOutputProps{
//   		value: fmt.Sprintf("https://%v.execute-api.%v.amazonaws.com/prod/books", restApi.restApiId, this.region),
//   	})
//   	return this
//   }
//
//   type resourceNestedStackProps struct {
//   	nestedStackProps
//   	restApiId *string
//   	rootResourceId *string
//   }
//
//   type petsStack struct {
//   	nestedStack
//   	methods []method
//   }
//
//   func newPetsStack(scope construct, props resourceNestedStackProps) *petsStack {
//   	this := &petsStack{}
//   	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-PetsStack"), props)
//
//   	api := awscdk.RestApi.fromRestApiAttributes(this, jsii.String("RestApi"), &restApiAttributes{
//   		restApiId: props.restApiId,
//   		rootResourceId: props.rootResourceId,
//   	})
//
//   	method := api.root.addResource(jsii.String("pets")).addMethod(jsii.String("GET"), awscdk.NewMockIntegration(&integrationOptions{
//   		integrationResponses: []integrationResponse{
//   			&integrationResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   		passthroughBehavior: awscdk.PassthroughBehavior_NEVER,
//   		requestTemplates: map[string]*string{
//   			"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   		},
//   	}), &methodOptions{
//   		methodResponses: []methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   	})
//
//   	this.methods.push(method)
//   	return this
//   }
//
//   type booksStack struct {
//   	nestedStack
//   	methods []*method
//   }
//
//   func newBooksStack(scope construct, props resourceNestedStackProps) *booksStack {
//   	this := &booksStack{}
//   	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-BooksStack"), props)
//
//   	api := awscdk.RestApi.fromRestApiAttributes(this, jsii.String("RestApi"), &restApiAttributes{
//   		restApiId: props.restApiId,
//   		rootResourceId: props.rootResourceId,
//   	})
//
//   	method := api.root.addResource(jsii.String("books")).addMethod(jsii.String("GET"), awscdk.NewMockIntegration(&integrationOptions{
//   		integrationResponses: []*integrationResponse{
//   			&integrationResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   		passthroughBehavior: awscdk.PassthroughBehavior_NEVER,
//   		requestTemplates: map[string]*string{
//   			"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   		},
//   	}), &methodOptions{
//   		methodResponses: []*methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   	})
//
//   	this.methods.push(method)
//   	return this
//   }
//
//   type deployStackProps struct {
//   	nestedStackProps
//   	restApiId *string
//   	methods []*method
//   }
//
//   type deployStack struct {
//   	nestedStack
//   }
//
//   func newDeployStack(scope construct, props deployStackProps) *deployStack {
//   	this := &deployStack{}
//   	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-DeployStack"), props)
//
//   	deployment := awscdk.NewDeployment(this, jsii.String("Deployment"), &deploymentProps{
//   		api: awscdk.RestApi.fromRestApiId(this, jsii.String("RestApi"), props.restApiId),
//   	})
//   	if *props.methods {
//   		for _, method := range *props.methods {
//   			deployment.node.addDependency(method)
//   		}
//   	}
//   	awscdk.NewStage(this, jsii.String("Stage"), &stageProps{
//   		deployment: deployment,
//   	})
//   	return this
//   }
//
//   NewRootStack(awscdk.NewApp())
//
type NestedStackProps struct {
	// A description of the stack.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Simple Notification Service (SNS) topics to publish stack related events.
	NotificationArns *[]*string `field:"optional" json:"notificationArns" yaml:"notificationArns"`
	// The set value pairs that represent the parameters passed to CloudFormation when this nested stack is created.
	//
	// Each parameter has a name corresponding
	// to a parameter defined in the embedded template and a value representing
	// the value that you want to set for the parameter.
	//
	// The nested stack construct will automatically synthesize parameters in order
	// to bind references from the parent stack(s) into the nested stack.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Policy to apply when the nested stack is removed.
	//
	// The default is `Destroy`, because all Removal Policies of resources inside the
	// Nested Stack should already have been set correctly. You normally should
	// not need to set this value.
	RemovalPolicy RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The length of time that CloudFormation waits for the nested stack to reach the CREATE_COMPLETE state.
	//
	// When CloudFormation detects that the nested stack has reached the
	// CREATE_COMPLETE state, it marks the nested stack resource as
	// CREATE_COMPLETE in the parent stack and resumes creating the parent stack.
	// If the timeout period expires before the nested stack reaches
	// CREATE_COMPLETE, CloudFormation marks the nested stack as failed and rolls
	// back both the nested stack and parent stack.
	Timeout Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

