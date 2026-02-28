package app_test

import (
    "database/sql"
    "errors"
    "testing"

    "lab3/app"
    "lab3/app/mocks"

    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/require"
)

// Заглушка для sql.Result
type dummyResult struct{}

func (d dummyResult) LastInsertId() (int64, error) { return 0, nil }
func (d dummyResult) RowsAffected() (int64, error) { return 0, nil }

func TestRun(t *testing.T) {
    type Test struct {
        Name      string
        ExecCall  func(string, ...any) (sql.Result, error)
        ActualErr error
    }

    errDB := errors.New("db error")

    tests := []Test{
        {
            Name: "Success",
            ExecCall: func(string, ...any) (sql.Result, error) {
                return dummyResult{}, nil
            },
        },
        {
            Name: "Insert error",
            ExecCall: func(q string, args ...any) (sql.Result, error) {
                if len(q) >= 6 && q[:6] == "INSERT" {
                    return dummyResult{}, errDB
                }
                return dummyResult{}, nil
            },
            ActualErr: errDB,
        },
    }

    for _, tt := range tests {
        t.Run(tt.Name, func(t *testing.T) {
            ctrl := gomock.NewController(t)

            // --- Mock DB ---
            db := mocks.NewMockDB(ctrl)
            db.EXPECT().
                Exec(gomock.Any(), gomock.Any()).
                DoAndReturn(tt.ExecCall).
                AnyTimes()

            db.EXPECT().
                Query(gomock.Any(), gomock.Any()).
                DoAndReturn(func(q string, args ...any) (app.Rows, error) {
                    rows := mocks.NewMockRows(ctrl)

                    // --- Ожидания для Rows ---
                    rows.EXPECT().Next().Return(false).AnyTimes()
                    rows.EXPECT().Scan(gomock.Any()).Return(nil).AnyTimes()
                    rows.EXPECT().Close().Return(nil).AnyTimes()

                    return rows, nil
                }).
                AnyTimes()

            // --- Mock Logger ---
            l := mocks.NewMockLogger(ctrl)
            l.EXPECT().Info(gomock.Any()).AnyTimes()
            l.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()

            // --- Run ---
            a := app.New(db, l)
            err := a.Run()

            require.ErrorIs(t, err, tt.ActualErr)
        })
    }
}
