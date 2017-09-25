#ifndef GOQUICK_QT_H
#define GOQUICK_QT_H 1
#include <stdlib.h>
#ifdef __cplusplus
extern "C" {
#endif//__cplusplus
    
    void *NewGuiApplication(int, char**);
    void DeleteGuiApplication(void*);
    int GuiApplicationExec(void*);

    void *NewQmlApplicationEngine();
    void DeleteQmlApplicationEngine(void*);
    void QmlApplicationEngineLoadUrl(void*, char*);
    void QmlApplicationEngineLoadFile(void*, char*);
    void QmlApplicationEngineLoadData(void*, char*);

    void *NewQuickView(void*, void*);
    void DeleteQuickView(void*);
    void QuickViewSetSource(void *, char*);
    void QuickViewShow(void *);
    
#ifdef __cplusplus
}
#endif//__cplusplus
#endif//GOQUICK_QT_H
