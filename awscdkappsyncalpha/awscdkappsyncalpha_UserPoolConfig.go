// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
)

// Configuration for Cognito user-pools in AppSync.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appsync_alpha "github.com/aws/aws-cdk-go/awscdkappsyncalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var userPool userPool
//
//   userPoolConfig := &userPoolConfig{
//   	userPool: userPool,
//
//   	// the properties below are optional
//   	appIdClientRegex: jsii.String("appIdClientRegex"),
//   	defaultAction: appsync_alpha.userPoolDefaultAction_ALLOW,
//   }
//
// Experimental.
type UserPoolConfig struct {
	// The Cognito user pool to use as identity source.
	// Experimental.
	UserPool awscognito.IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// the optional app id regex.
	// Experimental.
	AppIdClientRegex *string `field:"optional" json:"appIdClientRegex" yaml:"appIdClientRegex"`
	// Default auth action.
	// Experimental.
	DefaultAction UserPoolDefaultAction `field:"optional" json:"defaultAction" yaml:"defaultAction"`
}

