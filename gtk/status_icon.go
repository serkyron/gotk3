// Same copyright and license as the rest of the files in this project
// This file contains style related functions and structures

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

/*
 * GtkStatusIcon
 */

// StatusIcon is a representation of GTK's GtkStatusIcon.
type StatusIcon struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkStatusIcon.
func (v *StatusIcon) native() *C.GtkStatusIcon {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStatusIcon(p)
}

func wrapStatusIcon(obj *glib.Object) *StatusIcon {
	return &StatusIcon{obj}
}

// StatusIconNew is a wrapper around gtk_status_icon_new().
func StatusIconNew() (*StatusIcon, error) {
	c := C.gtk_status_icon_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapStatusIcon(obj), nil
}

// StatusIconNewFromFile is a wrapper around gtk_status_icon_new_from_file().
func StatusIconNewFromFile(file string) (*StatusIcon, error) {
	cstr := C.CString(file)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_status_icon_new_from_file((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapStatusIcon(obj), nil
}

// Wrapper around gtk_status_icon_set_visible
func (v *StatusIcon) SetVisible(visible bool) {
	C.gtk_status_icon_set_visible(v.GObject, gbool(visible))
}
