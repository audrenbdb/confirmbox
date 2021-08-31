// +build linux

package confirmbox


// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include <stdlib.h>
// static GtkWidget* msgdlg(GtkWindow *parent, GtkDialogFlags flags, GtkMessageType type, GtkButtonsType buttons, char *msg) {
// 	return gtk_message_dialog_new(parent, flags, type, buttons, "%s", msg);
// }
// static GtkWidget* filedlg(char *title, GtkWindow *parent, GtkFileChooserAction action, char* acceptText) {
// 	return gtk_file_chooser_dialog_new(title, parent, action, "Cancel", GTK_RESPONSE_CANCEL, acceptText, GTK_RESPONSE_ACCEPT, NULL);
// }
import "C"
import (
	"unsafe"
)

func init() {
	C.gtk_init(nil, nil)
}

func closeDialog(dlg *C.GtkWidget) {
	C.gtk_widget_destroy(dlg)
	for C.gtk_events_pending() != 0 {
		C.gtk_main_iteration()
	}
}

func isConfirmed(title, content string) bool {
	cContent := C.CString(content)
	defer C.free(unsafe.Pointer(cContent))
	dlg := C.msgdlg(nil, 0, C.GTK_MESSAGE_QUESTION, C.GTK_BUTTONS_YES_NO, cContent)
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.gtk_window_set_title((*C.GtkWindow)(unsafe.Pointer(dlg)), cTitle)
	defer closeDialog(dlg)
	return C.gtk_dialog_run((*C.GtkDialog)(unsafe.Pointer(dlg))) == C.GTK_RESPONSE_YES
}
