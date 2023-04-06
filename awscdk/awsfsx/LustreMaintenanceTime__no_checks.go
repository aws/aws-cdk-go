//go:build no_runtime_type_checking

package awsfsx

// Building without runtime type checking enabled, so all the below just return nil

func validateNewLustreMaintenanceTimeParameters(props *LustreMaintenanceTimeProps) error {
	return nil
}

