package awscodebuild


// Environment type for Windows Docker images.
//
// Example:
//   var ecrRepository repository
//
//
//   codebuild.NewProject(this, jsii.String("Project"), &projectProps{
//   	environment: &buildEnvironment{
//   		buildImage: codebuild.windowsBuildImage.fromEcrRepository(ecrRepository, jsii.String("v1.0"), codebuild.windowsImageType_SERVER_2019),
//   		// optional certificate to include in the build image
//   		certificate: &buildEnvironmentCertificate{
//   			bucket: s3.bucket.fromBucketName(this, jsii.String("Bucket"), jsii.String("my-bucket")),
//   			objectKey: jsii.String("path/to/cert.pem"),
//   		},
//   	},
//   })
//
type WindowsImageType string

const (
	// The standard environment type, WINDOWS_CONTAINER.
	WindowsImageType_STANDARD WindowsImageType = "STANDARD"
	// The WINDOWS_SERVER_2019_CONTAINER environment type.
	WindowsImageType_SERVER_2019 WindowsImageType = "SERVER_2019"
)

