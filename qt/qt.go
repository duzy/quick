package qt

//#include "qt.h"
import "C"
import (
        "runtime"
        "unsafe"
        "strings"
        "fmt"
)

type cstrs struct {
        v []*C.char
}

func (cs *cstrs) append_nil() {
        cs.v = append(cs.v, nil)
}

func (cs *cstrs) append(s string) {
        cs.v = append(cs.v, C.CString(s))
}

func (cs *cstrs) free() {
        for _, p := range cs.v {
                C.free(unsafe.Pointer(p))
        }
}

func (cs *cstrs) ptr() **C.char {
        return &cs.v[0]
}

func cstrings(a []string) *cstrs {
        var v []*C.char
        for _, s := range a {
                v = append(v, C.CString(s))
        }
        return &cstrs{v}
}

type GuiApplication struct {
        void unsafe.Pointer        
}

func (app *GuiApplication) Exec() int {
        return int(C.GuiApplicationExec(app.void))
}

func (app *GuiApplication) delete() {
        C.DeleteGuiApplication(app.void); app.void = nil
}

func NewGuiApplication(args []string) (app *GuiApplication) {
        argv := cstrings(args); defer argv.free()
        argv.append_nil() // null-terminated
        app = &GuiApplication{
                void: C.NewGuiApplication(C.int(len(args)), argv.ptr()),
        }
        return
}

type QmlApplicationEngine struct {
        void unsafe.Pointer
}

func (qae *QmlApplicationEngine) Load(source interface{}) error {
        s := fmt.Sprintf("%v", source)
        cs := C.CString(s); defer C.free(unsafe.Pointer(cs))
        // FIXME: better classification method.
        if strings.Contains(s, ":/") {
                C.QmlApplicationEngineLoadUrl(qae.void, cs);
        } else if !strings.Contains(s, "\n") {
                C.QmlApplicationEngineLoadFile(qae.void, cs);
        } else {
                C.QmlApplicationEngineLoadData(qae.void, cs);
        }
        return nil
}

func (qae *QmlApplicationEngine) delete() {
        C.DeleteQmlApplicationEngine(qae.void); qae.void = nil
}

func NewQmlApplicationEngine() (qae *QmlApplicationEngine) {
        qae = &QmlApplicationEngine{
                void: C.NewQmlApplicationEngine(),
        }
        return
}

type QuickView struct {
        void unsafe.Pointer
}

func (qv *QuickView) SetSource(source interface{}) error {
        s := fmt.Sprintf("%v", source)
        if s == "" {
                return nil
        } else if strings.HasPrefix(s, ":/") {
                s = strings.Replace(s, ":/", "qrc:///", 1)
        }
        cs := C.CString(s); defer C.free(unsafe.Pointer(cs))
        C.QuickViewSetSource(qv.void, cs)
        return nil
}

func (qv *QuickView) Show() {
        C.QuickViewShow(qv.void)
}

func (qv *QuickView) delete() {
        C.DeleteQuickView(qv.void); qv.void = nil
}

func NewQuickView(parent *QuickView) *QuickView {
        var parptr unsafe.Pointer
        if parent != nil {
                parptr = parent.void
        }
        void := C.NewQuickView(nil, parptr)
        view := &QuickView{ void }
        runtime.SetFinalizer(view, (*QuickView).delete)
        return view
}
