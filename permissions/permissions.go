package permissions

import "os"

type FilePermission os.FileMode

// FilePermissions are defined in a way where
// 1. User (U), Group (G), Others (O) are always first
// 2. Followed by allowed permissions Read (R), Write (W), Execute (E)
// e.g. Permissions for User to Read and Group to Execute granting nothing to Others will be URGE = 0410
const (
	URWGROR FilePermission = 0644
)
