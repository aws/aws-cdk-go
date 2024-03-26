package awscdkintegtestsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for defining an AssertionsProvider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assertionsProviderProps := &AssertionsProviderProps{
//   	Handler: jsii.String("handler"),
//   	LogRetention: awscdk.Aws_logs.RetentionDays_ONE_DAY,
//   	Uuid: jsii.String("uuid"),
//   }
//
// Experimental.
type AssertionsProviderProps struct {
	// The handler to use for the lambda function.
	// Default: index.handler
	//
	// Experimental.
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// How long, in days, the log contents will be retained.
	// Default: - no retention days specified.
	//
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// This determines the uniqueness of each AssertionsProvider.
	//
	// You should only need to provide something different here if you
	// _know_ that you need a separate provider.
	// Default: - the default uuid is used.
	//
	// Experimental.
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

