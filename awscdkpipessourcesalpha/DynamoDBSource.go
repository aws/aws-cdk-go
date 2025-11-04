package awscdkpipessourcesalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipessourcesalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

// A source that reads from an DynamoDB stream.
//
// Example:
//   var targetQueue Queue
//   table := ddb.NewTableV2(this, jsii.String("MyTable"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: ddb.AttributeType_STRING,
//   	},
//   	DynamoStream: ddb.StreamViewType_NEW_IMAGE,
//   })
//
//   pipeSource := sources.NewDynamoDBSource(table, &DynamoDBSourceParameters{
//   	StartingPosition: sources.DynamoDBStartingPosition_LATEST,
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: pipeSource,
//   	Target: awscdkpipestargetsalpha.NewSqsTarget(targetQueue),
//   })
//
// Experimental.
type DynamoDBSource interface {
	StreamSource
	// The dead-letter SQS queue or SNS topic.
	// Experimental.
	DeadLetterTarget() interface{}
	// The ARN of the source resource.
	// Experimental.
	SourceArn() *string
	// Base parameters for streaming sources.
	// Experimental.
	SourceParameters() *StreamSourceParameters
	// Bind the source to a pipe.
	// Experimental.
	Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.SourceConfig
	// Retrieves the ARN from the dead-letter SQS queue or SNS topic.
	// Experimental.
	GetDeadLetterTargetArn(deadLetterTarget interface{}) *string
	// Grants the pipe role permission to publish to the dead-letter target.
	// Experimental.
	GrantPush(grantee awsiam.IRole, deadLetterTarget interface{})
	// Grant the pipe role read access to the source.
	// Experimental.
	GrantRead(grantee awsiam.IRole)
}

// The jsii proxy struct for DynamoDBSource
type jsiiProxy_DynamoDBSource struct {
	jsiiProxy_StreamSource
}

func (j *jsiiProxy_DynamoDBSource) DeadLetterTarget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDBSource) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDBSource) SourceParameters() *StreamSourceParameters {
	var returns *StreamSourceParameters
	_jsii_.Get(
		j,
		"sourceParameters",
		&returns,
	)
	return returns
}


// Experimental.
func NewDynamoDBSource(table awsdynamodb.ITableV2, parameters *DynamoDBSourceParameters) DynamoDBSource {
	_init_.Initialize()

	if err := validateNewDynamoDBSourceParameters(table, parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamoDBSource{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-sources-alpha.DynamoDBSource",
		[]interface{}{table, parameters},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoDBSource_Override(d DynamoDBSource, table awsdynamodb.ITableV2, parameters *DynamoDBSourceParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-sources-alpha.DynamoDBSource",
		[]interface{}{table, parameters},
		d,
	)
}

// Determines if the source is an instance of SourceWithDeadLetterTarget.
// Experimental.
func DynamoDBSource_IsSourceWithDeadLetterTarget(source awscdkpipesalpha.ISource) *bool {
	_init_.Initialize()

	if err := validateDynamoDBSource_IsSourceWithDeadLetterTargetParameters(source); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-pipes-sources-alpha.DynamoDBSource",
		"isSourceWithDeadLetterTarget",
		[]interface{}{source},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDBSource) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.SourceConfig {
	if err := d.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *awscdkpipesalpha.SourceConfig

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDBSource) GetDeadLetterTargetArn(deadLetterTarget interface{}) *string {
	if err := d.validateGetDeadLetterTargetArnParameters(deadLetterTarget); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getDeadLetterTargetArn",
		[]interface{}{deadLetterTarget},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDBSource) GrantPush(grantee awsiam.IRole, deadLetterTarget interface{}) {
	if err := d.validateGrantPushParameters(grantee, deadLetterTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"grantPush",
		[]interface{}{grantee, deadLetterTarget},
	)
}

func (d *jsiiProxy_DynamoDBSource) GrantRead(grantee awsiam.IRole) {
	if err := d.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"grantRead",
		[]interface{}{grantee},
	)
}

