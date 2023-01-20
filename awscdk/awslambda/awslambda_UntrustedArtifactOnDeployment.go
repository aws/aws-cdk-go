package awslambda


// Code signing configuration policy for deployment validation failure.
type UntrustedArtifactOnDeployment string

const (
	// Lambda blocks the deployment request if signature validation checks fail.
	UntrustedArtifactOnDeployment_ENFORCE UntrustedArtifactOnDeployment = "ENFORCE"
	// Lambda allows the deployment of the code package, but issues a warning.
	//
	// Lambda issues a new Amazon CloudWatch metric, called a signature validation error and also stores the warning in CloudTrail.
	UntrustedArtifactOnDeployment_WARN UntrustedArtifactOnDeployment = "WARN"
)

