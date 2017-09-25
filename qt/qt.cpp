#include <QtGui/qguiapplication.h>
#include <QtQml/qqmlapplicationengine.h>
#include <QtQuick/qquickview.h>
#include "qt.h"

void *NewGuiApplication(int argc, char** argv)
{
    return new QGuiApplication(argc, argv);
}

void DeleteGuiApplication(/*QGuiApplication*/void *ptr)
{
    delete reinterpret_cast<QGuiApplication*>(ptr);
}

int GuiApplicationExec(void* ptr)
{
    auto app = reinterpret_cast<QGuiApplication*>(ptr);
    return app->exec();
}

void *NewQmlApplicationEngine()
{
    return new QQmlApplicationEngine();
}

void DeleteQmlApplicationEngine(void* ptr)
{
    delete reinterpret_cast<QQmlApplicationEngine*>(ptr);
}

void QmlApplicationEngineLoadUrl(void* ptr, char* url)
{
    auto qae = reinterpret_cast<QQmlApplicationEngine*>(ptr);
    qae->load(QUrl(url));
}

void QmlApplicationEngineLoadFile(void* ptr, char* file)
{
    auto qae = reinterpret_cast<QQmlApplicationEngine*>(ptr);
    qae->load(QString(file));
}

void QmlApplicationEngineLoadData(void* ptr, char* data)
{
    auto qae = reinterpret_cast<QQmlApplicationEngine*>(ptr);
    qae->loadData(QByteArray(data));
}

void *NewQuickView(/*QQmlEngine*/void* enginePtr, /*QWindow*/void *parentPtr)
{
    auto engine = reinterpret_cast<QQmlEngine*>(enginePtr);
    auto parent = reinterpret_cast<QWindow*>(parentPtr);
    auto view = new QQuickView(engine, parent);
# if 0
    view->setResizeMode(QQuickView::SizeViewToRootObject);
# else
    view->setResizeMode(QQuickView::SizeRootObjectToView);
# endif
    return view;
}

void DeleteQuickView(void *ptr)
{
    delete reinterpret_cast<QQuickView*>(ptr);
}

void QuickViewSetSource(void *ptr, char *src)
{
    auto view = reinterpret_cast<QQuickView*>(ptr);
    view->setSource(QString(src));
}

void QuickViewShow(void *ptr)
{
    auto view = reinterpret_cast<QQuickView*>(ptr);
    view->show();
}
