package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for a new build.
//
// Example:
//   var bucket bucket
//
//   build := gamelift.NewBuild(this, jsii.String("Build"), &BuildProps{
//   	Content: gamelift.Content_FromBucket(bucket, jsii.String("sample-asset-key")),
//   })
//
//   awscdk.NewCfnOutput(this, jsii.String("BuildArn"), &CfnOutputProps{
//   	Value: build.BuildArn,
//   })
//   awscdk.NewCfnOutput(this, jsii.String("BuildId"), &CfnOutputProps{
//   	Value: build.BuildId,
//   })
//
// Experimental.
type BuildProps struct {
	// The game build file storage.
	// Experimental.
	Content Content `field:"required" json:"content" yaml:"content"`
	// Name of this build.
	// Default: No name.
	//
	// Experimental.
	BuildName *string `field:"optional" json:"buildName" yaml:"buildName"`
	// Version of this build.
	// Default: No version.
	//
	// Experimental.
	BuildVersion *string `field:"optional" json:"buildVersion" yaml:"buildVersion"`
	// The operating system that the game server binaries are built to run on.
	// Default: No version.
	//
	// Experimental.
	OperatingSystem OperatingSystem `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// The IAM role assumed by GameLift to access server build in S3.
	//
	// If providing a custom role, it needs to trust the GameLift service principal (gamelift.amazonaws.com) and be granted sufficient permissions
	// to have Read access to a specific key content into a specific S3 bucket.
	// Below an example of required permission:
	// {
	//  "Version": "2012-10-17",
	//  "Statement": [{
	//        "Effect": "Allow",
	//        "Action": [
	//            "s3:GetObject",
	//            "s3:GetObjectVersion"
	//        ],
	//        "Resource": "arn:aws:s3:::bucket-name/object-name"
	//  }]
	// }.
	// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/security_iam_id-based-policy-examples.html#security_iam_id-based-policy-examples-access-storage-loc
	//
	// Default: - a role will be created with default permissions.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// A server SDK version you used when integrating your game server build with Amazon GameLift.
	// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-custom-intro.html
	//
	// Default: - 4.0.2
	//
	// Experimental.
	ServerSdkVersion *string `field:"optional" json:"serverSdkVersion" yaml:"serverSdkVersion"`
}

