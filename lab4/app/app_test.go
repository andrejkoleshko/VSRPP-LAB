package app_test

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "lab4/app"
    "lab4/app/mocks"

    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/require"
)

func TestHandleIndex(t *testing.T) {
    ctrl := gomock.NewController(t)

    db := mocks.NewMockDB(ctrl)
    rows := mocks.NewMockRows(ctrl)

    // Ожидания для Query
    db.EXPECT().
        Query("SELECT id, model, company, price FROM products").
        Return(rows, nil)

    rows.EXPECT().Next().Return(false)
    rows.EXPECT().Close().Return(nil)

    logger := mocks.NewMockLogger(ctrl)
    logger.EXPECT().Info(gomock.Any()).AnyTimes()
    logger.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()

    a := app.New(db, logger)

    req := httptest.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()

    a.HandleIndexForTest(w, req)

    resp := w.Result()
    require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHandleCreatePOST(t *testing.T) {
    ctrl := gomock.NewController(t)

    db := mocks.NewMockDB(ctrl)

    db.EXPECT().
        Exec("INSERT INTO products (model, company, price) VALUES (?, ?, ?)",
            "iPhone", "Apple", 1000).
        Return(nil, nil)

    logger := mocks.NewMockLogger(ctrl)
    logger.EXPECT().Info(gomock.Any()).AnyTimes()
    logger.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()

    a := app.New(db, logger)

    form := "model=iPhone&company=Apple&price=1000"
    req := httptest.NewRequest("POST", "/create", strings.NewReader(form))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    w := httptest.NewRecorder()

    a.HandleCreateForTest(w, req)

    resp := w.Result()
    require.Equal(t, http.StatusSeeOther, resp.StatusCode)
}

func TestHandleDelete(t *testing.T) {
    ctrl := gomock.NewController(t)

    db := mocks.NewMockDB(ctrl)

    db.EXPECT().
        Exec("DELETE FROM products WHERE id = ?", 5).
        Return(nil, nil)

    logger := mocks.NewMockLogger(ctrl)
    logger.EXPECT().Info(gomock.Any()).AnyTimes()
    logger.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()

    a := app.New(db, logger)

    req := httptest.NewRequest("GET", "/delete?id=5", nil)
    w := httptest.NewRecorder()

    a.HandleDeleteForTest(w, req)

    resp := w.Result()
    require.Equal(t, http.StatusSeeOther, resp.StatusCode)
}
