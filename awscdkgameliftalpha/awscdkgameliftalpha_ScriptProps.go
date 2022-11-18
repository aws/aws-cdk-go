// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for a new realtime server script.
//
// Example:
//   var bucket bucket
//
//   gamelift.NewScript(this, jsii.String("Script"), &scriptProps{
//   	content: gamelift.content.fromBucket(bucket, jsii.String("sample-asset-key")),
//   })
//
// Experimental.
type ScriptProps struct {
	// The game content.
	// Experimental.
	Content Content `field:"required" json:"content" yaml:"content"`
	// The IAM role assumed by GameLift to access server script in S3.
	//
	// If providing a custom role, it needs to trust the GameLift service principal (gamelift.amazonaws.com) and be granted sufficient permissions
	// to have Read access to a specific key content into a specific S3 bucket.
	// Below an example of required permission:
	// {
	//   "Version": "2012-10-17",
	//   "Statement": [{
	//         "Effect": "Allow",
	//         "Action": [
	//             "s3:GetObject",
	//             "s3:GetObjectVersion"
	//         ],
	//         "Resource": "arn:aws:s3:::bucket-name/object-name"
	//   }]
	// }.
	// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/security_iam_id-based-policy-examples.html#security_iam_id-based-policy-examples-access-storage-loc
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Name of this realtime server script.
	// Experimental.
	ScriptName *string `field:"optional" json:"scriptName" yaml:"scriptName"`
	// Version of this realtime server script.
	// Experimental.
	ScriptVersion *string `field:"optional" json:"scriptVersion" yaml:"scriptVersion"`
}

