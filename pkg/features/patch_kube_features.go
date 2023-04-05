package features

import (
	"os"
)

func OpenShiftStartCSIMigrationVSphere() bool {
	_, migrationEnabled := os.LookupEnv("OPENSHIFT_DO_VSPHERE_MIGRATION")
	return migrationEnabled
}
