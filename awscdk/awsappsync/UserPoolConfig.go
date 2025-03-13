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
//   userPoolConfig := &UserPoolConfig{
//   	UserPool: userPool,
//
//   	// the properties below are optional
//   	AppIdClientRegex: jsii.String("appIdClientRegex"),
//   	DefaultAction: awscdk.Aws_appsync.UserPoolDefaultAction_ALLOW,
//   }
//
type UserPoolConfig struct {
	// The Cognito user pool to use as identity source.
	UserPool awscognito.IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// the optional app id regex.
	// Default: -  None.
	//
	AppIdClientRegex *string `field:"optional" json:"appIdClientRegex" yaml:"appIdClientRegex"`
	// Default auth action.
	// Default: ALLOW.
	//
	DefaultAction UserPoolDefaultAction `field:"optional" json:"defaultAction" yaml:"defaultAction"`
}

