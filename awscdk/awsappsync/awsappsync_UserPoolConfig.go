package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
)

// Configuration for Cognito user-pools in AppSync.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var userPool userPool
//
//   userPoolConfig := &userPoolConfig{
//   	userPool: userPool,
//
//   	// the properties below are optional
//   	appIdClientRegex: jsii.String("appIdClientRegex"),
//   	defaultAction: awscdk.Aws_appsync.userPoolDefaultAction_ALLOW,
//   }
//
type UserPoolConfig struct {
	// The Cognito user pool to use as identity source.
	UserPool awscognito.IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// the optional app id regex.
	AppIdClientRegex *string `field:"optional" json:"appIdClientRegex" yaml:"appIdClientRegex"`
	// Default auth action.
	DefaultAction UserPoolDefaultAction `field:"optional" json:"defaultAction" yaml:"defaultAction"`
}

