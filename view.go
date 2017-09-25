package quick

import (
        "github.com/duzy/quick/qt"
)

type view struct {
        *qt.QuickView
}

func NewView(parent interface{}) View {
        return &view{ qt.NewQuickView(nil) }
}
