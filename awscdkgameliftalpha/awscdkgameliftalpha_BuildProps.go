// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for a new build.
//
// Example:
//   var bucket bucket
//
//   gamelift.NewBuild(this, jsii.String("Build"), &buildProps{
//   	content: gamelift.content.fromBucket(bucket, jsii.String("sample-asset-key")),
//   })
//
// Experimental.
type BuildProps struct {
	// The game build file storage.
	// Experimental.
	Content Content `field:"required" json:"content" yaml:"content"`
	// Name of this build.
	// Experimental.
	BuildName *string `field:"optional" json:"buildName" yaml:"buildName"`
	// Version of this build.
	// Experimental.
	BuildVersion *string `field:"optional" json:"buildVersion" yaml:"buildVersion"`
	// The operating system that the game server binaries are built to run on.
	// Experimental.
	OperatingSystem OperatingSystem `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// The IAM role assumed by GameLift to access server build in S3.
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
}

