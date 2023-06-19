//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

func (s *jsiiProxy_StateGraph) validateRegisterPolicyStatementParameters(statement awsiam.PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StateGraph) validateRegisterStateParameters(state State) error {
	if state == nil {
		return fmt.Errorf("parameter state is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StateGraph) validateRegisterSuperGraphParameters(graph StateGraph) error {
	if graph == nil {
		return fmt.Errorf("parameter graph is required, but nil was provided")
	}

	return nil
}

func validateNewStateGraphParameters(startState State, graphDescription *string) error {
	if startState == nil {
		return fmt.Errorf("parameter startState is required, but nil was provided")
	}

	if graphDescription == nil {
		return fmt.Errorf("parameter graphDescription is required, but nil was provided")
	}

	return nil
}

