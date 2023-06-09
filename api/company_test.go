package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/kombolewis/kompani/db/mock"
	db "github.com/kombolewis/kompani/db/sqlc"
	"github.com/kombolewis/kompani/utils"
	"github.com/stretchr/testify/require"
)

func TestGetCompanyApi(t *testing.T) {
	company := randomCompany()

	testCases := []struct {
		name          string
		id            int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			id:   company.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCompany(gomock.Any(), gomock.Eq(company.ID)).
					Times(1).
					Return(company, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatch(t, recorder.Body, company)

			},
		},
		{
			name: "NotFound",
			id:   company.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCompany(gomock.Any(), gomock.Eq(company.ID)).
					Times(1).
					Return(db.Company{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			id:   company.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCompany(gomock.Any(), gomock.Eq(company.ID)).
					Times(1).
					Return(db.Company{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			id:   0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCompany(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/company/%d", tc.id)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomCompany() db.Company {
	return db.Company{
		ID:          utils.RandomID(),
		Name:        utils.RandomName(),
		Description: sql.NullString{String: utils.RandomDescription(), Valid: true},
		Amount:      utils.RandomAmount(),
		Registered:  true,
		Type:        utils.RandomType(),
	}
}

func requireBodyMatch(t *testing.T, body *bytes.Buffer, company db.Company) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotCompany db.Company

	err = json.Unmarshal(data, &gotCompany)
	require.NoError(t, err)
	require.Equal(t, company, gotCompany)
}
