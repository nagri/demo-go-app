package jokesapp

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nagri/demo-go-app/internal/app/jokesapp/mock_jokesapp"
	"github.com/nagri/demo-go-app/internal/pkg/jokesapi"
	"go.uber.org/mock/gomock"
)

func TestMain(t *testing.T) {
	// twoPartJokeStub := &jokesapi.Joke{
	// 	Error:    false,
	// 	Category: "Misc",
	// 	Type:     "twopart",
	// 	Setup:    "Arguing with a woman is like reading a software's license agreement.",
	// 	Delivery: "In the end you ignore everything and click \"I agree\"",
	// 	ID:       71,
	// 	Safe:     false,
	// 	Lang:     "en",
	// }

	t.Run("Test if Handler is running on /ping", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/ping", nil)
		response := httptest.NewRecorder()

		Ping(response, request)

		got := response.Body.String()
		want := "pong"
		assertCorrectMessage(t, got, want)
	})

}

func TestJokeAPI(t *testing.T) {
	ctrl := gomock.NewController(t)

	mo := mock_jokesapp.NewMockJokeinterface(ctrl)

	mo.EXPECT().GetJoke().Return(&jokesapi.Joke{}, nil)
	mo.GetJoke()

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
