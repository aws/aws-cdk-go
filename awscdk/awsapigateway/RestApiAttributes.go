package awsapigateway


// Attributes that can be specified when importing a RestApi.
//
// Example:
//   import "github.com/aws/constructs-go/constructs"
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   	restApi := awscdk.NewRestApi(this, jsii.String("RestApi"), &RestApiProps{
//   		CloudWatchRole: jsii.Boolean(true),
//   		Deploy: jsii.Boolean(false),
//   	})
//   	restApi.Root.AddMethod(jsii.String("ANY"))
//
//   	petsStack := NewPetsStack(this, &resourceNestedStackProps{
//   		restApiId: restApi.RestApiId,
//   		rootResourceId: restApi.RestApiRootResourceId,
//   	})
//   	booksStack := NewBooksStack(this, &resourceNestedStackProps{
//   		restApiId: restApi.*RestApiId,
//   		rootResourceId: restApi.*RestApiRootResourceId,
//   	})
//   	NewDeployStack(this, &deployStackProps{
//   		restApiId: restApi.*RestApiId,
//   		methods: petsStack.methods.concat(booksStack.methods),
//   	})
//
//   	awscdk.NewCfnOutput(this, jsii.String("PetsURL"), &CfnOutputProps{
//   		Value: fmt.Sprintf("https://%v.execute-api.%v.amazonaws.com/prod/pets", restApi.*RestApiId, this.Region),
//   	})
//
//   	awscdk.NewCfnOutput(this, jsii.String("BooksURL"), &CfnOutputProps{
//   		Value: fmt.Sprintf("https://%v.execute-api.%v.amazonaws.com/prod/books", restApi.*RestApiId, this.*Region),
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
//   	api := awscdk.RestApi_FromRestApiAttributes(this, jsii.String("RestApi"), &RestApiAttributes{
//   		RestApiId: props.restApiId,
//   		RootResourceId: props.rootResourceId,
//   	})
//
//   	method := api.Root.AddResource(jsii.String("pets")).AddMethod(jsii.String("GET"), awscdk.NewMockIntegration(&IntegrationOptions{
//   		IntegrationResponses: []integrationResponse{
//   			&integrationResponse{
//   				StatusCode: jsii.String("200"),
//   			},
//   		},
//   		PassthroughBehavior: awscdk.PassthroughBehavior_NEVER,
//   		RequestTemplates: map[string]*string{
//   			"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   		},
//   	}), &MethodOptions{
//   		MethodResponses: []methodResponse{
//   			&methodResponse{
//   				StatusCode: jsii.String("200"),
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
//   	api := awscdk.RestApi_FromRestApiAttributes(this, jsii.String("RestApi"), &RestApiAttributes{
//   		RestApiId: props.restApiId,
//   		RootResourceId: props.rootResourceId,
//   	})
//
//   	method := api.Root.AddResource(jsii.String("books")).AddMethod(jsii.String("GET"), awscdk.NewMockIntegration(&IntegrationOptions{
//   		IntegrationResponses: []*integrationResponse{
//   			&integrationResponse{
//   				StatusCode: jsii.String("200"),
//   			},
//   		},
//   		PassthroughBehavior: awscdk.PassthroughBehavior_NEVER,
//   		RequestTemplates: map[string]*string{
//   			"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   		},
//   	}), &MethodOptions{
//   		MethodResponses: []*methodResponse{
//   			&methodResponse{
//   				StatusCode: jsii.String("200"),
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
//   	deployment := awscdk.NewDeployment(this, jsii.String("Deployment"), &DeploymentProps{
//   		Api: awscdk.RestApi_FromRestApiId(this, jsii.String("RestApi"), props.restApiId),
//   	})
//   	if *props.methods {
//   		for _, method := range *props.methods {
//   			deployment.Node.AddDependency(method)
//   		}
//   	}
//   	awscdk.NewStage(this, jsii.String("Stage"), &StageProps{
//   		Deployment: Deployment,
//   	})
//   	return this
//   }
//
//   NewRootStack(awscdk.NewApp())
//
type RestApiAttributes struct {
	// The ID of the API Gateway RestApi.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The resource ID of the root resource.
	RootResourceId *string `field:"required" json:"rootResourceId" yaml:"rootResourceId"`
	// The name of the API Gateway RestApi.
	// Default: - ID of the RestApi construct.
	//
	RestApiName *string `field:"optional" json:"restApiName" yaml:"restApiName"`
}

