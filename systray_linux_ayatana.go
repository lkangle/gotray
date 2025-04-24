// +build linux,!legacy_appindicator
//go:build linux && !legacy_appindicator
// +build linux,!legacy_appindicator

package gotray

/*
#cgo linux pkg-config: ayatana-appindicator3-0.1

#include "systray.h"
*/
import "C"
