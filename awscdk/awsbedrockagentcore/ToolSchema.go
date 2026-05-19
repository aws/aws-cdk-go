package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// ****************************************************************************                      TOOL SCHEMA CLASS ***************************************************************************.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   lambdaFunction := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_22_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String(`
//   	        exports.handler = async (event) => {
//   	            return {
//   	                statusCode: 200,
//   	                body: JSON.stringify({ message: 'Hello from Lambda!' })
//   	            };
//   	        };
//   	    `)),
//   })
//
//   // Create a gateway target with Lambda and tool schema
//   target := agentcore.GatewayTarget_ForLambda(this, jsii.String("MyLambdaTarget"), &GatewayTargetLambdaProps{
//   	GatewayTargetName: jsii.String("my-lambda-target"),
//   	Description: jsii.String("Target for Lambda function integration"),
//   	Gateway: gateway,
//   	LambdaFunction: lambdaFunction,
//   	ToolSchema: agentcore.ToolSchema_FromLocalAsset(path.join(__dirname, jsii.String("schemas"), jsii.String("my-tool-schema.json"))),
//   })
//
type ToolSchema interface {
	// The account ID of the S3 bucket owner for cross-account access.
	BucketOwnerAccountId() *string
	// The inline tool schema definition as a string, if using an inline schema.
	//
	// Can be in JSON or YAML format.
	InlineSchema() *[]*ToolDefinition
	// The S3 location of the tool schema file, if using an S3-based schema.
	//
	// Contains the bucket name and object key information.
	S3File() *awss3.Location
	// Bind the schema to a construct.
	Bind(scope constructs.Construct)
	// Grant permissions to the role.
	GrantPermissionsToRole(role awsiam.IRole)
}

// The jsii proxy struct for ToolSchema
type jsiiProxy_ToolSchema struct {
	_ byte // padding
}

func (j *jsiiProxy_ToolSchema) BucketOwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ToolSchema) InlineSchema() *[]*ToolDefinition {
	var returns *[]*ToolDefinition
	_jsii_.Get(
		j,
		"inlineSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ToolSchema) S3File() *awss3.Location {
	var returns *awss3.Location
	_jsii_.Get(
		j,
		"s3File",
		&returns,
	)
	return returns
}


func NewToolSchema_Override(t ToolSchema, s3File *awss3.Location, bucketOwnerAccountId *string, inlineSchema *[]*ToolDefinition) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.ToolSchema",
		[]interface{}{s3File, bucketOwnerAccountId, inlineSchema},
		t,
	)
}

// Creates a Tool Schema from an inline string.
func ToolSchema_FromInline(schema *[]*ToolDefinition) InlineToolSchema {
	_init_.Initialize()

	if err := validateToolSchema_FromInlineParameters(schema); err != nil {
		panic(err)
	}
	var returns InlineToolSchema

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.ToolSchema",
		"fromInline",
		[]interface{}{schema},
		&returns,
	)

	return returns
}

// Creates a tool Schema from a local file.
func ToolSchema_FromLocalAsset(path *string) ToolSchema {
	_init_.Initialize()

	if err := validateToolSchema_FromLocalAssetParameters(path); err != nil {
		panic(err)
	}
	var returns ToolSchema

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.ToolSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates a Tool Schema from an S3 File.
func ToolSchema_FromS3File(bucket awss3.IBucket, objectKey *string, bucketOwnerAccountId *string) S3ToolSchema {
	_init_.Initialize()

	if err := validateToolSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ToolSchema

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.ToolSchema",
		"fromS3File",
		[]interface{}{bucket, objectKey, bucketOwnerAccountId},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_ToolSchema) Bind(scope constructs.Construct) {
	if err := t.validateBindParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"bind",
		[]interface{}{scope},
	)
}

func (t *jsiiProxy_ToolSchema) GrantPermissionsToRole(role awsiam.IRole) {
	if err := t.validateGrantPermissionsToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"grantPermissionsToRole",
		[]interface{}{role},
	)
}

