package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/khilmi-aminudin/simplebankv1/db/sqlc"
	"github.com/khilmi-aminudin/simplebankv1/utils"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymetricKey:    utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
