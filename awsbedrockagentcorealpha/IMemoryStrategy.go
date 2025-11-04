package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrockagentcore"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Interface for Memory strategies.
// Experimental.
type IMemoryStrategy interface {
	// Grants the necessary permissions to the role.
	//
	// Returns: The Grant object for chaining.
	// Experimental.
	Grant(grantee awsiam.IGrantable) awsiam.Grant
	// Renders internal attributes to CloudFormation.
	// Experimental.
	Render() *awsbedrockagentcore.CfnMemory_MemoryStrategyProperty
	// The description of the memory strategy.
	// Experimental.
	Description() *string
	// The name of the memory strategy.
	// Experimental.
	Name() *string
	// The type of memory strategy.
	// Experimental.
	StrategyType() MemoryStrategyType
}

// The jsii proxy for IMemoryStrategy
type jsiiProxy_IMemoryStrategy struct {
	_ byte // padding
}

func (i *jsiiProxy_IMemoryStrategy) Grant(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemoryStrategy) Render() *awsbedrockagentcore.CfnMemory_MemoryStrategyProperty {
	var returns *awsbedrockagentcore.CfnMemory_MemoryStrategyProperty

	_jsii_.Invoke(
		i,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IMemoryStrategy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemoryStrategy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemoryStrategy) StrategyType() MemoryStrategyType {
	var returns MemoryStrategyType
	_jsii_.Get(
		j,
		"strategyType",
		&returns,
	)
	return returns
}

