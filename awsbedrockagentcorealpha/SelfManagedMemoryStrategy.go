package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrockagentcore"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Use AgentCore memory for event storage with custom triggers.
//
// Define memory processing logic in your own environment.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("memoryBucket"), &BucketProps{
//   	BucketName: jsii.String("test-memory"),
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	AutoDeleteObjects: jsii.Boolean(true),
//   })
//
//   topic := sns.NewTopic(this, jsii.String("topic"))
//
//   // Create a custom semantic memory strategy
//   selfManagedStrategy := agentcore.MemoryStrategy_UsingSelfManaged(&SelfManagedStrategyProps{
//   	Name: jsii.String("selfManagedStrategy"),
//   	Description: jsii.String("self managed memory strategy"),
//   	HistoricalContextWindowSize: jsii.Number(5),
//   	InvocationConfiguration: &InvocationConfiguration{
//   		Topic: topic,
//   		S3Location: &Location{
//   			BucketName: bucket.BucketName,
//   			ObjectKey: jsii.String("memory/"),
//   		},
//   	},
//   	TriggerConditions: &TriggerConditions{
//   		MessageBasedTrigger: jsii.Number(1),
//   		TimeBasedTrigger: cdk.Duration_Seconds(jsii.Number(10)),
//   		TokenBasedTrigger: jsii.Number(100),
//   	},
//   })
//
//   // Create memory with custom strategy
//   memory := agentcore.NewMemory(this, jsii.String("MyMemory"), &MemoryProps{
//   	MemoryName: jsii.String("my-custom-memory"),
//   	Description: jsii.String("Memory with custom strategy"),
//   	ExpirationDuration: cdk.Duration_Days(jsii.Number(90)),
//   	MemoryStrategies: []IMemoryStrategy{
//   		selfManagedStrategy,
//   	},
//   })
//
// Experimental.
type SelfManagedMemoryStrategy interface {
	IMemoryStrategy
	// The description of the memory strategy.
	// Experimental.
	Description() *string
	// Historical context window size for self managed memory strategy.
	// Experimental.
	HistoricalContextWindowSize() *float64
	// Invocation configuration for self managed memory strategy.
	// Experimental.
	InvocationConfiguration() *InvocationConfiguration
	// The name of the memory strategy.
	// Experimental.
	Name() *string
	// The type of memory strategy.
	// Experimental.
	StrategyType() MemoryStrategyType
	// Trigger conditions for self managed memory strategy.
	// Experimental.
	TriggerConditions() *TriggerConditions
	// Grants the necessary permissions to the role.
	//
	// [disable-awslint:no-grants].
	//
	// Returns: The Grant object for chaining.
	// Experimental.
	Grant(grantee awsiam.IGrantable) awsiam.Grant
	// Renders internal attributes to CloudFormation.
	// Experimental.
	Render() *awsbedrockagentcore.CfnMemory_MemoryStrategyProperty
}

// The jsii proxy struct for SelfManagedMemoryStrategy
type jsiiProxy_SelfManagedMemoryStrategy struct {
	jsiiProxy_IMemoryStrategy
}

func (j *jsiiProxy_SelfManagedMemoryStrategy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfManagedMemoryStrategy) HistoricalContextWindowSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"historicalContextWindowSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfManagedMemoryStrategy) InvocationConfiguration() *InvocationConfiguration {
	var returns *InvocationConfiguration
	_jsii_.Get(
		j,
		"invocationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfManagedMemoryStrategy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfManagedMemoryStrategy) StrategyType() MemoryStrategyType {
	var returns MemoryStrategyType
	_jsii_.Get(
		j,
		"strategyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfManagedMemoryStrategy) TriggerConditions() *TriggerConditions {
	var returns *TriggerConditions
	_jsii_.Get(
		j,
		"triggerConditions",
		&returns,
	)
	return returns
}


// Experimental.
func NewSelfManagedMemoryStrategy(strategyType MemoryStrategyType, props *SelfManagedStrategyProps) SelfManagedMemoryStrategy {
	_init_.Initialize()

	if err := validateNewSelfManagedMemoryStrategyParameters(strategyType, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SelfManagedMemoryStrategy{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.SelfManagedMemoryStrategy",
		[]interface{}{strategyType, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSelfManagedMemoryStrategy_Override(s SelfManagedMemoryStrategy, strategyType MemoryStrategyType, props *SelfManagedStrategyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.SelfManagedMemoryStrategy",
		[]interface{}{strategyType, props},
		s,
	)
}

func (s *jsiiProxy_SelfManagedMemoryStrategy) Grant(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grant",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SelfManagedMemoryStrategy) Render() *awsbedrockagentcore.CfnMemory_MemoryStrategyProperty {
	var returns *awsbedrockagentcore.CfnMemory_MemoryStrategyProperty

	_jsii_.Invoke(
		s,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

