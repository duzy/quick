package quick

import (
        "github.com/duzy/quick/qt"
)

type application struct {
        *qt.GuiApplication
        *qt.QmlApplicationEngine
}

func NewApplication(args ...string) Application {
        return &application{ 
                qt.NewGuiApplication(args),
                qt.NewQmlApplicationEngine(),
        }
}
