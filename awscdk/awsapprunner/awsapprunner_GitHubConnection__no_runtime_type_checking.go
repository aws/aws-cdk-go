//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapprunner

// Building without runtime type checking enabled, so all the below just return nil

func validateGitHubConnection_FromConnectionArnParameters(arn *string) error {
	return nil
}

func validateNewGitHubConnectionParameters(arn *string) error {
	return nil
}

